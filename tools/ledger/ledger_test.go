package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"gopkg.in/yaml.v3"
)

// fakeSource serves kep.yaml text keyed by KEP number, under the directory "<kep>-test".
type fakeSource map[int]string

func (f fakeSource) kepYAML(_ context.Context, sig string, kep int) (string, []byte, error) {
	s, ok := f[kep]
	if !ok {
		return "", nil, fmt.Errorf("KEP %d: no keps/%s directory", kep, sig)
	}
	return fmt.Sprintf("%d-test", kep), []byte(s), nil
}

type fakeIssues map[int]issue

func (f fakeIssues) issue(_ context.Context, n int) (issue, error) {
	is, ok := f[n]
	if !ok {
		return issue{}, errors.New("not found")
	}
	return is, nil
}

func kep(n int, sig, stage, alpha, beta, stable string, authors ...string) string {
	var b strings.Builder
	fmt.Fprintf(&b, "title: Test\nkep-number: %d\nowning-sig: %s\nstage: %s\nauthors:\n", n, sig, stage)
	for _, a := range authors {
		fmt.Fprintf(&b, "  - \"@%s\"\n", a)
	}
	fmt.Fprintf(&b, "milestone:\n  alpha: %q\n  beta: %q\n  stable: %q\n", alpha, beta, stable)
	return b.String()
}

