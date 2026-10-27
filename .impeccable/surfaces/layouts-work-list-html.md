---
version: 1
slug: "layouts-work-list-html"
primary_target: "layouts/work/list.html"
related_targets: ["layouts/partials/ledger/ledger.html","layouts/talks/list.html","layouts/blogs/list.html"]
---

# Surface brief: work page (release ledger)

Scope: `/work/` ("Features I've shipped"), plus the talks, blogs and 404 pages. Until 2026-10-27 this was the home page; Anish moved it here so home could introduce him in plain language.
Visitor mode: Experience. Detail pages (talk, post) read as documents.

Audience: Kubernetes maintainers, SIG peers and conference organizers checking what Anish shipped.
Job: see which features he carried through which releases, then open the KEP, post, talk or contact.
Proof: kubernetes/enhancements kep.yaml milestones, kubernetes.io release posts, talk recordings,
agent-primitives/sessions release tags. Nothing unverified is drawn.
Constraints: only agent-primitives/sessions on the agent track; permalinks kept.

Direction: "Release ledger", pinned by Anish on 2026-10-26 (overrides the roll).
Memorable moment: the staircase of feature lines firming from dashed alpha to solid GA across v1.25 to v1.38.

Unresolved: sessions role wording. 4412 GA is not drawn: kep.yaml lists v1.38, but v1.38 dropped the KEP.

## Direction contract

THESIS: The record is the page. Features advance across a fixed grid of Kubernetes releases; refuses the avatar, bio, card-grid personal-site default.

OWN-WORLD: Release-notes and enhancements-tracking typography on the site's spruce-tinted ground. Jost names and numerals, DM Sans text (self-hosted), tabular figures; version numerals set large and light. Three stage colours (alpha rust, beta ochre, GA teal) are the only colour inside the ledger, each also coded by line weight (dashed, thin, solid block); saffron marks links and focus as on every page (type and colour changed from Archivo and neutral greys to spruce by Anish, 2026-10-27). Hairline rules, ghost cells for empty releases, no cards, gradients, icons-as-glyphs or terminal chrome. Raised by seven-segment display: absence is drawn, every release cell exists and only stages light it. Raised by Game Boy four-shade field: total palette commitment, no fourth stage hue, focus and selection drawn from the site accent.

STORY: The visitor sees five features climbing through releases, reads who wrote each KEP, opens the post or talk for a stage, then the agent-platform track, then the dated record of posts and talks.

FIRST VIEWPORT: Top bar: name, Work, Blogs, Talks, Email, theme toggle. The page title "Features I've shipped" and one line on what the ledger shows. Below it the ledger takes full width: release axis 1.25 to 1.38 set as large tabular numerals, legend, and the first feature rows. Primary action: the stage links inside the rows.

FORM: Release ledger, owner-pinned; the roll's assigned candidate was overridden by the brief. seed:8661f46ec84f7a5a. Raised by pulp rack: hovering or focusing a feature promotes its posts and talks in the record below and quiets the rest (the signature interaction; no motion under reduced motion). Raised by akari: under 900px each feature folds flat into a stacked stage list in the same order.

FINISH: unreviewed and undocumented is unfinished; this build ends with the finish review, the verdict, DESIGN.md, and every shipping raster carrying its provenance
