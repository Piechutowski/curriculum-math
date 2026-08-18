# Discrete Mathematics — Lecture 3: Mathematical Induction

> Translated from the handwritten lecture notes `pdf/lect3-indukmat-MatDysk.pdf`.

## The domino metaphor

The teacher's title doodle: an infinite row of standing dominoes, and the question whether they all end up knocked over:

![[lec03_p01_domino.svg]]

"We want to show that e.g. **all of them are knocked over**."

**Method 1: naive** — we inspect each one separately.

**Method 2: inductive**

- **a)** we check whether the **first one** is knocked over;
- **b)** we show that knocking over the **$k_0$-th** causes the **$(k_0+1)$-th** to be knocked over.

Then:

- a) allows us to "get started";
- b) is then applied implicitly:
  - since the 1st is knocked over (by a)), then by b) so is the 2nd,
  - since the 2nd, then by b) so is the 3rd,
  - and so on.

## The Principle of Mathematical Induction (MI)

**Principle of Mathematical Induction:** to prove a theorem $P(n)$ (indexed e.g. by $n \in \mathbb{N}$) for $n \geq n_0$, it suffices to show:

- **(i)** the **base step**: $P(n_0)$;
- **(ii)** the **inductive step**:

$$
\forall\, n \geq n_0: \quad \text{if } P(n), \text{ then } P(n+1)
$$

(the so-called **weak form** of MI).

> **Remark:** it often happens that a certain property $P(n)$ holds not for all $n \geq 1$ but only from some point on, $n \geq n_0$.

**Remark:** sometimes one has to use a **stronger version** of MI, replacing (ii) by:

$$
\text{if } \ \forall\, n \geq n_0: \ P(n_0),\ P(n_0+1),\ \ldots,\ P(n) \ \text{ then } \ P(n+1)
$$

## The strong form

In the strong form: sometimes it is not enough to assume the truth of the thesis for the **previous** index, but one has to accept its truth for **several predecessors**, e.g. $n$, $n-1$, $n-2$.

The Principle of Mathematical Induction is used in many applications — proving mathematical theorems:

- proving **equalities**
- proving **inequalities**
- proving certain **geometric properties**
- proving **properties of algorithms**
- proving **properties of numbers**
- proving **recurrence rules**
- proving **combinatorial rules**
- etc. (discrete mathematics)

## Example — sum of the natural numbers

$$
\boxed{\ \sum_{k=1}^{n} k = \frac{n(n+1)}{2}, \qquad n \geq 1\ }
$$

(the sum of the natural numbers).

**(i)** For $k = 1$:

$$
\sum_{k=1}^{1} k = 1 \quad \text{and} \quad \frac{1 \cdot (2)}{2} = 1, \qquad 1 = 1 \quad \text{o.k.}
$$

**(ii)** Take $k_0 \geq 1$. We assume:

$$
\sum_{k=1}^{k_0} k = \frac{k_0(k_0+1)}{2} \qquad (*)
$$

$\Downarrow$ question?

$$
\sum_{k=1}^{k_0+1} k = \frac{(k_0+1)(k_0+2)}{2} \qquad (**)
$$

We start from $(*)$:

$$
\sum_{k=1}^{k_0} k = \frac{k_0(k_0+1)}{2} \qquad \Big|\ + (k_0+1)
$$

$$
\sum_{k=1}^{k_0} k + (k_0+1) = \frac{k_0(k_0+1)}{2} + (k_0+1)
$$

$$
\sum_{k=1}^{k_0+1} k = (k_0+1)\left[\frac{k_0}{2} + 1\right] = (k_0+1)\,\frac{(k_0+2)}{2}
$$

(here we used $\sum_{i=1}^{n} a_i + a_{n+1} = \sum_{i=1}^{n+1} a_i$), which is exactly $(**)$:

$$
\sum_{k=1}^{k_0+1} k = \frac{(k_0+1)(k_0+2)}{2} \qquad \square
$$

## Example — sum of squares of the even numbers

$$
\boxed{\ \sum_{k=1}^{n} (2k)^2 = \frac{2}{3}\, n(n+1)(2n+1), \qquad n \geq 1\ } \qquad (1)
$$

(the formula for the sum of squares of the even numbers).

**(i)**

