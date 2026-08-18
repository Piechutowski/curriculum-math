# Discrete Mathematics — Lecture 14b: Binary Trees and Closing Curiosities

> Translated from the handwritten lecture notes `pdf/lect14-b-drzewaBin+ciekawostki-MatDysk.pdf`.

## Trees

A **tree** is:

- an **undirected** graph (a *connected* one),
- **without loops**,
- **without cycles**.

The teacher's margin doodle of a tree:

![[lec14b_p01_tree-doodle.svg|140]]

## Recursive definition of the set of binary trees $\mathcal{T}$

- $s$ — symbolizes an arbitrary content of a tree node,
- $E$ — the symbol of the **empty tree**.

1. **Base step:** the empty binary tree
   - a) $E \in \mathcal{T}$
   - b) $(E, s, E) \in \mathcal{T}$ — (a binary tree with one node)
2. **Inductive step:**

$$
T_1, T_2 \in \mathcal{T} \;\Rightarrow\; T \stackrel{\text{def}}{=} (T_1, s, T_2) \in \mathcal{T}
$$

- $T_1$ — we call it the **left subtree** of the tree $T$,
- $T_2$ — we call it the **right subtree** of the tree $T$,
- $s$ — the content $s$ sits in the **root** of $T$.

Moreover:

- if $T_1 \neq E$ $\Rightarrow$ the root of $T_1$ is the **left child** of the root of the tree $T$;
- if $T_2 \neq E$ $\Rightarrow$ the root of $T_2$ is the **right child** of the root of the tree $T$.

**Remark:** formally, a tree $T$ is a string over the alphabet $\{s, E, (, )\}$ defined above according to schemes 1) and 2).

In computer science, in the construction $(T_1, s, T_2)$:

- $s$ represents a data structure,
- the $T_i$ represent references.

$T_1$ is different from $T_2$ when the strings defining them are different.

**Depth of a node** $\stackrel{\text{def}}{=}$ the level of nesting of the node minus $1$. The maximal depth is the **height of the tree**, $H(T)$.

## Example

**a)** A tree $T_a$ with a root and two one-node subtrees:

![[lec14b_p03_tree-ta.svg|260]]

$$
T_a = (\underbrace{(E, s, E)}_{T_{a_1}},\; s,\; \underbrace{(E, s, E)}_{T_{a_2}})
$$

$H(T_a) = 1$; $T_a$ has $2$ subtrees (each has $1$ node).

**b)** A tree $T_b$ that is a chain of left children:

![[lec14b_p03_tree-tb.svg|220]]

$$
T_b = (\underbrace{(((E,s,E),\,s,\,E),\,s,\,E)}_{T_{b_1}},\; s, \underbrace{E}_{T_{b_2}})
$$

$4$ nodes; $1$ non-empty subtree $T_{b_1}$; $H(T_b) = 3$.

## Counting binary trees

Let $\mathcal{T}_n$ denote the set of all binary trees with $n$ **nodes**, and let

$$
t(n) = |\mathcal{T}_n|
$$

> Margin note (see later): $t(3) = \frac{1}{3+1}\binom{2 \cdot 3}{3} = \frac{1}{4} \cdot \frac{6!}{3!\,3!} = \frac{6 \cdot 5 \cdot 4}{4 \cdot 3!} = 5$.

**Example** ($n = 3$). The circle $\odot$ marks the root:

![[lec14b_p03_five-trees.svg]]

We have $5$ trees: $t(3) = 5$.

**Note:** if we did not divide a graph (tree) into left and right subtrees, then trees $(2)$–$(5)$ would be equivalent.

## The recurrence for $t(n)$

Our goal — we want to compute $t(n) = \,?$

Initial conditions:

$$
\begin{cases}
t(0) = 1 & \text{because } E \in \mathcal{T} \ \text{(the empty subtree)} \\
t(1) = 1 & \text{because } (E, s, E) \in \mathcal{T} \ \text{(the tree having 1 node)}
\end{cases}
$$

