# Module 02 — Rosen & Krithivasan, *Discrete Mathematics* · Checklist & Map

<!-- progress:02-discrete-mathematics -->
![02 · Discrete Mathematics — Rosen progress](progress/02-discrete-mathematics.svg?v=2-0-91)

**0 / 91 lessons complete · 0%** — 0 of 13 chapters finished · [completion log](LOG.md)
<!-- endprogress -->

**Resource:** the book you own.
**How to tick:** one checkbox per book section. Tick only when the rhythm is done — section
read, study notes written, and the section's exercise set worked. Each chapter also gets one
**wrap-up** box: Review Questions + a slice of Supplementary Exercises + **one Computer
Project** (they are the bridge to your programming goals; no other book you own has them).
After editing checkboxes, run `go run ./tools/progress` to refresh the bars.

Do the chapters in order for the full pass; per-goal sprint paths are at the bottom.

Tags: `[GFX]` graphics/simulation · `[CMP]` compilers · `[NET]` networking · `[DB]` databases ·
`[CRY]` cryptography.

---

## Ch. 1 — The Foundations: Logic and Proofs `[DB]` `[CMP]`
Predicates and quantifiers (1.4–1.5) are the mathematics of SQL — a `WHERE` clause is a
predicate, a nested query is a nested quantifier. Normal forms (1.9) return in query
optimization and circuit minimization; program correctness (1.10) is reasoning about programs.
Work 1.7–1.8 slowly — proof skill makes the rest of the book readable.

- [ ] 1.1 Propositional Logic
- [ ] 1.2 Applications of Propositional Logic
- [ ] 1.3 Propositional Equivalences
- [ ] 1.4 Predicates and Quantifiers
- [ ] 1.5 Nested Quantifiers
- [ ] 1.6 Rules of Inference
- [ ] 1.7 Introduction to Proofs
- [ ] 1.8 Proof Methods and Strategy
- [ ] 1.9 Normal Forms
- [ ] 1.10 Program Correctness
- [ ] Ch. 1 wrap-up: Review Questions · Supplementary Exercises · one Computer Project

## Ch. 2 — Basic Structures: Sets, Functions, Sequences, Sums, Matrices `[DB]` `[CMP]`
Sets (2.1–2.2) are the raw material of the relational model — a table is a set of tuples.
Cardinality (2.6) proves there are more problems than programs — the cleanest route to "some
things are uncomputable."

- [ ] 2.1 Sets
- [ ] 2.2 Set Operations
- [ ] 2.3 Functions
- [ ] 2.4 Primitive and Partial Recursive Functions
- [ ] 2.5 Sequences and Summations
- [ ] 2.6 Cardinality of Sets
- [ ] 2.7 Matrices
- [ ] Ch. 2 wrap-up: Review Questions · Supplementary Exercises · one Computer Project

## Ch. 3 — Algorithms `[CMP]` `[DB]` `[NET]`
Big-O and complexity are the shared vocabulary of every index choice, routing decision, and
optimization pass. Short chapter, permanent vocabulary.

- [ ] 3.1 Algorithms
- [ ] 3.2 The Growth of Functions
- [ ] 3.3 Complexity of Algorithms
- [ ] Ch. 3 wrap-up: Review Questions · Supplementary Exercises · one Computer Project

## Ch. 4 — Number Theory and Cryptography `[CRY]` **core** · `[DB]` `[NET]`
The reason this book is on your shelf, for crypto — through to actual RSA in 4.6. Bonus payoffs
in 4.5: hash functions and pseudorandom generators `[DB]`, check digits `[NET]`. Do this chapter
completely and slowly.

- [ ] 4.1 Divisibility and Modular Arithmetic
- [ ] 4.2 Integer Representations and Algorithms
- [ ] 4.3 Primes and Greatest Common Divisors
- [ ] 4.4 Solving Congruences
- [ ] 4.5 Applications of Congruences
- [ ] 4.6 Cryptography
- [ ] Ch. 4 wrap-up: Review Questions · Supplementary Exercises · one Computer Project

## Ch. 5 — Induction and Recursion `[CMP]` **core**
The most important chapter for compiler work. Structural induction (5.3) is how you reason about
grammars, syntax trees, and program transformations.

- [ ] 5.1 Mathematical Induction
- [ ] 5.2 Strong Induction and Well-Ordering
- [ ] 5.3 Recursive Definitions and Structural Induction
- [ ] 5.4 Recursive Algorithms
- [ ] Ch. 5 wrap-up: Review Questions · Supplementary Exercises · one Computer Project

## Ch. 6 — Counting `[CRY]` `[NET]` `[DB]`
Keyspace sizes and brute-force costs are counting problems; the pigeonhole principle (6.2)
proves hash collisions are guaranteed, not unlucky. Feeds the Probability module — do this
chapter before starting module 05.

- [ ] 6.1 The Basics of Counting
- [ ] 6.2 The Pigeonhole Principle
- [ ] 6.3 Permutations and Combinations
- [ ] 6.4 Binomial Coefficients and Identities
- [ ] 6.5 Generalized Permutations and Combinations
- [ ] 6.6 Generating Permutations and Combinations
- [ ] Ch. 6 wrap-up: Review Questions · Supplementary Exercises · one Computer Project

## Ch. 7 — Advanced Counting Techniques `[CMP]` `[DB]`
Recurrences (7.1–7.3) are algorithm analysis: divide-and-conquer recurrences give the cost of
sorting, searching, and B-tree operations. Generating functions (7.4): read, don't master, on
first pass. Inclusion–exclusion returns in query cardinality estimation.

