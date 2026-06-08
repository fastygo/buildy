# BuildY UX/UI Mockup Progress Tracker

Status: active LLM progress tracker  
Created: 2026-06-08  
Last updated: 2026-06-08  
Scope: fixture-driven BuildY home mockup and design system.

## How To Use This File

This file must be included in every Plan-mode step for BuildY UX mockups.

The LLM must use it as the progress source of truth:

- [ ] Read this file before proposing or implementing a BuildY mockup step.
- [ ] Read `.project/checklist-plan.md` for the copy-paste Plan chunks.
- [ ] Work on one Plan Chunk at a time unless the user explicitly asks otherwise.
- [ ] Mark a checkbox as done only after the matching work exists in code, fixtures, docs, or verified mock output.
- [ ] Leave unchecked items untouched when they are planned but not implemented.
- [ ] Add a short Progress Log entry after every completed chunk.
- [ ] Do not mark implementation work complete from design discussion alone.

Updating this file (checkboxes, Progress Log) is allowed when completing a Plan Chunk — see `@BuildY/.cursor/rules/agent-conduct.mdc`.

## Source Plan

Copyable Plan-mode tasks live in:

```text
.project/checklist-plan.md
```

This progress file tracks completion for those chunks. Do not duplicate the full
Plan prompts here. Keep this file short enough for repeated LLM context.

## Current Target Mockup

The current target is the BuildY home screen mockup:

- [ ] Desktop app shell with top bar and left icon rail.
- [ ] Hero chat command center.
- [ ] Tool cards below the chat.
- [ ] Showcase grid for AppCMS, AppCRM, AppSuite, Director, Contentor, Webmaster.
- [ ] Prototype-mode notice.
- [ ] Dark theme.
- [ ] Mobile menu.
- [ ] Language switch.

Showcase names are fixture presets only. They are not repo dependencies.

## Workspace And Dependency Guardrails

Active BuildY workspace:

```text
e:\_@Go\.WorkSpace-BuildY\@buildy.code-workspace
```

Workspace folders:

- [x] `@BuildY` — product app, UX, project model, preview workspace.
- [x] `@Templ` — primary renderer and generated UI component layer.
- [x] `@Framework` — app host, HTTP runtime, sessions, middleware.
- [x] `@Platform` — BFF contracts, ScreenModel, actions, fixtures, panel/toolset.
- [x] `@AppCMS` — CMS-backed runtime reference and integration target.
- [x] `@Blank` — reference-only example for shell, mobile menu, theme toggle, and language switch patterns.

Rules:

- [ ] Do not add `@Blank` as a BuildY dependency.
- [ ] Do not copy product scope from `@Blank`; use it only as a UI implementation reference.
- [ ] Do not add showcase apps as BuildY dependencies.
- [ ] Do not add real GitHub, billing, deploy, live AppCMS sync, or sandbox execution.
- [ ] Keep BuildY mock data in BuildY-owned fixtures.
- [ ] Put reusable Platform changes behind a proven generic need.

## Reference-Only Implementation Patterns

Use these as references when the matching Plan Chunk is active:

- [ ] `@Blank/internal/ui/layout/shell.templ` — shell and mobile sheet pattern.
- [ ] `@Blank/internal/ui/layout/header.templ` — header layout and mobile trigger pattern.
- [ ] `@Blank/web/static/js/theme.js` — dark/light theme toggle pattern.
- [ ] `@Blank/internal/ui/components/toggles/language_switch.templ` — compact language switch pattern.
- [ ] `@Blank/internal/fixtures/locale/en.json` — locale fixture shape.
- [ ] `@Blank/internal/fixtures/locale/ru.json` — locale fixture shape.

Do not copy `@Blank` wholesale. Adapt only the smallest useful pattern.

## Mandatory Product Requirements

These requirements apply to every home mockup chunk:

- [ ] Dark theme is mandatory.
- [ ] Mobile menu is mandatory.
- [ ] Language switching is mandatory.
- [ ] The home screen must render from fixtures.
- [ ] User-facing copy must be localizable.
- [ ] The UI must communicate "prototype workspace", not "browser IDE".
- [ ] Generated source areas, if shown, must be readonly.
- [ ] Showcase cards must remain presets, not imports.

## Plan Chunk Progress

