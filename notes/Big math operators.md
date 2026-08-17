---
aliases:
  - Indexed operators
  - Iterated operators
  - Big operators
tags:
  - math/notation
  - reference
---

# Big math operators

Every "big" symbol in mathematics — $\sum$, $\prod$, $\bigwedge$, $\bigcup$, $\int$ — is the same
idea wearing different clothes: **take a binary operation and apply it across a whole indexed
family instead of just two things**. Learn the pattern once and every one of them reads itself.

This note collects every indexed operator you will meet across the curriculum
(`01-algebra-trigonometry-map.md`, `02-discrete-mathematics-map.md`, `04-calculus.md`,
`05-linear-algebra.md`, `06-probability.md`, `07-numerical-methods.md`), with its binary
counterpart, its identity element, and the one gotcha each.

---

## 1. The pattern

### Naming

The binary form gets the plain name; the big form gets a qualifier.

| Binary | Name | Big form | Name |
|---|---|---|---|
| $p \wedge q$ | conjunction ("and", "wedge") | $\bigwedge$ | big wedge · $n$-ary / indexed conjunction |
| $p \vee q$ | disjunction ("or") | $\bigvee$ | big vee · indexed disjunction |
| $a + b$ | addition | $\sum$ | big sigma · summation |
| $a \cdot b$ | multiplication | $\prod$ | big pi · product |
| $A \cup B$ | union | $\bigcup$ | big union · indexed union |

Nobody says "small conjunction". The family name is **iterated** or **indexed operators**.

### The fold analogy

$\sum_{i=1}^{n} a_i$ is a fold (reduce) over a range: bottom is the starting index, top is the
ending index, **and the top is inclusive** — unlike Go's `for i := 0; i < n; i++`.

```go
acc := 0                      // the identity, not a random zero
for i := 1; i <= n; i++ {     // note: <=, the top index is included
    acc = acc + a[i]          // the binary operation
}
```

Three parts, always the same: an **identity** to start from, a **range** to walk, a **binary
operation** to fold with. The analogy hides two things worth knowing.

### What makes it well defined

For $\bigodot_{i \in I} a_i$ to mean anything without extra parentheses, the operation $\odot$ must
be:

- **associative** — $(a \odot b) \odot c = a \odot (b \odot c)$, so no bracketing is needed;
- **commutative** — $a \odot b = b \odot a$, so the *order* of the index set does not matter.

$\wedge$, $\vee$, $+$, $\times$, $\cup$, $\cap$, $\max$ are all both. Matrix multiplication and
function composition are associative but **not** commutative, so their big forms need an explicit
order convention (§8). Floating-point addition is commutative but **not** associative (§14).

### The empty range

Every one of these has an **identity element** — the value of the operator over an empty index set.
This matters far more than it sounds like:

$$\bigodot_{i \in \varnothing} a_i = e \qquad\text{where } e \odot a = a \odot e = a$$

Empty sum is $0$, empty product is $1$, empty conjunction is *true*, empty disjunction is *false*.
It is why $0! = 1$ and $x^0 = 1$ (both empty products), why "for all elements of the empty set, $P$
holds" is vacuously true, and why a `reduce` over an empty slice must return the identity rather
than panic.

> A triple $(S, \odot, e)$ with an associative $\odot$ and identity $e$ is a **monoid**. Add
> commutativity and you have exactly the structure a big operator needs. Every row of the master
> table below is a commutative monoid — which is why they all behave the same.

---

## 2. Master table