- [ ] 7.1 Applications of Recurrence Relations
- [ ] 7.2 Solving Linear Recurrence Relations
- [ ] 7.3 Divide-and-Conquer Algorithms and Recurrence Relations
- [ ] 7.4 Generating Functions
- [ ] 7.5 Inclusion–Exclusion
- [ ] 7.6 Applications of Inclusion–Exclusion
- [ ] Ch. 7 wrap-up: Review Questions · Supplementary Exercises · one Computer Project

## Ch. 8 — Relations `[DB]` **core** · `[CMP]` `[NET]`
The mathematical foundation of the relational model, nearly by name: an n-ary relation (8.2)
*is* a database table. Transitive closure (8.4) is reachability — recursive queries and "can A
reach B." Partial orders (8.6) are dependency ordering and the lattices dataflow analysis runs
on.

- [ ] 8.1 Relations and Their Properties
- [ ] 8.2 n-ary Relations and Their Applications
- [ ] 8.3 Representing Relations
- [ ] 8.4 Closures of Relations
- [ ] 8.5 Equivalence Relations
- [ ] 8.6 Partial Orderings
- [ ] Ch. 8 wrap-up: Review Questions · Supplementary Exercises · one Computer Project

## Ch. 9 — Graphs `[NET]` **core** · `[CMP]` `[DB]`
Networks are graphs. Shortest paths (9.6) is routing — Dijkstra is what link-state protocols
compute. Graph coloring (9.8) is register allocation — one of the most beautiful
theory-to-practice landings in computing.

- [ ] 9.1 Graphs and Graph Models
- [ ] 9.2 Graph Terminology and Special Types of Graphs
- [ ] 9.3 Representing Graphs and Graph Isomorphism
- [ ] 9.4 Connectivity
- [ ] 9.5 Euler and Hamilton Paths
- [ ] 9.6 Shortest-Path Problems
- [ ] 9.7 Planar Graphs
- [ ] 9.8 Graph Coloring
- [ ] Ch. 9 wrap-up: Review Questions · Supplementary Exercises · one Computer Project

## Ch. 10 — Trees `[DB]` **core** · `[CMP]` `[NET]`
10.2 contains **B-trees** — what relational databases are physically made of — and **Huffman
coding** (compression). Postorder traversal of an AST (10.3) is stack-machine code. Spanning
trees (10.4–10.5): Ethernet's Spanning Tree Protocol, minimum-cost network design.

- [ ] 10.1 Introduction to Trees
- [ ] 10.2 Applications of Trees
- [ ] 10.3 Tree Traversal
- [ ] 10.4 Spanning Trees
- [ ] 10.5 Minimum Spanning Trees
- [ ] Ch. 10 wrap-up: Review Questions · Supplementary Exercises · one Computer Project

## Ch. 11 — Boolean Algebra `[CMP]` `[DB]`
Minimization (11.4) is simplifying boolean conditions — done by compilers and query planners
alike; gates ground you for how hardware executes everything above.

- [ ] 11.1 Boolean Functions
- [ ] 11.2 Representing Boolean Functions
- [ ] 11.3 Logic Gates
- [ ] 11.4 Minimization of Circuits
- [ ] Ch. 11 wrap-up: Review Questions · Supplementary Exercises · one Computer Project

## Ch. 12 — Algebraic Structures and Coding Theory `[CRY]` **core** · `[NET]`
Groups, rings, fields (12.3–12.5) are the language modern crypto is written in — AES lives in
GF(2⁸), elliptic curves live here someday. Then 12.7–12.8: error-correcting codes and CRC —
the bridge between your crypto goal and your networking goal.

- [ ] 12.1 Introduction
- [ ] 12.2 The Structure of Algebras
- [ ] 12.3 Semigroups, Monoids and Groups
- [ ] 12.4 Homomorphisms, Normal Subgroups and Congruence Relations
- [ ] 12.5 Rings, Integral Domains and Fields
- [ ] 12.6 Quotient and Product Algebras
- [ ] 12.7 Coding Theory
- [ ] 12.8 Polynomial Rings and Polynomial Codes
- [ ] Ch. 12 wrap-up: Review Questions · Supplementary Exercises · one Computer Project

## Ch. 13 — Modeling Computation `[CMP]` **core** · `[NET]`
The compiler chapter. Grammars (13.1) = parsing; finite-state machines (13.2–13.3) = lexers
*and* protocol state machines (TCP's connection lifecycle is a textbook FSM); 13.4 = what regex
can and cannot match; Turing machines (13.5) = the limits of what any compiler can ever do.

- [ ] 13.1 Languages and Grammars
- [ ] 13.2 Finite-State Machines with Output
- [ ] 13.3 Finite-State Machines with No Output
- [ ] 13.4 Language Recognition
- [ ] 13.5 Turing Machines
- [ ] Ch. 13 wrap-up: Review Questions · Supplementary Exercises · one Computer Project

---

## Per-goal sprint paths

| Goal | Path through Rosen |
|---|---|
| Compilers `[CMP]` | 1 → 2 → 5 → 3 → 9 → 10 → 13 |
| Databases `[DB]` | 1 → 2 → 3 → 8 → 10 → 11 → 7 |
| Cryptography `[CRY]` | 1 → 4 → 6 → 12 (then Probability §§2–3) |
| Networking `[NET]` | 1 → 9 → 10 → 6 → 12.7–12.8 (then Probability §§4–8) |
