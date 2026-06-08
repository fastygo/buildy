# BuildY Website — Audience Model

Status: product reference  
Scope: target audience, ICP, buyer personas, end users, and target market for BuildY marketing and UX.

## Vocabulary

| Term | Definition |
|------|------------|
| **Target Audience** | Broad set of people and organizations BuildY is designed to serve. |
| **ICP** (Ideal Customer Profile) | Composite profile of a *company* that fits BuildY best — size, stack, workflow, and buying context. |
| **Buyer Persona** | A specific *decision-maker* inside the ICP who evaluates, approves, or purchases BuildY. |
| **End User** | The person who works in BuildY day to day — prompts, edits models, previews, exports. |
| **Target Market** | Market segments BuildY prioritizes now and later, with explicit MVP focus. |

## Product anchor

BuildY is a **prototype workspace**, not a browser IDE. It generates **deployable Go/Templ/AppCMS** applications with CMS-backed content, BFF preview, and GitHub-owned source continuity.

First proof:

```text
prompt -> schema/pattern -> AppCMS-backed landing -> BFF/Templ preview -> generated source export
```

See `.project/INDEX.md` and `architecture/00-north-star.md` for stack truth.

## Documents

### Audience

| Folder | Role |
|--------|------|
| [`target-audience/`](target-audience/overview.md) | Who BuildY serves at the highest level |
| [`icp/`](icp/overview.md) | Ideal company profiles |
| [`buyer-persona/`](buyer-persona/overview.md) | Decision-makers who buy or adopt |
| [`end-user/`](end-user/overview.md) | Daily operators inside the workspace |
| [`target-market/`](target-market/overview.md) | Prioritized market segments |

### Capabilities & messaging

| Folder | Role |
|--------|------|
| [`capabilities/`](capabilities/INDEX.md) | What BuildY does — capability map and MVP scope |
| [`messaging/`](messaging/homepage-ia.md) | Public site IA — home and prototypes pages |
| [`templates/`](templates/saas-landing-blocks.md) | First template spec — blocks and content types |
| [`workspace/`](workspace/version-history-ux.md) | PromptRun history, revert, diff UX |
| [`competitive/`](competitive/comparison-framework.md) | Positioning vs AI builders, IDE agents, no-code |

## Product stance

BuildY is **not** a clone of any AI app builder. It targets **prototype → CMS-backed model → export** on the FastyGo stack, with **source and license continuity** for the customer (GitHub-owned repo, readable Go/Templ output).

UX research may borrow patterns from the broader AI-builder category (Lovable, Bolt, v0, Base44, and others) without anchoring docs or messaging to one vendor. Competitive mentions belong in [`competitive/comparison-framework.md`](competitive/comparison-framework.md) only.

## Usage

- **Homepage / positioning:** `messaging/homepage-ia.md` + `target-audience/` + `competitive/`.
- **Product pages:** `messaging/prototypes-page.md` + `capabilities/INDEX.md`.
- **First template implementation:** `templates/saas-landing-blocks.md`.
- **Workspace UX:** `workspace/version-history-ux.md` + `end-user/`.
- **Sales / onboarding:** `icp/` + relevant `buyer-persona/`.
