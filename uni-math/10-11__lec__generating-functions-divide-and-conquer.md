# Discrete Mathematics — Lectures 10–11: Generating Functions and Divide-and-Conquer

> Translated from the handwritten lecture notes `pdf/lect10-11FunkTworz-DzielRzadz-MatDysk.pdf`.

## Recursions II — first-order linear recurrences with variable coefficients

$$
\boxed{\,a_n = c_n\, a_{n-1} + d_n\,} \qquad (*)
$$

E.g. **non-constant coefficients** (sequential search):

$$
c_n = \frac{n-1}{n}, \qquad d_n = 1
$$

- There is no general theory, but $(*)$ is a frequent case in computer science (that is, recurrences of order 1).

### Theorem 2

If $\{a_n\}$ satisfies

$$
\begin{cases}
a_n = c_n\, a_{n-1} + d_n, & n \geq 1 \\
a_0 = d_0
\end{cases}
$$

then

$$
a_n = \sum_{j=0}^{n} \left( \prod_{i=j+1}^{n} c_i \right) d_j \qquad P(n)
$$

**Proof** — we use induction with respect to $n$.

**Base** ($n = 0$, so $j = 0$):

$$
a_0 = \sum_{j=0}^{0} \left( \prod_{i=j+1}^{0} c_i \right) d_j
    = \underbrace{\left( \prod_{i=1}^{0} c_i \right)}_{\text{empty product, } =\,1} d_0
    = 1 \cdot d_0 = d_0 \quad \text{OK}
$$

(the empty multiplication — by convention we take it $= 1$).

**Step:** we assume $P(n)$ and show $P(n+1)$:

$$
\begin{aligned}
a_{n+1} &= c_{n+1}\, a_n + d_{n+1} \\
&= c_{n+1} \left[ \sum_{j=0}^{n} \left( \prod_{i=j+1}^{n} c_i \right) d_j \right] + d_{n+1} \\
&= \left[ \sum_{j=0}^{n} \left( \prod_{i=j+1}^{n+1} c_i \right) d_j \right] + 1 \cdot d_{n+1} \\
&= \left[ \sum_{j=0}^{n} \left( \prod_{i=j+1}^{n+1} c_i \right) d_j \right]
   + \underbrace{\left( \prod_{i=n+2}^{n+1} c_i \right)}_{1} d_{n+1} \\
a_{n+1} &= \sum_{j=0}^{n+1} \left( \prod_{i=j+1}^{n+1} c_i \right) d_j \qquad P(n+1). \ \square
\end{aligned}
$$

Unrolling the recurrence (to see the pattern):

$$
\begin{aligned}
a_0 &= a_0 \\
a_1 &= c_1 a_0 + d_1 \\
a_2 &= c_2 a_1 + d_2 = c_2 c_1 a_0 + c_2 d_1 + d_2 \\
a_3 &= c_3 a_2 + d_3 = c_3 (c_2 c_1 a_0 + c_2 d_1 + d_2) + d_3 \\
    &= c_3 c_2 c_1 a_0 + c_3 c_2 d_1 + c_3 d_2 + d_3 \\
&\ \ \vdots
\end{aligned}
$$

### Example

$$
\boxed{\,a_n = n\, a_{n-1} + n!\,}
$$

Here the coefficient $c_n = n$ is **not constant**, and the recurrence is **inhomogeneous** — even if $c_n \equiv \text{const}$ (here $c_n = n$), finding a particular solution is a problem.

Theorem 2 gives the solution: with $c_n = n$, $d_n = n!$ (so $c_i = i$, $d_j = j!$):

$$
a_n = \sum_{j=0}^{n} \left( \prod_{i=j+1}^{n} i \right) j!
    = \sum_{j=0}^{n} \frac{n!}{j!}\, j!
    = \sum_{j=0}^{n} n!
    = (n+1)\, n! = \underline{(n+1)!}
$$

## Divide-and-conquer algorithms

- the input problem is divided into 2 or more parts;
- these parts are solved separately;
- then we put the results together;
- algorithms of this kind are well suited to **parallel computation**.

**Example** — finding the largest number in a set of numbers $X$:

