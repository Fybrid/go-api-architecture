# AGENTS.md

This is a lightweight guide for agents to work safely and efficiently in this repository. Aim for minimal, focused changes that align with the existing structure and style.

## Principles
- Scope minimization: Focus strictly on the request; avoid unrelated changes.
- Consistency: Follow existing naming, layering, and error-handling patterns.
- Readability: Prefer small, clear functions and comments; avoid unnecessary abstractions.

## Workflow
- Clarify impact and proceed with small patches.
- Skim only relevant files before implementing.
- After changes, build/run quickly and share a brief summary.

## Go Guidelines
- Error handling: Use `if err != nil { ... }`; keep wrapping/logging minimal but contextual.
- Context: Pass `context.Context` from callers and respect cancellation.
- Layers: Keep separation across HTTP router → handler → (service/usecase) → repository.
- Formatting: Assume `gofmt`-equivalent formatting (enable editor auto-format).

## Run & Verify
- Entry point: `cmd/app/main.go`
- Run: `go run ./cmd/app`
- Tests: `go test ./...` (if tests exist)

## Review Checklist (good to include in final note)
- Purpose of change and evidence it meets acceptance criteria
- Files touched and key modifications
- How to verify (commands, endpoints, expected results)
- Impact analysis, known limitations, and next steps

## Do Nots / Cautions
- No unrelated refactors, dependency additions, or license changes.
- No large renames or file moves without prior agreement.
- Ask before using external networks or performing destructive operations.

## Clarify With Requestor (when needed)
- Expected inputs/outputs, reproduction steps, acceptance criteria
- In/Out of scope and priorities

## Working Checklist
- Confirmed goals and acceptance criteria
- Implemented with minimal diffs (no unrelated changes)
- Verified locally via run/test (as applicable)
- Updated README or docs when necessary
- Summarized impact and follow-up actions
