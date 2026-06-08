# BuildY North Star

## Product Thesis

BuildY is an AI-assisted product prototyping workspace for FastyGo applications.

It generates real Go/Templ/AppCMS applications from schemas, patterns, prompts,
and connectors. The browser experience is optimized for product design, preview,
content modeling, and generated-source review. It is not intended to replace a
developer's GitHub repository or local IDE.

## Positioning

BuildY should offer a fast prompt-to-preview flow familiar from AI builders, but
it must not adopt a React/Next/shadcn/Supabase deliverable or a hosted clone
architecture. BuildY is a distinct product: prototype workspace, exportable
FastyGo source, and licensing-friendly code ownership — not a white-label of
any single competitor.

BuildY's result is not a static landing page with hardcoded content. The result is:

```text
deployable FastyGo source code
AppCMS-backed content model
Templ UI
BFF-backed preview
GitHub-owned source
CMS-editable public pages
```

## Stack Identity

BuildY uses the FastyGo stack as its default product platform:

```text
Framework -> Platform BFF -> AppCMS -> Templ -> BuildY preview workspace
```

The browser preview is powered by schemas, patterns, AppCMS content models, and
BFF actions. It does not need to execute arbitrary user Go code for the MVP.

## Core Principles

1. **Source code and preview are separate.** BuildY can show a file tree and file
   contents, but GitHub/local IDE remain the real source-code workflow.
2. **Preview is behavior-aware.** A prototype can run CRUD, validation, navigation,
   content rendering, and connector simulations through BFF and preview storage.
3. **AppCMS is a product asset.** Generated landings and sites are backed by CMS
   content types, entries, settings, media, and public rendering contracts.
4. **Templ is primary.** Generated UI should be Go/Templ-first, not Next-first.
5. **Platform remains generic.** BuildY-specific logic belongs in BuildY, not in
   `fastygo/platform`.
6. **Schemas beat arbitrary code.** Most UX/UI should be generated from structured
   schemas, reusable patterns, and typed BFF contracts.
7. **GitHub owns source continuity.** BuildY can generate, sync, and open diffs or
   pull requests, but long-term development continues in the repository.

## Non-Goals

BuildY is not:

- a full online IDE;
- a Codespaces replacement;
- a general arbitrary-code execution platform;
- a React/Next/shadcn clone;
- a static landing generator with hardcoded copy;
- a CMS fork;
- a reason to put product-specific builder logic into Platform.

## Target User Journey

1. User describes a product, site, or landing page.
2. BuildY proposes an app structure: pages, blocks, content types, fields, and theme.
3. User edits the model through a prototype workspace.
4. BuildY renders a live BFF/Templ preview with real CMS-backed content.
5. User reviews generated files and diffs.
6. BuildY exports or syncs the source code to GitHub.
7. The generated app is deployed with AppCMS runtime.
8. Content can be edited in BuildY or in the deployed AppCMS admin and synchronized.
9. Developers continue advanced implementation in GitHub/local IDE.

## Strategic Advantage

Typical AI builders often produce frontend code that later needs a backend. BuildY
should generate a working CMS-backed product skeleton from the beginning:

```text
UX idea -> content model -> BFF prototype -> Templ UI -> deployable AppCMS app
```

This makes the prototype less disposable and closer to a real product.

## First Proof

The first proof should be a CMS-backed landing builder:

- one project;
- one landing page template;
- generated AppCMS content types;
- editable landing block data;
- BFF/Templ preview;
- generated source tree view;
- exportable source snapshot.

Billing, GitHub Apps, cloud deployment, and isolated build workers come later.