1. Split $X = X_1 \cup X_2$ so that $X_1 \cap X_2 = \emptyset$ and $\overline{\overline{X_1}} \approx \overline{\overline{X_2}}$ (halves of roughly equal size).
2. Find $X_i^{\max} \in X_i$.
3. Compare $X_1^{\max}$ with $X_2^{\max}$.

### The cost recurrence

Let $T(n)$ denote the "time" needed to execute a given algorithm. Assume $n = 2^k$:

$$
T(n) = T\!\left(\tfrac{n}{2}\right) + T\!\left(\tfrac{n}{2}\right) + F(n) \qquad (*)
$$

where:

- $n$ — the given size,
- $T\!\left(\tfrac{n}{2}\right)$ — the "time" to execute the algorithm on each part separately,
- $F(n)$ — the "time" needed to combine the two results.

$(*)$ is a linear inhomogeneous recurrence relation (but of unspecified order):

$$
\boxed{\,T_n = 2\, T_{n/2} + F(n)\,} \qquad (**)
$$

We will now formulate a theorem which solves $(**)$. It can be generalized to e.g.:

$$
T_n = b\, T_{n/2} + F(n), \qquad T_n = 3\, T_{n/3} + F(n)
$$

### Theorem 1

Let $S_n$ be a sequence satisfying a recurrence relation of the form:

$$
\begin{cases}
S_{2n} = 2 \cdot S_n + f(n), & n \in \mathbb{Z}_+ \\
S_1 \ \text{given}
\end{cases}
$$

Then:

**a)**

$$
S_{2^m} = 2^m \left[ S_1 + \frac{1}{2} \sum_{i=0}^{m-1} \frac{f(2^i)}{2^i} \right]
$$

**b)** In particular, when $\boxed{f(n) = A + Bn}$ ($A$ and $B$ constants), then

$$
\boxed{\; S_{2^m} = 2^m \cdot S_1 + \left(2^m - 1\right) A + \frac{B}{2}\, 2^m \cdot m \;}
$$

