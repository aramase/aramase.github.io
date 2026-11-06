---
version: 1
slug: "layouts-index-html"
primary_target: "layouts/index.html"
related_targets: ["layouts/_default/baseof.html","layouts/partials/header.html","layouts/partials/footer.html","data/home.yaml"]
---

# Surface brief: home

Scope: `/` plus the shared chrome every page inherits (header, footer).
Visitor mode: Persuade and Experience. Approved by Anish on 2026-10-27 (Slack 1a).

Audience: conference organizers, peers and maintainers looking Anish up, most of whom don't know
the individual Kubernetes features.
Job: learn who he is and what he works on in plain language, pick a talk or post, get in touch.
Proof: the bio (Anish's text, verbatim), selected talks and posts already on the site, links to
the matching areas on the work page (`/work/#kubernetes`, `#secrets`, `#agents`).
Constraints: no timelines and no KEP numbers here (they live on `/work/`); bio verbatim; only
agent-primitives/sessions and kontxt named on the agent side; permalinks kept.

FIRST VIEWPORT: Header strip (ink, the inverse of the page): name, Work, Posts, Talks, Email,
the "Inverse" theme toggle; the clock rule under it. The name as the h1 in condensed display type,
the bio beside the headshot, then "What I work on" (Kubernetes identity and security, secrets and
workload identity, infrastructure for AI agents, one line and one link each). Below: selected talks
and posts (three each, from `data/home.yaml`), "For event organizers" (the bio again as a
select-all panel, and a headshot download with its pixel size), and contact (email; profiles in the
footer).

Type and colour: "timing", chosen by Anish on 2026-11-05 from the rendered concepts, replacing
spruce. The page is an engineering datasheet: Sofia Sans for text, Sofia Sans Condensed for names,
headings and numerals (self-hosted, OFL), a cool neutral paper and ink, and one magenta signal for
link underlines, roles, focus and selection. Dark mode is the inverse print. DESIGN.md holds the
tokens and the contrast table.
