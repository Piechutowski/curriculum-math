# Discrete Mathematics — Lecture 2: Functions

> Translated from the handwritten lecture notes `pdf/lect2-funkcje-MatDysk.pdf`.

## Functions — basic notions

**Function** $f : A \to B$ (mapping the set $A$ into $B$) — a transformation which assigns to every element $a \in A$, **uniquely**, a certain element $b = f(a) \in B$.

- $A$ — the **domain** of the function: $\operatorname{Dom}(f)$ — the set of those elements on which $f$ is defined.

The set of all values of $f$ is a subset of $B$ — we call

$$
B = \operatorname{Im}(f)
$$

— the **codomain** of the function $f$.

> **Note:** strictly speaking $\operatorname{Im}(f)$ (the image — the set of all values of $f$) may be a *proper* subset of the codomain $B$, as example b) below itself shows.

**Example:**

**a)** $f : \mathbb{R} \to \mathbb{R}$

$$
f(x) =
\begin{cases}
x; & x \geq 0 \\
-x; & x < 0
\end{cases}
$$

$$
f(x) = |x| \quad \text{(another notation)}
$$

$$
\operatorname{Dom}(f) = \mathbb{R}, \qquad \operatorname{Im}(f) = \mathbb{R}_+ \cup \{0\}
$$

**b)** A mapping between two finite sets $X$ and $Y$, drawn with arrows:

![[lec02_p02_function-mapping.svg]]

$$
f(x_1) = y_1, \quad f(x_2) = y_4, \quad f(x_3) = y_5, \quad f(x_4) = y_2 \qquad \text{— a function}
$$

$$
\begin{aligned}
\operatorname{Dom}(f) &= \{x_1, x_2, x_3, x_4\} \\
\operatorname{Im}(f) &= \{y_1, y_2, y_4, y_5\} \subseteq Y = \{y_1, y_2, y_3, y_4, y_5\}
\end{aligned}
$$

**c)** Here $x_1$ is sent to two different points:

![[lec02_p02_not-a-function.svg]]

$f$ is **not a function**, because $f(x_1)$ does not take a unique value!!!

**d)** $f : \mathbb{N} \to \mathbb{R}$ — a **number sequence**:

$$
f(n) = a_n
$$

(to a natural number a real number is assigned).

$$
f_1(n) = \frac{1}{n}, \qquad f_2(n) = \frac{1}{n^2}, \qquad f_3(n) = n^3
$$

Representation of these sequences — as points on a number line and as discrete graphs:

![[lec02_p03_sequence-graphs.svg]]

The graph of a sequence is **not a line but a set of points**, because the domain is discrete. $\square$

## Graph of a function

**Graph of a function** $f : A \to B$ — the subset $\operatorname{Graph}(f) \subset A \times B$ such that

$$
(a, b) \in \operatorname{Graph}(f) \ \text{ if } \ \boxed{b = f(a)}
$$

**Example:**

**a)** $f(x) = |x|$ — the absolute value:

![[lec02_p03_abs-graph.svg|200]]

**b)** the signum function:

$$
f(x) = \operatorname{sgn}(x) =
\begin{cases}
1 & x > 0 \\
0 & x = 0 \\
-1 & x < 0
\end{cases}
$$

![[lec02_p03_sgn-graph.svg|220]]

**c)** $f(n) = n^2$ — a discrete set of points:

![[lec02_p03_squares-points.svg|200]]

## Injections and surjections

A function $f : A \to B$ is **one-to-one** (an **injection**):

$$
\forall\, x_1 \neq x_2 \ \Rightarrow \ f(x_1) \neq f(x_2).
$$

Or: if $f(x_1) = f(x_2)$ then $x_1 = x_2$.

We say that $f : A \to S \subseteq B$ is a **mapping onto** the subset $S$ if

$$
\forall\, y \in S \ \ \exists\, x \in A \quad f(x) = y
$$

— a **surjection**

("for every" $\forall$, "there exists" $\exists$).

**a)** "onto" $f : X \to Y$, but not an injection:

![[lec02_p04_onto-not-injective.svg|300]]

