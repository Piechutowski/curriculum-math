# Discrete Mathematics — Lecture 6: Asymptotics

> Translated from the handwritten lecture notes `pdf/lect6-AsympPL-MatDysk.pdf`.

We need criteria for estimating certain quantities, e.g. in order to compare algorithms that accomplish the same goal.

## Definition 1 — big-$O$ notation

Let two number sequences be given: $a_n = f(n)$ and $b_n = g(n)$. We say that

$$
\boxed{\ a_n = O(b_n)\ }
$$

when

$$
\exists M > 0 \ \ \exists n_0 \ \ \forall n \geq n_0 \qquad |a_n| \leq M\, b_n .
$$

That is, $a_n$ is estimated **from above** by $b_n$ from some point on.

**Remark:** the estimate may be non-unique, because:

**Example:**

$$
\begin{aligned}
a_n = 5n^2 &\leq n^3 \qquad n \geq n_0 \\
5n^2 &\leq n^4 \qquad n \geq n_0'
\end{aligned}
$$

so

$$
a_n = O(n^3) \ \text{ and } \ a_n = O(n^4).
$$

One can see that

$$
a_n = O\!\left(n^{2+\varepsilon}\right) \qquad \forall \varepsilon \geq 0. \ \square
$$

## More complicated forms of $a_n$

Sometimes, however, $a_n$ may come in a more complicated form.

**a)**

$$
\boxed{\ a_n = \sum_{i=1}^{n} i\ } \stackrel{(*)}{=} n\,\frac{(1+n)}{2}
$$

From $(*)$ one sees that $a_n = O(n^2)$ (**quadratic growth**).

**b)**

$$
\boxed{\ T_{n+1} = 2T_n + 1\ } \qquad \text{(towers of Hanoi)}
$$

From the explicit (closed) form we showed that

$$
T_n = 2^n - 1 \quad \Longrightarrow \quad T_n = O(2^n)
$$

— **non-polynomial (exponential) growth**; an NP-problem.

**c)**

$$
\boxed{\ T(1) = 1, \qquad T(n) = 2\,T\!\left(\left\lfloor \tfrac{n}{2} \right\rfloor\right)\ }
$$

We showed the explicit form:

$$
T(n) = 2^k \quad \text{where } 2^k \leq n \quad \Longrightarrow \quad \boxed{\ T(n) = O(n)\ }
$$

> **Note:** the handwritten page –3– is missing from the scanned PDF (it jumps from page –2– to page –4–). Example **d)** below — the harmonic sums $S_n = 1 + \frac12 + \cdots + \frac1n$ together with the bound $S_{2^k} \leq k+1$ — evidently begins on that missing page; only its conclusion survives and is translated here.

**d)** (harmonic sums, continued)

**Question:** what happens with the sequence for indices

$$
n_k = 2^k < n < 2^{k+1} = n_{k+1} \ ?
$$

But

$$
S_{n+1} > S_n \quad (*)
$$

since

$$
1 + \tfrac{1}{2} + \cdots + \tfrac{1}{n} + \tfrac{1}{n+1} \ > \ 1 + \tfrac{1}{2} + \cdots + \tfrac{1}{n}
\qquad \boxed{\tfrac{1}{n+1} > 0} \ \ \text{OK}.
$$

For $n < 2^k$:

$$
S_n \stackrel{(*)}{<} S_{2^k} \stackrel{(*)}{<} S_{2^k+1} \stackrel{(*)}{<} S_{2^k+2} \stackrel{(*)}{<} \cdots < S_{2^{k+1}} \quad (\blacktriangle)
$$

$$
S_n < \begin{cases} k+1 & n = 2^k \quad (\blacktriangle_1) \\ k+1 & 2^{k-1} < n < 2^k \quad (\blacktriangle) \end{cases}
$$

Setting $2^k = n$:

$$
\boxed{\ \lg_2 n = k\ }
$$

$$
S_n < \lg_2 n + 1 \leq 2\lg_2 n
$$

$$
\boxed{\ S_n = O(\lg_2 n)\ }
$$

**e)**

$$
\boxed{\ t_n = n + \frac{n}{2} + \cdots + \frac{n}{n} = n\left(1 + \frac{1}{2} + \cdots + \frac{1}{n}\right)\ }
= n\, S_n \ \Longrightarrow \ t_n \leq n \cdot 2\lg_2 n = \underline{\underline{O(n \lg_2 n)}}
$$

## What the notation is used for

The above notion $a_n = O(b_n)$ is useful for estimating, when $n \gg c$ (a constant):

