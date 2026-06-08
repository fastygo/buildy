# Target Market — Overview

Target market defines **which segments BuildY prioritizes** — TAM logic, sequencing, and what we defer.

## Market definition

BuildY competes in the intersection of:

1. **AI-assisted product builders** (Lovable, Bolt, v0, Base44, Cursor agents)  
2. **Headless CMS + static/SSR site generators** (Contentful + Next, etc.)  
3. **Go/full-stack app platforms** (internal frameworks, AppCMS ecosystem)  

BuildY's wedge: **#1 speed** with **#3 stack ownership** and **#2 content model** baked in from prototype — not bolted on after launch.

## Segment priority matrix

| Segment | Priority | MVP fit | Doc |
|---------|----------|---------|-----|
| SaaS landing & marketing sites | **P0** | Core first proof | [`segments-primary.md`](segments-primary.md) |
| FastyGo product teams | **P0** | Stack-native adopters | [`segments-primary.md`](segments-primary.md) |
| Agency client delivery | **P1** | Templates + export | [`segments-secondary.md`](segments-secondary.md) |
| Internal tools & dashboards | **P2** | Post-landing expansion | [`segments-secondary.md`](segments-secondary.md) |
| Full product apps (auth, billing) | **P2** | After landing + export proven | [`segments-secondary.md`](segments-secondary.md) |
| Enterprise multi-team | **P3** | SSO, audit, isolation — later | [`segments-deferred.md`](segments-deferred.md) |

## Geographic and language focus

- **Primary:** global English-first SaaS and product teams  
- **Secondary:** bilingual products (EN/RU) — aligns with BuildY fixture/locale mockup work  
- **Defer:** regulated verticals requiring compliance certifications before sell  

## Competitive frame

| Alternative | BuildY differentiation |
|-------------|------------------------|
| AI prompt builders (Lovable, Bolt, v0, Base44) | FastyGo stack, CMS-native, Go export, license-clear handoff |
| Webflow / Framer | Real app + admin, not page-only |
| Raw Cursor/Claude in IDE | Structured schema, preview workspace, CMS sync |
| AppCMS alone | Generation + prototype speed from prompts/templates |

## Sizing heuristic (qualitative)

**Serviceable obtainable market (early):** teams already building or evaluating FastyGo + teams frustrated with disposable AI frontend output who need CMS + repo ownership.

Do not optimize messaging for mass non-technical "vibe code anything" market until deploy/billing layers exist.

## Go-to-market sequence

```text
1. CMS-backed SaaS landing (prove loop)
2. FastyGo community + agencies (stack believers)
3. Internal admin / dashboard patterns (reuse blocks)
4. GitHub sync, team seats, billing
5. Enterprise and marketplace connectors
```

## Related docs

- [`../target-audience/overview.md`](../target-audience/overview.md)
- [`../icp/overview.md`](../icp/overview.md)
- [`../../planning/roadmap.md`](../../planning/roadmap.md)
