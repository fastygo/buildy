# BuildY Decision Register

## Accepted Decisions

### D001 — BuildY Is A Product App, Not Platform Core

Decision: BuildY-specific project, prompt, preview, sync, billing, and repository
logic belongs in `github.com/fastygo/buildy`.

Reason: Platform must remain a stable BFF model constructor. BuildY is a product
that consumes Platform, AppCMS, Framework, and Templ.

Status: accepted.

### D002 — Preview Workspace, Not Browser IDE

Decision: BuildY's browser workspace is a preview/prototype workspace. It can show
generated source and diffs, but GitHub/local IDE remain the real code workflow.

Reason: Go is compiled, and arbitrary browser-side code execution would create
security and operational complexity too early.

Status: accepted.

### D003 — Templ First

Decision: Generated UI should be Templ-first.

Reason: Templ is the primary renderer in the FastyGo stack and fits the user's
production workflow across many sites/apps.

Status: accepted.

### D004 — AppCMS-Backed Landing Pages

Decision: Generated landings should map blocks and fields to AppCMS content types,
entries, settings, and collections.

Reason: BuildY should generate operational CMS-backed apps, not static pages with
hardcoded data.

Status: accepted.

### D005 — Separate Code, Content, And Schema Sync

Decision: BuildY should track three sync streams independently:

- code sync;
- content sync;
- schema sync.

Reason: GitHub source, live CMS content, and schema/migration state have different
owners, risks, and conflict strategies.

Status: accepted.

### D006 — Start Without Billing And GitHub App Automation

Decision: MVP should prove CMS-backed preview and source export before billing,
GitHub App integration, and deployment workers.

Reason: product/UX proof comes first; operational SaaS layers can follow after
the core workflow is proven.

Status: accepted.

### D007 — No Arbitrary User Code Execution In MVP

Decision: MVP preview should run schemas, BFF models, AppCMS-compatible preview
data, and Templ rendering, not arbitrary uploaded/generated user code.

Reason: isolated code execution requires a separate build worker and security model.

Status: accepted.

## Open Questions

### Q001 — Initial Storage

Question: Should preview storage start as in-memory, SQLite-per-project, or a
shared SQLite database?

Default: SQLite-per-project or scoped SQLite tables once the first preview model
exists. In-memory is acceptable for a very short proof only.

### Q002 — AppCMS Integration Shape

Question: Should the first AppCMS-backed project generate a standalone AppCMS app
or a module/profile that can be mounted by a later suite host?

Default: start with standalone generated app shape because it is easier to export
and deploy. Keep module/profile boundaries clean without adding suite dependencies
to the BuildY prototype.

### Q003 — Content Snapshot Format

Question: Should content snapshots use AppCMS native export format, BuildY manifest
format, or a compatibility wrapper?

Default: BuildY manifest wrapping AppCMS-compatible content snapshots, so sync can
track project/schema metadata without changing AppCMS internals.

### Q004 — Generated Source Ownership

Question: When GitHub integration exists, does BuildY keep the project model as
source of truth or does GitHub become source of truth immediately after export?

Default: BuildY owns structured model during prototype. GitHub owns source after
export. Re-import/diff is explicit, not automatic overwrite.

### Q005 — First Template

Question: Which landing template should prove the MVP?

Default: SaaS landing page with hero, features, pricing, testimonials, FAQ, CTA,
and SEO settings.

### Q006 — Connector MVP

Question: Which connectors are needed in the first proof?

Default: no external connector for first proof; add one safe connector simulation
after CMS-backed CRUD preview works.

## Risk Register

| Risk | Impact | Mitigation |
|------|--------|------------|
| BuildY becomes a generic online IDE | Scope explosion | Keep preview workspace and source ownership docs visible |
| AppCMS logic duplicated in BuildY | Maintenance drift | Use AppCMS models/snapshots/APIs; do not copy domain logic |
| Platform changes for every BuildY feature | Platform instability | Product logic in BuildY; Platform changes only for repeated generic primitives |
| Generated landing is static/hardcoded | Weak product value | Enforce block-to-CMS field mapping |
| Sync overwrites live content | Data loss | Checksums, diffs, explicit conflict resolution |
| Arbitrary code execution too early | Security/ops risk | Defer build workers until after R4/R5 |
| Billing distracts from core UX | Slow MVP | Defer billing to R7 |

## Decision Update Rule

When implementation changes an assumption:

1. update this register;
2. update `planning/roadmap.md` if phase order changes;
3. update `progress.md` with the current state;
4. do not silently move product-specific logic into Platform.
