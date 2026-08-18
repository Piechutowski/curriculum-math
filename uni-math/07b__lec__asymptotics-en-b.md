# Discrete Mathematics — Lecture 7: Asymptotics (Part B)

> Translated from the handwritten lecture notes `pdf/lect7asymEN-BMatDysk.pdf`.

> The original notes carry the header "Algorithms & Complexity — Lecture 4"; they continue the discussion of asymptotic estimates of sequences arising in the analysis of algorithms.

## From implicit to explicit form — or straight to an estimate

Sometimes, if the counting (of memory consumption, of the number of basic operations, etc.) is given, one does **not** need to re-transform the recursive scheme from

$$
\text{implicit form} \;\Longrightarrow\; \text{explicit form}.
$$

**Example:**

$$
(*) \quad
\begin{cases}
a_n = a_{n-2} + a_{n-3}, & n \geq 3 \\
a_0 = a_1 = a_2 = 1
\end{cases}
$$

We search for some **upper bound** to "assess" the estimate of the rate of "explosion" of $a_n$. One can plot on a computer the pairs

$$
(n_0, a_{n_0}), \ (n_0+1, a_{n_0+1}), \ \ldots
$$

and try to compare them with plots of curves

$$
\beta \ln x, \quad \beta x^n, \quad \beta x^\alpha, \quad \beta \sqrt[d]{n}, \quad \beta a^n, \quad \text{etc.}
$$

This helps to formulate the hypothesis on the upper bound, which can then be inductively proved and also later improved.

## Inductive proof of an upper bound for $(*)$

In the case of $(*)$, $a_n$ takes the values

$$
1,\ 1,\ 1,\ 2,\ 2,\ 3,\ 4,\ 5,\ 7,\ 9,\ 12,\ 16,\ 21,\ 28,\ 37,\ 49,\ \ldots
$$

> The handwritten list skips the value $16$ ($a_{11} = a_9 + a_8 = 9 + 7 = 16$); it is restored here.

We set up the hypothesis (this is the hardest part here, as proving it is easier):

$$
\boxed{\ P(n): \quad 0 \leq a_n \leq \left(\tfrac{4}{3}\right)^{n} \ }
$$

We now prove $P(n)$ by using mathematical induction.

**I.** We verify $P(0)$, $P(1)$, $P(2)$:

$$
\begin{aligned}
P(0):&\quad 0 \leq a_0 = 1 \leq \left(\tfrac{4}{3}\right)^{0} = 1 \quad \text{yes} \\
P(1):&\quad 0 \leq a_1 = 1 \leq \left(\tfrac{4}{3}\right)^{1} = \tfrac{4}{3} \quad \text{yes} \\
P(2):&\quad 0 \leq a_2 = 1 \leq \left(\tfrac{4}{3}\right)^{2} = \tfrac{16}{9} \quad \text{yes}
\end{aligned}
$$

**II.** Assume now that

$$
(\Delta) \quad
\begin{cases}
P(n-1): & 0 \leq a_{n-1} \leq \left(\tfrac{4}{3}\right)^{n-1} \\[2pt]
P(n-2): & 0 \leq a_{n-2} \leq \left(\tfrac{4}{3}\right)^{n-2} \\[2pt]
P(n-3): & 0 \leq a_{n-3} \leq \left(\tfrac{4}{3}\right)^{n-3}
\end{cases}
$$

are all true. We want to show that from $(\Delta)$ we can infer:

$$
\boxed{\ P(n): \quad 0 \leq a_n \leq \left(\tfrac{4}{3}\right)^{n} \ }
$$

By the recursive formula

$$
a_n = a_{n-2} + a_{n-3},
$$

since by $(\Delta)$ we have $a_{n-2} \geq 0$ and $a_{n-3} \geq 0$, also $a_n \geq 0$ (by $(\Delta)$). Furthermore:

$$
\begin{aligned}
a_n = a_{n-2} + a_{n-3}
&\underset{\text{by }\Delta}{\leq} \left(\tfrac{4}{3}\right)^{n-2} + \left(\tfrac{4}{3}\right)^{n-3} \\[4pt]
&= \underbrace{\left(\tfrac{4}{3}\right)^{2}\!\!\cdot\!\left(\tfrac{3}{4}\right)^{2}}_{\text{as }1}\cdot\left(\tfrac{4}{3}\right)^{n-2}
 + \underbrace{\left(\tfrac{4}{3}\right)^{3}\!\!\cdot\!\left(\tfrac{3}{4}\right)^{3}}_{\text{as }1}\cdot\left(\tfrac{4}{3}\right)^{n-3} \\[4pt]