$$
\sum_{k=1}^{1} (2k)^2 = (2 \cdot 1)^2 = \underline{\underline{4}}
$$

$$
\frac{2}{3} \cdot 1 \cdot (1+1)(2 \cdot 1 + 1) = \frac{2}{3} \cdot 2 \cdot 3 = \underline{\underline{4}} \quad \text{o.k.}
$$

**(ii)** Assume

$$
\sum_{k=1}^{n} (2k)^2 = \frac{2}{3}\, n(n+1)(2n+1) \quad \text{true for } n \geq 1
$$

$\Downarrow$ question?

$$
\sum_{k=1}^{n+1} (2k)^2 = \frac{2}{3}(n+1)(n+2)\big(2(n+1)+1\big) \qquad (*)
$$

— is it also true?

We add $\big(2(n+1)\big)^2 = a_{n+1}$ to both sides of the assumption:

$$
\underbrace{\sum_{k=1}^{n} (2k)^2 + \underbrace{\big(2(n+1)\big)^2}_{a_{n+1}}}_{\sum_{k=1}^{n+1}(2k)^2} = \frac{2}{3}\, n(n+1)(2n+1) + \big(2(n+1)\big)^2
$$

$$
\sum_{k=1}^{n+1} (2k)^2 = \underbrace{\frac{2}{3}\, n(n+1)(2n+1) + \big(2(n+1)\big)^2}_{(\triangle)}
$$

Question: is $(\triangle)$ equal to

$$
\underbrace{\frac{2}{3}(n+1)(n+2)\big(2(n+1)+1\big)}_{(\triangle\triangle)} \ ?
$$

We transform $(\triangle)$:

$$
\begin{aligned}
(\triangle) &= \frac{2}{3}(n+1)\left[ n(2n+1) + \frac{3}{2} \cdot 4(n+1) \right] \\
&= \frac{2}{3}(n+1)\big[ n(2n+1) + 6n + 6 \big] \\
&= \frac{2}{3}(n+1)\big[ n\big(2(n+1)+1\big) - 2n + 6n + 6 \big] \\
&= \frac{2}{3}(n+1)\big[ n\big(2(n+1)+1\big) + \underline{4n + 6} \big]
\end{aligned}
$$

but $4n + 6 = 2\big(2(n+1)+1\big)$ (this is what we need — see $(*)$ above), so:

$$
= \frac{2}{3}(n+1)\Big[ n\big(2(n+1)+1\big) + 2\big(2(n+1)+1\big) \Big]
$$

we **factor** (pull out in front of the bracket), using $a(b+c) = ab + ac$:

$$
= \frac{2}{3}(n+1)\big(2(n+1)+1\big)\big[\, n+2 \,\big]
$$

— and this is exactly the right-hand side of (ii). By MI, the formula (1) is true for $n \geq 1$. $\square$

## Other equalities

**Other equalities** (to check):

$$
\begin{aligned}
1.\quad & \sum_{k=1}^{n} (3k-2) = \frac{n(3n-1)}{2} \\
2.\quad & \sum_{k=1}^{n} k(k+1) = \frac{n(n+1)(n+2)}{3} \\
3.\quad & \sum_{k=1}^{n} k(k-1) = \frac{(n-1)(n)(n+1)}{3} \\
4.\quad & \sum_{k=1}^{n} (-1)^k k^2 = (-1)^n\, \frac{n(n+1)}{2} \\
5.\quad & \sum_{k=1}^{n} k(k+1)(k+2) = \frac{n(n+1)(n+2)(n+3)}{4} \\
6.\quad & \sum_{k=1}^{n} k(k-1)(k-2) = \frac{(n-2)(n-1)\,n\,(n+1)}{4} \\
7.\quad & \sum_{k=1}^{n} \frac{2}{(k+2)k} = \frac{3}{2} - \frac{(2n+3)}{(n+1)(n+2)} \\
8.\quad & \sum_{k=1}^{n} \frac{1}{(2k-1)(2k+1)} = \frac{n}{2n+1} \\
9.\quad & \sum_{k=1}^{n} \frac{1}{(3k-2)(3k+1)} = \frac{n}{3n+1} \\
10.\quad & \sum_{k=0}^{n-1} \frac{1}{(n+k)(n+k+1)} = \frac{1}{2n} \\
11.\quad & \sum_{k=0}^{n-1} \frac{1}{(n-2k)(n-2k+2)} = \frac{n}{4-n^2} \qquad \text{for odd } n > 0
\end{aligned}
$$