- **memory** usage,
- the number of **basic operations**,
- the amount of **time** used by a specific algorithm.

It rarely happens that a given algorithm is the best with respect to **all** criteria.

## Hierarchy of growth of powers

The following inequalities are satisfied:

$$
\sqrt[4]{n} \leq \sqrt[3]{n} \leq \sqrt{n} \leq n < n^2 \leq n^3 \leq \cdots \leq n^m
$$

$$
n^{\alpha} \leq n^{\beta} \qquad 0 \leq \alpha \leq \beta .
$$

The teacher's three sketches of $x^\alpha$ vs $x^\beta$ — for exponents above $1$ (convex curves crossing at $(1,1)$, with $x^\alpha < x^\beta$ past $1$), for $\alpha = 1$ (the straight line), and for exponents between $0$ and $1$ (concave curves):

![[lec06_p04_power-growth.svg]]

Moreover:

$$
n^{\alpha} \leq 2^n \qquad \forall n \geq n_0
$$

$$
\lg_2 n \leq n
$$

$$
2^n < n! < n^n
$$

A margin doodle comparing the line $y = x$ with the slowly growing $y = \lg x$:

![[lec06_p05_log-vs-linear.svg|200]]

## Bounded sequences: $O(1)$

$$
a_n = O(1) \ \equiv \ |a_n| \leq M \quad \forall n \geq n_0
$$

— the sequence $a_n$ is **bounded**.

**Example:**

$$
a_n = \begin{cases} \dfrac{1}{n} & n \text{ even} \\[4pt] 2 + (-1)^n & n \text{ odd} \end{cases}
\qquad\Longrightarrow\qquad a_n = O(1).
$$

## Theorem — a hierarchy of known sequences

**Thm.** A hierarchy of certain known sequences, ordered in such a way that each of them is of order $O$ of all the sequences to its right:

$$
\boxed{\ 1,\ \ \lg_2 n,\ \ \sqrt[m]{n}\,{}^{(*)},\ \ n,\ \ n\lg_2 n,\ \ n^2,\ \ n^3,\ \ n^4,\ \ n^m\,{}^{(**)},\ \ 2^n,\ \ n!,\ \ n^n\ }
$$

$(*)$ among the roots: $\sqrt[m_1]{n},\ \sqrt[m_2]{n}$ are ordered with $m_2 < m_1$ (the larger the root index, the slower the growth);
$(**)$ among the powers: $n^{m_1} < n^{m_2}$ for $m_1 < m_2$.

E.g. the hierarchy gives $\lg_2 n = O(n!)$.

## Names of complexities

- $O(n)$ — **linear**
- $O(n\sqrt{n})$ — **"one-and-a-half" linear** (i.e. $n^{3/2}$)
- $O(n^2)$ — **quadratic**
- $O(\lg_2 n)$ — **logarithmic**
- $O(n \lg_2 n)$ — **log-linear**
- $O(2^n)$ — **exponential**

The standard complexities are

$$
O(n^r \lg_2 n), \qquad O(a^n), \qquad O(n^r).
$$

> Margin note: *the battle* — to get an algorithm whose cost $T$ lands between $O(n)$ and $O(n^2)$.

## Theorem 2 — algebra of $O$

**Thm 2.**

$$
\begin{aligned}
\text{a)}\quad & a_n = O(b_n) \ \Longrightarrow\ c \cdot a_n = O(b_n) \\[4pt]
\text{b)}\quad & \left.\begin{aligned} a_n &= O(b_n) \\ c_n &= O(b_n) \end{aligned}\right\} \ \Longrightarrow\ a_n \pm c_n = O(b_n) \\[4pt]
\text{c)}\quad & \left.\begin{aligned} a_n &= O(b_n) \\ c_n &= O(d_n) \end{aligned}\right\} \ \Longrightarrow\ a_n \cdot c_n = O(b_n \cdot d_n) \\[4pt]
\text{d)}\quad & \left.\begin{aligned} a_n &= O(b_n) \\ b_n &= O(c_n) \end{aligned}\right\} \ \Longrightarrow\ a_n = O(c_n)
\end{aligned}
$$

**Proof** of e.g. b):

$$
\begin{aligned}
|a_n| \leq K_1 b_n, \quad |c_n| \leq K_2 b_n
\quad&\Longrightarrow\quad |a_n| + |c_n| \leq (K_1 + K_2)\, b_n, \\
\text{but } |a_n + c_n| \leq |a_n| + |c_n|
\quad&\Longrightarrow\quad a_n + c_n = O(b_n). \ \square
\end{aligned}
$$

## Examples

