# Discrete Mathematics — Lecture 8: Recurrences

> Translated from the handwritten lecture notes `pdf/lect8rekurencjeEN-MatDysk.pdf`.
> The original pages are headed "Algorithms & Complexity — Lecture 5"; the notes are written mostly in English.

## Goal of the lecture

We shall: analyze how to transform a **recursive formula** (an *implicit* form) into an **explicit formulation**, so that we can assess

$$
a_n = O(f(n)) \stackrel{?}{=}
$$

## Example: the Tower of Hanoi problem

How to shift $n$ blocks (here $n = 3$) from the 1st peg to the 3rd peg, with the aid of peg 2, subject to the constraint: *"only a smaller block can be put on a larger one"*.

The starting position — the whole tower must travel from peg 1 to peg 3:

![[lec08en_p01_hanoi-setup.svg]]

$a_n$ — the **number of shifts**; $a_n = O(?)$.

The teacher's sequence of doodles showing the full solution for $n = 3$ (each frame shows the next move as a dashed arrow):

![[lec08en_p01_hanoi-seven-moves.svg]]

so $a_3 = 7$. But for arbitrary $n$: $a_n = \ ?$

### Finding the recursive formula

We try to find the recursive formula (for the specific algorithm). The four stages of moving $n$ disks:

1. all $n$ disks (the largest one plus a stack of $n-1$) sit on peg 1;
2. move the top $n - 1$ disks to peg 2 — **cost $T_{n-1}$**;
3. move the largest disk from peg 1 to peg 3 — **cost $1$**;
4. move the $n - 1$ disks from peg 2 onto peg 3 — **cost $T_{n-1}$**.

![[lec08en_p02_hanoi-recursion.svg]]

Total cost: $T_n = T_{n-1} + 1 + T_{n-1}$, i.e.

$$
\boxed{\ T_n = 2\,T_{n-1} + 1\ } \quad (\ast)
$$

Perhaps now, having got $(\ast)$ and a general theory, we can transform

$$
T_n = 2\,T_{n-1} + 1 \quad \Longrightarrow \quad T_n = \text{explicit formula},
$$

which enables us to assess $T_n = O(f(n))$.

## Difference equations — definitions

**Def.** The **degree (order) of a difference equation** is the maximal difference between the smallest and the biggest coefficient index in

$$
a_n = f(a_{n-1}, a_{n-2}, \ldots, a_{n-k},\, n). \quad (\ast) \ \text{(i)}
$$

Here it is $k$.

If

$$
a_n = \sum_{i=1}^{k} f_i(n)\, a_{n-i} + g(n) \quad (\blacktriangle) \ \text{(ii)}
$$

(i.e. a special case of $(\ast)$), then such a scheme $(\blacktriangle)$ is called a **linear difference scheme**, **inhomogeneous** (as $g(n) \neq 0$) of **order $k$**.

- (iii) If $g(n) = 0$, then it is **homogeneous**.
- (iv) If $f_i(n) = \alpha_i$ (constants), then

$$
a_n = \sum_{i=1}^{k} \alpha_i\, a_{n-i} + g(n)
$$

is called a **linear difference scheme of order $k$ with constant coefficients**, inhomogeneous (as $g(n) \neq 0$) — (IV) IH; if $g(n) = 0$, homogeneous — (V) H.

## Examples

1. $E_n = E_{n-1}\,\dfrac{n-1}{n} + 1$ — $k = 1$; linear, inhomogeneous, **non-constant** coefficients.
2. $a_{n+1} = \dfrac{1}{2}\left(a_n + \dfrac{4}{a_n}\right)$ — $k = 1$; **non-linear**.
3. $a_n = a_{n-1} \cdot a_{n-2}$ — $k = 2$; non-linear.
4. $a_n = 2a_{n-1} + 1$ — $k = 1$; linear, inhomogeneous, with constant coefficients.
5. $a_n = a_{n-10}$ — $k = 10$; linear, homogeneous, with constant coefficients: $\alpha_{10} = 1$, $\alpha_1 = \ldots = \alpha_9 = 0$.
6. $a_n = A \cdot T\!\left(\lfloor \tfrac{n}{2} \rfloor\right)$ — $k = \infty$, so this is an example of a scheme of **not fixed order** (it is a **Divide & Conquer** scheme); homogeneous.
7. $a_n = 1 + a_{\lfloor \frac{n}{2} \rfloor}$ — $k = \infty$; inhomogeneous.
8. $a_n = a_{n-1} + \underbrace{n^3 - 2n^2 + 3n + 2}_{g(n)}$ — $k = 1$; linear, inhomogeneous, constant coefficients.
9. $a_{n+2} = a_{n+1} + a_n$ — $k = 2$; linear, constant coefficients, homogeneous.

