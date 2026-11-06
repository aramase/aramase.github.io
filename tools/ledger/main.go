// Command ledger regenerates data/ledger.yaml, the release ledger on /work/, from the
// kep.yaml of each KEP tracked in ledger/keps.yaml, merged with the hand-maintained
// facts in ledger/overrides.yaml.
//
// Stages come from kep.yaml `milestone`, bounded by kep.yaml `stage`:
//   - a stage in a released Kubernetes version is drawn as shipped;
//   - a GA stage in the release in development is drawn as a target only while the KEP
//     issue's milestone is that release and it carries lead-opted-in;
//   - anything later, or past the KEP's current stage, is a plan and is left out.
//
// The output carries no source commits, and a date only while a target is drawn, so it
// changes when the ledger does, and once a run while a target is live.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "ledger:", err)
		os.Exit(1)
	}
}

func run() error {
	root := flag.String("root", ".", "site repository root; the paths below are relative to it")
	keps := flag.String("keps", "ledger/keps.yaml", "tracked KEP list")
	overrides := flag.String("overrides", "ledger/overrides.yaml", "hand-maintained facts merged into the output")
	out := flag.String("out", "data/ledger.yaml", "output file")
	enhancements := flag.String("enhancements", "", "local kubernetes/enhancements clone; empty reads kubernetes/enhancements on GitHub")
	ref := flag.String("ref", "master", "kubernetes/enhancements ref to read when -enhancements is empty")
	latest := flag.Int("latest-release", 0, "latest released Kubernetes minor (37 for v1.37); 0 reads the kubernetes/kubernetes tags")
	check := flag.Bool("check", false, "exit non-zero if the output file is not up to date, ignoring the checked date, without writing it")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	gh := newGitHub(os.Getenv("GITHUB_TOKEN"))
	var src source = &remoteSource{gh: gh, ref: *ref}
	if *enhancements != "" {
		src = localSource{root: *enhancements}
	}

	tracked, err := loadTracked(filepath.Join(*root, *keps))
	if err != nil {
		return err
	}
	ov, err := loadOverrides(filepath.Join(*root, *overrides))
	if err != nil {
		return err
	}
	if *latest == 0 {
		if *latest, err = latestRelease(ctx); err != nil {
			return err
		}
	}

	g := generator{src: src, issues: gh, latest: *latest, log: os.Stderr}
	l, err := g.build(ctx, tracked, ov)
	if err != nil {
		return err
	}
	data := render(l)

	path := filepath.Join(*root, *out)
	if *check {
		cur, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if withoutChecked(cur) != withoutChecked(data) {
			return fmt.Errorf("%s is out of date; run make ledger", *out)
		}
		return nil
	}
	return os.WriteFile(path, data, 0o644)
}

// checkedRE matches the line render writes for the checked date.
var checkedRE = regexp.MustCompile(`(?m)^  checked: .*\n`)

// withoutChecked drops the checked date, which changes on every run while a target is
// drawn, so -check compares only what the ledger says.
func withoutChecked(b []byte) string {
	return checkedRE.ReplaceAllString(string(b), "")
}
