# Discrete Mathematics — Exercises 3: Solutions

> Translated from the handwritten lecture notes `pdf/Ćwicz3-RozwiązaniaMatDysk.pdf`.

## Problem 1 — image, preimage, domain and range

$$
X = \{1, \bullet, \heartsuit, \pi, *\}, \qquad Y = \{8, 9, 10, 11, 12, \, !\, \}
$$

$$
f(1) = 10, \quad f(\bullet) = 9, \quad f(\heartsuit) = \, !\,, \quad f(\pi) = 10, \quad f(*) = 8
$$

$$
A = \{\bullet, \heartsuit, \pi\}, \qquad B = \{10, \, !\, \}
$$

- $f(A) = \{9, \, !\,, 10\}$ — the **image** of the set $A$ under $f$;
- $f^{-1}(\{10, \, !\, \}) = \{1, \pi, \heartsuit\}$ — the **preimage** of the set $B$ under $f^{-1}$.

$$
\mathrm{Dom}(f) = X
$$

— the **domain** of $f$ (because $f$ is defined on the whole of $X$).

$$
\mathrm{Rg}(f) = Y \setminus \{11, 12\} = \{8, 9, 10, \, !\, \}
$$

— since neither $11$ nor $12$ is a value of $f$ at any argument.

$f$ is **not injective**, because for the pair $1 \neq \pi$ we have

$$
f(1) = f(\pi) = 10.
$$

$f$ is **not a surjection** of $X$ onto $Y$, because $\mathrm{Rg}(f) \neq Y$. In fact there is **no** surjection from $X$ onto $Y$ at all, because

$$
\overline{\overline{X}} = 5 \quad \text{while} \quad \overline{\overline{Y}} = 6.
$$

And so, since there is no surjection, there is **no bijection** from $X$ to $Y$ either.

## Problem 2 — the functions $f_1(x) = \frac{1}{x-1}$ and $f_2(x) = \sqrt[4]{x}$

$$
f_1(x) = \frac{1}{x - 1}, \qquad f_2(x) = \sqrt[4]{x}
$$

To draw $f_1$ we take the graph of $\frac{1}{x}$ and **shift it to the right by 1**. The teacher's first rough sketches — the dashed graph of $\frac{1}{x}$, the graph of $f_2$ through $(1,1)$, and the shifted hyperbola through $(2,1)$ and $(0,-1)$:

![[sol03_p01_graph-doodles.svg]]

The clean graphs: $f_1$ with the dashed vertical asymptote $x = 1$, and $f_2$:

![[sol03_p02_f1-f2-graphs.svg]]

### Domain and range

$$
\begin{aligned}
\mathrm{Dom}(f_1) &= \mathbb{R} \setminus \{1\} & \mathrm{Dom}(f_2) &= [0, +\infty) \\
\mathrm{Rg}(f_1) &= \mathbb{R} \setminus \{0\} & \mathrm{Rg}(f_2) &= [0, +\infty) = \mathbb{R}_+ \cup \{0\}
\end{aligned}
$$

### Injectivity

$f_1$ **is injective** (one can see that no line parallel to the axis $OX$ cuts the graph of $f_1$ in 2 places — this is the *geometric* argument). A similar argument works for $f_2$.

But *analytically* as well:

**(•)** For $f_1$:

$$
\begin{aligned}
\frac{1}{x-1} &= y \qquad (x \neq 1) \\
1 &= y \cdot (x - 1) \qquad (y \neq 0), \ \text{so} \\
\frac{1}{y} &= x - 1 \\
\frac{1}{y} + 1 &= x
\end{aligned}
$$

— that is, to one $y$ there corresponds one $x$.

**(••)** For $f_2$:

$$
y = \sqrt[4]{x}, \qquad y^4 = x
$$

— to one $y$ there corresponds $1$ value of $x$.

**Or:** does the implication

$$
x_1 \neq x_2 \ \stackrel{?}{\Rightarrow} \ f_i(x_1) \neq f_i(x_2), \qquad i = 1, 2
$$

hold, i.e.

$$
\frac{1}{x_1 - 1} \stackrel{?}{\neq} \frac{1}{x_2 - 1}, \qquad \sqrt[4]{x_1} \stackrel{?}{\neq} \sqrt[4]{x_2} \ ?
$$

**Yes**, because:

$$
x_1 \neq x_2 \ \Rightarrow \ x_1 - 1 \neq x_2 - 1 \ \Rightarrow \ \frac{1}{x_1 - 1} \neq \frac{1}{x_2 - 1} \ \iff \ f_1(x_1) \neq f_1(x_2)
$$

and for $x_1 \neq x_2$ indeed

$$
\sqrt[4]{x_1} \neq \sqrt[4]{x_2}, \quad \text{i.e.} \quad f_2(x_1) \neq f_2(x_2).
$$

So the $f_i$ $(i = 1, 2)$ are **injective**.

### Inverse functions

Note that from the arguments (•) and (••) we also get the formulas for the **inverse functions**:

$$
f_1^{-1} : \mathbb{R} \setminus \{0\} \to \mathbb{R} \setminus \{1\}, \qquad f_1^{-1}(y) = \frac{1}{y} + 1
$$

$$
f_2^{-1} : \mathbb{R}_+ \cup \{0\} \to \mathbb{R}_+ \cup \{0\}, \qquad f_2^{-1}(y) = y^4
$$

### Surjectivity

$f_1$ is **not** onto from $\mathbb{R} \setminus \{1\}$ onto $\mathbb{R}$, because