| Big | Binary | Identity (empty range) | Where you meet it |
|---|---|---|---|
| $\sum$ | $+$ | $0$ | everywhere |
| $\prod$ | $\times$ | $1$ | 02 Ch. 4, 06, 07 |
| $\bigwedge$ | $\wedge$ | $\top$ (true) | 02 Ch. 1, Ch. 11 |
| $\bigvee$ | $\vee$ | $\bot$ (false) | 02 Ch. 1, Ch. 11 |
| $\bigoplus$ (XOR) | $\oplus$ | $0$ / false | 02 Ch. 12, 07 §5 |
| $\forall$ | $\wedge$ | true | 02 Ch. 1 |
| $\exists$ | $\vee$ | false | 02 Ch. 1 |
| $\bigcup$ | $\cup$ | $\varnothing$ | 02 Ch. 2, Ch. 8; 06 §2 |
| $\bigcap$ | $\cap$ | $U$ (the universe) | 02 Ch. 2; 06 §2 |
| $\biguplus$, $\bigsqcup$ | $\uplus$ | $\varnothing$ | 02 Ch. 6; 06 §2 |
| $\prod$ (Cartesian) | $\times$ | $\{()\}$ | 02 Ch. 2, Ch. 8 |
| $\mathop{\triangle}$ | $\triangle$ | $\varnothing$ | 02 Ch. 2 |
| $\max$, $\bigvee$ | $\max$ | $-\infty$ | 04, 06 §7, 07 |
| $\min$, $\bigwedge$ | $\min$ | $+\infty$ | 04, 02 Ch. 9 |
| $\sup$ / $\inf$ | — | $-\infty$ / $+\infty$ | 04 §§2, 11 |
| $\bigoplus$ (direct sum) | $\oplus$ | $\{\mathbf{0}\}$ | 05 §6 |
| $\bigotimes$ | $\otimes$ | the scalar field | 05 §10 |
| $\prod$ (ordered, matrices) | matrix $\times$ | $I$ | 05 §§3–4 |
| $\bigcirc$ (composition) | $\circ$ | $\mathrm{id}$ | 01 §10; 05 §4 |
| $\int$ | — | $\int_a^a = 0$ | 04 §§5–8 |
| $\mathop{\ast}$ (convolution) | $\ast$ | $\delta_0$ | 06 §6 |
| $\gcd$ | $\gcd$ | $0$ | 02 Ch. 4 |
| $\operatorname{lcm}$ | $\operatorname{lcm}$ | $1$ | 02 Ch. 4 |

---

## 3. Notation mechanics

### Ways to write the index set

$$\sum_{i=1}^{n} a_i \qquad \sum_{i \in S} a_i \qquad \sum_{\substack{1 \le i \le n \\ i \text{ odd}}} a_i
\qquad \sum_{d \mid n} d \qquad \sum_{i < j} a_{ij}$$

The range form is a special case of the set form: $\sum_{i=1}^{n} = \sum_{i \in \{1, \dots, n\}}$.
When the index set is empty the whole thing collapses to the identity — no special case needed.

### Iverson bracket

$[P]$ is $1$ when $P$ is true and $0$ when false. It turns a condition into a factor, which lets you
drop conditions out of the index:

$$\sum_{\substack{i \in S \\ P(i)}} a_i \;=\; \sum_{i \in S} a_i \, [P(i)]$$

Handy whenever a condition is getting in the way of an algebraic manipulation.

### The five manipulations you will use constantly

$$
\begin{aligned}
\text{pull out a constant} \quad & \sum_{i \in I} c\,a_i = c \sum_{i \in I} a_i \\[4pt]
\text{split the range} \quad & \sum_{i=1}^{n} a_i = \sum_{i=1}^{k} a_i + \sum_{i=k+1}^{n} a_i \\[4pt]
\text{reindex (shift)} \quad & \sum_{i=1}^{n} a_i = \sum_{j=0}^{n-1} a_{j+1} \\[4pt]
\text{swap order} \quad & \sum_{i \in I} \sum_{j \in J} a_{ij} = \sum_{j \in J} \sum_{i \in I} a_{ij} \\[4pt]
\text{telescope} \quad & \sum_{i=1}^{n} \big(f(i+1) - f(i)\big) = f(n+1) - f(1)
\end{aligned}
$$

Telescoping is the discrete Fundamental Theorem of Calculus, and it is the trick behind half the
induction proofs in `02-discrete-mathematics-map.md` Ch. 5.

### Reading a nested big operator

Read outside-in, and note that the inner range may depend on the outer index:

$$\sum_{i=1}^{n} \sum_{j=1}^{i} a_{ij} \quad\text{— a triangle, not a rectangle}$$

Swapping order here means re-describing the same region: $\sum_{i=1}^n \sum_{j=1}^i =
\sum_{j=1}^{n} \sum_{i=j}^{n}$.

---

## 4. Arithmetic — $\sum$ and $\prod$

### $\sum$ — summation ("big sigma")

Binary $+$ · identity $0$ · `\sum`

$$\sum_{i=1}^{n} a_i = a_1 + a_2 + \dots + a_n$$

Closed forms worth memorising (Greene §15, Rosen 2.5, Thomas 5.2):

$$\sum_{i=1}^{n} i = \frac{n(n+1)}{2} \qquad
\sum_{i=1}^{n} i^2 = \frac{n(n+1)(2n+1)}{6} \qquad
\sum_{i=0}^{n} r^i = \frac{1 - r^{n+1}}{1 - r} \;\;(r \ne 1)$$

$$\sum_{i=0}^{\infty} r^i = \frac{1}{1-r} \quad (\lvert r \rvert < 1)$$

That last one is the geometric series — expected retransmissions on a lossy link
(`06-probability.md` §9 lesson 46) and amortised cost analysis both come straight out of it.