(the teacher's margin note: $\ \tfrac{1}{2}\, 2^m\, \dfrac{1 - \left(\tfrac12\right)^m}{1 - \tfrac12} = 2^m - 1$).

So if $2^m = n$:

$$
S_n = n\, S_1 + (n-1)\, A + \frac{B}{2}\, n \lg_2 n
$$

Margin special cases:

- $A \neq 0$ and $B = 0$: $\quad S_{2^m} = 2^m S_1 + A\,(2^m - 1)$
- $A = 0$ and $B \neq 0$: $\quad S_{2^m} = 2^m S_1 + \frac{B}{2}\, 2^m m$

**Proof.**

**I.** $P(1)$, i.e. $m = 1$ — we must start the induction from $m = 1$, because the formula uses $m - 1 \geq 0$. Take $S_1 = 1$ (it may be arbitrary). From the recurrence:

$$
S_2 = 2 \cdot S_1 + f(1) = 2 + f(1)
$$

From the formula:

$$
S_2 = 2 \left[ 1 + \frac{1}{2} \sum_{i=0}^{0} \frac{f(2^i)}{2^i} \right]
    = 2 + 2 \cdot \frac{f(1)}{2} = 2 + f(1) \quad \text{OK}
$$

**II.** We assume $P(m)$ and show $P(m+1)$. Assume

$$
S_{2^m} = 2^m \left[ S_1 + \frac{1}{2} \sum_{i=0}^{m-1} \frac{f(2^i)}{2^i} \right]
$$

Then:

$$
\begin{aligned}
S_{2^{m+1}} &= 2 S_{2^m} + f(2^m) \\
&\overset{P(m)}{=} 2^{m+1} \left[ S_1 + \frac{1}{2} \sum_{i=0}^{m-1} \frac{f(2^i)}{2^i} \right] + f(2^m) \\
&= 2^{m+1} \left[ S_1 + \frac{1}{2} \sum_{i=0}^{m-1} \frac{f(2^i)}{2^i} \right] + 2^{m+1} \cdot \frac{1}{2} \cdot \frac{f(2^m)}{2^m} \\
&= 2^{m+1} \left[ S_1 + \frac{1}{2} \sum_{i=0}^{m-1} \frac{f(2^i)}{2^i} + \frac{1}{2}\, \frac{f(2^m)}{2^m} \right] \\
&= 2^{m+1} \left[ S_1 + \frac{1}{2} \left( \sum_{i=0}^{m-1} \frac{f(2^i)}{2^i} + \frac{f(2^m)}{2^m} \right) \right] \\
&= \underbrace{2^{m+1} \left[ S_1 + \frac{1}{2} \sum_{i=0}^{m+1-1} \frac{f(2^i)}{2^i} \right]}_{P(m+1)}
\end{aligned}
$$

By the principle of mathematical induction, $P(m)$ is true for $m \geq 1$. $\square$

**The special case** $f(n) = A + Bn$, so $f(2^m) = A + B\, 2^m$:

$$
S_{2^m} = 2^m \left[ S_1 + \frac{1}{2} \sum_{i=0}^{m-1} \frac{A + 2^i B}{2^i} \right]
$$

But

$$
\begin{aligned}
\sum_{i=0}^{m-1} \frac{A + 2^i B}{2^i}
&= A \sum_{i=0}^{m-1} \frac{1}{2^i} + B \sum_{i=0}^{m-1} 1 \\
&= A\, \frac{1 - \left(\tfrac{1}{2}\right)^m}{1 - \tfrac{1}{2}} + B\, \underbrace{(1 + \ldots + 1)}_{m \text{ times}} \\
&= 2A \left( 1 - \frac{1}{2^m} \right) + m B
\end{aligned}
$$

Hence

$$
\begin{aligned}
S_{2^m} &= 2^m \left[ S_1 + \frac{1}{2} \cdot 2 A \left( 1 - \frac{1}{2^m} \right) + \frac{m}{2} B \right] \\
&= 2^m S_1 + A \left( 2^m - 1 \right) + 2^{m-1} m B
\end{aligned}
$$

$$
\boxed{\; S_{2^m} = S_1\, 2^m + A\,(2^m - 1) + 2^m m\, \frac{B}{2} \;}
\qquad \text{for } f(n) = A + Bn
$$

## Recursions III — applying the theorem

### Example: the largest number in a set via divide and conquer

The "divide and conquer" algorithm for finding the largest number in a set:

$$
T_{2n} = 2 T_n + A
$$

where the constant $A$ denotes the time to compare the largest numbers from the two halves.

So $B = 0$ in $f(n) = A + Bn$, and therefore

$$
\Rightarrow \quad T(2^m) = T_{2^m} = 2^m \cdot T(1) + \left(2^m - 1\right) A
$$

where $T(1)$ is the time needed to find the largest element in a one-element set. That is:

- $T(1)$ — the cost of checking a single element,
- $A$ — the cost of checking (comparing) two elements.

For $n = 2^m$:

$$
\boxed{\; T(n) = n\, T(1) + (n-1)\, A \;}
$$

(we check $n$ elements, and we perform $n-1$ comparisons).

Note that if we had used ordinary scanning of the elements one by one, remembering at each step the largest number found so far, we would have:

- **a)** $n - 1$ **comparisons** (each of time cost $A$) — at the 1st position there is no comparison;
- **b)** $n$ inspections of a single element (each of time cost $T(1)$).

Thus the "divide and conquer" method gives **no improvement** here! — unless we use several processors, and the gain comes from executing the subtasks in parallel.

### Example: sorting a list by divide and conquer (Merge-sort)

The list of length $\tilde n = 2n$ is split into two halves $L_1$, $L_2$; $T(n)$ is the time needed to sort a sublist of length $n$:

![[lec1011_p09_merge-sort-split.svg]]

How do we merge the sorted lists $L_1^{(1)} = [e_1, e_2, \ldots, e_n]$ and $L_2^{(1)} = [f_1, f_2, \ldots, f_n]$?

![[lec1011_p10_merge-lists.svg]]

$$
g_1 = \min\{e_1, f_1\} = \min\left\{ L_1^{(1)}, L_2^{(1)} \right\}
$$

$$
\left.
\begin{aligned}
L_1^{(2)} &:= L_1^{(1)} \\
L_2^{(2)} &:= L_2^{(1)} \setminus \{f_1\}
\end{aligned}
\right\} \ \text{if } g_1 = f_1
$$

