# BuildY Capabilities — Index

Status: product reference  
Scope: what BuildY offers, mapped to user journey, MVP slice, and architecture.

## Capability map

```text
Describe  →  Model  →  Preview  →  Review  →  Export  →  Continue in IDE
   │           │          │           │          │              │
 prompt    schema/     L1–L3      generated   GitHub/      AppCMS live
           blocks      preview      files      snapshot      + eng
```

## Documents

| Doc | Capability area |
|-----|-----------------|
| [`../messaging/homepage-ia.md`](../messaging/homepage-ia.md) | Public homepage — how capabilities are introduced |
| [`../messaging/prototypes-page.md`](../messaging/prototypes-page.md) | Prototyping story — PRD to handoff |
| [`../templates/saas-landing-blocks.md`](../templates/saas-landing-blocks.md) | First template — CMS blocks and content types |
| [`../workspace/version-history-ux.md`](../workspace/version-history-ux.md) | PromptRun history, revert, performance |
| [`../competitive/comparison-framework.md`](../competitive/comparison-framework.md) | Positioning vs AI builders and no-code |

## Core capabilities (summary)

### 1. Project lifecycle

| Capability | Description | MVP |
|------------|-------------|-----|
| Create project | Name, template, theme baseline | ✓ |
| Project list / detail | Workspace entry, metadata, status | ✓ |
| Template gallery | SaaS landing starter (expand later) | ✓ (one template) |

### 2. Prompt → model

| Capability | Description | MVP |
|------------|-------------|-----|
| Natural language input | Describe product, landing, audience | ✓ |
| Schema generation | Pages, blocks, content types, field mappings | ✓ |
| Pattern library | Reusable block/section patterns from Templ staging | ✓ (landing set) |
| Context import | PRD paste; connectors (Notion/Linear) | Later |

Domain objects: `Project`, `Page`, `Block`, `ContentTypeMapping`, `FieldMapping`, `Theme`, `PromptRun`.

### 3. Preview workspace (not IDE)

| Level | Capability | MVP |
|-------|------------|-----|
| **L1 Schema** | Page/block layout, nav, theme | ✓ |
| **L2 Behavior** | BFF CRUD, validation, preview store content | ✓ |
| **L3 Source** | Readonly file tree, diffs, artifact list | ✓ |
| **L4 Build** | `go test/build` in isolated worker | Later |

See [`../../architecture/02-preview-workspace.md`](../../architecture/02-preview-workspace.md).

### 4. CMS-backed content

| Capability | Description | MVP |
|------------|-------------|-----|
| Block → field mapping | Hero, features, pricing, FAQ, etc. map to AppCMS types | ✓ |
| Preview content CRUD | Edit entries in preview store | ✓ |
| SEO / settings fields | Meta title, description, OG | ✓ |
| Locale-ready copy | EN/RU via content fields (fixture pattern) | ✓ (mockup) |
| Live AppCMS sync | Preview ↔ deployed content | Later |

### 5. Generated source

| Capability | Description | MVP |
|------------|-------------|-----|
| File tree viewer | Go/Templ paths, origin stage labels | ✓ |
| Diff / change list | `GeneratedChange` per PromptRun | ✓ |
| Export snapshot | Downloadable source bundle | ✓ |
| GitHub sync | Push branch, open PR | Later |

Output stack: `framework + platform + templ + AppCMS-compatible modules`.

### 6. Version history

| Capability | Description | MVP |
|------------|-------------|-----|
| PromptRun log | Labeled runs with preview-before-revert | ✓ |
| Recent runs quick access | Last N edits | ✓ |
| Lazy history UI | Scroll-loaded run list | ✓ (when needed) |
| Model-level rollback | Revert schema/content to run | ✓ |

See [`../workspace/version-history-ux.md`](../workspace/version-history-ux.md).

### 7. Handoff & continuity

| Capability | Description | MVP |
|------------|-------------|-----|
| Eng-readable codegen | Templ/ui primitives, ui8px policy | ✓ |
| GitHub as source owner | Export; IDE is continuation path | ✓ |
| AppCMS admin post-deploy | Marketing edits without codegen | ✓ (via export target) |
| Deploy automation | One-click publish | Later |

## Explicit non-capabilities (current slice)

- Full browser IDE / arbitrary Go editing  
- Arbitrary code execution in cloud  
- React/Next primary output  
- Realtime multi-user co-editing  
- Billing, GitHub App, isolated build workers (until proven loop)  

## First proof checklist

From [`../../progress.md`](../../progress.md):

- [ ] One SaaS landing template with all MVP blocks  
- [ ] BFF/Templ preview with preview store CRUD  
- [ ] Generated files view with diffs  
- [ ] Export snapshot passes `go test ./...`  
- [ ] No BuildY imports into Platform  

## Related

- Audience: [`../target-audience/overview.md`](../target-audience/overview.md)
- North star: [`../../architecture/00-north-star.md`](../../architecture/00-north-star.md)
