# Discrete Mathematics — Exercises 2: Solutions

> Translated from the handwritten lecture notes `pdf/Ćwicz2-RozwiązaniaMatDysk.pdf`.

## Problem 1 — the family of intervals $A_i = \left[0, \frac{2}{i}\right)$

$$
A_i = \left[0, \frac{2}{i}\right)
$$

The first few sets of the family:

$$
\begin{aligned}
i = 1: &\quad [0, 2) \\
i = 2: &\quad [0, 1) \\
i = 3: &\quad \left[0, \tfrac{2}{3}\right) \\
i = 4: &\quad \left[0, \tfrac{1}{2}\right) \\
i = 5: &\quad \left[0, \tfrac{2}{5}\right)
\end{aligned}
$$

The intervals $A_1, \ldots, A_5$ drawn on number lines (filled dot = endpoint included, open circle = endpoint excluded):

![[sol02_p01_intervals-a1-a5.svg]]

Notice that

$$
A_5 \subseteq A_4 \subseteq A_3 \subseteq A_2 \subseteq A_1
$$

$$
\Rightarrow \quad \bigcup_{i=1}^{5} A_i = A_1 = [0, 2)
\qquad
\bigcap_{i=1}^{5} A_i = A_5 = \left[0, \tfrac{2}{5}\right)
$$

For a general index $i = n$ the set $A_n = \left[0, \frac{2}{n}\right)$ looks like this:

![[sol02_p01_interval-an.svg|340]]

$$
A_1 \supseteq A_2 \supseteq A_3 \supseteq \ldots \supseteq A_n \supseteq \ldots
$$

$$
\bigcup_{i=1}^{\infty} A_i = A_1 = [0, 2)
$$

$$
\bigcap_{i=1}^{\infty} A_i = \{0\}
\qquad
\text{— one can see that the only common element is } \{0\}.
$$

### $\left(\bigcup A_i\right) \cap B$ for $B = \left(-1, \tfrac{1}{2}\right)$

$$
\left( \bigcup_{i=1}^{\infty} A_i \right) \cap B = [0, 2) \cap \left(-1, \tfrac{1}{2}\right) = \left[0, \tfrac{1}{2}\right)
$$

Or alternatively:

$$
\left( \bigcup_{i=1}^{\infty} A_i \right) \cap B
\stackrel{*}{=}
\bigcup_{i=1}^{\infty} (A_i \cap B) = \bigcup_{i=1}^{\infty} \ldots
\qquad \text{← here one must split the union}
$$

where $(*)$ is the distributive law:

$$
(A \cup B) \cap C \stackrel{*}{=} (A \cap C) \cup (B \cap C)
$$

Splitting the union at $i = 5$ (for $i \geq 5$ we have $\tfrac{2}{i} \leq \tfrac{2}{5} < \tfrac{1}{2}$, so $A_i \cap B = A_i$):

$$
\begin{aligned}
&= \left( \bigcup_{i=1}^{4} (A_i \cap B) \right) \cup \left( \bigcup_{i=5}^{\infty} (A_i \cap B) \right) \\
&= (A_1 \cap B) \cup (A_2 \cap B) \cup (A_3 \cap B) \cup (A_4 \cap B) \cup \bigcup_{i=5}^{\infty} \left[0, \tfrac{2}{i}\right) \\
&= \underbrace{\left( \left[0, \tfrac{1}{2}\right) \cup \left[0, \tfrac{1}{2}\right) \cup \left[0, \tfrac{1}{2}\right) \cup \left[0, \tfrac{1}{2}\right) \right)}_{\left[0, \tfrac{1}{2}\right)} \cup \left[0, \tfrac{2}{5}\right) \\
&= \left[0, \tfrac{1}{2}\right).
\end{aligned}
$$

We get the same result.

### $\left(\bigcap A_i\right) \cup B$

$$
\left( \bigcap_{i=1}^{\infty} A_i \right) \cup B = \{0\} \cup \left(-1, \tfrac{1}{2}\right) = \left(-1, \tfrac{1}{2}\right)
$$

One can also proceed similarly to the above:

$$
\left( \bigcap_{i=1}^{\infty} A_i \right) \cup B = \bigcap_{i=1}^{\infty} (A_i \cup B)
$$

To be checked at home.

## Problem 2 — proofs by mathematical induction

### a) Sum of the first $n$ odd numbers

$$
\sum_{i=1}^{n} (2i - 1) = \underbrace{1 + 3 + \ldots + (2n - 1)}_{P(n)} = n^2
$$

**Mathematical induction.**

**I.** $n = 1$: $\quad P(1): \ 1 \stackrel{?}{=} 2 \cdot 1 - 1$

$$
1 = 1 \quad \text{yes}
$$

**II.** If $P(n)$ is true $\stackrel{?}{\Rightarrow}$ $P(n+1)$ is true:

$$
\sum_{i=1}^{n} (2i - 1) = n^2
\quad \stackrel{?}{\Rightarrow} \quad
\sum_{i=1}^{n+1} (2i - 1) = (n+1)^2
$$

Assume $P(n): \ \sum_{i=1}^{n} (2i-1) = n^2$ and add the next term, $1 + 2(n+1) - 1$:

$$
\underbrace{\sum_{i=1}^{n} (2i-1) + 2(n+1) - 1}_{\sum_{i=1}^{n+1} (2i-1)} = n^2 + 2(n+1) - 1
$$

$$
\begin{aligned}
\sum_{i=1}^{n+1} (2i-1) &= n^2 + 2n + 2 - 1 \\
\sum_{i=1}^{n+1} (2i-1) &= n^2 + 2n + 1 \\
\sum_{i=1}^{n+1} (2i-1) &= (n+1)^2
\end{aligned}
$$

i.e. $P(n+1)$ is also true, provided $P(n)$ is true.

By mathematical induction, $P(n)$ is true for $n \geq 1$. $\square$

### b) $2^n > (n+1)^2$ for $n \geq n_0 = \ ?$

$$
P(n): \quad 2^n > (n+1)^2, \qquad n \geq n_0 = \ ?
$$

We check from which $n_0$ the statement $P(n)$ has a chance of being true:

$$
\begin{aligned}
2^1 &> (1+1)^2 \quad \text{no} \\
2^2 &> (2+1)^2 \quad \text{no} \\
2^3 &> (3+1)^2 \quad \text{no} \\
2^4 &> (4+1)^2 \quad \text{no} \\
2^5 &> (5+1)^2 \quad \text{no} \\
64 = 2^6 &> (6+1)^2 = 49 \quad \text{yes}
\end{aligned}
$$

**I.** $n_0 = 6$: $\quad P(6): \ 2^6 > (6+1)^2$ — yes.

**II.** If $P(n)$ is true $\Rightarrow$ $P(n+1)$ is true, for $n \geq 6$:

$$
\text{if } 2^n > (n+1)^2 \ \Rightarrow \ 2^{n+1} > ((n+1)+1)^2
$$

Take $P(n)$ and multiply both sides by $2$:

$$
P(n): \quad 2^n > (n+1)^2 \quad \Big|\ \cdot 2
$$

$$
\begin{aligned}
2^n \cdot 2 &> 2(n+1)^2 \\
2^{n+1} &> 2(n+1)^2 \stackrel{?}{>} (n+2)^2 \quad \text{— does this hold for } n \geq 6?
\end{aligned}
$$

Check the inequality $(\bullet)$:

$$
\begin{aligned}
2(n+1)^2 &\stackrel{(\bullet)}{>} (n+2)^2 \\
2(n^2 + 2n + 1) &> n^2 + 4n + 4 \\
2n^2 + 4n + 2 &> n^2 + 4n + 4 \\
n^2 - 2 &> 0 \quad \text{— is this true (at least for } n \geq 6\text{)?}
\end{aligned}
$$

Consider $f(x) = x^2 - 2$ (a quadratic $ax^2 + bx + c = 0$):

$$
\Delta = b^2 - 4ac = 0^2 - 4 \cdot 1 \cdot (-2) = 8
$$

$$
x_1 = \frac{0 - \sqrt{8}}{2} = \frac{-2\sqrt{2}}{2} = -\sqrt{2}
\qquad
x_2 = \frac{0 + \sqrt{8}}{2} = \frac{2\sqrt{2}}{2} = \sqrt{2}
$$

The parabola $x^2 - 2$ with its roots $\pm\sqrt{2}$:

![[sol02_p04_parabola-sqrt2.svg|220]]

$$
x \in (-\infty, -\sqrt{2}) \cup (\sqrt{2}, +\infty)
\quad \Rightarrow \quad
x^2 - 2 > 0
$$

but

$$
n \in \{6, 7, 8, 9, \ldots\} \subseteq (\sqrt{2}, +\infty)
$$

hence $n^2 - 2 > 0$ for $n \in \{6, 7, 8, 9, \ldots\}$

$$
\Rightarrow \quad 2(n+1)^2 > (n+2)^2
\quad \Rightarrow \quad
2^{n+1} > (n+2)^2,
$$

since $2^{n+1} > 2(n+1)^2$.

By mathematical induction, $P(n)$ is true for $n \geq 6$. $\square$

### c) $a_n \leq 3^n$ for $a_0 = 1$, $a_1 = 2$, $a_2 = 3$, $a_n = a_{n-1} + a_{n-2} + a_{n-3}$

**I.**

$$
\begin{aligned}
P(0)&: \quad 1 \leq 3^0 = 1 \quad \text{yes} \\
P(1)&: \quad 2 \leq 3^1 \quad \text{yes} \\
P(2)&: \quad 3 \leq 3^2 \quad \text{yes}
\end{aligned}
$$

**II.**

$$
\left.
\begin{aligned}
P(n-1)&: \quad a_{n-1} \leq 3^{n-1} \\
P(n-2)&: \quad a_{n-2} \leq 3^{n-2} \\
P(n-3)&: \quad a_{n-3} \leq 3^{n-3}
\end{aligned}
\right\}
\quad \stackrel{?}{\Rightarrow} \quad
P(n): \ a_n \leq 3^n
$$

But from the recurrence relation:

$$
a_n = a_{n-1} + a_{n-2} + a_{n-3} \leq 3^{n-1} + 3^{n-2} + 3^{n-3}
$$

$$
a_n \leq 3^{n-3} \underbrace{\left( 3^2 + 3^1 + 1 \right)}_{13} = 3^{n-3} \cdot 13
$$

but $13 \leq 3^3 = 27$, so

$$
3^{n-3} \cdot 13 \leq 3^{n-3} \cdot 27 = 3^{n-3} \cdot 3^3 = 3^n
$$

Therefore

$$
a_n \leq 3^n,
$$

which is $P(n)$.

By the **generalized (strong) principle of mathematical induction**:

$$
a_n \leq 3^n \quad \forall n \in \{0, 1, 2, \ldots\}. \qquad \square
$$

### d) $5 \mid 2 \cdot 4^n + 3 \cdot 9^n$, $\ n \geq 0$

$$
P(n): \quad \underbrace{\left( 2 \cdot 4^n + 3 \cdot 9^n \right)}_{\text{is this expression divisible by } 5?} \Big/ \ 5, \qquad n \geq 0
$$

**I.** $P(0)$:

$$
2 \cdot 4^0 + 3 \cdot 9^0 = 2 \cdot 1 + 3 \cdot 1 = 5 \quad \text{yes — divisible by } 5
$$

**II.**

$$
P(n): \ 5 \mid 2 \cdot 4^n + 3 \cdot 9^n
\quad \stackrel{?}{\Rightarrow} \quad
P(n+1): \ 5 \mid 2 \cdot 4^{n+1} + 3 \cdot 9^{n+1}
$$

$$
\begin{aligned}
2 \cdot 4^{n+1} + 3 \cdot 9^{n+1}
&= \overbrace{2 \cdot 4}^{8} \cdot 4^n + \overbrace{3 \cdot 9}^{27} \cdot 9^n \\
&= 2 \cdot 4^n + 6 \cdot 4^n + 3 \cdot 9^n + \underbrace{24}_{15 + 9} \cdot 9^n \\
&= \underbrace{\left( 2 \cdot 4^n + 3 \cdot 9^n \right)}_{5 \text{ divides, since } P(n)}
 + \underbrace{3 \left( 2 \cdot 4^n + 3 \cdot 9^n \right)}_{\substack{5 \text{ divides, since } P(n) \\ \text{so this divides too}}}
 + \underbrace{15 \cdot 9^n}_{\substack{5 \text{ divides} \\ 15 \cdot 9^n}}
\end{aligned}
$$

A sum of $3$ terms each divisible by $5$ is itself divisible by $5$, which means that $P(n+1)$ is true.

By mathematical induction, $P(n)$ is true $\forall n \geq 0$ (natural). $\square$

## Problem 3 — from the implicit to the explicit form

$$
p_0 = 3, \quad p_1 = 7 \quad \text{and} \quad p_n = 3p_{n-1} - 2p_{n-2}
$$

$$
\Rightarrow \quad P(n): \quad \boxed{\ p_n = 2^{n+2} - 1\ }, \qquad n \geq 2
$$

**I.**

$$
p_2 = 3 \cdot p_1 - 2 p_0 = 3 \cdot 7 - 2 \cdot 3 = 15 \stackrel{?}{=} 2^{2+2} - 1 = 16 - 1 = \underline{\underline{15}}
$$

$P(2)$ — YES.

$$
p_3 = 3 p_2 - 2 \cdot p_1 = 3 \cdot 15 - 2 \cdot 7 = 31 \stackrel{?}{=} 2^{3+2} - 1 = 32 - 1 = \underline{\underline{31}}
$$

— YES.

**II.**

$$
\begin{aligned}
P(n-1)&: \quad p_{n-1} \stackrel{*}{=} 2^{(n-1)+2} - 1 = 2^{n+1} - 1 \\
P(n-2)&: \quad p_{n-2} \stackrel{\triangle}{=} 2^{(n-2)+2} - 1 = 2^{n} - 1
\end{aligned}
\qquad \stackrel{?}{\Downarrow} \qquad
P(n): \quad p_n = 2^{n+2} - 1
$$

From the recurrence relation:

$$
\begin{aligned}
p_n = 3 p_{n-1} - 2 p_{n-2}
&\stackrel{* \text{ and } \triangle}{=} 3 \cdot \left( 2^{n+1} - 1 \right) - 2 \left( 2^n - 1 \right) \\
&= 3 \cdot 2^{n+1} - 3 - 2 \cdot 2^n + 2 \\
&= 3 \cdot 2^{n+1} - 1 - 2^{n+1} \\
&= 2 \cdot 2^{n+1} - 1 \\
&= 2^{n+2} - 1
\end{aligned}
$$

$$
\Rightarrow \quad \boxed{\ p_n = 2^{n+2} - 1\ } \quad P(n)
$$

By mathematical induction:

$$
P(n): \quad 2^{n+2} - 1 = p_n \qquad (\bullet)
$$

Note that here one had to **guess** the formula $(\bullet)$, i.e. how to pass from the **implicit** (recursive) form to the **explicit** (closed) form (in order to then prove it by induction). We will show later how to do this!