&= \left(\tfrac{4}{3}\right)^{n}\cdot\tfrac{9}{16} + \left(\tfrac{4}{3}\right)^{n}\cdot\tfrac{27}{64}
 = \left(\tfrac{4}{3}\right)^{n}\left(\tfrac{36}{64} + \tfrac{27}{64}\right) \\[4pt]
&= \left(\tfrac{4}{3}\right)^{n}\cdot\tfrac{63}{64} \;<\; \left(\tfrac{4}{3}\right)^{n}
\quad\Longrightarrow\quad \underbrace{0 \leq a_n \leq \left(\tfrac{4}{3}\right)^{n}}_{P(n)}
\end{aligned}
$$

Thus by P.M.I. (principle of mathematical induction):

$$
0 \leq a_n \leq \left(\tfrac{4}{3}\right)^{n} \quad \forall\, n \geq 0
$$

$$
\Downarrow
$$

$$
a_n = O\!\left(\left(\tfrac{4}{3}\right)^{n}\right)
$$

**Is it sharp?** If not, can we improve it, i.e. find $f(n)$ such that $a_n = O(f(n))$ where $0 \leq f(n) < \left(\tfrac{4}{3}\right)^n$? What is the sharp result, i.e. $a_n = \Theta(f(n))$?

All such questions are standard in the analysis of the algorithm's complexity. The answers are usually not trivial, and each time they require a separate analysis. For some classes of recursive schemes (common in computer science) one can derive a general theory:

$$
\text{implicit form} \;\overset{\text{how to transform}}{\Longrightarrow}\; \text{explicit form}
$$

## Estimating via a subsequence

Sometimes one chooses from

$$
a_0, a_1, a_2, a_3, \ldots, a_n, a_{n+1}, \ldots
$$

only some terms, as it may be easier to analyze them:

$$
a_{k_1}, a_{k_2}, \ldots, a_{k_n}, \ldots
$$

**a)** if $a_{k_n} = O(f(n))$ for $k_n \in \mathbb{N}$,

**b)** $a_{k_n} \leq a_m \leq a_{k_{n+1}}$ for all terms with indices $k_n < m < k_{n+1}$,

then

$$
\Downarrow
$$

$$
\boxed{\ a_n = O(f(n)) \ } \quad \text{for all } n \in \mathbb{N}.
$$

### (i) The harmonic sums

Let us look at the following case:

$$
S_n = 1 + \tfrac{1}{2} + \cdots + \tfrac{1}{n}
$$

From the theory of series:

$$
\sum_{k=1}^{\infty} \frac{1}{k} = \lim_{k \to \infty}\left(1 + \tfrac{1}{2} + \cdots + \tfrac{1}{k}\right) = +\infty
$$

— the **harmonic series**, divergent. But how quickly?

We choose the terms $S_n$ with indices $n = 2^k$, $k \in \mathbb{N}$:

$$
\begin{aligned}
S_{2^0} = S_1 &= 1 \\
S_{2^1} = S_2 &= 1 + \tfrac{1}{2} = S_1 + \tfrac{1}{2} < S_1 + \tfrac{1}{1} = 2 \\
S_{2^2} = S_4 &= 1 + \tfrac{1}{2} + \tfrac{1}{3} + \tfrac{1}{4} = S_2 + \tfrac{1}{3} + \tfrac{1}{4} < S_2 + \left(\tfrac{1}{2} + \tfrac{1}{2}\right) = S_2 + 1 < 3 \\
S_{2^3} = S_8 &= 1 + \tfrac{1}{2} + \tfrac{1}{3} + \tfrac{1}{4} + \tfrac{1}{5} + \tfrac{1}{6} + \tfrac{1}{7} + \tfrac{1}{8} \\
&= S_4 + \tfrac{1}{5} + \tfrac{1}{6} + \tfrac{1}{7} + \tfrac{1}{8} < S_4 + \tfrac{1}{4} + \tfrac{1}{4} + \tfrac{1}{4} + \tfrac{1}{4} \\
&= S_4 + 1 < 3 + 1 = 4. \\
&\ \ \vdots
\end{aligned}
$$

Looking at the above tendency we set up the theoretical hypothesis:

$$
\boxed{\ P(k): \quad S_{2^k} < k + 1 \ } \qquad \text{(obviously } S_n \geq 0\text{)}.
$$

**I.** $P(0)$: $S_{2^0} = S_1 < 2$, indeed $1 < 2$. OK.

**II.** $P(k): S_{2^k} < k+1 \ \overset{?}{\Longrightarrow}\ P(k+1): S_{2^{k+1}} < k+2$:

$$
S_{2^{k+1}} = S_{2^k} + \Bigl(\underbrace{\frac{1}{2^k+1} + \frac{1}{2^k+2} + \cdots + \frac{1}{2^k+2^k}}_{2\cdot 2^k \,=\, 2^{k+1}}\Bigr)
\underset{\text{by } P(k)}{<} (k+1) + \Bigl(\frac{1}{2^k+1} + \frac{1}{2^k+2} + \cdots + \frac{1}{2^{k+1}}\Bigr)
$$

but

$$
(\bullet) \quad
\left.
\begin{aligned}
\frac{1}{2^k+1} &< \frac{1}{2^k} \\
\frac{1}{2^k+2} &< \frac{1}{2^k} \\
&\ \vdots \\
\frac{1}{2^{k+1}} &< \frac{1}{2^k}
\end{aligned}
\right\} \quad 2^k \text{ terms}
$$

so by $(\bullet)$:

$$
S_{2^{k+1}} < (k+1) + \underbrace{\frac{1}{2^k} + \cdots + \frac{1}{2^k}}_{2^k \text{ times}} = (k+1) + \frac{2^k}{2^k} = k+2.
$$

Therefore:

**a)**

$$
\underbrace{S_{2^{k+1}} < (k+1) + 1}_{P(k+1)}
$$

Thus by P.M.I., $P(k)$ holds for all $k \in \mathbb{N}$.

But note that between $S_{2^n}$ and $S_{2^{n+1}}$ we have the terms:

**b)**

$$
S_{2^n} < S_{2^n+1} < S_{2^n+2} < \cdots < S_{2^{n+1}-1} < S_{2^{n+1}}
$$

Thus, since $n = 2^k \Rightarrow k = \log_2 n$:

$$
S_n = O(k+1) \quad \forall\, n \in \mathbb{N} \quad \text{(since a) and b))}
$$

hence

$$
\boxed{\ S_n = O(\log_2 n) \ }
$$

So the explosion to $\infty$ is controlled from above by a logarithmic curve (modulo multiplication by a constant). The teacher's small sketch of this slowly growing curve:

![[lec07b_p08_log-growth.svg|200]]

This is a slow explosion rate (e.g. slower than the linear one).

### (ii) A related sum

For

$$
t_n = n + \frac{n}{2} + \frac{n}{3} + \cdots + \frac{n}{n}
$$

we have

$$
t_n = \underbrace{n}_{O(n)} \cdot \underbrace{\left(1 + \tfrac{1}{2} + \tfrac{1}{3} + \cdots + \tfrac{1}{n}\right)}_{O(\log_2 n)}
$$

thus

$$
\boxed{\ t_n = O(n \log_2 n) \ }
$$

Note that we previously showed that:

## Divide-and-conquer recurrences

**1)**

$$
\begin{cases}
T(1) = A, & A > 0 \\
T(n) = 2\, T\!\left(\lfloor \tfrac{n}{2} \rfloor\right)
\end{cases}
\qquad \text{implicit form}
$$

$$
\Downarrow \ \text{proved by M.I.}
$$

$$
\begin{cases}
\text{a)} \ T(n) = A \cdot 2^k \\
\text{b)} \ 2^k \leq n \\
\text{c)} \ k \in \mathbb{N} \text{ maximal such that } 2^k \leq n
\end{cases}
\qquad \text{explicit form}
$$

It is hard to see from the implicit form that $T(n) = O(?)$ — this is a **Divide & Conquer** type algorithm. But from the explicit form we have:

$$
T(n) = A \cdot 2^k \quad (\text{as } 2^k \leq n)
$$

$$
\underset{\text{obvious}}{0 \leq} T(n) \leq A \cdot n \quad \Longrightarrow \quad \boxed{\ T(n) = O(n) \ }
$$

**2)** We can also see, analogously:

$$
\underbrace{\begin{cases}
T(1) = A \\
T(n) = 4\, T\!\left(\lfloor \tfrac{n}{2} \rfloor\right)
\end{cases}}_{\text{I.F.}}
\quad\Longrightarrow\quad
\underbrace{\begin{cases}
T(n) = A \cdot 4^k \\
2^k \leq n \\
k \text{ max s.t. } 2^k \leq n
\end{cases}}_{\text{E.F.}}
$$

