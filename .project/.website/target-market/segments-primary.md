# Target Market — Primary Segments (P0)

Segments BuildY must win for **first proof** and initial GTM.

---

## 1. SaaS landing & product marketing sites

### Description

Startups and scale-ups need credible **marketing surfaces**: homepage, pricing, features, social proof, docs entry, waitlist CTA. Often the first "product" investors and customers see.

### Why P0

Matches north-star first proof:

```text
CMS-backed landing → BFF/Templ preview → generated source export
```

### Typical buyer / user

- Buyer: founder, marketing lead  
- End user: prototype author → content marketer post-launch  

### Key requirements

- Block library: hero, features, pricing, testimonials, FAQ, CTA  
- SEO and meta fields in content model  
- Theme variants (dark/light, brand tokens)  
- Export as deployable AppCMS app  

### Comparable market narratives

Founder and marketing audiences often search for **landing pages** and **branded campaign sites** — BuildY answers with CMS structure, not one-shot HTML.

### BuildY win message

Structured landing, not hardcoded AI HTML — editable forever in AppCMS.

---

## 2. FastyGo product teams

### Description

Teams building on `framework + platform + templ + AppCMS` who need **faster iteration** on pages and content models without abandoning stack conventions.

### Why P0

BuildY is a **FastyGo product app**, not a generic builder. Early adopters come from ecosystem gravity and Blank/AppCMS reference patterns.

### Typical buyer / user

- Buyer: eng lead + VP product  
- End user: PM prototype author, Go developer  

### Key requirements

- Generated code passes ui8px / templ spec / render tests  
- BFF preview uses Platform contracts  
- No Platform import of BuildY internals  
- Schema sync story toward AppCMS migrations  

### Comparable market narratives

PM and design buyers expect **prototype-before-roadmap** and **behavior over mockups** — BuildY delivers via schema + BFF preview.

### BuildY win message

The only AI prototype workspace whose output **is** your production architecture.

---

## Segment overlap

Many founder-led SaaS companies are **simultaneously** segment 1 and future segment 2 as they adopt FastyGo for the full product. Messaging can lead with landing pain, expand to full app skeleton.

## Success metrics (P0)

- ≥1 complete landing template shipped in BuildY fixtures  
- Export passes `go test ./...` and ui8px validation  
- Non-developer edits block content in preview  
- Eng lead approves sample PR from export  

## Related

- ICP: [`../icp/icp-founder-led-saas.md`](../icp/icp-founder-led-saas.md), [`../icp/icp-fastygo-product-team.md`](../icp/icp-fastygo-product-team.md)