**a)** A polynomial:

$$
a_n = W_k(n) = n^k a_k + n^{k-1} a_{k-1} + \cdots + n a_1 + a_0
$$

$$
\begin{aligned}
|a_n| &\leq |a_k|\, n^k + \cdots + n\,|a_1| + |a_0| \\
&\leq |a_k|\, n^k + \cdots + n^k |a_1| + n^k |a_0| \\
&= \left( |a_k| + \cdots + |a_1| + |a_0| \right) n^k = \tilde{C}\, n^k \\
&\Longrightarrow \ a_n = O(n^k)
\end{aligned}
$$

### Lemma 1

$$
\boxed{\ O(a_n) + O(b_n) = O\!\left(\max\{a_n, b_n\}\right)\ }
$$

**b)**

$$
\begin{aligned}
f_n = f(n) = O(n^5) \qquad & f_n + g_n = O(n^5) \\
g_n = g(n) = O(n^4) \qquad & f_n \cdot g_n = O(n^9)
\end{aligned}
$$

**c)**

$$
a_n = \ln n! = \ln\bigl(n (n-1) \cdots 2 \cdot 1\bigr) = \ln n + \ln(n-1) + \cdots + \ln 2 + \ln 1 = \sum_{i=1}^{n} \ln i .
$$

The teacher's sketch: unit-width rectangles $S^2, S^3, \ldots, S^n$ tucked under the curve $y = \ln x$ between $1$ and $n+1$, next to the hatched area $S_1 = \int_1^{n+1} \ln x \, dx$:

![[lec06_p07_ln-sum-rectangles.svg]]

$$
S = S^1 + S^2 + \cdots + S^n \quad \Longrightarrow \quad \boxed{\ S \leq S_1\ }
$$

where

$$
S^i = 1 \cdot \ln i, \qquad \underline{S^1 = 0}, \quad S^2 = 1 \cdot \ln 2, \quad S^3 = 1 \cdot \ln 3, \ \ldots, \ S^n = 1 \cdot \ln n .
$$

Hence

$$
a_n = \sum_{i=1}^{n} \ln i \ \leq \ \int_{1}^{n+1} \ln x \, dx = \int_{1}^{n+1} x' \ln x \, dx
= x \ln x \Big|_1^{n+1} - \int_{1}^{n+1} x \cdot \frac{1}{x}\, dx
$$

(integration by parts)

$$
\begin{aligned}
&= (n+1)\ln(n+1) - \underbrace{1 \cdot \ln 1}_{=\,0} - \ x \Big|_1^{n+1} \\
&= (n+1)\ln(n+1) - (n+1-1) \\
&= (n+1)\ln(n+1) - n
\end{aligned}
$$

One sees that

$$
\boxed{\ a_n = O(n \lg_2 n)\ }
$$

Often it is the case that we know the optimal complexity cannot be smaller than, e.g., $O(f(n))$. **The battle** is to find an algorithm of a complexity close to it.

## Another application: approximating the length of a curve

**Example:** we need to approximate the length of an **unknown curve** of which we know the **trace**, i.e. interpolation points ($n$ temporarily frozen).

The teacher's sketch of the curve $\gamma$ with its known interpolation points:

![[lec06_p09_curve-trace.svg]]

In mathematics one can use various curves passing through the $p_i$ (denote them $\gamma_s^n$). One can show that as $n \to \infty$, for specific curves,

$$
\lim_{n \to \infty} \left( d(\gamma_s^n) - d(\gamma) \right) = 0 \quad (*)
$$

i.e. we approximate the length of the unknown curve $\gamma$ by the interpolant.

But there are many families of curves (e.g. Lagrange polynomials, splines, piecewise-glued Lagrange functions, etc.).

The question now is not only **whether** $(*)$ holds, but whether it happens slowly or fast, i.e.:

$$
g(n) = d(\gamma_s^n) - d(\gamma) = O\!\left(\frac{1}{n^{\alpha}}\right), \qquad \underline{\alpha > 0}.
$$

$$
0 \ \leq \ \left| d(\gamma_s^n) - d(\gamma) \right| \ \leq \ K\, \frac{1}{n^{\alpha}}
$$

As $n \to \infty$ the right-hand side $\to 0$, so the middle $\to 0$ as well. But from the right-hand inequality we can **read off how fast**, e.g.:

$$
\begin{aligned}
\alpha = 1: \quad & g(n) = O\!\left(\tfrac{1}{n}\right) \\
\alpha = 2: \quad & g(n) = O\!\left(\tfrac{1}{n^2}\right) \quad \text{faster — quadratically} \\
\alpha = 3: \quad & g(n) = O\!\left(\tfrac{1}{n^3}\right) \quad \text{cubic, quite fast}
\end{aligned}
$$

