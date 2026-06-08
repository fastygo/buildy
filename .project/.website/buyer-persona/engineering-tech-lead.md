# Buyer Persona: Engineering / Tech Lead

## Profile

| Field | Detail |
|-------|--------|
| Title | Tech Lead, Staff Engineer, Architect, VP Engineering |
| Company | FastyGo adopter or Go-friendly SaaS |
| Reports to | CTO or founder |
| Tools today | GitHub, Go toolchain, IDE, AppCMS admin, CI |
| BuildY role | **Gatekeeper** — can block or accelerate adoption |

## Goals

- Protect **architecture boundaries** (Framework, Platform, AppCMS, Templ).
- Ensure prototypes **merge into production** with bounded diff, not rewrite.
- Avoid duplicate CMS logic or Platform contamination with BuildY-specific code.
- Maintain **security and review** — no unreviewed arbitrary code execution.

## Pain points

- PM prototypes in React that eng refuses to merge.
- No-code tools create **unmaintainable** output; eng becomes cleanup crew.
- Fear of "another codegen toy" that ignores BFF contracts and migrations.
- Time wasted re-implementing landing pages that marketing already "approved" in Figma.

## BuildY value proposition

> "Codegen that speaks your stack: Go/Templ views, Platform BFF patterns, AppCMS content types. Review files and diffs in GitHub — BuildY never replaces IDE."

## Decision criteria

1. Generated code readability and idiomatic Templ  
2. No improper imports into Platform from BuildY product logic  
3. Clear separation: preview store vs live AppCMS  
4. GitHub sync and diff workflow  
5. Schema/migration story for content types  

## Triggers to approve

- Explicit north-star alignment with existing monorepo layout  
- Successful export reviewed in PR with acceptable changes  
- PM/design iteration happens in BuildY, not in production branch  

## Red lines (will reject)

- Browser IDE expectation for all Go editing  
- Demand for React/Next output as primary  
- Arbitrary user Go execution in cloud without sandbox policy  

## Evaluation checklist

- [ ] Generated `internal/views` uses templ/ui primitives  
- [ ] Content types map to AppCMS-compatible models  
- [ ] BFF actions follow Platform panel/toolset boundaries  
- [ ] No secrets or env baked into generated tree  
- [ ] Handoff doc: how to run, test, deploy  

## Related

- ICP: [`../icp/icp-fastygo-product-team.md`](../icp/icp-fastygo-product-team.md)
- End user: [`../end-user/go-developer.md`](../end-user/go-developer.md)
