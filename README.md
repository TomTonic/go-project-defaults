# Go Project Defaults

Standard configurations, workflows, and templates for consistent, secure Go projects.

## Copy or Reference?

Most tools (golangci-lint, pre-commit, GitHub Actions, editors) require config
files to live **inside the consuming repository**. There is no way around
copying those files. However, Renovate supports **shared presets** that can be
referenced from a single line — so you get central control without copying.

| Method | Files | What it means |
|---|---|---|
| **Reference** (single source of truth) | Renovate config | Your project's `renovate.json` contains one `extends` line pointing here. Changes to the preset take effect everywhere automatically. |
| **Copy** (into each repo) | Everything else | Tools require these files locally. Copy them once, then customize `TODO:` markers. Renovate keeps the SHA-pinned action versions up to date. |

## What's Included

### Configuration Files (copy into your repo)

| File | Purpose |
|---|---|
| `.editorconfig` | Consistent editor settings (tabs for Go, spaces for YAML/JSON, LF line endings) |
| `.gitattributes` | Consistent line endings and binary file handling |
| `.gitignore` | Standard Go ignores (binaries, test artifacts, IDE files, `dist/`) |
| `.golangci.yml` | Balanced golangci-lint v2 config (errcheck, govet, gosec, revive, and more) |
| `.goreleaser.yml` | Cross-platform release builds — **only for CLI projects** |
| `.pre-commit-config.yaml` | Pre-commit hooks: gitleaks, golangci-lint, whitespace, YAML check |
| `.yamllint.yml` | YAML linting rules (120 char warning, truthy values, comments) |

### Renovate (reference — do NOT copy `default.json`)

The shared Renovate preset lives in `default.json` in this repository.
**Do not copy it.** Instead, put this minimal `renovate.json` in your project:

```json
{
  "$schema": "https://docs.renovatebot.com/renovate-schema.json",
  "extends": [
    "github>TomTonic/go-project-defaults"
  ]
}
```

That single `extends` line pulls in the full config: automerge rules,
grouping, commit message format, vulnerability alerts, and more.
Any change to `default.json` in this repo takes effect across all
projects on the next Renovate run.

You can override any setting in your project's `renovate.json`:

```json
{
  "extends": ["github>TomTonic/go-project-defaults"],
  "schedule": ["after 10pm"],
  "packageRules": [
    {
      "description": "Project-specific: pin this dependency",
      "matchPackageNames": ["example.com/some-lib"],
      "allowedVersions": "<2.0.0"
    }
  ]
}
```

A ready-to-copy example is provided as `renovate.json` in this repository.

### Workflows (copy into your repo's `.github/workflows/`)