**Where:** Greene §15 · Rosen 2.5, 6.4 · Thomas 5.2 (sigma notation), Ch. 10 (series) ·
Probability §4 (expectation) · Numerical Methods §5.

### $\prod$ — product ("big pi")

Binary $\times$ · identity $1$ · `\prod`

$$\prod_{i=1}^{n} a_i = a_1 \cdot a_2 \cdots a_n \qquad n! = \prod_{i=1}^{n} i \qquad 0! = \prod_{i \in \varnothing} i = 1$$

The bridge between the two: $\log$ turns $\prod$ into $\sum$.

$$\log \prod_{i=1}^{n} a_i = \sum_{i=1}^{n} \log a_i$$

This is not a curiosity — it is how you compute a product of thousands of probabilities without
underflowing a `float64` (`07-numerical-methods.md` §1), and it is why log-likelihoods exist.

**Where:** Rosen 4.3 (prime factorisation) · Probability §3 (independence) · Numerical Methods
§4 (Lagrange interpolation).

---

## 5. Logic — $\bigwedge$, $\bigvee$, $\bigoplus$, $\forall$, $\exists$

`02-discrete-mathematics-map.md` Ch. 1 and Ch. 11.

### $\bigwedge$ — big wedge · indexed conjunction

Binary $\wedge$ · identity $\top$ · `\bigwedge`

$$\bigwedge_{i=1}^{n} p_i = p_1 \wedge p_2 \wedge \dots \wedge p_n
\qquad \bigwedge_{i \in \varnothing} p_i = \top$$

### $\bigvee$ — big vee · indexed disjunction

Binary $\vee$ · identity $\bot$ · `\bigvee`

$$\bigvee_{i=1}^{n} p_i = p_1 \vee p_2 \vee \dots \vee p_n
\qquad \bigvee_{i \in \varnothing} p_i = \bot$$

De Morgan lifts to the big forms unchanged — the single most useful identity here:

$$\neg \bigwedge_{i \in I} p_i = \bigvee_{i \in I} \neg p_i
\qquad\qquad \neg \bigvee_{i \in I} p_i = \bigwedge_{i \in I} \neg p_i$$

### $\forall$ and $\exists$ are big operators in disguise

$$\forall x \in S,\; P(x) \;\equiv\; \bigwedge_{x \in S} P(x)
\qquad\qquad \exists x \in S,\; P(x) \;\equiv\; \bigvee_{x \in S} P(x)$$

The identities explain the notorious edge case exactly: over an empty domain $\forall$ is
**vacuously true** (empty conjunction $= \top$) and $\exists$ is **false** (empty disjunction
$= \bot$). And quantifier negation is just De Morgan:
$\neg \forall x\, P(x) \equiv \exists x\, \neg P(x)$.

Nested quantifiers (Rosen 1.5) are nested big operators, so the same rule applies: $\forall\exists$
and $\exists\forall$ are **not** interchangeable — you may only swap two operators of the *same*
kind.

### $\bigoplus$ — indexed XOR

Binary $\oplus$ · identity $0$ (false) · `\bigoplus`

$$\bigoplus_{i=1}^{n} b_i = b_1 \oplus b_2 \oplus \dots \oplus b_n = \Big(\sum_{i=1}^{n} b_i\Big) \bmod 2$$

XOR is addition in $\mathbb{F}_2$, which is why parity bits, checksums, CRCs (Rosen 12.7–12.8) and
stream ciphers are all the same operator. Also: XOR-folding a list is how you find the one
unpaired element, because $b \oplus b = 0$.

---

## 6. Sets — $\bigcup$, $\bigcap$, $\biguplus$, Cartesian $\prod$, $\mathop{\triangle}$

`02-discrete-mathematics-map.md` Ch. 2 and Ch. 8; `06-probability.md` §2 (events are sets).

### $\bigcup$ — indexed union

Binary $\cup$ · identity $\varnothing$ · `\bigcup`

$$\bigcup_{i \in I} A_i = \{\, x : \exists i \in I,\; x \in A_i \,\}$$

Note the shape: $\bigcup$ *is* $\exists$ applied to membership, and $\bigcap$ is $\forall$. Set
operations and logic are the same subject in two notations.

### $\bigcap$ — indexed intersection

Binary $\cap$ · identity $U$ · `\bigcap`

$$\bigcap_{i \in I} A_i = \{\, x : \forall i \in I,\; x \in A_i \,\}$$

> **The one genuine gotcha in this note.** The empty intersection is the *universe* $U$, not
> $\varnothing$ — every $x$ vacuously satisfies "belongs to all of no sets". This means $\bigcap$ is
> only well defined relative to an ambient universe, and in naive set theory
> $\bigcap_{i \in \varnothing} A_i$ has no answer at all (there is no set of everything). In
> practice: always intersect *within* a stated universe. In `06-probability.md` that universe is
> the sample space $\Omega$; in a database it is the table's domain.

