# Target Audience — Overview

BuildY serves teams that need to move from **product intent** to **deployable FastyGo applications** without treating prototypes as throwaway frontend demos.

## Core audience statement

BuildY is for product-led teams, founders, and delivery partners who want:

- speed comparable to AI app builders;
- output that is **real source code** (Go, Templ, AppCMS), not a locked no-code sandbox;
- **CMS-editable** marketing and product surfaces after launch;
- a clear path from browser prototype to **GitHub + local IDE** continuation.

## Primary audience clusters

### 1. Founders and solo builders

**Need:** validate a SaaS or product landing quickly, keep optionality to hire engineers later.  
**BuildY fit:** generate a CMS-backed landing + app skeleton; export source; deploy on FastyGo stack.  
**Not a fit:** teams that only want a static Webflow-style site with zero code path.

### 2. Product managers and product ops

**Need:** interactive prototypes with real flows before roadmap commitment; artifacts engineering can extend.  
**BuildY fit:** schema-driven pages, BFF-backed preview, generated file review — the "missing middle" between Figma and production.  
**Not a fit:** PMs who only need clickable wireframes with no content model or export.

### 3. Marketing and growth leads

**Need:** ship and iterate landing pages without waiting on engineering for every copy or block change.  
**BuildY fit:** block-level CMS mapping so marketing edits content in preview or live AppCMS, not in code.  
**Not a fit:** teams satisfied with hardcoded AI-generated pages and no structured content.

### 4. Designers and design engineers

**Need:** prototype behavior, states, and permissions — not just static screens.  
**BuildY fit:** preview workspace shows real navigation, content gaps, and block variants tied to schema.  
**Not a fit:** pure brand/design-only workflows with no interest in product structure.

### 5. Engineering leads and platform-minded developers

**Need:** guardrails, readable generated code, no vendor lock-in, alignment with existing Go/Templ/AppCMS investments.  
**BuildY fit:** GitHub-owned source, Templ-first UI, Platform BFF contracts, AppCMS as runtime — not a React fork.  
**Not a fit:** teams standardized on Next.js/Supabase who will not adopt FastyGo.

### 6. Agencies and FastyGo delivery partners

**Need:** repeatable client delivery: landing + admin + export + handoff.  
**BuildY fit:** project templates, CMS-backed patterns, diff review, client-owned repos.  
**Not a fit:** agencies selling only visual design without implementation continuity.

## Shared traits (qualification)

| Trait | Why it matters |
|-------|----------------|
| Values **source ownership** | BuildY assumes GitHub/IDE is the long-term home |
| Thinks in **pages, blocks, content types** | Matches CMS-backed landing model |
| Accepts **Go/Templ** as target stack | Core generated output |
| Wants **preview with behavior** | BFF + preview store, not screenshot mockups |
| Tolerates **structured iteration** | Schemas and patterns over arbitrary code editing in browser |

## Anti-audience (explicit non-fit)

- Teams wanting a full **browser IDE** or arbitrary Go execution in the cloud.
- Orgs that need **React/Next/shadcn** as the primary deliverable.
- Buyers seeking **one-click hosting** with no engineering involvement (MVP defers deploy automation).
- Enterprises requiring **realtime multi-user editing** in v1 (later layer).

## Messaging hierarchy

1. **Outcome:** prototype → preview → export → deployable FastyGo app.  
2. **Differentiator:** CMS-backed from day one, not hardcoded landing copy.  
3. **Continuity:** browser workspace hands off to GitHub; AppCMS runs the live site.  
4. **Speed:** Prompt-to-preview flow without proprietary stack lock-in.

## Related docs

- ICP detail: [`../icp/overview.md`](../icp/overview.md)
- Decision-makers: [`../buyer-persona/overview.md`](../buyer-persona/overview.md)
- Daily users: [`../end-user/overview.md`](../end-user/overview.md)
- Segments: [`../target-market/overview.md`](../target-market/overview.md)
