# Version History & PromptRun UX

Status: workspace product spec  
BuildY domain: `PromptRun`, `GeneratedFile`, `GeneratedChange`, preview sessions.

## Problem

Users iterate with prompts and schema edits. A bad run breaks preview or generated output. Without history they cannot:

- find the last good state;
- compare what changed;
- revert without manual reconstruction.

## Capability summary

| Feature | User outcome |
|---------|--------------|
| Run log | Every generation/edit pass is a labeled `PromptRun` |
| Preview before revert | Open preview at run N before restoring |
| Quick recent runs | Jump to last 8 runs from workspace shell |
| Artifact diffs | See which files changed per run |
| Model rollback | Restore schema/content to run boundary |

## PromptRun model

Aligns with [`../../progress.md`](../../progress.md):

```text
PromptRun
  id
  project_id
  created_at
  label              # auto + user-editable
  prompt_text        # optional summary
  status             # pending | running | success | failed
  preview_session_id
  parent_run_id      # optional fork
```

### Auto labels (examples)

- `Initial SaaS landing from template`
- `Added pricing tier — Pro`
- `FAQ block — 3 new items`
- `Theme — dark variant`
- `Revert to run #4`

User can rename labels (descriptive version labels).

## Workspace UI

### Run timeline (sidebar or panel)

```text
┌ PromptRun history ──────────────┐
│ ● Run 8  FAQ expanded      now   │
│ ○ Run 7  Pricing tweak     2h    │
│ ○ Run 6  Hero copy         1d    │
│ …                                │
│ [Load more]                      │
└──────────────────────────────────┘
```

**Performance:** lazy-load older runs on scroll — do not render full history in DOM at once.

### Run detail view

Route: `/builder/projects/{id}/runs/{runID}`

Shows:

- label, timestamp, status, prompt excerpt;
- **Preview at this run** (primary action);
- changed artifacts list (`GeneratedChange`);
- **Restore model to this run** (secondary, confirm dialog);
- link to full file diff.

### Preview before revert flow

```text
1. User selects Run 5 from timeline
2. Click "Preview this version"
3. Preview iframe loads snapshot bound to Run 5
4. User confirms → "Restore" writes schema/content forward from Run 5
5. New PromptRun created: "Restored from Run 5"
```

Never silent overwrite — always create a new run entry for audit trail.

## GeneratedChange linkage

Per [`../../architecture/02-preview-workspace.md`](../../architecture/02-preview-workspace.md):

```text
GeneratedChange
  run_id
  path
  operation: create | update | delete
  before_hash / after_hash
  source_model_refs   # Page, Block, ContentType IDs
```

UI groups changes:

- **Structure** — schema, blocks, routes  
- **Content** — preview store entries  
- **Source** — `.templ`, `.go`, seed files  

## Quick access: last 8 runs

Floating strip or command palette:

```text
Recent: [8] [7] [6] [5] [4] [3] [2] [1]
```

Each chip: hover shows label + time; click → preview at run.

## Modes of iteration (BuildY-specific)

| Mode | Creates PromptRun | Notes |
|------|-------------------|-------|
| Prompt | yes | Natural language → schema/source |
| Structured edit | yes | Block add/remove, field mapping |
| Content edit | optional | Minor content-only edits may batch or minor-run |
| Restore | yes | Always new run documenting revert |

**Planning before codegen:** allow conversational/planning step that does **not** emit `GeneratedChange` until user confirms — reduces eager codegen noise.

## Failed runs

- Keep failed run in timeline with error summary  
- Do not advance "current" pointer unless user restores prior success  
- Offer: retry with adjusted prompt | revert to last success  

## MVP vs later

| Feature | MVP | Later |
|---------|-----|-------|
| Run list + labels | ✓ | |
| Preview at run | ✓ | |
| Restore schema/content | ✓ | |
| File-level diff viewer | ✓ (basic) | side-by-side |
| Lazy infinite history | when >20 runs | |
| SSE run logs | | ✓ |
| Branch/fork runs | | ✓ |
| PR generation from run | | ✓ |

## Security notes

- Runs scoped to `project_id`; no cross-project leakage  
- Exported snapshots include run metadata for traceability  
- Restore requires project editor role (when auth exists)  

## Related

- [`../capabilities/INDEX.md`](../capabilities/INDEX.md)
- [`../messaging/prototypes-page.md`](../messaging/prototypes-page.md)
- [`../../architecture/02-preview-workspace.md`](../../architecture/02-preview-workspace.md)
