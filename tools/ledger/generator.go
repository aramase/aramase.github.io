package main

import (
	"context"
	"fmt"
	"io"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	stageAlpha = "alpha"
	stageBeta  = "beta"
	stageGA    = "ga"

	// self is left out of a KEP's authors to give its co-authors.
	self = "aramase"

	leadOptedIn = "lead-opted-in"
)

var stages = []string{stageAlpha, stageBeta, stageGA}

// rank orders stages; kep.yaml writes GA as "stable". It is -1 for anything else.
func rank(s string) int {
	if s == "stable" {
		s = stageGA
	}
	return slices.Index(stages, s)
}

// kepYAML is the part of a kep.yaml the ledger reads.
type kepYAML struct {
	Number    string   `yaml:"kep-number"`
	Authors   []string `yaml:"authors"`
	OwningSIG string   `yaml:"owning-sig"`
	Stage     string   `yaml:"stage"`
	Milestone struct {
		Alpha  string `yaml:"alpha"`
		Beta   string `yaml:"beta"`
		Stable string `yaml:"stable"`
	} `yaml:"milestone"`
}

func (k kepYAML) milestone(stage string) string {
	switch stage {
	case stageAlpha:
		return k.Milestone.Alpha
	case stageBeta:
		return k.Milestone.Beta
	}
	return k.Milestone.Stable
}

type issueReader interface {
	issue(ctx context.Context, n int) (issue, error)
}

type generator struct {
	src    source
	issues issueReader
	// latest is the newest released Kubernetes minor; latest+1 is in development.
	latest int
	// now gives the time of the run; Checked is its UTC date. Nil means time.Now.
	now func() time.Time
	log io.Writer
}

type stage struct {
	Stage   string
	Release int
	Target  bool
	annotation
}

type entry struct {
	ID, Name  string
	KEP       int
	Path, SIG string
	Coauthors []string
	Stages    []stage
	// Notes are the reasons for exclusions, written as comments.
	Notes []string
}

type ledger struct {
	// First and Last are the releases of the oldest and newest released stage on the grid;
	// a target is not released, so it never moves Last.
	First, Last int
	// InDevelopment is the release in development while a target is drawn, else 0.
	// Checked is the UTC date of the run that confirmed the target, set only with it: the
	// site stops drawing the target once that date is old, so an unmerged update cannot
	// leave a stale "in development" on the page.
	InDevelopment int
	Checked       time.Time
	Features      []entry
	Subprojects   []entry
}

func (g generator) logf(format string, args ...any) {
	if g.log != nil {
		fmt.Fprintf(g.log, format+"\n", args...)
	}
}

func (g generator) build(ctx context.Context, keps []tracked, ov map[int]override) (ledger, error) {
	var l ledger
	for kep := range ov {
		if !slices.ContainsFunc(keps, func(t tracked) bool { return t.KEP == kep }) {
			return ledger{}, fmt.Errorf("overrides: KEP %d is not in the tracked list", kep)
		}
	}
	for _, t := range keps {
		e, err := g.entry(ctx, t, ov[t.KEP])
		if err != nil {
			return ledger{}, err
		}
		if t.Kind == kindSubproject {
			l.Subprojects = append(l.Subprojects, e)
			continue
		}
		l.Features = append(l.Features, e)
	}
	if len(l.Features) == 0 {
		return ledger{}, fmt.Errorf("no tracked KEP is a feature on the release grid")
	}
	// The axis ends at the newest released stage, not at the newest release: it only
	// moves when the work does. A target is the one stage in the release in development,
	// so the in-development release, and the date that vouches for it, exist only while
	// a target does.
	for _, e := range l.Features {
		for _, s := range e.Stages {
			if s.Target {
				l.InDevelopment = s.Release
				continue
			}
			if l.First == 0 || s.Release < l.First {
				l.First = s.Release
			}
			l.Last = max(l.Last, s.Release)
		}
	}
	if l.Last == 0 {
		return ledger{}, fmt.Errorf("no tracked KEP has a stage in a released Kubernetes version")
	}
	if l.InDevelopment > 0 {
		now := time.Now
		if g.now != nil {
			now = g.now
		}
		y, m, d := now().UTC().Date()
		l.Checked = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	}
	return l, nil
}

