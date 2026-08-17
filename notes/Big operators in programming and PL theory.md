---
aliases:
  - Big operators for programmers
  - Monoids in programming
  - PL theory operators
tags:
  - math/notation
  - programming
  - pl-theory
  - reference
---

# Big operators in programming and PL theory

Companion to [[Big math operators]]. That note covers the operators of the mathematics curriculum
(modules 01–07). This one covers the ones that only show up once you are writing compilers, type
checkers, static analysers and distributed systems — where the *monoid itself* stops being a
background assumption and becomes the thing you are designing.

The shift in perspective: in mathematics the identity element is a footnote about empty ranges. In
programming it is **the base case of your recursion, the initial value of your accumulator, the
`⊥` your dataflow analysis starts from, and the reason `all([])` is `True`**. Getting it wrong is
a bug, not a technicality.

---

## 1. Master table

| Big | Binary | Identity | Where |
|---|---|---|---|
| `fold` / catamorphism | any monoid $\odot$ | $e$ | §2 |
| $\prod$ (concatenation) | $\cdot$ | $\varepsilon$ (empty string) | §3 free monoid |
| $\bigcup$ over languages | $\cup$ | $\varnothing$ | §4 |
| $\prod$ over languages | $L_1 L_2$ | $\{\varepsilon\}$ | §4 |
| $L^{*} = \bigcup_{n \ge 0} L^n$ | — | — | §4 Kleene star |
| $\bigsqcup$ (join) | $\sqcup$ | $\bot$ | §5 dataflow, CRDTs |
| $\mathop{\sqcap}$ (meet) | $\sqcap$ | $\top$ | §5 must-analysis |
| $\operatorname{lfp}$, $\mu$ | — | $\bot$ | §5 fixpoints |
| $\operatorname{gfp}$, $\nu$ | — | $\top$ | §5 coinduction |
| $\prod$ (product type) | $\times$ | `unit` / `()` | §6 |
| $\coprod$, $\sum$ (sum type) | $+$ | `void` / `never` | §6 |
| $\Pi_{x:A}$ | $\to$ | — | §6 dependent function |
| $\Sigma_{x:A}$ | $\times$ | — | §6 dependent pair |
| $\bigwedge$ (intersection types) | $\wedge$ | $\top$ (`unknown`) | §6 |
| $\bigvee$ (union types) | $\vee$ | $\bot$ (`never`) | §6 |
| $\bigwedge$ (rule premises) | $\wedge$ | $\top$ (an axiom) | §7 |
| $\bigotimes$, $\bigoplus$ (linear logic) | $\otimes$, $\oplus$ | $\mathbf{1}$, $\mathbf{0}$ | §8 |
| $\varprojlim$ / $\varinjlim$ (limits) | — | terminal / initial object | §9 |

---

## 2. The fold is the point

Every big operator in [[Big math operators]] is a fold. In programming the direction matters, so
there are two, and they only agree when the operation is associative with an identity:

$$\texttt{foldr}\;(\odot)\;e\;[a_1, \dots, a_n] = a_1 \odot (a_2 \odot (\dots \odot (a_n \odot e)))$$
$$\texttt{foldl}\;(\odot)\;e\;[a_1, \dots, a_n] = (((e \odot a_1) \odot a_2) \odot \dots) \odot a_n$$

The `Monoid` abstraction in Haskell, Rust's `Sum`/`Product` wrappers, Go's `slices` reductions and
`reduce` in every language are all the same three-field structure: a type, an associative binary
operation, an identity.

```go
// The general big operator. Every entry in the master tables is one instance of this.
func Fold[T any](xs []T, op func(a, b T) T, identity T) T {
    acc := identity          // the empty-range value — never a made-up zero
    for _, x := range xs {
        acc = op(acc, x)
    }
    return acc
}
```

### What each law buys you at runtime

