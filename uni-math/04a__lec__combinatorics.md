# Discrete Mathematics — Lecture 4A: Combinatorics — Theoretical Supplement, Part I

> Translated from the handwritten lecture notes `pdf/lect4A-Kombinat-MatDysk.pdf`.

## Injective functions

**Injective (one-to-one) function:**

$$
f: X \xrightarrow{\ 1\text{-}1\ } Y
$$

$$
\forall x, y \qquad x \neq y \;\Rightarrow\; f(x) \neq f(y)
$$

The teacher's sketch — two distinct points $x_1, x_2 \in X$ go to two distinct points $y_1, y_2 \in Y$:

![[lec04a_p01_injective-function.svg]]

Margin reminder:

$$
\binom{n}{k} = \frac{n!}{k!\,(n-k)!}, \qquad 0! = 1
$$

## Theorem 1 (variations without repetition)

Let $r$ and $k$ be natural numbers such that $r \leq k$. Then there exist exactly

$$
\overline{\overline{V}}{}_r^{\,k} \;=\; \frac{k!}{(k-r)!} \;=\; k(k-1)\cdots(k-r+1) \qquad \left(\text{i.e. } \binom{k}{r}\cdot r!\right)
$$

distinct **injective** mappings of the set $\{x_1, x_2, \ldots, x_r\}$ into the set $\{y_1, y_2, \ldots, y_k\}$:

$$
V_r^{\,k} = \left\{ A :\; A: \{x_1, x_2, \ldots, x_r\} \xrightarrow{\ 1\text{-}1\ } \{y_1, y_2, \ldots, y_k\},\ A \text{ is injective} \right\}
$$

> **Notation:** the double bar $\overline{\overline{V}}{}_r^{\,k}$ denotes the cardinality of the set $V_r^{\,k}$ (as in Lecture 1). The teacher writes the sub-/superscripts both ways ($V_k^{\,r}$ and $V_r^{\,k}$); here they are normalized to $V_r^{\,k}$: $r$ = size of the source set, $k$ = size of the target set.

### Proof — Method 1

$k \geq r$.

$$
\{x_1, x_2, \ldots, x_r\} \xrightarrow{\quad f \quad} \{y_1, y_2, \ldots, y_k\}
$$

How do we build an injective function?

- $f(x_1) = \,?$ — we have $k$ possibilities;
- $f(x_2) = \,?$ (with the value $f(x_1)$ frozen) — we have $k-1$ possibilities;
- $f(x_3) = \,?$ (with the values $f(x_1)$ and $f(x_2)$ frozen) — we have $k-2$ possibilities;
- $\vdots$
- $f(x_r) = \,?$ (with the values $f(x_1), f(x_2), \ldots, f(x_{r-1})$ frozen) — we have $k-r+1$ possibilities.

So:

$$
\overline{\overline{V}}{}_r^{\,k} = k\cdot(k-1)(k-2)\cdots(k-r+1) = \frac{k!}{(k-r)!} = \frac{k!}{(k-r)!\;r!}\cdot r! = \binom{k}{r}\, r! \qquad \square
$$

Alternative method — by induction (omitted).

## Lemma 1 — permutations

If $k = r$, then $V_r^{\,r}$ is called the set of **permutations** $P_r$. From Theorem 1:

$$
\overline{\overline{P}}_r = \binom{k}{k}\cdot k! = k! \qquad \square
$$

## Theorem 2

The number of all permutations of a $k$-element set equals $k! = \overline{\overline{P}}_r$.

## Theorem 3 (all functions)

Denote by $W_r^{\,k}$ the set of **all** functions from the set $X \to Y$, where $\overline{\overline{X}} = r$ and $\overline{\overline{Y}} = k$. Then

$$
\overline{\overline{W}}{}_r^{\,k} = k^{\,r}
$$

> Here $k \geq 1$ and $r \geq 1$; there are **no restrictions** on the relation between $k$ and $r$. Now the functions need not be injective.

### Proof

$$
\{x_1, x_2, \ldots, x_r\} \xrightarrow{\quad f \quad} \{y_1, y_2, \ldots, y_k\}
$$

- $x_1 \longrightarrow \{y_1, y_2, \ldots, y_k\}$ — can be mapped in $k$ ways;
- $x_2 \longrightarrow \{y_1, y_2, \ldots, y_k\}$ — can be mapped in $k$ ways;
- $\vdots$
- $x_r \longrightarrow \{y_1, y_2, \ldots, y_k\}$ — can be mapped in $k$ ways.

$\Rightarrow$ the number of possible functions is:

$$
\underbrace{k \cdot k \cdot k \cdots k}_{r \text{ times}} = k^{\,r} \qquad \square
$$

## Theorem 4 (combinations)

For a given set $X$ of cardinality $\overline{\overline{X}} = n$, the number of all $k$-element subsets of the set $X$ ($k$-element **combinations** of the set $X$) equals:

$$
\binom{n}{k} = \overline{\overline{C}}{}_k^{\,n}
$$

### Proof

$$
X = \{e_1, e_2, \ldots, e_n\}, \qquad C_k^{\,n} = \{\, X_{i_k} \subseteq X \,\}
$$

$$
X_{i_k} = \{e_{i_1}, e_{i_2}, \ldots, e_{i_k}\} \subseteq X
$$

— a subset of the set $X$. It can be identified with a function:

![[lec04a_p05_index-boxes.svg|300]]

$$
f(1) = e_{i_1}, \quad f(2) = e_{i_2}, \quad \ldots, \quad f(k) = e_{i_k}
$$

and, say,

$$
\tilde{f}(1) = e_{i_2}, \quad \tilde{f}(2) = e_{i_1}, \quad \ldots, \quad \tilde{f}(k) = e_{i_k}
$$

— **different functions**, but they **represent the same subset**.

Note, however, that the order of the elements in $X_{i_k}$ is irrelevant, so any two functions $g$ that identify indices with the elements of $X_{i_k}$ must be classified as the same — there are $k!$ such possibilities.

$$
\Rightarrow \quad \overline{\overline{C}}{}_k^{\,n} = \binom{n}{k}\cdot k! \cdot \frac{1}{k!} = \binom{n}{k} \qquad \square
$$

## Theorem 5 (number of all subsets)

For a set $X$ of finite cardinality $\overline{\overline{X}} = n$, the number of **all** subsets is

$$
\overline{\overline{\mathcal{P}(X)}} = 2^{\overline{\overline{X}}}
$$

($\mathcal{P}(X)$ — the **power set**.)

### Proof — from the Binomial theorem (Method 1)

$$
(a+b)^n = \binom{n}{0}a^n b^0 + \binom{n}{1}a^{n-1} b^1 + \cdots + \binom{n}{n}a^0 b^n
$$

Set $a = b = 1$:

$$
2^n = \underbrace{\binom{n}{0}}_{\substack{\text{number of } 0\text{-elem.} \\ \text{subsets of } X}} + \underbrace{\binom{n}{1}}_{\substack{\text{number of } 1\text{-elem.} \\ \text{subsets of } X}} + \binom{n}{2} + \cdots + \underbrace{\binom{n}{n}}_{\substack{\text{number of } n\text{-elem.} \\ \text{subsets of } X}}
$$

### (Method 2 — characteristic functions)

Each subset $A \subseteq X$ is identified with a function $\chi_A: X \to \{0, 1\}$:

$$
A \subseteq X \quad \xrightarrow[\text{with a function}]{\text{we identify each subset } A} \quad \chi_A(x) = \begin{cases} 1 & x \in A \\ 0 & x \notin A \end{cases}
$$

— the **characteristic function**. The number of subsets $\equiv$ the number of functions from $X$ into $\{0,1\}$. By the previous theorem (Theorem 3):

$$
\overline{\overline{\mathcal{P}(X)}} = 2^{\overline{\overline{X}}}, \qquad \text{since } \overline{\overline{\{0,1\}}} = 2. \qquad \square
$$

## Theorem 6 (identical elements in distinguishable boxes)

There are

$$
\binom{n+k-1}{k-1}
$$

ways of placing $n$ **identical** elements into $k$ **distinguishable** boxes.

### Proof

Take a given distribution of the objects $(*)$ and encode it as a $0/1$ sequence $(**)$ — the 1s represent the boundaries between the boxes:

![[lec04a_p07_stars-and-bars.svg]]

**Remark:** two consecutive 1s ("$11$") denote an **empty urn**. We have $n$ zeros and $k-1$ ones.

There are as many distributions $(*)$ as there are sequences $(**)$. But there are as many sequences $(**)$ as there are choices of $(k-1)$-element subsets (each 1 represents a chosen element).

Example with $n = 5$, $k = 3$ (the seven positions are $e_1 e_2 e_3 e_4 e_5 e_6 e_7$):

$$
0\,0\,0\,1\,0\,1\,0 \;\longleftrightarrow\; \{e_4, e_6\} \qquad\qquad 0\,1\,1\,0\,0\,0\,0 \;\longleftrightarrow\; \{e_2, e_3\}
$$

So the sequences $(**)$ correspond to the $(k-1)$-element subsets of an $(n+k-1)$-element set — call these $({*}{*}{*})$. The cardinality of $({*}{*}{*})$:

$$
\binom{n+k-1}{k-1} = \text{cardinality of } (**) = \text{cardinality of } (*) \qquad \square
$$

> **Remark:** if the balls were distinguishable, e.g. $\{1, 2, \ldots, n\}$, then the placement of the 1s alone would no longer decide the distribution (what stands between the 1s would matter as well).

## Permutations with repetitions