$$
\mathrm{Rg}(f_1) = \mathbb{R} \setminus \{0\} \neq \mathbb{R}.
$$

$f_2$ is **not** onto from $\mathbb{R}_+ \cup \{0\}$ onto $\mathbb{R}$, because

$$
\mathrm{Rg}(f_2) = \mathbb{R}_+ \cup \{0\} \neq \mathbb{R}
$$

(but it **is** onto $\mathbb{R}_+ \cup \{0\}$).

### Images and preimages of intervals

$$
f_2\big(\underbrace{[0, 2]}_{A_1}\big) = \big[0, \sqrt[4]{2}\,\big] \qquad \text{— see Fig. 1.}
$$

The segment $A_1$ on the $x$-axis and its image on the $y$-axis:

![[sol03_p04_fig1-image-f2.svg|340]]

$$
f_1\big(\underbrace{(1, +\infty)}_{A_2}\big) = (0, +\infty) \qquad \text{— see Fig. 2.}
$$

The ray $A_2$ (open at $1$) and its image $(0,+\infty)$ on the $y$-axis:

![[sol03_p04_fig2-image-f1.svg|360]]

$$
f_2^{-1}\big(\underbrace{[0, 16]}_{B_1}\big) = \big[0, 16^4\big]
$$

$$
f_1^{-1}\big(\underbrace{(0, 1)}_{B_2}\big) = (2, +\infty) \qquad \text{— see Fig. 3.}
$$

The segment $B_2$ on the $y$-axis and its preimage on the $x$-axis:

![[sol03_p04_fig3-preimage-f1.svg|400]]

## Problem 3 — counting

### a) Three tickets for eight seats

$$
\overline{X} = \{b_1, b_2, b_3\}, \qquad \overline{Y} = \{m_1, m_2, m_3, m_4, m_5, m_6, m_7, m_8\}
$$

($b_i$ — the tickets, $m_j$ — the seats.)

The question is whether an allocation of tickets to seats is a function from $\overline{Y}$ to $\overline{X}$ (option **a**) or from $\overline{X}$ to $\overline{Y}$ (option **b**). The teacher's sketch of an example allocation — ticket $b_2$ on seat $m_1$, ticket $b_1$ on seat $m_6$, ticket $b_3$ on seat $m_4$:

![[sol03_p05_ticket-seats.svg|380]]

**a)** e.g. $f(m_1) = b_2$, but $f(m_2) = \,?$ — because there is **no ticket here**. So an allocation is *not* a function from seats to tickets.

So it must be option **b)**:

$$
g(b_1) = m_6, \qquad g(b_2) = m_1, \qquad g(b_3) = m_4
$$

Hence $g : \overline{X} \to \overline{Y}$, and it must be an **injection** (every allocation of the tickets is a different injective function).

The number of allocations is the number of injective functions from $\overline{X}$ to $\overline{Y}$, i.e. **variations without repetition**:

$$
\binom{8}{3} \cdot 3! = \frac{8! \, 3!}{3! \, 5!} = \frac{8 \cdot 7 \cdot 6 \cdot 5!}{5!} = 8 \cdot 7 \cdot 6.
$$

### b) Digits on eight cells

The digits $\{4, 1, 2, 5\}$ are to be written into eight cells $c_1, c_2, \ldots, c_8$:

![[sol03_p06_eight-cells.svg|420]]

$$
\overline{X} = \{c_1, c_2, \ldots, c_8\}, \qquad \overline{Y} = \{4, 1, 2, 5\}
$$

Every allocation, e.g.

![[sol03_p06_example-allocation.svg|420]]

is a function $\overline{X} \to \overline{Y}$:

$$
\begin{aligned}
&f(c_1) = 4, \quad f(c_2) = 1, \quad f(c_3) = 2, \quad f(c_4) = 5, \\
&f(c_5) = 4, \quad f(c_6) = 4, \quad f(c_7) = 5, \quad f(c_8) = 2.
\end{aligned}
$$

Conversely it is **not** a function $g : \overline{Y} \to \overline{X}$, because

$$
g(4) = \,? \qquad c_1 \ \text{or} \ c_5 \ \text{or} \ c_6\,?
$$

The digits $\{4, 1, 2, 5\}$ **may repeat** — so the functions may be non-injective. Hence we count **all** functions from $\overline{X}$ to $\overline{Y}$ (the so-called **variations with repetitions**):

$$
|\overline{Y}|^{|\overline{X}|} = 4^8.
$$

> The manuscript writes $8^4$ here, but with $|\overline{X}| = 8$ cells and $|\overline{Y}| = 4$ digits the formula $|\overline{Y}|^{|\overline{X}|}$ gives $4^8$.

### c)

$$
20!
$$

— the number of **permutations** (of 20 objects).

### d) Balls in a row

$3$ yellow, $4$ white and $3$ black balls. An example arrangement:

![[sol03_p07_balls-row.svg|460]]

$$
\frac{10!}{3! \, 4! \, 3!}
$$

— **permutations with repetitions**.

### e) A delegation of women and men

$4$ women and $6$ men $= 10$ people.

**(i)** $\dbinom{4}{2}$ — in this many ways one can choose a $2$-person group of women from the group of $4$ women.

**(ii)** $\dbinom{6}{3}$ — in this many ways one can choose a $3$-person group of men from the group of $6$ men.

For each choice from (i) we may pick any of the $\binom{6}{3}$ possibilities from (ii). So by the **product rule** the number of choices is:

$$
\binom{4}{2} \cdot \binom{6}{3} = \frac{4!}{2! \, 2!} \cdot \frac{6!}{3! \, 3!}.
$$
