# Discrete Mathematics — Lecture 10: Recurrences and Generating Functions

> Translated from the handwritten lecture notes `pdf/lect10rekurencje-funk.pdf`.

> The original handwritten pages are headed "Algorithms & Complexity — Lecture 7"; in this course's file numbering it is Lecture 10.

## The method of generating functions

We present a method called **the method of generating functions (GF)**, which may apply to:

- **non-linear** difference schemes,
- schemes of **non-fixed order**,
- **convolved** schemes having more than one recursive scheme.

It has one main weakness — to be explained later.

For a sequence we take the mapping

$$
\{a_n\}_{n \in \mathbb{N}} \;\longmapsto\; f(x) = \sum_{n=0}^{\infty} a_n x^n \qquad (\blacksquare)
$$

— a **function series**:

$$
f(x) = \lim_{n \to \infty} \sum_{k=0}^{n} a_k x^k
$$

We know from calculus that for such **power series**

$$
\exists\, R \geq 0 \quad \forall\, \bar{x} \in (-R, R) \quad \sum_{n=0}^{\infty} a_n \bar{x}^{\,n} \ \text{ is convergent}
$$

(for a fixed $\bar{x}$ this is a **number series**).

## Radius of convergence

$R$ — called the **radius of convergence**. It can be:

- $0 < R < \infty$, or
- $R = 0$ (then only for $\bar{x} = 0$ we have convergence), or
- $R = \infty$ (then for all $x \in \mathbb{R}$ we have convergence).

> The notes write "$R = -\infty$" in the last case — an evident slip for $R = \infty$.

In general, if

$$
\lim_{n} \left| \frac{a_{n+1}}{a_n} \right| = g_1 \quad (1)
\qquad \text{or if} \qquad
\lim_{n} \sqrt[n]{|a_n|} = g_2 \quad (2)
$$

then

$$
R = \frac{1}{g}
$$

(if both $g_1$ and $g_2$ exist, they satisfy $g_1 = g_2$).

## Properties of power series

For $f(x) = \sum_{n=0}^{\infty} a_n x^n$, $x \in (-R, R)$:

$$
f'(x) = \sum_{n=1}^{\infty} a_n \cdot n\, x^{n-1}
$$

If

$$
g_1(x) = \sum_{n=1}^{\infty} a_n x^n, \quad x \in (-R_1, R_1),
\qquad
g_2(x) = \sum_{n=1}^{\infty} \hat{a}_n x^n, \quad x \in (-R_2, R_2),
$$

then

$$
g_1(x) + g_2(x) = \sum_{n=1}^{\infty} (a_n + \hat{a}_n)\, x^n,
\qquad x \in \bigl(-\min(R_1, R_2),\ \min(R_1, R_2)\bigr)
$$

## Example — a linear recurrence solved by GF

We illustrate the method of generating functions in several examples.

**Example:**

$$
(*) \quad
\begin{cases}
a_n - 3a_{n-1} + 2a_{n-2} = 0 \\
a_0 = 0, \qquad a_1 = 1
\end{cases}
$$

A linear scheme of order 2, homogeneous, with constant coefficients and with initial values $\Rightarrow$ it has a **unique solution** $\{a_n\}_{n \geq 0}$ satisfying $(*)$.

Take $\{a_n\}_{n \geq 0} \longmapsto f(x) = \sum_{k=0}^{\infty} a_k x^k$:

$$
\begin{aligned}
f(x) &= a_0 + a_1 x + a_2 x^2 + \ldots + a_n x^n + \ldots \\
-3x\, f(x) &= -3x \sum_{k=0}^{\infty} a_k x^k = -3a_0 x - 3a_1 x^2 - \ldots - 3a_{n-1} x^n - \ldots \\
2x^2 f(x) &= 2x^2 \sum_{k=0}^{\infty} a_k x^k = 2a_0 x^2 + 2a_1 x^3 + \ldots + 2a_{n-2} x^n + \ldots
\end{aligned}
$$

Thus we have:

$$
\begin{aligned}
f(x) - 3x f(x) + 2x^2 f(x)
&= a_0 + a_1 x + a_2 x^2 + a_3 x^3 + a_4 x^4 + \ldots \\
&\quad - 3a_0 x - 3a_1 x^2 - 3a_2 x^3 - 3a_3 x^4 - \ldots \\
&\quad + 2a_0 x^2 + 2a_1 x^3 + 2a_2 x^4 + \ldots \\[6pt]
&= a_0 + a_1 x - 3a_0 x + x^2 (a_2 - 3a_1 + 2a_0) + x^3 (a_3 - 3a_2 + 2a_1) \\
&\quad + x^4 (a_4 - 3a_3 + 2a_2) + \ldots + x^n (a_n - 3a_{n-1} + 2a_{n-2}) + \ldots
\end{aligned}
$$

But all the expressions in (brackets) **vanish**, as $a_n$ satisfies $(*)$. Thus

$$
(\circ) \quad f(x) - 3x f(x) + 2x^2 f(x) = a_0 + a_1 x - 3a_0 x
$$

Since $a_0 = 0$ and $a_1 = 1$, in the function equation $(\circ)$ we compute $f(x)$ to get:

$$
\boxed{\; f(x) = \frac{x}{1 - 3x + 2x^2} \;}
$$

## The general trick — and its weakness

The general trick is now:

$$
\sum_{n=1}^{\infty} a_n x^n \;=\; f(x) \;=\; \sum_{n=1}^{\infty} b_n x^n
$$

On the left — as we assumed, $f(x)$ is built from $\{a_n\}_{n \geq 0}$ (see $\blacksquare$); on the right — we know that each $f$ can be expanded by the Taylor formula:

$$
b_n = \frac{f^{(n)}(0)}{n!}
$$

As $f$ has a **unique** series representation, then

$$
\boxed{\; a_n = b_n \;}
$$

**However**, there is a weakness in the above:

- $a_n$ is given recursively in $(*)$;
- to compute $b_n = \dfrac{f^{(n)}(0)}{n!}$ we need recursively

$$
f^{(n)}(0) = \Bigl( f^{(n-1)}(0) \Bigr)' \qquad (*\!*\!*)
$$

  which is also relying on previous derivatives — and it is hard to see some general pattern;
- we do not want to replace one recursion $(*)$ with another one $(*\!*\!*)$ once we transform the implicit form into the explicit form.

Sometimes $\dfrac{f^{(n)}(0)}{n!}$ can be done **without** recursion, or with a recursion in which we see a general pattern which can be proved by mathematical induction. Indeed, in our example this is the case:

$$
f(x) = \frac{x}{1 - 3x + 2x^2} \qquad (\text{a rational function})
$$

We calculate $\Delta$ for the denominator, as we want to factorize it:

$$
\Delta = 1, \qquad x_1 = \tfrac{1}{2} \ \text{ and } \ x_2 = 1
$$

$$
1 - 3x + 2x^2 = 2\left(x - \tfrac{1}{2}\right)(x - 1) = (2x - 1)(x - 1) = (1 - 2x)(1 - x)
$$

Thus

$$
f(x) = \frac{x}{(1 - x)(1 - 2x)}
$$

The next step is to use the so-called **partial fraction decomposition**.

## Partial fraction decomposition

$$
\frac{x}{(1-x)(1-2x)}
= \frac{A}{1-x} + \frac{B}{1-2x}
= \frac{(1-2x)A + (1-x)B}{(1-x)(1-2x)}
= \frac{x(-2A - B) + (A + B)}{(1-x)(1-2x)}
\qquad \forall x
$$

$$
1 \cdot x + 0 = x(-2A - B) + (A + B) \qquad \forall x
$$

Two polynomials are the same if their coefficients are equal:

$$
\begin{cases} -2A - B = 1 \\ A + B = 0 \end{cases}
\;\Rightarrow\; A = -B
\;\Rightarrow\; 2B - B = 1
\;\Rightarrow\; \boxed{\, B = 1, \quad A = -1 \,}
$$

Hence

$$
f(x) = \frac{-1}{1-x} + \frac{1}{1-2x}
$$

Taking into account that the sum of a geometric progression for $|q| < 1$ amounts to

$$
\frac{1}{1-q} = 1 + q^1 + q^2 + q^3 + \ldots + q^n + \ldots
$$

we get

