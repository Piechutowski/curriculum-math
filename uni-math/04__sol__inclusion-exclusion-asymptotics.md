# Discrete Mathematics — Exercises 4: Solutions

> Translated from the handwritten lecture notes `pdf/Ćwicz4-RozwiązaniaMatDysk.pdf`.

## Problem 1 — numbers not divisible by 2, 3 or 5

$$
\overline{A \cup B \cup C} = \overline{A} + \overline{B} + \overline{C} - \overline{A \cap B} - \overline{A \cap C} - \overline{B \cap C} + \overline{A \cap B \cap C}
$$

($\overline{A}$ — the number of elements in $A$; this is the **inclusion–exclusion principle** for $n = 3$.)

$$
\overline{A} = \left\lfloor \frac{100}{2} \right\rfloor = \lfloor 50 \rfloor = 50
$$

— that is how many numbers from the range $\{1, 2, \ldots, 100\}$ are divisible by $2$.

$$
\overline{B} = \left\lfloor \frac{100}{3} \right\rfloor = \lfloor 33.(3) \rfloor = 33
$$

— that is how many numbers from the range $\{1, 2, \ldots, 100\}$ are divisible by $3$.

$$
\overline{C} = \left\lfloor \frac{100}{5} \right\rfloor = \lfloor 20 \rfloor = 20
$$

— that is how many numbers from the range $\{1, 2, \ldots, 100\}$ are divisible by $5$.

$\overline{A \cap C}$ — the number of numbers divisible by $2$ **and** $5$, i.e. divisible by $10$, among $\{1, \ldots, 100\}$:

$$
\overline{A \cap C} = \left\lfloor \frac{100}{10} \right\rfloor = \lfloor 10 \rfloor = 10
$$

$$
\overline{A \cap B} = \left\lfloor \frac{100}{6} \right\rfloor = \left\lfloor 16\tfrac{4}{6} \right\rfloor = 16
$$

— the number of numbers divisible by $2$ and $3$ among $\{1, \ldots, 100\}$.

$$
\overline{B \cap C} = \left\lfloor \frac{100}{15} \right\rfloor = \left\lfloor 6\tfrac{10}{15} \right\rfloor = 6
$$

— the number of numbers divisible by $3$ and $5$ from the range $\{1, \ldots, 100\}$.

$$
\overline{A \cap B \cap C} = \left\lfloor \frac{100}{30} \right\rfloor = \left\lfloor 3\tfrac{10}{30} \right\rfloor = 3
$$

— the number of numbers divisible by $2$, $3$ and $5$ from the range $\{1, 2, \ldots, 100\}$.

The number of those divisible by $2$ or $3$ or $5$ is $\overline{A \cup B \cup C}$; so the number of those which are **not** divisible by $2$ or $3$ or $5$ is:

$$
(*) \quad N - \overline{A \cup B \cup C} = 100 - \left[ 50 + 33 + 20 - 10 - 16 - 6 + 3 \right] = 100 - 74 = 26
$$

(where $N = 100$).

These numbers are:

$$
\begin{aligned}
&1, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 49, 53, \\
&59, 61, 67, 71, 73, 77, 79, 83, 89, 91 \text{ and } 97.
\end{aligned}
$$

But as the formula $(*)$ shows, they can be counted without finding them.

## Problem 2 — students playing football or basketball

**a)**

$$
\overline{A \cup B} = \overline{A} + \overline{B} - \overline{A \cap B}
$$

(the inclusion–exclusion principle for $n = 2$)

$$
\overline{A \cup B} = 4 + 5 - 2 = \underline{\underline{7}}
$$

— that many people play football or basketball.

**b)**

$$
30 - \overline{A \cup B} = 30 - 7 = 23
$$

— that many students play neither of these games.

## Problem 3 — socks and the pigeonhole principle

