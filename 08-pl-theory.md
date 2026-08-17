# Module 08 — Mathematics of Programming Languages

<!-- progress:08-pl-theory -->
![08 · PL Theory — Sipser · Pierce · Winskel progress](progress/08-pl-theory.svg?v=4-0-50)

**0 / 50 lessons complete · 0%** — 0 of 9 sections finished · [completion log](LOG.md)
<!-- endprogress -->

**Books:** three spines, used by section. Michael Sipser, *Introduction to the Theory of
Computation* (Cengage India) — §§2–4. Benjamin Pierce, *Types and Programming Languages* ("TAPL",
MIT Press) — §§5–6, 8. Glynn Winskel, *The Formal Semantics of Programming Languages* (MIT Press)
— §§6–7; compact free alternative: Nielson & Nielson, *Semantics with Applications*. Free
alternative for §§5–8 with machine-checked proofs: *Software Foundations*, Vol. 1–2 (online).
**Prerequisites:** Rosen Ch. 1, 2 and 5 (logic, sets, induction) — hard requirements. Rosen 8.6
(partial orders) and Ch. 9 (graphs) before §7. Rosen Ch. 13 is the four-section trailer for
§§2–4 of this module — helpful first, not required. Fully independent of Calculus, Linear
Algebra and Probability: this is the discrete track's second half.
**Why this module exists:** compiler engineering is the one goal whose core mathematics appears
nowhere else in the curriculum beyond Rosen Ch. 13's trailer. This module is that mathematics in
full: automata (lexers), grammars (parsers), computability (what no compiler can ever do), the
lambda calculus and type theory (what type checkers prove), semantics (what "correct
compilation" even means), and lattice fixpoints (what optimizers compute). The operator
cheat-sheet for the whole module is
[notes/Big operators in programming and PL theory.md](notes/Big%20operators%20in%20programming%20and%20PL%20theory.md).

**The rhythm applies to every lesson:** Lesson → Study Notes → Practice Test. Practice here
means proofs and constructions by hand — build the automaton, reduce the term, prove the case.
Where a lesson has a natural 10–30 line program (a DFA simulator, a CYK table, a unifier),
module 07's rule applies: write it, as part of the practice test.

Tags: `[GFX]` graphics/simulation · `[CMP]` compilers · `[NET]` networking · `[DB]` databases ·
`[CRY]` cryptography. No tag = foundation later lessons stand on.

---

## Section 1: Words, Trees, and Induction — the Bridge

Everything in this module is defined inductively and proved by induction. This section makes
that muscle explicit before the machines arrive.

- [ ] **1. Alphabets, strings, and languages** `[CMP]` — Σ*, the free monoid; union,
   concatenation, and star on languages, and their identities (∅ vs {ε} — not the same thing).
   *(Sipser Ch. 0)*
- [ ] **2. Structural induction** `[CMP]` — induction on strings, trees, and derivations; Rosen
   5.3 promoted to this module's standard proof method. *(Sipser Ch. 0)*
- [ ] **3. Well-founded induction and termination** — induction over any order with no infinite
   descent; how termination arguments are actually phrased. *(Winskel Ch. 3)*

## Section 2: Regular Languages — the Lexer's Mathematics `[CMP]` `[NET]` `[DB]`

The smallest infinite language class, and the one that runs everywhere: lexers `[CMP]`, protocol
state machines `[NET]`, LIKE and regex matching `[DB]`. Rosen 13.3 was the trailer; this is the
feature.

- [ ] **4. Deterministic finite automata** `[CMP]` `[NET]` — the five-tuple; running a DFA is a
   fold of δ over the input; TCP's connection lifecycle as a worked example. *(Sipser 1.1)*
- [ ] **5. Nondeterministic finite automata; the subset construction** `[CMP]` — nondeterminism
   as tracking a *set* of states; the construction inside every regex compiler. *(Sipser 1.2)*
- [ ] **6. Kleene's theorem: regular expressions ↔ automata** `[CMP]` `[DB]` — two notations, one
   class; both directions of the translation, by hand. *(Sipser 1.3)*
- [ ] **7. Closure properties; the product construction** — union, intersection, complement; why
   complement demands determinism first. *(Sipser 1.2–1.3)*
- [ ] **8. The pumping lemma for regular languages** `[CMP]` — finite memory cannot count:
   {aⁿbⁿ} is not regular, which is exactly why no regex can match balanced parentheses.
   *(Sipser 1.4)*
- [ ] **9. Myhill–Nerode and DFA minimization** `[CMP]` — the coarsest equivalence on states; the
   canonical minimal DFA and the algorithm that finds it. *(Sipser Problem 1.52;
   Hopcroft–Ullman)*
- [ ] **10. The mathematics of a lexer** `[CMP]` `[NET]` — maximal munch and rule priority; why
    scanning is linear-time, and where backtracking regex engines go exponential instead
    (ReDoS).

## Section 3: Context-Free Languages — the Parser's Mathematics `[CMP]`

One level up the hierarchy: grammars that count. The mathematics under every parser generator
ever shipped.

- [ ] **11. Context-free grammars and derivations** `[CMP]` — productions, derivations, parse
    trees; leftmost and rightmost derivations. *(Sipser 2.1)*
- [ ] **12. Ambiguity** `[CMP]` — two parse trees for one string; how precedence and
    associativity are *encoded in the grammar*; the dangling else. *(Sipser 2.1)*
- [ ] **13. Chomsky normal form; the CYK algorithm** — normalizing a grammar; parsing as dynamic
    programming in O(n³). *(Sipser 2.1)*
- [ ] **14. Pushdown automata** `[CMP]` — finite control plus a stack; equivalence with CFGs; why
    "recursive descent" and "a stack" are the same sentence. *(Sipser 2.2)*
- [ ] **15. The pumping lemma for CFLs** `[CMP]` — what one stack cannot do: {aⁿbⁿcⁿ}; why C's
    typedef context-sensitivity haunts parsers. *(Sipser 2.3)*
- [ ] **16. Deterministic CFLs; FIRST/FOLLOW as fixpoint equations** `[CMP]` — the LL/LR
    boundary; computing FIRST and FOLLOW sets is Kleene iteration (§7) in disguise.
    *(Sipser 2.4; any compilers text)*
- [ ] **17. The Chomsky hierarchy** — regular ⊂ context-free ⊂ context-sensitive ⊂ recursively
    enumerable; which machine buys which level, and where real languages actually sit.
    *(Sipser survey)*

## Section 4: Computability — What No Compiler Can Do `[CMP]`

The permanent limits. Every static-analysis design decision traces back to this section.

- [ ] **18. Turing machines; the Church–Turing thesis** — the honest definition of "algorithm";
    Rosen 13.5 done properly. *(Sipser 3.1–3.3)*
- [ ] **19. Decidable vs. recognizable; the universal machine** `[CMP]` — deciders halt,
    recognizers may loop; a TM that runs TMs is the first interpreter. *(Sipser 4.1)*
- [ ] **20. Diagonalization and the halting problem** `[CMP]` — the proof, cold; the program that
    asks "do I halt?". *(Sipser 4.2)*
- [ ] **21. Reductions** — proving new problems undecidable from old ones; the workhorse move of
    the field. *(Sipser 5.1, 5.3)*
- [ ] **22. Rice's theorem** `[CMP]` — every non-trivial question about program *behavior* is
    undecidable: the theorem that forces optimizers and analyzers to approximate (§7 is the
    mathematics of approximating soundly). *(Sipser Problem 5.28)*
- [ ] **23. Orientation: P, NP, and NP-completeness** `[CMP]` `[DB]` — the coarse map of the
    feasible; register allocation is graph coloring (Rosen 9.8), SAT is the first hard problem,
    query optimization is full of NP. One lesson, no more. *(Sipser Ch. 7)*

## Section 5: The Untyped Lambda Calculus `[CMP]`

Three constructs — variable, abstraction, application — and all of computation. The core
calculus of functional languages, and the substrate type theory lives on.

- [ ] **24. Syntax, α-equivalence, and capture-avoiding substitution** `[CMP]` — the most
    deceptively hard definition in PL; why naive substitution is wrong and every compiler has a
    renaming pass. *(TAPL 5.1, 5.3)*
- [ ] **25. β-reduction; normal forms; evaluation strategies** `[CMP]` — call-by-value,
    call-by-name, normal order: same term, different costs, different termination. *(TAPL 5.1)*
- [ ] **26. Church encodings** — booleans, numerals, pairs from functions alone; computation
    bootstrapped from nothing. *(TAPL 5.2)*
- [ ] **27. Fixed-point combinators** `[CMP]` — Y and Z; recursion manufactured rather than
    assumed; the λ-calculus twin of §7's fixpoints. *(TAPL 5.2)*
- [ ] **28. Confluence (Church–Rosser)** — reduction order cannot change the answer, only
    whether you reach it; uniqueness of normal forms. *(statement + proof sketch; Selinger's
    free lecture notes)*
- [ ] **29. λ-calculus ≡ Turing machines** — the two models compute the same functions; why
    "functional vs. imperative" is a question of taste, never of power. *(Sipser 3 + TAPL 5)*

## Section 6: Semantics — What Programs Mean `[CMP]`

Rules on paper that pin down exactly what an interpreter must do. "Correct compilation" is a
theorem only after this section defines the words in it.

- [ ] **30. Small-step operational semantics** `[CMP]` — inference rules, derivation trees,
    induction on derivations; evaluation as a relation. *(TAPL 3)*
- [ ] **31. Big-step semantics; equivalence of the two styles** `[CMP]` — e ⇓ v; proving the two
    presentations agree. *(TAPL 3; Winskel Ch. 2)*
- [ ] **32. Denotational semantics** — orientation: programs as mathematical functions,
    compositionally; why loops need §7's domains before this is rigorous. *(Winskel Ch. 5)*
- [ ] **33. Axiomatic semantics: Hoare logic** `[CMP]` — {P} c {Q}, weakest preconditions, loop
    invariants; Rosen 1.10 grown into a working proof system. *(Winskel Ch. 6–7)*

## Section 7: Lattices and Fixpoints — the Optimizer's Mathematics `[CMP]` `[DB]`

The mathematics static analysis actually runs on. Rice's theorem (lesson 22) says exact answers
are impossible; this section is how to be soundly, usefully approximate — and it powers
recursive queries `[DB]` for free.

- [ ] **34. Partial orders and lattices, working fluency** — Rosen 8.6 upgraded: join, meet,
    complete lattices, ⊥ and ⊤, and the empty join being ⊥. *(Winskel Ch. 5;
    Nielson–Nielson–Hankin App. A)*
- [ ] **35. Monotone functions; the Knaster–Tarski theorem** — every monotone map on a complete
    lattice has a least fixpoint: the theorem that makes recursive definitions legal at all.
    *(Winskel Ch. 5)*
- [ ] **36. Kleene iteration; the ascending chain condition** `[CMP]` — lfp as the join of ⊥,
    f(⊥), f²(⊥), …; when the worklist algorithm terminates. *(Winskel Ch. 5)*
- [ ] **37. Dataflow analysis as a lattice fixpoint** `[CMP]` — live variables and reaching
    definitions on a control-flow graph; may vs. must, join vs. meet. *(any compilers text;
    Nielson–Nielson–Hankin Ch. 2)*
- [ ] **38. Abstract interpretation** `[CMP]` — orientation: Galois connections, soundness, and
    widening — approximation with a proof that it never lies. *(Nielson–Nielson–Hankin Ch. 4)*
- [ ] **39. Fixpoints in the database** `[DB]` — Datalog, recursive CTEs, and transitive closure
    (Rosen 8.4) as least fixpoints; why recursive SQL is guaranteed to terminate.
- [ ] **40. CPOs, domains, and Scott continuity** — orientation: the mathematics that makes
    denotational recursion (lesson 32) rigorous; fix(f) as the join of the finite unfoldings.
    *(Winskel Ch. 8)*

## Section 8: Type Systems `[CMP]`

The mathematics of "well-typed programs don't go wrong" — the front end's crown jewel.

- [ ] **41. Typing relations; type safety = progress + preservation** `[CMP]` — the judgment
    Γ ⊢ e : τ as an inductively defined relation; the safety recipe stated. *(TAPL 8)*
- [ ] **42. The simply typed lambda calculus, with proofs** `[CMP]` — STLC's rules; prove
    progress and preservation properly, once, in full. *(TAPL 9)*
- [ ] **43. Curry–Howard** `[CMP]` — propositions as types, proofs as programs; why STLC
    normalizes, and what typed languages give up for it (no Y). *(TAPL 9)*
- [ ] **44. The algebra of data types** `[CMP]` — products, sums, unit, void; counting
    inhabitants; Option as 1 + T. *(TAPL 11)*
- [ ] **45. Recursive types** `[CMP]` — μ-types, iso- vs. equi-recursive; List a = μX. 1 + a×X.
    *(TAPL 20)*
- [ ] **46. Subtyping** `[CMP]` — the subtype order; co- and contravariance and why function
    arguments flip; top and bottom types. *(TAPL 15)*
- [ ] **47. Parametric polymorphism: System F** `[CMP]` — ∀ types; generics as Π; parametricity
    orientation ("theorems for free"). *(TAPL 23)*
- [ ] **48. Type inference: unification and Hindley–Milner** `[CMP]` — solving type equations;
    most general unifiers; let-polymorphism: the algorithm inside ML, Haskell, and Rust.
    *(TAPL 22)*

## Section 9: Capstone

- [ ] **49. Orientation: what lies beyond** — dependent types (Π/Σ for real), linear and effect
    types, category theory as the unifying grammar; what each would buy and when to come back
    for it.
- [ ] **50. Exit project: a tiny language, complete, on paper** `[CMP]` — grammar, lexer DFA,
    small-step semantics, and typing rules for a language with booleans, naturals, and
    functions; prove progress + preservation for two cases; run live-variables to fixpoint on a
    three-block program in it. Every section of this module, on one sheet of paper.

---

## Exit criteria

You are done with this module when you can, cold:

- Take a regular expression to an NFA, subset-construct the DFA, minimize it — and prove a given
  language non-regular with the pumping lemma.
- Explain the halting problem's diagonal argument to another programmer in five minutes, and
  state Rice's theorem with two concrete consequences for static analyzers.
- β-reduce a term under call-by-value and under normal order, and derive plus on Church
  numerals.
- Write the small-step rules and typing rules for a boolean/arithmetic language, and prove
  progress and preservation for its if-rule.
- State Knaster–Tarski, and run live-variables analysis to its fixpoint on a four-block CFG.
- Infer the type of λf. λx. f (f x) by unification, showing the constraints and their solution.
