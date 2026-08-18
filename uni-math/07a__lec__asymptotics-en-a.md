# Discrete Mathematics — Lecture 7: Asymptotics of Sequences (Part A)

> Translated from the handwritten lecture notes `pdf/lect7-AsympEn-A-MatDysk.pdf`. (The original pages are headed "Algorithms & Complexity — Lecture 3" and were already written in English.)

## Motivation

We now introduce formal tools to measure — with some sort of relaxation on multiplication by constants — the **rate of growth** of $a_n$ once all indices $n \ge n_0$ (i.e. for sufficiently large $n$).

Recall that a sequence $\{a_n\}_{n \ge 0}$ can be viewed as a function

$$
f : \mathbb{N} \to \mathbb{R} \quad \text{s.t.} \quad f(n) = a_n, \qquad \mathbb{N} = \{0, 1, 2, \ldots\}
$$

## Definition 1 — domination (big $O$)

Let $f, g : \mathbb{N} \to \mathbb{R}$. We say that **$g$ dominates $f$** (or $f$ is **dominated by** $g$) if there exist a constant $M \in \mathbb{R}_+$ and an index $k_0 \in \mathbb{N}$ such that

$$
\boxed{\,|f(n)| \le M \cdot |g(n)|\,} \qquad \forall n \ge k_0
$$

$$
\left[\ \text{or, in sequence notation:} \quad |a_n| \le M |b_n| \quad \forall n \ge k_0\ \right]
$$

If both $f(n) \ge 0$ (i.e. $a_n \ge 0$) and $g(n) \ge 0$ (i.e. $b_n \ge 0$), then we have a shorter notation:

$$
0 \le a_n \le M b_n
$$

In computer science $a_n$ usually represents positive values, as it counts things (time, memory consumption, number of basic calculations, etc.).

**Geometrically** Def. 1 means (for simplicity assume $f, g \ge 0$): as we consider the values $f(1) = a_1$, $f(2) = a_2$, $f(3) = a_3, \ldots$ and $g(1) = b_1$, $g(2) = b_2$, $g(3) = b_3, \ldots$ — there is a point (namely $k_0$) after which the size of $f(n)$ is bounded above by the positive multiple $M$ of the size of $g(n)$.

We say, if Def. 1 holds:

$$
a_n = O(b_n) \qquad \left[\, f(n) = O(g(n)) \,\right]
$$

— "$a_n$ is **of order** $b_n$".

**Example:**

$$
a_n = 5n^2 + 3n + 1, \qquad b_n = n^2
$$

Obviously

$$
|a_n| \le 5n^2 + 3n^2 + n^2 = 9n^2 \qquad \forall n \ge 1 \ (= k_0)
$$

$$
\Rightarrow \quad a_n = O(n^2), \quad \text{i.e.} \quad \boxed{a_n = O(b_n)}
$$

Note however that we also have

$$
|b_n| = n^2 \le 5n^2 + 3n + 1
$$

so in fact $\boxed{b_n = O(a_n)}$.

In such a case we say $a_n = \Theta(b_n)$: "$a_n$ is of **sharp order** $b_n$". Here $a_n$ is dominated by $b_n$ **and** $b_n$ is dominated by $a_n$.

## Definition 2 — sharp order (big $\Theta$)

We say $a_n = \Theta(b_n)$ (**of sharp order**) if $a_n = O(b_n)$ and $b_n = O(a_n)$. The latter means:

$$
\exists M_1\, \exists k_0\ \forall n \ge k_0: \quad |a_n| \le M_1 |b_n|
$$

$$
\text{and} \quad \exists M_2\, \exists \tilde{k}_0\ \forall n \ge \tilde{k}_0: \quad |b_n| \le M_2 |a_n|
$$

For the case when both $a_n, b_n \ge 0$ it suffices to show the existence of $M_1, M_2 > 0$ and $k_0$ such that for all $n \ge k_0$:

$$
\boxed{\ 0 \le M_1 b_n \le a_n \le M_2 b_n\ } \quad \iff \quad a_n = \Theta(b_n) \ \text{ and } \ b_n = \Theta(a_n)
$$

## Example — polynomials

$$
a_n = a_t n^t + a_{t-1} n^{t-1} + \ldots + a_2 n^2 + a_1 n + a_0, \qquad t \in \mathbb{N},\ a_t \neq 0,\ a_t, a_{t-1}, \ldots, a_0 \in \mathbb{R}
$$

