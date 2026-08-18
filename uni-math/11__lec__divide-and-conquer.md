# Discrete Mathematics — Lecture 11: Divide and Conquer

> Translated from the handwritten lecture notes `pdf/lect11DzielRzad-MatDysk.pdf`.
> (The original pages are headed "Algorithms & Complexity — Lecture 8".)

## Divide & Conquer algorithms

We address now **Divide & Conquer algorithms**.

**Example:** finding the maximal element of a list $L(2n)$ of length $2n$. Split the list into two halves $L_1(n)$ and $L_2(n)$:

![[lec11_p01_list-split.svg|340]]

- a) find $\max L_1(n)$
- b) find $\max L_2(n)$
- c) $\max L(2n) = \max\{m_1, m_2\}$

— and continue deep down (apply the same idea to each half).

### The Divide & Conquer paradigm

- the initial problem is **divided** into $2$ (or more) subparts;
- each subpart is **solved separately** (we can use here one processor, or we can execute the above tasks in parallel with multiple processors);
- next we **merge** the results from the subparts;
- we **nest** this approach.

## The recurrence for the running time

Let $T(n)$ denote the time to execute a particular task (or the number of operations, or the size of consumed memory) for a given algorithm, where $n$ denotes e.g. the length of a list, the size of the task to be performed, etc.

For $n = 2^k$ (and two subproblems):

$$
T(n) = \underbrace{T\!\left(\tfrac{n}{2}\right) + T\!\left(\tfrac{n}{2}\right)}_{\substack{\text{time to execute the algorithm} \\ \text{for each part separately}}} + \underbrace{f(n)}_{\substack{\text{time needed to} \\ \text{merge the results}}}
$$

$$
\Downarrow
$$

$$
(*) \qquad \boxed{\ T(n) = 2\,T\!\left(\tfrac{n}{2}\right) + f(n)\ }
$$

This is the recurrence-formula example for the **D&C** (**d**ivide & **c**onquer) paradigm.

We formulate now a theorem concerning $(*)$.

- It can be generalized to $T_n = 3\,T\!\left(\tfrac{n}{3}\right) + f(n)$ and more.

## Theorem (semi-explicit solution of the recurrence)

Let $S_n$ be the sequence satisfying:

$$
(*) \qquad
\begin{cases}
S_{2n} = 2 S_n + f(n) \\
S_1 \ \text{given}
\end{cases}
$$

Then for $n = 2^m$, $m \in \mathbb{N}$:

$$
(**) \qquad \boxed{\ S_{2^m = n} = 2^m \left[ S_1 + \frac{1}{2} \sum_{i=0}^{m-1} \frac{f(2^i)}{2^i} \right]\ }
$$

Note that still in $(**)$ we need to calculate the values $f(2^0), f(2^1), f(2^2), \ldots, f(2^{m-1})$, which in $S_{2^m}$ is also happening ($S_{2^{m-1}}, S_{2^{m-2}}, \ldots, S_{2^1}, S_{2^0}$). But $(**)$ can be called **semi-explicit**, as it helps for the special cases of

$$
f(n) = A + Bn
$$

(common in computer science) to find a full **explicit** form of $S_{2^m}$ and thus assess

$$
S_{2^m = n} = O(?)
$$

## Proof of $(**)$ by mathematical induction

We shall prove now $(**)$ by mathematical induction (running over $m$, where $n = 2^m$).

**I) (Initial step.)** By $(*)$, $S_1$ is given; for $m = 0$ in $(**)$ the upper limit is $m - 1 = -1$, so there is **no sum** (and $(**)$ reduces trivially to $S_1$).

So we need an induction from $m = 1$:

by $(*)$:

$$
S_2 = 2 \cdot S_1 + f(1)
$$

in $(**)$:

$$
S_2 = 2^1 \left[ S_1 + \frac{1}{2} \sum_{i=0}^{0} \frac{f(2^i)}{2^i} \right]
= 2 S_1 + \frac{1}{2} \cdot 2 \cdot \frac{f(2^0)}{2^0} = 2 S_1 + f(1)
$$

so OK.

**II)** Assume now that