func TestBuild(t *testing.T) {
	src := fakeSource{
		1: kep(1, "sig-auth", "stable", "v1.25", "v1.27", "v1.29", "ritazh", "aramase"),
		2: kep(2, "sig-api-machinery", "beta", "v1.36", "v1.37", "", "aramase", "BenTheElder"),
		3: kep(3, "sig-storage", "stable", "", "v1.35", "v1.38", "aramase"),
		4: kep(4, "sig-auth", "stable", "", "", "v1.22", "tam7t", "aramase"),
	}
	keps := []tracked{
		{KEP: 1, SIG: "sig-auth", ID: "one", Name: "One"},
		{KEP: 2, SIG: "sig-api-machinery", ID: "two", Name: "Two"},
		{KEP: 3, SIG: "sig-storage", ID: "three", Name: "Three"},
		{KEP: 4, SIG: "sig-auth", ID: "four", Name: "Four", Kind: kindSubproject},
	}
	ov := map[int]override{
		1: {Alpha: &annotation{PRs: []int{10, 11}, Posts: []string{"/blogs/a"}}},
		3: {GA: &annotation{PRs: []int{30}}},
		4: {GA: &annotation{Talks: []string{"/talks/x", "/talks/y"}}},
	}
	issues := fakeIssues{3: {Milestone: "v1.38", Labels: []string{"stage/stable", leadOptedIn}}}
	// 23:30 on 2026-10-01 in UTC-7 is 2026-10-02 in UTC.
	now := time.Date(2026, 10, 1, 23, 30, 0, 0, time.FixedZone("UTC-7", -7*60*60))
	g := generator{src: src, issues: issues, latest: 37, now: func() time.Time { return now }}

	got, err := g.build(context.Background(), keps, ov)
	if err != nil {
		t.Fatal(err)
	}
	want := ledger{
		// The target is in v1.38 but not released, so the axis still ends at v1.37.
		First: 25, Last: 37, InDevelopment: 38, Checked: time.Date(2026, 10, 2, 0, 0, 0, 0, time.UTC),
		Features: []entry{
			{ID: "one", Name: "One", KEP: 1, Path: "sig-auth/1-test", SIG: "SIG Auth", Coauthors: []string{"ritazh"}, Stages: []stage{
				{Stage: stageAlpha, Release: 25, annotation: annotation{PRs: []int{10, 11}, Posts: []string{"/blogs/a"}}},
				{Stage: stageBeta, Release: 27},
				{Stage: stageGA, Release: 29},
			}},
			{ID: "two", Name: "Two", KEP: 2, Path: "sig-api-machinery/2-test", SIG: "SIG API Machinery", Coauthors: []string{"BenTheElder"}, Stages: []stage{
				{Stage: stageAlpha, Release: 36},
				{Stage: stageBeta, Release: 37},
			}},
			{ID: "three", Name: "Three", KEP: 3, Path: "sig-storage/3-test", SIG: "SIG Storage", Coauthors: []string{}, Stages: []stage{
				{Stage: stageBeta, Release: 35},
				{Stage: stageGA, Release: 38, Target: true, annotation: annotation{PRs: []int{30}}},
			}},
		},
		Subprojects: []entry{
			{ID: "four", Name: "Four", KEP: 4, Path: "sig-auth/4-test", SIG: "SIG Auth", Coauthors: []string{"tam7t"}, Stages: []stage{
				{Stage: stageGA, Release: 22, annotation: annotation{Talks: []string{"/talks/x", "/talks/y"}}},
			}},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("build:\n got %+v\nwant %+v", got, want)
	}
}

// The axis ends at the newest released stage. A release with no new stage adds no column,
// and the in-development release and checked date exist only while a target does.
func TestBuildAxisEnd(t *testing.T) {
	src := fakeSource{
		1: kep(1, "sig-auth", "stable", "v1.25", "v1.27", "v1.29", "aramase"),
		2: kep(2, "sig-api-machinery", "beta", "v1.36", "v1.37", "", "aramase"),
	}
	keps := []tracked{
		{KEP: 1, SIG: "sig-auth", ID: "one", Name: "One"},
		{KEP: 2, SIG: "sig-api-machinery", ID: "two", Name: "Two"},
	}
	for _, latest := range []int{37, 38, 41} {
		g := generator{src: src, issues: fakeIssues{}, latest: latest}
		l, err := g.build(context.Background(), keps, nil)
		if err != nil {
			t.Fatal(err)
		}
		if l.First != 25 || l.Last != 37 || l.InDevelopment != 0 || !l.Checked.IsZero() {
			t.Errorf("latest v1.%d: first %d, last %d, inDevelopment %d, checked %v; want 25, 37, 0, zero", latest, l.First, l.Last, l.InDevelopment, l.Checked)
		}
		if out := string(render(l)); strings.Contains(out, "\n  inDevelopment:") || strings.Contains(out, "\n  checked:") {
			t.Errorf("latest v1.%d: output names a release in development without a target", latest)
		}
	}
}

// A GA milestone in the release in development that is not drawn as a target, because the
// release does not track the KEP or an exclusion drops it, adds no in-development release,
// no checked date and no column.
func TestBuildNoLiveTarget(t *testing.T) {
	src := fakeSource{
		1: kep(1, "sig-auth", "stable", "v1.33", "v1.34", "v1.38", "aramase"),
	}
	keps := []tracked{{KEP: 1, SIG: "sig-auth", ID: "one", Name: "One"}}
	cases := []struct {
		name   string
		issues fakeIssues
		ov     map[int]override
	}{
		{name: "milestone without lead-opted-in", issues: fakeIssues{1: {Milestone: "v1.38"}}},
		{name: "lead-opted-in on another milestone", issues: fakeIssues{1: {Milestone: "v1.39", Labels: []string{leadOptedIn}}}},
		{name: "no milestone", issues: fakeIssues{1: {}}},
		{
			name:   "excluded",
			issues: fakeIssues{1: {Milestone: "v1.38", Labels: []string{leadOptedIn}}},
			ov:     map[int]override{1: {Exclude: []exclusion{{Stage: stageGA, Release: 38, Note: "dropped"}}}},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g := generator{src: src, issues: tc.issues, latest: 37, now: func() time.Time { return time.Date(2026, 10, 1, 9, 0, 0, 0, time.UTC) }}
			l, err := g.build(context.Background(), keps, tc.ov)
			if err != nil {
				t.Fatal(err)
			}
			if l.First != 33 || l.Last != 34 || l.InDevelopment != 0 || !l.Checked.IsZero() {
				t.Errorf("first %d, last %d, inDevelopment %d, checked %v; want 33, 34, 0, zero", l.First, l.Last, l.InDevelopment, l.Checked)
			}
		})
	}
}