Since $|x + y| \le |x| + |y|$:

$$
\begin{aligned}
|a_n| &= |a_t n^t + a_{t-1} n^{t-1} + \ldots + a_2 n^2 + a_1 n + a_0| \\
&\le |a_t n^t| + |a_{t-1} n^{t-1}| + \ldots + |a_2 n^2| + |a_1 n| + |a_0| \\
&= n^t |a_t| + n^{t-1} |a_{t-1}| + \ldots + n^2 |a_2| + n |a_1| + |a_0| \\
&\le n^t |a_t| + n^t |a_{t-1}| + \ldots + n^t |a_2| + n^t |a_1| + n^t |a_0| \\
&= n^t \big( \underbrace{|a_t| + |a_{t-1}| + \ldots + |a_2| + |a_1| + |a_0|}_{M} \big) = M n^t
\end{aligned}
$$

$$
\Rightarrow \quad |a_n| \le M n^t \quad \forall n \ge 1 \ (= k_0) \quad \Rightarrow \quad \boxed{a_n = O(n^t)}
$$

On the other hand $|a_n| \gtrsim n^t |a_t|$, so

$$
\Rightarrow \quad \boxed{a_n = \Theta(n^t)}
$$

> **Note:** the teacher wrote the lower bound simply as $|a_n| \ge n^t |a_t|$; strictly speaking this holds only up to a constant — e.g. $|a_n| \ge \frac{|a_t|}{2}\, n^t$ for all sufficiently large $n$ — which is exactly what $\Theta$ requires.

In particular, if $a_n$ does not explode at all, i.e. $|a_n| \le M \cdot 1$ (a **bounded sequence**), then

$$
\boxed{a_n = O(1)}
$$

## Typical rates of growth

In computer science we have typical rates of growth for $a_n$. If $a_n$ is of order:

| Order | Name |
|---|---|
| $O(1)$ | constant |
| $O(\lg_2 n)$ | logarithmic |
| $O(\sqrt{n})$ | square-root |
| $O(n)$ | linear order |
| $O(n \lg_2 n)$ | "$n \lg_2 n$" |
| $O(n^2)$ | quadratic |
| $O(n^3)$ | cubic |
| $O(n^4)$ | quartic |
| $O(n^m)$ | polynomial, $m = 0, 1, 2, 3, \ldots$ |
| $O(n^\alpha)$ | power $\alpha$, $\alpha \in \mathbb{R}$, $\alpha \ge 1$ |
| $O(a^n)$ | exponential, $a > 1$ |
| $O(n!)$ | factorial |

These are the **most common orders** in computer science, but one can put them on an axis from left to right to symbolize which one is slower and which one is faster:

![[lec07a_p06_growth-orders-axis.svg]]

Here $n \lg_2 n$ sits between $n$ and $n^2$ because

$$
1 \le \lg_2 n \le n \quad (n \ge 2) \qquad \Rightarrow \qquad n \le n \lg_2 n \le n^2
$$

One can squeeze $\infty$-many other orders onto the axis:

- e.g. between $n^2$ and $n^3$: $\ n^2 \lg_2 n$;
- e.g. between $n$ and $n^2$: $\ n^\alpha$ for any $1 < \alpha < 2$ (say $n^{3/2}$ or $n^{1.9}$).

To see the explosion of growth it is good to look at the graphs of the functions. Say $a_n = O(n^2)$ — it means $a_n$ is below $n^2$ (but it still can be below $n$); the dots mark where the values $a_n$ can sit under the parabola:

![[lec07a_p06_O-n2-parabola.svg|280]]

If $a_n = \Theta(n^2)$, then it is between two parabolas $M_1 n^2$ and $M_2 n^2$:

![[lec07a_p06_theta-n2-parabolas.svg|290]]

## Useful properties

Some useful properties (not difficult to prove):

$$
a_n = O(b_n) \ \Rightarrow \ c \cdot a_n = O(b_n)
$$

$$
\left.
\begin{aligned}
a_n &= O(b_n) \\
c_n &= O(b_n)
\end{aligned}
\right\} \ \Rightarrow \ a_n \pm c_n = O(b_n)
$$

$$
\left.
\begin{aligned}
a_n &= O(b_n) \\
c_n &= O(d_n)
\end{aligned}
\right\} \ \Rightarrow \ a_n \cdot c_n = O(b_n \cdot d_n)
$$