$$
P(m): \quad S_{2^m} = 2^m \left[ S_1 + \frac{1}{2} \sum_{i=0}^{m-1} \frac{f(2^i)}{2^i} \right]
$$

holds. We ask whether

$$
P(m+1): \quad S_{2^{m+1}} = 2^{m+1} \left[ S_1 + \frac{1}{2} \sum_{i=0}^{m} \frac{f(2^i)}{2^i} \right]
$$

follows. Indeed:

$$
S_{2^{m+1}} = S_{2 \cdot 2^m} \stackrel{(*)}{=} 2 S_{2^m} + f(2^m)
\stackrel{P(m)}{=} 2 \left[ 2^m \left[ S_1 + \frac{1}{2} \sum_{i=0}^{m-1} \frac{f(2^i)}{2^i} \right] \right] + f(2^m)
$$

Thus

$$
\begin{aligned}
S_{2^{m+1}} &= 2^{m+1} \left[ S_1 + \frac{1}{2} \sum_{i=0}^{m-1} \frac{f(2^i)}{2^i} \right] + \frac{f(2^m)}{2 \cdot 2^m}\, 2^{m+1} \\[4pt]
&= 2^{m+1} \left[ S_1 + \frac{1}{2} \left( \sum_{i=0}^{m-1} \frac{f(2^i)}{2^i} + \frac{f(2^m)}{2^m} \right) \right] \\[4pt]
&= 2^{m+1} \left[ S_1 + \frac{1}{2} \sum_{i=0}^{m} \frac{f(2^i)}{2^i} \right] \\[4pt]
&= \underbrace{2^{m+1} \left[ S_1 + \frac{1}{2} \sum_{i=0}^{(m+1)-1} \frac{f(2^i)}{2^i} \right]}_{P(m+1)}
\end{aligned}
$$

Thus by the principle of mathematical induction $P(m)$ is true for all $m \geq 1$. $\square$

## The case $f(n) = A + Bn$

Let us apply $(**)$ to

$$
\text{c)} \quad f(n) = A + Bn,
$$

and then to:

- a) $A \neq 0$ and $B = 0$
- b) $A = 0$ and $B \neq 0$

— a) & b) (or c) are the common cases for algorithms in computer science.

For $f(n) = A + Bn$:

$$
\begin{aligned}
S_{2^m} &= 2^m \left[ S_1 + \frac{1}{2} \sum_{i=0}^{m-1} \frac{A + 2^i B}{2^i} \right] \\[4pt]
&= 2^m \left[ S_1 + \frac{1}{2} A \sum_{i=0}^{m-1} \frac{1}{2^i} + \frac{B}{2} \sum_{i=0}^{m-1} 1 \right] \\[4pt]
&= 2^m S_1 + \frac{1}{2}\, 2^m\, \frac{1 - \left(\frac{1}{2}\right)^m}{1 - \frac{1}{2}}\, A + 2^m \frac{B}{2} \underbrace{(1 + 1 + \ldots + 1)}_{m\text{-times}}
\end{aligned}
$$

(here we used the sum of the first $m$ terms of a geometric sequence: $\dfrac{1 - q^m}{1 - q} = 1 + q + \ldots + q^{m-1}$)

$$
= 2^m S_1 + 2^m \left( 1 - \left(\tfrac{1}{2}\right)^m \right) A + m\, 2^m\, \frac{B}{2}
$$

Thus:

$$
\boxed{\ S_{2^m} = 2^m S_1 + (2^m - 1) A + m\, 2^m\, \frac{B}{2}\ }
$$

Since $n = 2^m$, i.e. $m = \lg_2 n$:

$$
(*\!*\!*) \qquad \boxed{\ S_n = n\, S_1 + (n-1) A + n \lg_2 n\, \frac{B}{2}\ }
$$

We consider now the special subcase $B = 0$ and $A \neq 0$ (with the specific example).

## Example 1 — finding a maximal number in a list

**Divide & Conquer** for finding a maximal number in the list. Here

$$
f(n) = A + 0 \cdot B = A
$$

— $A$ is the time to compare $m_1$ and $m_2$, so here:

$$
\boxed{\ T(2n) = 2T(n) + A\ }
$$

By $(*\!*\!*)$ we have

