# Discrete-Math Exam — Dated Study Plan

**The course:** classes every 2 weeks starting **September 11**, two lectures per 2-week block —
15 numbered lectures across the semester. **Assumed exam: late January / early February**
(winter session) — correct this line when the date is announced, and adjust block boundaries to
the real syllabus once the course publishes it.

**Strategy, per topic:** Greene lessons first (foundation) → Rosen sections (the advanced,
rigorous treatment) → uni lecture notes to fill gaps — *except* asymptotics, recurrences and
generating functions, where the lectures exceed Rosen and are the main text. Every topic ends
the same way: **solve that topic's class exercises (`__ex__` files) closed-book**, then check
against `__sol__`. Being able to do the class problems without notes *is* being ready.

**The one rule this calendar adds: stay one block ahead.** Finish a topic's Greene + Rosen
*before* the uni block that covers it. Then the lecture is a second pass, the class is a third,
and the exercise session confirms rather than teaches.

**Standing rules**

1. Exercises closed-book, always. Open `__sol__` only after a real attempt.
2. A topic is done when its `__ex__` problems work without notes — not when the reading is done.
3. Doubled lectures (6 vs 7A/B, 8–9, 10–11): primary = whichever version your semester actually
   delivers; the other is a source of extra worked examples.
4. Fall behind? Cut Rosen exercises before class exercises. Never skip the closed-book gates.

---

## Phase 0 — Head start (Aug 17 – Sep 10, before classes begin)

Four weeks with no lectures competing for attention. Foundation plus the deep logic/proofs pass
from Rosen Ch. 1 — not on the exam directly, but it is the language every proof this course
demands is written in, and induction (a certainty on the exam) is a proof method.

| Week | Greene | Rosen |
|---|---|---|
| Aug 17–23 | 146, 149, 152, 155 · 431, 433, 435, 437 (inequalities; sequences & series) | 1.1, 1.2, 1.3 (propositional logic) |
| Aug 24–30 | 266, 268, 270, 272, 297, 301 · 334, 336, 338, 340 · 445, 447 (functions, inverses, sets) | 1.4, 1.5, 1.6 (predicates, quantifiers, inference) |
| Aug 31 – Sep 6 | 344, 346, 352, 354, 356, 358, 360, 362, 364 (exponentials & logarithms) | **1.7, 1.8 — proofs, the in-depth week.** Work many exercises; this is the payoff of the head start |
| Sep 7–10 | 439, 441, 443 (binomial, counting, probability) | 2.1, 2.2, 2.3 (sets & functions — now you are ahead of Lectures 1–2) |

*Optional if time remains: Rosen 1.9 (normal forms) and the Ch. 1 wrap-up exercises.*

## Semester — biweekly blocks (from Sep 11)

Each block: the uni delivers two lectures; you have already done the Greene + Rosen row *before*
the block starts (the "prepare" column of the *next* block is that block's homework). Close
every block by passing its exercise set closed-book.

| Block | Uni delivers | Your gate | Prepare for next block (Greene + Rosen) |
|---|---|---|---|
| **1** · Sep 11–24 | Lec 1–2: sets, functions | **Exercises 1** | Rosen 5.1, 5.2 (induction — Greene done in Phase 0) |
| **2** · Sep 25 – Oct 8 | Lec 3, 4A: induction, combinatorics | **Exercises 2** | Rosen 6.1, 6.2, 6.3, 6.5 · Greene 184, 187, 190, 193, 196, 199, 202, 205 (complex numbers & quadratics — early, they feed recurrences) |
| **3** · Oct 9–22 | Lec 4B, 5: counting, pigeonhole | **Exercises 3** | Rosen 3.2, 3.3 · Greene 310, 312, 314, 316, 318, 324 (polynomial roots) · Greene 449, 457, 464, 466, 473, 477 (unit-circle trig — the run-up to De Moivre) |
| **4** · Oct 23 – Nov 5 | Lec 6, 7A/7B: asymptotics *(lectures = main text)* | **Exercises 4** | Rosen 7.1, 7.2 · Greene 369, 371 (systems) · **Greene 533, 535, 537, 539 (polar form & De Moivre — required for recurrences with complex characteristic roots)** · **377, 379 (partial fractions — non-negotiable before generating functions)** |
| **5** · Nov 6–19 | Lec 8–9: recurrences *(lectures = main text)* | **first half of Exercises 5** | Rosen 7.3, 7.4 |
| **6** · Nov 20 – Dec 3 | Lec 10–11: generating functions, D&C *(lectures = main text)* | **Exercises 5 complete** | Rosen 9.1, 9.2, 9.4, 9.5 (graphs) |
| **7** · Dec 4–17 | Lec 12, 13a/b: graphs | **Exercises 6** | Rosen 10.1, 10.2, 10.3 (trees) |
| **8** · Dec 18 – Jan 7 *(holidays)* | Lec 14, 14b: trees | tree problems from Ex 6 + Rosen 10.1–10.3 exercises | — |

## Exam run-up (January)

- **Early January:** Lecture 15 (the course's own review) as a **closed-book mock exam**. Every
  miss names a topic — redo that topic's `__ex__` set and lecture examples.
- **Mid January:** second pass on the two weakest topics; re-sit the mock. Ready = **above 80%
  cold** on Lecture 15's problems. Recurrences and generating functions get priority in any
  spare hour — they are the heaviest-weighted, hardest material.

---

## What this plan deliberately leaves out

Rosen Ch. 4, 8, 11–13, and the ~118 Greene lessons not listed above are **not on this exam**.
Most of trigonometry stays deferred, with one exception that *is* in the plan: the unit-circle
path to **polar form and De Moivre's theorem** (Greene 449–477 selection plus §25), because
recurrences with complex characteristic roots — an exam topic — are written in exactly that
form. Everything else returns to the normal roadmap (README) afterwards — the crypto, database
and compiler gates still run through it. Rosen Ch. 1 is in the plan by choice: it sharpens the
proof-writing this exam grades, and it was next on your list anyway.
