# ICP: Product Team on FastyGo

## Snapshot

| Field | Value |
|-------|-------|
| Company type | Product-led SaaS, internal platform team, or FastyGo early adopter |
| Team size | 5–40; 2–10 engineers familiar with Go |
| Geography | Global; may operate EN/RU or multi-locale products |
| Budget | Team plan; values toolchain consolidation |
| Technical maturity | Uses or evaluates `framework + platform + templ + AppCMS` |

## Company profile

An organization already committed — or strongly leaning — toward the FastyGo product stack. Product, design, and engineering share a goal: **ship features faster** without duplicating CMS logic or spawning throwaway React prototypes that never merge into production.

They treat AppCMS as the content runtime and want marketing and product surfaces to stay compatible with admin, codex, and public rendering contracts.

## Jobs to be done

1. Prototype **new pages, blocks, and content types** before sprint commitment.
2. Align PM, design, and engineering on **schema + preview** instead of lengthy spec docs.
3. Generate **Templ UI and BFF actions** that engineers refine rather than rewrite.
4. Sync preview content with **live AppCMS** when ready to ship.

## Why BuildY wins

- Native fit: generated artifacts match **Platform BFF + Templ + AppCMS** boundaries.
- Preview workspace is **behavior-aware** (CRUD, validation, navigation) via preview store.
- Reinforces `.project` guardrail: BuildY orchestrates, Platform stays generic.
- Engineers trust output because it follows **existing repo patterns** (Blank/AppCMS reference).

## Why BuildY loses

- Team wants prototyping inside **Figma only** with no codegen step.
- Platform/AppCMS not adopted and no appetite to learn Go stack.
- Expectation of full IDE in browser for arbitrary Go edits.

## Typical buying committee

| Role | Involvement |
|------|-------------|
| VP Product / Head of Product | Buyer or sponsor |
| Engineering lead / architect | Technical gatekeeper |
| Product manager | Champion, daily end user |
| Design lead | Contributor, preview reviewer |

## Success metrics

- Reduced cycle from idea doc to reviewable prototype  
- Fewer "prototype rewrite" sprints  
- Generated files merge to main with acceptable diff size  
- Content types map cleanly to AppCMS migrations  

## Market framing

**Schema-first FastyGo prototyping** for PM and design audiences — internal tools remain a later expansion.

## Related personas

- Buyers: [`../buyer-persona/product-manager.md`](../buyer-persona/product-manager.md), [`../buyer-persona/engineering-tech-lead.md`](../buyer-persona/engineering-tech-lead.md)
- End users: [`../end-user/prototype-author.md`](../end-user/prototype-author.md), [`../end-user/go-developer.md`](../end-user/go-developer.md)
