# SaaS Landing Template — Blocks & Content Model

Status: product spec (first template)  
BuildY anchor: **CMS-backed landing** — first proof in `progress.md`.

## Template identity

| Field | Value |
|-------|-------|
| Template ID | `saas-landing-v1` |
| Category | `templates/saas` |
| Target ICP | Founder-led SaaS, FastyGo product team |
| Pages (MVP) | `home` (single landing route; multi-page later) |
| Output | Go/Templ views + AppCMS content types + seed entries |

## Template page (marketing gallery)

Structure for `/templates/saas/saas-landing-v1` public page:

```text
Hero title + one-line outcome
Preview screenshot / live fixture link
Key Highlights (4 bullets)
Features & Capabilities (block list)
CTA: Preview · Create from template
```

### Key highlights (marketing copy)

1. **CMS-backed from generation** — every section maps to editable content types, not hardcoded strings.
2. **SaaS conversion sections** — hero, features, pricing, proof, FAQ, CTA out of the box.
3. **BFF/Templ preview** — see real navigation and content gaps before export.
4. **Export-ready FastyGo source** — Go/Templ/AppCMS-compatible tree for GitHub handoff.

## Landing blocks (MVP — required)

From north star first proof — standard SaaS landing section patterns, scoped to **one page**:

| Order | Block ID | Purpose | CMS content type |
|-------|----------|---------|------------------|
| 1 | `hero` | Value prop, subhead, primary/secondary CTA | `landing_hero` |
| 2 | `logo_strip` | Social proof logos (optional MVP) | `landing_logos` |
| 3 | `features_grid` | 3–6 feature cards | `landing_feature` |
| 4 | `pricing` | 2–3 tiers, bullets, highlight tier | `landing_pricing_tier` |
| 5 | `testimonials` | Quotes, name, role, avatar | `landing_testimonial` |
| 6 | `faq` | Question/answer pairs | `landing_faq_item` |
| 7 | `cta_banner` | Final conversion section | `landing_cta` |
| 8 | `footer` | Links, legal placeholders, social | `landing_footer` |

Global / page-level: **`seo_settings`** — title, description, OG image, canonical.

## Field mappings (example: `landing_hero`)

| Field | Type | Required | Notes |
|-------|------|----------|-------|
| `headline` | string | yes | H1 |
| `subheadline` | text | yes | Supporting line |
| `primary_cta_label` | string | yes | |
| `primary_cta_href` | string | no | Default `#signup` |
| `secondary_cta_label` | string | no | |
| `secondary_cta_href` | string | no | |
| `media_image` | media ref | no | Hero visual |

Repeat similar explicit mappings for each block in implementation spec.

## Staging location (BuildY app)

Per buildy rules:

```text
internal/ui/blocks/marketing/   # hero, features, pricing, …
internal/views/                 # landing page compose
internal/fixtures/locale/       # EN/RU copy for mockup phase
```

Promote to shared packages only after freeze.

## Sections deferred (post-MVP)

From services template (pet-sitter) — **do not** include in v1:

| Block | Why deferred |
|-------|--------------|
| Booking wizard | Product app scope, not landing proof |
| Availability calendar | Domain-specific |
| Client dashboard | Authenticated app surface |
| Multi-page nav (11 pages) | Expand after single landing works |
| Contact form backend | Connector / BFF action later |

Revisit for `services-website-v1` template.

## Preview behaviors (L2)

| Block | Preview interactions |
|-------|---------------------|
| `hero` | CTA links navigate or scroll |
| `features_grid` | Renders empty state if no items |
| `pricing` | Highlight tier styling from field |
| `faq` | Accordion expand/collapse |
| `testimonials` | Carousel optional later |

## Generated artifacts (L3)

Expected paths (illustrative):

```text
internal/views/landing.templ
internal/ui/blocks/marketing/*.templ
internal/appschema/landing_content_types.go   # or yaml/json seed
internal/fixtures/seed/landing_en.json
cmd/server/main.go                            # route registration
```

Each file tagged with origin: `schema` | `templ` | `seed`.

## Template metadata (fixture)

For BuildY template gallery card:

```json
{
  "id": "saas-landing-v1",
  "title": "SaaS Product Landing",
  "subtitle": "CMS-backed conversion page with pricing and FAQ",
  "category": "saas",
  "blocks": ["hero", "features_grid", "pricing", "testimonials", "faq", "cta_banner"],
  "locales": ["en", "ru"]
}
```

## Acceptance criteria

- [ ] All MVP blocks render in preview with seed content  
- [ ] Empty states visible when content entries removed  
- [ ] SEO fields appear in page `<head>` via shell partial  
- [ ] Export snapshot includes content type definitions + seed  
- [ ] `bun run verify` + render tests pass on landing view  
- [ ] Non-developer can change hero headline in preview store  

## Related

- [`../capabilities/INDEX.md`](../capabilities/INDEX.md)
- [`../messaging/homepage-ia.md`](../messaging/homepage-ia.md)
- [`../../architecture/03-cms-backed-sync.md`](../../architecture/03-cms-backed-sync.md)
