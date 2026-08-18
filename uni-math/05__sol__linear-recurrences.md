# Discrete Mathematics — Exercises 5: Solutions

> Translated from the handwritten lecture notes `pdf/Ćwicz5-RozwiązaniaMatDysk.pdf`.

## Problem 1 — growth rates of sequences

### (i)

$$
a_n = 1^2 + 3^2 + \ldots + (2n-1)^2 = n(2n-1)(2n+1)\cdot\frac{1}{3}
$$

This should be shown by induction

$$
\Rightarrow \quad a_n = O(n^3)
$$

— **cubic complexity**.

### (ii)

$$
a_n = \sum_{i=1}^{n} 2^{i-1} = 2^n - 1 \quad \Rightarrow \quad a_n = O(2^n)
$$

— **exponential complexity**; in fact, since we have an *equality*, $a_n = \Theta(2^n)$.

### (iii)

$$
a_1 = 1, \quad a_2 = 2, \qquad a_n = a_{n-1} + a_{n-2}, \quad n \geq 3
$$

Inductively it is easy to show that

$$
(*) \qquad 0 \leq a_n \leq \left(\tfrac{7}{4}\right)^n \quad \Rightarrow \quad a_n = O\!\left(\left(\tfrac{7}{4}\right)^n\right)
$$

— **exponential complexity**. Maybe it is slower (we only have an *inequality* here) $\Rightarrow$ so is

$$
a_n \stackrel{?}{=} \Theta\!\left(\left(\tfrac{7}{4}\right)^n\right)
$$

That requires further analysis.

**I) Mathematical induction** (base cases):

$$
a_1 = 1 \leq \left(\tfrac{7}{4}\right)^1 \quad \text{OK}
\qquad
a_2 = 2 \leq \left(\tfrac{7}{4}\right)^2 = \tfrac{49}{16} \quad \text{OK}
$$

**II)** Inductive step — assume

$$
0 \leq a_{n-1} \leq \left(\tfrac{7}{4}\right)^{n-1}, \qquad
0 \leq a_{n-2} \leq \left(\tfrac{7}{4}\right)^{n-2}
\quad \Rightarrow \quad 0 \leq a_n \leq \left(\tfrac{7}{4}\right)^{n} \ ?
$$

Then:

$$
a_n = a_{n-1} + a_{n-2}
\ \stackrel{\text{II}}{\leq}\
\left(\tfrac{7}{4}\right)^{n-1} + \left(\tfrac{7}{4}\right)^{n-2}
= \left(\tfrac{4}{7}\right)\left(\tfrac{7}{4}\right)^{n} + \left(\tfrac{4}{7}\right)^{2}\left(\tfrac{7}{4}\right)^{n}
$$

$$
a_n \leq \left(\tfrac{7}{4}\right)^{n} \left[ \tfrac{16}{49} + \tfrac{28}{49} \right]
= \left(\tfrac{7}{4}\right)^{n} \cdot \tfrac{16+28}{49}
$$

$$
a_n \leq \left(\tfrac{7}{4}\right)^{n} \cdot \tfrac{44}{49} \leq \left(\tfrac{7}{4}\right)^{n}
\quad \Rightarrow \quad a_n \leq \left(\tfrac{7}{4}\right)^{n}
$$

Moreover:

$$
a_{n-1} \geq 0, \quad a_{n-2} \geq 0 \quad \Rightarrow \quad a_n = a_{n-1} + a_{n-2} \geq 0
$$

$$
0 \leq a_n \leq \left(\tfrac{7}{4}\right)^{n}
$$

By virtue of mathematical induction, $a_n$ satisfies $(*)$

$$
\Rightarrow \quad a_n = O\!\left(\left(\tfrac{7}{4}\right)^n\right)
$$

We have **not** shown sharpness here!!! $a_n \stackrel{?}{=} \Theta(\cdot)$

## Problem 2 — a double characteristic root

$$
a_{n+2} = 4a_{n+1} - 4a_n, \quad n \geq 0, \qquad a_0 = 1, \quad a_1 = 3
$$

We substitute $a_n = C r^n$:

$$
C r^{n+2} = 4 C r^{n+1} - 4 C r^{n} \quad \big/ \, : C r^n
$$

