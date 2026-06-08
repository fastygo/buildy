# Target Market — Secondary Segments (P1–P2)

Segments to pursue **after** landing proof or in parallel where reuse is high.

---

## P1: Agency client delivery

### Description

Studios delivering websites and MVPs for clients who need **repo handoff + CMS admin**.

### Entry path

- Reuse SaaS landing templates with client branding  
- Multi-project workspace (future)  
- Export to client-owned GitHub  

### ICP / personas

- [`../icp/icp-digital-agency.md`](../icp/icp-digital-agency.md)  
- [`../buyer-persona/agency-principal.md`](../buyer-persona/agency-principal.md)  

---

## P2: Internal tools & operational dashboards

### Description

PM/ops teams prototyping **admin views, roadmap hubs, metrics dashboards** before eng builds internal tools.

### Why deferred from MVP

North star and progress explicitly prioritize **CMS-backed landing** first. Internal tools share workspace UX but different block patterns and BFF complexity.

### Entry path

- Extend block library (`internal/ui/blocks/dashboard/`)  
- Reuse preview store CRUD patterns from landing work  

### BuildY positioning when ready

"Same workspace — from landing to internal tool — one stack, one export path."

---

## P2: Expanded product apps (beyond landing)

### Description

Founders and product teams moving from landing to **authenticated app shells**: onboarding, settings, basic CRUD modules.

### Scope boundary

Includes auth patterns, module descriptors, AppCMS-backed entities — **not** arbitrary marketplace, billing, or realtime collab in early phases.

### BuildY gate

Ship only after:

1. Landing export stable  
2. GitHub sync proven  
3. Eng merge rate acceptable  

---

## P2: Design-led product exploration

### Description

Designers prototyping permissions, unhappy paths, and multi-step flows tied to schema.

### BuildY angle

Preview workspace shows **real empty/error states** from content model, not Figma overlays.

---

## Segment comparison

| Segment | Reuses landing work? | Eng intensity |
|---------|----------------------|---------------|
| Agency delivery | High (templates) | Medium |
| Internal tools | Medium (workspace) | High |
| Full product app | Low initially | Very high |
| Design exploration | High (preview UX) | Low–medium |

## Related

- [`segments-deferred.md`](segments-deferred.md)
- [`overview.md`](overview.md)
