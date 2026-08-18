# Discrete Mathematics — Lecture 9: Recurrences (continued)

> Translated from the handwritten lecture notes `pdf/lect9rekurencjeEN-MatDysk.pdf`.

> **Note:** the handwritten pages are headed "Algorithms & Complexity — Lecture 6"; the material directly continues the previous lecture on linear recurrences (order-2 schemes and the space $V_\infty$ of solutions).

## Third-order homogeneous schemes ($k = 3$)

For $k = 3$ we have $\dim(V_\infty^3) = 3$.

We may have the scheme:

$$
(\clubsuit) \quad a_n = \alpha_1 a_{n-1} + \alpha_2 a_{n-2} + \alpha_3 a_{n-3}
$$

Substituting the trial solution $a_n = C r^n$ gives:

$$
(\bullet) \quad r^3 - \alpha_1 r^2 - \alpha_2 r - \alpha_3 = 0
$$

— the **characteristic equation** corresponding to the scheme $(\clubsuit)$.

### Case A — three different real roots

The equation $(\bullet)$ has 3 different real roots:

$$
r_1 \neq r_2 \neq r_3 \quad \text{(the multiplicity of each root is } 1\text{)}
$$

$$
\Rightarrow \quad (*) \quad a_n = C_1 r_1^n + C_2 r_2^n + C_3 r_3^n
$$

Here $r_1^n,\ r_2^n,\ r_3^n$ form a **basis of $V_\infty^3$**, and $a_n$ depends on 3 parameters.

Once $\hat{a}_0, \hat{a}_1, \hat{a}_2$ are given, we eliminate the 3 parameters from $(*)$:

$$
\begin{cases}
\hat{a}_0 = a_0 = C_1 r_1^0 + C_2 r_2^0 + C_3 r_3^0 \\
\hat{a}_1 = a_1 = C_1 r_1^1 + C_2 r_2^1 + C_3 r_3^1 \\
\hat{a}_2 = a_2 = C_1 r_1^2 + C_2 r_2^2 + C_3 r_3^2
\end{cases}
$$

and after computing $\tilde{C}_1, \tilde{C}_2, \tilde{C}_3$ from this system, the unique solution is:

$$
a_n = \tilde{C}_1 r_1^n + \tilde{C}_2 r_2^n + \tilde{C}_3 r_3^n
$$

### Case B — a double root and a single root

The roots of $(\bullet)$ are $r_1 = r_2 = r$ (multiplicity 2) and $r_3$ different from $r$ (multiplicity 1):

$$
(\spadesuit) \quad a_n = C_1 r^n + C_2\, n\, r^n + C_3 r_3^n
$$

If $\hat{a}_0, \hat{a}_1, \hat{a}_2$ are given, we can eliminate the parameters $C_1, C_2, C_3$ from:

$$
\begin{cases}
\hat{a}_0 = a_0 = C_1 r^0 + C_2 \cdot 0 \cdot r^0 + C_3 r_3^0 \\
\hat{a}_1 = a_1 = C_1 r^1 + C_2 \cdot 1 \cdot r^1 + C_3 r_3^1 \\
\hat{a}_2 = a_2 = C_1 r^2 + C_2 \cdot 2 \cdot r^2 + C_3 r_3^2
\end{cases}
$$

Once $\tilde{C}_1, \tilde{C}_2, \tilde{C}_3$ are computed, the unique solution to $(\spadesuit)$ with $\hat{a}_0, \hat{a}_1, \hat{a}_2$ given reads as:

$$
a_n = \tilde{C}_1 r^n + \tilde{C}_2\, n\, r^n + \tilde{C}_3 r_3^n
$$

### Case C — a triple root

The roots of $(\bullet)$ are $r_1 = r_2 = r_3 = r$ (multiplicity of the root is 3). Then:

$$
a_n = C_1 r^n + C_2\, n\, r^n + C_3\, n^2 r^n
$$

The solution depends on 3 parameters ($\dim V_\infty^3 = 3$). If $\hat{a}_0, \hat{a}_1$ and $\hat{a}_2$ are given, then we can eliminate $C_1, C_2, C_3$ from:

$$
\begin{cases}
\hat{a}_0 = a_0 = C_1 r^0 + C_2 \cdot 0 \cdot r^0 + C_3\, 0^2\, r^0 \\
\hat{a}_1 = a_1 = C_1 r^1 + C_2 \cdot 1 \cdot r^1 + C_3\, 1^2\, r^1 \\
\hat{a}_2 = a_2 = C_1 r^2 + C_2 \cdot 2 \cdot r^2 + C_3\, 2^2\, r^2
\end{cases}
$$

Again, upon eliminating and computing $\tilde{C}_1, \tilde{C}_2, \tilde{C}_3$, we obtain a unique solution:

