---
name: anishram.com
description: Anish Ramasekar's site as a modern product changelog in Geist, with each Kubernetes feature shown as an Alpha, Beta, GA stepper and a filterable timeline.
colors:
  bg: "oklch(0.995 0.002 250)"
  surface: "oklch(0.972 0.005 250)"
  ink: "oklch(0.21 0.015 255)"
  ink-2: "oklch(0.45 0.015 255)"
  line: "oklch(0.905 0.006 255)"
  line-2: "oklch(0.80 0.008 255)"
  accent: "oklch(0.49 0.13 250)"
  accent-ink: "oklch(0.99 0.004 250)"
  accent-soft: "oklch(0.945 0.03 250)"
  bg-dark: "oklch(0.165 0.008 255)"
  surface-dark: "oklch(0.205 0.01 255)"
  ink-dark: "oklch(0.95 0.005 255)"
  ink-2-dark: "oklch(0.73 0.012 255)"
  line-dark: "oklch(0.29 0.01 255)"
  line-2-dark: "oklch(0.42 0.012 255)"
  accent-dark: "oklch(0.77 0.11 245)"
  accent-ink-dark: "oklch(0.18 0.02 255)"
  accent-soft-dark: "oklch(0.27 0.045 250)"
typography:
  hero:
    fontFamily: "Geist, system-ui, sans-serif"
    fontSize: "clamp(2.6rem, 6.4vw, 4.75rem)"
    fontWeight: 680
    lineHeight: 1
    letterSpacing: "-0.04em"
  page-title:
    fontFamily: "Geist, system-ui, sans-serif"
    fontSize: "clamp(2.5rem, 6vw, 4rem)"
    fontWeight: 680
    lineHeight: 1
    letterSpacing: "-0.04em"
  doc-title:
    fontFamily: "Geist, system-ui, sans-serif"
    fontSize: "clamp(2rem, 4.6vw, 3rem)"
    fontWeight: 680
    lineHeight: 1.08
    letterSpacing: "-0.035em"
  section-title:
    fontFamily: "Geist, system-ui, sans-serif"
    fontSize: "1.5rem"
    fontWeight: 640
    lineHeight: 1.2
    letterSpacing: "-0.025em"
  item-title:
    fontFamily: "Geist, system-ui, sans-serif"
    fontSize: "1.1875rem"
    fontWeight: 640
    lineHeight: 1.3
    letterSpacing: "-0.02em"
  body:
    fontFamily: "Geist, system-ui, sans-serif"
    fontSize: "1rem"
    fontWeight: 400
    lineHeight: 1.6
  meta:
    fontFamily: "Geist, system-ui, sans-serif"
    fontSize: "0.9375rem"
    fontWeight: 400
  numeral:
    fontFamily: "Geist Mono, ui-monospace, monospace"
    fontSize: "0.875em"
    fontWeight: 400
    letterSpacing: "-0.01em"
rounded:
  control: "10px"
  chip: "999px"
  panel: "14px"
  kit: "18px"
  photo: "28px"
spacing:
  wrap: "74rem"
  gutter: "clamp(1rem, 4vw, 2.5rem)"
  section: "clamp(2.5rem, 5vw, 4rem)"
components:
  header:
    backgroundColor: "{colors.bg}"
    textColor: "{colors.ink-2}"
  button-primary:
    backgroundColor: "{colors.accent}"
    textColor: "{colors.accent-ink}"
    rounded: "{rounded.control}"
  button:
    backgroundColor: "{colors.bg}"
    textColor: "{colors.ink}"
    rounded: "{rounded.control}"
  stepper:
    textColor: "{colors.ink-2}"
  timeline:
    textColor: "{colors.ink}"
  panel:
    backgroundColor: "{colors.surface}"
    rounded: "{rounded.panel}"
---

# Design System: anishram.com

## Overview

**North star: a modern product changelog.** Someone landing from a talk learns within seconds what Anish works on and what he has shipped, through parts every visitor already knows: a hero, clear sections, a "how it works" explainer, then the list. Owner-chosen on 2026-11-06 ("changelog", concepts round 3), replacing "timing".

- One accent hue (blue, 250) on a cool neutral ground; everything else is ink, ink-2 and lines.
- Geist for text and headings, Geist Mono for release numbers and dates. Both self-hosted woff2 with their OFL licences in `static/fonts/`.
- Stages are named in words (Alpha, Beta, GA, "Skipped", "Not yet") and drawn with three marks (ring, half disc, full disc), never told apart by hue.
- Dark mode is a near-black product UI with the accent lifted; layout does not change.

## Colors

Tokens live in `assets/css/site.css` (`:root`, the `prefers-color-scheme` block and `[data-theme="dark"]`).

| Token | Use | Light | Dark |
|---|---|---|---|
| `bg` | page ground, header | `#fcfeff` | `#0c0f12` |
| `surface` | explainer, kit, notes, nav hover | `#f3f6f9` | `#14171c` |
| `ink` | text | `#14191f` | `#eceff2` |
| `ink-2` | secondary text, dates, meta | `#50565e` | `#a3a8af` |
| `line` | section and row rules, timeline rail | `#dde0e4` | `#282c30` |
| `line-2` | control borders, underlines, "Not yet" dash | `#babec3` | `#494e54` |
| `accent` | links that act, marks, steps, focus, selection | `#1263a6` | `#75bbf6` |
| `accent-ink` | text on `accent` | `#fafcfe` | `#0c121a` |
| `accent-soft` | role chip | `#deefff` | `#14283c` |