$$
r^2 = 4r - 4
$$

The **characteristic polynomial**:

$$
r^2 - 4r + 4 = 0
$$

$$
\Delta = 0 \quad \Rightarrow \quad r^{\pm} = 2 \ \text{ — a double root}
$$

$$
h_n = C_1 2^n + C_2\, n\, 2^n
$$

From the initial conditions:

$$
\begin{cases}
1 = a_0 = C_1 2^0 + C_2 \cdot 0 \cdot 2^0 \\
3 = a_1 = C_1 2^1 + C_2 \cdot 1 \cdot 2^1
\end{cases}
\qquad\Rightarrow\qquad
\begin{cases}
1 = C_1 \\
3 = 2C_1 + 2C_2
\end{cases}
\qquad C_1 = 1
$$

$$
3 = 2 \cdot 1 + 2C_2 \quad \Rightarrow \quad 1 = 2C_2 \quad \Rightarrow \quad C_2 = \tfrac{1}{2}
$$

$$
\Downarrow
$$

$$
a_n = 2^n + \tfrac{1}{2}\, n\, 2^n = 2^n + n\, 2^{n-1}
\qquad \Rightarrow \qquad a_n = \Theta(n\, 2^n)
$$

Check that $a_0 = 1$ and $a_1 = 3$, and that $a_n$ satisfies $a_{n+2} = 4a_{n+1} - 4a_n$.

> The handwriting says "check whether $a_1 = 1$ and $a_2 = 3$" — an index slip: the given initial conditions are $a_0 = 1$, $a_1 = 3$.

## Problem 3 — a third-order recurrence

$$
\begin{cases}
2a_{n+3} = a_{n+2} + 2a_{n+1} - a_n & (\blacktriangle) \\
a_0 = 0, \quad a_1 = 1, \quad a_2 = 2
\end{cases}
$$

We substitute $a_n = C r^n$ into $(\blacktriangle)$ and obtain the **characteristic equation**:

$$
\underbrace{2r^3 - r^2 - 2r + 1}_{W(r)} = 0
$$

$r = 1$ and $r = -1$ satisfy it $\Rightarrow$

$$
W(r) = (r-1)(r+1) \cdot W_1(r)
$$

Dividing through, we get $W_1(r) = 2r - 1$:

$$
\Rightarrow \quad W(r) = (2r-1)(r+1)(r-1), \qquad r = \tfrac{1}{2}
$$

**3 real roots:** $r = 1, -1, \tfrac{1}{2}$.

$$
a_n = C_1 (1)^n + C_2 (-1)^n + C_3 \left(\tfrac{1}{2}\right)^n
$$

$$
\begin{cases}
0 = a_0 = C_1 1^0 + C_2 (-1)^0 + C_3 \left(\tfrac{1}{2}\right)^0 \\
1 = a_1 = C_1 1^1 + C_2 (-1)^1 + C_3 \left(\tfrac{1}{2}\right)^1 \\
2 = a_2 = C_1 1^2 + C_2 (-1)^2 + C_3 \left(\tfrac{1}{2}\right)^2
\end{cases}
$$

3 equations, 3 unknowns:

$$
\begin{cases}
0 = C_1 + C_2 + C_3 \\
1 = C_1 - C_2 + \tfrac{1}{2} C_3 \\
2 = C_1 + C_2 + \tfrac{1}{4} C_3
\end{cases}
$$

After the computations!

$$
C_1 = \tfrac{5}{2}, \qquad C_2 = \tfrac{1}{6}, \qquad C_3 = -\tfrac{8}{3}
$$

$$
\Rightarrow \quad a_n = \tfrac{5}{2} + \tfrac{1}{6}(-1)^n + \left(-\tfrac{8}{3}\right)\left(\tfrac{1}{2}\right)^n
\quad \Rightarrow \quad a_n = O(1)
$$

## Problem 4 — nonhomogeneous recurrences

$$
(\bullet) \qquad a_{n+2} - 10a_{n+1} + 21a_n = f(n)
$$

**Roots:**

$$
r^2 - 10r + 21 = 0, \qquad \Delta > 0 \ \Rightarrow\ r_1 \neq r_2, \qquad r_1 = 3 \ \text{ and } \ r_2 = 7
$$