For $n$ vertices (the root takes one vertex; the remaining $n-1$ vertices split between the left and right subtree):

- $a_0$) zero vertices on the left, $\ n-1$ vertices on the right;
- $a_1$) $1$ of the vertices on the left, $\ n-2$ vertices on the right;
- $\ \vdots$
- $a_k$) $k$ vertices on the left, $\ n-k-1$ vertices on the right;
- $\ \vdots$
- $a_{n-1}$) $n-1$ vertices on the left, $\ 0$ vertices on the right.

The teacher's sketches of the cases (triangles stand for subtrees with the indicated number of vertices):

![[lec14b_p04_split-cases.svg]]

Therefore the recurrence relation is the following:

$$
(*) \quad
\begin{cases}
\ t(n) = t(0)\,t(n-1) + t(1)\,t(n-2) + \ldots + t(n-2)\,t(1) + t(n-1)\,t(0), \\[4pt]
\oplus \ \text{ initial values } \ t(0) = t(1) = 1.
\end{cases}
$$

> **Note:** a **nonlinear** recurrence relation! Degree $n$.

We want to find the explicit (closed) form of $t(n)$ from $(*)$:

$$
(\blacksquare) \quad \boxed{\ t(n) = \frac{1}{n+1}\binom{2n}{n}\ }
$$

## Proof (generating functions)

**a)** We first define the **generating function**

$$
F(x) = \sum_{n=0}^{\infty} t(n)\, x^n
\qquad \left(\text{an application to the nonlinear scheme } (*)\right)
$$

**b)** Let us compute

$$
x F^2(x) = x \left( \sum_{i=0}^{\infty} t(i)\,x^i \right) \left( \sum_{j=0}^{\infty} t(j)\,x^j \right)
$$

The multiplication of the $2$ series — **Cauchy's formula**:

$$
\left( \sum_{n=0}^{\infty} a_n \right)\left( \sum_{n=0}^{\infty} b_n \right) = \sum_{n=0}^{\infty} c_n,
\qquad
c_n = a_0 b_n + a_1 b_{n-1} + a_2 b_{n-2} + \ldots + a_n b_0,
$$

under the condition that $\sum_{n=0}^{\infty} a_n$ and $\sum_{n=0}^{\infty} b_n$ converge absolutely.

So (the inner sum is the result of multiplying out the coefficients $c_k$):

$$
\begin{aligned}
x F^2(x)
&= x \sum_{k=0}^{\infty} \Bigl( \underbrace{\sum_{i+j=k} t(i)\,t(j)}_{\text{this is, by } (*),\ t(k+1)} \Bigr) x^k
&& \left(c_k = t(k+1) \text{ from the recurrence}\right) \\
&= x \sum_{k=0}^{\infty} t(k+1)\, x^k \\
&= \sum_{k=0}^{\infty} t(k+1)\, x^{k+1}
&& \left(\text{we change the index: } k+1 = n;\ k=0 \Rightarrow n=1,\ k=\infty \Rightarrow n=+\infty\right) \\
&= \sum_{n=1}^{\infty} t(n)\, x^n + \underbrace{t(0)\,x^0 - t(0)\,x^0}_{\text{by } (*),\ t(0)=1} \\
&= \sum_{n=0}^{\infty} t(n)\, x^n - 1 \;=\; F(x) - 1.
\end{aligned}
$$

$$
\boxed{\ x F^2(x) = F(x) - 1\ }
$$

The quadratic equation $x F^2(x) - F(x) + 1 = 0$ (with respect to $F(x)$) gives $2$ solutions:

$$
F(x) = \frac{1 \pm \sqrt{1 - 4x}}{2x}
\qquad
\left(1 - 4x \geq 0,\ \ 1 \geq 4x,\ \ x \leq \tfrac14\right)
$$

Expand the square root using Newton's generalized binomial series

