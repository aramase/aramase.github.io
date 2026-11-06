---
version: 1
slug: "layouts-index-html"
primary_target: "layouts/index.html"
related_targets: ["layouts/_default/baseof.html","layouts/partials/header.html","layouts/partials/footer.html","layouts/partials/email-block.html","layouts/partials/pick.html","data/home.yaml"]
---

# Surface brief: home

Scope: `/` plus the shared chrome every page inherits (header, footer).
Visitor mode: Persuade. Approved by Anish on 2026-10-27 (Slack 1a); design "changelog" chosen on 2026-11-06.

Audience: conference organizers, peers and maintainers looking Anish up, most of whom don't know
the individual Kubernetes features.
Job: learn who he is and what he works on in plain language, pick a talk or post, get in touch.
Proof: the bio (Anish's text, verbatim), selected talks and posts already on the site, links to
the matching areas on the work page (`/work/#kubernetes`, `#secrets`, `#agents`).
Constraints: no timelines and no KEP numbers here (they live on `/work/`); bio verbatim; only
agent-primitives/sessions and kontxt named on the agent side; permalinks kept.

FIRST VIEWPORT: sticky header (avatar and name, Work, Talks and posts, Contact, theme toggle).
The hero: name as the h1, the bio, "See my work" and "Get in touch", the colour headshot beside it
(above it on phones). Below: the three focus areas, picked talks and posts (three each, from
`data/home.yaml`), the organizer kit (copy bio, download headshot with its pixel size) and the
contact section, whose email block shows the address as selectable text with Copy and
"Open in mail app".

Type and colour: Geist and Geist Mono on a cool neutral ground with one blue accent; dark mode is
a near-black product UI. DESIGN.md holds the tokens and the contrast table.