Hexes are the sRGB rendering of the OKLCH values. `theme-color` follows the theme (`#fcfdff` light, `#0c0f12` dark).

### Contrast (WCAG 2.1 AA, sRGB rendering of the tokens in Chromium)

| Pair | Use | Need | Light | Dark |
|---|---|---|---|---|
| ink on bg | body text | 4.5 | 17.47 | 16.65 |
| ink-2 on bg | secondary text, dates | 4.5 | 7.33 | 8.03 |
| ink on surface | kit, explainer, notes | 4.5 | 16.29 | 15.56 |
| ink-2 on surface | secondary text on a panel | 4.5 | 6.83 | 7.51 |
| accent on bg | links, step links | 4.5 | 6.18 | 9.33 |
| accent-ink on accent | primary button, Copy | 4.5 | 6.08 | 9.13 |
| bg on ink | pressed filter, copied state, skip link | 4.5 | 17.47 | 16.65 |
| ink on accent-soft | role chip | 4.5 | 15.06 | 13.01 |
| accent on bg | focus ring, stage marks | 3 | 6.18 | 9.33 |

`line-2` borders (1.85 light, 2.29 dark) are decorative: every control carries a text label, so the border does not identify it.

## Typography

- Hero h1 680, clamp(2.6rem, 6.4vw, 4.75rem), -0.04em. Page h1 680, clamp(2.5rem, 6vw, 4rem). Doc h1 680, clamp(2rem, 4.6vw, 3rem).
- Section h2 640, 1.5rem, -0.025em. Feature and sub-heads 640, 1.1875rem. Area and entry names 620, about 1.0625 to 1.125rem.
- Body 400, 1rem/1.6, `ss01`. Bio clamp(1.0625rem, 1.5vw, 1.25rem). Doc prose 1.0625rem/1.7, 46rem column.
- Meta 0.9375rem in `ink-2`; small 0.875rem; chips 0.75 to 0.8125rem.
- `time` and `.v` (release numbers) in Geist Mono at 0.875em, tabular figures.

## Layout

- Container 74rem with a clamp(1rem, 4vw, 2.5rem) gutter. Sections are separated by a 1px `line` rule and clamp(2.5rem, 5vw, 4rem) padding.
- Home: hero (text, photo), three areas, talks and posts side by side, organizer kit, contact.
- Work: page head with jump links; Kubernetes (explainer, features, co-chair role); Projects (two columns); timeline.
- Lists (talks, posts, tags): a 7rem year column beside the rows.
- Breakpoints: under 60rem features, areas, project columns and the explainer stack. Under 44rem the hero photo moves above the text, the second nav item and the toggle label hide (the toggle keeps its accessible name), picks and list rows stack, the year label stops being sticky.
- No horizontal overflow at 390px.

## Shapes and depth

- Radii: 10px controls, 999px chips and pills, 14px panels, 18px kit, 28px hero photo (22px on phones), 50% avatars and nodes.
- One shadow token (`--shadow`), on the hero photo only. Everything else is flat, separated by lines and the `surface` fill.

## Components

- **Header:** sticky, `bg` with a bottom `line`. Avatar and name, nav (Work, Talks and posts, Contact; `aria-current` on the current page) and the theme toggle (`button.theme`, `aria-pressed`, label "Dark").
- **Buttons:** `.btn` (outlined), `.btn--primary` (accent), `.btn--small`; `.more` is an accent text link with an arrow.
- **Email block:** the address as selectable text, Copy (Clipboard API, then `execCommand`, then select and show "Press ⌘C/Ctrl+C"), "Open in mail app" (mailto) and a live status line.
- **Explainer:** "How a Kubernetes feature ships", a `surface` panel with the three marks and one line each.
- **Feature row:** name, line, KEP, SIG, co-authors and PRs on the left; the stepper on the right.
- **Stepper:** three steps, Alpha, Beta, GA. A reached step has a solid accent bar and mark, shows `v1.N` and the month, and links its posts and talks. A missing stage below the highest reached reads "Skipped" (thin line). A stage above it, or a GA target, reads "Not yet" (dashed bar).
- **Entry:** project or role, with a role chip (`accent-soft`); a subproject adds its KEP, stage and talk count.
- **Timeline:** year groups with a sticky year on desktop, a rail with a node per event (stage mark for releases, dot for posts and talks). Filters All, Releases, Posts, Talks are `aria-pressed` buttons with counts and a live "N entries" count. Without JS every event shows and the filters stay hidden.
- **List row (`.pick`):** date, title, venue or source with Recording and Slides links, optional summary and achievement, and tags linking to `/work/#<id>`.
- **Doc page:** back link, title, meta, action buttons, prose, and the "On the work page" note for ledger pages; posts end with a newer and older pager.

## Accessibility

- One h1 per page. The skip link is the first Tab stop and targets `main#main`.
- Focus: 2px `accent` outline at 3px offset.
- The theme toggle stores the choice in `localStorage` ("theme") and sets `aria-pressed`. Without JS the OS preference applies.
- Stage marks are `aria-hidden`; every stage and status exists as text.
- Motion is limited to 160ms hover transitions, removed under `prefers-reduced-motion: reduce`.

## Do's and Don'ts

- **Do** render steppers and the timeline from `data/ledger.yaml` and lists from content, so new KEPs, releases, talks and posts need no template edit.
- **Do** check any new colour pair against the contrast table.
- **Don't** add a second hue or tell stages apart by colour.
- **Don't** put KEP numbers or steppers on home.
- **Don't** load fonts or scripts from third parties, add analytics or a JS framework.
- **Don't** show a stage, target or role the data does not record.
