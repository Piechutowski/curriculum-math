# Module 02 — Study Map: Rosen & Krithivasan, *Discrete Mathematics and Its Applications*

**Resource:** the book you own.
**How to use this file:** do the chapters in order for the full pass (each chapter leans on the
previous ones), keeping the rhythm — for this book that means: read the section, write study
notes, then work the section exercises, and at chapter end do the *Review Questions* and a slice
of *Supplementary Exercises*. **Do the Computer Projects** at the end of each chapter — they are
the bridge between this curriculum and your programming goals, and no other book you own has
them.

Tags: `[GFX]` graphics/simulation · `[CMP]` compilers · `[NET]` networking · `[DB]` databases ·
`[CRY]` cryptography.

---

## Ch. 1 — The Foundations: Logic and Proofs `[DB]` `[CMP]`
The chapter that trains the muscle every other chapter uses. Direct payoffs: predicates and
quantifiers (1.4–1.5) are the mathematics of SQL — a `WHERE` clause is a predicate, a nested
query is a nested quantifier `[DB]`. Normal forms (1.9) return in query optimization and circuit
minimization. Program correctness (1.10) is your first contact with reasoning about programs
`[CMP]`. Work 1.7–1.8 (proofs) slowly — proof skill is what makes the rest of the book readable.

## Ch. 2 — Sets, Functions, Sequences, Sums, Matrices `[DB]` `[CMP]`
Sets and set operations (2.1–2.2) are the raw material of the relational model — a table is a
set of tuples `[DB]`. Functions (2.3) formalize every mapping you will ever program. Cardinality
(2.6) proves there are more problems than programs — the cleanest route to "some things are
uncomputable" `[CMP]`. 2.4 (primitive/partial recursive functions) is computability foundations;
it pairs with Ch. 13.

## Ch. 3 — Algorithms `[CMP]` `[DB]` `[NET]`
Big-O and complexity (3.2–3.3) are the shared vocabulary of every index choice `[DB]`, routing
decision `[NET]`, and optimization pass `[CMP]`. Short chapter, permanent vocabulary.

## Ch. 4 — Number Theory and Cryptography `[CRY]` **core** · also `[DB]` `[NET]`
The reason this book is on your shelf, for crypto: divisibility and modular arithmetic (4.1),
primes and GCD / Euclidean algorithm (4.3), solving congruences and the Chinese Remainder
Theorem (4.4), and then 4.6 builds actual RSA. Do this chapter completely and slowly. Bonus
payoffs hiding in 4.5: hash functions and pseudorandom generators `[DB]`, check digits — the
ancestor of checksums `[NET]`.

## Ch. 5 — Induction and Recursion `[CMP]` **core**
The single most important chapter for compiler work. Structural induction (5.3) is how you
reason about grammars, syntax trees, and program transformations — "prove it for the leaves,
prove it for each way of building a bigger tree." Recursive algorithms with correctness proofs
(5.4) is the habit of mind compilers are written in.

## Ch. 6 — Counting `[CRY]` `[NET]` `[DB]`
Keyspace sizes and brute-force costs `[CRY]` are counting problems. The pigeonhole principle
(6.2) proves hash collisions are *guaranteed*, not unlucky `[DB]` `[CRY]`. This chapter feeds
the Probability module directly — do it before starting module 05.

## Ch. 7 — Advanced Counting Techniques `[CMP]` `[DB]`
Recurrence relations and their solutions (7.1–7.3) are algorithm analysis: divide-and-conquer
recurrences give the cost of sorting, searching, and B-tree operations `[DB]`. Generating
functions (7.4) are optional-depth on first pass — read, don't master. Inclusion–exclusion
(7.5–7.6) returns in query cardinality estimation `[DB]`.

## Ch. 8 — Relations `[DB]` **core** · `[CMP]` `[NET]`
The mathematical foundation of the relational model, nearly by name: an n-ary relation (8.2)
*is* a database table, and the section says so. Closures (8.4): transitive closure is
reachability — recursive queries `[DB]` and "can A reach B" in networks `[NET]`. Equivalence
relations (8.5) partition the world (think: hash buckets, type equivalence). Partial orders
(8.6) are dependency ordering — build systems, topological sort, and the lattices that dataflow
analysis in compilers runs on `[CMP]`.

## Ch. 9 — Graphs `[NET]` **core** · `[CMP]` `[DB]`
Networks are graphs; this is the networking chapter in disguise. Shortest paths (9.6) is
routing — Dijkstra's algorithm is what link-state protocols compute `[NET]`. Connectivity (9.4)
is network robustness. Graph coloring (9.8) is register allocation `[CMP]` — one of the most
beautiful theory-to-practice landings in computing. Control-flow graphs `[CMP]` and query graphs
`[DB]` are both this chapter's objects.

## Ch. 10 — Trees `[DB]` **core** · `[CMP]` `[NET]`
Applications of trees (10.2) contains **B-trees** — the data structure relational databases are
physically made of `[DB]` — and **Huffman coding**, the entry point to compression `[NET]`. Tree
traversal (10.3) is expression evaluation and code generation: postorder traversal of an AST is
stack-machine code `[CMP]`. Spanning trees (10.4–10.5): Ethernet's Spanning Tree Protocol and
minimum-cost network design `[NET]`.

## Ch. 11 — Boolean Algebra `[CMP]` `[DB]`
Minimization (11.4) is simplifying boolean conditions — done by compilers and query planners
alike. Logic gates ground you for how hardware executes everything above.

## Ch. 12 — Algebraic Structures and Coding Theory `[CRY]` **core** · `[NET]`
Groups, rings, and fields (12.3–12.5) are the language modern cryptography is written in —
"the multiplicative group mod p," "the field GF(2⁸)" (AES lives there), and someday elliptic
curves. Then 12.7–12.8 (coding theory, polynomial codes) is error detection and correction:
CRC — the checksum in Ethernet frames — is a polynomial code `[NET]`. This chapter is the
bridge between your crypto goal and your networking goal.

## Ch. 13 — Modeling Computation `[CMP]` **core** · `[NET]`
The compiler chapter. Languages and grammars (13.1) is parsing theory — BNF grammars are what
you will implement. Finite-state machines (13.2–13.3) are lexers/tokenizers `[CMP]` *and*
protocol state machines — TCP's connection lifecycle is a textbook FSM `[NET]`. Language
recognition (13.4) tells you what regular expressions can and cannot match, and why parsers need
more power. Turing machines (13.5) mark the limits of what any compiler can ever do.

---

## Per-goal sprint paths

Do chapters in order for the full pass. When you want to sprint toward one goal:

| Goal | Path through Rosen |
|---|---|
| Compilers `[CMP]` | 1 → 2 → 5 → 3 → 9 → 10 → 13 |
| Databases `[DB]` | 1 → 2 → 3 → 8 → 10 → 11 → 7 |
| Cryptography `[CRY]` | 1 → 4 → 6 → 12 (then Probability §§2–3) |
| Networking `[NET]` | 1 → 9 → 10 → 6 → 12.7–12.8 (then Probability §§4–8) |
