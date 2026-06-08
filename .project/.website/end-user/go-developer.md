# End User: Go Developer

## Profile

| Field | Detail |
|-------|--------|
| Archetype | Engineer who receives BuildY output and owns production quality |
| Typical titles | Backend Engineer, Full-Stack (Go), Platform Engineer, Tech Lead |
| Technical level | High — daily Go, Templ, Git, CI |
| Session length | 1–3 hours for initial review; ongoing IDE work |
| Success feeling | "I merged 80% of this and shipped in one sprint" |

## Primary workflow

```text
Notification: prototype approved for export
    → review generated file tree in BuildY or GitHub PR
    → run locally (Framework host + AppCMS)
    → fix gaps, add auth/billing/integrations in IDE
    → merge → deploy AppCMS runtime
```

## Tasks in BuildY

| Task | Workspace surface |
|------|-------------------|
| Assess export readiness | Generated files view, diffs |
| Validate architecture | Check imports, layer boundaries |
| Compare iterations | Snapshot diff between PromptRuns |
| Pull preview content | Export content snapshot for seed data |
| **Not** primary: | Full project authoring in browser |

## What they inspect first

```text
internal/views/          # page composition
internal/ui/blocks/      # section organisms
internal/appschema/      # content model hints
pkg/ / cmd/server/       # host wiring
go.mod                   # dependency sanity
```

## Motivations

- **Mergeable output** — idiomatic, small diffs  
- **Clear boundaries** — BuildY doesn't pollute Platform  
- **Honest preview** — BFF behavior matches what they'll ship  
- **No lock-in** — GitHub is source of truth  

## Frustrations to avoid

- Generated code ignores `.cursor/rules` / ui8px / templ specs  
- Hidden magic that doesn't appear in repo  
- Preview uses mocks that diverge from exported BFF  
- Massive unreviewable diffs on every prompt tweak  

## Handoff contract (ideal)

BuildY export includes:

- Runnable app skeleton  
- Documented content types and seed entries  
- List of intentional TODOs for eng  
- Test baseline (`go test ./...` passes)  

## Handoff contrast

Hosted AI builders often treat GitHub as optional. BuildY treats **GitHub + IDE as the normal continuation path** — Go/Templ/AppCMS output, eng persona **first-class**.

## Related

- Buyer: [`../buyer-persona/engineering-tech-lead.md`](../buyer-persona/engineering-tech-lead.md)
- ICP: [`../icp/icp-fastygo-product-team.md`](../icp/icp-fastygo-product-team.md)