It is also worth considering **permutations with repetitions**. Split $n$ cells into groups $G_1, G_2, \ldots, G_k$ with $n_1 = \overline{\overline{G_1}},\ n_2 = \overline{\overline{G_2}}, \ldots,\ n_k = \overline{\overline{G_k}}$:

![[lec04a_p08_grouped-cells.svg]]

— the elements **within each group** $G_1, G_2, \ldots, G_k$ are indistinguishable.

## Theorem 7

The number of variations (permutations, for $n = k$) **with repetitions** is:

$$
\overline{\overline{P}}{}^{\,n}_{n_1, n_2, \ldots, n_k} = \frac{n!}{n_1!\; n_2! \cdots n_k!}
$$

### Proof

The number of all permutations is $n!$. We eliminate the identical copies:

$$
\underbrace{n_1!}_{\text{that many within } G_1} \cdot \underbrace{n_2!}_{\text{within } G_2} \cdot\; \ldots \;\cdot\; n_k!
$$

$$
P^{\,n}_{n_1 n_2 \ldots n_k} = \frac{n!}{n_1!\; n_2! \cdots n_k!} \qquad \square
$$

## Ordered partitions

$A$ — a set. $(A_1, A_2, \ldots, A_k)$ — an **ordered partition** of $A$:

- a) $A_i \cap A_j = \emptyset$ — (a) and (b) together make a **partition**;
- b) $\displaystyle\bigcup_{i=1}^{k} A_i = A$;
- c) the order of the elements **within** each $A_i$ is irrelevant;
- d) the order of $(A_1, A_2, \ldots, A_k)$ **matters**:

$$
(A_1, A_2, \ldots, A_k) \neq \text{e.g. } (A_2, A_1, \ldots, A_k)
$$

(We assume fixed cardinalities $|A_k| = n_k$.)

## Theorem 9

> The teacher's numbering jumps from Theorem 7 straight to Theorem 9 — there is no Theorem 8 in the notes.

$A$ has $n$ elements, with

$$
n = n_1 + n_2 + \cdots + n_k.
$$

There exist

$$
\frac{n!}{n_1!\; n_2! \cdots n_k!}
$$

ordered partitions $(A_1, A_2, \ldots, A_k)$ of this set such that $\overline{\overline{A_i}} = n_i$.

### Proof

**Either** we identify each ordered partition with a permutation with repetitions [Theorem 7],

**or**:

$$
\overline{\overline{I}}_1 = \binom{n}{n_1}
$$

— that many ways to draw an $n_1$-element subset from the $n$-element set (into $G_1$).

$$
\overline{\overline{I}}_2 = \binom{n}{n_1} \cdot \binom{n-n_1}{n_2}
$$

— that many ways to draw an $n_2$-element subset into $G_2$, given the choice already made for $G_1$.

$$
\overline{\overline{I}}_3 = \binom{n}{n_1}\binom{n-n_1}{n_2}\binom{n-n_1-n_2}{n_3}
$$

— that many ways to draw an $n_3$-element subset into $G_3$, given the choices made for $G_1$ and $G_2$.

$$
\vdots
$$

$$
\overline{\overline{I}}_k = \binom{n}{n_1}\binom{n-n_1}{n_2}\binom{n-n_1-n_2}{n_3}\cdots\binom{n-n_1-n_2-\cdots-n_{k-1}}{n_k}
$$

— that many ways to draw an $n_k$-element subset into $G_k$, given the choices made for $G_1, G_2, \ldots, G_{k-1}$. (The last factor equals $1$, since $n - n_1 - n_2 - \cdots - n_{k-1} = n_k$.)

Expanding the binomial coefficients:

$$
I_k = \frac{n!}{n_1!\,(n-n_1)!} \cdot \frac{(n-n_1)!}{n_2!\,(n-n_1-n_2)!} \cdot \frac{(n-n_1-n_2)!}{n_3!\,(n-n_1-n_2-n_3)!} \cdot\; \cdots
$$

$$
\cdots\; \cdot \frac{(n-n_1-n_2-\cdots-n_{k-2})!}{n_{k-1}!\,(n-n_1-n_2-\cdots-n_{k-2}-n_{k-1})!} \cdot \frac{(n-n_1-n_2-\cdots-n_{k-1})!}{n_k!\,\underbrace{(n-n_1-n_2-\cdots-n_{k-1}-n_k)!}_{0! \,=\, 1}}
$$

— the factorials cancel in a telescoping fashion (each numerator cancels the $(\cdots)!$ in the previous denominator), so:

$$
I_k = \frac{n!}{n_1!\; n_2! \cdots n_k!} \qquad \square
$$

## Theorem 10 (the product rule)

Sets $S_i$ with cardinalities $\overline{\overline{S_i}} = k_i$.

$\Rightarrow$ the number of ordered $k$-tuples (sequences):

$$
|S_1 \times S_2 \times \cdots \times S_k| = \prod_{i=1}^{k} \overline{\overline{S_i}}
$$