$$
\begin{aligned}
-\frac{1}{1-x} &= -\bigl(1 + x + x^2 + x^3 + \ldots + x^n + \ldots\bigr) \\
\frac{1}{1-2x} &= 1 + 2^1 x^1 + 2^2 x^2 + 2^3 x^3 + \ldots + 2^n x^n + \ldots
\end{aligned}
$$

Consequently

$$
\begin{aligned}
f(x) &= -\bigl(1 + x + x^2 + \ldots + x^n + \ldots\bigr)
      + \bigl(1 + 2x + 2^2 x^2 + \ldots + 2^n x^n + \ldots\bigr) \\
&= x + (2^2 - 1)x^2 + (2^3 - 1)x^3 + \ldots + (2^n - 1)x^n + \ldots \\
&= \underbrace{(2^0 - 1)}_{=\,0} x^0 + (2^1 - 1)x^1 + (2^2 - 1)x^2 + (2^3 - 1)x^3 + \ldots + (2^n - 1)x^n + \ldots \\
&= \sum_{n=0}^{\infty} (2^n - 1)\, x^n
\end{aligned}
$$

So, on one hand and on the other hand,

$$
\sum_{n=0}^{\infty} a_n x^n \;=\; f(x) \;=\; \sum_{n=0}^{\infty} (2^n - 1)\, x^n
$$

and comparing coefficients of both series:

$$
\boxed{\; a_n = 2^n - 1, \qquad a_0 = 0 \;}
$$

$$
\boxed{\; a_n = \Theta(2^n) \;}
$$

## Example — a convolved system of two recurrences (a particle reactor)

We give an example in which the method of generating functions can be applied to a **convolved system of 2 recursive schemes**.

**Example:** In a reactor of particles let:

- $a_n$ — denote the number of **high-energy neutrons** ($N_H$) at time $t_n$ — $n$ milliseconds passed from $t_0 = 0$;
- $b_n$ — denote the number of **low-energy neutrons** ($N_L$) at time $t_n$ — $n$ milliseconds passed from $t_0 = 0$.

Assume that at $t_0 = 0$ in the reactor we "inject":

- $a_0 = 1$ — one neutron of high energy,
- $b_0 = 0$ — no neutrons of low energy.

In physics:

**a)** A high-energy neutron, after hitting the atoms' kernel [nucleus] of the radioactive material, produces $2\,N_H$ and $1\,N_L$:

![[lec10_p09_high-energy-neutron.svg]]

**b)** A low-energy neutron ($N_L$), after hitting the atoms' kernel of the radioactive material, produces $1\,N_H$ and $1\,N_L$:

![[lec10_p10_low-energy-neutron.svg]]

a) and b) have the following recurrent representation at time $t_{n+1}$:

$$
a_{n+1} = \underbrace{2a_n}_{\substack{\text{contribution of } a_n \\ \text{neutrons at time } t_n}} + \underbrace{b_n}_{\substack{\text{contribution of } b_n \\ \text{neutrons at } t_n}}
\qquad\qquad
b_{n+1} = a_n + b_n
$$

So

$$
(\star) \quad
\begin{cases}
a_{n+1} = 2a_n + b_n, & a_0 = 1 \\
b_{n+1} = a_n + b_n, & b_0 = 0
\end{cases}
$$

> The notes write "$b_0 = 1$" next to the system, but this contradicts the injection described just above ($b_0 = 0$, no low-energy neutron) and the computation of $g(x)$ below, which uses $b_0 = 0$. We keep $b_0 = 0$.

If we use $c_n = (a_n, b_n)$, the scheme $(\star)$ **cannot** be transformed into a single scheme of the form

$$
c_{n+1} = \bar{k} \cdot c_n + \bar{k}_2
$$

## Solving the system with generating functions

$$
\{a_n\}_{n \geq 0} \;\longmapsto\; f(x) = \sum_{n=0}^{\infty} a_n x^n \quad (**)
\qquad\qquad
\{b_n\}_{n \geq 0} \;\longmapsto\; g(x) = \sum_{n=0}^{\infty} b_n x^n \quad (*\!*\!*)
$$

Multiply both recurrences by $x^{n+1}$:

$$
\begin{cases}
a_{n+1}\, x^{n+1} = 2a_n\, x^{n+1} + b_n\, x^{n+1} \\
b_{n+1}\, x^{n+1} = a_n\, x^{n+1} + b_n\, x^{n+1}
\end{cases}
$$

We take the infinite sum:

$$
\begin{cases}
\displaystyle \sum_{n=0}^{\infty} a_{n+1} x^{n+1} = 2 \sum_{n=0}^{\infty} a_n x^{n+1} + \sum_{n=0}^{\infty} b_n x^{n+1} \\[8pt]
\displaystyle \sum_{n=0}^{\infty} b_{n+1} x^{n+1} = \sum_{n=0}^{\infty} a_n x^{n+1} + \sum_{n=0}^{\infty} b_n x^{n+1}
\end{cases}
$$

Substituting $n + 1 = k$ (so $n = 0 \Rightarrow k = 1$ and $n = \infty \Rightarrow k = \infty$):

$$
\begin{cases}
\displaystyle \sum_{k=1}^{\infty} a_k x^k = 2x \sum_{n=0}^{\infty} a_n x^n + x \sum_{n=0}^{\infty} b_n x^n \\[8pt]
\displaystyle \sum_{k=1}^{\infty} b_k x^k = x \sum_{n=0}^{\infty} a_n x^n + x \sum_{n=0}^{\infty} b_n x^n
\end{cases}
$$

Renaming $k = n$ and using $(**)$ and $(*\!*\!*)$:

$$
\begin{cases}
\displaystyle \sum_{n=0}^{\infty} a_n x^n - a_0 x^0 = 2x\, f(x) + x\, g(x) \\[8pt]
\displaystyle \sum_{n=0}^{\infty} b_n x^n - b_0 x^0 = x\, f(x) + x\, g(x)
\end{cases}
$$

i.e.

$$
\begin{cases}
f(x) - 1 = 2x\, f(x) + x\, g(x) \\
g(x) = x\, f(x) + x\, g(x)
\end{cases}
$$

This is a system of 2 equations with 2 **unknown functions** $f(x)$ and $g(x)$. Upon solving it:

$$
\boxed{\; f(x) = \frac{1 - x}{x^2 - 3x + 1} \;}
\qquad\qquad
\boxed{\; g(x) = \frac{x}{x^2 - 3x + 1} \;}
$$

Again

$$
f(x) = \sum_{n=0}^{\infty} \frac{f^{(n)}(0)}{n!}\, x^n,
\qquad
g(x) = \sum_{n=0}^{\infty} \frac{g^{(n)}(0)}{n!}\, x^n,
$$

but it is difficult to see a general pattern for $\dfrac{f^{(n)}(0)}{n!}$ and $\dfrac{g^{(n)}(0)}{n!}$. So again we use a trick based on **partial fractions** applied to $f(x)$ and $g(x)$, which luckily are rational functions $\dfrac{W_n(x)}{W_m(x)}$ (ratios of polynomials).

## Partial fractions with irrational roots

So once $\Delta$ is computed for $x^2 - 3x + 1$, we get two roots (as $\Delta > 0$):

$$
\gamma = \frac{3 + \sqrt{5}}{2}, \qquad \delta = \frac{3 - \sqrt{5}}{2}
$$

and thus (for $ax^2 + bx + c = 0$ with $a = 1$):

$$
x^2 - 3x + 1 = (x - \delta)(x - \gamma) = (\delta - x)(\gamma - x)
$$

> **Footnote (from the notes):** as we try to solve real physical problems, the roots may be in an inconvenient form.

Partial fractions yield:

$$
f(x) = \frac{1 - x}{x^2 - 3x + 1}
= \frac{A}{\gamma - x} + \frac{B}{\delta - x}
= \left( \frac{5 + \sqrt{5}}{10} \right) \frac{1}{\gamma - x}
+ \left( \frac{5 - \sqrt{5}}{10} \right) \frac{1}{\delta - x}
$$

$$
g(x) = \frac{x}{x^2 - 3x + 1}
= \frac{\tilde{A}}{\gamma - x} + \frac{\tilde{B}}{\delta - x}
= \left( \frac{-5 - 3\sqrt{5}}{10} \right) \frac{1}{\gamma - x}
+ \left( \frac{-5 + 3\sqrt{5}}{10} \right) \frac{1}{\delta - x}
$$