**b)** an injection $f : X \to Y$, but not a surjection; but if $Y' = \{1, 2, 3\}$, then $f : X \to Y'$ is also a surjection:

![[lec02_p04_injective-not-onto.svg|300]]

**Example:** $f : \mathbb{R} \to \mathbb{R}$

**a)** $f_1(x) = x^2$ — **not** an injection, since $f(x) = f(-x)$! A horizontal line $y \equiv c$ hits the parabola twice:

![[lec02_p05_parabola-line.svg|320]]

**b)** $f_2(x) = x^3$ — an injection (also a surjection), so a **bijection**:

![[lec02_p05_cubic-line.svg|300]]

**c)** $f(n) = 2n - 1$. Is it an injection?

$$
n_1 \neq n_2 \ \stackrel{?}{\Rightarrow} \ f(n_1) \neq f(n_2)
$$

$$
2n_1 \neq 2n_2, \qquad 2n_1 - 1 \neq 2n_2 - 1 \qquad \text{yes} \ \square
$$

## Composition of functions

The scheme of composing $f : A \to B$ with $g : B \to C$:

![[lec02_p06_composition.svg]]

**Composition of functions** $h = g \circ f$:

$$
h : A \to C
$$

$$
\boxed{\, h(x) = g\big(f(x)\big) \,}
$$

$$
\operatorname{Im}(f) \subseteq \operatorname{Dom}(g) \quad \text{— important!}
$$

**Example:**

**a)** $h(x) = (x^3 + 2x)^7$ is the composition of

$$
h_1(x) = x^3 + 2x, \qquad h_2(x) = x^7
$$

$$
h(x) = h_2\big(h_1(x)\big) = h_2(x^3 + 2x) = (x^3 + 2x)^7.
$$

**b)**

$$
f(x) = x^4, \qquad g(y) = \sqrt{y^2 + 1}, \qquad h(z) = z^2 + 72
$$

$$
\big(h \circ (g \circ f)\big)(x) = h\big((g \circ f)(x)\big) = h\big(g(f(x))\big)
$$

$$
\begin{aligned}
h\big(g(f(x))\big) = h\big(g(x^4)\big) &= h\left(\sqrt{x^8 + 1}\right) \\
&= \left(\sqrt{x^8 + 1}\right)^2 + 72 \\
&= \underline{x^8 + 73} \qquad \square
\end{aligned}
$$

The property of breaking composite functions into simple functions is useful in many applications (e.g. computing the derivative of a function).

**Associativity of function composition:**

$$
\boxed{\, h \circ (g \circ f) = (h \circ g) \circ f \,}
$$

Usually there is **no commutativity** — not only

$$
\boxed{\, f \circ g \neq g \circ f \,} \ !
$$

but also $f \circ g$ may exist while $g \circ f$ may not exist!

**Example:**

$$
f_1(x) = \sqrt{x}, \qquad f_2(x) = -x^4 - 1
$$

$$
(f_2 \circ f_1)(x) = f_2\big(f_1(x)\big) = f_2(\sqrt{x}) = -x^2 - 1
$$

$$
(f_1 \circ f_2)(x) = f_1\big(f_2(x)\big) = f_1(\underbrace{-x^4 - 1}_{\text{a negative number!}})
$$

$\sqrt{-x^4 - 1}$ **does not exist** — the image of the function $f_2$ is not a subset of the domain of the function $f_1$! $\square$

**Example** (restrictions on the domain), with sketches of the graphs:

- **a)** $f(x) = \dfrac{1}{x}$, $\quad x \neq 0$
- **b)** $f(x) = \sqrt{x}$, $\quad x \geq 0$
- **c)** $f(x) = \ln x$, $\quad x > 0$
- **d)** $f(x) = \tan(x)$, $\quad x \neq \frac{\pi}{2} + k\pi$ (equivalently $x \neq -\frac{\pi}{2} + k\pi$), $k \in \mathbb{Z}$

![[lec02_p08_domain-restrictions.svg]]

## Inverse function

The **inverse function** of $f : A \to B$ is the function $f^{-1} : B \to A$ satisfying:

$$
\boxed{\
\begin{aligned}
f \circ f^{-1} &= \operatorname{id}_B \\
f^{-1} \circ f &= \operatorname{id}_A
\end{aligned}
\ }
$$