## Initial conditions and uniqueness

Usually ($k$ = order)

$$
a_n = f(a_{n-1}, a_{n-2}, \ldots, a_{n-k},\, n) \quad (\ast)
$$

is supplemented with $k$ **initial values**

$$
a_0, a_1, \ldots, a_{k-1} \quad (\ast\ast)
$$

in order to enforce **uniqueness** of $a_n$ satisfying $(\ast)$ and $(\ast\ast)$.

**Theorem.** Let $f : \mathbb{R}^k \times \mathbb{N} \to \mathbb{R}$ be a function. Then there exists **exactly one** sequence $\{a_n\}_{n=0}^{\infty}$ satisfying

$$
\begin{aligned}
\text{a)} \quad & a_n = f(a_{n-1}, a_{n-2}, \ldots, a_{n-k},\, n) \\
\text{b)} \quad & a_i = c_i \quad 0 \leq i \leq k-1
\end{aligned}
$$

**Proof** (by strong induction). $P(n)$: there exists exactly one $\{a_m\}_{m \geq 0}$ satisfying a) and b) up to index $n$.

**I)** $P(0), P(1), \ldots, P(k-1)$ are true, as only one sequence of numbers $a_0, a_1, \ldots, a_{k-1}$ satisfies the initial conditions.

**II)** Assume $P(j)$ holds $\forall j < n$. Is $P(n)$ true?

Based on:

$$
\begin{aligned}
P(n-1)\ \text{true} \quad & \text{as } n-1 < n \\
P(n-2)\ \text{true} \quad & \text{as } n-2 < n \\
& \ \ \vdots \\
P(n-k)\ \text{true} \quad & \text{as } n-k < n
\end{aligned}
$$

we have only one $a_{n-1}, a_{n-2}, \ldots, a_{n-k}$. Since

$$
a_n = f(a_{n-1}, a_{n-2}, \ldots, a_{n-k},\, n),
$$

$a_n$ is uniquely determined by $f$ and the unique $a_{n-1}, a_{n-2}, \ldots, a_{n-k}$ and $n$. This shows $P(n)$.

By the principle of mathematical induction, $\{a_n\}_{n \geq 0}$ is determined uniquely. $\square$

## The space of sequences

We develop now the general approach to solve:

$$
\boxed{\ a_n = \sum_{i=1}^{k} \alpha_i\, a_{n-i}\ }
$$

**Fact 1.**

$$
V_{\infty} = \left\{ \begin{pmatrix} a_0 \\ a_1 \\ \vdots \\ a_n \\ \vdots \end{pmatrix} : a_i \in \mathbb{R} \right\}
$$

— the space of vectors with $\infty$-many coefficients.

Assume $\{a_n\}_{n \geq 0} \in V_{\infty}$, $\{b_n\}_{n \geq 0} \in V_{\infty}$. Define:

- (i) $\{a_n\}_{n \geq 0} \oplus \{b_n\}_{n \geq 0} = \{c_n\}_{n \geq 0} \in V_{\infty}$, where $c_n = a_n + b_n$;
- (ii) $\lambda \odot \{a_n\}_{n \geq 0} = \{d_n\}_{n \geq 0} \in V_{\infty}$, where $d_n = \lambda \cdot a_n$.

It is easy to show that $(V_{\infty}, \oplus, \odot, \mathbb{R})$ defines a **vector space** over the field $K = \mathbb{R}$.

Take a subset $V_{\infty}^{k} \subseteq V_{\infty}$:

$$
V_{\infty}^{k} = \left\{ \{a_n\}_{n \geq 0} :\ a_n = \sum_{i=1}^{k} \alpha_i\, a_{n-i} \ (\ast) \right\}
$$

— the set of all sequences from $V_{\infty}$ which satisfy the recursive scheme $a_n = \sum_{i=1}^{k} \alpha_i a_{n-i}$.

### $V_{\infty}^{k}$ is a linear subspace of $V_{\infty}$

