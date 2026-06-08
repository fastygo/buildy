# Buyer Persona: Product Manager

## Profile

| Field | Detail |
|-------|--------|
| Title | Product Manager, Group PM, Product Owner |
| Company | Product-led SaaS or platform team on FastyGo |
| Reports to | VP Product, CPO, or founder |
| Tools today | Notion/Linear PRDs, Figma, Miro, Slack, sometimes SQL/Amplitude |
| BuildY role | Champion; drives team trial; heavy end user |

## Goals

- Answer **"should we build this?"** with interactive evidence before roadmap lock.
- Reduce spec ambiguity — engineers build from **schema + preview**, not prose alone.
- Align design, eng, and stakeholders on **edge cases and content structure** early.
- Avoid being blocked by eng capacity for every prototype iteration.

## Pain points

- Static prototypes don't capture **data states, permissions, CMS gaps**.
- Eng treats Figma prototypes as "not real" — rework after sprint start.
- Internal tool requests pile up; each becomes a custom eng project.
- Generic AI-builder output doesn't map to company's Go/AppCMS production path.

## BuildY value proposition

> "The missing middle between wireframes and production: CMS-backed pages, BFF preview, and generated Templ — so engineering extends instead of restarts."

## Decision criteria

1. Fidelity of preview (navigation, CRUD, content empty states)  
2. Generated artifacts understandable by eng lead  
3. Time saved vs eng prototype spikes  
4. Fit with AppCMS content workflow  
5. Ability to share preview link with stakeholders  

## Triggers to buy

- Major landing or onboarding redesign in next quarter  
- Recurring eng ask: "can someone prototype this flow first?"  
- Platform team standardizing on FastyGo stack  

## Jobs in evaluation

1. Import or describe PRD / screenshots  
2. Generate landing or flow with block + field mapping  
3. Share preview for stakeholder review  
4. Export file tree for eng estimation  
5. Track diffs when model changes  

## MVP focus

**CMS-backed landing + feature page prototypes** first; metrics dashboard and roadmap hub patterns are **later** expansions.

## Related

- ICP: [`../icp/icp-fastygo-product-team.md`](../icp/icp-fastygo-product-team.md)
- End user: [`../end-user/prototype-author.md`](../end-user/prototype-author.md)