| Law | Mathematical meaning | What it lets the machine do |
|---|---|---|
| associativity | no parentheses needed | **re-associate**: split the work into chunks, tree-reduce, SIMD, parallel `reduce` |
| commutativity | index order irrelevant | **reorder**: process chunks as they finish, out-of-order, unordered map/reduce |
| identity | empty range defined | **split anywhere**: empty chunks are free; no special case for empty input |
| idempotence ($a \odot a = a$) | duplicates harmless | **retry and re-deliver safely** — the CRDT property (§5) |

MapReduce is exactly this theorem: your reducer must be an associative, commutative monoid, or the
framework's freedom to shard and re-run makes your results nondeterministic. Same for
`std::reduce` vs `std::accumulate` in C++ — the parallel one *requires* associativity, the
sequential one does not.

### Scan — all the partial folds at once

$$\texttt{scan}\;(\odot)\;e\;[a_1,\dots,a_n] = \Big[\, e,\; a_1,\; a_1 \odot a_2,\; \dots,\; \bigodot_{i=1}^{n} a_i \,\Big]$$

Prefix sums, running maxima, and parallel-prefix carry-lookahead adders are all scans. A scan over
an associative operation parallelises in $O(\log n)$ depth — the identity is what makes the first
element of the output exist.

---

## 3. The free monoid — strings

Concatenation over an alphabet $\Sigma$ gives $\Sigma^{*}$, the **free monoid**: associative, has
identity $\varepsilon$ (the empty string), and is emphatically **not** commutative.

$$\prod_{i=1}^{n} s_i = s_1 s_2 \cdots s_n \qquad \prod_{i \in \varnothing} s_i = \varepsilon$$

"Free" means it assumes nothing but the monoid laws — which is why every fold factors through it:
`map` then `fold` is a monoid homomorphism $\Sigma^{*} \to M$. That is the whole content of
"MapReduce works".

> **Notation clash.** $\Sigma$ means *alphabet* in formal language theory and *summation* in
> arithmetic, and both appear on the same page in a compilers text. $\Sigma^{*}$ with a star is
> always the alphabet.

Non-commutativity is not academic: `"ab" + "cd" ≠ "cd" + "ab"`, so a parallel string join must
preserve chunk order even though it may re-associate freely.

---

## 4. Formal languages and automata

`02-discrete-mathematics-map.md` Ch. 13. A language is a set of strings, so all of §6 of
[[Big math operators]] applies — plus two operations of its own.

### Two monoids on languages

$$\bigcup_{i} L_i \quad \text{identity } \varnothing
\qquad\qquad \prod_{i=1}^{n} L_i = L_1 L_2 \cdots L_n \quad \text{identity } \{\varepsilon\}$$

Note the second identity carefully: the concatenation unit is $\{\varepsilon\}$, **the set
containing the empty string**, not the empty set. $\varnothing L = \varnothing$ (an annihilator),
while $\{\varepsilon\} L = L$. Confusing the two is a classic regex-engine bug.

Together they form a **semiring** $(\mathcal{P}(\Sigma^{*}), \cup, \cdot, \varnothing,
\{\varepsilon\})$ — the same structure as §13 of [[Big math operators]], which is why regular
expressions look like arithmetic: alternation is $+$, concatenation is $\times$.

### Kleene star is a big union

$$L^{*} = \bigcup_{n \ge 0} L^{n} \qquad\text{where}\qquad L^{0} = \{\varepsilon\}$$

$L^0 = \{\varepsilon\}$ is the empty concatenation — the identity again. It is why `a*` matches the
empty string, and why $\varnothing^{*} = \{\varepsilon\}$ rather than $\varnothing$.