$$
\left.
\begin{aligned}
a_n &= O(b_n) \\
b_n &= O(c_n)
\end{aligned}
\right\} \ \Rightarrow \ a_n = O(c_n)
$$

$$
\left.
\begin{aligned}
a_n &= O(c_n) \\
b_n &= O(d_n)
\end{aligned}
\right\} \ \Rightarrow \ a_n \pm b_n = O\big(\max\{c_n, d_n\}\big)
$$

**Note that** $a_n = O(b_n)$ only means that $a_n$ is bounded from above by $b_n$ — but $a_n$ can still be of lower order. On the sketch below $a_n = O(b_n)$, but also $a_n = O(c_n)$ and $a_n = O(d_n)$:

![[lec07a_p07_bound-line.svg|420]]

In particular:

$$
\boxed{\ a_n = O(b_n) \ \Rightarrow \ a_n = O(d_n)\ } \qquad \forall\, d_n \ \text{such that} \ b_n \le d_n
$$

If $a_n = O(b_n)$, then for $c_n \le b_n$:

- **(i)** $a_n$ *may* be of order $O(c_n)$, or
- **(ii)** $a_n$ *may not* be of order $O(c_n)$.

All depends on whether $a_n = \Theta(b_n)$:

- if **yes**, then (i) is false and (ii) is true;
- if **no**, then (i) is true and (ii) is false.

> **Note:** this dichotomy is the teacher's shorthand: when $a_n = \Theta(b_n)$, no order strictly smaller than $b_n$ can dominate $a_n$; when $a_n$ is not of sharp order $b_n$, the upper estimate can be improved.

## Example: $a_n = \ln(n!)$

Let us analyze an example. Since $\ln(x \cdot y) = \ln x + \ln y$:

$$
a_n = \ln(n!) = \ln\big[1 \cdot 2 \cdot 3 \cdots (n-1) \cdot n\big]
$$

$$
a_n = \ln 1 + \ln 2 + \ln 3 + \ldots + \ln(n-1) + \ln n
$$

### (i) Naive estimate

For $x \in [0, n]$ we have (a) $\ln x \le \ln n \le n$, and in general (b) $\ln x \le x$ — the logarithm graph stays below the line $y = x$:

![[lec07a_p08_lnx-below-x.svg|230]]

$$
0 \le a_n \stackrel{\text{a)}}{=} \ln 1 + \ldots + \ln n \le \underbrace{n + n + \ldots + n}_{n\text{-times}} = n^2
$$

$$
\Rightarrow \quad \boxed{a_n = O(n^2)} \quad (\bullet)
$$

Can we improve it?

### (ii) More intricate approach

$$
0 \le a_n = \ln 1 + \ln 2 + \ln 3 + \ldots + \ln n \le 1 + 2 + \ldots + n
$$

$$
= \frac{(1+n)\,n}{2} = \frac{1}{2}(n^2 + n) \le \frac{1}{2}(n^2 + n^2) = n^2
$$

$$
\Rightarrow \quad \boxed{a_n = O(n^2)} \quad (\bullet\bullet)
$$

So a more intricate approach does **not** improve the previous result $(\bullet)$.

### (iii) More advanced approach

This does **not** mean that $a_n = \Theta(n^2)$ (i.e. that we cannot improve this estimate).

Usually we now plot the points $(n, a_n)$ and compare them with the plotted graphs $x^2$, $x^{3/2}$, $x$, $x \lg_2 x$, $\lg_2 x$ (in Matlab or in Mathematica). What we will notice is that the points $(n, a_n)$ are well below $(n, K n^2)$ — in fact the cloud is close to $x \lg_2 x$.

We shall prove that now by resorting to **integration**. We show:

$$
(\text{A}) \quad a_n \le M_2\, n \lg_2 n \ \ \forall n \ge k_2, \qquad (\text{AA}) \quad M_1\, n \lg_2 n \le a_n \ \ \forall n \ge k_1
$$

$$
\Downarrow
$$

$$
\boxed{a_n = \Theta(n \lg_2 n)}
$$

### Upper estimate (A) — rectangles below $\ln x$

We compare the sum with an integral. The rectangles $S_1, S_2, \ldots, S_{n-1}$ (of width $1$, with heights $\ln 2, \ln 3, \ldots, \ln n$) fit below the graph of $y(x) = \ln x$:

![[lec07a_p10_lnx-upper-rectangles.svg]]

