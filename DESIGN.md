---
name: anishram.com
description: Anish Ramasekar's site set as an engineering datasheet in Sofia Sans, with his Kubernetes work drawn as a timing diagram.
colors:
  paper: "oklch(0.99 0.002 260)"
  ink: "oklch(0.19 0.012 260)"
  ink-2: "oklch(0.43 0.012 260)"
  rule: "oklch(0.78 0.006 260)"
  grid: "oklch(0.915 0.004 260)"
  sig: "oklch(0.50 0.20 350)"
  sig-fill: "oklch(0.50 0.20 350 / 0.085)"
  sig-fill-2: "oklch(0.50 0.20 350 / 0.16)"
  paper-dark: "oklch(0.18 0.01 260)"
  ink-dark: "oklch(0.95 0.004 260)"
  ink-2-dark: "oklch(0.77 0.01 260)"
  rule-dark: "oklch(0.42 0.01 260)"
  grid-dark: "oklch(0.255 0.01 260)"
  sig-dark: "oklch(0.78 0.14 350)"
  sig-fill-dark: "oklch(0.78 0.14 350 / 0.11)"
  sig-fill-2-dark: "oklch(0.78 0.14 350 / 0.2)"
typography:
  display:
    fontFamily: "Sofia Sans Condensed, Sofia Sans, system-ui, sans-serif"
    fontSize: "clamp(3.5rem, 10vw, 6rem)"
    fontWeight: 800
    lineHeight: 0.9
    letterSpacing: "-0.02em"
  page-title:
    fontFamily: "Sofia Sans Condensed, Sofia Sans, system-ui, sans-serif"
    fontSize: "clamp(3rem, 7vw, 5rem)"
    fontWeight: 800
    lineHeight: 0.9
    letterSpacing: "-0.02em"
  doc-title:
    fontFamily: "Sofia Sans Condensed, Sofia Sans, system-ui, sans-serif"
    fontSize: "clamp(2.25rem, 5vw, 3.5rem)"
    fontWeight: 800
    lineHeight: 1
    letterSpacing: "-0.015em"
  contact:
    fontFamily: "Sofia Sans Condensed, Sofia Sans, system-ui, sans-serif"
    fontSize: "clamp(1.75rem, 5vw, 3.25rem)"
    fontWeight: 800
    lineHeight: 1.1
  diagram-title:
    fontFamily: "Sofia Sans Condensed, Sofia Sans, system-ui, sans-serif"
    fontSize: "clamp(1.375rem, 2.4vw, 1.75rem)"
    fontWeight: 800
  section-title:
    fontFamily: "Sofia Sans Condensed, Sofia Sans, system-ui, sans-serif"
    fontSize: "1.625rem"
    fontWeight: 800
    letterSpacing: "0.01em"
  title:
    fontFamily: "Sofia Sans Condensed, Sofia Sans, system-ui, sans-serif"
    fontSize: "1.25rem"
    fontWeight: 700
  feature-name:
    fontFamily: "Sofia Sans Condensed, Sofia Sans, system-ui, sans-serif"
    fontSize: "1.125rem"
    fontWeight: 700
    lineHeight: 1.2
  wordmark:
    fontFamily: "Sofia Sans Condensed, Sofia Sans, system-ui, sans-serif"
    fontSize: "1.125rem"
    fontWeight: 700
    lineHeight: 1
    letterSpacing: "0.01em"
  bio:
    fontFamily: "Sofia Sans, system-ui, sans-serif"
    fontSize: "clamp(1.125rem, 1.6vw, 1.25rem)"
    fontWeight: 400
    lineHeight: 1.55
  body:
    fontFamily: "Sofia Sans, system-ui, sans-serif"
    fontSize: "1.0625rem"
    fontWeight: 400
    lineHeight: 1.55
    fontFeature: "tnum"
  table:
    fontFamily: "Sofia Sans, system-ui, sans-serif"
    fontSize: "1rem"
    fontWeight: 400
    lineHeight: 1.55
  meta:
    fontFamily: "Sofia Sans, system-ui, sans-serif"
    fontSize: "0.9375rem"
    fontWeight: 400
  axis-numeral:
    fontFamily: "Sofia Sans Condensed, Sofia Sans, system-ui, sans-serif"
    fontSize: "0.9375rem"
    fontWeight: 700
    lineHeight: 1
    fontFeature: "tnum"
  label:
    fontFamily: "Sofia Sans Condensed, Sofia Sans, system-ui, sans-serif"
    fontSize: "0.875rem"
    fontWeight: 700
    lineHeight: 1.2
  small:
    fontFamily: "Sofia Sans, system-ui, sans-serif"
    fontSize: "0.8125rem"
    fontWeight: 400
    lineHeight: 1.3
  micro:
    fontFamily: "Sofia Sans Condensed, Sofia Sans, system-ui, sans-serif"
    fontSize: "0.75rem"
    fontWeight: 600
    lineHeight: 1