$$
(x+y)^{\alpha} = \sum_{k=0}^{\infty} \binom{\alpha}{k} x^{\alpha-k} y^k \quad (|y| < |x|),
\qquad
(1+y)^{\alpha} = \sum_{k=0}^{\infty} \binom{\alpha}{k} y^k \quad (|y| < 1):
$$

$$
G_1(x) = \sqrt{1-4x} = (1-4x)^{\alpha = \frac12} = \sum_{n=0}^{\infty} \binom{1/2}{n} (-4x)^n = \sum_{n=0}^{\infty} \binom{1/2}{n} (-4)^n x^n
$$

Let us now compute the expression $\binom{1/2}{n}(-4)^n$:

$$
\binom{1/2}{n} \cdot (-4)^n
= \frac{\overbrace{\tfrac12 \left(-\tfrac12\right)\left(-\tfrac32\right) \cdots \left(\tfrac12 - n + 1\right)}^{n \ \text{factors}}}{n!} \, (-4)^n
$$

> **Note:** $\displaystyle \binom{\alpha}{n} \stackrel{\text{df}}{=} \frac{\alpha(\alpha-1)(\alpha-2)\cdots(\alpha-n+1)}{n!}, \qquad \binom{\alpha}{0} \stackrel{\text{df}}{=} 1.$

Since the last factor is $\tfrac12 - n + 1 = \tfrac{3-2n}{2} = -\left(\tfrac{2n-3}{2}\right)$, and $(-4)^n = (-1)^n \cdot 2^n \cdot 2^n$:

$$
\begin{aligned}
\binom{1/2}{n}(-4)^n
&= \frac{\tfrac12 \overbrace{\left(-\tfrac12\right)\left(-\tfrac32\right)\cdots\left(\tfrac{3-2n}{2}\right)}^{n-1 \ \text{factors} \ (**)}}{n!} \; (-1)^n\, 2^n\, 2^n \\[4pt]
&= \frac{\overset{(**)}{(-1)^n \cdot (-1)^{n-1}}\; 1 \cdot 1 \cdot 3 \cdot \ldots \cdot (2n-3)}{n! \; 2^n} \; 2^n\, 2^n \\[4pt]
&= -\,\frac{1 \cdot 3 \cdot \ldots \cdot (2n-3) \cdot 2^n}{n!} \\[4pt]
&= -\,\frac{1 \cdot 3 \cdot \ldots \cdot (2n-3) \;\; 2^n\, (n-1)!}{n!\,(n-1)!} \\[4pt]
&= -\,\frac{\bigl[1 \cdot 3 \cdot \ldots \cdot (2n-3)\bigr] \; \overbrace{\bigl(1 \cdot 2 \cdot 3 \cdots (n-1)\bigr)}^{n-1 \ \text{factors}} \; \overbrace{2 \cdot 2 \cdots 2}^{n \ \text{factors}}}{n!\,(n-1)!} \\[4pt]
&= -\,\frac{\bigl[1 \cdot 3 \cdot \ldots \cdot (2n-3)\bigr]\,\bigl[2 \cdot 4 \cdot 6 \cdots (2n-2)\bigr] \cdot 2}{n!\,(n-1)!} \\[4pt]
&= -\,\frac{1 \cdot 2 \cdot 3 \cdot 4 \cdot 5 \cdot 6 \cdots (2n-3)(2n-2) \cdot 2}{n!\,(n-1)!} \\[4pt]
&= -\,\frac{(2n-2)! \; 2}{n\,(n-1)!\,(n-1)!} \qquad (n \neq 0) \\[4pt]
&= -\,\frac{2}{n} \binom{2n-2}{n-1}
\end{aligned}
$$

Indeed:

$$
\binom{2n-2}{n-1} = \frac{(2n-2)!}{(n-1)!\,\bigl[(2n-2)-(n-1)\bigr]!} = \frac{(2n-2)!}{(n-1)!\,(n-1)!}
$$