$$
\left.
\begin{aligned}
L_1^{(2)} &:= L_1^{(1)} \setminus \{e_1\} \\
L_2^{(2)} &:= L_2^{(1)}
\end{aligned}
\right\} \ \text{if } g_1 = e_1
$$

$$
g_2 = \min\left\{ L_1^{(2)}, L_2^{(2)} \right\}, \quad \ldots \ \text{and so on.}
$$

We have:

- $A_1 \cdot n$ — comparisons of two elements,
- $A_2 \cdot n$ — the cost of getting to the $n$ elements in $L_1$,
- $A_3 \cdot n$ — likewise, in $L_2$.

$$
(A_1 + A_2 + A_3)\, n = B n
$$

Hence for Merge-sort:

$$
\boxed{\; T(2n) = 2\, T(n) + Bn \;}
$$

Applying Theorem 1 ($A \equiv 0$ — here that constant vanishes):

$$
T(2^m) = 2^m \cdot T(1) + \frac{B}{2}\, 2^m \cdot m
$$

If $2^m = n$, then $\lg_2 n = m$:

$$
\boxed{\; T(n) = n\, T(1) + \frac{B}{2}\, n \lg_2 n \;}
$$

We can see that for large $n$ (even if $B$ is small), since $\lg_2 n \geq 1$ and $n \lg_2 n \geq n$, the dominant cost is

$$
n \lg_2 n
$$

and it comes from **merging** the lists.

For comparison, other list-sorting algorithms, e.g. **Insertion Sort** or **Bubble Sort**, have a number of operations $\approx n^2$ ($> n \lg_2 n$!). So "divide and conquer" + recurrence analysis allows us to propose Merge-Sort and demonstrate its superiority — at least with respect to this criterion: time usage for large $n$!

### Between powers of two

Note that Theorem 1 says what happens only for

$$
\boxed{\, n = 2^m \,}
$$

There is a gap:

$$
\underbrace{2^m}_{n} < n + 1 < \ldots < 2n - 1 < \underbrace{2^{m+1}}_{2n}
$$

Often, however, it turns out that either

**(i)** $S_n \leq S_{n+1}$ (the sequence $S_n$ is monotone) — e.g. when $S_n$ represents the time of performing operations;

**(ii)** or only across the gap:

$$
S_{2^m = n} \leq S_{n+1}, \quad \ldots, \quad S_{2n-1} \leq S_{2^{m+1} = 2n}
$$

If we know that $S_n \leq S_{n+1}$ and $S_{2^m = n} \leq f(n)$, then

$$
\Rightarrow \quad f\big(\underbrace{2^{m-1}}_{n/2}\big) \;\leq\; S_{\tilde n} \;\leq\; f\big(\underbrace{n}_{2^m}\big)
\qquad \forall\, \tilde n \leq 2^m = n
$$

(i.e. for every $\tilde n$ between $n/2$ and $n$). This will help with **asymptotic analysis**!

### Example: estimating a sequence without a closed form

$$
(**) \quad
\begin{cases}
a_n = a_{n-2} + a_{n-3}, & n \geq 3 \\
a_0 = a_1 = a_2 = 1
\end{cases}
$$

Understanding the behaviour of a sequence (since we often want an estimate from above — e.g. a bound on the amount of time, number of operations, etc.) does not always require finding a closed form.

The first terms:

$$
1,\ 1,\ 1,\ 2,\ 2,\ 3,\ 4,\ 5,\ 7,\ 9,\ 12,\ 16,\ 21,\ 28,\ 37,\ 49,\ \ldots
$$

> The manuscript's list of first terms accidentally omits $a_{11} = 16$; it is restored above.

**Claim:**

$$
P(n): \quad \boxed{\, a_n \leq \left(\tfrac{4}{3}\right)^n \,}
$$

**I.** $P(0)$, $P(1)$, $P(2)$: $\ a_0 = a_1 = a_2 = 1 \leq \left(\tfrac43\right)^n$ for $n = 0, 1, 2$.

**II.** (strong induction) Assume $P(k)$: $a_k \leq \left(\tfrac43\right)^k$ for all $k < n$; we show $\left(\tfrac43\right)^n \geq a_n$:

$$
\begin{aligned}
a_n \overset{(**)}{=} a_{n-2} + a_{n-3}
&\underset{P(n-2),\ P(n-3)}{\leq} \left(\tfrac{4}{3}\right)^{n-2} + \left(\tfrac{4}{3}\right)^{n-3} \\
&= \left(\tfrac{4}{3}\right)^n \left[ \frac{1}{\left(\tfrac43\right)^2} + \frac{1}{\left(\tfrac43\right)^3} \right]
 = \left(\tfrac{4}{3}\right)^n \left[ \left(\tfrac34\right)^2 + \left(\tfrac34\right)^3 \right] \\
&= \left(\tfrac{4}{3}\right)^n \cdot \frac{63}{64} < \left(\tfrac{4}{3}\right)^n \qquad P(n). \ \square
\end{aligned}
$$

(We have some slack here: the factor $\tfrac{63}{64} < 1$.)

## Generating functions

To a sequence we associate a power series:

$$
\text{sequence } \{a_n\}_{n \in \mathbb{N}} \quad \longrightarrow \quad
f(x) = \sum_{n=0}^{\infty} a_n x^n = \lim_{n \to \infty} \sum_{k=0}^{n} a_k x^k
$$

(power series).

From analysis it is known that for every $\sum_{n=0}^{\infty} a_n x^n$ there exists $R \geq 0$ such that for all $x \in (-R, R)$ the series

$$
\sum_{n=0}^{\infty} a_n x^n \ \text{ is convergent.}
$$

$R$ — may be $0$ or $\infty$.

A formula for $R$: e.g. if

$$
\lim_{n} \left| \frac{a_{n+1}}{a_n} \right| \ \text{exists} \ = g
\qquad \text{or} \qquad
\lim_{n} \sqrt[n]{|a_n|} \ \text{exists} \ = g,
$$

then

$$
R = \frac{1}{g}
$$

**Remark:**

$$
f(x) = \sum_{n=0}^{\infty} a_n x^n, \ x \in (-R, R)
\quad \Rightarrow \quad
f'(x) = \sum_{n=1}^{\infty} a_n\, n\, x^{n-1}
$$

**Remark:** series may be added for $x \in \big(-\min(R_1, R_2),\ \min(R_1, R_2)\big)$, where

- $R_1$ — the radius of convergence for $\sum a_n x^n$,
- $R_2$ — the radius of convergence for $\sum b_n x^n$.

### Example: solving a recurrence with a generating function

$$
(*) \quad
\begin{cases}
a_n - 3 a_{n-1} + 2 a_{n-2} = 0, & n \geq 2 \\
a_0 = 0, \quad a_1 = 1
\end{cases}
$$

We form the generating function:

$$
\begin{aligned}
f(x) &= \sum_{k=0}^{\infty} a_k x^k = \big( a_0 + a_1 x + a_2 x^2 + \ldots \big) \\
-3x\, f(x) &= -3 \sum_{k=0}^{\infty} a_k x^{k+1} = \big( -3a_0 x - 3a_1 x^2 - 3a_2 x^3 - \ldots \big) \\
2x^2 f(x) &= 2 \sum_{k=0}^{\infty} a_k x^{k+2} = \big( 2a_0 x^2 + 2a_1 x^3 + 2a_2 x^4 + \ldots \big)
\end{aligned}
$$

Adding the three rows column by column:

$$
\begin{aligned}
f(x) - 3x f(x) + 2x^2 f(x)
= \ & a_0 + a_1 x - 3 a_0 x
+ x^2 \underbrace{(a_2 - 3a_1 + 2a_0)}_{=\,0 \text{ by } (*)}
+ x^3 \underbrace{(a_3 - 3a_2 + 2a_1)}_{=\,0 \text{ by } (*)} \\
& + x^4 \underbrace{(a_4 - 3a_3 + 2a_2)}_{=\,0 \text{ by } (*)} + \ldots
\end{aligned}
$$

So

$$
f(x) - 3x\, f(x) + 2x^2 f(x) = a_0 + a_1 x - 3 a_0 x
$$

and with $a_0 = 0$, $a_1 = 1$:

$$
\boxed{\; f(x) = \frac{x}{1 - 3x + 2x^2} \;}
$$