rounded:
  none: "0px"
spacing:
  space-1: "0.25rem"
  space-2: "0.5rem"
  space-3: "0.75rem"
  space-4: "1rem"
  space-5: "1.25rem"
  space-6: "1.5rem"
  space-7: "2.5rem"
  section: "clamp(3rem, 6vw, 5rem)"
  gutter: "clamp(1rem, 3.5vw, 3rem)"
  page: "84rem"
  label-col: "15rem"
  doc: "46rem"
  measure: "70ch"
components:
  strip:
    backgroundColor: "{colors.ink}"
    textColor: "{colors.paper}"
    height: "3.5rem"
  theme-toggle:
    textColor: "{colors.paper}"
    typography: "{typography.meta}"
    rounded: "{rounded.none}"
    height: "44px"
    padding: "0.55rem 0.75rem"
  skip-link:
    backgroundColor: "{colors.paper}"
    textColor: "{colors.ink}"
    rounded: "{rounded.none}"
    padding: "0.6rem 0.9rem"
  clock-rule:
    backgroundColor: "{colors.sig}"
    height: "12px"
  section-head:
    textColor: "{colors.ink}"
    typography: "{typography.section-title}"
    padding: "0 0 0.5rem"
  diagram:
    backgroundColor: "{colors.paper}"
    textColor: "{colors.ink}"
    rounded: "{rounded.none}"
  trace:
    backgroundColor: "{colors.sig-fill}"
    textColor: "{colors.sig}"
  edge-label:
    textColor: "{colors.ink}"
    typography: "{typography.small}"
  feature-tag:
    textColor: "{colors.ink}"
    typography: "{typography.small}"
    rounded: "{rounded.none}"
    padding: "0.25rem 0.5rem"
  table-head:
    textColor: "{colors.ink-2}"
    typography: "{typography.label}"
    padding: "0.5rem 0.75rem 0.5rem 0"
  table-cell:
    textColor: "{colors.ink}"
    typography: "{typography.table}"
    padding: "0.75rem 0.75rem 0.75rem 0"
  speaker-bio:
    backgroundColor: "{colors.sig-fill}"
    textColor: "{colors.ink}"
    typography: "{typography.body}"
    rounded: "{rounded.none}"
    padding: "1rem 1.25rem"
  headshot:
    backgroundColor: "{colors.paper}"
    rounded: "{rounded.none}"
    padding: "6px"
    width: "10rem"
  role:
    textColor: "{colors.sig}"
    typography: "{typography.meta}"
  doc-page:
    textColor: "{colors.ink}"
    typography: "{typography.body}"
    width: "46rem"
---

# Design System: anishram.com

## Overview

**Creative North Star: "The datasheet and its timing diagram"**

The site reads like a component datasheet: a black header strip with a square-wave clock rule under it, a big condensed title, tabulated records, ruled sections and one signal colour. The proof sits on `/work/`: each Kubernetes feature Anish shipped is drawn as a digital signal that steps up a level (alpha, beta, GA) on the release that moved it. Maturity reads as height, so a visitor who does not know the Kubernetes vocabulary still sees which features went furthest.

Density is that of a reference document: tight vertical rhythm, hairline rules between rows, 2px ink rules under section heads, tabular figures. Nothing is rounded, nothing casts a shadow. The page is paper and ink; the signal colour marks only what is live: the trace, link underlines, roles, stage names, focus and selection. Dark mode is the inverse print of the same sheet, not a second palette.

Home introduces Anish in plain language with the title block, three areas, selected talks and posts, and a block for event organizers (a select-all bio and a headshot download). KEP numbers and the diagram stay on `/work/`.