$$
\begin{aligned}
\operatorname{id}_A(x) = x \qquad \forall\, x \in A \\
\operatorname{id}_B(x) = x \qquad \forall\, x \in B
\end{aligned}
$$

The teacher's scheme — $f$ and $f^{-1}$ acting between $A$ and $B$ in both directions:

![[lec02_p09_inverse-diagram.svg|300]]

**Not all functions are invertible!**

$$
f(x) = x^2, \qquad f : \mathbb{R} \to \mathbb{R}
$$

$f^{-1}$ does not exist:

$$
f(1) = 1, \quad f(-1) = 1 \quad \text{(not an injection)}, \qquad f^{-1}(1) \to
\begin{cases}
-1 \\
1
\end{cases}?
$$

But for

$$
f : \mathbb{R}_+ \cup \{0\} \to \mathbb{R}, \qquad f(x) = x^2 \quad \text{then} \quad f^{-1}(x) = \sqrt{x}
$$

("injection"). Sketches — the full parabola (not an injection), the restricted half-parabola (an injection), and $\sqrt{x}$:

![[lec02_p09_invertibility-graphs.svg]]

**Example:** $f(x) = 2x + 1$; what is $f^{-1}$?

$$
2x + 1 = y \quad \Rightarrow \quad x = \frac{y - 1}{2} \qquad \Rightarrow \qquad f^{-1}(x) = \frac{x - 1}{2}
$$

Let us check:

$$
(f \circ f^{-1})(x) = f\left(\frac{x-1}{2}\right) = 2 \cdot \frac{x-1}{2} + 1 = \underline{\underline{x}} = \operatorname{id}(x)
$$

$$
(f^{-1} \circ f)(x) = f^{-1}(2x + 1) = \frac{2x + 1 - 1}{2} = x \qquad \square
$$

**Theorem:** A function $f : A \to B$ is invertible if and only if it is one-to-one and maps $A$ **onto** $B$.

## Important functions

**Example:**

**(i)** The function

$$
\chi_A(x) =
\begin{cases}
1 & x \in A \\
0 & x \notin A
\end{cases}
$$

is called the **characteristic function** of the set $A$:

![[lec02_p10_characteristic-function.svg|300]]

**(ii)** The **floor function**:

$$
\lfloor x \rfloor =
\begin{cases}
x\,; & x \ \text{an integer} \\[4pt]
\text{the nearest integer} < x\,; & x \ \text{not an integer}
\end{cases}
$$

$$
\lfloor 5 \rfloor = 5, \qquad \lfloor -5 \rfloor = -5, \qquad \lfloor 5.1 \rfloor = 5, \qquad \lfloor -5.3 \rfloor = -6
$$

Its staircase graph, between the dotted lines $y = x$ and $y = x - 1$:

![[lec02_p11_floor-graph.svg]]

$$
\boxed{\, x - 1 < \lfloor x \rfloor \leq x \,}
$$

**(iii)** The **ceiling function**:

$$
\lceil x \rceil =
\begin{cases}
x\,; & x \ \text{an integer} \\[4pt]
\text{the nearest integer} > x\,; & x \ \text{not an integer}
\end{cases}
$$

$$
\lceil 5 \rceil = 5, \qquad \lceil -5 \rceil = -5, \qquad \lceil 5.1 \rceil = 6, \qquad \lceil -5.3 \rceil = -5 \qquad \square
$$

## Restriction of a function

For $f : A \to B$ and $C \subseteq A$:

$$
\boxed{\, g = f|_C \,}
$$

— the **restriction** of $f$ to the subset $C \subseteq A$:

$$
g(x) = f(x) \quad x \in C, \qquad \operatorname{Dom}(g) = C
$$

**Example:**

$$
f(x) = |x|, \qquad g_1 = f|_{\mathbb{R}_+}, \qquad g_2 = f|_{\mathbb{R}_- \cup \{0\}}
$$

![[lec02_p12_abs-restrictions.svg]]

## Glued (piecewise) functions