## Example: rates of convergence

**a)** Taylor's formula with remainder:

$$
f(x) = f(a) + \frac{f'(a)}{1!}(x-a) + \frac{f''(a)}{2!}(x-a)^2 + \cdots + \frac{f^{(n)}(a)}{n!}(x-a)^n + \frac{f^{(n+1)}(\xi)}{(n+1)!}(x-a)^{n+1}
$$

$$
\boxed{\ 0 < \xi < x \ \text{ or } \ x < \xi < 0\ }
$$

For $f(x) = e^x$, $a = 0$, $x > 0$:

$$
e^x = e^0 + \frac{e^0}{1!}x + \frac{e^0}{2!}x^2 + \cdots + \frac{e^0}{n!}x^n + \frac{e^{\xi}}{(n+1)!}x^{n+1}, \qquad 0 < \xi < x .
$$

At $x = 1$:

$$
e^1 = \underbrace{1 + \frac{1}{1!} + \frac{1}{2!} + \cdots + \frac{1}{n!}}_{a_n} + \frac{e^{\xi}}{(n+1)!},
\qquad 0 < e^{\xi} < e^1 < \underline{\underline{3}} .
$$

$$
0 \ \leq \ |e - a_n| = \frac{e^{\xi}}{(n+1)!} \ \leq \ \frac{3}{(n+1)!} \ \xrightarrow[n \to \infty]{} \ 0,
$$

so $|e - a_n| \to 0$ as well. We have, moreover, that

$$
\boxed{\ |e - a_n| = O\!\left(\frac{1}{n!}\right)\ }
$$

i) i.e. $a_n$ approximates the number $e$;
ii) it approximates **very fast** — the order of convergence is $\frac{1}{n!}$.

**b)** Let us estimate the area under $f(x) = x^2$:

$$
\boxed{\ \text{Area} = \int_0^1 x^2 \, dx\ }
$$

The hatched region under the parabola between $0$ and $1$:

![[lec06_p11_area-under-x2.svg|280]]

Approximating it with right-endpoint rectangles of width $\frac{1}{n}$:

![[lec06_p12_riemann-sum.svg]]

$$
S_n = \sum_{i=0}^{n-1} \frac{1}{n} \cdot f\!\left( \frac{i+1}{n} \right) \quad \longleftarrow \ \text{Riemann's sum}
$$

$$
\lim_{n \to \infty} \sum_{i=0}^{n-1} \frac{1}{n}\, f\!\left( \frac{i+1}{n} \right) = \int_0^1 x^2 \, dx
$$

$$
E_n = \text{error}_n = S_n - \int_0^1 x^2\, dx = E_n^0 + E_n^1 + E_n^2 + \cdots + E_n^{n-1}
$$

— the sum of the little areas between the rectangles and the curve. We have convergence to Area! **But how fast?**

$$
E_n = \sum_{i=0}^{n-1} \Biggl[\, \frac{1}{n} \left( \frac{i+1}{n} \right)^2 - \int_{i/n}^{(i+1)/n} x^2 \, dx \,\Biggr]
$$

The teacher's margin doodle: on each cell $\left[\frac{i}{n}, \frac{i+1}{n}\right]$ the term in brackets is the hatched sliver $E_n^i$ between the rectangle's top and the curve:

![[lec06_p12_error-cell.svg|220]]

$$
\begin{aligned}
E_n &= \sum_{i=0}^{n-1} \left[ \frac{1}{n^3}\left( i^2 + 2i + 1 \right) - \frac{1}{3}\left( \frac{i+1}{n} \right)^{3} + \frac{1}{3}\left( \frac{i}{n} \right)^{3} \right] \\
&= \sum_{i=0}^{n-1} \frac{1}{n^3} \left[ i^2 + 2i + 1 + \frac{-i^3 - 3i^2 - 3i - 1}{3} + \frac{1}{3} i^3 \right] \\
&= \sum_{i=0}^{n-1} \frac{1}{n^3} \left[ i^2 + 2i + 1 - \frac{1}{3}i^3 - i^2 - i - \frac{1}{3} + \frac{1}{3} i^3 \right] \\
&= \sum_{i=0}^{n-1} \frac{1}{n^3} \left[ i + \frac{2}{3} \right]
\end{aligned}
$$

$$
\begin{aligned}
&= \frac{1}{n^3} \underbrace{\sum_{i=0}^{n-1} i}_{\text{arithmetic sequence}} + \frac{1}{n^3} \sum_{i=0}^{n-1} \frac{2}{3}
= \frac{(n-1)+0}{2}\, n\, \frac{1}{n^3} + \frac{2}{3}\, \frac{1}{n^2} \\[4pt]
&= \frac{(n-1)\,n}{2}\, \frac{1}{n^3} + \frac{2}{3}\, \frac{1}{n^2}
= \frac{1}{2n} - \frac{1}{2n^2} + \frac{2}{3n^2}
= \frac{1}{2n} + \frac{1}{6n^2} \ \leq \ \frac{1}{2n} + \frac{1}{6n}
\end{aligned}
$$

$$
\text{Error}_n = O\!\left(\frac{1}{n}\right) \qquad \text{— linear convergence.}
$$

## Definition 2 — the $\Theta$ notation (sharp estimate)

Let us note:

$$
a_n = 5n = O(n) = O(n^{1+\varepsilon}) \qquad \forall \varepsilon > 0 .
$$

We want a **sharp** estimate, i.e. the "smallest" one.

**Def 2.**

$$
f(n) = \Theta(g(n))
$$

$$
\exists c_1 \text{ and } \exists c_2 \ \ \forall n > n_0 \qquad c_2\, g(n) \leq f(n) \leq c_1\, g(n)
$$

The sketch: $f(n)$ sandwiched between $c_2 g(n)$ and $c_1 g(n)$:

![[lec06_p13_theta-sandwich.svg|300]]

$\Theta(g(n))$ (the **sharp estimate**) is the asymptotically **exact** estimate.

## Definition 3 — $\mathrm{Ord}$ and little-$o$

**Def 3.**

**a)** $a_n = \mathrm{Ord}(b_n)$ if

