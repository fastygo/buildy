# Preview Workspace

## Definition

The BuildY preview workspace is a behavior-aware product prototyping surface.

It may show generated source files, file trees, diffs, logs, and rendered pages,
but it is not a full online IDE and not a general-purpose code sandbox.

## Why This Matters

Go is compiled, and FastyGo's product philosophy is schema-driven:

```text
schemas + patterns + BFF models + Templ rendering
```

The MVP should avoid executing arbitrary user Go code in the browser workflow.
Arbitrary code execution requires isolation, build workers, resource limits, secret
boundaries, and operational complexity. That belongs later.

## Source Code vs Preview

| Concern | Source code | Preview workspace |
|---------|-------------|-------------------|
| Owner | GitHub/local IDE | BuildY |
| Purpose | Long-term implementation | Product design and validation |
| Editing | Developer tools | Structured model/pattern changes |
| Execution | Build/test/deploy pipeline | BFF/Templ preview runtime |
| Risk | Normal repository risk | Controlled preview behavior |
| User expectation | Real codebase | Prototype and generated artifacts |

## Preview Levels

### L1 — Schema Preview

Fastest feedback loop.

```text
project schema -> Platform BFF ScreenModel -> Templ render
```

Use for:

- page structure;
- sections and blocks;
- table/form layout;
- navigation;
- basic theme checks.

No generated source build is required.

### L2 — Behavior Preview

Adds live product behavior through BFF actions and preview storage.

```text
BFF actions -> preview store -> AppCMS-compatible content -> Templ preview
```

Use for:

- CRUD flows;
- form validation;
- list filtering;
- content editing;
- connector simulations;
- preview data reset/seed.

This is the most important BuildY differentiator: the prototype is not static.

### L3 — Source Preview

Shows generated source artifacts without turning the browser into the source editor.

Features:

- generated file tree;
- readonly file content;
- diff view;
- changed artifact list;
- apply/revert generated changes at the model level;
- export snapshot.

The source preview should answer: "What will BuildY generate?" not "Can I edit all
code here forever?"

### L4 — Build Preview

Later stage.

```text
generated repo -> templ generate -> go test/build -> preview binary
```

Use for:

- compile validation;
- generated app smoke tests;
- deploy artifact creation;
- release candidates.

Requires build workers and stronger isolation. Do not start here.

## Preview Store

The preview store is the draft runtime database for BuildY projects.

Initial options:

- in-memory for very early tests;
- SQLite per project for MVP;
- snapshot files for export/import;
- live AppCMS sync later.

The preview store should hold:

- content entries;
- content type definitions;
- page/block instances;
- media references;
- settings;
- connector mock results;
- draft/published state.

## Generated File Viewer

The file viewer should show:

- path;
- language/type;
- generated content;
- origin stage (`schema`, `templ`, `module`, `seed`, `test`);
- whether the file is new, changed, or deleted;
- links to related schema objects.

Do not make this the primary editor. Editing should happen through structured
models first, then GitHub/local IDE after export.

## Diff Model

BuildY should track generated changes as artifacts:

```text
GeneratedChange
  project_id
  run_id
  path
  operation: create | update | delete
  before_hash
  after_hash
  content
  source_model_refs
```

This allows review, rollback, export, and future PR generation.

## Preview Routes

Suggested early routes:

```text
/builder/projects
/builder/projects/{projectID}
/builder/projects/{projectID}/preview
/builder/projects/{projectID}/files
/builder/projects/{projectID}/runs/{runID}
/builder/projects/{projectID}/content
```

The preview iframe should render through a product-owned route that uses BFF and
Templ, not through a generic frontend dev server.

## Live Updates

Start simple:

- refresh preview after generation;
- poll run status;
- show logs after each step.

Later:

- SSE for run logs;
- preview refresh events;
- connector test streaming.

Do not require WebSocket/realtime for the MVP.

## Security Notes

MVP preview is safe because it avoids arbitrary user code execution.

Later build previews must isolate:

- filesystem;
- network;
- environment variables;
- secrets;
- CPU/memory/time;
- generated binaries.

Possible future isolation strategies:

- Docker workers;
- short-lived VMs;
- Firecracker;
- restricted remote build queue;
- WASM for selected plugin-like execution.

## MVP Rule

Start with L1, L2, and L3. Add L4 only after CMS-backed landing generation is proven.