$$
a_n = \tilde{C}_1 r^n + \tilde{C}_2\, n\, r^n + \tilde{C}_3\, n^2 r^n
$$

## Example

### a)

$$
(*) \quad a_n = 2a_{n-1} + 5a_{n-2} - 6a_{n-3}
$$

Once we substitute $a_n = c r^n$ we obtain a characteristic polynomial for $(*)$:

$$
r^3 - 2r^2 - 5r + 6 = (r-1)(r+2)(r-3) = 0
$$

**Attention:** finding roots of schemes of order $> 2$ generates a technical difficulty in finding roots of the corresponding characteristic polynomial. The higher the order $k$ is, the higher the order of the polynomial we have.

**Hints:** if we search for roots of a polynomial with integer coefficients, and the roots are to be integer, they must be **divisors of $a_0$**, where:

$$
W_n(x) = a_n x^n + a_{n-1} x^{n-1} + \ldots + a_1 x^1 + a_0
$$

Once you find one root $\bar{x}$, we divide $W_n(x)$ by $(x - \bar{x})$ and get a polynomial of degree $n-1$, i.e. $W_n(x) = (x - \bar{x}) \cdot \widehat{W}_{n-1}(x)$, and we continue the procedure.

So:

$$
(*) \quad a_n = C_1\, 1^n + C_2 (-2)^n + C_3\, 3^n
$$

If initial conditions are not given, then $a_n$ is of the form $(*)$ and depends on 3 parameters.

### b)

$$
\begin{cases}
b_{n+1} = 5b_n - 3b_{n-1} - 9b_{n-2} & (*) \\
b_0 = 6, \quad b_1 = -8, \quad b_2 = -22 & (**)
\end{cases}
$$

Characteristic equation:

$$
r^3 - 5r^2 + 3r + 9 = 0
$$

$$
(r+1)(r-3)^2 = 0
$$

$$
\Rightarrow \quad b_n = C_1 (-1)^n + C_2\, 3^n + C_3\, n\, 3^n \quad \text{— general solution to } (*)
$$

Taking into account $(**)$:

$$
\begin{cases}
C_1 + C_2 = 6 \\
-C_1 + 3C_2 + 3C_3 = -8 \\
C_1 + 9C_2 + 18C_3 = -22
\end{cases}
$$

$$
\Rightarrow \quad C_1 = 5, \qquad C_2 = 1, \qquad C_3 = -2
$$

hence:

$$
b_n = 5(-1)^n + 3^n - 2n\, 3^n
$$

## The general homogeneous scheme of order $k$

**In general:**

$$
(\clubsuit) \quad a_n = \sum_{i=1}^{k} c_i\, a_{n-i}
$$

$$
\Rightarrow \quad \text{characteristic polynomial} \quad W_k(r) = r^k - \sum_{i=1}^{k} c_i\, r^{k-i}
$$

> Margin note (for $k = 4$, the possible root patterns):
> $r_1 = r_2 = r_3 = r_4$; $\quad r_1 = r_2 \neq r_3 = r_4$; $\quad r_1 = r_2 = r_3 \neq r_4$; $\quad r_1 \neq r_2 \neq r_3 \neq r_4$.

### Theorem

Assume the characteristic polynomial $W_k(r)$ has $r_1, r_2, \ldots, r_{l_k}$ different roots (real or complex), each with the multiplicity $m_1, m_2, \ldots, m_{l_k}$ such that $m_1 + m_2 + \ldots + m_{l_k} = k$, so:

$$
W_k(r) = \prod_{i=1}^{l_k} (r - r_i)^{m_i}
$$

Then each solution to $(\clubsuit)$ is a linear combination of:

$$
\begin{aligned}
& r_1^n,\ n r_1^n,\ \ldots,\ n^{m_1 - 1} r_1^n \\
& r_2^n,\ n r_2^n,\ \ldots,\ n^{m_2 - 1} r_2^n \\
& \vdots \\
& r_{l_k}^n,\ n r_{l_k}^n,\ \ldots,\ n^{m_{l_k} - 1} r_{l_k}^n
\end{aligned}
$$

which depends on the parameters $C_1, C_2, \ldots, C_k$. If additionally $\hat{a}_0, \hat{a}_1, \ldots, \hat{a}_{k-1}$ are given, then $C_1, C_2, \ldots, C_k$ can be computed from the respective system of $k$ equations in $k$ unknowns.

## Inhomogeneous schemes

Still we do not know how to solve an **inhomogeneous** difference scheme — in particular we do not know how to handle the Tower of Hanoi problem:

$$
T_n = 2T_{n-1} + 1
$$

In general, how to deal with:

$$
(\mathrm{IH}) \quad a_n = \sum_{i=1}^{k} c_i\, a_{n-i} + f(n), \qquad f \not\equiv 0
$$

— an **inhomogeneous linear scheme with constant coefficients**.

We know (from the general scheme $a_n = f(a_{n-1}, \ldots, a_{n-k}, n)$ of the previous lecture) that if $a_0, a_1, \ldots, a_{k-1}$ are given, then $(\mathrm{IH})$ has a unique solution. So $a_n$ depends on $k$ parameters.

$$
(\mathrm{H}) \quad h_n = \sum_{i=1}^{k} c_i\, h_{n-i}
$$

— the corresponding **homogeneous scheme**; $h_n$ depends on $k$ parameters.

If $b_n$ is a special solution to $(\mathrm{IH})$, then:

$$
a_n = h_n + b_n \ \text{ satisfies } (\mathrm{IH}) \ \text{ and depends on } k \text{ parameters}
$$

And vice versa: if $a_n$ satisfies $(\mathrm{IH})$ (depends on $k$ parameters), we take $b_n$ a special solution:

$$
a_n = \underbrace{a_n - b_n}_{\text{satisfies } (\mathrm{H})\ (\text{depends on } k \text{ param.})} + \ b_n
$$

where $b_n$ is a special solution of $(\mathrm{IH})$.

### Theorem

Assume $b_n$ is a specific solution to $(\mathrm{IH})$. Then $a_n$ is a general solution of $(\mathrm{IH})$:

$$
\iff \quad a_n = b_n + h_n
$$

where $h_n$ is the general solution of $(\mathrm{H})$.

> The handwritten note reads "$b_n = a_n + h_n$" here — an evident slip of notation; the intended statement is $a_n = b_n + h_n$, as used everywhere else in the lecture.

## The solution method (steps I–IV)

**So**, for the problem:

$$
(1) \quad
\begin{cases}
a_n = \sum_{i=1}^{k} c_i\, a_{n-i} + f(n) \\
a_0, a_1, \ldots, a_{k-1} \ \text{[optionally given]}
\end{cases}
$$

**Solution for (1):**

**I.** Solve first the homogeneous system:

$$
h_n = \sum_{i=1}^{k} c_i\, h_{n-i}
$$

*Note: $h_n$ depends on $k$ parameters.*

**II.** Find a specific solution $b_n$ for the inhomogeneous equation:

$$
b_n = \sum_{i=1}^{k} c_i\, b_{n-i} + f(n)
$$

— by the **prediction method** (see later).

**III.** The general solution to the inhomogeneous equation:

$$
a_n = \underbrace{h_n}_{\text{I}} + \underbrace{b_n}_{\text{II}}
$$

— depends on the $k$ parameters $C_1, \ldots, C_k$.

**IV.** [Optionally] If $a_0, a_1, \ldots, a_{k-1}$ are given, the parameters $C_1, C_2, \ldots, C_k$ are computed from $k$ equations in $k$ unknowns.

## Example

$$
(\mathrm{IH}) \quad
\begin{cases}
a_n - 2a_{n-1} = 3^n & (*) \\
a_0 = 4 & (**)
\end{cases}
\qquad (k = 1)
$$

**I)** $r - 2 = 0$ (the characteristic equation for $h_n - 2h_{n-1} = 0$):

$$
r = 2 \qquad \Rightarrow \qquad h_n = C \cdot 2^n
$$

**II)** Prediction method: since $f(n) = 3^n$, we guess the solution (see the table at the end of this lecture) as:

$$
b_n = A\, 3^n, \qquad A = \,?
$$

(as $f(n) = \bar{r}^n \Rightarrow b_n = A \bar{r}^n$ if $\bar{r} \neq r$).

> **Attention:** $3 \neq r = 2$. If it were equal, we would treat it as multiplicity 2, i.e. take $b_n = A\, n\, 3^n$.

So we substitute into the inhomogeneous equation $(*)$:

$$
A\, 3^n - 2A\, 3^{n-1} = 3^n
$$

$$
3A - 2A = 3
$$

$$
A = 3
$$

$$
\Rightarrow \quad b_n = 3^{n+1}
$$

**III)**

$$
a_n = h_n + b_n = C\, 2^n + 3^{n+1}
$$

— the general form of the solution to $(*)$, depending on 1 parameter ($k = 1$ here).

**IV)** Since we know $a_0 = 4$:

$$
4 = a_0 = C\, 2^0 + 3^{0+1} \quad \Rightarrow \quad C = 1
$$

$\Rightarrow$ a unique solution to $(*)$ & $(**)$:

$$
a_n = 3^{n+1} + 2^n \qquad \Rightarrow \qquad a_n = \Theta(3^n)
$$

## Example — the Hanoi problem