De Morgan again, and distributivity:

$$\overline{\bigcup_{i \in I} A_i} = \bigcap_{i \in I} \overline{A_i}
\qquad B \cap \bigcup_{i \in I} A_i = \bigcup_{i \in I} (B \cap A_i)$$

### $\biguplus$ / $\bigsqcup$ — disjoint union

Binary $\uplus$ · identity $\varnothing$ · `\biguplus`, `\bigsqcup`

Union *plus the promise* that the pieces do not overlap — which is exactly what makes counting
easy. The sum rule (Rosen 6.1) and the probability axiom of countable additivity
(`06-probability.md` §2) are the same statement:

$$\Big\lvert \biguplus_{i=1}^{n} A_i \Big\rvert = \sum_{i=1}^{n} \lvert A_i \rvert
\qquad\qquad P\Big( \biguplus_{i=1}^{\infty} A_i \Big) = \sum_{i=1}^{\infty} P(A_i)$$

When the sets *do* overlap you pay the inclusion–exclusion tax — a formula that stacks three big
operators at once (Rosen 7.5, Probability §2 lesson 7):

$$\Big\lvert \bigcup_{i=1}^{n} A_i \Big\rvert
= \sum_{\varnothing \ne J \subseteq \{1,\dots,n\}} (-1)^{\lvert J \rvert + 1} \Big\lvert \bigcap_{j \in J} A_j \Big\rvert$$

### Indexed Cartesian product

Binary $\times$ · identity $\{()\}$, the set containing only the empty tuple · `\prod`

$$\prod_{i=1}^{n} A_i = A_1 \times A_2 \times \dots \times A_n
= \{\, (a_1, \dots, a_n) : a_i \in A_i \,\}$$

An element is a tuple; an $n$-ary relation (Rosen 8.2) is a subset of one of these — which is
literally the definition of a database table. Cardinality: $\lvert \prod A_i \rvert = \prod \lvert
A_i \rvert$, the product rule of counting. The identity being a one-element set (not the empty set)
is what makes $\lvert A^0 \rvert = 1$ come out right.

### $\mathop{\triangle}$ — indexed symmetric difference

Binary $\triangle$ · identity $\varnothing$ · `\triangle`

$$\mathop{\triangle}_{i=1}^{n} A_i = \{\, x : x \text{ lies in an \textit{odd} number of the } A_i \,\}$$

Associative and commutative, though it never looks like it. It is XOR on indicator functions — the
same operator as §5, seen through $x \in A \leftrightarrow [x \in A]$.

---

## 7. Order — $\max$, $\min$, $\sup$, $\inf$

### $\max$ and $\min$

Binary $\max$ / $\min$ · identity $-\infty$ / $+\infty$ · `\max`, `\min`

$$\max_{i \in I} a_i \qquad \min_{i \in I} a_i \qquad
\max_{i \in \varnothing} a_i = -\infty \qquad \min_{i \in \varnothing} a_i = +\infty$$

Associative, commutative, **and idempotent** ($\max(a,a) = a$) — the extra property that makes them
lattice operations rather than merely monoid operations. The infinite identities are why shortest-
path algorithms initialise distances to $+\infty$: that is not a sentinel hack, it is the identity
of $\min$ (`02-discrete-mathematics-map.md` 9.6).

### $\sup$ and $\inf$

The completion of $\max$/$\min$ to sets with no attained extreme: $\sup$ is the least upper bound,
$\inf$ the greatest lower bound. $\sup\{x : x < 1\} = 1$ even though the max does not exist. Needed
the moment limits get rigorous (`04-calculus.md` §2 lesson 11, and monotone convergence in §11).

### Lattices — the same symbols, one level up

In a partially ordered set (Rosen 8.6) and in Boolean algebra (Rosen Ch. 11), $\bigvee$ is **join**
(least upper bound) and $\bigwedge$ is **meet** (greatest lower bound), with identities $\bot$
(bottom) and $\top$ (top). The symbol reuse is deliberate: propositional logic is the two-element
Boolean lattice, $\max/\min$ is the lattice on a chain, and $\cup/\cap$ is the lattice on
$\mathcal{P}(U)$. Three subjects, one algebra.

### $\operatorname*{argmax}$ — the near miss

$$\operatorname*{argmax}_{i \in I} a_i = \text{the } i \text{ at which } a_i \text{ is largest}$$