func (g generator) entry(ctx context.Context, t tracked, o override) (entry, error) {
	dir, raw, err := g.src.kepYAML(ctx, t.SIG, t.KEP)
	if err != nil {
		return entry{}, err
	}
	var k kepYAML
	if err := yaml.Unmarshal(raw, &k); err != nil {
		return entry{}, fmt.Errorf("KEP %d: kep.yaml: %w", t.KEP, err)
	}
	if k.Number != strconv.Itoa(t.KEP) {
		return entry{}, fmt.Errorf("KEP %d: kep.yaml kep-number is %q", t.KEP, k.Number)
	}
	if k.OwningSIG != t.SIG {
		return entry{}, fmt.Errorf("KEP %d: kep.yaml owning-sig is %q, the tracked list says %q", t.KEP, k.OwningSIG, t.SIG)
	}

	e := entry{ID: t.ID, Name: t.Name, KEP: t.KEP, Path: t.SIG + "/" + dir, SIG: sigName(t.SIG), Coauthors: []string{}}
	for _, a := range k.Authors {
		a = strings.TrimPrefix(strings.TrimSpace(a), "@")
		if !strings.EqualFold(a, self) {
			e.Coauthors = append(e.Coauthors, a)
		}
	}

	if e.Stages, e.Notes, err = g.stages(ctx, t.KEP, k, o.Exclude); err != nil {
		return entry{}, err
	}
	if len(e.Stages) == 0 {
		return entry{}, fmt.Errorf("KEP %d: kep.yaml gives no stage in a released Kubernetes version", t.KEP)
	}

	for _, name := range stages {
		a := o.stage(name)
		if a == nil {
			continue
		}
		i := slices.IndexFunc(e.Stages, func(s stage) bool { return s.Stage == name })
		if i < 0 {
			return entry{}, fmt.Errorf("overrides: KEP %d has %s facts, but the ledger has no %s stage for it", t.KEP, name, name)
		}
		e.Stages[i].annotation = *a
	}

	if t.Kind == kindSubproject {
		// A subproject shows its latest stage, with the talks from every stage.
		var talks []string
		for _, s := range e.Stages {
			if len(s.PRs) > 0 || len(s.Posts) > 0 {
				return entry{}, fmt.Errorf("overrides: KEP %d is a subproject, which carries talks only", t.KEP)
			}
			talks = append(talks, s.Talks...)
		}
		last := e.Stages[len(e.Stages)-1]
		last.annotation = annotation{Talks: talks}
		e.Stages = []stage{last}
	}
	return e, nil
}

// stages reads kep.yaml milestones into drawn stages; see the package comment for the rules.
func (g generator) stages(ctx context.Context, kep int, k kepYAML, exclude []exclusion) ([]stage, []string, error) {
	current := rank(k.Stage)
	if current < 0 {
		return nil, nil, fmt.Errorf("KEP %d: kep.yaml stage %q is not alpha, beta or stable", kep, k.Stage)
	}
	var out []stage
	var notes []string
	used := make([]bool, len(exclude))
	prev := 0
	for i, name := range stages {
		m := k.milestone(name)
		if m == "" {
			continue
		}
		rel, err := parseRelease(m)
		if err != nil {
			return nil, nil, fmt.Errorf("KEP %d: kep.yaml milestone %s: %w", kep, name, err)
		}
		if rel < prev {
			return nil, nil, fmt.Errorf("KEP %d: kep.yaml milestone %s (v1.%d) is earlier than the stage before it (v1.%d)", kep, name, rel, prev)
		}
		prev = rel
		if i > current {
			g.logf("KEP %d: skip %s v1.%d: kep.yaml stage is %s", kep, name, rel, k.Stage)
			continue
		}
		if j := slices.IndexFunc(exclude, func(x exclusion) bool { return x.Stage == name && x.Release == rel }); j >= 0 {
			used[j] = true
			notes = append(notes, exclude[j].Note)
			g.logf("KEP %d: skip %s v1.%d: excluded in overrides", kep, name, rel)
			continue
		}
		switch {
		case rel <= g.latest:
			out = append(out, stage{Stage: name, Release: rel})
		case rel == g.latest+1 && name == stageGA:
			tracked, err := g.releaseTracks(ctx, kep, rel)
			if err != nil {
				return nil, nil, err
			}
			if !tracked {
				g.logf("KEP %d: skip GA target v1.%d: the KEP issue is not on the v1.%d milestone with %s", kep, rel, rel, leadOptedIn)
				continue
			}
			out = append(out, stage{Stage: name, Release: rel, Target: true})
		default:
			g.logf("KEP %d: skip %s v1.%d: not released, and only a GA target is drawn", kep, name, rel)
		}
	}
	for j, x := range exclude {
		if !used[j] {
			g.logf("warning: KEP %d: the exclusion of %s v1.%d matches no kep.yaml milestone; remove it", kep, x.Stage, x.Release)
		}
	}
	return out, notes, nil
}

// releaseTracks reports whether release v1.<rel> tracks the KEP: the issue's milestone is
// that release and the release lead opted it in. kep.yaml alone is not enough, since it
// keeps a target after the release drops the KEP.
func (g generator) releaseTracks(ctx context.Context, kep, rel int) (bool, error) {
	is, err := g.issues.issue(ctx, kep)
	if err != nil {
		return false, fmt.Errorf("KEP %d: reading the KEP issue: %w", kep, err)
	}
	return is.Milestone == fmt.Sprintf("v1.%d", rel) && slices.Contains(is.Labels, leadOptedIn), nil
}

var releaseRE = regexp.MustCompile(`^v1\.(\d+)$`)

func parseRelease(m string) (int, error) {
	sub := releaseRE.FindStringSubmatch(strings.TrimSpace(m))
	if sub == nil {
		return 0, fmt.Errorf("%q is not a release like v1.30", m)
	}
	return strconv.Atoi(sub[1])
}

// sigName turns a SIG directory into its display name: sig-api-machinery is SIG API Machinery.
func sigName(dir string) string {
	acronyms := map[string]string{"api": "API", "cli": "CLI", "k8s": "K8s"}
	words := strings.Split(strings.TrimPrefix(dir, "sig-"), "-")
	for i, w := range words {
		if a, ok := acronyms[w]; ok {
			words[i] = a
		} else if w != "" {
			words[i] = strings.ToUpper(w[:1]) + w[1:]
		}
	}
	return "SIG " + strings.Join(words, " ")
}
