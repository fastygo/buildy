# BuildY

**Your repo. Your decisions. Your product.**

BuildY is a browser workspace for creating software products. You describe the product, shape the model, preview real behavior, review the generated source, and export everything into **your own GitHub repository** as code you control.

BuildY does not claim ownership of the products you create. You stay the founder, the decision maker, and the rights holder of what you ship.

## You own what you build

You connect your GitHub (or work locally), choose the dependencies and patterns, write the prompts, define the UX, content models, business logic, and final project layout. The result is a standalone product in your repository.

You are the author and rights holder of the software product in the parts created through your creative decisions, prompts, configurations, UX/UI, business logic, copy, settings, and the final composition of the project. Rights to third-party components, libraries, frameworks, and templates remain with their respective rights holders and are governed by their applicable licenses.

You choose the license for your product, provided it is compatible with the licenses of the components you include.

**Your repo. Your decisions. Your product.**  
BuildY is a browser workspace for creating software products. You own what you build, subject to the licenses of the components you choose.

### Quick license compatibility notes
- Permissive dependencies (MIT, Apache, BSD) + your product under AGPLv3, commercial, or any other license you prefer is typically fine (preserve notices where required).
- Copyleft components (GPL/AGPL) inside your product generally require the product to follow compatible terms when distributed or offered as a service.
- Components without a clear license should be avoided — "it's on GitHub" does not grant usage rights.
- Source-available / BSL-style components have their own specific terms; read them.

BuildY will never promise "any license always works." The rule is simple: **you select a license for your product that respects the licenses of the pieces you chose**.

## How it works

1. **Start** — Open a project with a prompt ("SaaS landing for a developer tool with usage-based billing") or pick a proven starting pattern.
2. **Model** — Work with structured product building blocks: pages, sections, content types, fields, relations, themes, and connector contracts. Everything that should be editable later is wired to a content model from the start.
3. **Preview with real behavior** — See a live, CMS-backed preview. Edit content, submit forms, filter lists, exercise validation and flows. The preview is not a static mockup — it runs against the same BFF patterns and content runtime your final app will use.
4. **Iterate with intent** — Use prompts to evolve the product. BuildY proposes precise, reviewable changes to the model. You accept, tweak, or revert at the model level; generated source follows.
5. **Inspect the source** — At any moment you can browse the generated Go code, Templ templates, content schemas, BFF contracts, seed data, and manifests. It is readable, conventional code — not a black box.
6. **Export to your repo** — Push a clean snapshot, open a branch, or create a PR in your connected GitHub repository. From that point the code is yours in your repository.
7. **Ship and keep iterating** — Deploy the generated app (it includes a full content runtime). Edit live content in the CMS admin or pull changes back into BuildY for the next round of product work. Developers continue in their normal Git + IDE workflow.

BuildY is the fast loop for product discovery, modeling, and generation. Your repository is where the product lives long-term.

## What you actually get

- A real, deployable Go project using the patterns that power production FastyGo-style applications: explicit BFF layer, type-safe Templ UI, and a structured content model.
- CMS-backed landings and apps by default. Hero text, features, pricing, testimonials, FAQ — every visible block maps to typed content fields and collections you (or your customers) can edit after launch without touching code.
- Behavior preview that validates product logic before you commit to the repository.
- Full visibility into every generated file and the exact diff a change will produce.
- Clean export that does not require you to keep using BuildY to run the product.

The goal is not to trap you in a proprietary builder. The goal is to give you a high-quality, operational starting point that you own completely.

## Who this is for

- Founders who want to ship a real product, not another Figma-to-code handoff or a throwaway frontend prototype.
- Product teams that need both fast iteration in the browser and a maintainable backend they can reason about.
- Developers who like Go, value type-safe templates, and want the content layer to be a first-class part of the product instead of an afterthought.
- Anyone tired of AI tools that output "something that looks like the app" but leaves you with a rewrite or vendor lock-in.

You make the product decisions. BuildY handles the mechanical generation, wiring, and preview so you can move faster without giving up control.

## Technical approach (for the people who ship the code)

BuildY is a structured generation workspace, not a general-purpose code execution environment (at least in the early phases).

- **Project model** — Serializable definitions for projects, pages, blocks, content types, fields, themes, connectors, and generation runs.
- **Generation pipeline** — Prompt / intent → product plan → app schema → CMS content model + mappings → BFF screen models → Templ blueprints → concrete source files and manifests.
- **Preview runtime** — Schema-driven L1 (structure), behavior-aware L2 (BFF actions + preview content store), and source L3 (generated file tree + diffs). No arbitrary user code execution required for the core product-design loop.
- **Three independent sync streams** — Code (GitHub owns source after export), Content (preview store ↔ live AppCMS), Schema (model → content type definitions/migrations).
- **Output is conventional code** — Generated projects are meant to be readable, testable, and extendable in a normal Go + Templ + AppCMS workflow.

Later layers (GitHub App automation, isolated build workers, deployment targets, billing, and entitlements) are planned but intentionally deferred until the core prompt-to-preview-to-export loop is proven.

## Current status

BuildY is in active early development. The immediate focus is the end-to-end prototype:

prompt + structured model → CMS-backed behavior preview → readable generated source → exportable project snapshot.

Detailed architecture, invariants, roadmap, and decision register live in the `.project/` directory. This is the single source of truth for direction and scope.

No billing, no GitHub App, and no arbitrary code execution sandboxes in the first pass. Those come after the product loop is solid.

## License

BuildY itself is licensed under the **GNU Affero General Public License v3.0** (AGPLv3).

See the [LICENSE](./LICENSE) file for the full text.

Key implications for BuildY (the tool):

- You can run, study, modify, and redistribute BuildY.
- If you operate a modified version of BuildY as a hosted service, the corresponding source must be offered to users interacting with that service.
- BuildY does **not** claim any ownership or special rights over the products you create inside it, beyond the narrow technical permissions required to display, generate, edit, build, test, deploy, and synchronize code at your explicit direction.

Your generated products are licensed by **you**, under terms you choose, subject only to the license compatibility of the components you elected to include.

## "BuildY doesn't own my product" (for legal/compliance folks)

BuildY requires only the technical rights necessary to operate the service you asked for: rendering previews, performing generation, showing diffs, syncing on your command, and similar operations on content and source you control.

BuildY asserts no copyright, ownership, or ongoing rights over the creative output (prompts, model decisions, final compositions, copy, UX flows, business logic) that constitute your product.

When you export to your repository, the code in that repository is yours to license, sell, modify, or host under whatever terms are compatible with the included dependencies.

## Contributing

The project is young. Architecture and scope are documented in `.project/`. Once the core workspace, CMS-backed generation, preview, and export are stable, external contributions will be welcome.

Until then, the fastest way to help is to stress-test the direction against real product-building workflows and report gaps in the decision register or roadmap.

## Links

- Issues and discussion: (repository issues once public)
- Architecture and progress: `.project/` in this repo
- Generated apps target the same open patterns used across the FastyGo ecosystem (Framework, Platform, Templ, AppCMS) — but your exported project stands on its own.

Build real products.  
Own the repo.  
Keep the decisions.