(Margin: $\Delta = 1$, $x_1 = \tfrac12$, $x_2 = 1$; $\ 1 - 3x + 2x^2 = 2\left(x - \tfrac12\right)(x-1) = (2x-1)(x-1) = (1-2x)(1-x)$.)

We decompose $f(x)$ into partial fractions:

$$
\frac{x}{(1-x)(1-2x)} = \frac{A}{1-x} + \frac{B}{1-2x}
$$

which gives

$$
f(x) = \frac{-1}{1-x} + \frac{1}{1-2x}
$$

But

$$
\frac{1}{1-x} = 1 + x + x^2 + \ldots + x^n + \ldots \qquad |x| < 1
$$

$$
-\frac{1}{1-x} = -\left( 1 + x + x^2 + \ldots + x^n + \ldots \right)
$$

Similarly

$$
\frac{1}{1-2x} = \left( 1 + 2x + 4x^2 + \ldots + (2x)^n + \ldots \right)
$$

Therefore:

$$
\begin{aligned}
f(x) &= -\left( 1 + x + x^2 + \ldots + x^n + \ldots \right)
      + \left( 1 + 2x + 4x^2 + \ldots + (2x)^n + \ldots \right) \\
&= x + \left(2^2 - 1\right) x^2 + \left(2^3 - 1\right) x^3 + \ldots + \left(2^n - 1\right) x^n + \ldots \\
&= \left(2^1 - 1\right) x^1 + \left(2^2 - 1\right) x^2 + \left(2^3 - 1\right) x^3 + \ldots + \left(2^n - 1\right) x^n + \ldots
\qquad \left(\text{since } \left(2^0 - 1\right) x^0 = 0\right) \\
&= \sum_{n=0}^{\infty} a_n x^n \quad \text{(as we assumed)}
\end{aligned}
$$

$$
\Rightarrow \quad a_0 = 0 \quad \text{and} \quad \boxed{\, a_n = 2^n - 1 \,}
$$

For **systems** of linear recurrences this method is more complicated, but it has generalizations — to which the linear methods do not carry over.

### Example: neutrons in a reactor (a system of recurrences)

In a reactor, after $n$ milliseconds let:

- $a_n$ — denote the number of **high-energy** neutrons,
- $b_n$ — denote the number of **low-energy** neutrons,

$$
a_0 = 1, \qquad b_0 = 0
$$

From the principles of nuclear physics:

- **a)** if a (high-energy) neutron reacts with a nucleus of fissile material, then after absorption (one millisecond later) **2 high-energy** neutrons and **1 low-energy** neutron are produced;
- **b)** if a low-energy neutron reacts with the fissile material, then **1 low-energy** neutron and **1 high-energy** neutron are produced.

$$
\begin{cases}
a_{n+1} = 2 a_n + b_n \\
b_{n+1} = a_n + b_n \\
a_0 = 1, \quad b_0 = 0
\end{cases}
$$

Introduce the generating functions

$$
f(x) = \sum_{n=0}^{\infty} a_n x^n, \qquad g(x) = \sum_{n=0}^{\infty} b_n x^n
$$

Multiply both equations by $x^{n+1}$:

$$
\begin{cases}
a_{n+1}\, x^{n+1} = 2 a_n x^{n+1} + b_n x^{n+1} \\
b_{n+1}\, x^{n+1} = a_n x^{n+1} + b_n x^{n+1}
\end{cases}
$$

and sum over $n$:

$$
\Rightarrow \quad
\begin{aligned}
\sum_{n=0}^{\infty} a_{n+1} x^{n+1} &= 2x \sum_{n=0}^{\infty} a_n x^n + x \sum_{n=0}^{\infty} b_n x^n \\
\sum_{n=0}^{\infty} b_{n+1} x^{n+1} &= x \sum_{n=0}^{\infty} a_n x^n + x \sum_{n=0}^{\infty} b_n x^n
\end{aligned}
$$

$$
\begin{cases}
\displaystyle \sum_{n=1}^{\infty} a_n x^n = 2x f(x) + x\, g(x) \\[2ex]
\displaystyle \sum_{n=1}^{\infty} b_n x^n = x f(x) + x\, g(x)
\end{cases}
$$

