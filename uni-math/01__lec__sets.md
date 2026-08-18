# Discrete Mathematics — Lecture 1: Sets

> Translated from the handwritten lecture notes `pdf/lect1-zbiory-MatDysk.pdf`.

## Sets — basic notions

**Set** — a collection of objects called **elements**.

E.g. for a *finite* set:

$$
S = \{ e_1, e_2, \ldots, e_n \}
$$

$e_i \in S$ — membership of $e_i$ in the set $S$.

**Example:**

$$
\mathbb{N} = \{0, 1, 2, \ldots\}
$$

— the set of natural numbers (an *infinite* set).

$$
\mathbb{Z} = \{\ldots, -2, -1, 0, 1, 2, \ldots\}
$$

— the set of integers.

- $\mathbb{Z}_+$ (or $\mathbb{P}$) — the set of positive integers
- $\mathbb{Z}_-$ — the set of negative integers

$$
\mathbb{Q} = \left\{ w : w = \frac{m}{n},\ m, n \in \mathbb{Z},\ n \neq 0 \right\}
$$

— the set of rational numbers.

- $\mathbb{R}$ — the set of real numbers
- $\mathbb{C}$ — the set of complex numbers
- $\mathbb{IQ}$ — the set of irrational numbers

A non-mathematical example — the set of all people (a person belongs to the Earth):

![[lec01_p02_set-of-all-people.svg]]

## Definition 1 — set operations

### a) $A \subseteq B$ — ($A$ is a **subset** of $B$)

$$
A \subseteq B \iff \forall a \in A \Rightarrow a \in B
$$

Venn diagram:

![[lec01_p02_venn-subset.svg]]

### b) $A \cap B$ — (**intersection** of $A$ and $B$)

$$
a \in A \cap B \iff a \in A \text{ and } a \in B
$$

![[lec01_p02_venn-intersection.svg]]

### c) $A \cup B$ — (**union** of $A$ and $B$)

$$
a \in A \cup B \iff a \in A \text{ or } a \in B
$$

![[lec01_p02_venn-union-card.svg|300]]

### d) $A \setminus B$ — (**difference** of $A$ and $B$)

$$
a \in A \setminus B \iff a \in A \text{ and } a \notin B
$$

![[lec01_p03_venn-difference.svg]]

### e) $A = B$ — (**equality of sets**)

$A = B$ if they have exactly the same elements.

### f) $\emptyset$ — the **empty set**

It has no elements.

### g) $A \times B$ — (**Cartesian product** of $A$ and $B$)

The set of ordered pairs $(a, b)$ such that $a \in A$ and $b \in B$.

> **Note:** in general $(a,b) \neq (b,a)$.

### h) Finite unions and intersections

$$
\bigcup_{i=1}^{n} A_i = A_1 \cup A_2 \cup \ldots \cup A_n
$$

— union of a finite number of sets.

$$
\bigcap_{i=1}^{n} A_i = A_1 \cap A_2 \cap \ldots \cap A_n
$$

— intersection of a finite number of sets.

The teacher's margin doodles illustrating the union and intersection of several sets:

![[lec01_p03_union-intersection-doodles.svg|120]]

### i) $\mathcal{P}(S)$ — (**power set** of the set $S$)

The set of all subsets of the set $S$. Sometimes written with the alternative notation $2^S$.

![[lec01_p04_powerset-doodle.svg|120]]

### j) $A \oplus B \equiv (A \cup B) \setminus (A \cap B)$ — **symmetric difference**

**Example:**

$$
A = \{e_1, e_2, e_3\}, \qquad B = \{e_1, e_4\}, \qquad C = \{e_4, e_5\}
$$

$$
\begin{aligned}
A \cap B &= \{e_1\} \\
A \cup B &= \{e_1, e_2, e_3, e_4\} \\
A \setminus B &= \{e_2, e_3\} \\
B \setminus A &= \{e_4\} \\
\mathcal{P}(B) &= \{\, \emptyset, \{e_1\}, \{e_4\}, \{e_1, e_4\} \,\} \\
\mathcal{P}(C) &= \{\, \emptyset, \{e_4\}, \{e_5\}, \{e_4, e_5\} \,\}
\end{aligned}
$$

