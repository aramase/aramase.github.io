package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"

	"gopkg.in/yaml.v3"
)

// tracked is one line of ledger/keps.yaml.
type tracked struct {
	KEP  int    `yaml:"kep"`
	SIG  string `yaml:"sig"`
	ID   string `yaml:"id"`
	Name string `yaml:"name"`
	Kind string `yaml:"kind"`
}

const kindSubproject = "subproject"

var (
	sigDirRE = regexp.MustCompile(`^sig-[a-z0-9-]+$`)
	idRE     = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)
)

func loadTracked(path string) ([]tracked, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var doc struct {
		KEPs []tracked `yaml:"keps"`
	}
	if err := decodeStrict(b, &doc); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	if len(doc.KEPs) == 0 {
		return nil, fmt.Errorf("%s: no KEPs listed", path)
	}
	seenKEP := map[int]bool{}
	seenID := map[string]bool{}
	for _, t := range doc.KEPs {
		switch {
		case t.KEP <= 0:
			return nil, fmt.Errorf("%s: KEP %d: kep must be a positive number", path, t.KEP)
		case !sigDirRE.MatchString(t.SIG):
			return nil, fmt.Errorf("%s: KEP %d: sig %q is not a SIG directory like sig-auth", path, t.KEP, t.SIG)
		case !idRE.MatchString(t.ID):
			return nil, fmt.Errorf("%s: KEP %d: id %q must be lowercase words joined by hyphens", path, t.KEP, t.ID)
		case t.Name == "":
			return nil, fmt.Errorf("%s: KEP %d: name is empty", path, t.KEP)
		case t.Kind != "" && t.Kind != kindSubproject:
			return nil, fmt.Errorf("%s: KEP %d: kind %q is not %q", path, t.KEP, t.Kind, kindSubproject)
		case seenKEP[t.KEP]:
			return nil, fmt.Errorf("%s: KEP %d is listed twice", path, t.KEP)
		case seenID[t.ID]:
			return nil, fmt.Errorf("%s: id %q is used twice", path, t.ID)
		}
		seenKEP[t.KEP] = true
		seenID[t.ID] = true
	}
	return doc.KEPs, nil
}

// annotation is what a stage carries beyond kep.yaml.
type annotation struct {
	PRs   []int    `yaml:"prs"`
	Posts []string `yaml:"posts"`
	Talks []string `yaml:"talks"`
}

// exclusion drops a kep.yaml milestone that is not true.
type exclusion struct {
	Stage   string `yaml:"stage"`
	Release int    `yaml:"release"`
	Note    string `yaml:"note"`
}

// override is one KEP's entry in ledger/overrides.yaml.
type override struct {
	Alpha   *annotation `yaml:"alpha"`
	Beta    *annotation `yaml:"beta"`
	GA      *annotation `yaml:"ga"`
	Exclude []exclusion `yaml:"exclude"`
}

func (o override) stage(s string) *annotation {
	switch s {
	case stageAlpha:
		return o.Alpha
	case stageBeta:
		return o.Beta
	case stageGA:
		return o.GA
	}
	return nil
}

func loadOverrides(path string) (map[int]override, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	ov := map[int]override{}
	if err := decodeStrict(b, &ov); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	for kep, o := range ov {
		for _, e := range o.Exclude {
			if rank(e.Stage) < 0 || e.Release <= 0 || e.Note == "" {
				return nil, fmt.Errorf("%s: KEP %d: an exclusion needs a stage (alpha, beta or ga), a release and a note", path, kep)
			}
		}
	}
	return ov, nil
}

// decodeStrict rejects unknown fields, so a typo in a hand-edited file fails loudly.
// An empty document leaves v unchanged.
func decodeStrict(b []byte, v any) error {
	dec := yaml.NewDecoder(bytes.NewReader(b))
	dec.KnownFields(true)
	if err := dec.Decode(v); err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	return nil
}
