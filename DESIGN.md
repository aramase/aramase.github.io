---
name: anishram.com
description: Anish Ramasekar's site in spruce and saffron, Jost and DM Sans, with his Kubernetes work set as a release ledger.
colors:
  ground: "oklch(0.975 0.008 170)"
  ink: "oklch(0.25 0.04 195)"
  ink-2: "oklch(0.40 0.035 195)"
  ink-3: "oklch(0.48 0.03 195)"
  rule: "oklch(0.87 0.02 180)"
  ghost: "oklch(0.93 0.014 175)"
  band: "oklch(0.31 0.055 190)"
  band-ink: "oklch(0.97 0.012 160)"
  band-ink-2: "oklch(0.86 0.03 170)"
  band-rule: "oklch(0.42 0.05 190)"
  saffron: "oklch(0.83 0.14 78)"
  accent: "oklch(0.50 0.11 62)"
  on-saffron: "oklch(0.25 0.04 195)"
  alpha: "oklch(0.50 0.13 35)"
  beta: "oklch(0.50 0.10 110)"
  ga: "oklch(0.43 0.09 200)"
  ground-dark: "oklch(0.205 0.03 200)"
  ink-dark: "oklch(0.95 0.012 165)"
  ink-2-dark: "oklch(0.83 0.02 175)"
  ink-3-dark: "oklch(0.71 0.025 185)"
  rule-dark: "oklch(0.34 0.035 195)"
  ghost-dark: "oklch(0.27 0.03 197)"
  band-dark: "oklch(0.285 0.05 190)"
  band-rule-dark: "oklch(0.40 0.05 190)"
  accent-dark: "oklch(0.83 0.14 78)"
  alpha-dark: "oklch(0.77 0.12 40)"
  beta-dark: "oklch(0.82 0.12 115)"
  ga-dark: "oklch(0.80 0.09 190)"
typography:
  display:
    fontFamily: "Jost, Futura, ui-sans-serif, system-ui, -apple-system, Segoe UI, sans-serif"
    fontSize: "clamp(2.75rem, 1.4rem + 5.6vw, 5.5rem)"
    fontWeight: 500
    lineHeight: 0.95
    letterSpacing: "-0.02em"
  headline:
    fontFamily: "Jost, Futura, ui-sans-serif, system-ui, -apple-system, Segoe UI, sans-serif"
    fontSize: "clamp(2rem, 1.4rem + 2.6vw, 3rem)"
    fontWeight: 500
    lineHeight: 1.15
    letterSpacing: "-0.02em"
  section-title:
    fontFamily: "Jost, Futura, ui-sans-serif, system-ui, -apple-system, Segoe UI, sans-serif"
    fontSize: "1.625rem"
    fontWeight: 500
    lineHeight: 1.15
    letterSpacing: "-0.01em"
  focus-name:
    fontFamily: "Jost, Futura, ui-sans-serif, system-ui, -apple-system, Segoe UI, sans-serif"
    fontSize: "1.25rem"
    fontWeight: 500
    lineHeight: 1.15
  release-numeral:
    fontFamily: "Jost, Futura, ui-sans-serif, system-ui, -apple-system, Segoe UI, sans-serif"
    fontSize: "clamp(1.5rem, 0.9rem + 1.6vw, 2.5rem)"
    fontWeight: 300
    lineHeight: 1
    letterSpacing: "-0.02em"
    fontFeature: "tnum"
  year-numeral:
    fontFamily: "Jost, Futura, ui-sans-serif, system-ui, -apple-system, Segoe UI, sans-serif"
    fontSize: "clamp(2rem, 1.4rem + 2.6vw, 3rem)"
    fontWeight: 300
    lineHeight: 1
    letterSpacing: "-0.03em"
    fontFeature: "tnum"
  feature-name:
    fontFamily: "Jost, Futura, ui-sans-serif, system-ui, -apple-system, Segoe UI, sans-serif"
    fontSize: "1.0625rem"
    fontWeight: 650
    lineHeight: 1.15
  wordmark:
    fontFamily: "Jost, Futura, ui-sans-serif, system-ui, -apple-system, Segoe UI, sans-serif"
    fontSize: "1.125rem"
    fontWeight: 600
  nav:
    fontFamily: "Jost, Futura, ui-sans-serif, system-ui, -apple-system, Segoe UI, sans-serif"
    fontSize: "0.9375rem"
    fontWeight: 500
    letterSpacing: "0.01em"
  bio:
    fontFamily: "DM Sans, ui-sans-serif, system-ui, -apple-system, Segoe UI, sans-serif"
    fontSize: "clamp(1.125rem, 1rem + 0.5vw, 1.3125rem)"
    fontWeight: 400
    lineHeight: 1.55
  body:
    fontFamily: "DM Sans, ui-sans-serif, system-ui, -apple-system, Segoe UI, sans-serif"
    fontSize: "1.0625rem"
    fontWeight: 400
    lineHeight: 1.55
  small:
    fontFamily: "DM Sans, ui-sans-serif, system-ui, -apple-system, Segoe UI, sans-serif"
    fontSize: "0.8125rem"
    fontWeight: 400
    lineHeight: 1.45
  label:
    fontFamily: "Jost, Futura, ui-sans-serif, system-ui, -apple-system, Segoe UI, sans-serif"
    fontSize: "0.8125rem"
    fontWeight: 700
    letterSpacing: "0.06em"
