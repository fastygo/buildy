# End User: Content Marketer

## Profile

| Field | Detail |
|-------|--------|
| Archetype | Person who owns messaging, copy, and landing performance |
| Typical titles | Content Marketer, Growth Marketer, Marketing Ops, Copywriter |
| Technical level | Low — expects CMS-like editing, not Git |
| Session length | 15–45 minutes per edit pass; frequent short visits |
| Success feeling | Published copy matches brand without opening a ticket |

## Primary workflow

```text
Open project preview or live AppCMS admin
    → edit hero/pricing/FAQ/testimonial entries
    → preview changes → publish or sync to live
    → monitor (external analytics) → iterate blocks
```

## Tasks in BuildY (MVP → live)

| Phase | Where they work |
|-------|-----------------|
| Pre-launch | BuildY preview store — edit block content entries |
| Post-export | AppCMS admin — primary home for content |
| Coordinated launch | BuildY content sync → live AppCMS (later) |

## Content objects they touch

- Hero headline, subhead, CTA labels  
- Feature grid titles and descriptions  
- Pricing tiers, bullets, highlight flags  
- Testimonials (quote, name, role, avatar)  
- FAQ pairs  
- SEO title, description, OG fields  
- Locale variants (EN/RU) where enabled  

## Motivations

- **Autonomy** from engineering for copy and campaign swaps  
- **Safety** — preview before public  
- **Consistency** — blocks enforce layout; they control words  
- **Speed** for launches and regulatory copy updates  

## Frustrations to avoid

- Editing requires understanding Go/Templ files  
- Unstructured markdown blobs instead of typed fields  
- No preview of SEO/social cards  
- Breaking layout when changing text length (need sensible block variants)  

## Relationship to other personas

- **Prototype author** sets initial structure; marketer owns ongoing content  
- **Go developer** implements new block types when marketing needs new patterns  
- **Marketing lead buyer** cares that this persona is unblocked  

## Onboarding emphasis

Landing workflows should stress **field-level CMS** in help and first-run — not one-shot generated copy.

## Related

- Buyer: [`../buyer-persona/marketing-lead.md`](../buyer-persona/marketing-lead.md)
- ICP: [`../icp/icp-founder-led-saas.md`](../icp/icp-founder-led-saas.md)