// Without an injected clock the checked date is today's, in UTC.
func TestBuildCheckedDefaultsToNow(t *testing.T) {
	src := fakeSource{1: kep(1, "sig-auth", "stable", "", "v1.37", "v1.38", "aramase")}
	issues := fakeIssues{1: {Milestone: "v1.38", Labels: []string{leadOptedIn}}}
	g := generator{src: src, issues: issues, latest: 37}
	before := time.Now().UTC()
	l, err := g.build(context.Background(), []tracked{{KEP: 1, SIG: "sig-auth", ID: "one", Name: "One"}}, nil)
	if err != nil {
		t.Fatal(err)
	}
	after := time.Now().UTC()
	if l.InDevelopment != 38 || l.Last != 37 || l.Checked.Location() != time.UTC ||
		l.Checked.After(after) || l.Checked.Before(before.Truncate(24*time.Hour)) {
		t.Errorf("inDevelopment %d, last %d, checked %v; want 38, 37, today in UTC", l.InDevelopment, l.Last, l.Checked)
	}
}

func TestWithoutChecked(t *testing.T) {
	a := "kubernetes:\n  last: 37\n  inDevelopment: 38   # x\n  checked: 2026-10-01   # y\n  features:\n"
	b := "kubernetes:\n  last: 37\n  inDevelopment: 38   # x\n  checked: 2026-11-01   # y\n  features:\n"
	c := "kubernetes:\n  last: 37\n  features:\n"
	if withoutChecked([]byte(a)) != withoutChecked([]byte(b)) {
		t.Error("ledgers that differ only in the checked date compare different")
	}
	if withoutChecked([]byte(a)) == withoutChecked([]byte(c)) {
		t.Error("a ledger with a target compares equal to one without")
	}
}