**Key Characteristics:**
- One hue (signal magenta, 350) on a cool neutral sheet (hue 260); everything else is ink, ink-2 and rules.
- Sofia Sans for text, Sofia Sans Condensed for names, headings, numerals and labels. Both self-hosted, OFL.
- The timing diagram: release columns, a clock on the axis, one stepped trace per feature.
- Datasheet tables for posts and talks, collapsing to stacked rows under 48rem.
- Square corners, 1px and 2px rules, no shadows, no cards.

## Colors

A cool neutral sheet with a single magenta signal.

### Primary
- **Signal** (`sig`, light `#ae1173` / oklch(0.50 0.20 350); dark `#f990c4` / oklch(0.78 0.14 350)): the trace on the diagram, the clock rule under the strip, link underlines, roles, stage names in tags, list markers, focus rings, caret and selection.
- **Signal fill** (`sig-fill`, the signal at 8.5% light / 11% dark): the area under a trace, the speaker bio panel, level glyph fills.
- **Signal fill, raised** (`sig-fill-2`, 16% / 20%): the area under a trace while its row is hovered or focused.

### Neutral
- **Paper** (`paper`, `#fbfcfd` / dark `#0f1216`): page ground; text on the strip.
- **Ink** (`ink`, `#111419` / dark `#edeef1`): text, the header strip, 2px section rules, diagram frame.
- **Ink 2** (`ink-2`, `#4c5057` / dark `#b0b4bb`): secondary text, dates, table heads, the clock on the axis, the dotted pre-baseline.
- **Rule** (`rule`, `#b5b7bb` / dark `#4a4d53`): hairlines between rows and cells, tag borders.
- **Grid** (`grid`, `#e1e3e6` / dark `#202328`): release column lines inside diagram tracks only.

### Themes
- Light by default; dark under `prefers-color-scheme: dark` unless `data-theme="light"` is set, and always under `data-theme="dark"`.
- The "Inverse" toggle sets `data-theme`, stores it in `localStorage` ("theme") and updates `theme-color` to the strip colour (`#111419` light, `#edeef1` dark). Without JS the toggle stays `hidden` and the OS preference applies.
- The strip is always the inverse of the page: ink ground, paper text, in both themes.

### Contrast (WCAG 2.1 AA, computed from the OKLCH values converted to sRGB; fills composited over paper)

| Pair | Use | Need | Light | Dark |
|---|---|---|---|---|
| ink on paper | body text | 4.5 | 17.94 | 16.25 |
| ink-2 on paper | secondary text, dates | 4.5 | 7.88 | 9.07 |
| sig on paper | roles, stage names | 4.5 | 6.54 | 8.81 |
| sig on paper | trace, focus ring | 3 | 6.54 | 8.81 |
| ink on sig-fill | edge labels, speaker bio | 4.5 | 15.52 | 13.66 |
| ink on sig-fill-2 | edge labels, hovered row | 4.5 | 13.58 | 11.28 |
| ink-2 on sig-fill | secondary text over a fill | 4.5 | 6.81 | 7.63 |
| sig on sig-fill | trace over its fill | 3 | 5.66 | 7.41 |
| sig on sig-fill-2 | trace over hovered fill | 3 | 4.95 | 6.12 |
| ink-2 on paper | dotted baseline, axis clock | 3 | 7.88 | 9.07 |
| paper on ink | strip text | 4.5 | 17.94 | 16.25 |
| paper on ink | strip focus ring | 3 | 17.94 | 16.25 |
| toggle border on ink | toggle outline (45% paper in ink) | 3 | 3.81 | 3.36 |
| paper on sig | selection | 4.5 | 6.54 | 8.81 |
| ink on grid | text crossing a column line | 4.5 | 14.35 | 13.63 |

### Named Rules
**The One Signal Rule.** Magenta is the only hue. It marks what is live or actionable: the trace, links, roles, focus. Stages are told apart by height and name, never by colour.

**The Inverse Print Rule.** Dark mode swaps paper and ink and lifts the signal; it adds no colour and changes no layout.

**The Strip Rule.** The header strip is always ink on the page's paper, so its focus ring, selection and skip link carry their own paper treatment.

## Typography

**Display and label font:** Sofia Sans Condensed (fallback Sofia Sans, system-ui)
**Body font:** Sofia Sans (fallback system-ui)

