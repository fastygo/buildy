# BuildY Progress

Last updated: 2026-06-08

## Current Status

BuildY is a new Go module with product architecture documented under `.project`.

Current implementation state:

- `go.mod` exists for `github.com/fastygo/buildy`;
- no product code has been implemented yet;
- `.project` now defines north star, architecture, preview workspace, CMS sync, roadmap, and decisions.

## Current Focus

BuildY should first prove:

```text
CMS-backed landing model -> BFF/Templ preview -> generated source export
```

This means:

- create a product shell;
- define structured project/page/block/content model;
- integrate AppCMS-backed preview content;
- render through Templ and Platform BFF;
- show generated files and diffs;
- export source.

## Completed

### Product Direction

- BuildY positioned as a FastyGo product prototyping workspace.
- Browser experience defined as preview/prototype workspace, not a full IDE.
- AppCMS-backed landing model accepted as the first differentiator.
- Source code ownership assigned to GitHub/local IDE after export.
- Platform remains generic; BuildY-specific logic stays in BuildY.

### Documentation

- `INDEX.md`
- `architecture/00-north-star.md`
- `architecture/01-system-map.md`
- `architecture/02-preview-workspace.md`
- `architecture/03-cms-backed-sync.md`
- `planning/roadmap.md`
- `registries/decision-register.md`

## Next Best Actions

### 1. Bootstrap BuildY App Shell

Create the first app skeleton:

```text
cmd/server
pkg/app
pkg/templ
internal/appschema
internal/domain/project
internal/application/project
internal/storage
```

Use Framework host and Platform BFF patterns from AppCMS.

### Workspace And Dependency Guardrails

Keep BuildY focused on the active workspace only:

- `@BuildY`;
- `@Templ`;
- `@Framework`;
- `@Platform`;
- `@AppCMS`.

Do not add extra repos or module dependencies unless the current UX prototype
proves a concrete need. BuildY should depend on Framework, Platform, Templ, and
AppCMS-facing contracts only when implementation starts.

### 2. Define Project Domain Model

Start with:

```text
Project
Page
Block
ContentTypeMapping
FieldMapping
Theme
PromptRun
PreviewSession
GeneratedFile
```

Keep it simple and serializable.

### 3. Implement First BFF Screens

Minimum screens:

- project list;
- project detail;
- prototype workspace shell;
- generated files view;
- preview screen placeholder.

### 4. Implement CMS-Backed Landing Prototype

Generate one SaaS landing template:

- hero;
- features;
- pricing;
- testimonials;
- FAQ;
- CTA;
- SEO settings.

Map each visible block to CMS-backed structured content.

### 5. Add Preview Store

Use a simple preview storage approach:

- in-memory for first pass, or
- SQLite once the model stabilizes.

The preview store must support content entries and BFF actions for editing.

## Not Yet

Do not start these until the first CMS-backed landing preview works:

- GitHub App;
- Stripe/Paddle billing;
- isolated build workers;
- arbitrary code execution;
- deployment automation;
- realtime collaboration;
- full connector marketplace.

## Verification Targets

Early checks should include:

- `go test ./...` in BuildY;
- BFF screen tests for project/workspace screens;
- generated model serialization round-trip;
- preview content CRUD tests;
- generated source snapshot tests;
- no imports from BuildY into Platform.

## Focus Reminders

- BuildY is not a browser IDE.
- BuildY previews behavior through BFF and AppCMS-compatible content.
- Generated source should be readable and exportable.
- AppCMS is the content runtime, not a reference to copy into BuildY.
- Platform changes should be exceptional and justified by reusable primitives.
- Templ is the primary generated UI layer.