$$
\begin{cases}
T_n = 2T_{n-1} + 1 & (*) \\
T_1 = 1 & (**)
\end{cases}
$$

> The handwritten page writes $T_n = 2T_{n+1} + 1$ here — a slip; the Hanoi recurrence (as on the earlier page) is $T_n = 2T_{n-1} + 1$.

**I)** $h_n - 2h_{n-1} = 0$:

$$
r - 2 = 0 \quad \text{(characteristic polynomial)}, \qquad r = 2 \quad \Rightarrow \quad h_n = C \cdot 2^n
$$

**II)** As $f(n) = 1 = 1^n = \bar{r}^n$, the prediction method suggests to search for:

$$
b_n = A \cdot \bar{r}^n = A \qquad (\text{as } \bar{r} \neq 2 = r)
$$

To compute $A$ we plug it into $(*)$:

$$
A = 2A + 1 \quad \Rightarrow \quad A = -1 \quad \Rightarrow \quad b_n = -1
$$

**III)**

$$
a_n = C\, 2^n - 1
$$

— the general solution to $(*)$, depending on 1 parameter $C$ ($k = 1$).

**IV)** By $(**)$: $1 = a_1 = C\, 2^1 - 1 \ \Rightarrow \ C = 1$:

$$
a_n = 2^n - 1
$$

For $n = 3$: $a_3 = 7$ — we had that, see Lecture 5.

## Example — a second-order inhomogeneous scheme

$$
(*) \quad a_n + a_{n-1} - 6a_{n-2} = \underbrace{2^n - 1}_{f(n)}
$$

**I]** $h_n + h_{n-1} - 6h_{n-2} = 0$:

$$
r^2 + r - 6 = 0, \qquad r_1 = -3 \ \text{ and } \ r_2 = 2
$$

$$
\Rightarrow \quad h_n = C_1 (-3)^n + C_2\, 2^n
$$

**II]** Prediction method: $f(n) = 2^n - 1^n$, with $\bar{r} = 2$ and $\bar{\bar{r}} = 1$.

As $\bar{r} = r_2 = 2$, the prediction method gives:

$$
b_n = C\, n\, 2^n + D \cdot 1^n
$$

We plug $b_n$ into $(*)$ to compute $C$ and $D$:

$$
\left( C n\, 2^n + D \right) + \left( C(n-1)\, 2^{n-1} + D \right) - 6\left( C(n-2)\, 2^{n-2} + D \right) = 2^n - 1
$$

$$
\Downarrow
$$

$$
\frac{5}{2}\, C\, 2^n - 4D = 2^n - 1 \qquad \forall n \in \mathbb{N}
$$

$$
\frac{5}{2}\, C = 1 \ \text{ and } \ -4D = -1
$$

$$
\Rightarrow \quad C = \frac{2}{5} \ \text{ and } \ D = \frac{1}{4}
$$

$$
b_n = \frac{2}{5}\, n\, 2^n + \frac{1}{4}
$$

— a special solution to the inhomogeneous equation $(*)$.

**III]**

$$
a_n = h_n + b_n = C_1 (-3)^n + C_2\, 2^n + \frac{2}{5}\, n\, 2^n + \frac{1}{4}
$$

As we do not have $\hat{a}_0$ & $\hat{a}_1$, this is a final answer; $a_n = \Theta(3^n)$.

## The prediction method — guessing $b_n$

We finish this lecture by hinting how to guess $b_n$ given special forms of $f(n)$ in:

$$
a_n = \sum_{i=1}^{k} c_i\, a_{n-i} + f(n)
$$

Note that in $f$ we may have $\bar{r}^n$ with $\bar{r} = r$ (one of the roots of the characteristic polynomial).

The teacher's table of predictions of the special solution $b_n$ for the standard forms of $f(n)$:

![[lec09en_p14_prediction-table.svg]]

Things get trickier if a summand $f_1(n)$ of $f(n)$ is a constant multiple of a solution of the associated homogeneous system. This happens e.g. when $f(n)$ contains a summand such as $c\, r^n$ or $(c_1 + c_2 n)\, r^n$ and $r$ is a root of the characteristic polynomial.

Then we multiply such a solution by the smallest power of $n$, say $n^s$, for which no summand of $n^s f_1(n)$ is a solution to the associated homogeneous relation.

## Summary of the method I–IV

The method based on steps I–IV has its:

**Advantages:**

- the algorithmic set of steps to transform the implicit form into the explicit form of the recursive formula, permitting e.g. to find $a_n = O(f(n))$.

**Disadvantages:**

- applies only to the schemes $a_n = \sum_{i=1}^{k} c_i\, a_{n-i} + f(n)$;
- relies on finding roots of the characteristic polynomial (which can be hard).
