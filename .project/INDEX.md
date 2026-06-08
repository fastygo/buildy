# BuildY Project Index

Status: strategic reference  
Created: 2026-06-07  
Scope: BuildY as a FastyGo product prototyping workspace and CMS-backed landing/app builder.

## Purpose

This directory is the working memory for BuildY. It should let future chats and
agents recover the product direction without drifting into a generic online IDE,
generic no-code builder, or React/Next clone.

BuildY is a product app on the FastyGo stack:

```text
fastygo/framework + fastygo/platform + fastygo/templ + AppCMS
```

It helps users design and prototype real FastyGo applications in the browser, then
continue development in GitHub and a local IDE.

## Core Principle

BuildY is not a browser IDE.

BuildY is a **prototype workspace**:

- it shows generated source trees and file contents;
- it previews real BFF-backed behavior;
- it supports CRUD through AppCMS/BFF preview storage;
- it generates deployable Go/Templ/AppCMS source code;
- it syncs with GitHub as the source-code owner;
- it treats local IDE development as the normal continuation path.

## Documents

### Architecture

- `architecture/00-north-star.md` — product position, principles, and non-goals.
- `architecture/01-system-map.md` — main subsystems and ownership boundaries.
- `architecture/02-preview-workspace.md` — preview workspace, generated files, and behavior preview.
- `architecture/03-cms-backed-sync.md` — AppCMS integration, content/schema/code sync.

### Planning

- `planning/roadmap.md` — implementation phases and exit criteria.

### Registries

- `registries/decision-register.md` — accepted decisions and open questions.

### Progress

- `progress.md` — current state, next actions, and focus reminders.
- `progress-ux-mockup-fixtures.md` — UX/UI mockup fixture roadmap and workspace focus.

## Reference Stack

BuildY should reuse sibling packages deliberately:

| Layer | Role |
|-------|------|
| `@Framework` | HTTP host, sessions, middleware, runtime safety |
| `@Platform` | BFF contract, module descriptors, panel/toolset, renderer-agnostic models |
| `@Templ` | primary renderer and component authoring layer |
| `@AppCMS` | CMS-backed content model, admin, REST/codex, public content runtime |
| `@BuildY` | product UX, project lifecycle, generation pipeline, preview workspace, sync, billing |

BuildY must not move AppCMS domain logic into Platform or duplicate CMS behavior
inside the builder. It should orchestrate AppCMS-compatible modules, schemas,
content snapshots, and generated source.

## Working Vocabulary

| Term | Meaning |
|------|---------|
| Prototype workspace | Browser UI for prompts, schemas, generated files, diffs, and preview |
| Preview workspace | Runtime view of BFF-backed behavior, not a full code editor |
| Source code | Real Go/Templ project files owned by GitHub/local IDE |
| CMS-backed landing | Landing page whose blocks map to AppCMS content types and fields |
| Preview store | Isolated draft content/storage for BuildY previews |
| Live app | Deployed generated app with AppCMS runtime and public pages |
| Code sync | BuildY generated source to GitHub and GitHub diffs back to BuildY |
| Content sync | BuildY preview content to/from live AppCMS content |
| Schema sync | BuildY model/content types to AppCMS migrations or schema diffs |

## Focus Rule

BuildY should first prove:

```text
prompt -> schema/pattern -> AppCMS-backed landing -> BFF/Templ preview -> generated source
```

Do not start with billing, GitHub App complexity, arbitrary code execution, or
isolated build sandboxes. Those are later product layers.