Again, upon using the geometric series formula

$$
\frac{1}{1-q} = 1 + q + q^2 + \ldots + q^n + \ldots, \qquad |q| < 1
$$

(note that $\dfrac{1}{\gamma - x} = \dfrac{1}{\gamma} \cdot \dfrac{1}{1 - x/\gamma}$), we arrive at:

$$
\begin{aligned}
f(x) &= \sum_{n=0}^{\infty} \left[ \frac{5 + \sqrt{5}}{10} \cdot \frac{1}{\gamma} \left( \frac{x}{\gamma} \right)^n + \frac{5 - \sqrt{5}}{10} \cdot \frac{1}{\delta} \left( \frac{x}{\delta} \right)^n \right] \\
&= \sum_{n=0}^{\infty} \underbrace{\left[ \frac{5 + \sqrt{5}}{10} \left( \frac{3 + \sqrt{5}}{2} \right)^{-(n+1)} + \frac{5 - \sqrt{5}}{10} \left( \frac{3 - \sqrt{5}}{2} \right)^{-(n+1)} \right]}_{\hat{a}_n} x^n
= \sum_{n=0}^{\infty} \hat{a}_n x^n
\end{aligned}
$$

$$
\begin{aligned}
g(x) &= \sum_{n=0}^{\infty} \left[ \frac{-5 - 3\sqrt{5}}{10} \cdot \frac{1}{\gamma} \left( \frac{x}{\gamma} \right)^n + \frac{-5 + 3\sqrt{5}}{10} \cdot \frac{1}{\delta} \left( \frac{x}{\delta} \right)^n \right] \\
&= \sum_{n=0}^{\infty} \left[ \frac{-5 - 3\sqrt{5}}{10} \left( \frac{3 + \sqrt{5}}{2} \right)^{-(n+1)} + \frac{-5 + 3\sqrt{5}}{10} \left( \frac{3 - \sqrt{5}}{2} \right)^{-(n+1)} \right] x^n
= \sum_{n=0}^{\infty} \hat{b}_n x^n
\end{aligned}
$$

Thus

$$
\sum_{n=0}^{\infty} a_n x^n = \sum_{n=0}^{\infty} \hat{a}_n x^n \;\Rightarrow\; a_n = \hat{a}_n,
\qquad
\sum_{n=0}^{\infty} b_n x^n = \sum_{n=0}^{\infty} \hat{b}_n x^n \;\Rightarrow\; b_n = \hat{b}_n
$$

Hence

$$
\boxed{\;
\begin{aligned}
a_n &= \frac{5 + \sqrt{5}}{10} \left( \frac{3 + \sqrt{5}}{2} \right)^{-(n+1)} + \frac{5 - \sqrt{5}}{10} \left( \frac{3 - \sqrt{5}}{2} \right)^{-(n+1)} \\[6pt]
b_n &= \frac{-5 - 3\sqrt{5}}{10} \left( \frac{3 + \sqrt{5}}{2} \right)^{-(n+1)} + \frac{-5 + 3\sqrt{5}}{10} \left( \frac{3 - \sqrt{5}}{2} \right)^{-(n+1)}
\end{aligned}
\;}
$$

> Sanity check: since $\gamma\delta = 1$, we have $\gamma^{-(n+1)} = \delta^{\,n+1}$ and $\delta^{-(n+1)} = \gamma^{\,n+1}$; the formulas indeed give $a_0 = 1$, $b_0 = 0$, $a_1 = 2$, $b_1 = 1$.

## Closing remarks

Note that if

$$
f(x) = \frac{W_n(x)}{W_m(x)}, \qquad \deg\bigl(W_n(x)\bigr) = n, \quad \deg\bigl(W_m(x)\bigr) = m,
$$

then each polynomial can be factorized as a multiplication of **quadratic and linear polynomials** (this requires root finding for $W_m(x)$).

One can apply the method of generating functions to schemes of **non-fixed orders**, e.g.

$$
t(n) = t(0)\, t(n-1) + t(1)\, t(n-2) + \ldots + t(n-1)\, t(0)
$$

and exploit the **Cauchy formula for series multiplication**. To be discussed later.