We use here the **pigeonhole principle** (Dirichlet's drawer principle):

- a sock is a **pigeon**,
- a color is a **pigeonhole**.

We want to have $2$ pigeons in a hole, so to be certain one must take **6 socks**!

## Problem 4 — the recurrence $a_n = 16\, a_{\lfloor n/2 \rfloor}$

$$
\begin{cases}
a_n = 16\, a_{\lfloor \frac{n}{2} \rfloor} \\
a_1 = 3
\end{cases}
$$

— the **implicit (recursive) form**. Let us look:

$$
\begin{aligned}
a_1 &= 3 & (k = 0) \\
a_2 &= 16 \cdot a_{\lfloor \frac{2}{2} \rfloor} = 16 \cdot a_{\lfloor 1 \rfloor} = 16 \cdot a_1 = 16 \cdot 3 & (k = 1) \\
a_4 &= 16 \cdot a_{\lfloor \frac{4}{2} \rfloor} = 16 \cdot a_{\lfloor 2 \rfloor} = 16 \cdot a_2 = 16^2 \cdot 3 & (k = 2)
\end{aligned}
$$

(so for $n = 2^k$: $\ a_{2^1} = 3 \cdot 16^1$, $\ a_{2^2} = 3 \cdot 16^2$).

We will show that:

$$
(\bullet) \quad
\begin{cases}
\text{(i)} & a_m = 3 \cdot 16^k \quad \longleftarrow \text{explicit (closed) form} \\
\text{(ii)} & \text{where } 2^k \leq m \quad \longleftarrow \text{for arbitrary } m \\
\text{(iii)} & \text{and } k \text{ is the maximal natural number such that } 2^k \leq m
\end{cases}
$$

**Mathematical induction:**

### I. Base case, $n = 1$

$$
a_1 = 3 \cdot 16^0, \qquad k = 0, \qquad 2^0 \leq 1
$$

and there cannot be a bigger $k$ (since $2^1 > 1$), so $(\bullet)$ is true.

### II. Inductive step

Assume that $(\bullet)$ is true for all $k < n+1$. We use here the **strong version of mathematical induction**. From the recurrence relation:

$$
(\diamond) \quad a_{n+1} = 16\, a_{\left\lfloor \frac{n+1}{2} \right\rfloor}
$$

Two cases: **a)** $n+1$ even, **b)** $n+1$ odd.

**ad a)** $\frac{n+1}{2}$ is a natural number, so

$$
\left\lfloor \frac{n+1}{2} \right\rfloor = \frac{n+1}{2}
$$

$$
(\blacktriangle) \quad a_{n+1} \overset{(\diamond)}{=} 16\, a_{\lfloor \frac{n+1}{2} \rfloor} = 16\, a_{\frac{n+1}{2}}, \qquad \text{but } \frac{n+1}{2} < n+1,
$$

so from the induction hypothesis we have:

$$
\begin{aligned}
\text{(i}'\text{)} \quad & a_{\frac{n+1}{2}} = 3 \cdot 16^{k'} \\
\text{(ii}'\text{)} \quad & 2^{k'} \leq \frac{n+1}{2} \\
\text{(iii}'\text{)} \quad & k' \text{ is the maximal natural number with property (ii}'\text{)}
\end{aligned}
$$

Substituting (i$'$) into $(\blacktriangle)$ we have:

$$
a_{n+1} = 16 \cdot 3 \cdot 16^{k'} = 3 \cdot 16^{k'+1} = 3 \cdot 16^{l'}, \qquad l' = k'+1 \quad \Rightarrow \quad \boxed{a_{n+1} = 3 \cdot 16^{l'}}
$$

— and this is (i).

Using (ii$'$):

$$
2^{l'} = 2^{k'+1} = 2 \cdot 2^{k'} \leq n+1 \quad \Rightarrow \quad \boxed{2^{l'} \leq n+1}
$$

— and this is (ii).

The question is whether $l' = k'+1$ is maximal such that

$$
2^{l'} \leq n+1.
$$

Suppose there exists another $\hat{l}$ such that

$$
\underline{\hat{l} > l' = k'+1} \quad (*)
$$

and $2^{\hat{l}} \leq n+1$. Then

$$
2 \cdot 2^{\hat{l}-1} \leq n+1
$$

$$
2^{\hat{l}-1} \leq \frac{n+1}{2}
$$

but $k'$ was maximal such that $2^{k'} \leq \frac{n+1}{2}$, hence

$$
\hat{l} - 1 \leq k' \quad \Rightarrow \quad \underline{\hat{l} \leq k'+1} \quad (**)
$$

And $(*)$ and $(**)$ give a contradiction. So $(*)$ is false $\Rightarrow$ $l' = k'+1$ is maximal such that $2^{l'} \leq n+1$ — and this is (iii).

**ad b)** Now $n+1$ is odd

$$
\Rightarrow \quad \frac{n+1}{2} \text{ does not divide without remainder:} \qquad \left\lfloor \frac{n+1}{2} \right\rfloor \neq \frac{n+1}{2}
$$

But then $n$ is even, and

$$
\left\lfloor \frac{n+1}{2} \right\rfloor = \left\lfloor \frac{n}{2} + \frac{1}{2} \right\rfloor = \frac{n}{2}
$$

(since $\frac{n}{2}$ is an integer and $\frac{1}{2} \in (0,1)$). From $(\diamond)$ (the recurrence relation):

$$
(\blacktriangle\blacktriangle) \quad a_{n+1} = 16\, a_{\lfloor \frac{n+1}{2} \rfloor} = 16 \cdot a_{\frac{n}{2}}
$$

Now, similarly, $\frac{n}{2} < n+1$, so one can use the induction hypothesis:

$$
\begin{aligned}
\text{(i}''\text{)} \quad & a_{\frac{n}{2}} = 3 \cdot 16^{k''} \\
\text{(ii}''\text{)} \quad & 2^{k''} \leq \frac{n}{2} \\
\text{(iii}''\text{)} \quad & k'' \text{ is maximal such that (ii}''\text{) is true}
\end{aligned}
$$

From $(\blacktriangle\blacktriangle)$ and (i$''$) we have:

$$
a_{n+1} = 16 \cdot 3 \cdot 16^{k''} = 3 \cdot 16^{k''+1} = \underline{3 \cdot 16^{l''}}, \qquad \text{where } l'' = k''+1.
$$

And this is exactly (i).

Since (from (ii$''$))

$$
2^{k''} \leq \frac{n}{2},
$$

we get

$$
2^{k''+1} \leq 2 \cdot \frac{n}{2} = n < n+1,
$$

and by transitivity:

$$
2^{k''+1} \leq n+1, \quad \text{i.e. } \underline{2^{l''} \leq n+1}
$$

— and this is (ii).

Suppose there exists $l > k''+1 = l''$ $\ (\bullet\bullet)$ such that $2^l \leq n+1$:

$$
\underbrace{2^l}_{\text{even}} \leq \underbrace{n+1}_{\text{odd}} \quad \Downarrow \quad 2^l \leq n
$$

$$
2^{l-1} \leq \frac{n}{2},
$$

but $k''$ was maximal such that $2^{k''} \leq \frac{n}{2}$ (by (iii$''$)), from which it follows that

$$
l - 1 \leq k'' \quad \Downarrow \quad l \leq k''+1 = l''
$$

So $l \leq l''$, while from $(\bullet\bullet)$ we have $l > l''$. We have a contradiction; hence $2^{l''} \leq n+1$ is satisfied by $l''$ as the maximal natural number — and this is (iii).

> **Note:** the manuscript writes this last step as "$k'' \leq l-1 \Rightarrow l'' = k''+1 \leq l$", which does not produce the contradiction; the maximality of $k''$ actually gives $l-1 \leq k'' \Rightarrow l \leq l''$, exactly as in the analogous step of case a). Corrected above.

