---
version: 1
slug: "layouts-work-list-html"
primary_target: "layouts/work/list.html"
related_targets: ["layouts/partials/ledger/ledger.html","layouts/talks/list.html","layouts/blogs/list.html"]
---

# Surface brief: work page (release ledger)

Scope: `/work/` ("Work"; areas since 2026-10-27), plus the talks, posts and 404 pages. Until 2026-10-27 this was the home page; Anish moved it here so home could introduce him in plain language.
Visitor mode: Experience. Detail pages (talk, post) read as documents.

Audience: Kubernetes maintainers, SIG peers and conference organizers checking what Anish shipped.
Job: see what he works on beyond Kubernetes, which features he carried through which releases, then open the project, KEP, post or talk.
Proof: kubernetes/enhancements kep.yaml milestones, kubernetes.io release posts, talk recordings,
OWNERS, CODEOWNERS and the SIG Auth README for roles. Nothing unverified is drawn.
Constraints: only agent-primitives/sessions and aramase/kontxt under Agent infrastructure; the ledger is KEP-only, so projects and agent-platform work stay in the area entries; entries stay short (name, role, one line, link); permalinks kept.

Direction: "timing", chosen by Anish on 2026-11-05; it replaces the "Release ledger" grid. The ledger data, its rules (fold, cap, fresh targets) and its sources are unchanged; only the drawing changed.
Memorable moment: each feature as a signal stepping up from baseline to alpha, beta and GA on the release that moved it, so maturity reads as height.

Unresolved: sessions role. Shown as "Creator" (initial commit); "maintainer" rests on a GitHub team that is not public. 4412 GA is not drawn: kep.yaml lists v1.38, but v1.38 dropped the KEP; its stage list reads "beta since v1.34".

## Direction contract

THESIS: The record is the page, drawn as a timing diagram. Features are signals across a grid of Kubernetes releases that ends at Anish's newest stage; refuses the avatar, bio, card-grid personal-site default.

OWN-WORLD: An engineering datasheet. Sofia Sans text and Sofia Sans Condensed names, numerals and labels (self-hosted), tabular figures; paper and ink with one magenta signal for the trace, links, roles and focus. Stages are coded by height (baseline 280, alpha 195, beta 110, GA 25 in a 300-unit track) and by name, never by hue. A square-wave clock runs along the axis, one period per release, dashed for the release in development. Square corners, 1px and 2px rules, no shadows or cards.

STORY: The visitor reads the three areas from the "On this page" index, sees five features stepping up through releases, opens the post or talk at a step, then the SIG Auth, secrets and agent entries, then the Posts and Talks tables, where each row tags the feature it belongs to.

FIRST VIEWPORT: Header strip and clock rule. The page title "Work" with the area index. Below it the Kubernetes area: title, lede, then the timing diagram at full width: head with the key (alpha low, beta mid, GA high; GA target dashed only when drawn), the axis (at most 14 release columns, older ones folded into "≤N", "in dev" for a fresh target), and the first feature rows. Primary action: the post and talk links on each step.

FORM: Timing diagram, owner-chosen. Edge labels sit under each step, as wide as the stage runs; one-column stages run inward at the end or stack and truncate, and a tall label grows its row. The trace thickens on row hover or focus (no motion under reduced motion). Under 64rem each row stacks, edge labels hide and a stage list carries stage, release and links. Under 48rem the tables stack to rows.

FINISH: the build ends with screenshots at 1280 and 390 in both themes, the acceptance run (future talks and a fake feature and target in a /tmp copy), DESIGN.md, and the contrast table.
