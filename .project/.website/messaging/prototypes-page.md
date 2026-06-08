# BuildY Prototypes — Product Page

Status: marketing + product reference  
BuildY anchor: **proof of concept in hours** with **schema + BFF preview + GitHub handoff**.

## Page goal

Speak to **product managers, designers, founders** evaluating BuildY for pre-roadmap validation — emphasize that disposable throwaway prototypes are the problem BuildY avoids.

## Hero

**Eyebrow:** BuildY for prototyping

**Headline:** Proof of concept in hours, not weeks.

**Subhead:** Turn PRDs, briefs, and landing ideas into CMS-backed prototypes with real structure — then export Go/Templ source your engineering team can extend.

**CTA:** Get started → create project from SaaS landing template

## Section: Make your PRD real

**Message:** Describe the product or paste a PRD section. BuildY proposes pages, blocks, content types, and fields — then renders a working preview.

| Typical AI builder | BuildY |
|--------------------|--------|
| Connector pull from docs (later) | MVP: paste / prompt |
| Chat-only UI preview | Schema + L2 behavior preview via BFF |
| Pixel/visual-only edits | Block/content edits + structured model panel |

**Capability:** prompt → `Page` / `Block` / `ContentTypeMapping` generation.

## Section: Four beats

### Skip the eng backlog

Product and design iterate on **structure and content** in the prototype workspace without sprint allocation for throwaway spikes.

**BuildY detail:** edits go through schema and preview store — not production branch.

### Show polished structure, not wireframes

Themes, block variants, and Templ-rendered sections give stakeholders **production-shaped** feedback — hero, pricing tables, FAQ accordions.

**BuildY detail:** uses `templ/ui` patterns and ui8px policy — eng lead recognizes output.

### Take it for a spin

Share preview routes. Test navigation, empty content states, and block variants. Validate assumptions before merge.

**BuildY detail:** L2 preview — real BFF actions, preview store CRUD, not clickable Figma.

### Handoff to dev

Export generated tree. Engineers review diffs in GitHub, run locally, merge, deploy AppCMS runtime.

**BuildY detail:** handoff artifact is **Go/Templ + content model**, not a throwaway frontend spike.

## Section: The missing middle

Quote-style callout:

> BuildY sits between static prototypes and production: a living preview tied to a content model and exportable FastyGo source — not a demo that gets thrown away.

## Section: What you prototype first (MVP)

| In scope now | Later |
|--------------|-------|
| SaaS marketing landing | Internal dashboards |
| Block/content model edits | Full auth/billing modules |
| Preview store content | Live AppCMS bidirectional sync |
| Generated file review | GitHub PR automation |

Aligns with [`../../progress.md`](../../progress.md) first proof.

## FAQ (BuildY answers)

### How do prototypes become production code?

BuildY generates a **deployable FastyGo project skeleton**. After stakeholder sign-off, export to GitHub. Engineers add auth, integrations, tests, and deploy with AppCMS. The landing content model already matches admin fields.

### Do we need engineering to build prototypes?

No for **landing structure and preview content**. Yes for **merge, deploy, and product depth** — but eng receives schema-aware source, not a blank rewrite spec.

### Can we use real data in prototypes?

MVP: **preview store** with seed entries and CRUD. Later: connector pull from Notion/Linear; sync with live AppCMS content.

### What if generation goes wrong?

Use **PromptRun history**: preview prior run, revert model-level changes, iterate with refined prompts. See [`../workspace/version-history-ux.md`](../workspace/version-history-ux.md).

### Is BuildY a browser IDE?

**No.** Structured model editing + readonly source view. Long-form code work happens in GitHub/local IDE.

### How is this different from other AI builders?

See [`../competitive/comparison-framework.md`](../competitive/comparison-framework.md) — CMS-native, Go/Templ output, AppCMS runtime, export ownership.

## Proof points (when available)

- Export diff size / merge rate  
- Time from prompt to previewable landing  
- Content edits without developer (post-export AppCMS)  

## CTA block

Repeat: `Ask BuildY to create a SaaS landing for…` + Get started.

## Related

- [`homepage-ia.md`](homepage-ia.md)
- [`../end-user/prototype-author.md`](../end-user/prototype-author.md)
- [`../../architecture/02-preview-workspace.md`](../../architecture/02-preview-workspace.md)