Hence:

$$
\begin{aligned}
G_1(x) = \sqrt{1-4x}
&= \sum_{n=1}^{\infty} -\frac{2}{n} \binom{2n-2}{n-1} x^n + 1 \\[2pt]
&= \left\{
\begin{array}{l}
\text{change of index:} \\
n = \tilde{n} + 1, \quad n = 1 \Rightarrow \tilde{n} = 0 \\
n - 1 = \tilde{n}, \quad n = \infty \Rightarrow \tilde{n} = \infty
\end{array}
\right\} \\[2pt]
&= \sum_{\tilde{n}=0}^{\infty} -\frac{2}{\tilde{n}+1} \binom{2\tilde{n}}{\tilde{n}} x^{\tilde{n}+1} + 1 \\[2pt]
&= \sum_{n=0}^{\infty} -\frac{2}{n+1} \binom{2n}{n} x^{n+1} + 1 \qquad \text{(we return to the old index)}
\end{aligned}
$$

Now recall the definition

$$
(\blacksquare) \quad \boxed{\ F(x) = \sum_{n=0}^{\infty} t(n)\,x^n\ }
$$

— this is the definition of $F(x)$; all $t(n) \geq 0$, so for every $x$ in the convergence interval $I$ we take $x > 0$, $x \in I$, and then

$$
(\bullet) \quad F(x) \geq 0,
\qquad
F(x) = \frac{1 \pm \sqrt{1-4x}}{2x} = \frac{1 \pm G(x)}{2x}.
$$

To preserve the condition $(\bullet)$ we take

$$
\begin{aligned}
F(x) &= \frac{1 - G(x)}{2x} = \frac{1}{2x} - \frac{G(x)}{2x} \\[4pt]
&= \cancel{\frac{1}{2x}} - \frac{\displaystyle\sum_{n=0}^{\infty} -\frac{2}{n+1}\binom{2n}{n} x^{n+1}}{2x} - \cancel{\frac{1}{2x}}
\end{aligned}
$$

$$
\boxed{\ F(x) = \sum_{n=0}^{\infty} \frac{1}{n+1} \binom{2n}{n} x^n\ }
$$

Comparing with $(\blacksquare)$ $F(x) = \sum_{n=0}^{\infty} t(n)\,x^n$:

$$
(\blacktriangle) \quad \boxed{\ t(n) = \frac{1}{n+1} \binom{2n}{n}\ }
$$

Which ends the proof. $\square$

The numbers $t(n)$ are called the **Catalan numbers** (Eugène Catalan, 1814–1894, considered this formula in the context of placing parentheses in the expression $x_1 x_2 x_3 \cdots x_n$).

## The first 11 Catalan numbers

$$
(\circ) \quad
\begin{cases}
t(1) = t(0) = 1, & t(2) = 2, \quad t(3) = 5, \\
t(4) = 14, & t(5) = 42, \quad t(6) = 132, \\
t(7) = 429, & t(8) = 1430, \quad t(9) = 4862, \\
t(10) = 16796.
\end{cases}
$$

The formula $(\blacktriangle)$ is practically impossible to guess by following only $(\circ)$ and not applying the above analysis.

It applies to permutations on a stack in computer science.

## The wheel of fortune

**Casino:**

- a roulette with a wheel of numbers numbered from $[1 \ldots 1000]$;
- we draw a number (each with the same probability); then:
  - a) if $\lfloor \sqrt[3]{n} \rfloor \mid n$ (i.e. $\lfloor \sqrt[3]{n} \rfloor$ divides $n$) — we receive $5$ PLN;
  - b) in the opposite case we lose $1$ PLN.

**Can one get rich in this game?**

The average win: let

$$
I_1 = \left\{ \text{numbers satisfying } \lfloor \sqrt[3]{n} \rfloor \mid n \right\},
\qquad
I_2 = \left\{ \text{numbers not satisfying } \lfloor \sqrt[3]{n} \rfloor \mid n \right\},
$$

