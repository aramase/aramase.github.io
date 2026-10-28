# Release ledger

`data/ledger.yaml` drives the Kubernetes ledger on `/work/`. It is generated; don't edit it.

- `keps.yaml`: the tracked KEPs, one line each.
- `overrides.yaml`: what kep.yaml does not record (PRs, posts, talks), and exclusions for
  kep.yaml milestones that are not true.
- `tools/ledger`: the generator. Stages come from each KEP's kep.yaml `milestone`, up to its
  `stage`. A stage counts once its release has a v1.N.0 tag; a GA stage in the release in
  development is drawn as a target only while the KEP issue is on that milestone with
  `lead-opted-in`. The page derives its release columns from the data: they run from
  `first` to `last`, the oldest and newest released stage, so a release with no new stage
  changes nothing. The ledger is KEP-only; projects belong in `data/work.yaml`.

With a target, the generator also writes `inDevelopment` (the release in development) and
`checked`, the UTC date of the run that confirmed the target. The site draws targets (the
"in dev" column, the dashed GA box and "GA target" on post and talk pages) only through the
35th day after `checked`, so a monthly pull request left unmerged falls back to released
stages instead of showing a stale target. The build fails if `inDevelopment` is set without
`checked`. `-check` ignores the `checked` line.

Regenerate with `make ledger` (reads GitHub; set `GITHUB_TOKEN` if the anonymous API rate
limit runs out), or `make ledger ENHANCEMENTS=$HOME/go/src/k8s.io/enhancements`
for a local clone (read as checked out). `make ledger-test` runs the generator's tests.

`.github/workflows/ledger.yml` runs the generator on the 1st of each month and on demand,
and opens a pull request from `ledger/update` when the ledger changed, which is every month
while a target is live, since `checked` moves (manual runs on other branches only generate
it). It rebuilds that branch on each run, and fails rather than overwrite a commit someone
else pushed to it; put corrections in `overrides.yaml` instead.

To track a KEP, add its line to `keps.yaml`, its PRs, posts and talks to `overrides.yaml`,
and run `make ledger`.
