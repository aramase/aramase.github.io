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
the work page, the Secrets Store CSI Driver and agent-primitives/sessions.
Constraints: no timelines and no KEP numbers here (they live on `/work/`); bio verbatim; only
agent-primitives/sessions named on the agent side; permalinks kept.

FIRST VIEWPORT: Top bar: name, Work, Posts, Talks, Email, theme toggle. Headshot, the name as
the h1, the bio. Below: "What I work on" (Kubernetes identity and security, secrets management,
infrastructure for AI agents, one line and one link each), selected talks and posts (three each,
from `data/home.yaml`), and contact (email and profiles).

Type and colour: spruce, picked by Anish on 2026-10-27 (Slack 1a) from three rendered directions.
Jost for names and headings, DM Sans for text, a deep spruce band full width behind the header and
hero, and saffron as the second colour for the name bar, links and focus. DESIGN.md holds the tokens.