$$
\begin{cases}
\displaystyle \sum_{n=0}^{\infty} a_n x^n - a_0 = 2x f(x) + x\, g(x) \\[2ex]
\displaystyle \sum_{n=0}^{\infty} b_n x^n - b_0 = x f(x) + x\, g(x)
\end{cases}
$$

$$
(*) \quad
\begin{cases}
f(x) - \underbrace{a_0}_{1} = 2x f(x) + x\, g(x) \\
g(x) - \underbrace{b_0}_{0} = x f(x) + x\, g(x)
\end{cases}
$$

Solving the system $(*)$ we obtain:

$$
\boxed{\; f(x) = \frac{1-x}{x^2 - 3x + 1} \;}
\qquad
\boxed{\; g(x) = \frac{x}{x^2 - 3x + 1} \;}
$$

Factor the denominator:

$$
x^2 - 3x + 1 = (x - \gamma)(x - \delta), \qquad
\gamma = \frac{3 + \sqrt{5}}{2}, \quad \delta = \frac{3 - \sqrt{5}}{2}
$$

Decomposing into partial fractions:

$$
\begin{aligned}
f(x) = \frac{1-x}{x^2 - 3x + 1}
&= \left( \frac{5 + \sqrt{5}}{10} \right) \frac{1}{\gamma - x}
 + \left( \frac{5 - \sqrt{5}}{10} \right) \frac{1}{\delta - x} \\[1ex]
g(x) = \frac{x}{x^2 - 3x + 1}
&= \left( \frac{-5 - 3\sqrt{5}}{10} \right) \frac{1}{\gamma - x}
 + \left( \frac{-5 + 3\sqrt{5}}{10} \right) \frac{1}{\delta - x}
\end{aligned}
$$

Expanding each fraction as a geometric series ($\frac{1}{\gamma - x} = \frac1\gamma \sum_n \left(\frac x\gamma\right)^n$):

$$
\begin{aligned}
f(x) &= \sum_{n=0}^{\infty} \left[
\frac{5 + \sqrt{5}}{10}\, \frac{1}{\gamma} \left( \frac{x}{\gamma} \right)^n
+ \frac{5 - \sqrt{5}}{10}\, \frac{1}{\delta} \left( \frac{x}{\delta} \right)^n
\right] \\
&= \sum_{n=0}^{\infty} \left[
\frac{5 + \sqrt{5}}{10} \left( \frac{3 + \sqrt{5}}{2} \right)^{-(n+1)}
+ \frac{5 - \sqrt{5}}{10} \left( \frac{3 - \sqrt{5}}{2} \right)^{-(n+1)}
\right] x^n \\
&= \sum_{n=0}^{\infty} a_n x^n
\end{aligned}
$$

Hence (and similarly for $b_n$):

$$
\begin{cases}
\displaystyle a_n = \frac{5 + \sqrt{5}}{10} \left( \frac{3 + \sqrt{5}}{2} \right)^{-(n+1)}
 + \frac{5 - \sqrt{5}}{10} \left( \frac{3 - \sqrt{5}}{2} \right)^{-(n+1)} \\[2ex]
\displaystyle b_n = \frac{-5 - 3\sqrt{5}}{10} \left( \frac{3 + \sqrt{5}}{2} \right)^{-(n+1)}
 + \frac{-5 + 3\sqrt{5}}{10} \left( \frac{3 - \sqrt{5}}{2} \right)^{-(n+1)}
\end{cases}
$$

### Theorem — operations on generating functions

If

$$
\{a_n\}_{n=0}^{\infty} \longrightarrow f_1(x), \qquad
\{b_n\}_{n=0}^{\infty} \longrightarrow f_2(x)
$$

then:

**a)** (linearity)

$$
\{\alpha\, a_n + \beta\, b_n\}_{n=0}^{\infty} \longrightarrow \alpha f_1(x) + \beta f_2(x)
$$

**b)** (product / convolution)

$$
\{c_n\}_{n=0}^{\infty} \longrightarrow f_1(x) \cdot f_2(x),
\qquad \text{where} \quad c_n = \sum_{j=0}^{n} a_j\, b_{n-j}
$$

Formula **b)** will come in handy in the further part of the lecture.