### Chunk 01 — Design System Foundation

Source: `.project/checklist-plan.md`, Plan Chunk 01.

- [ ] Define BuildY visual tokens for light theme.
- [ ] Define BuildY visual tokens for dark theme.
- [ ] Define spacing, radius, border, shadow, and typography rules.
- [ ] Define status colors for info, success, warning, error, and muted states.
- [ ] Define home-screen component inventory.
- [ ] Confirm no backend behavior is included in this chunk.

Acceptance:

- [ ] Dark theme is explicitly covered.
- [ ] Mobile menu is listed as a first-class component.
- [ ] Language switch is listed as a first-class component.
- [ ] The foundation supports fixture-driven rendering.

### Chunk 02 — Responsive App Shell

Source: `.project/checklist-plan.md`, Plan Chunk 02.

- [ ] Define desktop shell structure.
- [ ] Define top bar structure.
- [ ] Define left icon rail structure.
- [ ] Define main scroll/content region.
- [ ] Define mobile shell structure.
- [ ] Define off-canvas mobile menu behavior.
- [ ] Define keyboard and close behavior expectations for the mobile menu.
- [ ] Confirm shell supports dark theme classes.

Acceptance:

- [ ] Desktop layout matches the target mockup structure.
- [ ] Mobile menu is implemented or fully specified.
- [ ] Mobile menu is not treated as optional.
- [ ] Shell can host later BuildY screens.

### Chunk 03 — Dark Theme And Language Switch

Source: `.project/checklist-plan.md`, Plan Chunk 03.

- [ ] Add or specify theme toggle.
- [ ] Add or specify dark theme persistence.
- [ ] Cover dark styling for background, cards, borders, chat input, badges, and showcase cards.
- [ ] Add or specify `EN` / `RU` language switch.
- [ ] Add or specify locale fixtures.
- [ ] Ensure visible home copy is localizable.
- [ ] Place theme and language controls in top bar.
- [ ] Place theme and language controls in mobile menu.

Acceptance:

- [ ] Dark theme is visually designed, not just technically toggled.
- [ ] Language switch works from fixtures or is fully specified for fixtures.
- [ ] No hardcoded user-facing copy remains in the home template once implemented.

### Chunk 04 — Home Screen Fixtures

Source: `.project/checklist-plan.md`, Plan Chunk 04.

- [ ] Define `HomeFixture`.
- [ ] Define workspace identity fixture data.
- [ ] Define user profile fixture data.
- [ ] Define navigation fixture data.
- [ ] Define hero chat fixture data.
- [ ] Define prompt suggestion fixture data.
- [ ] Define tool card fixture data.
- [ ] Define showcase app fixture data.
- [ ] Define prototype-mode notice fixture data.
- [ ] Define locale fixture data.
- [ ] Define theme metadata fixture data.

Showcase fixture presets:

- [ ] AppCMS.
- [ ] AppCRM.
- [ ] AppSuite.
- [ ] Director.
- [ ] Contentor.
- [ ] Webmaster.

Acceptance:

- [ ] Home screen can render entirely from fixtures.
- [ ] Showcase fixture presets do not imply imports or repo dependencies.

### Chunk 05 — Hero Chat Command Center

Source: `.project/checklist-plan.md`, Plan Chunk 05.

- [ ] Render or specify hero heading.
- [ ] Render or specify hero subtitle.
- [ ] Render or specify large prompt input box.
- [ ] Render or specify input action icons.
- [ ] Render or specify primary generate CTA.
- [ ] Render or specify prompt suggestion chips.
- [ ] Render or specify right-side workflow preview.
- [ ] Add disabled state fixture.
- [ ] Add loading state fixture.
- [ ] Add validation error state fixture.

Acceptance:

- [ ] Chat is the primary visual focus.
- [ ] Chat feels like a build command center, not support chat.
- [ ] Hero chat works in dark theme.
- [ ] Hero chat works in mobile layout.

### Chunk 06 — Tool Cards Row

Source: `.project/checklist-plan.md`, Plan Chunk 06.

- [ ] Render or specify Request Plan card.
- [ ] Render or specify CMS Model card.
- [ ] Render or specify Preview card.
- [ ] Render or specify Source Code card.
- [ ] Render or specify Change Review card.
- [ ] Render or specify Export card.
- [ ] Add icon, title, and description for each card.
- [ ] Add responsive card layout.
- [ ] Add dark theme card treatment.