$$
a_n^{(h)} = C_1 3^n + C_2 7^n
$$

— the general form of the solution of the homogeneous equation

$$
a_{n+2} - 10a_{n+1} + 21a_n = 0.
$$

### Special solutions

**Case $f(n) = 5$** $\Rightarrow$ $A_0 = a_n^{s}$. We substitute $a_n^s$ into $(\bullet)$:

$$
\Rightarrow \quad A_0 - 10A_0 + 21A_0 = 5
\quad \Rightarrow \quad 12A_0 = 5
\quad \Rightarrow \quad A_0 = \tfrac{5}{12}
$$

**General solution:**

$$
a_n = C_1 3^n + C_2 7^n + \tfrac{5}{12}
$$

We do not have $a_0$ and $a_1$, so $C_1, C_2$ remain as parameters.

**Case $f(n) = 3n^2 - 2$** $\Rightarrow$ the special solution is

$$
A_3 n^2 + A_2 n + A_1 = a_n^{s}
$$

We substitute $a_n^s$ into $(\bullet)$. For all $n$:

$$
A_3(n+2)^2 + A_2(n+2) + A_1 + \left[A_3(n+1)^2 + A_2(n+1) + A_1\right](-10) + 21\left[A_3 n^2 + A_2 n + A_1\right] = 3n^2 + 0\cdot n + (-2)
$$

We group both sides by powers of $n$; for all $n$:

$$
K_1 n^2 + K_2 n + K_3 = 3n^2 + 0 \cdot n + (-2)
$$

So since this holds for all $n$, the coefficients are the same:

$$
\begin{cases}
K_1 = 3 \\
K_2 = 0 \\
K_3 = -2
\end{cases}
\quad \longrightarrow \quad \text{from these equations we compute } \tilde{A}_3, \tilde{A}_2 \text{ and } \tilde{A}_1
$$

**Without initial data:**

$$
a_n = C_1 3^n + C_2 7^n + \tilde{A}_3 n^2 + \tilde{A}_2 n + \tilde{A}_1
$$

If we add $a_0$ and $a_1$, then we eliminate the coefficients $C_1$ and $C_2$ from the system of equations:

$$
\begin{cases}
a_0 = C_1 3^0 + C_2 7^0 + \tilde{A}_3 \cdot 0^2 + \tilde{A}_2 \cdot 0 + \tilde{A}_1 \\
a_1 = C_1 3^1 + C_2 7^1 + \tilde{A}_3 \cdot 1^2 + \tilde{A}_2 \cdot 1 + \tilde{A}_1
\end{cases}
$$

We compute $\tilde{C}_1$ and $\tilde{C}_2$ and we have exactly one formula:

$$
\boxed{\ a_n = \tilde{C}_1 3^n + \tilde{C}_2 7^n + \tilde{A}_3 n^2 + \tilde{A}_2 n + \tilde{A}_1\ }
$$

## Problem 5 — resonance with the homogeneous solution

$$
(\bullet) \qquad a_{n+1} - a_n = n, \qquad a_2 = 1
$$

**Homogeneous part:**

$$
a_{n+1}^{(h)} - a_n^{(h)} = 0 \quad \Rightarrow \quad r - 1 = 0 \ \text{ — the characteristic equation}
$$

$$
\Rightarrow \quad a_n^{(h)} = C \cdot 1^n = C
$$

**Special solution** of the nonhomogeneous equation

$$
a_{n+1}^{(s)} - a_n^{(s)} = \underbrace{n}_{f(n)}
$$

is of the form $a_n^{(s)} = A_1 n + A_0$, but since the term $A_0$ is itself a solution of the homogeneous equation

$$
\Rightarrow \quad a_n^{(s)} = n^{\alpha}\,(n A_1 + A_0)
$$

chosen so that no term is a solution of that equation; here $\alpha = 1$:

$$
a_n^{(s)} = n(n A_1 + A_0) = n^2 A_1 + n A_0
$$

We substitute into $(\bullet)$:

$$
A_1(n+1)^2 + A_0(n+1) = A_1 n^2 + A_0 n + n
$$

