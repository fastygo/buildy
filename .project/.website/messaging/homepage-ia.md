# BuildY Homepage — Information Architecture

Status: marketing reference  
BuildY anchor: prototype workspace → CMS-backed landing → export, not one-click hosted deploy.

## Page goal

Convert qualified visitors (founders, PMs, marketing leads, Go-friendly teams) to **start first project** — not to imply BuildY replaces GitHub or AppCMS admin long-term.

## Primary message

**Headline option A (speed + stack):**  
Build real FastyGo prototypes from a prompt.

**Headline option B (outcome):**  
From landing idea to CMS-backed source — in the workspace.

**Subhead:**  
Describe your product. Preview pages with real content models. Export Go/Templ/AppCMS source to GitHub. Continue in your IDE.

## Hero

| Element | BuildY content |
|---------|----------------|
| Prompt shell | `Ask BuildY to create a SaaS landing for…` |
| Primary CTA | Get started → project create / template pick |
| Secondary CTA | View templates → SaaS landing gallery (MVP: one template) |
| Trust line | Prototype workspace · CMS-backed · GitHub export |

Do **not** promise: one-click publish, vendor-managed DB auto-provision, arbitrary app types.

## Section: Meet BuildY (3 steps)

Use an export-aware flow (not generic "idea → live → ship"):

### 1. Start with intent

Describe a landing, product surface, or paste a PRD outline. Pick a template (SaaS landing) or start from pattern.

**Capability shown:** prompt input, template gallery entry.

### 2. Preview real structure

See pages, blocks, and CMS field mappings render through BFF/Templ — not a static mockup. Edit content in preview store.

**Capability shown:** L1 schema + L2 behavior preview.

### 3. Review and export

Inspect generated file tree and diffs. Export snapshot or sync to GitHub. Deploy AppCMS app; continue advanced work locally.

**Capability shown:** L3 source preview, export, handoff.

## Section: Discover templates

MVP gallery (expand over time):

| Card | Label | BuildY template ID |
|------|-------|-------------------|
| SaaS landing | Hero, features, pricing, FAQ, CTA | `saas-landing-v1` |
| *(later)* | Services website | deferred |
| *(later)* | Product docs shell | deferred |

Each card: screenshot, 1-line outcome, **Preview** + **Create from template**.

Reference block inventory: [`../templates/saas-landing-blocks.md`](../templates/saas-landing-blocks.md).

## Section: Why BuildY (differentiators)

Three columns:

| Column | Message |
|--------|---------|
| **CMS from day one** | Blocks map to AppCMS content types — marketing edits after export |
| **Behavior preview** | BFF-backed CRUD and navigation, not screenshot prototypes |
| **Source you own** | Go/Templ output in your repo — engineers extend, not restart |

Optional fourth: **FastyGo native** — Framework, Platform BFF, Templ, AppCMS.

## Section: Built for (audience chips)

Link to audience positioning (no separate pages required for MVP):

- Founders → [`../buyer-persona/founder-ceo.md`](../buyer-persona/founder-ceo.md)
- Product managers → [`../buyer-persona/product-manager.md`](../buyer-persona/product-manager.md)
- Marketing → [`../buyer-persona/marketing-lead.md`](../buyer-persona/marketing-lead.md)
- Engineering leads → [`../buyer-persona/engineering-tech-lead.md`](../buyer-persona/engineering-tech-lead.md)

## Section: Social proof (placeholder)

Until BuildY case studies exist:

- FastyGo stack logos / "Built on Framework + Platform + Templ + AppCMS"
- Fixture-led demo quotes (clearly labeled preview/mockup phase)
- Skip inflated vanity metrics from the AI-builder category

## Section: Final CTA

Repeat hero prompt shell + Get started.

Footer: docs link, security/trust (when ready), privacy (legal deferred).

## Nav (marketing site)

```text
Product
  └ Prototypes     → /product/prototypes (see prototypes-page.md)
  └ Templates      → /templates
  └ Security       → /security (later)
Audiences
  └ Founders / PMs / Marketing / Developers
Docs
Pricing (later)
Login / Get started
```

## MVP scope for BuildY app mockup

Home mockup in `@BuildY` fixtures should reflect:

- hero + prompt shell;
- 3-step "Meet BuildY";
- template card for SaaS landing;
- EN/RU locale strings in fixtures;
- dark theme + mobile sheet (per buildy rules).

Workspace routes (`/builder/...`) are separate from marketing home.

## Messaging avoid list

| Common AI-builder trope | BuildY stance |
|-------------------------|---------------|
| "Deploy to the world with one click" | "Export and deploy AppCMS app on your stack" |
| Broad template gallery before first proof | One proven template first |
| "Vibe code anything" | "Prototype FastGo apps with CMS-backed landings" |
| Vendor-scale vanity metrics | Honest early-adopter framing |

## Related

- [`prototypes-page.md`](prototypes-page.md)
- [`../capabilities/INDEX.md`](../capabilities/INDEX.md)