1. $\{a_n\}_{n \geq 0},\, \{b_n\}_{n \geq 0} \in V_{\infty}^{k} \ \Rightarrow\ \{a_n\}_{n \geq 0} \oplus \{b_n\}_{n \geq 0} \in V_{\infty}^{k}$
2. $\{a_n\}_{n \geq 0} \in V_{\infty}^{k} \ \Rightarrow\ \lambda \odot \{a_n\}_{n \geq 0} \in V_{\infty}^{k}$

**Proof of 1)** Since $\{a_n\}_{n \geq 0}$ and $\{b_n\}_{n \geq 0} \in V_{\infty}^{k}$:

$$
a_n = \sum_{i=1}^{k} \alpha_i\, a_{n-i} \quad \text{and} \quad b_n = \sum_{i=1}^{k} \alpha_i\, b_{n-i}
$$

$$
\underbrace{a_n + b_n}_{c_n} = \sum_{i=1}^{k} \alpha_i\, a_{n-i} + \sum_{i=1}^{k} \alpha_i\, b_{n-i}
= \sum_{i=1}^{k} \alpha_i \underbrace{(a_{n-i} + b_{n-i})}_{c_{n-i}}
$$

$\Downarrow$ $\{c_n\}_{n \geq 0}$ also satisfies $(\ast)$, so $\{c_n\}_{n \geq 0} = \{a_n\}_{n \geq 0} \oplus \{b_n\}_{n \geq 0} \in V_{\infty}^{k}$.

**Proof of 2)** $\lambda \odot \{a_n\}_{n \geq 0} = \{\lambda a_n\}_{n \geq 0}$:

$$
a_n = \sum_{i=1}^{k} \alpha_i\, a_{n-i} \quad \Big|\, \cdot \lambda
$$

$$
\underbrace{\lambda a_n}_{c_n} = \sum_{i=1}^{k} \alpha_i \underbrace{\lambda a_{n-i}}_{c_{n-i}}
$$

so $\{c_n\}_{n \geq 0} = \lambda \odot \{a_n\}_{n \geq 0}$ also satisfies $(\ast)$ $\Rightarrow$ $\{c_n\}_{n \geq 0} \in V_{\infty}^{k}$. $\square$

## The mapping $\varphi$ is a linear isomorphism

**Mapping:**

$$
\varphi : \begin{pmatrix} a_0 \\ a_1 \\ \vdots \\ a_{k-1} \end{pmatrix} \longmapsto \{a_n\}_{n \geq 0} \in V_{\infty}^{k}
$$

— the sequence is uniquely determined by the initial vector via the scheme:

$$
(\ast\ast) \quad
\begin{cases}
a_0,\ a_1,\ \ldots,\ a_{k-1} & \text{(the given initial values)} \\
a_k = \alpha_1 a_{k-1} + \alpha_2 a_{k-2} + \ldots + \alpha_k a_0 \\
a_{k+1} = \alpha_1 a_k + \alpha_2 a_{k-1} + \ldots + \alpha_k a_1 \\
\quad\vdots \\
a_n = \alpha_1 a_{n-1} + \alpha_2 a_{n-2} + \ldots + \alpha_k a_{n-k} \\
\quad\vdots
\end{cases}
$$

**$\varphi$ is linear.** Indeed, take two initial vectors and their images

$$
\varphi\!\begin{pmatrix} a_0 \\ \vdots \\ a_{k-1} \end{pmatrix} = \{a_n\}_{n \geq 0},
\qquad
\varphi\!\begin{pmatrix} b_0 \\ \vdots \\ b_{k-1} \end{pmatrix} = \{b_n\}_{n \geq 0},
$$

where $\{b_n\}$ is generated by the same scheme: $b_n = \alpha_1 b_{n-1} + \alpha_2 b_{n-2} + \ldots + \alpha_k b_{n-k}$. Then

$$
\varphi\!\begin{pmatrix} a_0 \\ \vdots \\ a_{k-1} \end{pmatrix} \oplus \varphi\!\begin{pmatrix} b_0 \\ \vdots \\ b_{k-1} \end{pmatrix} = \{c_n\}_{n \geq 0},
$$

which is the sequence