Kleene star is also the semiring's *closure* operator, the exact analogue of $\sum_{n \ge 0} x^n =
1/(1-x)$: solving $X = LX \cup M$ gives $X = L^{*}M$ (Arden's lemma), which is literally the
geometric series in a different semiring.

### The extended transition function is a fold

$$\hat{\delta}(q, \varepsilon) = q \qquad\qquad \hat{\delta}(q, wa) = \delta\big(\hat{\delta}(q, w),\, a\big)$$

That is `foldl` with the state as accumulator and $\delta$ as the operation. Running a DFA *is*
folding the input string — and the base case $\hat{\delta}(q, \varepsilon) = q$ is the identity
law. Every lexer you write is this fold.

For an NFA the accumulator is a *set* of states and the step is a big union over the frontier:

$$\hat{\delta}(S, a) = \bigcup_{q \in S} \delta(q, a)$$

which is the subset construction (Rosen 13.3) written as one line.

### Product and closure constructions

- Intersection of two DFAs = the product automaton on $Q_1 \times Q_2$ — a Cartesian $\prod$.
- Complement = swap accepting states — needs a *total, deterministic* automaton, which is why
  determinisation comes first.
- Weighted automata swap the semiring: over min-plus you get shortest-path parsing; over the
  probability semiring $([0,1], +, \times)$ you get HMMs; over the Boolean semiring you get plain
  recognition. One algorithm, many meanings.

---

## 5. Lattices, dataflow analysis, and fixpoints

The compiler back end's native mathematics. `02-discrete-mathematics-map.md` 8.6 (partial orders)
is the prerequisite; §7 of [[Big math operators]] is the same lattice algebra.

### Join and meet

$$\bigsqcup_{i \in I} x_i \quad \text{identity } \bot
\qquad\qquad \mathop{\sqcap}_{i \in I} x_i \quad \text{identity } \top$$

A dataflow analysis is a lattice of facts plus a transfer function. The classic equation is a big
join over predecessor blocks:

$$\mathrm{in}[B] = \bigsqcup_{P \in \mathrm{pred}(B)} \mathrm{out}[P]
\qquad\qquad \mathrm{out}[B] = f_B\big(\mathrm{in}[B]\big)$$

The entry block has no predecessors, so $\mathrm{in}[\text{entry}] = \bigsqcup \varnothing = \bot$
— **the empty join, doing real work**. That is not an initialisation convention someone picked; it
is forced by the algebra.

| Analysis | Direction | Combine | Lattice bottom |
|---|---|---|---|
| reaching definitions | forward | $\cup$ (may) | $\varnothing$ |
| available expressions | forward | $\cap$ (must) | all expressions |
| live variables | backward | $\cup$ (may) | $\varnothing$ |
| constant propagation | forward | $\sqcup$ on the flat lattice | $\bot$ = "no information" |

"May" analyses join, "must" analyses meet — and flipping the lattice upside down turns one into the
other, which is why textbooks disagree about which end $\top$ is at. Always state your order.

### Fixpoints — $\mu$, $\nu$, and Kleene iteration

The solution to a dataflow system is the **least fixpoint** of a monotone function on the lattice.
Kleene's theorem computes it as a big join over iterates:

$$\operatorname{lfp}(f) \;=\; \bigsqcup_{n \ge 0} f^{n}(\bot)
\qquad\qquad \operatorname{gfp}(f) \;=\; \mathop{\sqcap}_{n \ge 0} f^{n}(\top)$$

Start at the identity, apply until nothing changes — that *is* the worklist algorithm. It
terminates when the lattice has no infinite ascending chains (the **ascending chain condition**);
when it doesn't, abstract interpretation forces termination with a widening operator $\nabla$ that
deliberately over-jumps to a fixpoint.

$\mu X.\, F(X)$ and $\nu X.\, F(X)$ are the same two fixpoints in the notation used for recursive
types (§6), modal $\mu$-calculus, and model checking: $\mu$ for finite/inductive things
(termination, reachability), $\nu$ for infinite/coinductive ones (safety, streams).

### Scott continuity — "commutes with the big operator"

Denotational semantics needs $f$ to preserve directed joins:

$$f\Big(\bigsqcup_{d \in D} d\Big) = \bigsqcup_{d \in D} f(d)$$

Read it as: *continuity is the statement that a function commutes with its big operator*. Exactly
the same shape as linearity of expectation ($\mathbb{E}$ commutes with $\sum$) and Fubini (swap two
$\int$s). Recursion is then $\operatorname{fix}(f) = \bigsqcup_{n} f^{n}(\bot)$ — the meaning of a
recursive definition is the join of its finite unfoldings.

### CRDTs — the identity/idempotence payoff in distributed systems

A state-based CRDT is a **join-semilattice**: merge is associative, commutative and *idempotent*.
Those three laws are precisely what makes re-delivered, out-of-order, duplicated network messages
harmless, so replicas converge without coordination:

$$\text{state} = \bigsqcup_{\text{messages received}} \text{payload}$$

Idempotence is the one law ordinary monoids don't require, and it is the entire reason the
technique works. Same family: G-counters, OR-sets, version vectors, and `max`-based last-write-wins
registers.

---

## 6. Type theory — $\prod$, $\sum$, and the algebra of data types

### Products and sums of types

| Type | Big form | Identity | In code |
|---|---|---|---|
| product | $\prod_{i} T_i$ | `unit` / `()` — one value | struct, record, tuple |
| sum | $\coprod_{i} T_i$, $\sum_i T_i$ | `void` / `never` — no values | enum, tagged union, sealed trait |

The identities are the punchline: the empty product is the **unit type** (exactly one value — the
empty tuple), and the empty sum is the **empty type** (no values at all). An empty struct is
inhabited; an empty enum is not. That mirrors $\prod_{\varnothing} = 1$ and $\sum_{\varnothing} =
0$ in arithmetic, and it is why they are called *algebraic* data types:

$$\lvert A \times B \rvert = \lvert A \rvert \cdot \lvert B \rvert \qquad
\lvert A + B \rvert = \lvert A \rvert + \lvert B \rvert \qquad
\lvert A \to B \rvert = \lvert B \rvert^{\lvert A \rvert}$$

Cardinality arithmetic really does predict type isomorphisms: $A \times (B + C) \cong A \times B +
A \times C$ is distributivity, and `Option<T>` is $1 + T$.

### Dependent $\Pi$ and $\Sigma$

When the second component's *type* depends on the first component's *value*, $\times$ and $\to$
grow into big operators indexed by a type:

$$\Pi_{x : A}\, B(x) \quad\text{— generalises } A \to B \text{ (a function whose result type varies)}$$
$$\Sigma_{x : A}\, B(x) \quad\text{— generalises } A \times B \text{ (a pair whose second type varies)}$$

With $A$ finite these literally are the indexed product and sum: $\Pi_{i \in \{1,2\}} B(i) \cong
B(1) \times B(2)$.

### Curry–Howard — the same table, twice

| Logic | Type theory | Set/§ in [[Big math operators]] |
|---|---|---|
| $\wedge$ | product / pair | $\cap$ |
| $\vee$ | sum / `Either` | $\cup$ |
| $\Rightarrow$ | function type | — |
| $\top$ | `unit` | $U$ |
| $\bot$ | `void` | $\varnothing$ |
| $\forall$ | $\Pi$ (generic / polymorphic type) | $\bigwedge$ |
| $\exists$ | $\Sigma$ (existential — an abstract data type) | $\bigvee$ |

$\forall$ being $\Pi$ is not an analogy: a generic function `fn id<T>(x: T) -> T` is an element of
$\Pi_{T}\, (T \to T)$, and an existential $\exists T.\,(\dots)$ is a module hiding its
representation — the interface/`dyn Trait`/vtable pattern, formalised.

### Recursive types as fixpoints

$$\texttt{List}\;a \;=\; \mu X.\; 1 + a \times X
\qquad\qquad \texttt{Stream}\;a \;=\; \nu X.\; a \times X$$

$\mu$ (least, inductive, finite — you can fold it) and $\nu$ (greatest, coinductive, possibly
infinite — you can unfold it) are §5's fixpoints applied to type constructors. `1` there is the
unit type: the empty list is the empty product, again.

### Intersection and union types

TypeScript's `A & B` and `A | B`, and the subtyping lattice generally:

$$\bigwedge_{i} T_i \quad \text{identity } \top \;(\texttt{unknown}/\texttt{any})
\qquad\qquad \bigvee_{i} T_i \quad \text{identity } \bot \;(\texttt{never})$$

Type inference joining the types of `if`/`match` branches is a $\bigsqcup$ over the subtyping
lattice; an empty match arm list yields `never`, which is the empty join. Rust's `!` type,
Kotlin's `Nothing` and TypeScript's `never` are all the same $\bot$.

---

## 7. Semantics and inference rules

An inference rule is a **big conjunction of premises** with one conclusion:

$$\dfrac{\displaystyle \bigwedge_{i=1}^{n} \big(\Gamma \vdash e_i : \tau_i\big)}{\Gamma \vdash e : \tau}$$

An **axiom** is a rule with no premises — the empty conjunction, $\top$. That is why axioms are
drawn with a bare line above them and need no justification: the empty-range identity yet again.

The two standard styles are two big-operator recursions:

- **Small-step** $e \to e'$: fold the transitive closure $\to^{*} = \bigcup_{n \ge 0} \to^{n}$
  (§4's Kleene star on a relation, and Rosen 8.4's transitive closure).
- **Big-step** $e \Downarrow v$: the derivation tree *is* the fold, evaluated bottom-up.

Denotational semantics writes $[\![ e ]\!]\rho$ for the meaning of $e$ in environment $\rho$, and
compositionality is the demand that $[\![\cdot]\!]$ be a homomorphism — the meaning of a compound
expression is a fold over the meanings of its parts. Compilers are that homomorphism, made
executable.

---

## 8. Linear logic in one page

The connectives you meet in session types, Rust's ownership story and separation logic. Every one
of the four has a unit, which is the identity of its big form:

| Connective | Reads as | Unit | Intuition |
|---|---|---|---|
| $\otimes$ (tensor) | "both, and you must use both" | $\mathbf{1}$ | a pair of resources |
| ⅋ (par) | "both, concurrently" | $\bot$ | the dual of tensor |
| $\&$ (with) | "either, *consumer* chooses" | $\top$ | an external choice / API surface |
| $\oplus$ (plus) | "either, *producer* chose" | $\mathbf{0}$ | a tagged union of resources |

The point for programming: in linear logic a hypothesis must be used *exactly once*, which is the
type-theoretic account of ownership, move semantics and file-handle protocols. The $\&$/$\oplus$
distinction is exactly the difference between "the caller picks which method to call" and "the
callee returns one of these variants".

Orientation only — but knowing that $\otimes$ here is *not* §8's tensor product of vector spaces
and *not* §5's XOR saves real confusion.

---

## 9. Category theory — where all of this is one definition

You do not need this to write compilers, but it explains why the tables keep rhyming.

- A **product** $\prod_i A_i$ and a **coproduct** $\coprod_i A_i$ are defined by universal
  properties, not by construction. In `Set` they are Cartesian product and disjoint union; in
  types they are structs and enums; in lattices they are meet and join. Same definition, three
  readings.
- The **terminal object** (identity of $\prod$) is `unit`; the **initial object** (identity of
  $\coprod$) is `void`. The empty product and empty sum from §6 are exactly these.
- A **limit** $\varprojlim$ and **colimit** $\varinjlim$ generalise every big operator in both
  notes: $\sum$, $\prod$, $\bigcup$, $\bigsqcup$, $\operatorname{lfp}$ are all (co)limits.
- A **catamorphism** is `fold` generalised to any recursive type; `foldr` on lists is the special
  case. "Bananas, lenses, envelopes and barbed wire" is the classic paper.
- "A monad is a monoid in the category of endofunctors" is the same monoid definition (§2) with
  $\odot$ = functor composition and $e$ = the identity functor. The joke is true; it is just
  §2 one level up.

---

## 10. Analysis of algorithms

Programming-adjacent, and mostly $\sum$ doing familiar work.

$$T(n) = \sum_{i=1}^{n} c_i \qquad\qquad
T(n) = a\,T(n/b) + f(n) \;\Rightarrow\; T(n) = \sum_{i=0}^{\log_b n} a^{i} f(n/b^{i})$$

The Master theorem (Rosen 7.3) is a classification of that geometric-ish sum by which end
dominates. Amortised analysis is the same move: the aggregate method bounds
$\sum_{i} \text{cost}_i$ directly, and the potential method telescopes
$\sum_i (\Phi_i - \Phi_{i-1}) = \Phi_n - \Phi_0$ — §3's telescoping identity, used to pay for
expensive operations with cheap ones.

**Big-O is not a big operator.** $O(f)$ is a *set of functions*, and $g = O(f)$ is an abuse of
notation for $g \in O(f)$. It has no binary partner and no identity; it appears here only because
the symbol looks like one and the family name collides.

---

## 11. Empty-input behaviour, by language

The most practical payoff in either note. All of these follow from the identity, not from taste:

| Expression | Value | Identity of |
|---|---|---|
| `sum([])` | `0` | $+$ |
| `math.prod([])` | `1` | $\times$ |
| `all([])` | `True` | $\wedge$ |
| `any([])` | `False` | $\vee$ |
| `"".join([])` | `""` | concatenation |
| `set().union(*[])` | `set()` | $\cup$ |
| `max([])` | **raises** | $-\infty$ is not in the type |
| `functools.reduce(op, [])` | **raises** without `initial` | the missing identity |

Two lessons: `max([])` and bare `reduce` fail precisely because the identity is unrepresentable or
unsupplied, and the fix is always to name the identity (`float('-inf')`, an explicit `initial=`).
Any `reduce` you write without an explicit identity has an undefined empty case — that is the bug
class this whole note exists to prevent.

---

## 12. Where it breaks

Everything in §14 of [[Big math operators]] still applies. Three additions specific to code:

1. **`foldr` vs `foldl` diverge for non-monoids.** Subtraction is not associative:
   `foldl (-) 0 [1,2,3]` is $-6$, `foldr (-) 0 [1,2,3]` is $2$. If your operation is a genuine
   monoid the two agree, and *that* is the practical test for whether you may parallelise.

2. **Side effects destroy commutativity.** A reducer that logs, mutates shared state, or allocates
   in order is not commutative even if its return value is. Parallel `reduce` will then produce
   correct numbers and wrong side effects.

3. **Lattice height decides termination.** A fold over a list ends because the list ends; a
   fixpoint iteration (§5) ends only if the lattice has no infinite ascending chains. Adding a
   "harmless" unbounded fact (an integer range, a set of allocation sites) to an analysis lattice
   turns a terminating pass into a hang — the reason widening exists.

---

## 13. Quick recall

- Which three laws must a MapReduce reducer satisfy, and what does each one let the framework do?
- Why is the identity of language concatenation $\{\varepsilon\}$ and not $\varnothing$?
- Write Kleene star as a big union, and say what $L^0$ is.
- Why does a dataflow analysis start the entry block at $\bot$? Derive it, don't recall it.
- State Kleene's fixpoint formula, and name the condition that makes the iteration terminate.
- What is the empty product of types? The empty sum? Which one is inhabited?
- Give the Curry–Howard partner of $\forall$, $\exists$, $\wedge$, $\vee$, $\bot$.
- Why is `all([])` true and `any([])` false — in one sentence, from the algebra.
- Name the extra law CRDT merge needs beyond monoid, and what network misbehaviour it tolerates.
- Why do `foldl` and `foldr` agree on `+` but not on `-`?