## Inequalities

**Example:**

**(i)** For $x, y \geq 0$:

$$
\boxed{\ (x+y)^n = \sum_{k=0}^{n} \binom{n}{k} x^{n-k} y^k\ }
$$

— **Newton's binomial theorem**. Here (for $n \geq k \geq 0$):

$$
\binom{n}{k} = \frac{n!}{k!\,(n-k)!}, \qquad n! = 1 \cdot 2 \cdot \ldots \cdot n, \qquad 0! = 1
$$

The proof, e.g., goes by mathematical induction — one uses here:

$$
\binom{n}{k} = \binom{n-1}{k} + \binom{n-1}{k-1}
$$

**(ii)** **Bernoulli's inequality:**

$$
\boxed{\ (1+x)^n \geq 1 + nx\ } \qquad x \geq 0, \quad n \geq 0
$$

(here it holds even for $x \geq -1$).

**Step 0:**

$$
(1+x)^0 \geq 1 + 0 \cdot x \qquad \left(0^0 = 1 \ \text{by def.}\right)
$$

$$
1 \geq 1 \quad \text{o.k.}
$$

**Inductive step:** if $(1+x)^n \geq 1 + nx$, does it follow that

$$
(1+x)^{n+1} \stackrel{?}{\geq} 1 + (n+1)x \qquad (*)
$$

Since

$$
(1+x)^n \geq 1 + nx \quad \text{and} \quad (1+x) \geq 0 \quad (x \geq 0;\ \text{true also for } x \geq -1),
$$

we may multiply both sides by $(1+x)$:

$$
\underbrace{(1+x)^n \cdot (1+x)}_{(1+x)^{n+1} \ =\ \text{the left side of } (*)} \geq (1+nx)(1+x) = 1 + (n+1)x + \underbrace{n x^2}_{\text{but } \geq 0}
$$

$$
\underline{\underline{(1+x)^{n+1}}} \geq 1 + (n+1)x + nx^2 \geq \underline{\underline{1 + (n+1)x}}
$$

From transitivity: $(1+x)^{n+1} \geq 1 + (n+1)x$.

By mathematical induction, Bernoulli's inequality is true

$$
\forall\, n \geq 0 \ \text{ and } \ x \geq 0 \quad (\text{also for } x \geq -1). \qquad \square
$$

**Other inequalities** (examples):

$$
\begin{aligned}
1.\quad & 2^n > n^2 \quad \text{for } n > 4 \\
2.\quad & \sum_{k=1}^{n} \frac{1}{\sqrt{k}} > \sqrt{n} \\
3.\quad & 2^n > n+1
\end{aligned}
$$

## Divisibility

1. $4$ divides $3^n + 1$ — for $n > 0$ and $n$ **odd**
2. $5$ divides $2 \cdot 4^n + 3 \cdot 9^n$ — for $n \geq 0$
3. $7$ divides $2^{n+2} + 3^{2n+1}$ — for $n \geq 0$
4. $13$ divides $4^{2n+1} + 3^{n+2}$ — for $n \geq 0$

**Ad 1.**

$$
3^1 + 1 = 4, \qquad 4 \mid 4 \quad \text{o.k.}
$$

If $4 \mid 3^k + 1$ ($k$ odd), does it follow that $4 \mid 3^{k+2} + 1$? ($k+2$ is the next odd number, and it is odd too.)

$$
3^{k+2} + 1 = 3^k \cdot 9 + 1 = 3^k \cdot 9 + 9 - 8 = \underbrace{9\,(3^k + 1)}_{4 \text{ divides}} - \underbrace{8}_{4 \text{ divides}}
$$

so $4$ divides the whole expression — because if two numbers are divisible by $4$, then their difference is too.

So $4$ divides $3^n + 1$ (for $n > 0$ and $n = 2k-1$) by the principle of mathematical induction.

> **Remark:** only for $n$ odd. $\square$

## Recurrences

$$
a_{n+1} = 2a_n + a_{n-1}, \qquad n > 1
$$