**Character:** a technical grotesque in two widths. Condensed carries names, numerals and labels the way a datasheet sets part numbers; the normal width keeps running text comfortable. Both are variable woff2 (weights 1-1000), `font-display: swap`, the regular and condensed files preloaded.

### Hierarchy
- **Display** (Condensed 800, clamp(3.5rem, 10vw, 6rem), 0.9): the name on home.
- **Page title** (Condensed 800, clamp(3rem, 7vw, 5rem), 0.9): h1 on Work, Talks, Posts and list pages.
- **Doc title** (Condensed 800, clamp(2.25rem, 5vw, 3.5rem), 1): talk and post pages.
- **Contact** (Condensed 800, clamp(1.75rem, 5vw, 3.25rem)): the email address on home.
- **Section title** (Condensed 800, 1.625rem, 2px ink rule below): area and section heads, year heads.
- **Diagram title** (Condensed 800, clamp(1.375rem, 2.4vw, 1.75rem)).
- **Title** (Condensed 700, 1.25rem): entry names, area names, sub-heads.
- **Feature name** (Condensed 700, 1.125rem, 1.2): the label column of the diagram; also the wordmark.
- **Body** (Sans 400, 1.0625rem, 1.55, tabular figures); bio clamp(1.125rem, 1.6vw, 1.25rem); tables 1rem; prose up to 70ch.
- **Meta** (Sans 400, 0.9375rem, ink-2): venues, sources, captions, footer.
- **Label** (Condensed 700, 0.875rem): table heads, edge-label stage names, axis "release v1.".
- **Small** (Sans, 0.8125rem): feature tags and edge-label links.
- **Micro** (Condensed 600, 0.75rem): "in dev" under the axis numeral; axis numerals under 48rem.

### Named Rules
**The Two Widths Rule.** Condensed for anything you scan (names, numbers, dates, labels); normal width for anything you read.

**The Tabular Rule.** Figures are tabular everywhere, so release numbers and dates line up in columns.

## Layout

- Page container 84rem with a clamp(1rem, 3.5vw, 3rem) gutter; sections separated by clamp(3rem, 6vw, 5rem).
- Home: a title grid (name across; bio beside the photo; areas across), a two-column talks/posts block, the organizer block (bio panel beside the headshot), then contact.
- Work: page head with an "On this page" jump list; per area the lede, the diagram (for the ledger area), then the entries in an auto-fit grid (min 22rem); then the Posts and Talks tables.
- Diagram: a 15rem label column and a track of `--cols` equal release columns (at most 14; older releases fold into "≤N").
- Doc pages: a 46rem column, prose at 70ch.

### Responsive
- Under 64rem the diagram stacks: the label column sits over the track, edge labels hide and the `.stages` text list carries stage, release and links.
- Under 48rem the nav wraps under the wordmark, home and organizer columns stack, tables stack each row and hide their head visually (clipped, so the cells keep their column headers for screen readers; empty cells collapse to zero height), even axis numerals are hidden (except the release in development) and "in dev" goes.
- Under 30rem the title photo shrinks to 6.5rem.
- No horizontal overflow at 390px.

## Elevation & Depth

Flat. No shadows anywhere. Depth comes from rules (1px rule, 2px ink), the inverse strip, and the signal fill under a trace. The only box-shadow is the skip link's focus ring over the strip.

## Shapes

Square corners throughout (0px). Frames are 1px ink (diagram, photo mat with 6px paper padding); row dividers are 1px rule; section heads sit on a 2px ink rule. The recurring silhouette is the step: the square-wave clock, the trace, the level glyph and the toggle mark.

## Components

### Header strip and clock rule
- Ink strip, min-height 3.5rem: wordmark left, primary nav (Work, Posts, Talks, Email) and the toggle right. Current page underlined in paper.
- Under the strip a 12px clock rule: a repeating square wave in `sig`, drawn as a CSS mask, `aria-hidden`.
- Skip link: first Tab stop, paper box with an ink ring, shown over the strip.

### Theme toggle ("Inverse")
- Visible text "Inverse", accessible name "Inverse theme", `aria-pressed`. 44px min height, 1px border at 45% paper in ink, a half-filled mark that flips sides when pressed. `hidden` until JS runs (`.toggle[hidden] { display: none }`).

### Links
- Inherit text colour; 2px `sig` underline at 0.2em offset; hover turns the text `sig`. In the strip, underline in paper.

