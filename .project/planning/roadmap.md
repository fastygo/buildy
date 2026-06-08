# BuildY Roadmap

## Current Stage

BuildY is at concept and product architecture stage.

Current module:

```text
github.com/fastygo/buildy
```

The first implementation target is not a full SaaS. It is a CMS-backed prototype
workspace that can generate and preview a real AppCMS-powered landing.

## Roadmap Summary

```text
R0 Product memory
R1 Prototype workspace shell
R2 CMS-backed landing model
R3 Behavior preview
R4 Generated source preview/export
R5 AppCMS live sync proof
R6 GitHub integration
R7 Account, billing, and entitlements
R8 Build/deploy workers
```

## R0 — Product Memory And Scope

Goal: keep the product direction stable before coding.

Work:

- maintain `.project` docs;
- define core terms;
- document non-goals;
- align with Platform BFF and AppCMS boundaries;
- keep progress and decisions updated.

Exit criteria:

- `.project/INDEX.md` exists;
- north star, system map, preview model, sync model, and roadmap exist;
- future agents can recover BuildY's direction without `.manual` notes.

Status: in progress.

## R1 — Prototype Workspace Shell

Goal: create the first BuildY product app shell.

Work:

- set up `cmd/server`;
- use `fastygo/framework` host;
- add Templ shell;
- create routes for project list and project detail;
- add stub auth/session or local developer mode;
- build initial BFF screens for projects, runs, and preview.

Suggested routes:

```text
/builder
/builder/projects
/builder/projects/{projectID}
/builder/projects/{projectID}/preview
/builder/projects/{projectID}/files
/builder/projects/{projectID}/runs
```

Exit criteria:

- user can open BuildY shell;
- user can create/select a project;
- project dashboard has placeholders for prompt, preview, file tree, and runs;
- no AppCMS sync or billing required.

Do not:

- add payment;
- add GitHub OAuth;
- execute arbitrary code;
- edit Platform kernel.

## R2 — CMS-Backed Landing Model

Goal: generate a structured AppCMS-backed landing model.

Work:

- define BuildY project schema;
- define page/block schema;
- define content mapping model;
- add first landing template;
- map blocks to AppCMS content types/fields;
- generate seed content.

First landing blocks:

- hero;
- features;
- pricing;
- testimonials;
- FAQ;
- CTA;
- SEO settings.

Exit criteria:

- a project can produce a structured landing model;
- landing blocks map to CMS fields/collections;
- no hardcoded visible content is required in generated Templ;
- model can be serialized and reloaded.

## R3 — Behavior Preview

Goal: make the prototype behave like a real CMS-backed app.

Work:

- preview storage;
- AppCMS-compatible content entries;
- BFF PageRuntime for preview pages;
- BFF actions for editing preview content;
- Templ preview iframe;
- validation errors and success states;
- reset/seed preview data.

Exit criteria:

- editing hero/pricing/FAQ data changes preview;
- CRUD-like changes use BFF actions;
- preview proves behavior without compiling arbitrary user code;
- preview content can be exported as a snapshot.

## R4 — Generated Source Preview And Export

Goal: show and export real generated source.

Work:

- generated file manifest;
- source tree viewer;
- readonly file viewer;
- diff viewer;
- export zip or local folder;
- generated `go.mod`, `cmd/server`, module descriptors, Templ files, seeds;
- generation run history.

Exit criteria:

- user can review generated files;
- user can export a source snapshot;
- generated source represents the same project model as preview;
- source tree is not positioned as the primary code editor.

## R5 — AppCMS Live Sync Proof

Goal: prove BuildY preview content can sync with a deployed AppCMS app.

Work:

- content snapshot format;
- content import/export;
- schema version/checksum;
- manual push to live AppCMS;
- manual pull from live AppCMS;
- conflict report.

Exit criteria:

- BuildY can push content to a live/generated AppCMS app;
- BuildY can pull changed content back;
- conflicts are detected and shown;
- code sync and content sync remain separate.

## R6 — GitHub Integration

Goal: connect generated source to real repository ownership.

Work:

- GitHub repository connection;
- export to repo;
- branch per generation run;
- commit generated changes;
- pull repository diff;
- optional PR creation.

Exit criteria:

- BuildY can create or update a repo branch;
- generated changes are reviewable in GitHub;
- user can continue in local IDE;
- BuildY does not overwrite developer changes silently.

## R7 — Account, Billing, And Entitlements

Goal: make BuildY a SaaS product.

Work:

- user account;
- organization/workspace;
- plans;
- subscription provider integration;
- usage metering;
- entitlements;
- invoices and customer portal;
- limits for projects, runs, preview minutes, and repo connections.

Exit criteria:

- backend enforces entitlements;
- account portal exists;
- billing state affects actions securely;
- free/paid plan boundaries are clear.

Do not enforce billing only in the UI.

## R8 — Build And Deploy Workers

Goal: validate and deploy generated apps.

Work:

- isolated build worker;
- `templ generate`;
- `go test`;
- `go build`;
- deployment adapter;
- logs;
- resource limits;
- secret boundary.

Exit criteria:

- generated app can be built in isolation;
- build logs are visible in BuildY;
- deploy target can receive artifact;
- arbitrary code execution is isolated from BuildY core.

## First MVP

MVP should include:

- R1 shell;
- R2 CMS-backed landing model;
- R3 behavior preview;
- partial R4 source preview/export.

MVP should not include:

- billing;
- GitHub App automation;
- arbitrary code execution;
- production deployment workers;
- complex realtime.

## Progress Tracking

Update `../progress.md` after each implementation pass:

- completed scope;
- tests/verification;
- changed assumptions;
- next best action;
- blockers.