$$
|I_1| = W, \qquad |I_2| = P, \qquad P = 1000 - W.
$$

$$
\begin{aligned}
\hat{S} &= \sum_{i \in I_1} P(i) \cdot 5 \; - \sum_{i \in I_2} P(i) \cdot 1 \\
&= \frac{1}{1000} \left( 5W - P \right) \\
&= \frac{5W - (1000 - W)}{1000} \\
&= \frac{6W - 1000}{1000} \;\geq\; 0 \quad (\text{when?}) \\
&\qquad\qquad W \geq \frac{1000}{6} = 167.
\end{aligned}
$$

So out of the $1000$ numbers in the set $I = I_1 \cup I_2$ there must be at least $167$ numbers satisfying the condition

$$
(*) \quad \lfloor \sqrt[3]{n} \rfloor \mid n
$$

**How can we count these numbers?**

## Iverson notation

We use the advantages of **Iverson notation** in determining the sum (when the logical condition is satisfied we have $1$; when it is not, or it makes no sense, we put $0$):

$$
\begin{aligned}
W &= \sum_{n=1}^{1000} \bigl[\, n \text{ is a winning number (cond. } * ) \,\bigr] \\[2pt]
&= \sum_{1 \le n \le 1000} \Bigl[\, \bigl\lfloor \sqrt[3]{n} \bigr\rfloor \mid n \,\Bigr] \\[2pt]
&= \sum_{k,\,n} \underbrace{\Bigl[\, k = \bigl\lfloor \sqrt[3]{n} \bigr\rfloor \,\Bigr]}_{\text{cond. 1}} \; \underbrace{\bigl[\, k \mid n \,\bigr]}_{\text{cond. 2}} \; \underbrace{\bigl[\, 1 \le n \le 1000 \,\bigr]}_{\text{cond. 3}}
\qquad \text{← now we do not have to watch the ranges of the indices} \\[2pt]
&= \sum_{k,\,m,\,n} \underbrace{\bigl[\, k^3 \le n < (k+1)^3 \,\bigr]}_{\text{cond. 1'}} \; \underbrace{\bigl[\, n = km \,\bigr]}_{\text{cond. 2'}} \; \underbrace{\bigl[\, 1 \le n \le 1000 \,\bigr]}_{\text{cond. 3}}
\end{aligned}
$$

(The weakened inequality $1 \le n \le 1000$ we strengthen to $1 \le n < 1000$; for $n = 1000$, it **is** a number satisfying $(*)$, since $\lfloor\sqrt[3]{1000}\rfloor = 10 \mid 1000$.)

$$
\begin{aligned}
W &= 1 + \sum_{k,\,m,\,n} \bigl[\, k^3 \le n < (k+1)^3 \,\bigr] \bigl[\, n = km \,\bigr] \bigl[\, 1 \le n < 1000 \,\bigr] \\[2pt]
&= 1 + \sum_{k,\,m,\,n} \bigl[\, k^3 \le n < (k+1)^3 \,\bigr] \bigl[\, n = km \,\bigr] \bigl[\, 1 \le k < 10 \,\bigr] \\[2pt]
&= 1 + \sum_{k,\,m} \bigl[\, k^3 \le km < (k+1)^3 \,\bigr] \bigl[\, 1 \le k < 10 \,\bigr] \\[2pt]
&= 1 + \sum_{k,\,m} \Bigl[\, k^2 \le m < \tfrac{k^3 + 3k^2 + 3k + 1}{k} \,\Bigr] \bigl[\, 1 \le k < 10 \,\bigr] \\[2pt]
&= 1 + \sum_{k,\,m} \Bigl[\, m \in \bigl[\, k^2,\; k^2 + 3k + 3 + \tfrac1k \,\bigr) \Bigr] \bigl[\, 1 \le k < 10 \,\bigr]
\end{aligned}
$$