**Remark:** a broader theory (of which $\{a_n\}_{n=0}^{\infty} \to f(x)$ is a special case) is the **$\mathcal{Z}$-transform**:

$$
\{\tilde a_n\}_{n=-\infty}^{+\infty} \ \overset{\mathcal{Z}}{\longrightarrow}\ \sum_{n=-\infty}^{+\infty} \frac{\tilde a_n}{z^n}
$$

Let us define

$$
\tilde a_n =
\begin{cases}
0 & n > 0 \\
a_{-n} & n \leq 0
\end{cases}
$$

Then

$$
\mathcal{Z}\{\tilde a_n\}
= \sum_{n=-\infty}^{0} \tilde a_n\, z^{-n} + \sum_{n=1}^{\infty} 0 \cdot \frac{1}{z^n}
= \sum_{n=-\infty}^{0} a_{-n}\, z^{-n}
\ \overset{[k = -n]}{=}\ \sum_{k=0}^{\infty} a_k z^k = f(z)
$$

> In the manuscript the two cases of $\tilde a_n$ are written with the opposite sign convention ($0$ for $n<0$, $a_{-n}$ for $n \geq 0$), but the final substitution $k = -n$ shows the sequence must be supported on $n \leq 0$, as written above.

### Generalized binomial coefficients

**Recall:** for $n, r \in \mathbb{Z}^+$, $n \geq r \geq 0$:

$$
\binom{n}{r} \overset{\text{df}}{=} \frac{n!}{r!\,(n-r)!} = \frac{n(n-1)\cdots(n-r+1)}{r!}
\qquad \boxed{0! = 1}
$$

If $n \in \mathbb{R}$, $r \in \mathbb{Z}^+$:

$$
\binom{n}{r} \overset{\text{df}}{=} \frac{n(n-1)\cdots(n-r+1)}{r!}
$$

For example, for $n \in \mathbb{Z}^+$:

$$
\begin{aligned}
\binom{-n}{r} &= \frac{(-n)(-n-1)\cdots(-n-r+1)}{r!} \\
&= \frac{(-1)^r\, n (n+1)(n+2)\cdots(n+r-1)}{r!} \\
&= \frac{(-1)^r\, (n+r-1)!}{(n-1)!\; r!} = (-1)^r \binom{n+r-1}{r}
\end{aligned}
$$

### Useful formulas for generating functions

(proofs by induction)

$$
(1+x)^n = \binom{n}{0} + \binom{n}{1} x + \binom{n}{2} x^2 + \ldots + \binom{n}{n} x^n
$$

$$
(1+x^m)^n = \binom{n}{0} + \binom{n}{1} x^m + \binom{n}{2} x^{2m} + \ldots + \binom{n}{n} x^{m \cdot n}
$$

$$
1 - x^{n+1} = (1-x)\left( 1 + x + x^2 + \ldots + x^n \right)
$$

$$
\frac{1}{1-x} = \sum_{i=0}^{\infty} x^i \qquad |x| < 1
$$

And the negative-exponent expansions:

$$
\begin{aligned}
\frac{1}{(1+x)^n} = (1+x)^{-n}
&= \binom{-n}{0} + \binom{-n}{1} x + \binom{-n}{2} x^2 + \ldots
= \boxed{\ \sum_{i=0}^{\infty} \binom{-n}{i} x^i\ } \\
&= 1 + (-1)\binom{n+1-1}{1} x^1 + (-1)^2 \binom{n+2-1}{2} x^2 + \ldots \\
&= \boxed{\ \sum_{i=0}^{\infty} (-1)^i \binom{n+i-1}{i}\, x^i\ }
\end{aligned}
$$

$$
\begin{aligned}
\frac{1}{(1-x)^n}
&= \binom{-n}{0} + \binom{-n}{1} (-x) + \binom{-n}{2} (-x)^2 + \ldots
= \boxed{\ \sum_{i=0}^{\infty} \binom{-n}{i} (-1)^i x^i\ } \\
&= 1 + (-1)\binom{n+1-1}{1}(-x) + (-1)^2 \binom{n+2-1}{2} (-x)^2 + \ldots \\
&= \boxed{\ \sum_{i=0}^{\infty} \binom{n+i-1}{i}\, x^i\ }
\end{aligned}
$$
