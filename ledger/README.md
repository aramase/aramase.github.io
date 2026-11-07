# Release ledger

`data/ledger.yaml` drives the Kubernetes ledger on `/work/`. It is generated; don't edit it.

- `keps.yaml`: the tracked KEPs, each with an optional `line`, the one-sentence description
  shown under its name.
- `overrides.yaml`: what kep.yaml does not record (PRs, posts, talks), and exclusions for
  kep.yaml milestones that are not true.
- `tools/ledger`: the generator. Stages come from each KEP's kep.yaml `milestone`, up to its
  `stage`. A stage counts once its release has a v1.N.0 tag; a GA stage in the release in
  development is written as a target only while the KEP issue is on that milestone with
  `lead-opted-in`. `releases` maps each release with a released stage to the commit date of
  its v1.N.0 tag in kubernetes/kubernetes, read with one shallow, tree-less fetch of those
  tags; the page dates its stages and timeline from it. The ledger is KEP-only; projects
  belong in `data/work.yaml`.

With a target, the generator also writes `inDevelopment` (the release in development) and
`checked`, the UTC date of the run that confirmed the target. The page shows a target like
any stage not reached yet ("Not yet"), so a monthly pull request left unmerged shows nothing
stale. `-check` ignores the `checked` line.

Regenerate with `make ledger` (reads GitHub and needs `git`; set `GITHUB_TOKEN` if the anonymous API rate
limit runs out), or `make ledger ENHANCEMENTS=$HOME/go/src/k8s.io/enhancements`
for a local clone (read as checked out). `make ledger-test` runs the generator's tests.

`.github/workflows/ledger.yml` runs the generator on the 1st of each month and on demand,
and opens a pull request from `ledger/update` when the ledger changed, which is every month
while a target is live, since `checked` moves (manual runs on other branches only generate
it). It rebuilds that branch on each run, and fails rather than overwrite a commit someone
else pushed to it; put corrections in `overrides.yaml` instead.

To track a KEP, add its line to `keps.yaml`, its PRs, posts and talks to `overrides.yaml`,
and run `make ledger`.