func TestStages(t *testing.T) {
	note := "dropped by the release"
	cases := []struct {
		name      string
		kep       string
		exclude   []exclusion
		issues    fakeIssues
		want      []stage
		wantNotes []string
		wantLog   string
		wantErr   string
	}{
		{
			name: "released stages",
			kep:  kep(9, "sig-auth", "stable", "v1.29", "v1.30", "v1.34"),
			want: []stage{{Stage: stageAlpha, Release: 29}, {Stage: stageBeta, Release: 30}, {Stage: stageGA, Release: 34}},
		},
		{
			name: "empty milestones are skipped",
			kep:  kep(9, "sig-auth", "stable", "", "v1.35", "v1.36"),
			want: []stage{{Stage: stageBeta, Release: 35}, {Stage: stageGA, Release: 36}},
		},
		{
			name:    "a stage past kep.yaml stage is a plan",
			kep:     kep(9, "sig-auth", "beta", "v1.30", "v1.31", "v1.33"),
			want:    []stage{{Stage: stageAlpha, Release: 30}, {Stage: stageBeta, Release: 31}},
			wantLog: "skip ga v1.33: kep.yaml stage is beta",
		},
		{
			name:   "GA target the release tracks",
			kep:    kep(9, "sig-auth", "stable", "v1.33", "v1.34", "v1.38"),
			issues: fakeIssues{9: {Milestone: "v1.38", Labels: []string{leadOptedIn}}},
			want:   []stage{{Stage: stageAlpha, Release: 33}, {Stage: stageBeta, Release: 34}, {Stage: stageGA, Release: 38, Target: true}},
		},
		{
			name:    "GA target without lead-opted-in",
			kep:     kep(9, "sig-auth", "stable", "v1.33", "v1.34", "v1.38"),
			issues:  fakeIssues{9: {Milestone: "v1.38", Labels: []string{"stage/stable"}}},
			want:    []stage{{Stage: stageAlpha, Release: 33}, {Stage: stageBeta, Release: 34}},
			wantLog: "skip GA target v1.38",
		},
		{
			name:    "GA target the release dropped",
			kep:     kep(9, "sig-auth", "stable", "v1.33", "v1.34", "v1.38"),
			issues:  fakeIssues{9: {Labels: []string{leadOptedIn}}},
			want:    []stage{{Stage: stageAlpha, Release: 33}, {Stage: stageBeta, Release: 34}},
			wantLog: "skip GA target v1.38",
		},
		{
			name:    "the KEP issue cannot be read",
			kep:     kep(9, "sig-auth", "stable", "v1.33", "v1.34", "v1.38"),
			issues:  fakeIssues{},
			wantErr: "reading the KEP issue",
		},
		{
			name:    "beta in the release in development is not drawn",
			kep:     kep(9, "sig-auth", "beta", "v1.37", "v1.38", ""),
			want:    []stage{{Stage: stageAlpha, Release: 37}},
			wantLog: "skip beta v1.38: not released",
		},
		{
			name:    "a later release is not drawn",
			kep:     kep(9, "sig-auth", "stable", "v1.36", "v1.37", "v1.39"),
			want:    []stage{{Stage: stageAlpha, Release: 36}, {Stage: stageBeta, Release: 37}},
			wantLog: "skip ga v1.39: not released",
		},
		{
			name:      "exclusion",
			kep:       kep(9, "sig-auth", "stable", "v1.33", "v1.34", "v1.38"),
			exclude:   []exclusion{{Stage: stageGA, Release: 38, Note: note}},
			want:      []stage{{Stage: stageAlpha, Release: 33}, {Stage: stageBeta, Release: 34}},
			wantNotes: []string{note},
			wantLog:   "skip ga v1.38: excluded in overrides",
		},
		{
			name:    "an exclusion that matches nothing is reported",
			kep:     kep(9, "sig-auth", "stable", "v1.33", "v1.34", "v1.36"),
			exclude: []exclusion{{Stage: stageGA, Release: 38, Note: note}},
			want:    []stage{{Stage: stageAlpha, Release: 33}, {Stage: stageBeta, Release: 34}, {Stage: stageGA, Release: 36}},
			wantLog: "warning: KEP 9: the exclusion of ga v1.38 matches no kep.yaml milestone",
		},
		{
			name:    "malformed milestone",
			kep:     kep(9, "sig-auth", "stable", "1.30", "", "v1.32"),
			wantErr: `milestone alpha: "1.30" is not a release like v1.30`,
		},
		{
			name:    "milestones out of order",
			kep:     kep(9, "sig-auth", "stable", "v1.30", "v1.29", "v1.32"),
			wantErr: "milestone beta (v1.29) is earlier than the stage before it (v1.30)",
		},
		{
			name:    "unknown kep.yaml stage",
			kep:     kep(9, "sig-auth", "", "v1.30", "", ""),
			wantErr: `stage "" is not alpha, beta or stable`,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var log bytes.Buffer
			g := generator{issues: tc.issues, latest: 37, log: &log}
			var k kepYAML
			if err := yaml.Unmarshal([]byte(tc.kep), &k); err != nil {
				t.Fatal(err)
			}
			got, notes, err := g.stages(context.Background(), 9, k, tc.exclude)
			if tc.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
					t.Fatalf("error = %v, want it to contain %q", err, tc.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("stages = %+v, want %+v", got, tc.want)
			}
			if !reflect.DeepEqual(notes, tc.wantNotes) {
				t.Errorf("notes = %q, want %q", notes, tc.wantNotes)
			}
			if !strings.Contains(log.String(), tc.wantLog) {
				t.Errorf("log = %q, want it to contain %q", log.String(), tc.wantLog)
			}
		})
	}
}