$$
(\ast\!\ast\!\ast) \quad
\begin{cases}
a_0 + b_0,\ a_1 + b_1,\ \ldots,\ a_{k-1} + b_{k-1} \\
a_k + b_k = \alpha_1 (a_{k-1} + b_{k-1}) + \ldots + \alpha_k (a_0 + b_0) \\
a_{k+1} + b_{k+1} = \alpha_1 (a_k + b_k) + \ldots + \alpha_k (a_1 + b_1) \\
\quad\vdots \\
a_n + b_n = \alpha_1 (a_{n-1} + b_{n-1}) + \ldots + \alpha_k (a_{n-k} + b_{n-k}) \\
\quad\vdots
\end{cases}
$$

But also

$$
\varphi\!\left(\begin{pmatrix} a_0 \\ \vdots \\ a_{k-1} \end{pmatrix} + \begin{pmatrix} b_0 \\ \vdots \\ b_{k-1} \end{pmatrix}\right)
= \varphi\!\begin{pmatrix} a_0 + b_0 \\ \vdots \\ a_{k-1} + b_{k-1} \end{pmatrix} = (\ast\!\ast\!\ast).
$$

So

$$
\varphi\!\begin{pmatrix} a_0 \\ \vdots \\ a_{k-1} \end{pmatrix} \oplus \varphi\!\begin{pmatrix} b_0 \\ \vdots \\ b_{k-1} \end{pmatrix}
= \varphi\!\left(\begin{pmatrix} a_0 \\ \vdots \\ a_{k-1} \end{pmatrix} + \begin{pmatrix} b_0 \\ \vdots \\ b_{k-1} \end{pmatrix}\right).
$$

Similarly it is easy to show that

$$
\lambda \odot \varphi\!\begin{pmatrix} a_0 \\ \vdots \\ a_{k-1} \end{pmatrix} = \varphi\!\left(\lambda \cdot \begin{pmatrix} a_0 \\ \vdots \\ a_{k-1} \end{pmatrix}\right).
$$

Thus $\varphi : V^k \longrightarrow V_{\infty}^{k}$ is **linear**, where the space

$$
V^k = \left\{ \begin{pmatrix} a_0 \\ \vdots \\ a_{k-1} \end{pmatrix} : a_i \in \mathbb{R} \right\}.
$$

- It is also **injective**, as $(a_0, \ldots, a_{k-1})$ uniquely determines $\{a_n\}_{n \geq 0} \in V_{\infty}^{k}$ (proved by induction).
- It is also **surjective**, as each sequence $\{a_n\}_{n \geq 0} \in V_{\infty}^{k}$ is an image of $(a_0, a_1, \ldots, a_{k-1})$.

Thus $\varphi$ is a **linear isomorphism**, hence

$$
\dim(V^k) = \dim(V_{\infty}^{k}).
$$

Since

$$
\begin{pmatrix} a_0 \\ a_1 \\ \vdots \\ a_{k-1} \end{pmatrix}
= a_0 \begin{pmatrix} 1 \\ 0 \\ \vdots \\ 0 \end{pmatrix}
+ a_1 \begin{pmatrix} 0 \\ 1 \\ \vdots \\ 0 \end{pmatrix}
+ \ldots
+ a_{k-1} \begin{pmatrix} 0 \\ 0 \\ \vdots \\ 1 \end{pmatrix}
\quad \longleftarrow \text{canonical basis in } V^k,
$$

we get

$$
\boxed{\ \dim(V_{\infty}^{k}) = k\ }
$$

Thus any sequence satisfying $a_n = \sum_{i=1}^{k} \alpha_i a_{n-i}$ is a **linear combination** (so it depends on $k$ parameters) of the basis sequences $\{e_n^i\}_{n \geq 0}$, $1 \leq i \leq k$, each satisfying $a_n = \sum_{i=1}^{k} \alpha_i a_{n-i}$; together with $a_0, \ldots, a_{k-1}$ the sequence is pinned down uniquely.

We shall look at that for different $k$.

## Case I: $k = 2$ (or $k = 1$)

$$
\boxed{\ a_n = a\, a_{n-1} + b\, a_{n-2}\ } \quad (\star)
$$

with $a_0, a_1$ given — linear homogeneous of order 1 or 2.

### (i) Take $b = 0$ and $a \neq 0$ $\Rightarrow$ $k = 1$

$$
a_n = a\, a_{n-1}, \qquad \dim(V_{\infty}^{1}) = 1
$$

$$
\frac{a_n}{a_{n-1}} = a \quad \text{so a geometric sequence:} \quad a_n = A \cdot a^n
$$