Observe: for an interval $[\alpha \,..\, \beta)$, $\beta \geq \alpha$, we have $\lfloor \beta \rfloor - \lfloor \alpha \rfloor + 1$ integers (when $\beta$ is not an integer). So:

$$
\begin{aligned}
W &= 1 + \sum_{1 \le k < 10} \Bigl( \bigl\lfloor k^2 + 3k + 3 + \tfrac1k \bigr\rfloor - \bigl\lfloor k^2 \bigr\rfloor + 1 \Bigr) \\[2pt]
&= 1 + \sum_{1 \le k < 10} \Bigl( \cancel{k^2} + 3k + 3 + 1 + \underbrace{\bigl\lfloor \tfrac1k \bigr\rfloor}_{\text{"0"}} - \cancel{k^2} \Bigr) \\[2pt]
&= 1 + \sum_{1 \le k < 10} \underbrace{(3k + 4)}_{a_k} \\[2pt]
&= 1 + \frac{\overset{a_1}{7} + \overset{a_9}{31}}{2} \cdot 9 \\[2pt]
&= \underline{\underline{172}}
\end{aligned}
$$

> Aside on $a_k$: $a_{k+1} - a_k = 3k + 3 + 4 - 3k - 4 = 3 = r$, an arithmetic sequence, so $S_n = \frac{a_1 + a_n}{2} \cdot n$.

> Small check for $k = 1$: here $\beta = 1 + 3 + 3 + 1 = 8$ is an integer and is excluded from the interval, and indeed $m \in \{1, \ldots, 7\}$ gives $7 = 3 \cdot 1 + 4$ numbers, so the term $3k + 4$ is correct for every $k$.

The number of numbers $\in \mathbb{N}$ satisfying $(*)$ in the set $[1, 2, \ldots, 1000]$ is

$$
172 > 167 \quad (\text{on average this many must occur in order to win}).
$$

$$
\frac{6 \cdot 172 - 1000}{1000} = 3.2 \ \text{gr} \quad (\text{we win on average}).
$$

After $100$ games we expect to win $3.2$ PLN.

## End of the lecture: 8 coins in 2 weighings

$8$ coins ($2$ weighings only): how to decide which one is heavier (there is exactly one heavier coin)?

### Method 1 (also good for 9 coins!)

We form two triples of coins:

![[lec14b_p14_two-triples.svg|380]]

We weigh any $2$ triples of coins:

- a) if they are of the same mass, then among the $2$ remaining coins there is $1$ heavier. One additional weighing decides which one is heavier.
- b) if one triple is heavier, we choose it. We pick any $2$ coins from this triple and weigh them:
  - if they weigh the same — the $3$rd, rejected coin is the heavier one;
  - if they differ — this $2$nd weighing itself decides which coin is heavier.

### Method 2

**a)** We weigh a triple against a triple; if they are equal, the $2$nd weighing (of the two remaining coins) decides:

![[lec14b_p15_method2-first-weighing.svg|560]]

**b)** If one triple is heavier — we choose that triple (the two leftover coins are then known to be of equal weight, the same as the other genuine coins). We split the chosen triple (arbitrarily): one of its coins we **swap** for a leftover coin ($Z$ — the swapped-in one), next to another we **add** the second leftover coin ($D$ — the added one), and the third coin is set aside. (We keep watching on which of the two pans the swapped and the added coin lie.)

![[lec14b_p15_method2-split.svg|600]]

**c)** We weigh the pan with $Z$ against the pan with $D$; if there is **equilibrium**, then the swapped-out, set-aside coin is the heavier one:

![[lec14b_p15_method2-second-weighing.svg|360]]

If instead:

- (i) the pan with the swapped-in coin $Z$ sinks — the triple's coin lying on that pan is the heavier one;
- (ii) the pan with the added coin $D$ sinks — the triple's coin lying on that pan is the heavier one.

![[lec14b_p16_method2-outcomes.svg|480]]