| File | Trigger | Purpose |
|---|---|---|
| `ci.yml` | push, PR | Lint (golangci-lint) + go mod tidy check + test with race detector |
| `coverage.yml` | push | Test coverage with gist-based badge (no dead branch!) |
| `codeql.yml` | push to main, PR, weekly | GitHub CodeQL security analysis |
| `dependency-review.yml` | PR | Flag known-vulnerable dependencies in PRs |
| `scorecard.yml` | push to main, weekly | OpenSSF Scorecard supply-chain audit |
| `grype-me.yml` | daily, manual | Vulnerability scan via [grype_me](https://github.com/TomTonic/grype_me) with badge |
| `fuzz.yml` | manual (schedule disabled) | Go fuzz testing for all `Fuzz*` targets |
| `release.yml` | version tag (`v*`) | GoReleaser binary builds — **only for CLI projects** |

### Markdown Templates (copy into your repo)

| File | Purpose |
|---|---|
| `SECURITY.md` | Vulnerability reporting instructions (GitHub private reporting) |
| `CONTRIBUTING.md` | Contribution guidelines, commit conventions, testing standards |
| `AGENTS.md` | AI coding assistant guidelines (Codex, Claude, etc.) |
| `.github/copilot-instructions.md` | GitHub Copilot-specific project instructions |

## Quick Start

1. **Renovate**: Add `renovate.json` with the one-line `extends` reference (see [Renovate section](#renovate-reference--do-not-copy-defaultjson) above). Done — no copying needed.
2. **Everything else**: Copy the files you need into your Go project.
3. Search for `TODO:` comments and customize (project name, thresholds, paths).
4. Set up required secrets and variables (see below).
5. Let Renovate run once — it will pin any unpinned action tags to SHAs.

## Secrets & Variables Setup

All badges use **gist-based storage** — no dead branches, no external services.
You can reuse a single gist for all badges across all your projects.

### Step-by-step

1. **Create a GitHub Gist** at [gist.github.com](https://gist.github.com/) with any initial file (e.g. `init.txt` containing `{}`).
2. **Create a PAT** at [GitHub Settings → Tokens](https://github.com/settings/tokens) with `gist` scope.
3. **Add to each repository:**

| Type | Name | Value | Used by |
|---|---|---|---|
| Secret | `GIST_TOKEN` | Your gist PAT | coverage, grype_me |
| Variable | `COVERAGE_GIST_ID` | Gist ID | coverage.yml |
| Variable | `GRYPE_BADGE_GIST_ID` | Gist ID (can be the same) | grype-me.yml |

### Badge Markdown for Your README

**Coverage:**
```markdown
[![Coverage](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/<USER>/<GIST_ID>/raw/<REPO>-coverage.json)](https://gist.github.com/<USER>/<GIST_ID>)
```

**Vulnerabilities:**
```markdown
[![Vulnerabilities](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/<USER>/<GIST_ID>/raw/<REPO>-grype.json)](https://gist.github.com/<USER>/<GIST_ID>#file-<REPO>-grype-md)
```

Replace `<USER>`, `<GIST_ID>`, and `<REPO>` with your values.

## Customization Notes

### Renovate: Automerge Strategy

The shared preset (`default.json`) automerges **patch and minor** updates
when CI is green. This is a good balance for most Go projects — the Go
ecosystem has strong backward compatibility guarantees.

To switch a specific project to a more conservative strategy (**patches only**,
manual review for minor), add overrides in that project's `renovate.json`:

```json
{
  "extends": ["github>TomTonic/go-project-defaults"],
  "packageRules": [
    {
      "matchUpdateTypes": ["minor"],
      "automerge": false,
      "reviewersFromCodeOwners": true
    }
  ]
}
```

The preset's automerge rule still applies to patches; the project-level
override disables it for minor updates only.

### Coverage Threshold

Default: **80%**. Change the `COVERAGE_THRESHOLD` env variable in
`.github/workflows/coverage.yml`.

### Fuzz Testing

The fuzz workflow is **disabled by default** (manual trigger only).
To enable scheduled runs:

1. Make sure your project has fuzz tests (`func FuzzXxx(f *testing.F)`).
2. Uncomment the `schedule` trigger in `.github/workflows/fuzz.yml`.
3. Adjust `FUZZ_TIME` to control how long each target runs.

### GoReleaser (CLI Projects Only)

The `.goreleaser.yml` and `release.yml` workflow are for CLI tools that produce
binaries. **Libraries don't need these** — just use git tags.

Customize `.goreleaser.yml`:
- Set `project_name`.
- Adjust `goos`/`goarch` for your target platforms.
- Modify `ldflags` to match your version variables.

## Files You Probably Don't Need

| Skip if... | Files |
|---|---|
| Library (no binaries) | `.goreleaser.yml`, `.github/workflows/release.yml` |
| No fuzz tests | `.github/workflows/fuzz.yml` |
| No AI agent interaction expected | `AGENTS.md` |

## Design Decisions

- **SHA-pinned actions** with version comments. Renovate updates these automatically.
- **`step-security/harden-runner`** in every workflow for supply-chain hardening.
- **Go version from `go.mod`** (`go-version-file`) — no version duplication.
- **Gist-based badges** instead of branch-based — no dead `badges` branch cluttering your repo.
- **`katexochen/go-tidy-check`** in CI catches forgotten `go mod tidy`.
- **golangci-lint v2** config with shadow detection, gosec, revive, and gocritic — strict enough to catch real problems, relaxed enough not to annoy.
- **`grype_me`** runs daily against `latest_release` for continuous vulnerability monitoring.
