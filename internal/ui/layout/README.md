# `registry:layout`

Application-owned structural shell — **stays in the app** after blocks/widgets are extracted.

## BuildY packages

| Area | Files | Role |
|------|-------|------|
| Shell | `shell.templ` | Document frame, mobile sheet, main slot |
| Header | `header.templ`, `nav.templ` | Top bar, desktop nav, mobile trigger |
| Footer | `footer.templ` | App footer |
| Glue | `props.go`, `helpers.go` | LayoutData, nav helpers |

## Rules

- Compose with `github.com/fastygo/templ/ui` for primitives.
- Preserve `data-ui8kit-*` hooks for `theme.js` / `ui8kit.js` / `@ui8kit/aria` (see `buildy-aria.mdc`).
- Routing and locale resolution live in `internal/site/` — layout receives resolved props only.

For sidebar-style shell, see git branch **`sidebar`**.