**Moral:** discarded information also contains information.

## Example: an electric ladder circuit

Consider the electric circuit below. Question: $v_n = \,?$ for $0 \le n \le K$.

![[lec14b_p16_ladder-circuit.svg]]

1. **Kirchhoff's law** (at the node where the voltage is $v_{n+1}$):

$$
I = I_1 + I_2
$$

2. **Ohm's law**: the voltage drop $V_2 - V_1$ across a resistor $R$ gives the current (intensity) through the resistor:

$$
I = \frac{V_2 - V_1}{R}
$$

So:

$$
I_1 = \frac{v_{n+1} - 0}{1}, \qquad
I_2 = \frac{v_{n+1} - v_n}{3}, \qquad
I = \frac{v_{n+2} - v_{n+1}}{3},
$$

$$
(*) \quad \boxed{\ \frac{v_{n+2} - v_{n+1}}{3} = \frac{v_{n+1}}{1} + \frac{v_{n+1} - v_n}{3}\ }
\qquad 0 \le n \le K-2.
$$

From $(*)$ we have the recursive relation

$$
(**) \quad \boxed{\ v_{n+2} - 5 v_{n+1} + v_n = 0\ }
$$

— a homogeneous recurrence scheme of order $2$ (linear, constant coefficients).

Whatever $v_0$ is:

$$
\frac{v_1 - v_0}{3} = \frac{v_0 - 0}{1} \iff \boxed{\ v_1 = 4 v_0\ }
$$

The characteristic equation for $(**)$:

$$
r^2 - 5r + 1 = 0, \qquad
r_1 = \frac{5 + \sqrt{21}}{2}, \qquad
r_2 = \frac{5 - \sqrt{21}}{2},
$$

$$
v_n = C_1 r_1^n + C_2 r_2^n.
$$

$$
\begin{cases}
v_0 = C_1 + C_2 \\
v_1 = C_1 r_1 + C_2 r_2 \quad (= 4v_0)
\end{cases}
$$

$$
\Rightarrow \quad
\begin{aligned}
C_1 &= v_0\,(r_2 - 4)(r_2 - r_1)^{-1} \\
C_2 &= v_0\,(r_1 - 4)(r_1 - r_2)^{-1}
\end{aligned}
$$

$$
\Rightarrow \quad (\blacksquare\blacksquare) \quad
\boxed{\ v_n = v_0 \, \frac{r_2 - 4}{r_2 - r_1} \, r_1^n + v_0 \, \frac{r_1 - 4}{r_1 - r_2} \, r_2^n\ }
$$

But $v_K = V$, so from $(\blacksquare\blacksquare)$:

$$
v_K = v_0 \, \frac{r_2 - 4}{r_2 - r_1} \, r_1^K + v_0 \, \frac{r_1 - 4}{r_1 - r_2} \, r_2^K.
$$

We can compute $v_0$ by substituting $V$ into $(\blacksquare\blacksquare)$:

$$
v_0 = \frac{(r_2 - r_1)\,V}{(r_2 - 4)\,r_1^K - (r_1 - 4)\,r_2^K}
$$

Hence, finally:

$$
\boxed{\ v_n = V \left[ \frac{(r_2 - 4)\,r_1^n - (r_1 - 4)\,r_2^n}{(r_2 - 4)\,r_1^K - (r_1 - 4)\,r_2^K} \right] \ }
$$

## A task for the future

- We fly on a plane (we are afraid that there is a bomb on board).
- We take a bomb with us (the chance that there are now $2$ bombs on board is smaller than the chance that there is one bomb).
- We have lowered the risk (but we are still afraid).
- So we take $2$ bombs with us (the chance that there are $3$ bombs on board is smaller than of $2$).
- Again we have lowered the risk.
- Proceeding analogously, at a fixed level of risk, we take $n$ bombs on board in order to feel safer.
- **Why is this reasoning flawed?**
