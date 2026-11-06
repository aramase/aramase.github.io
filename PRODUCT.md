# Product

<!-- impeccable:product-schema 1 -->

Source: owner request and critique plan, 2026-10-26; home rework approved by Anish, 2026-10-27.
No interview ran; facts marked (inferred) come from those briefs and the repository, not from a
confirmed answer.

## Platform

web

## Users

- Conference organizers, peers and maintainers looking Anish up, many without knowing the individual features (owner, 2026-10-27).
- Kubernetes maintainers and SIG peers checking what he has shipped, release by release.
- Engineers arriving from a kubernetes.io release post or a talk recording.

## Product Purpose

anishram.com is Anish Ramasekar's personal site. The home page says who he is and what he works
on in plain language, with a few selected talks and posts and his contact links. The work page
(`/work/`, "Work") holds the full record by area: Kubernetes identity and security (the SIG Auth
co-chair role and the release ledger of features he carried through releases), secrets and
workload identity, and agent infrastructure, then the dated posts and talks. Success is a
visitor understanding who he is and what he works on, then reaching the talk, post, KEP or
contact they came for.

## Positioning

The person first, then the record. His work moves through Kubernetes releases (alpha, beta, GA), each
stage backed by a KEP, a release blog post or a talk. No template personal site has that shape.

## Operating Context

- Hugo static site on GitHub Pages at anishram.com, rebuilt daily for future-dated posts.
- Content lives in `content/blogs/` (each entry links to an external post) and `content/talks/`.
- Ledger facts come from kubernetes/enhancements kep.yaml and KEP READMEs, never from memory.
- `tools/ledger` generates `data/ledger.yaml` from the KEPs in `ledger/keps.yaml` and the facts in
  `ledger/overrides.yaml` (`make ledger`).

## Capabilities and Constraints

- Existing permalinks (`/blogs/<slug>/`, `/talks/<slug>/`, `/tags/`, `/categories/`) must keep working.
- No timelines or KEP numbers on the home page; those live on `/work/`.
- Light and dark themes, following the OS by default with a manual toggle (inferred: plan recommendation, incumbent behaviour).
- Agent work names only agent-primitives/sessions and aramase/kontxt, both under Agent infrastructure (owner, 2026-10-27). Nothing internal, and no other agent project, is named or linked.
- Roles appear only with a public source (OWNERS, CODEOWNERS, the SIG README, commit history), noted in `data/work.yaml`. No IETF authorship or standards-body role is claimed.

## Brand Commitments

- The home bio is Anish's text of 2026-10-27, verbatim. No invented claims, titles or roles.
- Headshot: `assets/images/anish-ramasekar.png` (supplied 2026-10-26).
- The terminal theme and typewriter effect are retired.

## Evidence on Hand

- 9 blog entries (7 kubernetes.io release posts, 2 Microsoft Open Source posts).
- 14 talks with recordings; 8 with slides; KubeCon EU 2021 talk ranked in the top 10 (CNCF transparency report, linked from the talk page).
- KEPs authored or co-authored: 2907, 3299, 3331, 4412, 5538, 5793 (kep.yaml authors).
- Roles (checked 2026-10-27): SIG Auth co-chair (kubernetes/community sig-auth README); Secrets Store CSI Driver approver (OWNERS); `*` code owner of the Azure Key Vault provider and Azure Workload Identity (CODEOWNERS); creator of agent-primitives/sessions and aramase/kontxt (initial commits).
- Absent, must not be fabricated: 2026 talks, speaker kit, testimonials.

## Product Principles

- Every claim on the page links to its source.
- Leave out what cannot be verified rather than approximate it.
- Plain language on the home page; detail (KEPs, stages, PRs) on the work page.

## Accessibility & Inclusion

- WCAG 2.1 AA contrast, one h1 per page, working skip link, reduced motion respected, usable at 390px.