Thus

$$
0 \leq T(n) = A \cdot 4^k = A \cdot 2^k \cdot 2^k \leq A \cdot n^2
\quad\Longrightarrow\quad \boxed{\ T(n) = O(n^2) \ }
$$

**And in general:**

$$
\text{I.F.}
\begin{cases}
T(1) = A \\
T(n) = (B = 2^{\ell}) \cdot T\!\left(\lfloor \tfrac{n}{2} \rfloor\right)
\end{cases}
\quad\Longrightarrow\quad
\text{E.F.}
\begin{cases}
\text{a)} \ T(n) = A \cdot B^{k} \\
\text{b)} \ 2^k \leq n \\
\text{c)} \ k \text{ max, } 2^k \leq n
\end{cases}
$$

$$
0 \leq T(n) = A \cdot B^k = A\,(\underbrace{2 \cdot 2 \cdots 2}_{\ell\text{-times}})^{k}
= A\,\underbrace{2^k \cdot 2^k \cdots 2^k}_{\ell\text{-times}}
\underset{\text{b)}}{\leq} A\,\underbrace{n \cdot n \cdots n}_{\ell\text{-times}}
$$

$$
\Longrightarrow \quad \boxed{\ T(n) = O(n^{\ell}) \ }
$$

If $B \neq 2^{\ell}$, e.g. $B = 3$:

$$
\text{I.F.}
\begin{cases}
T(1) = A \\
T(n) = 3\, T\!\left(\lfloor \tfrac{n}{2} \rfloor\right)
\end{cases}
\quad\Longrightarrow\quad
\begin{cases}
\text{a)} \ T(n) = A \cdot 3^{k} \\
\text{b)} \ 2^k \leq n \\
\text{c)} \ k \text{ max such that } 2^k \leq n
\end{cases}
$$

then, using $3 = 2^{\log_2 3}$:

$$
0 \leq T(n) = A \cdot 3^k = A\left(2^{\log_2 3}\right)^{k} = A \cdot 2^{k \log_2 3}
= A\left(2^{k}\right)^{\log_2 3}
\underset{\text{(b)}}{\leq} A\, n^{\log_2 3}
$$

$$
\Downarrow
$$

$$
\boxed{\ T(n) = O\!\left(n^{\log_2 3}\right) \ }
$$

Note $1 < \log_2 3 < 2$, so we have an order of explosion faster than linear and slower than quadratic.

## Tools: integration, series, telescoping

We showed that to examine the complexity of algorithms (and possibly the sharpness of estimates):

- **calculus of integration** is needed sometimes (e.g. $a_n = \ln n!$);
- or **series theory**, to assess exactly the sum or its upper bound, e.g.:

$$
a_0\left(1 + q + q^2 + \cdots + q^n + \cdots\right) = \frac{a_0}{1-q}, \qquad |q| < 1
$$

(sum of a geometric progression);

$$
t_n = 1 + \tfrac{1}{2} + \cdots + \tfrac{1}{n} \leq \text{upper bound was found.}
$$

In general, exact counting of sums can be difficult. Sometimes

$$
\boxed{\ \textbf{a telescoping technique can be useful:} \ }
$$

Indeed, observe

$$
S_n = \frac{1}{1 \cdot 2} + \frac{1}{2 \cdot 3} + \frac{1}{3 \cdot 4} + \cdots + \frac{1}{(n-1)n} + \frac{1}{n(n+1)}.
$$

If

$$
a_k = \frac{1}{k(k+1)} \underset{\substack{\uparrow \\ \text{observe that}}}{=} \frac{1}{k} - \frac{1}{k+1}
$$

thus

$$
S_n = \left(\frac{1}{1} - \cancel{\frac{1}{2}}\right) + \left(\cancel{\frac{1}{2}} - \cancel{\frac{1}{3}}\right) + \left(\cancel{\frac{1}{3}} - \cancel{\frac{1}{4}}\right) + \cdots + \left(\cancel{\frac{1}{n-1}} - \cancel{\frac{1}{n}}\right) + \left(\cancel{\frac{1}{n}} - \frac{1}{n+1}\right) = 1 - \frac{1}{n+1}
$$

$$
S_n = 1 - \frac{1}{n+1} \quad \Longrightarrow \quad S_n = \Theta(1).
$$