Acceptance:

- [ ] Cards explain prompt -> plan -> model -> preview -> source -> review -> export.
- [ ] Cards remain static/mock in this phase.
- [ ] Cards render from fixtures.

### Chunk 07 — Showcase Apps Grid

Source: `.project/checklist-plan.md`, Plan Chunk 07.

- [ ] Render or specify AppCMS showcase card.
- [ ] Render or specify AppCRM showcase card.
- [ ] Render or specify AppSuite showcase card.
- [ ] Render or specify Director showcase card.
- [ ] Render or specify Contentor showcase card.
- [ ] Render or specify Webmaster showcase card.
- [ ] Add thumbnail mock area for each card.
- [ ] Add category badge for each card.
- [ ] Add three capability bullets for each card.
- [ ] Add primary CTA for each card.
- [ ] Add secondary preview CTA for each card.
- [ ] Add mobile showcase layout.
- [ ] Add dark theme showcase treatment.

Acceptance:

- [ ] Grid matches the target mockup density.
- [ ] Cards are fixture presets only.
- [ ] Mobile layout remains usable.

### Chunk 08 — Visual QA States

Source: `.project/checklist-plan.md`, Plan Chunk 08.

- [ ] Happy path state.
- [ ] Loading showcase state.
- [ ] Empty showcase state.
- [ ] Prompt validation error state.
- [ ] Generation disabled state.
- [ ] Prototype mode notice state.
- [ ] Mobile menu open state.
- [ ] Dark theme active state.
- [ ] Russian locale active state.
- [ ] English locale active state.

Acceptance:

- [ ] Every state is represented by fixtures.
- [ ] No backend behavior is required.
- [ ] Screens remain Templ-first and mock-driven.

## Later UX Prototype Roadmap

Do not start these until the home mockup foundation is usable:

- [ ] Projects dashboard.
- [ ] New project wizard.
- [ ] Prototype workspace layout.
- [ ] CMS-backed landing builder.
- [ ] Content model designer.
- [ ] Preview panel.
- [ ] Generated source viewer.
- [ ] Diff and change review.
- [ ] Preview data CRUD.
- [ ] Sync center mock.
- [ ] Export and repository placeholders.
- [ ] Account and billing placeholders.
- [ ] Error, empty, and permission states.
- [ ] Design system consolidation.

## Suggested BuildY-Owned Mockup Files

Implementation may use this shape when code starts:

```text
internal/fixtures/
  home.go
  locale.go
  showcase.go
  tools.go
  workspace.go

internal/ui/layout/
  shell.templ
  header.templ
  rail.templ
  mobile_menu.templ

internal/ui/home/
  hero_chat.templ
  tool_cards.templ
  showcase_grid.templ
  prototype_notice.templ

web/static/js/
  theme.js
  ui.js
```

This is a suggested structure, not a requirement. Follow existing repo patterns
once implementation begins.

## Definition Of Done For Home Mockup

- [ ] Home screen renders from fixtures.
- [ ] Desktop layout matches the target mockup.
- [ ] Dark theme is implemented and visually checked.
- [ ] Mobile menu is implemented and visually checked.
- [ ] Language switch is implemented for English and Russian.
- [ ] Hero chat, tool cards, showcase grid, and prototype notice are present.
- [ ] Showcase presets do not introduce extra module dependencies.
- [ ] No live GitHub, billing, deployment, AppCMS sync, or sandbox behavior exists.
- [ ] `go test ./...` passes once BuildY code exists.

## Progress Log

### 2026-06-08

- [x] Reviewed current workspace composition.
- [x] Focused BuildY product scope on BuildY, Templ, Framework, Platform, and AppCMS.
- [x] Kept Blank as a reference-only workspace folder for UI shell, theme, mobile menu, and locale patterns.
- [x] Confirmed the CMS module is inside `@AppCMS/pkg/module/cms.go`.
- [x] Created `.project/checklist-plan.md` for copyable Plan-mode chunks.
- [x] Converted this file into an LLM-friendly checkbox progress tracker.
- [ ] Start Plan Chunk 01 — Design System Foundation.