func TestBuildErrors(t *testing.T) {
	ok := kep(1, "sig-auth", "beta", "v1.30", "v1.31", "")
	cases := []struct {
		name    string
		src     fakeSource
		keps    []tracked
		ov      map[int]override
		wantErr string
	}{
		{
			name:    "kep-number mismatch",
			src:     fakeSource{1: kep(2, "sig-auth", "beta", "v1.30", "", "")},
			keps:    []tracked{{KEP: 1, SIG: "sig-auth", ID: "a", Name: "A"}},
			wantErr: `kep.yaml kep-number is "2"`,
		},
		{
			name:    "owning SIG mismatch",
			src:     fakeSource{1: kep(1, "sig-node", "beta", "v1.30", "", "")},
			keps:    []tracked{{KEP: 1, SIG: "sig-auth", ID: "a", Name: "A"}},
			wantErr: `owning-sig is "sig-node", the tracked list says "sig-auth"`,
		},
		{
			name:    "missing KEP directory",
			src:     fakeSource{},
			keps:    []tracked{{KEP: 1, SIG: "sig-auth", ID: "a", Name: "A"}},
			wantErr: "KEP 1: no keps/sig-auth directory",
		},
		{
			name:    "no released stage",
			src:     fakeSource{1: kep(1, "sig-auth", "alpha", "v1.39", "", "")},
			keps:    []tracked{{KEP: 1, SIG: "sig-auth", ID: "a", Name: "A"}},
			wantErr: "no stage in a released Kubernetes version",
		},
		{
			name:    "facts for a stage the ledger does not draw",
			src:     fakeSource{1: ok},
			keps:    []tracked{{KEP: 1, SIG: "sig-auth", ID: "a", Name: "A"}},
			ov:      map[int]override{1: {GA: &annotation{PRs: []int{1}}}},
			wantErr: "KEP 1 has ga facts, but the ledger has no ga stage",
		},
		{
			name:    "overrides for an untracked KEP",
			src:     fakeSource{1: ok},
			keps:    []tracked{{KEP: 1, SIG: "sig-auth", ID: "a", Name: "A"}},
			ov:      map[int]override{2: {}},
			wantErr: "KEP 2 is not in the tracked list",
		},
		{
			name:    "subproject with PRs",
			src:     fakeSource{1: ok, 2: kep(2, "sig-auth", "stable", "", "", "v1.22")},
			keps:    []tracked{{KEP: 1, SIG: "sig-auth", ID: "a", Name: "A"}, {KEP: 2, SIG: "sig-auth", ID: "b", Name: "B", Kind: kindSubproject}},
			ov:      map[int]override{2: {GA: &annotation{PRs: []int{1}}}},
			wantErr: "subproject, which carries talks only",
		},
		{
			name:    "only subprojects",
			src:     fakeSource{2: kep(2, "sig-auth", "stable", "", "", "v1.22")},
			keps:    []tracked{{KEP: 2, SIG: "sig-auth", ID: "b", Name: "B", Kind: kindSubproject}},
			wantErr: "no tracked KEP is a feature",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g := generator{src: tc.src, issues: fakeIssues{}, latest: 37}
			_, err := g.build(context.Background(), tc.keps, tc.ov)
			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("error = %v, want it to contain %q", err, tc.wantErr)
			}
		})
	}
}