### Generalization to arithmetic sequences

One can generalize that technique to the schemes involving sums

$$
(\blacktriangle) \quad \sum_{k=0}^{\infty} \frac{1}{a_k a_{k+1}}
\qquad \text{provided } a_{k+1} - a_k = r \neq 0 \ \text{ (arithmetic seq.)}
$$

$$
= \lim_{n} S_n, \qquad S_n = \frac{1}{a_0 a_1} + \cdots + \frac{1}{a_n a_{n+1}}.
$$

Note:

$$
\frac{1}{a_k} - \frac{1}{a_{k+1}} = \frac{a_{k+1} - a_k}{a_k a_{k+1}} = \frac{r}{a_k a_{k+1}}
$$

$$
\Downarrow
$$

$$
\frac{1}{a_k a_{k+1}} = \frac{1}{r}\left(\frac{1}{a_k} - \frac{1}{a_{k+1}}\right)
$$

Hence:

$$
\begin{aligned}
S_n &= \frac{1}{a_0 a_1} + \frac{1}{a_1 a_2} + \frac{1}{a_2 a_3} + \cdots + \frac{1}{a_{n-1} a_n} + \frac{1}{a_n a_{n+1}} \\[4pt]
&= \frac{1}{r}\left(\frac{1}{a_0} - \cancel{\frac{1}{a_1}}\right) + \frac{1}{r}\left(\cancel{\frac{1}{a_1}} - \cancel{\frac{1}{a_2}}\right) + \frac{1}{r}\left(\cancel{\frac{1}{a_2}} - \cancel{\frac{1}{a_3}}\right)
+ \cdots + \frac{1}{r}\left(\cancel{\frac{1}{a_{n-1}}} - \cancel{\frac{1}{a_n}}\right) + \frac{1}{r}\left(\cancel{\frac{1}{a_n}} - \frac{1}{a_{n+1}}\right) \\[4pt]
&= \frac{1}{r}\left(\frac{1}{a_0} - \frac{1}{a_{n+1}}\right)
\end{aligned}
$$

In the 1st example, $\frac{1}{n(n+1)}$ is the special case

$$
a_k = k, \qquad a_{k+1} = k+1, \qquad a_{k+1} - a_k = 1 = r.
$$

But telescoping can also be applied to other sums, different than $(\blacktriangle)$:

$$
\sum_{n=1}^{\infty}\left(\sqrt[n]{n} - \sqrt[n+1]{n+1}\right) = \lim_{n} S_n
$$

$$
S_n = \left(\sqrt[1]{1} - \cancel{\sqrt[2]{2}}\right) + \left(\cancel{\sqrt[2]{2}} - \cancel{\sqrt[3]{3}}\right) + \cdots + \left(\cancel{\sqrt[n]{n}} - \sqrt[n+1]{n+1}\right) = 1 - \sqrt[n+1]{n+1}
$$

### Exercises

One can apply the above telescoping method to the following examples:

**a)**

$$
\sum_{k=1}^{n} \frac{1}{(2k-1)(2k+1)} = \frac{n}{2n+1}
$$

**b)**

$$
\sum_{k=1}^{n} \frac{1}{(3k-2)(3k+1)} = \frac{n}{3n+1}
$$

**c)**

$$
\sum_{k=0}^{n-1} \frac{1}{(n+k)(n+k+1)} = \frac{1}{2n}
$$

**d)**

$$
\sum_{k=0}^{n-1} \frac{1}{(n-2k)(n-2k+2)} = \frac{n}{4-n^2} \qquad \text{for } n \text{ odd.}
$$

## The goals in the analysis of algorithms

Recall that in the analysis of algorithms we wish to (for $a_n$):

- **compute exactly** the number $a_n$ (at best), which counts e.g. memory consumption, number of iterations, execution time, etc. This however (the exact count) is usually difficult;
- then, if $a_n$ cannot be computed exactly, we try first to find an **upper bound** in terms of:

$$
a_n = O(f(n))
$$

- once the latter is found with a naive method, we try to **improve** it with a more advanced method (which does not always give such an improvement), so that:

$$
a_n = O(g(n)) \quad \text{where} \quad g(n) = O(f(n)).
$$

Any improvement is a step in a good direction;

- in the process of such improvement we keep in mind the **ultimate goal**:

$$
a_n = \Theta(h(n)) \qquad \text{"sharp estimate"}
$$