($a^n$ is the basis — 1 sequence). If $a_0$ is given $\Rightarrow a_0 = a^0 \cdot A \Rightarrow A = a_0$, so the unique sequence is

$$
\boxed{\ a_n = a_0\, a^n\ }
$$

### (ii) Assume now $a = 0$ and $b \neq 0$

$$
a_n = b\, a_{n-2}, \qquad \frac{a_n}{a_{n-2}} = b
$$

If $a_0$, $a_1$ are given:

$$
a_{2k} = b^k a_0, \qquad a_{2k+1} = b^k a_1
$$

**Example:**

(i) $S_n = 3 S_{n-1}$ (so $a = 3$), $S_0 = 5$:

$$
\boxed{\ S_n = 5 \cdot 3^n\ }
$$

(ii) $S_n = 3 S_{n-2}$, $S_0 = 5$, $S_1 = 2$:

$$
S_{2n} = 5 \cdot 3^n, \qquad S_{2n+1} = 2 \cdot 3^n
$$

### (iii) Case $k = 2$: $a \neq 0 \wedge b \neq 0$

We search for a solution of the form:

$$
\boxed{\ S_n = C \cdot r^n\ }
$$

Plugging into $(\star)$:

$$
C r^n = a\, C r^{n-1} + b\, C r^{n-2} \qquad \Big|\, : C r^{n-2}
$$

$$
\boxed{\ r^2 - a r - b = 0\ } \quad (\blacktriangle)
$$

— the **characteristic equation** for

$$
a_n = a\, a_{n-1} + b\, a_{n-2}. \quad (\blacktriangle\blacktriangle)
$$

### a) Assume $\Delta > 0$

$\exists\, r_1 \neq r_2$ (real numbers) solving $(\blacktriangle)$

$$
\Longrightarrow \quad
\{x_n\}_{n \geq 0} = \{\tilde{C}_1\, r_1^n\}_{n \geq 0}, \qquad
\{y_n\}_{n \geq 0} = \{\tilde{C}_2\, r_2^n\}_{n \geq 0}
$$

with $\tilde{C}_1$, $\tilde{C}_2$ arbitrary. These two sequences satisfy:

$$
S_n = a\, S_{n-1} + b\, S_{n-2}
$$

It is easy to check that they are **linearly independent** (and as $\dim(V_{\infty}^{k}) = 2$) they form the **basis** for $V_{\infty}^{k}$, and thus any solution to $(\blacktriangle\blacktriangle)$ is a linear combination of these basis vectors:

$$
\Downarrow
$$

$$
S_n = \tilde{C}_1\, r_1^n + \tilde{C}_2\, r_2^n
$$

($S_n$ depends on 2 parameters). If we add $S_0$ and $S_1$, we have a unique solution $S_n$ — we eliminate $\tilde{C}_1$ and $\tilde{C}_2$ by solving

$$
\begin{cases}
S_0 = \tilde{C}_1\, r_1^0 + \tilde{C}_2\, r_2^0 \\
S_1 = \tilde{C}_1\, r_1^1 + \tilde{C}_2\, r_2^1
\end{cases}
\qquad \det \begin{pmatrix} 1 & 1 \\ r_1 & r_2 \end{pmatrix} \neq 0,
$$

as we know $S_0$, $S_1$.

**Linear independence.** To see that $\{r_1^n\}_{n \geq 0}$ and $\{r_2^n\}_{n \geq 0}$ are linearly independent, take a linear combination:

$$
\alpha \odot \{r_1^n\}_{n \geq 0} \oplus \beta \odot \{r_2^n\}_{n \geq 0} = \{0\}_{n \geq 0} \quad (0, 0, \ldots, 0, \ldots)
$$

$$
\Downarrow
$$

$$
\alpha\, r_1^n + \beta\, r_2^n = 0 \qquad \forall n \in \mathbb{N}
$$

Take e.g. $n = 0$ and $n = 1$:

$$
\begin{cases}
\alpha + \beta = 0 \\
r_1 \alpha + r_2 \beta = 0
\end{cases}
\iff
\begin{bmatrix} 1 & 1 \\ r_1 & r_2 \end{bmatrix}
\begin{bmatrix} \alpha \\ \beta \end{bmatrix}
= \begin{bmatrix} 0 \\ 0 \end{bmatrix}
$$

$$
\det \begin{bmatrix} 1 & 1 \\ r_1 & r_2 \end{bmatrix} = r_2 - r_1 \neq 0 \quad \text{as } r_2 \neq r_1 \text{ (as } \Delta > 0\text{)}
$$