func TestKEPNumberQuoted(t *testing.T) {
	// kep.yaml writes kep-number both as a number and as a string.
	src := fakeSource{3331: strings.Replace(kep(3331, "sig-auth", "beta", "v1.29", "v1.30", ""), "kep-number: 3331", `kep-number: "3331"`, 1)}
	g := generator{src: src, issues: fakeIssues{}, latest: 37}
	if _, err := g.build(context.Background(), []tracked{{KEP: 3331, SIG: "sig-auth", ID: "a", Name: "A"}}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestRender(t *testing.T) {
	l := ledger{
		First: 25, Last: 37, InDevelopment: 38, Checked: time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC),
		Features: []entry{
			{ID: "kms-v2", Name: "KMS v2", KEP: 3299, Path: "sig-auth/3299-kms-v2-improvements", SIG: "SIG Auth", Coauthors: []string{"ritazh"}, Stages: []stage{
				{Stage: stageAlpha, Release: 25, annotation: annotation{PRs: []int{111126}, Posts: []string{"/blogs/a"}}},
				{Stage: stageGA, Release: 29, annotation: annotation{Posts: []string{"/blogs/b", "/blogs/c"}}},
			}},
			{ID: "t", Name: "No: plain", KEP: 4412, Path: "sig-auth/4412-x", SIG: "SIG Auth", Coauthors: []string{}, Stages: []stage{
				{Stage: stageGA, Release: 38, Target: true},
			}, Notes: []string{"No GA target: kep.yaml still lists stable: v1.38, but v1.38 dropped the KEP at its readiness deadline."}},
		},
		Subprojects: []entry{
			{ID: "sp", Name: "Secrets Store CSI Driver", KEP: 2907, Path: "sig-auth/2907-x", SIG: "SIG Auth", Coauthors: []string{"tam7t"}, Stages: []stage{
				{Stage: stageGA, Release: 22, annotation: annotation{Talks: []string{"/talks/x"}}},
			}},
		},
	}
	got := string(render(l))
	want := header + `
kubernetes:
  first: 25   # v1.25, the first release with a stage on the ledger
  last: 37    # v1.37, the newest released stage on the ledger
  inDevelopment: 38   # v1.38, in development: a target is drawn
  checked: 2026-10-01   # the target was confirmed on this date
  features:
    - id: kms-v2
      name: KMS v2
      kep: 3299
      path: sig-auth/3299-kms-v2-improvements
      sig: SIG Auth
      coauthors: ["ritazh"]
      stages:
        - stage: alpha
          release: 25
          prs: [111126]
          posts: ["/blogs/a"]
        - stage: ga
          release: 29
          posts:
            - "/blogs/b"
            - "/blogs/c"

    - id: t
      name: "No: plain"
      kep: 4412
      path: sig-auth/4412-x
      sig: SIG Auth
      coauthors: []
      stages:
        - stage: ga
          release: 38
          target: true
        # No GA target: kep.yaml still lists stable: v1.38, but v1.38 dropped the KEP
        # at its readiness deadline.

# Out-of-tree projects with a KEP, shown at the latest stage kep.yaml records.
subprojects:
  - id: sp
    name: Secrets Store CSI Driver
    kep: 2907
    path: sig-auth/2907-x
    sig: SIG Auth
    coauthors: ["tam7t"]
    stage: ga
    release: 22
    talks: ["/talks/x"]
`
	if got != want {
		t.Errorf("render:\n%s\nwant:\n%s", got, want)
	}

	// The output is YAML Hugo can read, with the values the template expects.
	var doc struct {
		Kubernetes struct {
			First, Last   int
			InDevelopment int       `yaml:"inDevelopment"`
			Checked       time.Time `yaml:"checked"`
			Features      []struct {
				Name      string
				Coauthors []string
				Stages    []struct {
					Stage   string
					Release int
					Target  bool
					PRs     []int `yaml:"prs"`
					Posts   []string
				}
			}
		}
		Subprojects []struct {
			Stage string
			Talks []string
		}
	}
	if err := yaml.Unmarshal([]byte(got), &doc); err != nil {
		t.Fatal(err)
	}
	k := doc.Kubernetes
	if k.First != 25 || k.Last != 37 || k.InDevelopment != 38 || !k.Checked.Equal(l.Checked) || len(k.Features) != 2 ||
		k.Features[1].Name != "No: plain" || !k.Features[1].Stages[0].Target ||
		!reflect.DeepEqual(k.Features[0].Stages[1].Posts, []string{"/blogs/b", "/blogs/c"}) ||
		len(doc.Subprojects) != 1 || doc.Subprojects[0].Talks[0] != "/talks/x" {
		t.Errorf("parsed output = %+v", doc)
	}
}

func TestScalar(t *testing.T) {
	for in, want := range map[string]string{
		"KMS v2":       "KMS v2",
		"SIG Auth":     "SIG Auth",
		"sig-auth/1-x": "sig-auth/1-x",
		"yes":          `"yes"`,
		"a: b":         `"a: b"`,
		"#x":           `"#x"`,
		"1.30":         `"1.30"`,
	} {
		if got := scalar(in); got != want {
			t.Errorf("scalar(%q) = %s, want %s", in, got, want)
		}
	}
}

func TestSigName(t *testing.T) {
	for in, want := range map[string]string{
		"sig-auth":          "SIG Auth",
		"sig-storage":       "SIG Storage",
		"sig-api-machinery": "SIG API Machinery",
		"sig-cli":           "SIG CLI",
	} {
		if got := sigName(in); got != want {
			t.Errorf("sigName(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestParseLatestRelease(t *testing.T) {
	out := strings.Join([]string{
		"a\trefs/tags/v1.9.0",
		"b\trefs/tags/v1.37.0",
		"c\trefs/tags/v1.37.1",
		"d\trefs/tags/v1.38.0-alpha.2",
		"e\trefs/tags/v1.38.0-rc.0",
		"",
	}, "\n")
	got, err := parseLatestRelease(out)
	if err != nil || got != 37 {
		t.Errorf("parseLatestRelease = %d, %v; want 37", got, err)
	}
	if _, err := parseLatestRelease("x\trefs/tags/v1.38.0-alpha.0\n"); err == nil {
		t.Error("want an error when there is no v1.N.0 tag")
	}
}

func TestLoadTracked(t *testing.T) {
	cases := map[string]string{
		"no KEPs listed":                "keps: []\n",
		"is not a SIG directory":        "keps:\n  - {kep: 1, sig: auth, id: a, name: A}\n",
		"must be lowercase words":       "keps:\n  - {kep: 1, sig: sig-auth, id: A_b, name: A}\n",
		"name is empty":                 "keps:\n  - {kep: 1, sig: sig-auth, id: a}\n",
		"kind \"x\" is not":             "keps:\n  - {kep: 1, sig: sig-auth, id: a, name: A, kind: x}\n",
		"KEP 1 is listed twice":         "keps:\n  - {kep: 1, sig: sig-auth, id: a, name: A}\n  - {kep: 1, sig: sig-auth, id: b, name: B}\n",
		"id \"a\" is used twice":        "keps:\n  - {kep: 1, sig: sig-auth, id: a, name: A}\n  - {kep: 2, sig: sig-auth, id: a, name: B}\n",
		"field nmae not found":          "keps:\n  - {kep: 1, sig: sig-auth, id: a, nmae: A}\n",
		"kep must be a positive number": "keps:\n  - {sig: sig-auth, id: a, name: A}\n",
	}
	dir := t.TempDir()
	for want, body := range cases {
		path := filepath.Join(dir, "keps.yaml")
		if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
		if _, err := loadTracked(path); err == nil || !strings.Contains(err.Error(), want) {
			t.Errorf("%q: error = %v, want it to contain %q", body, err, want)
		}
	}
}

func TestLoadOverrides(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "overrides.yaml")
	write := func(s string) {
		if err := os.WriteFile(path, []byte(s), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	write("# nothing yet\n")
	if ov, err := loadOverrides(path); err != nil || len(ov) != 0 {
		t.Errorf("empty file: %v, %v", ov, err)
	}

	write("1:\n  beta:\n    prs: [5]\n  exclude:\n    - {stage: ga, release: 38, note: dropped}\n")
	ov, err := loadOverrides(path)
	if err != nil {
		t.Fatal(err)
	}
	if got := ov[1]; got.Beta == nil || got.Beta.PRs[0] != 5 || got.Exclude[0].Release != 38 {
		t.Errorf("overrides = %+v", got)
	}

	write("1:\n  stable:\n    prs: [5]\n")
	if _, err := loadOverrides(path); err == nil || !strings.Contains(err.Error(), "field stable not found") {
		t.Errorf("unknown stage key: error = %v", err)
	}

	write("1:\n  exclude:\n    - {stage: ga, release: 38}\n")
	if _, err := loadOverrides(path); err == nil || !strings.Contains(err.Error(), "needs a stage") {
		t.Errorf("exclusion without a note: error = %v", err)
	}
}

func TestLocalSource(t *testing.T) {
	root := t.TempDir()
	for _, d := range []string{"keps/sig-auth/3299-kms-v2", "keps/sig-auth/32990-other", "keps/sig-auth/12-a", "keps/sig-auth/12-b"} {
		if err := os.MkdirAll(filepath.Join(root, d), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(root, "keps/sig-auth/3299-kms-v2/kep.yaml"), []byte("kep-number: 3299\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	src := localSource{root: root}

	dir, b, err := src.kepYAML(context.Background(), "sig-auth", 3299)
	if err != nil || dir != "3299-kms-v2" || string(b) != "kep-number: 3299\n" {
		t.Errorf("kepYAML(3299) = %q, %q, %v", dir, b, err)
	}
	if _, _, err := src.kepYAML(context.Background(), "sig-auth", 12); err == nil || !strings.Contains(err.Error(), "more than one") {
		t.Errorf("kepYAML(12): error = %v", err)
	}
	if _, _, err := src.kepYAML(context.Background(), "sig-auth", 7); err == nil || !strings.Contains(err.Error(), "no keps/sig-auth/7-* directory") {
		t.Errorf("kepYAML(7): error = %v", err)
	}
	if _, _, err := src.kepYAML(context.Background(), "sig-node", 1); err == nil {
		t.Error("kepYAML in a missing SIG directory: want an error")
	}
}

func TestRemoteSource(t *testing.T) {
	var listCalls int
	mux := http.NewServeMux()
	mux.HandleFunc("/repos/kubernetes/enhancements/contents/keps/sig-auth", func(w http.ResponseWriter, r *http.Request) {
		listCalls++
		if r.Header.Get("Authorization") != "Bearer tok" {
			t.Errorf("API request without the token")
		}
		if r.URL.Query().Get("ref") != "main" {
			t.Errorf("ref = %q", r.URL.Query().Get("ref"))
		}
		fmt.Fprint(w, `[{"name":"README.md","type":"file"},{"name":"3299-kms-v2","type":"dir"},{"name":"3331-structured","type":"dir"}]`)
	})
	mux.HandleFunc("/raw/kubernetes/enhancements/main/keps/sig-auth/", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "" {
			t.Errorf("the token was sent outside the API: %s", r.URL)
		}
		fmt.Fprintf(w, "path: %s\n", r.URL.Path)
	})
	mux.HandleFunc("/repos/kubernetes/enhancements/issues/4412", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"milestone":{"title":"v1.38"},"labels":[{"name":"stage/stable"},{"name":"lead-opted-in"}]}`)
	})
	mux.HandleFunc("/repos/kubernetes/enhancements/issues/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"milestone":null,"labels":[]}`)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	gh := &gitHub{client: srv.Client(), api: srv.URL, raw: srv.URL + "/raw", token: "tok"}
	src := &remoteSource{gh: gh, ref: "main"}
	ctx := context.Background()

	for _, n := range []int{3299, 3331} {
		dir, b, err := src.kepYAML(ctx, "sig-auth", n)
		if err != nil {
			t.Fatal(err)
		}
		if want := fmt.Sprintf("path: /raw/kubernetes/enhancements/main/keps/sig-auth/%s/kep.yaml\n", dir); string(b) != want {
			t.Errorf("kep.yaml = %q, want %q", b, want)
		}
	}
	if listCalls != 1 {
		t.Errorf("listed keps/sig-auth %d times, want once", listCalls)
	}
	if _, _, err := src.kepYAML(ctx, "sig-node", 1); err == nil || !strings.Contains(err.Error(), "404") {
		t.Errorf("missing SIG directory: error = %v", err)
	}

	is, err := gh.issue(ctx, 4412)
	if err != nil || is.Milestone != "v1.38" || !reflect.DeepEqual(is.Labels, []string{"stage/stable", leadOptedIn}) {
		t.Errorf("issue(4412) = %+v, %v", is, err)
	}
	if is, err := gh.issue(ctx, 1); err != nil || is.Milestone != "" || len(is.Labels) != 0 {
		t.Errorf("issue(1) = %+v, %v", is, err)
	}
	if _, err := gh.issue(ctx, 2); err == nil {
		t.Error("issue(2): want an error for a 404")
	}
}

func TestGitHubErrors(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, `{"message":"API rate limit exceeded"}`)
	}))
	defer srv.Close()
	ctx := context.Background()

	anon := &gitHub{client: srv.Client(), api: srv.URL, raw: srv.URL}
	_, err := anon.issue(ctx, 1)
	if err == nil || !strings.Contains(err.Error(), "403 Forbidden: API rate limit exceeded (set GITHUB_TOKEN") {
		t.Errorf("anonymous API error = %v", err)
	}
	if _, err := anon.get(ctx, srv.URL+"/raw", false); err == nil || strings.Contains(err.Error(), "GITHUB_TOKEN") {
		t.Errorf("raw error = %v, want no token hint", err)
	}
	authed := &gitHub{client: srv.Client(), api: srv.URL, raw: srv.URL, token: "tok"}
	if _, err := authed.issue(ctx, 1); err == nil || strings.Contains(err.Error(), "set GITHUB_TOKEN") {
		t.Errorf("authenticated API error = %v, want no token hint", err)
	}
}

func TestWrap(t *testing.T) {
	got := wrap("one two three four", 9)
	want := []string{"one two", "three", "four"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrap = %q, want %q", got, want)
	}
}
