# Buyer Persona — Overview

Buyer personas describe **people who approve budget, tool adoption, or stack direction** — not everyone who clicks inside BuildY.

## Persona map

| Persona | Primary ICP | Buying role |
|---------|-------------|-------------|
| [Founder / CEO](founder-ceo.md) | Founder-led SaaS | Economic buyer, self-serve |
| [Product Manager](product-manager.md) | FastyGo product team | Champion, team plan influencer |
| [Marketing Lead](marketing-lead.md) | Founder SaaS, agency | Champion for landing/CMS |
| [Engineering / Tech Lead](engineering-tech-lead.md) | FastyGo product team | Gatekeeper, handoff owner |
| [Agency Principal](agency-principal.md) | Digital agency | Economic buyer, methodology owner |

## Shared buyer concerns

1. **Will this become throwaway work?** → Generated Go/Templ + export answers yes/no for engineering.
2. **Who owns the code?** → GitHub / client repo; BuildY is workspace, not prison.
3. **Can non-devs update the site?** → AppCMS content model + admin.
4. **Does it fit our stack?** → FastyGo only; disqualify React-first buyers early.
5. **Security and lock-in?** → Source visible, no arbitrary cloud execution in MVP (expand later).

## Evaluation journey

```text
Awareness (prompt-builder category — speed story)
    → Qualification (FastyGo / CMS / export needs)
    → Trial (first CMS-backed landing project)
    → Technical review (eng lead reads generated files)
    → Purchase / expand (seats, projects, sync)
```

## Objection handling themes

| Objection | Response angle |
|-----------|----------------|
| "We already use Lovable / Bolt / v0 / Base44 / Cursor" | BuildY output is CMS-backed FastyGo app with export — different stack and handoff |
| "We need deploy now" | MVP proves preview + export; deploy is AppCMS/FastyGo hosting path |
| "PMs shouldn't generate code" | They edit schemas and content; eng owns repo merge |
| "Too niche stack" | ICP is teams *choosing* Go/Templ/AppCMS for ownership and CMS |

## Related docs

- [`../icp/overview.md`](../icp/overview.md)
- [`../end-user/overview.md`](../end-user/overview.md)