$$
T(n = 2^m) = \underbrace{2^m}_{=\,n} S_1 + (\underbrace{2^m}_{=\,n} - 1) A
$$

where $S_1$ is the time needed to find a maximal element in a list of length $1$.

$$
\boxed{\ T(n) = n\, S_1 + (n - 1) A\ }
$$

— we check $n$ elements in lists of $1$ element, and we have $n - 1$ comparisons. Hence

$$
\boxed{\ T(n) = \Theta(n)\ }
$$

Note that if we use a **naive algorithm** to find a maximal element of the list:

![[lec11_p07_naive-scan-array.svg|300]]

- take $a_1$ and $a_2$, compare $a_1 \leq a_2$; take $\tilde{a}_1 = \max\{a_1, a_2\}$;
- take now $a_3$ and compare with $\tilde{a}_1$: $\tilde{a}_2 = \max\{\tilde{a}_1, a_3\}$;
- … continue …

then we have for this naive algorithm:

- a) $n - 1$ comparisons (each with cost $A$),
- b) $n$ checks of each element.

$$
\Downarrow
$$

So the algorithm to find $\max\{L(2n)\}$:

- a) find $\max\{L_1(n)\}$ — $m_1$
- b) find $\max\{L_2(n)\}$ — $m_2$
- c) $m = \max\{m_1, m_2\}$

**does not give acceleration** over the standard sequential algorithm. Only one can get some time-execution acceleration if $2$ processors are used.

## Example 2 — sorting a list with MERGE-SORT

Sorting the list by using the method **MERGE-SORT** (based on the Divide & Conquer paradigm). The list is split into two halves $L_1$ and $L_2$, each sorted in time $T(n)$:

![[lec11_p08_merge-sort-split.svg|360]]

$T(n)$: time needed to sort $L_1(n)$ and $L_2(n)$ with length $n$.

**Merging** the two sorted lists $L_1^{(1)} = (e_1, e_2, \ldots, e_n)$ and $L_2^{(1)} = (f_1, f_2, \ldots, f_n)$:

![[lec11_p09_two-lists-merge.svg|480]]

$$
g_1 = \min\{e_1, f_1\} = \min\{L_1^{(1)}, L_2^{(1)}\}
$$

$$
\begin{cases}
L_1^{(2)} := L_1^{(1)} \\
L_2^{(2)} := L_2^{(1)} \setminus \{f_1\}
\end{cases}
\quad \text{if } g_1 = f_1
$$

$$
\begin{cases}
L_1^{(2)} := L_1^{(1)} \setminus \{e_1\} \\
L_2^{(2)} := L_2^{(1)}
\end{cases}
\quad \text{if } g_1 = e_1
$$

$$
g_2 = \min\{L_1^{(2)}, L_2^{(2)}\}
$$

and so on.

We have here:

- $A_1 \cdot n$ — comparisons of $2$ elements,
- $A_2 \cdot n$ — cost of getting $n$ elements from $L_1$,
- $A_3 \cdot n$ — cost of getting $n$ elements from $L_2$.

Total cost:

$$
A_1 n + A_2 n + A_3 n = \underbrace{(A_1 + A_2 + A_3)}_{B} \cdot n
$$

Here $f(n) = Bn$. So for **merge-sort**:

$$
\boxed{\ T(2n) = 2T(n) + Bn\ }
$$

By $(*\!*\!*)$ we have:

$$
\boxed{\ T(n) = n\, T(1) + \frac{B}{2}\, n \lg_2 n\ }
$$

$$
\boxed{\ T(n) = \Theta(n \lg_2 n)\ } \qquad (A = 0,\ B \neq 0)
$$

But other algorithms to sort $L(n)$:

- **bubble sort** — $\Theta(n^2)$
- **insertion sort** — $\Theta(n^2)$

So even on a **single processor** merge-sort gives a smaller complexity to sort a list than a bubble sort or insertion sort.

$$
\begin{cases}
T(n) = 4\, T\!\left( \left\lfloor \tfrac{n}{2} \right\rfloor \right) \\
T(1) = 1
\end{cases}
\quad \text{(divide \& conquer type)}
\qquad \Rightarrow \qquad \text{we showed} \quad T(n) = \Theta(n^2)
$$
