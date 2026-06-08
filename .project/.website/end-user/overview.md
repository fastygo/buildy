# End User — Overview

End users **operate BuildY weekly** — prompts, schema edits, preview review, export — regardless of who signed the contract.

## End user vs buyer

| | Buyer persona | End user |
|---|---------------|----------|
| Question | "Should we adopt BuildY?" | "How do I ship this landing?" |
| Frequency | Monthly / at renewal | Daily / weekly during projects |
| Success | ROI, stack fit, risk | Tasks completed, preview shared, export merged |

## Persona map

| Persona | Typical title | Primary jobs |
|---------|---------------|--------------|
| [Prototype Author](prototype-author.md) | PM, founder, designer, agency PM | Create project, prompt, edit model, share preview |
| [Content Marketer](content-marketer.md) | Content/growth marketer | Edit blocks, copy, SEO fields, locales |
| [Go Developer](go-developer.md) | Engineer, design engineer | Review generated files, merge PR, extend in IDE |

## Core jobs-to-be-done (workspace)

1. **Create project** — name, template, theme baseline  
2. **Describe product** — prompt, PRD paste, pattern selection  
3. **Shape model** — pages, blocks, content types, field mappings  
4. **Preview behavior** — BFF-backed navigation and content CRUD  
5. **Review source** — file tree, diffs, generated Templ  
6. **Export / sync** — snapshot or GitHub push  
7. **Iterate** — prompt runs, schema changes, re-preview  

## UX implications (from `.project`)

Minimum screens align with end-user flow:

- project list  
- project detail  
- prototype workspace shell  
- generated files view  
- preview placeholder → live preview  

BuildY is **not** optimized for arbitrary file editing — end users who need that switch to IDE (Go Developer persona).

## Permission model (conceptual)

| Action | Prototype author | Content marketer | Go developer |
|--------|------------------|------------------|--------------|
| Create project | ✓ | ○ | ✓ |
| Edit schema/blocks | ✓ | ○ | ✓ |
| Edit preview content | ✓ | ✓ | ○ |
| View generated files | ✓ | ○ | ✓ |
| Export / sync GitHub | ○ | ○ | ✓ |
| Merge to production | ○ | ○ | ✓ |

(○ = sometimes / with approval)

## Related docs

- [`../buyer-persona/overview.md`](../buyer-persona/overview.md)
- [`../../architecture/02-preview-workspace.md`](../../architecture/02-preview-workspace.md)