$$
\begin{aligned}
A \times A = \{ &(e_1,e_1), (e_1,e_2), (e_1,e_3), \\
&(e_2,e_1), (e_2,e_2), (e_2,e_3), \\
&(e_3,e_1), (e_3,e_2), (e_3,e_3) \}
\end{aligned}
$$

$$
A \cap B \cap C = \emptyset, \qquad A \cup B \cup C = \{e_1, e_2, e_3, e_4, e_5\}
$$

$$
A \oplus B = B \oplus A = \{e_2, e_3, e_4\}
$$

## Alphabets and words

**Example:**

$$
\Sigma = \{a, b, c, d, \ldots, z\}
$$

— the 26 letters of the English alphabet.

- A **word** over the alphabet $\Sigma$ — any finite sequence of letters from the alphabet $\Sigma$.
- The set of all words of a given alphabet is denoted $\Sigma^*$ — a **language** over the alphabet $\Sigma$ is any subset of $\Sigma^*$.
- The **length** of a word is its number of characters. $\square$

## The universe and the complement

**Remark:** $U$ — a fixed set (the **universe**).

![[lec01_p05_universe.svg]]

$U \setminus A$ — the **complement of $A$**, also denoted:

$$
\bar{A}, \quad A^c, \quad -A, \quad A'
$$

**Example:** $U = \mathbb{R}^2$ (the plane); $A$ is a disc, so $A'$ is the plane with a hole:

![[lec01_p05_plane-complement.svg]]

## Laws of set algebra

1. **Commutativity**
   - a) $A \cup B = B \cup A$
   - b) $A \cap B = B \cap A$
2. **Associativity**
   - a) $(A \cup B) \cup C = A \cup (B \cup C)$
   - b) $(A \cap B) \cap C = A \cap (B \cap C)$
3. **Distributivity**
   - a) $A \cup (B \cap C) = (A \cup B) \cap (A \cup C)$
   - b) $A \cap (B \cup C) = (A \cap B) \cup (A \cap C)$
4. **Idempotence laws**
   - a) $A \cup A = A$
   - b) $A \cap A = A$
5. **Identity laws**
   - a) $A \cup \emptyset = A$
   - b) $A \cup U = U$
   - c) $A \cap \emptyset = \emptyset$
   - d) $A \cap U = A$
6. **Double complement**

$$
(A^c)^c = A
$$

7. **Complement laws**
   - a) $A \cup A^c = U$
   - b) $A \cap A^c = \emptyset$
8. Complements of the universe and the empty set

$$
U^c = \emptyset, \qquad \emptyset^c = U
$$

9. **De Morgan's laws**
   - a) $(A \cup B)^c = A^c \cap B^c$
   - b) $(A \cap B)^c = A^c \cup B^c$

**Venn diagrams** easily show the truth of the above rules. E.g. for the distributive law — the left diagram shades $A \cap (B \cup C)$, the right one $(A \cap B) \cup (A \cap C)$, and both shaded regions coincide:

![[lec01_p07_venn-distributive.svg]]

## Proof of the distributive law

$$
(*) \quad A \cap (B \cup C) \stackrel{?}{=} (A \cap B) \cup (A \cap C)
$$

**Or a proof** using the rule (equality of sets via double inclusion):

$$
\tilde{A} = \tilde{B} \;\equiv\; \tilde{A} \subseteq \tilde{B} \ \text{ and } \ \tilde{B} \subseteq \tilde{A}
$$

So we must show:

$$
\begin{aligned}
\text{(i)} \quad & A \cap (B \cup C) \subseteq (A \cap B) \cup (A \cap C) \\
\text{(ii)} \quad & (A \cap B) \cup (A \cap C) \subseteq A \cap (B \cup C)
\end{aligned}
$$

(i) and (ii) together prove $(*)$.

**(i)** Let $a \in A \cap (B \cup C)$. Then $a \in A$ and $a \in B \cup C$, i.e. $a \in B$ or $a \in C$.

- **(i1)** if $a \in B$ and $a \in A$ $\Rightarrow$ $a \in A \cap B$, so $a \in (A \cap B) \cup (\text{any set})$, hence

