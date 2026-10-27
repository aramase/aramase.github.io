# Release ledger

`data/ledger.yaml` drives the Kubernetes ledger on `/work/`. It is generated; don't edit it.

- `keps.yaml`: the tracked KEPs, one line each.
- `overrides.yaml`: what kep.yaml does not record (PRs, posts, talks), and exclusions for
  kep.yaml milestones that are not true.
- `tools/ledger`: the generator. Stages come from each KEP's kep.yaml `milestone`, up to its
  `stage`. A stage counts once its release has a v1.N.0 tag; a GA stage in the release in
  development is drawn as a target only while the KEP issue is on that milestone with
  `lead-opted-in`. The page derives its release columns from the data.

Regenerate with `make ledger` (reads GitHub; set `GITHUB_TOKEN` if the anonymous API rate
limit runs out), or `make ledger ENHANCEMENTS=$HOME/go/src/k8s.io/enhancements`
for a local clone (read as checked out). `make ledger-test` runs the generator's tests.

`.github/workflows/ledger.yml` runs the generator on the 1st of each month and on demand,
and opens a pull request from `ledger/update` when the ledger changed (manual runs on other
branches only generate it). It rebuilds that branch on each run, and fails rather than
overwrite a commit someone else pushed to it; put corrections in `overrides.yaml` instead.

To track a KEP, add its line to `keps.yaml`, its PRs, posts and talks to `overrides.yaml`,
and run `make ledger`.
