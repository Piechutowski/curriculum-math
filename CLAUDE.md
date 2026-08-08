# CLAUDE.md — how this repository works

This file explains the whole system, for humans and for Claude sessions working in this repo.

## What this repository is

A personal mathematics curriculum with progress tracking. The owner (Szymon) is building the
math foundation for five programming goals: 2D graphics & physical simulation `[GFX]`, compiler
engineering `[CMP]`, computer networking `[NET]`, relational databases `[DB]`, and cryptography
`[CRY]`. The curriculum is modeled on Mr. Greene's college math course: one atomic skill per
lesson, and every lesson means three things — **Lesson → Study Notes → Practice Test**. A
checkbox is ticked only when all three are done.

## The files

| Path | What it is | Edited by |
|---|---|---|
| `README.md` | Roadmap, book list, study phases, "when can I start subject X" gates, progress dashboard | Hand, except the marker block |
| `01-algebra-trigonometry-map.md` | Checklist + study map for Mr. Greene's course (220 lessons, his numbering) | Tick checkboxes |
| `02-discrete-mathematics-map.md` | Checklist + study map for Rosen & Krithivasan (91 items) | Tick checkboxes |
| `03-calculus.md` | Full calculus curriculum, 89 lessons (Thomas' Calculus) | Tick checkboxes |
| `04-linear-algebra.md` | Full geometric linear algebra curriculum, 52 lessons (Kumaresan) | Tick checkboxes |
| `05-probability.md` | Full probability curriculum, 46 lessons (Veerarajan) | Tick checkboxes |
| `06-numerical-methods.md` | Bridging module, 18 lessons (Sastry) | Tick checkboxes |
| `LOG.md` | Completion log, newest day first | **Generated — never by hand** |
| `progress/log.csv` | Raw log: timestamp, module, lesson for every completion | **Generated — never by hand** |
| `progress/*.svg` | Progress-bar cards (one per module + `overall.svg`) | **Generated — never by hand** |
| `progress/progress.csv` | Raw counts per module | **Generated — never by hand** |
| `tools/progress/main.go` | The Go tool that generates all of the above | Hand |
| `.github/workflows/progress.yml` | GitHub Action that runs the tool on every markdown push | Hand |

## How progress tracking works

Every lesson is a markdown checkbox (`- [ ]` → `- [x]`). The Go tool
(`go run ./tools/progress`, run from the repo root, needs Go ≥ 1.21, stdlib only) does four
things on each run:

1. Counts checkboxes in the six module files.
2. Redraws the SVG progress cards in `progress/` and rewrites `progress/progress.csv`.
3. Rewrites the content between `<!-- progress:... -->` and `<!-- endprogress -->` markers in
   each module file and in `README.md` (the "N / M lessons complete" lines and the dashboard).
4. Reconciles the completion log (see below) and regenerates `LOG.md`.

**Never edit generated files or the text between the marker comments by hand** — the next run
overwrites it. Never remove the marker comments themselves; the tool fails loudly if a marker
is missing.

## How the log works

`progress/log.csv` is the source of truth for *when* each lesson was completed:

- When the tool sees a checkbox ticked that was not in the log, it appends an entry stamped
  with the current UTC time (minute precision).
- When a previously logged lesson is found unticked, its entry is removed (progress reverted).
- Entries keep their original timestamp forever — re-running the tool never re-stamps them.
- `LOG.md` is rendered from the CSV: grouped by date, newest day first, no time of day shown
  (full timestamps stay in the CSV).

A lesson is identified by module + label. The label is the checkbox line's bold span when it
has one (modules 03–06), otherwise the whole line (Greene and Rosen entries). **Renaming a
lesson line therefore drops and re-stamps its log entry** — avoid renaming ticked lessons.

## How to record completed lessons (three ways)

1. **On github.com, no tools needed:** open the module file → pencil icon → change `[ ]` to
   `[x]` → commit. The GitHub Action reruns the tool and pushes a follow-up commit
   (`[progress-bot]`) with updated bars and log within about a minute. The completion date is
   stamped server-side at that moment.
2. **Locally:** tick the boxes in an editor, then `go run ./tools/progress`, then commit and
   push everything it changed.
3. **Tell Claude:** list what you finished (e.g. "I did 131 and 134, and all of Section 4").
   Claude ticks the boxes, runs the tool, commits, and pushes.

Batching is fine — tick any number of boxes in one edit; they all get logged in that run.

## Rules for Claude sessions

- Work on the designated branch; commit and push after changes.
- After any checkbox change, always run `go run ./tools/progress` before committing, so the
  generated files stay in sync with the checklists.
- When the owner reports completed lessons in chat, tick exactly those, no more. Lesson
  numbers in module 01 are Mr. Greene's own course numbers (1–550 with gaps), not sequential.
- Don't untick or rename ticked lessons unless explicitly asked (it rewrites log history).
- If a new module or lesson list is added, register it in the `modules` table at the top of
  `tools/progress/main.go` (id, file, title, short name, gradient colors) and add the marker
  block near the top of the new file:

  ```markdown
  <!-- progress:NEW-ID -->
  <!-- endprogress -->
  ```

- Curriculum content conventions: atomic lessons; one checkbox per skill; goal tags
  `[GFX] [CMP] [NET] [DB] [CRY]` only where the payoff is direct (no tag = load-bearing
  foundation); every module ends with exit criteria; full standard course coverage — tags say
  *why*, never *whether*.

## Gotchas

- The Action skips commits whose message contains `[progress-bot]` — that's the loop guard.
  Don't put that string in normal commit messages.
- The tool must run from the repository root (it checks for `README.md`).
- Two checklists intentionally differ in granularity: module 01 = one box per Greene lesson;
  module 02 = one box per Rosen book section plus one wrap-up box per chapter (review
  questions + supplementary exercises + one Computer Project).
- The repo currently has no `main` branch; all work lives on the working branch until the
  owner promotes or merges it.