$$
\lim_{n \to \infty} \frac{a_n}{b_n} = c \neq 0
$$

— $a_n$ is **of the order (degree) of** $b_n$. The teacher's margin doodle — two sequences climbing together at the same rate:

![[lec06_p14_same-order.svg|160]]

**b)** If

$$
\lim_{n \to \infty} \frac{a_n}{b_n} = 0,
$$

then we write $a_n = o(b_n)$ — $a_n$ is **negligible** relative to $b_n$. We also write

$$
\mathrm{Ord}(a_n) < \mathrm{Ord}(b_n).
$$

### Theorem

**a)**

- $\mathrm{Ord}(n^r) < \mathrm{Ord}(n^s)$ for $r < s$
- $\mathrm{Ord}(n^r) < \mathrm{Ord}(n^r \lg n) < \mathrm{Ord}(n^{r+1})$
- $\mathrm{Ord}(a_1^{\,n}) < \mathrm{Ord}(a_2^{\,n})$ for $a_1 < a_2$
- $\mathrm{Ord}(n^r) < \mathrm{Ord}(a^n)$ for all $r$, $a > 1$

**Proof** of a):

$$
\lim_{n} \frac{n^r}{n^s} = \lim_{n} \frac{1}{n^{s-r}} = 0 . \ \checkmark
$$

### Lemma

**a)** If $\mathrm{Ord}(a_n) < \mathrm{Ord}(b_n)$, then

$$
a_n + b_n = \mathrm{Ord}(b_n),
$$

since

$$
\lim_{n} \frac{a_n + b_n}{b_n} = \underbrace{\lim_{n} \frac{a_n}{b_n}}_{=\,0} + 1 = 1 .
$$

**b)** If

$$
a_n = \mathrm{Ord}(b_n) \quad \text{and} \quad \mathrm{Ord}(b_n) < \mathrm{Ord}(c_n) \ \ (b_n = o(c_n)),
$$

then $\underline{\mathrm{Ord}(a_n) < \mathrm{Ord}(c_n)}$. Indeed:

$$
\begin{cases}
\displaystyle \lim_{n} \frac{a_n}{b_n} = C \neq 0 & \text{(this we know)} \\[8pt]
\displaystyle \lim_{n} \frac{b_n}{c_n} = 0
\end{cases}
$$

Hence

$$
\lim_{n} \frac{a_n}{c_n} = \lim_{n} \left( \frac{a_n}{b_n} \cdot \frac{b_n}{c_n} \right)
= \lim_{n} \frac{a_n}{b_n} \cdot \lim_{n} \frac{b_n}{c_n} = C \cdot 0 = 0
$$

$$
\Downarrow
$$

$$
\boxed{\ \mathrm{Ord}(a_n) < \mathrm{Ord}(c_n)\ } \qquad \square
$$