### Conclusion

By Mathematical Induction, (i)–(iii) are true

$$
\Rightarrow \quad a_n = 3 \cdot 16^{k}, \qquad \text{where } 2^k \leq n \text{ and } k \text{ is maximal.}
$$

But

$$
16^k = (2^4)^k = 2^{4k} = (2^k) \cdot (2^k) \cdot (2^k) \cdot (2^k) \leq n \cdot n \cdot n \cdot n = n^4
$$

$$
\Rightarrow \quad 0 \leq a_n \leq 3 \cdot n^4, \qquad \text{so} \quad \boxed{a_n = O(n^4)}
$$

— complexity of degree $4$ (but it may be lower).

The scheme was:

$$
\text{implicit form} \ \longrightarrow \ \text{explicit form} \ \longrightarrow \ \text{determining the asymptotics}
$$

From the above considerations it does **not** follow that $a_n = \Theta(n^4)$ — perhaps it is lower?

## Problem 5 — asymptotics of two sequences

### (i) $a_n = 5n^4 + 6 \ln n$

A sketch comparing the growth of $x^4$ and $\ln x$ (with the sample points $n = 1, 2, 3$ marked on both curves):

![[sol04_p08_growth-comparison.svg]]

One sees that the dominating term is $n^4$:

$$
\Rightarrow \quad a_n = O(n^4)
$$

One can show that the sharp bound $\Theta(n^4)$ holds — because we have the inequality — a complexity of degree $4$.

### (ii) $b_n = n^2 - 2\binom{n}{2}$

Crudely: $n^2$ is $O(n^2)$ and $2\binom{n}{2}$ is $O(n^2)$, so

$$
O(n^2) + O(n^2) = O(n^2).
$$

But one can do better. Since

$$
\binom{n}{2} = \frac{n!}{2!\,(n-2)!} = \frac{n \cdot (n-1) \cdot (n-2)!}{2!\,(n-2)!} = \frac{n^2 - n}{2},
$$

we get

$$
b_n = n^2 - 2 \cdot \frac{n^2 - n}{2} = n \quad \Rightarrow \quad \boxed{b_n = \Theta(n)}
$$

Sharpness: it cannot be better. $\square$
