# Target Market — Deferred Segments (P3+)

Explicit **not now** segments to keep GTM and engineering focused. Aligns with `.project/progress.md` "Not Yet" list.

---

## Enterprise multi-team organizations

**Profile:** 500+ FTE, SSO/SAML, audit logs, data residency, procurement.

**Why defer:** MVP has no team RBAC, billing, or compliance packaging.

**Revisit when:** Team seats, org workspaces, audit trail, and support tier exist.

---

## Consumer social & realtime products

**Profile:** Feeds, DMs, live presence, mobile-first social graphs.

**Why defer:** Not aligned with CMS-backed landing first proof; heavy infra.

**Revisit when:** Core loop proven and FastyGo patterns exist for realtime modules (if ever prioritized).

---

## E-commerce & marketplace at scale

**Profile:** Inventory, payments, multi-vendor, complex search.

**Why defer:** Payments, connectors, and worker isolation listed as later layers in north star.

**Revisit when:** Connector marketplace and billing integrations are productized.

---

## "Vibe code anything" mass market

**Profile:** Non-technical users wanting any app type with one-click hosted deploy, no Git concept.

**Why defer:** BuildY requires stack literacy for handoff; browser is prototype workspace, not replacement for all engineering.

**Revisit when:** Never as primary — may remain non-goal per north star.

---

## Isolated build workers & arbitrary code execution

**Profile:** Run user Go code in cloud sandboxes from browser IDE.

**Why defer:** Explicit non-goal for MVP; security and cost complexity.

**Revisit when:** Clear enterprise demand and sandbox architecture approved in decision register.

---

## GitHub App / deploy automation buyers

**Profile:** Buyers who purchase on "push to deploy" and GitHub-native CI alone.

**Why defer:** GitHub App complexity deferred in roadmap; export snapshot first.

**Revisit when:** GitHub sync stable and deploy story documented for AppCMS hosting.

---

## Realtime collaboration

**Profile:** Multiple users editing same prototype simultaneously (Figma-like).

**Why defer:** Listed in progress "Not Yet"; fixture-driven mockups first.

---

## Tracking rule

Before moving a segment from deferred → active:

1. Update [`../../planning/roadmap.md`](../../planning/roadmap.md)  
2. Add decision to [`../../registries/decision-register.md`](../../registries/decision-register.md)  
3. Adjust ICP priority in [`../icp/overview.md`](../icp/overview.md)
