# `registry:layout`

Application-owned structural shell — **stays in the app** after blocks/widgets are extracted.

## BuildY packages

| Area | Files | Role |
|------|-------|------|
| Shell | `shell.templ` | Document frame, sidebar, mobile sheet groups, main slot |
| Header | `header.templ`, `nav.templ` | Top bar, header nav, mobile trigger |
| Sidebar | `sidebar.templ` | Primary icon rail / first mobile sheet group |
| Footer | `footer.templ` | App footer |
| Glue | `props.go`, `helpers.go` | LayoutData, nav helpers |

## Rules

- Compose with `github.com/fastygo/templ/ui` for primitives.
- Preserve `data-ui8kit-*` hooks for `theme.js` / `ui8kit.js` / `@ui8kit/aria` (see `buildy-aria.mdc`).
- Routing and locale resolution live in `internal/site/` — layout receives resolved props only.
- Split navigation into `SidebarItems` (primary rail) and `HeaderNavItems` (top bar). Mobile sheet renders both groups sequentially.