$$
a \in (A \cap B) \cup (A \cap C)
$$

- **(i2)** if $a \in C$ and $a \in A$ $\Rightarrow$ $a \in A \cap C$, so $a \in (A \cap C) \cup (\text{any set})$, hence

$$
a \in (A \cap C) \cup (A \cap B)
$$

So in both cases (i1) and (i2) we have shown that if $a \in A \cap (B \cup C)$, then $a \in (A \cap B) \cup (A \cap C)$. Therefore (i) is true.

**(ii)** Now let

$$
a \in (A \cap B) \cup (A \cap C)
$$

- **(ii1)** $a \in A \cap B \Rightarrow a \in A$ and $a \in B$, so $a \in A$ and $a \in B \cup (\text{any set})$, hence $a \in A$ and $a \in B \cup C$, i.e.

$$
a \in A \cap (B \cup C)
$$

- **(ii2)** $a \in A \cap C \Rightarrow a \in A$ and $a \in C$, so $a \in A$ and $a \in C \cup (\text{any set})$, hence $a \in A$ and $a \in C \cup B$, i.e.

$$
a \in A \cap (B \cup C).
$$

We have shown that if $a \in (A \cap B) \cup (A \cap C) \Rightarrow a \in A \cap (B \cup C)$. So (ii) is proved. $\square$

## Cardinality of a set

**Remark:** for a finite set $A$:

$$
|A| \ \text{ or } \ \overline{\overline{A}}
$$

denotes the **number of elements** in the set.

**Example:**

$$
A = \{1, 2, 3, 4\} \quad \Rightarrow \quad \overline{\overline{A}} = 4
$$

$\overline{\overline{\mathbb{N}}}$ is an infinite number\* .

> \* infinite numbers can be distinguished using Cantor's numbers.

## Cartesian product of $k$ sets

**Remark:** for sets $A_1, A_2, \ldots, A_k$:

$$
A_1 \times A_2 \times \ldots \times A_k
$$

— the **$k$-fold Cartesian product**: the set of ordered **$k$-tuples** $(a_1, a_2, \ldots, a_k)$ with $a_i \in A_i$.

$A_1 \times A_2 \times \ldots \times A_k$ is also called the **product of the sets $A_k$**.

If $A_1 = A_2 = \ldots = A_k$, we write:

$$
A^k = \underbrace{A \times A \times \ldots \times A}_{k \text{ times}}
$$

## Inclusion–exclusion principle

(We will extend it later when we get to counting.)

**a)** $A$ and $B$ disjoint:

$$
|S \cup T| = |S| + |T|
$$

![[lec01_p10_disjoint-count.svg|200]]

**b)** in general:

$$
|S \cup T| = |S| + |T| - |S \cap T|
$$

![[lec01_p10_overlap-count.svg|250]]

## Infinite unions and intersections

Unions and intersections of sets can be generalized to **infinite** "summing" and "intersecting" of sets:

$$
\{A_i\}_{i=1}^{\infty} = \{A_i\}_{i \in \mathbb{N}}, \qquad \{B_i\}_{i=1}^{\infty} = \{B_i\}_{i \in \mathbb{N}}
$$

(when the index set is "equinumerous" with the set of natural numbers).

$$
\boxed{\ \bigcup_{i=1}^{\infty} A_i = A\ } \quad \stackrel{\text{def}}{\equiv} \quad x \in \bigcup_{i=1}^{\infty} A_i \ \text{ if there exists } k_0 \in \mathbb{N} \text{ such that } x \in A_{k_0}
$$

$$
\boxed{\ \bigcap_{i=1}^{\infty} B_i = B\ } \quad \stackrel{\text{def}}{\equiv} \quad x \in \bigcap_{i=1}^{\infty} B_i \ \text{ if for every } i \in \mathbb{N}, \ x \in A_i
$$

**Example:** $A_i = \left[0, \frac{1}{n}\right)$ for $n = 1, 2, \ldots$

$$
A_1 = [0, 1), \quad A_2 = \left[0, \tfrac{1}{2}\right), \quad \ldots, \quad A_n = \left[0, \tfrac{1}{n}\right), \ \ldots
$$