It looks like a big operator but is not a fold over a monoid — it returns an *index*, ties make it
set-valued, and it has no identity. Worth knowing precisely because it breaks the pattern.

---

## 8. Linear algebra — $\prod$ (ordered), $\bigoplus$, $\bigotimes$, $\bigcirc$

`05-linear-algebra.md`.

### Ordered $\prod$ of matrices — where commutativity dies

Binary matrix $\times$ · identity $I$

$$\prod_{i=1}^{n} A_i = A_1 A_2 \cdots A_n$$

Matrix multiplication is associative but **not** commutative, so this notation is only meaningful
once you fix a convention (here: increasing $i$ left to right — always state it). This is
`05-linear-algebra.md` §3 lesson 14 in one line: rotate-then-translate $\ne$
translate-then-rotate, and the 2D transform pipeline (§4 lesson 20) is exactly such an ordered
product

$$M = M_{\text{screen}} \, M_{\text{view}} \, M_{\text{world}} \, M_{\text{local}}$$

read right to left as the point travels left. Matrix powers $A^n = \prod_{i=1}^n A$ *are* safe to
reorder against themselves, which is why diagonalisation (§9 lesson 43) works.

### $\sum$ over vectors — linear combinations

Identity $\mathbf{0}$, the zero vector. Every core definition of the subject is one of these:

$$\mathbf{v} = \sum_{i=1}^{n} c_i \mathbf{v}_i \qquad
(A\mathbf{x})_i = \sum_{j} a_{ij} x_j \qquad
(AB)_{ij} = \sum_{k} a_{ik} b_{kj}$$

Span, independence, rank and the matrix product are all statements about that first sum.

### $\sum$ over permutations — the determinant

The Leibniz formula packs an $n!$-term sum with an $n$-term product inside (§5, Greene §13):

$$\det A = \sum_{\sigma \in S_n} \operatorname{sgn}(\sigma) \prod_{i=1}^{n} a_{i,\sigma(i)}$$

A perfect example of an index set that is neither a range nor obvious — it is the symmetric group.
Also a perfect example of why you compute determinants by elimination instead (§5 lesson 25):
$n!$ terms is not an algorithm.

### $\bigoplus$ — direct sum of subspaces

Binary $\oplus$ · identity $\{\mathbf{0}\}$, the zero subspace · `\bigoplus`

$$V = \bigoplus_{i=1}^{k} U_i \quad\text{means every } \mathbf{v} \in V \text{ decomposes uniquely as } \sum_{i=1}^{k} \mathbf{u}_i$$

Dimensions add: $\dim V = \sum \dim U_i$. This is the precise version of "the eigenbasis splits
space into independent directions" (§9). The same symbol also means block-diagonal assembly of
matrices, and — confusingly — XOR (§5). Context always disambiguates; the shared symbol is because
all three are coproduct-like.

### $\bigotimes$ — tensor / Kronecker product

Binary $\otimes$ · identity the scalar field (the $1 \times 1$ matrix $[1]$) · `\bigotimes`