So we have one solution to this system: $\alpha = \beta = 0$ $\Rightarrow$ $\{r_1^n\}_{n \geq 0}$ and $\{r_2^n\}_{n \geq 0}$ are **linearly independent**. $\square$

### b) Assume now $\Delta = 0$

$$
\Rightarrow \quad r_1 = r_2 = r \quad \text{(the multiplicity of } r_1 \text{ is two)}
$$

This time $\{r_1^n\}_{n \geq 0}$ and $\{r_1^n\}_{n \geq 0}$ are **linearly dependent**! So we need to find one more solution to $(\blacktriangle\blacktriangle)$ which is not linearly dependent on $r^n$:

$$
\{x_n\}_{n \geq 0} = \{C_1\, r^n\}_{n \geq 0}
$$

Take now

$$
\{y_n\}_{n \geq 0} = \{C_2\, n\, r^n\}_{n \geq 0}
$$

We check:

- **a)** $\{y_n\}_{n \geq 0}$ satisfies $a_n = a\, a_{n-1} + b\, a_{n-2}$;
- **b)** $\{y_n\}_{n \geq 0}$ and $\{x_n\}_{n \geq 0}$ are linearly independent.

$$
\Downarrow
$$

Any solution to $a_n = a\, a_{n-1} + b\, a_{n-2}$ (as $\dim(V_{\infty}^{k}) = 2$) is a linear combination:

$$
S_n = \tilde{C}_1\, r^n + \tilde{C}_2\, n\, r^n
$$

(depends on 2 parameters). If we add $S_0$ and $S_1$, we can eliminate $\tilde{C}_1$ and $\tilde{C}_2$ as follows:

$$
\begin{cases}
S_0 = \tilde{C}_1\, r^0 + \tilde{C}_2 \cdot 0 \cdot r^0 \\
S_1 = \tilde{C}_1\, r^1 + \tilde{C}_2 \cdot 1 \cdot r^1
\end{cases}
\qquad
\det \begin{pmatrix} 1 & 0 \\ r & r \end{pmatrix} = r \neq 0 \ \text{ for } r \neq 0.
$$

> If $r = 0$, then $r^2 - ar - b = 0$ would give $-b = 0$, but we assumed $b \neq 0$ — so indeed $r \neq 0$.

So we compute $\tilde{C}_1$ and $\tilde{C}_2$ from the system and we get a **unique** solution to

$$
\begin{cases}
a_n = a\, a_{n-1} + b\, a_{n-2} \\
a_0,\ a_1 \ \text{given.}
\end{cases}
$$

**We show a):** $n r^n$ satisfies $a_n = a\, a_{n-1} + b\, a_{n-2}$ $(\star)$, provided $r^n$ satisfies $(\star)$.

Note (the characteristic equation): if $\Delta = 0$, then

$$
r^2 - ar - b = 0 \quad (\ast) \qquad \iff \qquad (r - r_1)^2 = 0 \quad (\ast\ast)
$$

where $r_1$ is the (double) root. Expanding:

$$
r^2 - ar - b = 0 \iff r^2 - 2 r r_1 + r_1^2 = 0
\quad \Rightarrow \quad \boxed{a = 2 r_1} \quad \boxed{b = -r_1^2}
$$

Now we put $n\, r_1^n$ into $a_n = a\, a_{n-1} + b\, a_{n-2}$:

$$
\begin{aligned}
n\, r_1^n &= a (n-1)\, r_1^{n-1} + b (n-2)\, r_1^{n-2} \\
n\, r_1^n &= 2 r_1 (n-1)\, r_1^{n-1} + (-r_1^2)(n-2)\, r_1^{n-2} \\
n\, r_1^n &= 2 (n-1)\, r_1^{n} - (n-2)\, r_1^{n} \qquad \Big|\, : r_1^n \\
n &= 2n - 2 - n + 2 \\
n &= n \\
0 &= 0 \qquad \text{so a) holds. } \underline{\underline{\text{OK}}}
\end{aligned}
$$

**b)** $\{r_1^n\}_{n \geq 0}$ and $\{n\, r_1^n\}_{n \geq 0}$ are linearly independent:

$$
\alpha \odot \{r_1^n\}_{n \geq 0} \oplus \beta \odot \{n\, r_1^n\}_{n \geq 0} = \{0\}_{n \geq 0}
$$

