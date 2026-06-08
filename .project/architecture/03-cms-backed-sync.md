# CMS-Backed Sync

## Thesis

BuildY should generate CMS-backed products, not static mockups.

Every generated landing or site should map visible blocks and editable text to
AppCMS content types, fields, settings, and entries. After export and deploy, the
same content can be edited in the live AppCMS admin or in BuildY and synchronized.

## Example

A SaaS landing page should not generate hardcoded copy:

```text
Hero title -> CMS field
Hero subtitle -> CMS field
Features -> CMS collection
Pricing plans -> CMS collection
Testimonials -> CMS collection
FAQ items -> CMS collection
SEO title/description -> CMS settings
```

This makes the generated product operational immediately.

## Content Mapping Model

BuildY should represent page blocks as structured mappings:

```yaml
page: home
blocks:
  hero:
    source: cms.singleton.hero
    fields:
      title: text
      subtitle: textarea
      cta_label: text
      cta_href: url

  features:
    source: cms.collection.features
    fields:
      title: text
      description: textarea
      icon: text

  pricing:
    source: cms.collection.pricing_plans
    fields:
      name: text
      price: text
      features: repeater
      cta_label: text
      cta_href: url
```

The mapping should generate:

- AppCMS content type definitions;
- BFF preview screens/actions;
- Templ rendering inputs;
- seed content;
- export manifests;
- future migrations.

## Three Sync Streams

Keep code, content, and schema separate.

### 1. Code Sync

Purpose: source ownership and developer continuation.

```text
BuildY generated files -> GitHub repo
GitHub repo -> BuildY diff/import
```

Source of truth: GitHub/local IDE after export.

Typical artifacts:

- Go modules;
- Templ files;
- app profiles;
- content model manifests;
- seed files;
- tests/fixtures;
- Dockerfile/deploy manifests later.

### 2. Content Sync

Purpose: keep BuildY preview content and live AppCMS content aligned.

```text
BuildY preview store <-> Live AppCMS content store
```

Initial mode:

- manual push;
- manual pull;
- diff before apply;
- conflict warning.

Later mode:

- scheduled sync;
- webhook sync;
- per-entry conflict resolution.

### 3. Schema Sync

Purpose: evolve content types and generated app structure.

```text
BuildY schema -> AppCMS content type migrations
Live schema changes -> BuildY import/diff
```

During early prototype work, BuildY owns schema. After deployment, schema changes
should flow through migrations or explicit schema publish steps.

## Sync Artifacts

BuildY should define durable artifacts:

```text
ProjectManifest
ContentSchemaManifest
ContentSnapshot
GeneratedFileManifest
SyncState
ConflictReport
DeploymentBinding
```

### ProjectManifest

Describes app identity:

- project id;
- name;
- app type;
- modules;
- renderer;
- route bases;
- generated package/module names;
- connected repository.

### ContentSchemaManifest

Describes CMS types:

- content type id;
- fields;
- relations;
- validation;
- singleton/collection mode;
- public/admin visibility.

### ContentSnapshot

Describes content state:

- entries;
- settings;
- media references;
- locales;
- status/draft state;
- version/checksum.

### SyncState

Tracks synchronization:

- source (`buildy`, `github`, `live-appcms`);
- target;
- last synced revision;
- content checksum;
- schema checksum;
- source commit sha;
- last conflict report.

## AppCMS Integration Points

BuildY should use AppCMS through product-level contracts:

| Need | AppCMS role |
|------|-------------|
| Content types | AppCMS/module descriptors and schema manifests |
| CRUD preview | BFF actions over preview storage |
| Live admin | AppCMS admin after deploy |
| Public page data | AppCMS content services/public rendering |
| Import/export | snapshots or content sync API |
| SEO/media/settings | AppCMS domain/runtime |

Do not copy AppCMS domain logic into BuildY. BuildY orchestrates schema generation,
preview storage, source generation, and sync.

## Conflict Strategy

Conflicts should be explicit.

Examples:

- BuildY changed `hero.title`, live AppCMS also changed `hero.title`;
- GitHub changed a generated Templ file while BuildY regenerated it;
- live AppCMS added a content field unknown to BuildY schema;
- schema migration would remove fields that contain live content.

MVP conflict handling:

1. detect checksum mismatch;
2. show source/target diff;
3. allow choose BuildY wins / live wins / skip;
4. do not silently overwrite production content.

## MVP Sync Flow

1. User creates CMS-backed landing in BuildY.
2. BuildY creates preview content types and seed content.
3. User edits content in BuildY preview workspace.
4. BuildY exports generated source and content snapshot.
5. User deploys generated AppCMS app.
6. User manually pushes content snapshot to live app.
7. User manually pulls live content back into BuildY.

GitHub automation and webhooks come later.

## Future Sync Flow

```text
BuildY project <-> GitHub repository
BuildY preview content <-> Live AppCMS content
BuildY schema <-> AppCMS migrations
```

With:

- GitHub App integration;
- deployment callbacks;
- AppCMS webhooks;
- branch-based generated diffs;
- content sync jobs;
- conflict reports.

## Guardrails

- Code sync must not overwrite developer changes without a diff.
- Content sync must not overwrite production content without a conflict policy.
- Schema sync must not delete live fields/content silently.
- BuildY must not become the live CMS database.
- AppCMS remains the generated app content runtime.