Dimensions **multiply** where $\bigoplus$ adds them: $\dim(U \otimes V) = \dim U \cdot \dim V$.
Orientation only at this level (§10 lesson 51's neighbourhood), but it is the operator behind
Kronecker products, multi-qubit state spaces, and the "outer product" view of rank-1 matrices.

### $\bigcirc$ — indexed composition

Binary $\circ$ · identity $\mathrm{id}$ · `\bigcirc`

$$\bigcirc_{i=1}^{n} f_i = f_1 \circ f_2 \circ \dots \circ f_n$$

Associative, not commutative — same caveat as matrix products, and not by coincidence: matrices
*are* linear functions (§3 lesson 12), so the two are the same monoid. Greene §10 lesson 301 is the
binary case.

### Einstein summation — the invisible $\sum$

$a_{ij} x_j$ with a repeated index means $\sum_j a_{ij} x_j$. You will see it in graphics and
physics writing; it is pure notation, no new content — the $\sum$ is just omitted.

---

## 9. Calculus — $\int$, $\lim$, infinite $\sum$ and $\prod$

`04-calculus.md`.

### $\int$ — the continuous limit of $\sum$

Identity $\int_a^a f = 0$ · `\int`

The definite integral literally *is* a limit of sums — that is the definition (Thomas 5.1–5.3,
lessons 33–35):

$$\int_a^b f(x)\,dx = \lim_{n \to \infty} \sum_{i=1}^{n} f(x_i^*)\, \Delta x$$

Every $\sum$ rule has an $\int$ twin, which is the fastest way to learn the integral's properties:

| $\sum$ | $\int$ |
|---|---|
| $\sum_{i \in \varnothing} a_i = 0$ | $\int_a^a f = 0$ |
| $\sum c\,a_i = c\sum a_i$ | $\int c f = c \int f$ |
| split the range at $k$ | $\int_a^b = \int_a^c + \int_c^b$ |
| telescoping | Fundamental Theorem of Calculus |
| swap two sums | Fubini's theorem (swap two integrals) |

The elongated S is a literal S, for *summa*. Beyond this module: $\iint$, $\oint$, $\int_C$ —
survey only (lesson 89).

### Infinite $\sum$ — where commutativity finally breaks

$$\sum_{i=1}^{\infty} a_i := \lim_{n \to \infty} \sum_{i=1}^{n} a_i$$

An infinite series is **not** a fold — it is a *limit of* folds, and the fold laws only survive
under absolute convergence. Riemann's rearrangement theorem (lesson 75): a conditionally convergent
series can be reordered to converge to **any** value you like. So "order of iteration doesn't
matter" is a theorem about finite index sets, and it expires here. Same caveat for swapping two
infinite sums, or a sum with an integral: legal under absolute convergence, not otherwise.

### $\lim$, $\limsup$, $\liminf$

`\lim` · not folds — no binary operation and no identity. They are indexed by a *process*
($n \to \infty$, $x \to a$), not by a set, so nothing in §3 applies to them. Grouped here only
because they share the big-operator *look*. $\limsup$ and $\liminf$ do combine with $\sup$/$\inf$
as folds over tails:
$\limsup_n a_n = \inf_{n} \sup_{k \ge n} a_k$.

### Infinite $\prod$

$$\prod_{i=1}^{\infty} (1 + a_i) \quad\text{converges iff}\quad \sum_{i=1}^{\infty} \log(1+a_i) \text{ converges}$$

Rare in this curriculum; listed so the $\log$ bridge from §4 is visibly the general tool.

---

## 10. Probability — $\mathbb{E}$, $\bigcup$, $\prod$, $\ast$

`06-probability.md`.

### $\mathbb{E}$ — expectation, a $\sum$ or $\int$ in disguise

$$\mathbb{E}[X] = \sum_{x} x\, p(x) \qquad\text{(discrete)} \qquad\qquad
\mathbb{E}[X] = \int_{-\infty}^{\infty} x\, f(x)\,dx \qquad\text{(continuous)}$$

The single most useful property in the subject is a statement about big operators commuting —
**linearity of expectation** (§4 lesson 17):

$$\mathbb{E}\Big[\sum_{i=1}^{n} X_i\Big] = \sum_{i=1}^{n} \mathbb{E}[X_i]$$

with **no independence required**. That is what makes it the workhorse of average-case analysis:
decompose into indicator variables, sum the easy expectations. Variance does *not* have this
property unless the $X_i$ are uncorrelated:
$\operatorname{Var}(\sum X_i) = \sum \operatorname{Var}(X_i) + 2\sum_{i<j} \operatorname{Cov}(X_i, X_j)$.

### $\prod$ — independence

$$P\Big(\bigcap_{i=1}^{n} A_i\Big) = \prod_{i=1}^{n} P(A_i) \quad\text{(mutually independent events)}$$

The complement trick (§2 lesson 7) — the most-used computation in the module — is this plus De
Morgan:

$$P\Big(\bigcup_{i=1}^{n} A_i\Big) = 1 - P\Big(\bigcap_{i=1}^{n} \overline{A_i}\Big) = 1 - \prod_{i=1}^{n}\big(1 - P(A_i)\big)$$

The birthday problem, hash-collision probability and series/parallel reliability are all one
application of that line.

### Countable additivity and total probability

$$P\Big(\biguplus_{i=1}^{\infty} A_i\Big) = \sum_{i=1}^{\infty} P(A_i)
\qquad\qquad P(B) = \sum_{i} P(B \mid A_i)\, P(A_i)$$

The second (law of total probability, §3) needs $\{A_i\}$ to be a partition — a $\biguplus$ of the
whole sample space. Bayes' theorem is that denominator plus one line of algebra.

### $\ast$ — convolution

Binary $\ast$ · identity $\delta_0$ (point mass at $0$) · `\ast`

Associative and commutative, so the indexed form is well defined. The distribution of a sum of
independent random variables is the convolution of their distributions (§6 lesson 32):

$$p_{X+Y}(z) = \sum_{k} p_X(k)\, p_Y(z-k) \qquad\qquad f_{X+Y}(z) = \int f_X(t)\, f_Y(z-t)\,dt$$

$n$-fold convolution is written $f^{\ast n}$. This is why sums smooth out, and — via the central
limit theorem — why they end up normal.

### Generating functions — a $\sum$ that encodes a whole sequence

Rosen 7.4 and `05-linear-algebra.md` §9 lesson 44 attack the same recurrences from two sides:

$$G(x) = \sum_{n=0}^{\infty} a_n x^n$$

Convolution shows up here too: the coefficients of $G(x)H(x)$ are $\sum_{k} a_k b_{n-k}$ — the same
formula as above, which is not a coincidence.

---

## 11. Number theory and algebra — divisor sums, $\gcd$, $\operatorname{lcm}$

`02-discrete-mathematics-map.md` Ch. 4 and Ch. 12.

### Indexing by divisibility

$$\sum_{d \mid n} f(d) \qquad\qquad \prod_{p \mid n} p \qquad\qquad n = \prod_{p} p^{\,e_p}$$

The index set is "the divisors of $n$" or "the primes dividing $n$" — a set, not a range, and the
last one is an infinite product in which all but finitely many factors are $1$ (the identity doing
real work). Fundamental theorem of arithmetic, Rosen 4.3.

### $\gcd$ and $\operatorname{lcm}$ as big operators

Both are associative, commutative and idempotent, so $\gcd(a_1, \dots, a_n)$ is unambiguous. Their
identities are the ones nobody guesses:

$$\gcd(\varnothing) = 0 \qquad\qquad \operatorname{lcm}(\varnothing) = 1$$

because $\gcd(n, 0) = n$ and $\operatorname{lcm}(n, 1) = n$. Under divisibility order they are
exactly meet and join (§7): the positive integers form a lattice with $1$ at the bottom and $\gcd$
as $\wedge$.

### Direct products of algebraic structures

$\prod_i G_i$ (direct product) and $\bigoplus_i G_i$ (direct sum — the sub-object where all but
finitely many coordinates are the identity) — Rosen 12.6. For finitely many factors they coincide.
The Chinese remainder theorem is the statement that
$\mathbb{Z}_{mn} \cong \mathbb{Z}_m \times \mathbb{Z}_n$ when $\gcd(m,n) = 1$.

---

## 12. Numerical methods — big operators as code

`07-numerical-methods.md`. Every method in this module is a big operator you actually evaluate.

### Quadrature — approximating $\int$ by $\sum$

$$\int_a^b f(x)\,dx \approx \sum_{i=0}^{n} w_i\, f(x_i)$$

Trapezoid and Simpson (§5 lesson 16, Thomas lesson 55) differ only in the weights $w_i$. The whole
subject of numerical integration is "choose good $w_i$ and $x_i$".

### Lagrange interpolation — a $\sum$ of $\prod$s

$$P(x) = \sum_{j=0}^{n} y_j \prod_{\substack{k=0 \\ k \ne j}}^{n} \frac{x - x_k}{x_j - x_k}$$

Read the inner product: it is $1$ at $x = x_j$ and $0$ at every other node — the empty-product
identity is what makes the $k \ne j$ exclusion come out clean (§4 lesson 13).

### Bernstein / Bézier — a $\sum$ with binomial weights

$$B(t) = \sum_{i=0}^{n} \binom{n}{i} (1-t)^{\,n-i} t^{\,i}\, P_i$$

The weights sum to $1$ for every $t$ (binomial theorem), which is exactly why a Bézier curve stays
inside the convex hull of its control points (§4 lesson 15).

### Norms — $\sum$, and $\max$ as its limit

$$\lVert x \rVert_p = \Big(\sum_{i=1}^{n} \lvert x_i \rvert^{\,p}\Big)^{1/p}
\qquad\qquad \lVert x \rVert_\infty = \max_i \lvert x_i \rvert = \lim_{p \to \infty} \lVert x \rVert_p$$

The $\infty$-norm is a $\max$-fold and the $2$-norm is a $\sum$-fold — two different monoids, one
family of error measures (§1 lesson 2).

### Horner's method — a fold that beats the naive $\sum$

$$\sum_{i=0}^{n} a_i x^i = a_0 + x\big(a_1 + x(a_2 + \dots)\big)$$

Same value, $n$ multiplications instead of $O(n^2)$, and better numerically. Greene lesson 55
(synthetic division) is this algorithm. Proof that *how* you fold matters even when the math says
it doesn't.

---

## 13. Same shape, different monoid

Swap the underlying monoid and the same formula computes something completely different. Matrix
"multiplication" over a semiring $(\oplus, \otimes)$:

$$(A \otimes B)_{ij} = \bigoplus_{k} \big(a_{ik} \otimes b_{kj}\big)$$

| Semiring | $\oplus$ | $\otimes$ | Identities | The same formula now computes |
|---|---|---|---|---|
| ordinary | $+$ | $\times$ | $0$, $1$ | matrix product, walk counts (Rosen 9.3) |
| Boolean | $\vee$ | $\wedge$ | false, true | reachability, transitive closure (Rosen 8.4, Warshall) |
| min-plus (tropical) | $\min$ | $+$ | $+\infty$, $0$ | shortest paths (Rosen 9.6, Floyd–Warshall) |
| max-plus | $\max$ | $+$ | $-\infty$, $0$ | critical path, longest path in a DAG |

Transitive closure written with big operators makes the point:

$$R^* = \bigcup_{n \ge 1} R^n$$

— the same union-fold that §6 defines, over the powers of a relation. Recognising this pattern is
worth more than memorising any single algorithm: Dijkstra, Warshall and matrix exponentiation are
one algorithm in three monoids.

---

## 14. Where the rules break

The four exceptions worth carrying in your head.

1. **Non-commutative operations.** Matrix products and function composition (§8) are associative
   but not commutative — the index order is part of the meaning, so always state the convention.
   $\forall$ and $\exists$ do not commute with each other either (§5).

2. **Infinite index sets.** Rearrangement and swapping (§3) are theorems about *finite* folds.
   For infinite sums they need absolute convergence; conditionally convergent series can be
   rearranged to any limit at all (§9).

3. **Floating point is not associative.** In `float64`, $(a + b) + c \ne a + (b + c)$ in general —
   so a parallel or reordered $\sum$ gives a different answer from a sequential one, and
   catastrophic cancellation (`07-numerical-methods.md` §1 lesson 2) can destroy the result
   entirely.

   ```go
   // Kahan summation: keep a running compensation for the lost low-order bits.
   sum, c := 0.0, 0.0
   for _, x := range a {
       y := x - c
       t := sum + y
       c = (t - sum) - y   // what the addition just threw away
       sum = t
   }
   ```

   Sorting by magnitude before summing, or pairwise summation, are the other standard fixes. The
   mathematical identity $\sum_{i \in \varnothing} = 0$ survives; associativity does not.

4. **The empty intersection.** $\bigcap_{i \in \varnothing} A_i = U$ needs a stated universe, and
   without one it is undefined (§6). Watch for it whenever an intersection is built in a loop.

---

## 15. Curriculum cross-reference

| Module | Operators it leans on |
|---|---|
| `01-algebra-trigonometry-map.md` | $\sum$, $\prod$ (§15 series, §16 binomial theorem), $\circ$ (§10), $\sum$ inside determinants (§13) |
| `02-discrete-mathematics-map.md` | $\bigwedge$, $\bigvee$, $\forall$, $\exists$ (Ch. 1, 11); $\bigcup$, $\bigcap$, $\biguplus$, Cartesian $\prod$ (Ch. 2, 8); $\sum$, $\prod$ (Ch. 2, 6, 7); $\sum_{d \mid n}$, $\gcd$ (Ch. 4); $\bigoplus$ XOR (Ch. 12); semiring folds (Ch. 8–9) |
| `04-calculus.md` | $\int$, infinite $\sum$ and $\prod$, $\lim$, $\sup$/$\inf$, Taylor's $\sum$ |
| `05-linear-algebra.md` | $\sum$ (combinations, matrix product), ordered $\prod$, $\bigoplus$, $\bigotimes$, $\bigcirc$, $\sum_{\sigma \in S_n}$ |
| `06-probability.md` | $\mathbb{E}$ as $\sum$/$\int$, $\biguplus$ additivity, $\prod$ independence, $\ast$ convolution, $\sum$ generating functions |
| `07-numerical-methods.md` | $\sum$ quadrature, $\sum$-of-$\prod$ interpolation, $\max$-norm, Horner's fold, and every caveat in §14 |

---

## 16. Quick recall

Cold-recall questions for the practice-test half of the rhythm:

- Name the binary partner and the identity of: $\sum$, $\prod$, $\bigwedge$, $\bigvee$, $\bigcup$,
  $\bigcap$, $\gcd$, $\min$, $\ast$, $\bigcirc$.
- Why is $0! = 1$, and why is "every element of $\varnothing$ is prime" true?
- Why is $\bigcap_{i \in \varnothing} A_i$ the universe and not $\varnothing$?
- Which two properties does an operation need before its big form is unambiguous, and name one
  operator in this curriculum that fails each.
- Write De Morgan for $\bigcup$ and for $\bigwedge$; then write quantifier negation and notice they
  are the same statement.
- State linearity of expectation and say exactly what it does *not* require.
- Give the three semirings from §13 and what the same matrix formula computes in each.
- Why does a parallel $\sum$ over `float64` disagree with the sequential one?