rounded:
  none: "0px"
spacing:
  space-1: "0.25rem"
  space-2: "0.5rem"
  space-3: "0.75rem"
  space-4: "1rem"
  space-5: "1.5rem"
  space-6: "2rem"
  space-7: "3rem"
  space-8: "4rem"
  space-9: "6rem"
  gutter: "clamp(1.25rem, 4vw, 3rem)"
  page: "76rem"
  measure: "60ch"
  label-col: "minmax(14rem, 19rem)"
  doc: "48rem"
components:
  nav-link:
    textColor: "{colors.ink-2}"
    typography: "{typography.nav}"
    height: "44px"
    padding: "0 0.25rem"
  nav-link-current:
    textColor: "{colors.ink}"
    typography: "{typography.nav}"
    height: "44px"
  nav-link-band:
    textColor: "{colors.band-ink-2}"
    typography: "{typography.nav}"
    height: "44px"
  theme-toggle:
    textColor: "{colors.ink-2}"
    typography: "{typography.nav}"
    height: "44px"
  skip-link:
    backgroundColor: "{colors.ink}"
    textColor: "{colors.ground}"
    rounded: "{rounded.none}"
    padding: "0.75rem 1rem"
  skip-link-band:
    backgroundColor: "{colors.saffron}"
    textColor: "{colors.on-saffron}"
  band:
    backgroundColor: "{colors.band}"
    textColor: "{colors.band-ink}"
  hero-bio:
    textColor: "{colors.band-ink-2}"
    typography: "{typography.bio}"
  name-bar:
    backgroundColor: "{colors.saffron}"
    width: "3.5rem"
    height: "0.375rem"
  action-link:
    textColor: "{colors.accent}"
    typography: "{typography.small}"
    height: "44px"
  axis-numeral:
    textColor: "{colors.ink-2}"
    typography: "{typography.release-numeral}"
  axis-numeral-in-dev:
    textColor: "{colors.ink-3}"
    typography: "{typography.release-numeral}"
  stage-alpha:
    textColor: "{colors.alpha}"
    typography: "{typography.small}"
  stage-beta:
    textColor: "{colors.beta}"
    typography: "{typography.small}"
  stage-ga:
    textColor: "{colors.ga}"
    typography: "{typography.small}"
  line-beta:
    backgroundColor: "{colors.beta}"
    width: "1.75rem"
    height: "3px"
  line-ga:
    backgroundColor: "{colors.ga}"
    width: "1.75rem"
    height: "0.5rem"
  feature-row:
    textColor: "{colors.ink}"
    padding: "1rem 0 1.5rem"
  feature-row-unlit:
    textColor: "{colors.ink-3}"
  record-entry:
    textColor: "{colors.ink}"
    typography: "{typography.body}"
    padding: "0.75rem 0"
  record-entry-unlit:
    textColor: "{colors.ink-3}"
  dated-row:
    textColor: "{colors.ink}"
    typography: "{typography.feature-name}"
    padding: "1rem 0"
  doc-page:
    width: "48rem"
    padding: "4rem 0 0"
  doc-action:
    textColor: "{colors.ink}"
    height: "44px"
  area-entry:
    textColor: "{colors.ink-2}"
    typography: "{typography.body}"
    padding: "1rem 0"
  focus-row:
    textColor: "{colors.ink-2}"
    typography: "{typography.body}"
    padding: "1rem 0"
  pick:
    textColor: "{colors.ink}"
    typography: "{typography.body}"
    padding: "1rem 0"
  contact-email:
    textColor: "{colors.ink}"
    typography: "{typography.headline}"