$$
A_1 n^2 + (2A_1 + A_0)n + (A_1 + A_0) = A_1 n^2 + (A_0 + 1)n + 0
$$

$$
\Rightarrow \quad A_1 = A_1, \qquad 2A_1 + A_0 = A_0 + 1, \qquad A_1 + A_0 = 0
$$

$$
\Rightarrow \quad A_1 = \tfrac{1}{2}, \qquad A_0 = -\tfrac{1}{2}
$$

$$
a_n^{(s)} = \tfrac{1}{2} n^2 + \left(-\tfrac{1}{2}\right) n
$$

$\Rightarrow$ **general solution:**

$$
a_n = a_n^{(h)} + a_n^{s} = C + \tfrac{1}{2}n^2 + \left(-\tfrac{1}{2}\right)n
$$

$$
a_n = C + \tfrac{1}{2}\, n(n-1)
$$

We fix the parameter $C$ because we have $a_2 = 1$:

$$
1 = a_2 = C + \tfrac{1}{2}\cdot 2 \cdot (2-1)
$$

$$
\Downarrow
$$

$$
C = 0
$$

$$
\Downarrow
$$

$$
\boxed{\ a_n = \tfrac{1}{2}\, n(n-1)\ }
$$

## Problem 6 — a nonhomogeneous Fibonacci sequence

$$
a_n = a_{n-1} + a_{n-2} + 1, \qquad a_0 = 0, \quad a_1 = 1
$$

— the Fibonacci sequence, **nonhomogeneous**. $\Downarrow$ Transform it into closed (explicit) form.

$$
r^2 - r - 1 = 0
$$

— the **characteristic equation** (of $a_n = a_{n-1} + a_{n-2}$).

$$
\Delta = 5 > 0 \quad \Rightarrow \quad
r_1 = \frac{1+\sqrt{5}}{2}, \qquad r_2 = \frac{1-\sqrt{5}}{2}
$$

$$
a_n^{(h)} = C_1 \left(\frac{1+\sqrt{5}}{2}\right)^n + C_2 \left(\frac{1-\sqrt{5}}{2}\right)^n
$$

$$
f(n) = 1 \quad \Rightarrow \quad a_n^{(s)} = A
$$

We substitute into the nonhomogeneous equation:

$$
A = A + A + 1 \quad \Rightarrow \quad A = -1
$$

$$
\Rightarrow \quad a_n = C_1 \left(\frac{1+\sqrt{5}}{2}\right)^n + C_2 \left(\frac{1-\sqrt{5}}{2}\right)^n - 1
$$

We compute $C_1$ and $C_2$ from $a_0 = 0$ and $a_1 = 1$:

$$
\begin{cases}
0 = a_0 = C_1 \left(\dfrac{1+\sqrt{5}}{2}\right)^0 + C_2 \left(\dfrac{1-\sqrt{5}}{2}\right)^0 - 1 \\[2ex]
1 = a_1 = C_1 \left(\dfrac{1+\sqrt{5}}{2}\right)^1 + C_2 \left(\dfrac{1-\sqrt{5}}{2}\right)^1 - 1
\end{cases}
$$

$$
\Rightarrow \quad C_1 = \frac{3+\sqrt{5}}{2\sqrt{5}}, \qquad C_2 = \frac{\sqrt{5}-3}{2\sqrt{5}}
$$

$$
\Rightarrow \quad
a_n = \left(\frac{3+\sqrt{5}}{2\sqrt{5}}\right)\left(\frac{1+\sqrt{5}}{2}\right)^n + \left(\frac{\sqrt{5}-3}{2\sqrt{5}}\right)\left(\frac{1-\sqrt{5}}{2}\right)^n - 1
$$

> **Note:** the handwritten notes give $C_1 = \frac{1+\sqrt{5}}{2\sqrt{5}}$, $C_2 = \frac{\sqrt{5}-1}{2\sqrt{5}}$ — an arithmetic slip: those constants produce the sequence shifted by one index ($0, 0, 1, 2, 4, 7, \ldots$ instead of $0, 1, 2, 4, 7, 12, \ldots$). The constants above solve the displayed system and reproduce $a_0 = 0$, $a_1 = 1$. Equivalently, $a_n = F_{n+2} - 1$ in terms of the Fibonacci numbers.