Here computers can help: if we have $a_n = O(f(n))$, we can plot $(n, K f(n))$ ($K$ — a constant) against $(n, a_n)$ and try to see whether we can improve our estimate or not.

## Example: naive vs. sharp estimates

### a) Bubble sort

For **Bubble sort**:

$$
a_n = 1 + 2 + \cdots + (n-1)
$$

**(i)** If naive counting is used:

$$
0 \leq a_n = 1 + 2 + \cdots + (n-1) \leq \underbrace{n + n + \cdots + n + n}_{n-1}
$$

$$
0 \leq a_n \leq n(n-1) \quad (*) \qquad \Longrightarrow \qquad \boxed{\ a_n = O(n^2) \ } \quad (*)
$$

**(ii)** But the exact counting yields

$$
(**) \quad a_n = \frac{n(n-1)}{2}
$$

so it is better than $(*)$ (modulo a constant). But it also shows

$$
\boxed{\ a_n = \Theta(n^2) \ }
$$

In this example, since the exact result was "computable", we have an indication from $(**)$ about the sharp estimate immediately (the naive method gives that).

### b) The harmonic sums again

Take now $S_n = 1 + \tfrac{1}{2} + \cdots + \tfrac{1}{n}$; here it is hard to count $S_n$.

**(i)** Naive estimation:

$$
0 \leq S_n \leq \underbrace{1 + 1 + \cdots + 1}_{n} = n \qquad \Longrightarrow \qquad \boxed{\ S_n = O(n) \ }
$$

**(ii)** More advanced estimation (if you plot $(n, S_n)$ against $(n, n)$ one can see that $S_n$ is well below $n$). In this lecture we showed (with a trickier proof):

$$
S_n = O(\log_2 n)
$$

which is better, as $\log_2 n = O(n)$. So the naive method did **not** give us a sharp result.

Is now $S_n = O(\log_2 n)$ sharp? I.e. is $S_n = \Theta(\log_2 n)$?

A possible tool is to use the **"integration approach"**, which can prove both:

$$
\begin{aligned}
\text{(i)} \quad & 0 \leq S_n \leq M_2 \log_2 n \quad \Longrightarrow \quad S_n = O(\log_2 n) \\
\text{(ii)} \quad & M_1 \log_2 n \leq S_n \quad \Longrightarrow \quad \text{(i) \& (ii)}: \ S_n = \Theta(\log_2 n)
\end{aligned}
$$

We only show (ii), as (i) was proved by a different technique and follows also analogously as in (ii).

## Sharpness via the integration approach

We show now the sharpness:

$$
\boxed{\ S_n = \Theta(\log_2 n) \ }
$$

The teacher's picture — rectangles $P_1, P_2, \ldots, P_n$ of heights $1, \tfrac{1}{2}, \ldots, \tfrac{1}{n}$ erected over the unit intervals $[1,2], [2,3], \ldots, [n, n+1]$, lying above the graph of $f(x) = \tfrac{1}{x}$:

![[lec07b_p18_harmonic-integral.svg]]

$P_i$ — a rectangle, $A_i$ — the area of $P_i$. From the picture it is clear that $A_1 + A_2 + \cdots + A_n \geq$ the area under $f(x)$ between $1$ and $n+1$:

![[lec07b_p18_area-under-curve.svg|200]]

$$
(2-1)\cdot\tfrac{1}{1} + (3-2)\cdot\tfrac{1}{2} + (4-3)\cdot\tfrac{1}{3} + \cdots + \left[n-(n-1)\right]\cdot\tfrac{1}{n-1} + (n+1-n)\cdot\tfrac{1}{n}
\;\geq\; \int_{1}^{n+1} \frac{1}{x}\, dx
$$

$$
\Updownarrow
$$

$$
\underbrace{1 + \tfrac{1}{2} + \tfrac{1}{3} + \cdots + \tfrac{1}{n-1} + \tfrac{1}{n}}_{S_n}
\;\geq\; \ln x \,\Big|_{1}^{\,n+1} = \ln(n+1) - \underbrace{\ln 1}_{0} = \ln(n+1)
$$

$$
S_n \geq \ln(n+1) \geq \ln(n) = M_1 \log_2 n
$$

but

$$
\ln n = \log_e n = \underbrace{\log_e 2}_{M_1} \cdot \log_2 n
\qquad \Longrightarrow \qquad \boxed{\ S_n = \Theta(\log_2 n) \ }
$$