![[lec01_p12_intervals.svg]]

$$
\text{(i)} \ \bigcup_{i=1}^{\infty} A_i = [0, 1) \qquad \text{(ii)} \ \bigcap_{i=1}^{\infty} A_i = \{0\}
$$

**(i)** $\bigcup_{i=1}^{\infty} A_i = [0,1)$, since it suffices to take $k_0 = 1$: $x \in A_1$. No $x \notin [0,1)$ satisfies the condition that $\exists k_0 \geq 1$ with $x \in A_{k_0}$.

**(ii)** $\bigcap_{i=1}^{\infty} A_i = \{0\}$: for all $k \geq 1$, $x \in \{0\}$ works. No other $x \notin \{0\}$ satisfies the condition $\forall k \geq 1: x \in A_k$. $\square$

## General indexed families

**In general\*:** for a family $\{A_t\}_{t \in T}$:

$$
\boxed{\ \bigcup_{t \in T} A_t = A_1\ } \quad \stackrel{\text{def}}{\equiv} \quad x \in A_1 \equiv \exists\, t_0 \in T: \ x \in A_{t_0}
$$

("there exists")

$$
\boxed{\ \bigcap_{t \in T} A_t = A_2\ } \quad \stackrel{\text{def}}{\equiv} \quad x \in A_2 \equiv \forall\, t_0 \in T: \ x \in A_{t_0}
$$

("for every")

**(\*)** The index set may be infinite, and may even be "more numerous" than the natural numbers (e.g. $T = \mathbb{R}$).

Here one should introduce the notion of **cardinal numbers** (or ordinal numbers), which distinguish between different infinities, e.g.:

$$
\overline{\overline{\mathbb{N}}} = \aleph_0 \ \text{("aleph zero")}, \qquad \overline{\overline{\mathbb{R}}} = \mathfrak{c} \ \text{("continuum")}
$$

This is not the topic of this lecture.

## Examples of set operations

### a) Sets in the plane

$$
A = \{(x, y) : x^2 + y^2 \leq 1\}, \qquad B = \{(x, y) : x > 0\}, \qquad U = \mathbb{R}^2 \ \text{(the plane)}
$$

$A$ — the disc centered at $(0,0)$ with radius $R = 1$; $B$ — an open half-plane. The teacher's sketches of $A$, $B$, $A \cup B$, $A \cap B$, $A \setminus B$, $B \setminus A$, and (for $U = \mathbb{R}^2$) the complements $A'$, $B'$:

![[lec01_p14_plane-set-operations.svg]]

Here $A \not\subseteq B$ and $B \not\subseteq A$.

### b) Intervals on the real line

$$
A = (0, 2], \qquad B = [1, 3) \cup \{-1\}
$$

![[lec01_p15_number-lines.svg]]

For $U = \mathbb{R}$:

$$
\begin{aligned}
A \cup B &= \{-1\} \cup (0, 3) \\
A \cap B &= [1, 2] \\
A \setminus B &= (0, 1) \\
B \setminus A &= \{-1\} \cup (2, 3) \\
A' &= (-\infty, 0] \cup (2, +\infty) \\
B' &= (-\infty, -1) \cup (-1, 1) \cup [3, +\infty)
\end{aligned}
$$

### c) Cartesian products are not commutative

$$
A = \{1, 2\}, \qquad B = \{\bullet, \star\}
$$

$A \times B \neq B \times A$, because:

$$
\begin{aligned}
A \times B &= \{(1, \bullet), (1, \star), (2, \bullet), (2, \star)\} \\
B \times A &= \{(\bullet, 1), (\bullet, 2), (\star, 1), (\star, 2)\}
\end{aligned}
$$

$A$ and $B$ are disjoint because $A \cap B = \emptyset$; also $A \not\subseteq B$ and $B \not\subseteq A$.

### d) A power set

$$
A = \{\square, 1\} \quad \Rightarrow \quad \mathcal{P}(A) = \{\, \emptyset, \{\square\}, \{1\}, \{\square, 1\} \,\}
$$
