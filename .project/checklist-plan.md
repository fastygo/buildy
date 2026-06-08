Да, лучше превратить `progress-ux-mockup-fixtures.md` из общего roadmap в **пошаговый Plan-mode сценарий**: каждый блок должен быть самодостаточным, чтобы ты мог копировать его в чат и получать один маленький mockup-слой.

`@Blank` стоит зафиксировать как **reference-only пример**, не как активную зависимость BuildY. В нём полезны паттерны: `internal/ui/layout/shell.templ` для shell + mobile sheet, `web/static/js/theme.js` для dark/light toggle, `internal/ui/components/toggles/language_switch.templ` и `internal/fixtures/locale/*.json` для языков. `@Blank/go.mod` показывает базовый стек: `framework` + `templ`, но BuildY всё равно должен держать свой фокус на текущем продукте.

Я бы спроектировал файл так.

```md
## Design System Mockup Plan

Goal: reproduce the provided BuildY home mockup as fixture-driven Templ screens.

Reference image:
- BuildY home screen with top bar, left rail, hero chat, tool cards, showcase grid, prototype banner.

Reference-only implementation examples:
- `@Blank/internal/ui/layout/shell.templ`
- `@Blank/internal/ui/layout/header.templ`
- `@Blank/web/static/js/theme.js`
- `@Blank/internal/ui/components/toggles/language_switch.templ`
- `@Blank/internal/fixtures/locale/en.json`
- `@Blank/internal/fixtures/locale/ru.json`

Do not add `@Blank` as a BuildY dependency.
```

Дальше не одним большим ТЗ, а серией копируемых задач.

### 1. Foundation Prompt

```md
### Plan Chunk 01 — Design System Foundation

Goal:
Create BuildY design-system foundation for the home mockup.

Scope:
- Define visual tokens for light and dark themes.
- Define spacing, radius, border, shadow, typography, status colors.
- Define component inventory for the home screen.
- Do not implement backend behavior.

Required components:
- AppShell
- TopBar
- LeftRail
- MobileMenu
- ThemeToggle
- LanguageSwitch
- HeroChat
- ToolCard
- ShowcaseCard
- PrototypeModeBanner

Acceptance:
- Dark theme is mandatory.
- Mobile menu is mandatory.
- Language switch is mandatory.
- Everything is fixture-driven.
- No real GitHub, billing, deploy, live sync, or sandbox execution.
```

### 2. Shell Prompt

```md
### Plan Chunk 02 — Responsive App Shell

Goal:
Build the responsive BuildY shell matching the mockup.

Desktop layout:
- Fixed top bar.
- Left vertical icon rail.
- Main scroll area.
- Centered content width.
- Header actions: documentation, templates, settings, user menu.

Mobile layout:
- Top bar remains visible.
- Left rail collapses into mobile menu.
- Mobile menu opens as off-canvas sheet.
- Menu must be keyboard accessible and closable.

Reference:
Use `@Blank/internal/ui/layout/shell.templ` as behavior inspiration only.

Acceptance:
- Desktop, tablet, and mobile layouts are represented.
- Mobile menu is not optional.
- Shell supports dark theme classes.
```

### 3. Theme And Language Prompt

```md
### Plan Chunk 03 — Dark Theme And Language Switch

Goal:
Add theme and language UX for the BuildY home mockup.

Theme:
- Default supports light and dark.
- Dark mode must be visually designed, not only technically toggled.
- Theme toggle appears in top bar and mobile menu.

Language:
- Add language switch for `EN` and `RU`.
- All visible home-screen copy must come from locale fixtures.
- Language switch appears in top bar and mobile menu.

Reference:
Use `@Blank/web/static/js/theme.js` and `@Blank/internal/ui/components/toggles/language_switch.templ` as patterns.

Acceptance:
- No hardcoded user-facing copy in templates.
- Dark theme covers background, cards, borders, chat input, showcase cards, badges.
```

### 4. Home Fixture Prompt

```md
### Plan Chunk 04 — Home Screen Fixtures

Goal:
Define mock fixtures for the BuildY home screen.

Fixture shape:
- Workspace identity.
- User profile.
- Navigation items.
- Hero chat copy.
- Prompt suggestions.
- Tool cards.
- Showcase apps.
- Prototype mode notice.
- Locale strings.
- Theme metadata.

Showcase apps:
- AppCMS
- AppCRM
- AppSuite
- Director
- Contentor
- Webmaster

Important:
These showcase apps are mock presets, not repo dependencies.

Acceptance:
- Home screen can render entirely from fixtures.
- Showcase cards do not imply BuildY imports those apps.
```

### 5. Hero Chat Prompt

```md
### Plan Chunk 05 — Hero Chat Command Center

Goal:
Render the main hero chat area from the mockup.

Layout:
- Large heading: “What do you want to create?”
- Short explanatory subtitle.
- Large prompt input box.
- Attachment/tool icons inside the input.
- Primary CTA: Generate prototype.
- Prompt suggestion chips below input.
- Right-side visual workflow preview: plan, model, preview, source, export.

Behavior:
- Static mock only.
- Show disabled/loading/error variants through fixtures.

Acceptance:
- Chat is the primary visual focus.
- It feels like a build command center, not support chat.
- Works in dark theme and mobile layout.
```

### 6. Tool Cards Prompt

```md
### Plan Chunk 06 — Tool Cards Row

Goal:
Render the tool cards below the hero chat.

Cards:
- Request Plan
- CMS Model
- Preview
- Source Code
- Change Review
- Export

Each card includes:
- Icon
- Title
- Short description
- Status/availability metadata if useful

Acceptance:
- Tool cards explain the workflow:
  prompt -> plan -> model -> preview -> source -> review -> export.
- Cards are responsive.
- Cards support dark theme.
```

### 7. Showcase Prompt

```md
### Plan Chunk 07 — Showcase Apps Grid

Goal:
Render the showcase grid from the provided mockup.

Cards:
- AppCMS: CMS sites, pages, media, taxonomies.
- AppCRM: leads, contacts, deals, activities.
- AppSuite: launch hub and workspace switcher.
- Director: operations dashboard and approvals.
- Contentor: editorial planning and publishing.
- Webmaster: SEO, redirects, site health.

Each card:
- Thumbnail mock image area.
- Category badge.
- Description.
- Three capability bullets.
- Primary CTA: Use.
- Secondary CTA: Preview.

Acceptance:
- Grid matches the visual density of the mockup.
- Mobile layout becomes a single-column or horizontal card stack.
- Apps are fixture presets only, not dependencies.
```

### 8. QA Prompt

```md
### Plan Chunk 08 — Visual QA States

Goal:
Add required visual states for the BuildY home mockup.

States:
- Happy path.
- Loading showcase.
- Empty showcase.
- Prompt validation error.
- Generation disabled.
- Prototype mode notice.
- Mobile menu open.
- Dark theme active.
- Russian locale active.
- English locale active.

Acceptance:
- Every state is represented by fixtures.
- No backend behavior is required.
- Screens remain Templ-first and mock-driven.
```

Главная мысль: в `progress-ux-mockup-fixtures.md` нужно добавить не только roadmap, а раздел **“Plan Chunks”**. Каждый chunk должен иметь одинаковую форму: `Goal`, `Scope`, `Reference`, `Acceptance`, `Out of scope`. Тогда ты сможешь идти по одному блоку за раз и не терять фокус.