$$
\Downarrow \ \forall n \geq 0 \qquad \alpha\, r_1^n + \beta\, n\, r_1^n = 0
$$

Take $n = 0$ $\Rightarrow$ $\alpha = 0$; then (e.g. $n = 1$) $\Rightarrow$ $\beta = 0$. So $\{r_1^n\}_{n \geq 0}$ and $\{n\, r_1^n\}_{n \geq 0}$ are **linearly independent**. $\square$

### c) Assume $\Delta < 0$

This is the same as $\Delta > 0$, since we get $r_1 \neq r_2 \in \mathbb{C}$ (complex roots).

## Worked examples

### a) Distinct roots

$$
\boxed{\ a_n = a_{n-1} + 2 a_{n-2}\ } \quad (\ast)
$$

Try $a_n = C r^n$:

$$
C r^n = C r^{n-1} + 2 C r^{n-2} \quad \Rightarrow \quad r^2 = r + 2 \quad \Rightarrow \quad \boxed{r^2 - r - 2 = 0}
$$

— the characteristic equation. $\Delta > 0$ and $r_1 = 2 \neq r_2 = -1$.

So we have 2 linearly independent solutions $2^n$ and $(-1)^n$, so

$$
c_n = C_1 (-1)^n + C_2\, 2^n
$$

— an arbitrary solution to $(\ast)$. If we add $\boxed{a_0 = a_1 = 3}$, we can eliminate the parameters $C_1$ and $C_2$:

$$
\begin{cases}
3 = a_0 = C_1 (-1)^0 + C_2\, 2^0 \\
3 = a_1 = C_1 (-1)^1 + C_2\, 2^1
\end{cases}
\quad \Updownarrow \quad
\begin{cases}
3 = C_1 + C_2 \\
3 = -C_1 + 2 C_2
\end{cases}
\quad \Rightarrow \quad
\begin{aligned}
C_1 &= 1 \\
C_2 &= 2
\end{aligned}
$$

$$
\Rightarrow \quad \boxed{\ a_n = 2^{n+1} + (-1)^n\ } \quad \Rightarrow \quad \boxed{\ a_n = \Theta(2^n)\ }
$$

### b) A double root

$$
\boxed{\ S_n = 6 S_{n-1} - 9 S_{n-2}\ } \quad (\ast\ast)
$$

Characteristic polynomial:

$$
r^2 - 6r + 9 = 0 \quad \Rightarrow \quad \Delta = 0 \quad \Rightarrow \quad r = 3
$$

> The handwritten note has a sign slip here ($r^2 - 6r - 9 = 0$); for $S_n = 6S_{n-1} - 9S_{n-2}$ the characteristic equation is $r^2 - 6r + 9 = 0$, consistent with $\Delta = 0$ and $r = 3$ as used in the rest of the computation.

$$
a_n = 3^n \quad \text{and} \quad b_n = n\, 3^n
$$

— the basis vectors. An arbitrary solution to $(\ast\ast)$ is of the form:

$$
S_n = C_1\, 3^n + C_2\, n\, 3^n
$$

If the initial conditions are known, $\boxed{S_0 = 1 \ \text{and}\ S_1 = -3}$, then we can eliminate $C_1$ and $C_2$:

$$
\begin{cases}
1 = S_0 = C_1\, 3^0 + C_2 \cdot 0 \cdot 3^0 \\
-3 = S_1 = C_1\, 3^1 + C_2 \cdot 1 \cdot 3^1
\end{cases}
\quad \Rightarrow \quad C_1 = 1 \ \text{ and } \ C_2 = -2,
$$

and thus

$$
\boxed{\ S_n = 3^n - 2 \cdot n \cdot 3^n\ } \quad \Rightarrow \quad \boxed{\ S_n = \Theta(n\, 3^n)\ }
$$

— the unique sequence satisfying $(\ast\ast)$ and $S_0 = 1$, $S_1 = -3$.

## Higher orders and a final remark

The cases for $k > 2$ can be analyzed similarly, but the **multiplicities of roots** of the characteristic polynomial decide about the basis for $V_{\infty}^{k}$.

**Remark:** once you find $a_n$ in the explicit form (whatever method you use), **check** whether your solution satisfies the recursive scheme and the initial conditions, to eliminate the arithmetic errors which potentially may happen in your calculations!
