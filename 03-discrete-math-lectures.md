# Module 03 — Discrete Mathematics · University Lecture Course

<!-- progress:03-discrete-math-lectures -->
![03 · Discrete Math — Lecture Course progress](progress/03-discrete-math-lectures.svg?v=3-0-27)

**0 / 27 lessons complete · 0%** — 0 of 9 sections finished · [completion log](LOG.md)
<!-- endprogress -->

**Resource:** your own [`discreat-math`](https://github.com/Piechutowski/discreat-math) repository —
21 lectures transcribed to markdown from the original *Matematyka Dyskretna* slides, 6 exercise
sessions with full solutions, and the source PDFs under `pdf/`.
**Prerequisites:** none beyond Greene §§1–5. Runs in parallel with module 02.
**Why this module exists:** this is a real taught course covering the same ground as Rosen from a
different angle — heavier on recurrences, asymptotics and generating functions, lighter on logic
and number theory. Rosen is the reference; these lectures are the classroom pass with worked
examples and graded exercises. Where they overlap, the second pass is what makes the material
stick.

**The rhythm applies to every lesson:** Lesson → Study Notes → Practice Test. Here the exercise
sessions *are* the practice tests: work the `__ex__` file first, then check against the `__sol__`
file — never the other way round.

Tags: `[GFX]` graphics/simulation · `[CMP]` compilers · `[NET]` networking · `[DB]` databases ·
`[CRY]` cryptography. No tag = foundation later lessons stand on.

**Note on doubled lectures:** the course delivered lectures 8–11 twice — a combined
Polish-original version and split English versions — and asymptotics likewise appears as lecture
6 (PL) and lecture 7 A/B (EN). They are not translations of each other; the examples differ
(Tower of Hanoi vs. sequential search, for instance). Both are listed, and doing both is a
legitimate second pass. If you would rather not repeat the topic, tick the version you actually
studied and leave the other; the bar will simply read as a partly-done section.

---

## Section 1: Sets and Functions `[DB]` `[CMP]`
The vocabulary everything else is written in. Sets are the raw material of the relational model;
functions formalize every mapping you will program. Pairs with Rosen Ch. 2.

- [ ] **Lecture 1 — Sets** `[DB]` — elements, subsets, power sets, Venn reasoning, set algebra.
  *(`01__lec__sets.md`)*
- [ ] **Lecture 2 — Functions** `[CMP]` — mappings, injective/surjective/bijective, composition,
  inverses. *(`02__lec__functions.md`)*
- [ ] **Exercises 1 — Set Operations** — practice test for this section.
  *(`01__ex__set-operations.md`, solutions in `01__sol__`)*

## Section 2: Mathematical Induction `[CMP]`
The proof technique compiler work runs on: prove it for the base case, prove each step preserves
it. Structural induction over syntax trees is this idea one level up. Pairs with Rosen Ch. 5.

- [ ] **Lecture 3 — Mathematical Induction** `[CMP]` — the domino metaphor, base case and
  inductive step, strong induction. *(`03__lec__mathematical-induction.md`)*
- [ ] **Exercises 2 — Set Families and Induction** — practice test.
  *(`02__ex__set-families-induction.md`)*

## Section 3: Counting and Combinatorics `[CRY]` `[DB]` `[NET]`
Keyspace sizes, hash collisions, and the substrate of all probability. Pairs with Rosen Ch. 6.

- [ ] **Lecture 4A — Combinatorics: Theoretical Supplement, Part I** — injective functions and
  counting arguments built on them. *(`04a__lec__combinatorics.md`)*
- [ ] **Lecture 4B — Counting: Sum Rule, Product Rule, Inclusion–Exclusion** `[DB]` — the three
  rules everything else is assembled from. *(`04b__lec__inclusion-exclusion.md`)*
- [ ] **Lecture 5 — Counting (continued): the Pigeonhole Principle** `[DB]` `[CRY]` — why hash
  collisions are guaranteed rather than unlucky. *(`05__lec__pigeonhole-principle.md`)*
- [ ] **Exercises 3 — Functions and Counting** — practice test.
  *(`03__ex__functions-and-counting.md`)*

## Section 4: Asymptotics `[CMP]` `[DB]` `[NET]`
Big-O with proofs — the shared vocabulary of every index choice, routing decision and
optimization pass. Deeper here than in Rosen Ch. 3.

- [ ] **Lecture 6 — Asymptotics** `[CMP]` — big-O defined and applied. *(`06__lec__asymptotics.md`)*
- [ ] **Lecture 7A — Asymptotics of Sequences (Part A)** `[CMP]` — domination, growth
  comparison. *(`07a__lec__asymptotics-en-a.md`)*
- [ ] **Lecture 7B — Asymptotics (Part B)** `[CMP]` — from implicit form to explicit estimates.
  *(`07b__lec__asymptotics-en-b.md`)*
- [ ] **Exercises 4 — Counting and Asymptotics** — practice test.
  *(`04__ex__counting-and-asymptotics.md`)*

## Section 5: Recurrences `[CMP]` `[DB]`
The heart of algorithm analysis: cost equations you solve in closed form. This course goes
considerably further than Rosen 7.1–7.2. *(Doubled — see the note above.)*

- [ ] **Lectures 8–9 — Recurrences (combined)** `[CMP]` — sequential search, difference
  equations, existence and uniqueness, the space of sequences, inhomogeneous schemes.
  *(`08-09__lec__recurrences.md`)*
- [ ] **Lecture 8 — Recurrences (EN)** `[CMP]` — Tower of Hanoi, difference equations, the
  linear isomorphism φ, order-2 solutions. *(`08__lec__recurrences-en.md`)*
- [ ] **Lecture 9 — Recurrences, continued (EN)** `[CMP]` — order-k homogeneous schemes,
  inhomogeneous schemes, the four-step method. *(`09__lec__recurrences-en.md`)*
- [ ] **Exercises 5 — Recurrences and Asymptotics** — practice test.
  *(`05__ex__recurrences-and-asymptotics.md`)*

## Section 6: Generating Functions and Divide & Conquer `[CMP]` `[DB]`
Generating functions turn a recurrence into algebra; divide-and-conquer recurrences give the cost
of sorting, searching and B-tree operations. *(Doubled — see the note above.)*

- [ ] **Lectures 10–11 — Generating Functions and Divide-and-Conquer (combined)** `[CMP]` —
  variable-coefficient recurrences through to the D&C paradigm.
  *(`10-11__lec__generating-functions-divide-and-conquer.md`)*
- [ ] **Lecture 10 — Recurrences and Generating Functions** `[CMP]` — the generating-function
  method proper. *(`10__lec__recurrences-and-functions.md`)*
- [ ] **Lecture 11 — Divide and Conquer** `[CMP]` `[DB]` — splitting a problem and paying for
  the merge. *(`11__lec__divide-and-conquer.md`)*

## Section 7: Graphs `[NET]` `[CMP]` `[DB]`
Networks are graphs. Pairs with Rosen Ch. 9.

- [ ] **Lecture 12 — Graphs and Trees: Introduction** `[NET]` — digraphs, undirected graphs,
  the basic vocabulary. *(`12__lec__graphs.md`)*
- [ ] **Lecture 13a — Euler Graphs** `[NET]` — the Königsberg bridges and Euler's 1736
  answer: traversing every edge. *(`13a__lec__euler-graphs.md`)*
- [ ] **Lecture 13b — Hamiltonian Graphs and Trees** `[NET]` — visiting every vertex; rooted
  trees. *(`13b__lec__hamiltonian-graphs-trees.md`)*
- [ ] **Exercises 6 — Recurrences and Graphs** — practice test.
  *(`06__ex__recurrences-and-graphs.md`)*

## Section 8: Trees `[DB]` `[CMP]`
The structure databases are physically made of and compilers parse into. Pairs with Rosen Ch. 10.

- [ ] **Lecture 14 — Trees: Introduction** `[DB]` — trees as connected acyclic graphs;
  isomorphism. *(`14a__lec__trees-intro.md`)*
- [ ] **Lecture 14b — Binary Trees and Closing Curiosities** `[CMP]` `[DB]` — the recursive
  definition of binary trees: the shape of every AST and every index node.
  *(`14b__lec__binary-trees.md`)*

## Section 9: Exam Review
The course's own consolidation pass — treat it as a mock exam, closed-book, before ticking.

- [ ] **Lecture 15 — Review for the Exams** — worked problems across asymptotics, recurrences,
  counting and graphs. *(`15__lec__review.md`)*

---

## How this course and Rosen overlap

Roughly **half of Rosen has a lecture counterpart, and the whole lecture course lives inside
Rosen's scope.** Nothing here is outside the book — but large, goal-critical parts of the book
are outside the course. Read this table before deciding what to study twice and what you can
only get from one source.

### Shared ground — where each section has a Rosen counterpart

| Lecture course | Rosen sections | Which is deeper |
|---|---|---|
| §1 Sets and Functions | 2.1, 2.2, 2.3 | About equal |
| §2 Mathematical Induction | 5.1, 5.2 | About equal (Rosen adds structural induction — see below) |
| §3 Counting and Combinatorics | 6.1–6.3, 6.5, 7.5–7.6 | Rosen — more techniques, more exercises |
| §4 Asymptotics | 3.2, 3.3 | **Lectures** — three lectures with proofs vs. two sections |
| §5 Recurrences | 7.1, 7.2 | **Lectures**, substantially — difference equations as a vector space, the linear isomorphism, order-*k* schemes, a systematic four-step method |
| §6 Generating Functions and D&C | 7.3, 7.4 | **Lectures** — variable-coefficient recurrences get real treatment |
| §7 Graphs | 9.1, 9.2, 9.4, 9.5 | Comparable; the lectures give Euler graphs a full historical treatment |
| §8 Trees | 10.1, 10.2, 10.3 | Rosen — B-trees and Huffman coding are in 10.2 and have no lecture |
| §9 Exam Review | — | Lectures only |

### Rosen-only — no lecture covers these, and several are load-bearing for your goals

| Rosen | Topic | Why it matters |
|---|---|---|
| **Ch. 1** (all) | Logic, quantifiers, proof methods, normal forms, program correctness | `[DB]` predicate logic *is* SQL; the proof skill everything else assumes |
| 2.4–2.7 | Recursive functions, summations, cardinality, matrices | Cardinality is the route to "some problems are uncomputable" |
| 3.1 | Algorithms as objects | Vocabulary for the rest of the book |
| **Ch. 4** (all) | Number theory, modular arithmetic, RSA | `[CRY]` **your entire cryptography foundation** — the lectures contain none of it |
| 5.3–5.4 | Structural induction, recursive algorithms | `[CMP]` how you reason about grammars and syntax trees |
| 6.4, 6.6 | Binomial coefficients, generating permutations | Feeds probability (module 06) |
| **Ch. 8** (all) | Relations, closures, equivalences, partial orders | `[DB]` the mathematical definition of a table; `[CMP]` dataflow lattices |
| 9.3, 9.6–9.8 | Graph representation, shortest paths, planarity, coloring | `[NET]` Dijkstra is routing; `[CMP]` graph coloring is register allocation |
| 10.4–10.5 | Spanning trees, minimum spanning trees | `[NET]` Spanning Tree Protocol, network design |
| **Ch. 11** | Boolean algebra and circuit minimization | `[CMP]` `[DB]` condition simplification |
| **Ch. 12** | Groups, rings, fields, coding theory | `[CRY]` the language modern crypto is written in; `[NET]` CRC |
| **Ch. 13** | Grammars, finite-state machines, Turing machines | `[CMP]` **the compiler chapter** — parsing, lexers, the limits of computation |

### How to run them together

- **On shared topics, lecture first, Rosen second.** The lectures have worked examples and
  graded exercise sessions with full solutions; Rosen has the exercise volume and the precise
  statements. Two passes over the same idea from two angles is the cheapest real learning
  available in this whole curriculum — take it where it is offered.
- **Recurrences, asymptotics and generating functions: trust the lectures.** They go past the
  book. Read Rosen 7.1–7.4 afterwards as a summary, not as the main text.
- **Everything in the Rosen-only table has to come from Rosen.** Four of your five programming
  goals depend on chapters the lecture course never touches — cryptography most starkly, since
  Ch. 4 and Ch. 12 are its entire mathematical basis and the lectures offer no substitute.
- **Do not wait to finish one before starting the other.** The bars fill per section, so
  running both in parallel is legible: ticking Rosen Ch. 9 and lecture §7 in the same week
  lights up the corresponding segment on each of the two bars.

## Exit criteria

You are done with this module when you can, cold:

- Prove a summation formula by induction, writing base case and inductive step properly.
- Count a structured set three ways — product rule, inclusion–exclusion, and pigeonhole for the
  bound — and say which is tightest.
- Prove f(n) = O(g(n)) from the definition, with explicit constants.
- Solve a second-order linear recurrence, homogeneous and inhomogeneous, by hand.
- Turn a recurrence into a generating function and read the closed form back out.
- Decide whether a graph has an Euler circuit and justify it by Euler's degree condition.
- Sit Lecture 15's review problems closed-book and score above 80%.