- $a_1$ — an arbitrary integer,
- $a_2$ — an arbitrary integer.

Show that $a_n$ always gives an **integer**.

**(i)**

- $n = 1$: $a_1$ — yes, by assumption;
- $n = 2$: $a_2$ — yes, by assumption;

$a_1$ and $a_2$ are integers.

**(ii)** If $a_n$ and $a_{n-1}$ are integers

$$
\Downarrow \ ?
$$

is $a_{n+1}$ an integer?

**Remark:** here we use the strong version of MI! (More precisely, a stronger version.)

Yes — because if $a_n$ is an integer, then $2a_n \in \mathbb{Z}$; and if $2a_n \in \mathbb{Z}$ and $a_{n-1} \in \mathbb{Z}$, then

$$
2a_n + a_{n-1} \in \mathbb{Z}, \qquad \text{but this} \ = a_{n+1}
$$

## Multidimensional induction

Sometimes a certain property depends not on **one** natural number ($P(n)$) but on **more**, e.g. $P(n, m)$. Then we use **multidimensional induction**.

E.g. we **freeze one index** $n_0$ — the teacher's sketch of the lattice of pairs $(n,m)$ with the column above $n_0$ singled out:

![[lec03_p12_lattice-induction.svg]]

We check whether $P(n_0, m)$ is true $\forall\, m \geq 0$.

Then we freeze $m_0$ and prove similarly $P(n, m_0)$ $\forall\, n \geq 0$.

Remarks:

- **a)** the induction may start from $k_0 \geq 1$;
- **b)** it does not have to run over **all** $n \geq k_0$ (e.g. only over the odd numbers) — but the set must be infinite;
- **c)** one can also apply a **finite version**: $k_0, k_0+1, \ldots, n_0$.

## The stronger version of MI

**Stronger version:**

$$
\tilde{\text{I}}. \quad P(n_0),\ P(n_0+1),\ \ldots,\ P(n_0+k) \quad \text{true}
$$

$$
\tilde{\text{II}}. \quad \text{if } P(n),\ P(n+1),\ \ldots,\ P(n+k) \text{ true} \ \Downarrow \ P(n+k+1) \text{ true}
$$

Then, by mathematical induction:

$$
P(n) \ \text{true} \quad \forall\, n \geq n_0
$$

The example with the Fibonacci-type sequence (p. 11, the recurrence above) refers to this version (there $k = 1$).

**Strong version of mathematical induction:**

$$
\tilde{\tilde{\text{I}}}. \quad P(n_0) \ \text{true}
$$

$$
\tilde{\tilde{\text{II}}}. \quad \text{if } P(n) \text{ true } \forall\, n < \tilde{n} \ \Rightarrow \ P(\tilde{n}) \ \text{true}
$$

$$
\Downarrow
$$

$$
P(n) \ \text{true} \quad \forall\, n_0 \leq n
$$

We will give concrete examples for $\tilde{\text{I}}$ and $\tilde{\text{II}}$ later.

## Choosing $n_0$ — a final example

The choice of $n_0$ sometimes requires care. To finish, let us consider the example:

$$
\boxed{\ P(n): \quad 2^n > n+1\ } \ (*) \qquad n_0 = \,?
$$

For $n = 1$: $\ 2^1 > 2$ — **no**.

For $n = 2$: $\ 2^2 > 2+1$, i.e. $4 > 3$ — **yes**.

If it also checks out for $n_0 = 3, 4$, etc., then $P(n)$ is true $\Rightarrow$ we take $\underline{n_0 = 2}$.

**II.** $P(n) \stackrel{?}{\Rightarrow} P(n+1)$:

$$
2^n > n+1 \ \stackrel{?}{\Rightarrow} \ 2^{n+1} > n+2
$$

Indeed: from the inequality

$$
2^n > n+1 \qquad \Big|\ \cdot\, 2
$$

$$
2^n \cdot 2 > 2(n+1)
$$

$$
2^{n+1} > 2n+2 = n+2+\underbrace{n}_{\geq\, 0} \ \geq\ n+2
$$

Hence, from the transitivity of the inequality:

$$
2^{n+1} > (n+1) + 1
$$

And this is $P(n+1)$. So $(*)$ is true $\forall\, n \geq 2$. $\square$
