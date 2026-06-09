# BuildY UX/UI Mockup Progress Tracker

Status: active LLM progress tracker  
Created: 2026-06-08  
Last updated: 2026-06-09  
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
- [x] Hero chat command center.
- [x] Tool cards below the chat.
- [x] Showcase grid for AppCMS, AppCRM, AppSuite, Director, Contentor, Webmaster.
- [x] Prototype-mode notice.
- [x] Dark theme.
- [x] Mobile menu.
- [x] Language switch.

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

- [x] Do not add `@Blank` as a BuildY dependency.
- [x] Do not copy product scope from `@Blank`; use it only as a UI implementation reference.
- [x] Do not add showcase apps as BuildY dependencies.
- [x] Do not add real GitHub, billing, deploy, live AppCMS sync, or sandbox execution.
- [x] Keep BuildY mock data in BuildY-owned fixtures.
- [x] Put reusable Platform changes behind a proven generic need.

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

- [x] Dark theme is mandatory.
- [x] Mobile menu is mandatory.
- [x] Language switching is mandatory.
- [x] The home screen must render from fixtures.
- [x] User-facing copy must be localizable.
- [x] The UI must communicate "prototype workspace", not "browser IDE".
- [x] Generated source areas, if shown, must be readonly.
- [x] Showcase cards must remain presets, not imports.

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

- [x] Define `HomeFixture`.
- [x] Define workspace identity fixture data.
- [x] Define user profile fixture data.
- [ ] Define navigation fixture data.
- [x] Define hero chat fixture data.
- [x] Define prompt suggestion fixture data.
- [x] Define tool card fixture data.
- [x] Define showcase app fixture data.
- [x] Define prototype-mode notice fixture data.
- [x] Define locale fixture data.
- [x] Define theme metadata fixture data.

Showcase fixture presets:

- [x] AppCMS.
- [x] AppCRM.
- [x] AppSuite.
- [x] Director.
- [x] Contentor.
- [x] Webmaster.

Acceptance:

- [x] Home screen can render entirely from fixtures.
- [x] Showcase fixture presets do not imply imports or repo dependencies.

### Chunk 05 — Hero Chat Command Center

Source: `.project/checklist-plan.md`, Plan Chunk 05.

- [x] Render or specify hero heading.
- [x] Render or specify hero subtitle.
- [x] Render or specify large prompt input box.
- [x] Render or specify input action icons.
- [x] Render or specify primary generate CTA.
- [x] Render or specify prompt suggestion chips.
- [x] Render or specify right-side workflow preview.
- [ ] Add disabled state fixture.
- [ ] Add loading state fixture.
- [ ] Add validation error state fixture.

Acceptance:

- [x] Chat is the primary visual focus.
- [x] Chat feels like a build command center, not support chat.
- [x] Hero chat works in dark theme.
- [x] Hero chat works in mobile layout.

### Chunk 06 — Tool Cards Row

Source: `.project/checklist-plan.md`, Plan Chunk 06.

- [x] Render or specify Request Plan card.
- [x] Render or specify CMS Model card.
- [x] Render or specify Preview card.
- [x] Render or specify Source Code card.
- [x] Render or specify Change Review card.
- [x] Render or specify Export card.
- [x] Add icon, title, and description for each card.
- [x] Add responsive card layout.
- [x] Add dark theme card treatment.

Acceptance:

- [x] Cards explain prompt -> plan -> model -> preview -> source -> review -> export.
- [x] Cards remain static/mock in this phase.
- [x] Cards render from fixtures.

### Chunk 07 — Showcase Apps Grid

Source: `.project/checklist-plan.md`, Plan Chunk 07.

- [x] Render or specify AppCMS showcase card.
- [x] Render or specify AppCRM showcase card.
- [x] Render or specify AppSuite showcase card.
- [x] Render or specify Director showcase card.
- [x] Render or specify Contentor showcase card.
- [x] Render or specify Webmaster showcase card.
- [x] Add thumbnail mock area for each card.
- [x] Add category badge for each card.
- [x] Add three capability bullets for each card.
- [x] Add primary CTA for each card.
- [x] Add secondary preview CTA for each card.
- [x] Add mobile showcase layout.
- [x] Add dark theme showcase treatment.

Acceptance:

- [x] Grid matches the target mockup density.
- [x] Cards are fixture presets only.
- [x] Mobile layout remains usable.

### Chunk 08 — Visual QA States

Source: `.project/checklist-plan.md`, Plan Chunk 08.

- [x] Happy path state.
- [ ] Loading showcase state.
- [ ] Empty showcase state.
- [ ] Prompt validation error state.
- [ ] Generation disabled state.
- [x] Prototype mode notice state.
- [x] Mobile menu open state.
- [x] Dark theme active state.
- [x] Russian locale active state.
- [x] English locale active state.

Acceptance:

- [ ] Every state is represented by fixtures.
- [x] No backend behavior is required.
- [x] Screens remain Templ-first and mock-driven.

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

- [x] Home screen renders from fixtures.
- [x] Desktop layout matches the target mockup.
- [ ] Dark theme is implemented and visually checked.
- [ ] Mobile menu is implemented and visually checked.
- [x] Language switch is implemented for English and Russian.
- [x] Hero chat, tool cards, showcase grid, and prototype notice are present.
- [x] Showcase presets do not introduce extra module dependencies.
- [x] No live GitHub, billing, deployment, AppCMS sync, or sandbox behavior exists.
- [x] `go test ./...` passes once BuildY code exists.

## Progress Log

### 2026-06-08

- [x] Reviewed current workspace composition.
- [x] Focused BuildY product scope on BuildY, Templ, Framework, Platform, and AppCMS.
- [x] Kept Blank as a reference-only workspace folder for UI shell, theme, mobile menu, and locale patterns.
- [x] Confirmed the CMS module is inside `@AppCMS/pkg/module/cms.go`.
- [x] Created `.project/checklist-plan.md` for copyable Plan-mode chunks.
- [x] Converted this file into an LLM-friendly checkbox progress tracker.
- [ ] Start Plan Chunk 01 — Design System Foundation.

### 2026-06-09

- [x] Added persistent spec for the bilingual YC-quality runtime home redesign.
- [x] Expanded EN/RU home fixtures for workspace identity, hero chat, suggestions, workflow steps, tool cards, showcase presets, prototype notice, and ARIA labels.
- [x] Rebuilt runtime `HomePage` inside the existing shell without importing the isolated catalog block.
- [x] Added progressive chat textarea enhancement for auto-resize and internal scrolling on long prompts.
- [x] Verified with `bun run verify` (templ generate, CSS/JS build, ui8px lint, ARIA validation, Go tests).
