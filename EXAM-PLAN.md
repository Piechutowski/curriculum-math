# Discrete-Math Exam — Dated Study Plan

**The course:** eight weekend sessions (Sat–Sun), two lectures per session:
**Oct 10–11 · Oct 24–25 · Nov 7–8 · Nov 21–22 · Dec 5–6 · Dec 19–20 2026 · Jan 9–10 ·
Jan 23–24 2027.** Three-week holiday gap between sessions 6 and 7. **Assumed exam: February
2027 winter session** — correct this line when the date is announced. The lecture-to-session
mapping below is inferred from the lecture numbering; adjust it against the official syllabus
when published.

**Strategy, per topic:** Greene lessons first (foundation) → Rosen sections (the advanced,
rigorous treatment) → uni lecture notes to fill gaps — *except* asymptotics, recurrences and
generating functions, where the lectures exceed Rosen and are the main text. Every topic ends
the same way: **solve that topic's class exercises (`__ex__` files) closed-book**, then check
against `__sol__`. Being able to do the class problems without notes *is* being ready.

**The calendar's gift: an 8-week head start** (Aug 17 – Oct 9). Long enough to clear **every
Greene prerequisite — all ~60 lessons, De Moivre path and partial fractions included — plus
Rosen Ch. 1 in depth, before the first class.** During the semester, the only preparation left
is the Rosen section for the upcoming topic; every lecture becomes a second pass.

**Personal tick-list:** [EXAM-TODO.md](EXAM-TODO.md) mirrors this plan as checkboxes — outside
the tracking system, for crossing things off; real progress is still ticked in modules 01/02/03.

**Standing rules**

1. Exercises closed-book, always. Open `__sol__` only after a real attempt.
2. A topic is done when its `__ex__` problems work without notes — not when the reading is done.
3. Doubled lectures (6 vs 7A/B, 8–9, 10–11): primary = whichever version your semester actually
   delivers; the other is a source of extra worked examples.
4. Fall behind? Cut Rosen exercises before class exercises. Never skip the closed-book gates.

---

## Phase 0 — Head start (Aug 17 – Oct 9, before classes begin)

Eight clear weeks. Weeks 1–4 build the general foundation and the deep Rosen Ch. 1 logic/proofs
pass; weeks 5–8 clear the recurrence machinery — the semester's hardest dependency chain —
months before it is needed.

| Week | Greene | Rosen |
|---|---|---|
| Aug 17–23 | 146, 149, 152, 155 · 431, 433, 435, 437 (inequalities; sequences & series) | 1.1, 1.2, 1.3 (propositional logic) |
| Aug 24–30 | 266, 268, 270, 272, 297, 301 · 334, 336, 338, 340 · 445, 447 (functions, inverses, sets) | 1.4, 1.5, 1.6 (predicates, quantifiers, inference) |
| Aug 31 – Sep 6 | 344, 346, 352, 354, 356, 358, 360, 362, 364 (exponentials & logarithms) | **1.7, 1.8 — proofs, the in-depth week.** Work many exercises |
| Sep 7–13 | 439, 441, 443 (binomial, counting, probability) | 2.1, 2.2, 2.3 · *(optional: 1.9, Ch. 1 wrap-up)* |
| Sep 14–20 | 184, 187, 190 · 193, 196, 199, 202, 205 (complex numbers; the quadratic toolkit) | — |
| Sep 21–27 | 310, 312, 314, 316, 318, 324 (polynomial roots) · 369, 371 (systems) | — |
| Sep 28 – Oct 4 | 449, 457, 464, 466, 473, 477 (unit circle) · 533, 535, 537, 539 (polar form & **De Moivre**) | — |
| Oct 5–9 | **377, 379 (partial fractions — non-negotiable)** · buffer & review | 5.1, 5.2 (induction, ahead of the course) |

**By Oct 9 every Greene prerequisite is done.** Nothing below needs module 01 again.

## Semester — eight weekend sessions

Before each session: the Rosen prep from the previous row is done, so the lectures are a second
pass. After each session: review the lecture notes against your pass, then close the topic by
passing its exercise set closed-book before the next session.

| Session | Dates | Lectures (inferred) | Your gate before next session | Rosen prep for next session |
|---|---|---|---|---|
| **1** | Oct 10–11 | Lec 1–2: sets, functions | **Exercises 1** | 6.1, 6.2, 6.3, 6.5 (counting) |
| **2** | Oct 24–25 | Lec 3, 4A: induction, combinatorics | **Exercises 2** | 7.5, 7.6 (inclusion–exclusion) |
| **3** | Nov 7–8 | Lec 4B, 5: counting, pigeonhole | **Exercises 3** | 3.2, 3.3 (growth of functions) |
| **4** | Nov 21–22 | Lec 6, 7A/7B: asymptotics *(lectures = main text)* | **Exercises 4** | 7.1, 7.2 (recurrences) |
| **5** | Dec 5–6 | Lec 8–9: recurrences *(lectures = main text)* | **first half of Exercises 5** · complex-root recurrence via De Moivre, by hand | 7.3, 7.4 (generating functions) |
| **6** | Dec 19–20 | Lec 10–11: generating functions, D&C *(lectures = main text)* | **Exercises 5 complete** | 9.1, 9.2, 9.4, 9.5 (graphs) — holidays |
| **7** | Jan 9–10 | Lec 12, 13a/b: graphs | **Exercises 6** | 10.1, 10.2, 10.3 (trees) |
| **8** | Jan 23–24 | Lec 14, 14b, 15: trees + review | tree problems from Ex 6 + Rosen 10.1–10.3 exercises | — |

**The holiday gap (Dec 21 – Jan 8)** is the consolidation window, placed by luck right after
the hardest material: finish Exercises 5, redo every recurrence and generating-function problem
cold, read the graphs prep. Treat it as the first half of the exam run-up.

## Exam run-up (from Jan 25)

- **Week of Jan 25:** Lecture 15 (the course's own review) as a **closed-book mock exam**.
  Every miss names a topic — redo that topic's `__ex__` set and lecture examples.
- **Following week(s), to the exam:** second pass on the two weakest topics; re-sit the mock.
  Ready = **above 80% cold**. Recurrences and generating functions get priority in any spare
  hour — heaviest-weighted, hardest material.

---

## What this plan deliberately leaves out

Rosen Ch. 4, 8, 11–13, and the ~118 Greene lessons not listed above are **not on this exam**.
Most of trigonometry stays deferred, with one exception that *is* in the plan: the unit-circle
path to **polar form and De Moivre's theorem** (Greene 449–477 selection plus §25), because
recurrences with complex characteristic roots — an exam topic — are written in exactly that
form. Everything else returns to the normal roadmap (README) afterwards — the crypto, database
and compiler gates still run through it. Rosen Ch. 1 is in the plan by choice: it sharpens the
proof-writing this exam grades, and it was next on your list anyway.
