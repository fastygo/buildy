# Competitive Comparison Framework

Status: positioning reference  
BuildY use: homepage, sales, docs — differentiation on **stack, CMS, handoff, licensing**.

## Comparison thesis

Most AI builders (Lovable, Bolt, v0, Base44, and similar) optimize for **speed to hosted frontend**. BuildY optimizes for **speed to a CMS-backed FastyGo product skeleton** with **exportable source the customer owns** — relevant for prototype validation and downstream licensing.

The right choice depends on:

1. What stack you must own (Go/Templ vs React/Next)  
2. Whether marketing needs **structured CMS** after launch  
3. Whether preview must match **BFF/production architecture**  
4. Where development continues (IDE vs platform-only)  
5. Whether generated output must be **auditable and license-clear** in your repo  

## Master comparison table

| Dimension | Generic no-code (Bubble) | AI prompt builders (Lovable, Bolt, v0, Base44) | IDE agent (Cursor, Claude Code) | **BuildY** |
|-----------|---------------------------|--------------------------------------------------|----------------------------------|------------|
| Primary output | Proprietary visual app | React/TS or hosted app shell | Files in open repo | **Go/Templ/AppCMS source** |
| Code ownership | No export | Varies; often platform-tied | Full | **Full (Go repo)** |
| CMS / content model | Platform DB | Often hardcoded or light CMS | Manual setup | **AppCMS types from day one** |
| Preview type | In-platform | Dev-server-like UI | Local run | **BFF/Templ L1–L3 workspace** |
| Prototype → prod gap | Stay on platform | Frontend strong; backend bolt-on | Engineer-driven | **Schema → preview → export** |
| Learning curve | High (visual logic) | Low start | High (dev skills) | **Low for PM; eng reviews Go** |
| Best for | Non-tech logic apps | Rapid UI, hosted MVPs | Existing codebases | **FastyGo teams, CMS landings** |
| Weak for | Export, custom stack | Go shops, CMS-native | Non-devs, structured CMS gen | React-only shops, no-Go mandate |

## BuildY vs AI prompt builders (category)

Representative category tools: **Lovable, Bolt, v0, Base44** — each differs in hosting, export, and stack; compare individually in sales conversations.

| Feature | Typical AI prompt builder | **BuildY** |
|---------|---------------------------|------------|
| Stack | React/Next + vendor cloud or DB | **Framework + Platform + Templ + AppCMS** |
| Deploy story | One-click / vendor hosting | **Export → AppCMS deploy on your infra** (MVP: export-first) |
| Content after launch | Often code or light CMS | **AppCMS admin fields per block** |
| Preview | AI dev server / hosted preview | **BFF behavior + preview store** |
| Browser role | Near-IDE + visual edits | **Prototype workspace, readonly source** |
| Licensing posture | Vendor terms + generated code mix | **Customer-owned repo; Go/Templ artifacts for review** |

**When a prompt builder wins:** team wants fastest **hosted** MVP on React/Node, minimal Go interest, accepts vendor runtime.

**When BuildY wins:** team chose or evaluates FastyGo; landing must stay CMS-editable; eng must extend same architecture; **export and license clarity** matter.

## BuildY vs v0 / Bolt (UI codegen subset)

| Feature | v0 / Bolt | BuildY |
|---------|-----------|--------|
| Output | Components / Next apps | **Full Go app skeleton + CMS model** |
| Content model | Manual | **Generated with landing** |
| Preview | Component/story | **Product route via BFF** |
| Runtime | Node/Vercel bias | **Go binary + AppCMS** |

**When v0/Bolt wins:** React/Next shop adding UI pieces.

**When BuildY wins:** greenfield **product** with admin and public pages on AppCMS.

## BuildY vs Base44 (hosted app generation)

| Feature | Base44 (representative) | BuildY |
|---------|-------------------------|--------|
| Model | Prompt → hosted app on vendor stack | Prompt → **schema + preview → export** |
| Continuation | Primarily on platform | **GitHub + IDE** |
| CMS depth | App-centric content | **AppCMS content types from generation** |
| Fit | Quick standalone apps | **FastyGo product line + legal handoff** |

Qualify Base44 vs BuildY on **ownership, stack, and post-export workflow** — not feature parity of hosted runtime.

## BuildY vs Cursor / IDE agents

| Feature | Cursor agent | BuildY |
|---------|--------------|--------|
| User | Developer | **PM, founder, marketer + eng review** |
| Structure | Ad hoc files | **Schema, blocks, patterns** |
| Guardrails | Repo rules | **Platform/BFF/Templ/ui8px contracts** |
| Preview | Local | **Shared browser preview link** |

**Complementary, not exclusive:** BuildY generates baseline; Cursor continues in exported repo.

**When Cursor alone wins:** experienced eng, existing repo, no CMS generation need.

**When BuildY wins:** pre-repo prototype with **content model + landing** before eng sprint.

## BuildY vs Bubble (no-code platform)

| Feature | Bubble | BuildY |
|---------|--------|--------|
| Export | No source | **Go/Templ export** |
| Logic | Visual workflows | **BFF + Go (post-handoff)** |
| Lock-in | High | **Low — GitHub owns source** |
| CMS | Bubble DB | **AppCMS** |

**When Bubble wins:** non-technical solo builder, complex workflows, no dev team ever.

**When BuildY wins:** any scenario requiring **real codebase** and CMS-backed marketing site.

## Messaging templates

### One-liner

BuildY is the prototype workspace for **FastyGo** — CMS-backed landings, BFF preview, Go/Templ export you own.

### For PM audience

Generic AI demos give you a frontend spike. BuildY gives you a **content model, preview, and Go source** your team can ship.

### For eng audience

Review generated Templ and AppCMS types in a PR — don't rewrite landing from a Figma link.

### For marketing audience

Change hero and pricing in **AppCMS**, not in a redeploy queue.

## Objection → response

| Objection | Response |
|-----------|----------|
| "We already use Lovable / Bolt / v0 / Base44" | BuildY output is **CMS-backed FastyGo app** with export — different stack and handoff |
| "Hosted builders deploy faster" | BuildY optimizes **handoff quality, CMS, and license-clear source** — deploy on your AppCMS path |
| "We don't use Go" | BuildY is wrong fit; qualify out early |
| "We need mobile app" | Out of MVP scope; web landing + AppCMS first |
| "Can PM edit Go files?" | No — PM edits **schema and content**; eng edits Go in IDE |
| "Is this no-code?" | **Low-code prototype workspace** with full-code export |

## Content formats to produce later

- `/guides/buildy-vs-ai-builders` — category comparison (Lovable, Bolt, v0, Base44)  
- `/guides/buildy-vs-cursor` — complementary positioning  
- Sales one-pager PDF from **Master comparison table**  

## Related

- [`../target-market/overview.md`](../target-market/overview.md)
- [`../capabilities/INDEX.md`](../capabilities/INDEX.md)
- [`../../architecture/00-north-star.md`](../../architecture/00-north-star.md)