The areas of $S_1, S_2, \ldots, S_{n-1}$ are smaller than the area bounded by the $OX$ axis ($1 \le x \le n+1$) and the graph of $y(x) = \ln x$:

![[lec07a_p10_area-under-lnx.svg|210]]

Let $A_i =$ area of $S_i$. We have

$$
A_1 + A_2 + \ldots + A_{n-1} \le \int_1^{n+1} \ln x \, dx
$$

$$
(3-2)\ln 2 + (4-3)\ln 3 + \ldots + \big[(n+1) - n\big]\ln n \le \int_1^{n+1} \ln x \, dx
$$

$$
\underbrace{\ln 1}_{=\,0} + \ln 2 + \ln 3 + \ldots + \ln n \le \int_1^{n+1} \ln x \, dx
$$

(as $\ln 1 = 0$), i.e.

$$
(*) \quad a_n \le \int_1^{n+1} \ln x \, dx
$$

We use **integration by parts** (I.P.) to compute the right-hand side of $(*)$. Recall the I.P. formula:

$$
\int_a^b f'(t)\, g(t)\, dt = f(t)\, g(t) \Big|_a^b - \int_a^b f(t)\, g'(t)\, dt
$$

So:

$$
\begin{aligned}
\int_1^{n+1} \ln x \, dx &= \int_1^{n+1} x' \ln x \, dx \stackrel{\text{I.P.}}{=} x \ln x \Big|_1^{n+1} - \int_1^{n+1} x \cdot (\ln x)' \, dx \\
&= \big( (n+1)\ln(n+1) - 1 \cdot \underbrace{\ln 1}_{=\,0} \big) - \int_1^{n+1} x \cdot \tfrac{1}{x} \, dx \\
&= n \ln(n+1) + \ln(n+1) - x \Big|_1^{n+1} \\
&= n \ln(n+1) + \ln(n+1) - (n + 1 - 1) \\
&= n \ln(n+1) + \ln(n+1) - n
\end{aligned}
$$

Hence

$$
\begin{aligned}
0 \le a_n &\le n \ln(n+1) + \ln(n+1) - n \\
&\le n \ln(n+1) + n \ln(n+1) \\
&= 2n \ln(n+1) \\
&\le 2n \cdot 2 \ln n \\
&= 4 n \ln n \qquad \forall n \ge k_2
\end{aligned}
$$

Here we used $\ln(x+1) \le 2 \ln x$ for large $x$; indeed:

$$
\ln(x+1) \le 2\ln x \iff e^{\ln(x+1)} \le e^{2\ln x} \iff x + 1 \le x^2
$$

$$
x^2 - x - 1 = 0, \quad \Delta = 1 + 4 = 5, \quad x_{1,2} = \frac{1 \pm \sqrt{5}}{2}; \quad \text{we take } x_2 = \frac{1}{2} + \frac{\sqrt{5}}{2}
$$

The roots $x_1, x_2$ of the parabola $x^2 - x - 1$:

![[lec07a_p11_x2-vs-x-plus-1.svg|185]]

So for $x \ge \frac{1}{2} + \frac{\sqrt{5}}{2}$: $\ x^2 \ge x + 1 \iff \ln(x+1) \le 2\ln x$; and $k_2$ is any integer $> \frac{1}{2} + \frac{\sqrt{5}}{2}$.

Now convert to the base-2 logarithm:

$$
\ln n = \log_e n = \log_e 2 \cdot \lg_2 n = \ln 2 \cdot \lg_2 n
$$

$$
\Downarrow
$$

$$
0 \le a_n \le \underbrace{4 \ln 2}_{M_2} \cdot\, n \lg_2 n
$$

$$
\Downarrow
$$

$$
\boxed{a_n = O(n \lg_2 n)} \quad (1)
$$

So we now have a better estimate instead of $a_n = O(n^2)$:

![[lec07a_p12_improved-estimate-axis.svg|540]]

### Lower estimate (AA) — can we improve (iii) more?

**Question:** can we improve (iii) more? (i.e. can we shift the estimate further to the left on the axis?)

We will show now that **NO!** — i.e.

$$
\exists M_1\ \exists K_1: \quad M_1\, n \lg_2 n \le a_n \qquad \forall n \ge K_1
$$

**Again** we compare with an integral — but now the rectangles $S_1, S_2, \ldots, S_{n-1}$ (heights $\ln 2, \ln 3, \ldots, \ln n$) stick out **above** the graph of $y(x) = \ln x$:

![[lec07a_p12_lnx-lower-rectangles.svg]]

The areas of $S_1, S_2, \ldots, S_{n-1}$ summed up **exceed** the area under $y(x) = \ln x$ over $1 \le x \le n$:

![[lec07a_p12_area-under-lnx-n.svg|210]]

$$
A(S_1) + A(S_2) + \ldots + A(S_{n-1}) \ge \int_1^n \ln x \, dx
$$

$$
(2-1)\ln 2 + (3-2)\ln 3 + \ldots + \big[n - (n-1)\big]\ln n \ge \int_1^n \ln x \, dx
$$

$$
\underbrace{\underbrace{\ln 1}_{=\,0} + \ln 2 + \ln 3 + \ldots + \ln n}_{a_n} \ge x \ln x \Big|_1^n - \int_1^n x\,(\ln x)'\, dx
$$

$$
a_n \ge n \ln n - \underbrace{\ln 1}_{=\,0} - x \Big|_1^n = n \ln n - n + 1
$$

$$
a_n \ge n \ln n - n + 1 \ge n \ln n - n
$$

$$
a_n \ge n(\ln n - 1) \qquad \forall n \in \mathbb{N} \quad (\blacksquare)
$$

But $\ln x - 1 \ge \frac{1}{2} \ln x$ for $x \ge x_0$; indeed:

$$
e^{\ln x - 1} \ge \big(e^{\ln x}\big)^{1/2} \iff \frac{x}{e} \ge \sqrt{x} \iff \frac{x^2}{e^2} \ge x \iff x^2 - e^2 x \ge 0 \iff x(x - e^2) \ge 0 \iff x \ge e^2
$$

so we take $K_1 \ge e^2$. Then

$$
a_n \ge n(\ln n - 1) \ge n \cdot \frac{1}{2} \ln n \qquad \text{for } n \ge K_1 \ (\text{fixed})
$$

$$
a_n \ge \underbrace{\frac{1}{2} \ln 2}_{M_1} \cdot\, n \lg_2 n \qquad \forall n \ge K_1
$$

$$
\Downarrow
$$

$$
\boxed{a_n \ge M_1\, n \lg_2 n} \quad (2)
$$

$$
\Rightarrow \quad \text{by (1) and (2):} \qquad \boxed{\ a_n = \Theta(n \lg_2 n)\ }
$$

## The general path of asymptotic estimation

The example from above shows the general/usual path when estimating the asymptotics for $a_n$ coming from a specific algorithm:

- **a)** finding a "terse" (first, coarse) estimate;
- **b)** searching for improvements;
- **c)** ultimately proving sharpness.

The estimates move left along the axis of orders until the sharp order is reached:

![[lec07a_p14_estimation-path.svg|560]]

Usually more and more advanced techniques are needed in passing from a) $\Rightarrow$ c). In a search for improvements one can use computer plots of $(n, a_n)$ against plots of different $(n, f(n))$.

There are **no standard techniques** for all algorithms.

At best, if we can calculate $a_n$ **exactly** (see e.g. Bubble sort), this gives sharpness in one step, without passing through a) and b).

However, finding an exact formula is sometimes difficult. And also, when we compare 2 algorithms for performing a certain task, we are interested in how they differ for sufficiently large $n$ (modulo constant). For the latter,

$$
a_n = O(f(n)) \ \text{ for Algorithm 1} \qquad \text{and} \qquad \hat{a}_n = O(g(n)) \ \text{ for Algorithm 2}
$$

— these measurements suffice!

The criterion of using integration is nice, but it can be applied only if the sequence $a_n$ **is a sum** (which usually is the case in computer science — see e.g. Bubble sort). However, we may have situations in the analysis of algorithms where this technique cannot be applied. Then the analysis each time involves a new argument, adapted to the problem.

## Remark — counting and worst cases

**a)** Quite often the analysis of algorithms involves counting all possibilities (the **search space**). E.g. if we want to find a maximal element (of a list), the question is about the complexity of the **exhaustive search** algorithm: we simply take all $f(x_i)$.

**b)** The other case is when we want to assess the **worst-case scenario**, so that we can compare the performance of the algorithms against the worst case.

All of this is related to the analysis and complexity of algorithms and resorts to either:

- **exact** computation of all possibilities (exact sums, combinatorial counting), or
- **approximate** computation of all possibilities (inequalities, counting)

— they all involve determining $a_n = \,?$
