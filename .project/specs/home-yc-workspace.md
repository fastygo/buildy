# BuildY Home YC Workspace Spec

Status: active
Scope: runtime BuildY home page, English and Russian fixtures.

## Intent

Turn the runtime home page into a polished product prototype workspace suitable
for a YC-style demo. The user describes what they want to build, and BuildY
frames the flow as prompt -> plan -> model -> preview -> source -> export.

## References

- Provided English and Russian home mockup images.
- `internal/ui/blocks/workspace/home` as visual guidance only.
- BuildY rules in `.cursor/rules`.
- Templ component specs when choosing primitives and composites.

## Affected Surfaces

- `internal/fixtures/screens/home.en.json`
- `internal/fixtures/screens/home.ru.json`
- `internal/fixtures/fixtures.go`
- `internal/site/provider.go`
- `internal/views/models.go`
- `internal/views/home.templ`
- `internal/views/wireframe_render_test.go`
- `.ui8px/policy/*` only if validation proves a missing utility is needed.

## Data And Locale

All runtime copy must come from fixtures. English and Russian home fixtures must
share the same JSON shape.

Fixture groups:

- hero command center;
- chat composer labels and prompts;
- prompt suggestions;
- workflow steps;
- tool cards;
- showcase presets;
- prototype-mode notice;
- static ARIA labels.

## UI Structure

Use the existing app shell. Do not duplicate sidebar, header, mobile sheet,
theme toggle, language switch, or footer in `home.templ`.

Home sections:

1. Hero command center with modern AI chat composer.
2. Workflow preview rail.
3. Tool cards for plan, model, preview, source, diff, export.
4. Showcase grid with AppCMS, AppCRM, AppSuite, Director, Contentor, Webmaster.
5. Prototype-mode notice.

## Accessibility

- Use fixture-backed accessible names for chat, textarea, actions, suggestions,
  workflow, and showcase.
- Preserve the existing `@ui8kit/aria` dialog/sheet shell contract.
- Do not add a new page-level dialog or mobile sheet.

## Validation

Run:

```bash
bun run verify
```

This must cover templ generation, CSS build, JS build, ui8px validation, ARIA
validation, and Go tests.

## Non-Goals

- Do not import the catalog `home.Page` into runtime.
- Do not introduce real GitHub, billing, deploy, live AppCMS sync, or sandbox
  code execution.
- Do not add legacy UI dependencies or raw product copy in templates.