---

# Design System: anishram.com

## Overview

**Creative North Star: "Spruce and Saffron"**

The home page introduces Anish to people looking him up: photo, name, bio, three plain-language focus areas, a few selected talks and posts, and contact. Its first screen is a deep spruce band that runs full width behind the header and the hero; a short saffron bar under the name is the one bright mark. Below the band the page is a pale spruce-tinted ground with dark spruce ink, and saffron returns only where something can be acted on: link underlines, action links, focus rings, text selection.

The full record lives on `/work/` ("Work"), in three areas: Kubernetes identity and security, secrets and workload identity, and agent infrastructure. Each area opens on its title and a one- or two-line lede, then short entries (name, role, one line, link). The Kubernetes area carries the release ledger, where features advance across a grid of Kubernetes releases, read the way release notes and enhancement tracking read: hairline rules, numerals set large and light. Three stage hues say what stage a feature reached, and every stage is also drawn in a distinct line weight, so the ledger still reads in greyscale. The one interaction lights a feature's posts and talks in the record and unlights the rest.

**Key Characteristics:**
- Two tones carry the identity: spruce (the band) and saffron (action and focus). Stage hues appear only on stages.
- Two self-hosted variable faces: Jost (geometric, Futura lineage) names things; DM Sans carries text. Tabular figures for every date and version.
- Hairline rules for rows; no cards, no shadows, no radii. The band is the only filled surface.
- A release axis of large light numerals over a grid of ghost cells.
- Colour-only motion, removed under reduced motion.

## Colors

The CSS custom properties in `assets/css/site.css` are the source of truth, written in OKLCH. Every component reads the properties, never a literal colour.

### Primary
- **Spruce band** (`band` / `band-dark`): background of the header and hero on the home page only, full bleed. `band-ink` sets the name and wordmark on it, `band-ink-2` the bio and nav links, `band-rule` the header hairline and the photo's 1px outline.

### Secondary
- **Saffron** (`saffron`): the bar under the name, the contact email's 3px underline, nav hover and current underlines on the band, focus rings on the band, the home skip link fill, and the text selection fill (both with `on-saffron` text). Never used for text on the light ground, where it fails contrast.
- **Accent** (`accent` / `accent-dark`): saffron as a text colour. On light ground it is darkened to a deep amber that holds AA; in dark it is saffron itself. Used for action links (focus-row links, "All 14 talks", "All 9 posts"), link underlines (70% mix at rest, full on hover) and the focus ring off the band.

### Tertiary: stage hues
- **Alpha rust** (`alpha` / `alpha-dark`), **Beta olive** (`beta` / `beta-dark`), **GA teal** (`ga` / `ga-dark`): alpha, beta and GA stages, stage tags and line samples on `/work/` and in "On the ledger" notes. Alpha is warm, beta olive and GA cool; olive keeps beta clear of the amber accent, so no stage reads as a link.

### Neutral
- **Ground** (`ground` / `ground-dark`): page background; also the knockout behind each stage label so the ghost-cell hairline does not run through text. Mirrored as hex in the `theme-color` meta tags, `site.js`, `manifest.json` and `browserconfig.xml` (`#f2f9f6`, `#041b1d`).
- **Ink** (`ink` / `ink-dark`): headings, titles, body links, the 2px rules that open dated lists and the doc ledger note, the skip link fill off the home page.
- **Ink 2** (`ink-2` / `ink-2-dark`): secondary text: focus-area lines, meta lines, axis numerals, nav links off the band, footer.
- **Ink 3** (`ink-3` / `ink-3-dark`): dates, the in-development release numeral, notes, and the unlit state of the signature interaction. Lowest neutral text tone.
- **Rule** (`rule` / `rule-dark`): 1px hairlines between rows, ghost-cell ticks and baseline, unlit link underlines. Never used for text.
- **Ghost** (`ghost` / `ghost-dark`): the faintest divider, used between stacked stage rows in the folded ledger.