Often a function $f : \mathbb{R} \to \mathbb{R}$ is expressed by a non-homogeneous formula (a **"glued" function**):

$$
f(x) =
\begin{cases}
f_1(x) & x \geq a \\
f_2(x) & x < a
\end{cases}
$$

![[lec02_p12_glued-function.svg]]

(e.g. also $\chi_A(x)$ is glued).

$$
|x| =
\begin{cases}
x & x \geq 0 \\
-x & x < 0
\end{cases}
$$

Here $a = 0$, $f_1(x) = x$ and $f_2(x) = -x$.

## Image and preimage

For $f : A \to B$ and an arbitrary subset $A_1 \subset A$:

$$
f(A_1) = \{\, f(x) \in B \ : \ x \in A_1 \,\}
$$

— the **image of $A_1$** under the function $f$:

![[lec02_p13_image-doodle.svg|260]]

**Example:**

$$
f(x) = x^2, \qquad \operatorname{Dom}(f) = \mathbb{R}, \qquad A_1 = (0, 1] \subseteq \mathbb{R}
$$

$$
f(A_1) = (0, 1] \qquad \square
$$

![[lec02_p13_x2-image.svg|200]]

Similarly, for $B_1 \subset B$:

$$
f^{-1}(B_1) = \{\, x \in A \ : \ f(x) \in B_1 \,\}
$$

— the **preimage of $B_1$** under the function $f$:

![[lec02_p13_preimage-doodle.svg|260]]

**Example:**

$$
f(x) = x^2, \qquad B_1 = \{2\} \subseteq \mathbb{R}
$$

$$
f^{-1}(B_1) = \{-\sqrt{2}, \sqrt{2}\}. \qquad \square
$$

## Functions of several variables

Of course when e.g. the set is $B = \mathbb{R} \times \mathbb{R}$, then $f$ depends on 2 variables (or e.g. $A = \mathbb{N} \times \mathbb{N}$):

> **Note:** for $f$ to depend on 2 variables it is the *domain* that is the product, $A = \mathbb{R} \times \mathbb{R}$ — as in the examples below.

$$
f_1(x, y) = x^2 + y^2, \qquad f_1 : \underbrace{\mathbb{R}^2}_{\mathbb{R} \times \mathbb{R}} \to \mathbb{R}_+ \cup \{0\}
$$

Its graph is a **paraboloid**:

![[lec02_p14_paraboloid.svg|300]]

$$
f_2(m, n) = m \cdot n
$$

## Ways of representing functions

Functions can be represented:

- by an **explicit formula**, and/or
- by the **graph** of the function,
- for a **finite domain** — by giving all the values of the function at every argument.

## Worked example

**a)**

$$
X = \{\bullet, 1, \triangle\}, \qquad Y = \{0, 2, *, \square\}
$$

$$
f(\bullet) = 0, \qquad f(1) = \square, \qquad f(\triangle) = 2
$$

$f$ **is one-to-one**.

$f$ is **not** a mapping onto $Y$, because

$$
f(\,?\,) = * \qquad \text{has no argument } x \in X.
$$

For

$$
A = \{\bullet, 1\} \qquad f(A) = \{0, \square\}
$$

$$
B = \{0, \square\} \qquad f^{-1}(B) = \{\bullet, 1\}
$$

> **Note:** the handwritten notes give $f(A) = \{0, \square, 2\}$, but $2 = f(\triangle)$ and $\triangle \notin A$; for $A = \{\bullet, 1\}$ the image is $f(A) = \{0, \square\}$ (which is consistent with the very next line, $f^{-1}(\{0,\square\}) = \{\bullet, 1\}$).

$$
\operatorname{Dom}(f) = X, \qquad \operatorname{Im}(f) = \{0, \square, 2\}
$$

**b)** For a): $g = f^{-1}$ (the inverse function of $f$) exists, because $f$ is one-to-one.

$$
g = f^{-1} : \operatorname{Im}(f) \to \operatorname{Dom}(f)
$$

$$
\begin{aligned}
g(0) &= f^{-1}(0) = \bullet \\
g(\square) &= f^{-1}(\square) = 1 \\
g(2) &= f^{-1}(2) = \triangle
\end{aligned}
$$