### Timing diagram (signature)
- `figure.diagram` framed in 1px ink: head (title and key), axis, one row per feature, caption.
- Axis: "release v1." label, a square-wave clock with one period per release (dashed in the dev column), numerals with "≤N" for the fold and "in dev" for the release in development.
- Row: label column (name, KEP link and SIG, co-authors, implementation PRs) and a track. The track is a grid of the release columns with a fixed-height wave (9rem; 4rem under 64rem), so every row draws its levels at the same heights. It holds an SVG (`viewBox 0 0 cols*100 300`, `preserveAspectRatio="none"`, non-scaling strokes): a dotted pre-baseline in ink-2, the fill under the trace, the trace in `sig` at 2.25 (3.5 on row hover or focus), and a dashed GA target in the dev column when a fresh target exists.
- Levels: baseline 280, alpha 195, beta 110, GA 25.
- Edge labels sit just under the level a stage steps to, as wide as the stage runs: stage name, then post and talk links. A one-column stage in the last column stacks its lines against the end (`edge--end`); its grid area also takes the column before, so only a line wider than the last column runs inward across the rise; a one-column stage followed by another, or one in the last column right after a stage in the column before it, stacks its links and truncates them (`edge--tight`). The row's last stage may use the free track to its end. Labels are grid items in one row, offset to their level, so a stage with many links grows its row instead of running into the next one. They are positioned, so they paint above the SVG, and the SVG takes no pointer events, so every label link is clickable. Two stages in one release share one label ("beta, GA").
- Key: alpha = low, beta = mid, GA = high, plus "GA target" only when one is drawn.
- Caption: how to read it, the fold, and the dev release with its checked date. All text comes from `data/ledger.yaml`.

### Level glyph
- A 1.4rem step at the stage's height (low, mid, high; dashed for a target). Used in the key, the stage list, feature tags and the "On the work page" note.

### Datasheet tables (posts, talks)
- Columns: Date, Title, Venue or source (with Recording and Slides links and achievements), Feature (where the ledger references the page). Condensed ink-2 heads over a 1px ink rule; 1px rule between rows.
- Feature tag: 1px rule box, level glyph, feature name and stage name in `sig`; links to `/work/#<id>`; border turns `sig` on hover.

### Event organizer block
- Speaker bio: the home bio, rendered once more in a `sig-fill` panel with a 1px rule border and a `sig` left edge; `user-select: all` so one click selects it.
- Headshot: a 10rem thumbnail in a paper mat; the link downloads a JPG generated from the source image at native size, its pixel size in the link text.

### Doc page (talk, post, 404)
- Talk and post pages: back link, doc title, subtitle, a meta bar between a 2px ink and 1px rule, action links, then prose at 70ch with `sig` list markers and blockquote edge. Pages referenced by the ledger end with an "On the work page" note. The 404 uses the same column with its body and action links.

### Accessibility
- One h1 per page; skip link to `main#main`; landmarks labelled (Primary, Elsewhere, On this page).
- Tables keep `th scope="col"` headers at every width: under 48rem the head is clipped, not `display: none`, and empty cells stay in the table, so each cell keeps its column header.
- The diagram is a `figure` labelled by its title; each row is a list item with a heading, and every stage and link exists as text (the SVG is `aria-hidden`); folded and truncated releases stay in visually hidden text.
- Focus is a 2px `sig` outline at 3px offset (paper inside the strip).
- Motion: only the trace thickening transitions, and only under `prefers-reduced-motion: no-preference`.

## Do's and Don'ts

### Do:
- **Do** keep magenta to the trace, link underlines, roles, stage names, focus and selection.
- **Do** set names, numbers, dates and labels in Sofia Sans Condensed with tabular figures.
- **Do** render every diagram element from `data/ledger.yaml`, and every list from content, so new KEPs, releases, talks and posts need no template edit.
- **Do** keep the strip the inverse of the page in both themes.
- **Do** check new contrast pairs against the table above, compositing fills over paper.

### Don't:
- **Don't** add a second hue, or colour stages differently; height carries the stage.
- **Don't** round corners, add shadows or wrap content in cards.
- **Don't** put KEP numbers or the diagram on home.
- **Don't** load fonts or scripts from third parties, or add a JS framework.
- **Don't** draw a stage, target or role that the data does not record.