### Themes
- Light is the `:root` default; dark applies under `prefers-color-scheme: dark` unless `data-theme="light"` is set, and always under `data-theme="dark"`.
- Both themes define the same sixteen properties (`--ground`, `--ink`, `--ink-2`, `--ink-3`, `--rule`, `--ghost`, `--band`, `--band-ink`, `--band-ink-2`, `--band-rule`, `--saffron`, `--accent`, `--on-saffron`, `--alpha`, `--beta`, `--ga`). `band-ink`, `band-ink-2`, `saffron` and `on-saffron` are the same in both.
- In dark the ground drops below the band, so the band still reads as the lighter field at the top; stage hues are lifted lighter, not desaturated.

### Contrast (WCAG 2.1 AA, computed from the OKLCH values converted to sRGB)
- Light, on ground: ink 14.8, ink 2 8.5, ink 3 6.0, accent 5.8, alpha 6.0, beta 5.5, GA 7.1.
- Dark, on ground: ink 15.5, ink 2 10.7, ink 3 7.0, accent 10.4, alpha 8.2, beta 10.4, GA 9.9.
- On the band (light / dark): band ink 11.8 / 12.9, band ink 2 8.5 / 9.3.
- Selection and the home skip link: on-saffron on saffron 9.3. Skip link elsewhere: ground on ink 14.8 / 15.5.
- Focus rings (3:1 non-text): accent on ground 5.8 light, saffron on dark ground 10.4, saffron on the band 7.5 / 8.2 (the home skip link's saffron fill and ring against the band are the same pair).
- Saffron on the light ground is 1.6, which is why it never sets text there.

### Named Rules
**The Two Tones Rule.** Spruce owns the top of the home page; saffron marks the name, action and focus. No other brand hue, no tints beyond the 70% underline mix, no gradients.

**The Stage Hue Rule.** Alpha rust, beta olive and GA teal appear only where a stage is named or drawn. They never mark links, hover or decoration.

**The Double Code Rule.** Every stage is coded twice, by hue and by line weight: alpha is a 2px dashed line, beta a 3px solid line, GA a 0.5rem solid block, GA target a 0.5rem box outlined with a 1.5px dashed GA border. The stage name is always written as text beside the line.

**The Unlit Rule.** Quieting is done by reassigning the three stage properties and text colour to ink 3, never by opacity, blur or filter. Unlit content still passes AA.

## Typography

**Voice font:** Jost (variable, weight 100 to 900, self-hosted from `/fonts/jost-latin.woff2`, licence `/fonts/Jost-OFL.txt`) with `Futura, ui-sans-serif, system-ui, -apple-system, "Segoe UI", sans-serif`. Wordmark, nav and theme toggle, every h1 to h4, the contact email, and the index numerals (release axis, years).
**Text font:** DM Sans (variable, weight and optical size, self-hosted from `/fonts/dm-sans-latin.woff2`, licence `/fonts/DMSans-OFL.txt`) with `ui-sans-serif, system-ui, -apple-system, "Segoe UI", sans-serif`. Body, bio, meta lines, links, labels that are not headings.
Both are latin subsets, preloaded, `font-display: swap`. There is no monospace anywhere.

**Character:** Jost's geometric capitals and round bowls give the name and headings a confident, slightly mid-century voice at medium weight; DM Sans is a low-contrast grotesk that stays quiet at reading sizes.

### Hierarchy
- **Display** (Jost 500, `clamp(2.75rem, 1.4rem + 5.6vw, 5.5rem)`, line-height 0.95, tracking -0.02em): the name on the home page, followed by the saffron bar.
- **Headline** (Jost 500, `clamp(2rem, 1.4rem + 2.6vw, 3rem)`, tracking -0.02em): list and detail page titles (detail pages at line-height 1.08).
- **Section title** (Jost 500, 1.625rem, tracking -0.01em): "What I work on", "Selected talks", and the area titles on `/work/` ("Kubernetes identity and security", "Agent infrastructure"). The ledger title inside its area ("Kubernetes features I've shipped", with no release range) is Jost 500 at the focus-name size.
- **Focus name** (Jost 500, 1.25rem): the three area names on the home page.
- **Release numeral** (Jost 300, `clamp(1.5rem, 0.9rem + 1.6vw, 2.5rem)`, line-height 1, tabular): the ledger axis. The `v1.` prefix sits small in ink 3 at the left of the axis.
- **Year numeral** (Jost 300, headline size, tabular): year headings on the Posts and Talks lists; sticky beside its rows on wide screens.
- **Feature name** (Jost 650, 1.0625rem): ledger row names and area entry names; dated-row titles at 600.
- **Bio** (DM Sans 400, `clamp(1.125rem, 1rem + 0.5vw, 1.3125rem)`, line-height 1.55, max 36em): the home bio, at every width.
- **Body** (DM Sans 400, 1.0625rem, line-height 1.55): paragraphs, capped at 60ch; row summaries cap at 62ch.
- **Small** (DM Sans 400 to 650, 0.8125rem): meta lines, stage labels, legend, dates, footer, notes, action links (600).
- **Label** (Jost 700, 0.8125rem, 0.06em tracking, uppercase): headings that name a list: "Posts", "Talks", "Pinned", "On the ledger".
- **Wordmark** (Jost 600, 1.125rem) and **Nav** (Jost 500, 0.9375rem, 0.01em tracking, sentence case): the header.
- Headings h1 to h4 default to Jost 650, line-height 1.15, `text-wrap: balance`; paragraphs use `text-wrap: pretty`.

### Named Rules
**The Two Faces Rule.** Jost names things, DM Sans says things. No third face.

**The Tabular Rule.** Every date, version and release number is set with tabular figures, so columns of numerals align without a monospace face.

**The Light Numeral Rule.** Numerals that index the record (release axis, years) are set large at Jost 300; words that name things are set at 500 to 650. Weight separates index from content.

## Layout

- **Container:** max width 76rem, centred, side gutter `clamp(1.25rem, 4vw, 3rem)`.
- **Spacing scale:** nine steps from 0.25rem to 6rem (space-1 to space-9). Rows pad 1rem above and 1.5rem below; page heads pad 4rem above; the footer sits 6rem below content.
- **Section openers:** home sections, the work page's areas, its ledger and its record open on their title alone, with space instead of a rule. Dated lists (`.rows`) and the doc "On the ledger" note open on a 2px ink rule. Rows inside are separated by 1px rule hairlines.
- **Ledger grid:** a label column of `minmax(14rem, 19rem)` then one equal column per release (`--cols`, set from the data, at most 14). The axis, feature list and each feature row share it through subgrid, so every stage aligns to its release numeral.
- **Shared label column:** area entries reuse the 14rem to 19rem label column with a single fluid line column, so entry names align with the ledger's feature names.
- **Areas:** each `/work/` area is a section with 4rem below it; the ledger sits 3rem under the Kubernetes area's entries.
- **Home band:** the header and a `.band` wrapper around the hero share the band colour and run the full viewport width; the hero sits in the normal 76rem container inside it. A 13rem square headshot spans two rows beside the name and the bio; 4rem padding above and below; 4rem of ground before "What I work on".
- **Home sections:** focus rows use the shared label column, a fluid line column and a link column, divided by 1px rules; selected talks and posts sit in two equal columns with a 3rem gap, each list opening on a 1px rule.
- **Record:** two equal columns (Posts, Talks) with 3rem gaps; each entry is a 5.5rem date column plus title.
- **Dated lists:** a 9rem year column beside rows of date (5rem), title and summary (fluid), and links or source (`minmax(9rem, 12rem)`).
- **Doc pages:** single column, max 48rem, body capped at 60ch.

### Responsive
- **56rem, the fold:** home focus rows stack name, line and link. The ledger grid folds flat. The axis and ghost cells are hidden; each feature becomes a stacked stage list in the same order: release label (now visible, 650 weight, 4.25rem column), a 2.25rem line sample, the stage name, then links right-aligned. Stage rows are divided by ghost-tone hairlines. Area entries stack name and role over the line.
- **48rem:** the record and the home selected talks and posts drop to one column; year and pinned grids stack the year above its rows and the year loses its sticky position; dated rows become a single column; row links grow to 44px targets.
- **40rem:** the header nav moves to its own full-width row under a hairline (band rule on the home band) with links spread edge to edge; the home hero puts a 5.5rem photo beside the name and runs the bio full width; the contact email drops to 1.5rem to 2rem; the doc pager becomes one column.
- **22.5rem:** the home name drops to 2.25rem so "Ramasekar" fits beside the photo at 320px.

### Named Rules
**The Every Cell Exists Rule.** Each axis column in each feature row, the folded "≤N" column included, is drawn as exactly one ghost cell (a 1px rule tick and a 1px rule baseline), whether or not a stage occurred. Stages are laid over the cells; nothing is drawn to fill an empty release.

## Elevation & Depth

Flat. There are no shadows and no raised panels. The spruce band is the one filled surface, and it is a field, not a card: it touches both viewport edges and has no border radius or shadow. Elsewhere hierarchy comes from rule weight, ink tone and space. The only stacking is the ground-coloured knockout behind stage labels, which hides the ghost-cell baseline under text.

### Named Rules
**The Flat Paper Rule.** No `box-shadow`, no elevation, no raised panels. If something needs separating, use a rule or space.

## Shapes

- Square corners everywhere (0px). The only radius in the build is 1px on the focus outline.
- Form language is lines and bars: the saffron name bar, dashed and solid stage strokes, the GA block, hairline ticks. The headshot is a hard square crop with a 1px band-rule outline.
- The theme toggle's state mark is a drawn 0.75rem square (1.5px border in the current text colour, filled when dark is active), not a glyph.

### Named Rules
**The Square Corner Rule.** Nothing is rounded. A corner radius would turn a row into a card.

## Components

### Header
- Wordmark (site title, Jost 600, 44px tall) pushed left; nav and theme toggle right; 1px rule below.
- Nav links (Work, Posts, Talks, Email): Jost 500 in ink 2 at rest; hover and `aria-current="page"` switch to ink with a 2px underline offset 0.5em.
- On the home page the header carries `.site-header--band`: band background, band-rule hairline, nav in band ink 2, and hover or current in band ink with a saffron underline; focus rings turn saffron.
- Theme toggle: a button with the drawn square and the word "Dark"; `aria-pressed` tracks the effective theme. Hidden until the script runs.

### Footer
- 1px rule above, 2rem and 3rem vertical padding, small type in ink 2.
- Elsewhere links (GitHub, LinkedIn, X, Bluesky, YouTube, email, RSS) as plain text links on the left, copyright on the right, every item 44px tall. No logos. The home page omits them, since its Contact section carries the same links.

### Home hero, focus rows, selected work, contact
- Hero (inside the band): square headshot, the name as the display h1 in band ink with a 3.5rem by 0.375rem saffron bar under it, the bio as one paragraph in band ink 2. No role line; the bio carries the role.
- Focus rows ("What I work on"): area name (h3, focus name style), one plain-language line in ink 2, and one action link in accent at small 600 weight to the matching `/work/` area (`#kubernetes`, `#secrets`, `#agents`). No KEP numbers on the home page.
- Selected talks and posts: section title, then three entries each from `data/home.yaml` between 1px rules: title (600 link), a small meta line (date, then venue or source when the page has one) and, for a talk with one, its achievement line. An "All 14 talks" or "All 9 posts" action link in accent closes each column.
- Contact: the email address in Jost 500 at headline size with a 3px saffron underline, then the profile links and RSS at 600 weight, 44px tall.

### Links
- Inherit text colour; 1px underline in accent mixed to 70% with transparent, offset 0.22em. Hover thickens to 2px in full accent.
- Focus: 2px accent outline, 3px offset; saffron on the band.
- Selection: saffron fill with on-saffron text.

### Ledger grid and axis
- Section title plus legend on one line, wrapping when tight.
- Section title: "Kubernetes features I've shipped", with no release range; the axis carries the releases.
- The axis follows Anish's own stages, not the Kubernetes calendar. It starts at `first`, the oldest released stage, and ends at `last`, the newest released stage; targets never extend it. A Kubernetes release with no new stage leaves the ledger unchanged.
- An "in dev" column follows `last` only while a fresh target is drawn (see below).
- At most 14 columns. Past that the oldest releases fold into a leading "≤N" column: N is the newest folded release, the "≤" hangs left of the numeral in ink 3, and each feature shows only its latest stage from the folded releases there, written with its release ("GA v1.29"). A stage that began in the fold and continues runs from the "≤N" column.
- Axis: `v1.` prefix in small ink 3, then one light numeral per release in ink 2; the in-development release is ink 3 with a 0.625rem uppercase "in dev" line under it. Hidden from assistive tech (each stage carries its own release text).
- A note under the ledger, small ink 3, names the source of the stages; while a target is drawn it also explains the dashed GA target box and gives the date the target was checked.
- A target is drawn only while the KEP issue's milestone is the in-development release and it carries `lead-opted-in`; a kep.yaml target alone is not enough. The generator writes `inDevelopment` and `checked` only with a target, and the site stops drawing targets 35 days after `checked` (UTC days), everywhere at once: no "in dev" column, legend entry or note on the ledger, and no "GA target" in post and talk ledger notes; a feature whose only stage is a stale target drops off the ledger. The build fails if `inDevelopment` is set without `checked`.
- KEP-only: the ledger tracks Kubernetes KEPs Anish carried through releases. Projects and agent-platform work stay in the area entries, never on the release axis.

### Feature row
- Label column: feature name, then a meta line with the KEP link, SIG, authorship and implementation PR links (PR numbers tabular).
- Track: ghost cells, then stages spanning from their release to the next stage's release. An open stage (the latest, short of GA and not a target) ends at `last`, never in the "in dev" column, with a 2px end-cap in its hue, and its name reads "beta since v1.34" (the "since" part small ink 3 at 400, hidden in the folded list, where the release label already says it).
- A cue link ("2 posts, 1 talk in the record below") appears under the label only while the row is lit and has related entries.

### Stage
- Line first (alpha dashed, beta solid, GA block, GA target dashed box), then the stage name in its hue at 650, then links ("k8s.io", "Microsoft", "Talk") each with hidden full-title context for screen readers.
- Stage name, stage tag and line sample always share a hue and line weight.

### Ghost cells
- 0.5rem tall, 1px rule left tick, 1px rule baseline at mid-height; one per axis column per feature, the folded column included. Decorative and `aria-hidden`.

### Legend
- Small type in ink 2: a 1.75rem line sample for alpha, beta and GA (and GA target, only while a target is drawn) followed by its name. The same line samples reappear in the doc ledger note.

### Work page head and area entries
- Page head: "Work", one line, then an index of the three areas as accent links at small 600 weight, 44px tall.
- Area: section title, a lede of one or two lines in ink 2, then entries between 1px rules.
- Entry: name (h3, feature-name style, linked to the project), role in the small meta line, and one line in ink 2 capped at 62ch. No timelines, no release lists.
- An entry with a KEP outside the release grid (the Secrets Store CSI Driver, a subproject in `data/ledger.yaml`) adds the KEP link and its stage tag ("GA at v1.22") to the meta line and joins the cross-highlight, lighting its talks in the record.

### Record entry
- Date (small, ink 3, tabular) beside the title (550 weight link) and a meta line: source or venue, then one stage tag per ledger reference ("KEP-4412 beta" in the beta hue).
- Lit (related to the focused feature): title gains a 2px underline. Unlit: text and stage tags fall to ink 3, underlines to rule.
- Section headings "Posts" and "Talks" use the label style over a 1px rule; an "All posts" or "All 14 talks" link closes each column.

### Dated rows (Posts, Talks)
- Page head: headline title plus a one-line count ("9 posts, newest first").
- Year groups: light year numeral, rows under a 2px ink rule, 1px rule between rows.
- Row: date, title (Jost 600, no underline until hover) with venue or summary, aside with source and stage tags (blogs) or Recording and Slides links (talks). A talk credential line sits under the title in ink at weight 600.
- Talks list opens with a "Pinned" group (label-style heading in a 9rem column) using the same rows.

### Doc page (talk detail, post, 404)
- Back link ("All talks" or "All posts"), headline title, optional subtitle, then a meta strip between two 1px rules (date, venue, location, co-speaker).
- Actions ("Watch the recording", "Slides", "Playlist", "Read the full post on ...") as 650-weight text links, 44px tall.
- "On the ledger" note under a 2px ink rule: line sample, feature link back to its ledger row on `/work/`, stage tag, release and KEP.
- Post pages end with a Newer/Older pager under a 1px rule, its titles underlined like links; 404 reuses the doc layout with links home, to the work page, to Posts and to Talks.

### Signature interaction: the ledger cross-highlight
- Hovering a feature row or an area entry with a KEP with a mouse, or moving focus into it, lights it and its related record entries and unlights the other features, the other area entries, and unrelated entries.
- Touch does not trigger it; it ends on pointer leave (unless focus is inside), on focus leaving the row, or on Escape.
- Motion: colour, background colour, border colour and underline colour transition over 160ms with `cubic-bezier(0.16, 1, 0.3, 1)`. Nothing moves, scales or fades. Under `prefers-reduced-motion: reduce` all transitions and animations are removed. A theme change never transitions: `site.js` holds `.theme-switching` on the root across the switch so text and ground change in the same frame.
- Without JavaScript the ledger and record render fully lit; the site does not depend on the script.

### Named Rules
**The Colour-Only Motion Rule.** The only motion is a 160ms colour change on the cross-highlight. No transforms, no entrance animation, no scroll effects, and none at all under reduced motion.

### Accessibility
- WCAG 2.1 AA text contrast in both themes, including unlit text, the band and every stage hue on ground (lowest text pair: beta on light ground at 5.5:1; see Colors).
- One h1 per page; a skip link that appears on focus and targets the main region: ink fill and ground text, or on the home page, where it lands on the band, saffron fill, on-saffron text and a saffron ring.
- Visible 2px focus ring on every focusable element (accent, saffron on the band).
- Stage meaning never relies on hue alone: line weight and written stage names carry it too. Line samples, ghost cells and the axis are `aria-hidden`; stages and links carry text equivalents.
- Header, footer, doc actions, home action links, record "All" links and folded stage links are at least 44px tall.
- Inline ledger stage links on desktop are at least 24px tall (WCAG 2.2 AA target size); row links in list asides are 2rem.
- Theme follows the OS by default; a manual choice persists in `localStorage` and is applied before first paint.

## Do's and Don'ts

### Do:
- **Do** read every colour through the custom properties so both themes and the unlit state work without extra rules.
- **Do** keep saffron off text on the light ground; use `accent` for saffron-coloured text.
- **Do** pair every stage hue with its line weight and written name: alpha 2px dashed, beta 3px solid, GA 0.5rem block, GA target 0.5rem dashed box.
- **Do** draw every release cell for every feature, lit or not.
- **Do** set dates, versions and release numbers in tabular figures; set index numerals in Jost 300 and names at 500 to 650.
- **Do** keep interactive text targets at 44px where they stand alone (nav, footer, actions, folded stage links).
- **Do** quiet by switching to ink 3, and keep that state AA.

### Don't:
- **Don't** add a brand hue beyond spruce and saffron, or use a stage hue for links, hover or decoration.
- **Don't** use cards, panels, shadows or rounded corners; the band is the only filled surface.
- **Don't** use gradients of any kind.
- **Don't** use icon glyphs (icon fonts, SVG pictograms, emoji, arrows as icons); the drawn line samples, the saffron bar and the toggle's drawn square are the only marks.
- **Don't** use terminal costume: no monospace, prompt characters, typewriter effects, blinking cursors or window chrome.
- **Don't** animate position, size or opacity; motion is colour only, and none under reduced motion.
- **Don't** quiet content with opacity or blur.
- **Don't** add a third typeface.
