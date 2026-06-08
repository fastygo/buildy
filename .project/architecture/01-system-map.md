# BuildY System Map

## High-Level Architecture

```text
BuildY product app
  project dashboard
  prompt/workflow UX
  schema and block designer
  prototype workspace
  generated file viewer
  sync/export orchestration
        |
        v
Platform BFF contracts
  PageModel
  ScreenModel
  actions
  validation
  navigation/session
        |
        v
AppCMS runtime
  content types
  entries
  media
  settings
  public content
  admin/codex APIs
        |
        v
Templ renderer
  preview shell
  landing/page templates
  generated UI
        |
        v
GitHub/local IDE/deploy
  source ownership
  review
  production build
```

## Subsystems

### 1. Product Workspace

Owns the user-facing builder UX:

- project list;
- project dashboard;
- prompt/chat panel;
- app/page/block tree;
- schema designer;
- connector configuration;
- generated file viewer;
- diff viewer;
- preview iframe;
- run history and logs.

This is BuildY-specific product logic.

### 2. Project Model

Owns the structured source of the generated product:

```text
Workspace
Project
Site
Page
Block
ContentType
Field
Theme
Connector
PromptRun
PreviewSession
GeneratedFile
DeploymentTarget
RepositoryConnection
```

The project model should be stored as structured data, not only as generated text.
Generated source is an artifact of the model, not the only source of truth during
prototype design.

### 3. Generation Pipeline

Turns user intent and schemas into artifacts:

```text
Prompt -> ProductPlan -> AppSchema -> CMSModel -> TemplBlueprint -> GeneratedFiles
```

Pipeline stages:

1. parse intent;
2. propose app/site structure;
3. map blocks to AppCMS content types and fields;
4. generate Platform module descriptors;
5. generate Templ templates;
6. generate seed content and fixtures;
7. generate tests and export manifests.

### 4. Preview Runtime

Runs prototype behavior without promising a full online IDE:

- BFF PageRuntime for screens;
- preview storage for draft content;
- AppCMS-compatible content model;
- Templ preview rendering;
- BFF actions for CRUD/validation;
- connector mocks or safe connector adapters;
- preview iframe/proxy.

The preview runtime should prove product behavior before a real build/deploy.

### 5. AppCMS Integration

AppCMS provides the content runtime for generated apps:

- content types;
- entries and blocks;
- public pages;
- admin editing;
- REST/codex API;
- media references;
- snapshots/import/export;
- future sync API.

BuildY should orchestrate AppCMS rather than reimplement it.

### 6. Source Export And Repository Sync

GitHub/local IDE own the long-term source workflow.

BuildY may:

- generate source files;
- show a source tree;
- show diffs;
- export zip/local folder;
- create GitHub repo/branch/commit/PR;
- read repository changes back for diff/import.

BuildY should not become the primary code editor.

### 7. Account, Billing, And Portal

Later product layer:

- organizations/workspaces;
- users and roles;
- plans and entitlements;
- usage limits;
- payments and invoices;
- connected repositories;
- deployment targets.

This is a portal surface over the same BFF model, not a Platform feature.

## Ownership Boundaries

| Area | Owner |
|------|-------|
| BFF contract, ScreenModel, actions, fixtures | `fastygo/platform` |
| HTTP host, sessions, middleware | `fastygo/framework` |
| Templ components and generated templates | `fastygo/templ` + BuildY generated code |
| CMS content model/runtime | AppCMS, including `pkg/module/cms.go` |
| Builder UX, projects, prompts, preview sessions | `fastygo/buildy` |
| Source repository | GitHub/local developer workflow |
| Billing and entitlements | BuildY product app |

## Runtime Surfaces

| Surface | Purpose |
|---------|---------|
| Builder admin | BuildY internal product workspace |
| Project preview | Live behavior preview through BFF/Templ |
| Generated public site | AppCMS-backed public pages |
| Generated admin | AppCMS admin for deployed app |
| Account portal | Billing, teams, repository connections |
| Automation API | Future agent/CLI access to project model and runs |

## Key Invariant

BuildY can create product-specific schemas, files, actions, and previews, but it
must not require Platform kernel changes for each new generated product.
