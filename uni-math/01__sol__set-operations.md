# Discrete Mathematics — Exercises 1: Solutions

> Translated from the handwritten lecture notes `pdf/Ćwicz1-RozwiązaniaMatDysk.pdf`.

## Exercise 1 — operations on finite sets

Let

$$
U = \{*, \square, 1, \blacksquare, \clubsuit, \pi, \heartsuit, \spadesuit\}, \qquad C = \{*, \clubsuit\},
$$

$$
B = \{\heartsuit, \square, \pi\} \quad \text{and} \quad A = \{*, \square, 1, \blacksquare, \clubsuit\}.
$$

> The teacher's hand-drawn doodle symbols (an asterisk, an open and a filled square, a little flower, a heart, a spade) are rendered here as $*, \square, \blacksquare, \clubsuit, \heartsuit, \spadesuit$.

**Union and intersection:**

$$
\text{(i)} \quad A \cup B = \{*, \square, 1, \blacksquare, \clubsuit, \heartsuit, \pi\}
\qquad\qquad
\text{(ii)} \quad A \cap B = \{\square\}
$$

**Differences:**

$$
A \setminus B = \{*, 1, \blacksquare, \clubsuit\} \qquad\qquad B \setminus A = \{\heartsuit, \pi\}
$$

Note that $A \setminus B \neq B \setminus A$. But always

$$
A \cup B = B \cup A \quad \text{and} \quad A \cap B = B \cap A.
$$

**Complements** (with respect to $U$):

$$
A' = U \setminus A = \{\pi, \heartsuit, \spadesuit\}, \qquad B' = U \setminus B = \{*, 1, \blacksquare, \clubsuit, \spadesuit\}
$$

**Symmetric difference** — using (i) and (ii):

$$
A \oplus B = (A \cup B) \setminus (A \cap B) = \{*, 1, \blacksquare, \clubsuit, \heartsuit, \pi\}
$$

**Inclusions:**

- $C \subseteq A$ — **yes**, because every element of $C$ is also in $A$.
- $C \subseteq B$ — **no**, because e.g. $* \in C$ but $* \notin B$.

**Cartesian products:**

$$
C \times B = \{(*, \heartsuit),\ (*, \square),\ (*, \pi),\ (\clubsuit, \heartsuit),\ (\clubsuit, \square),\ (\clubsuit, \pi)\}
$$

$$
B \times C = \{(\heartsuit, *),\ (\heartsuit, \clubsuit),\ (\square, *),\ (\square, \clubsuit),\ (\pi, *),\ (\pi, \clubsuit)\}
$$

One can see that $C \times B \neq B \times C$.

**Power sets:**

$$
\mathcal{P}(\emptyset) = \{\emptyset\}
$$

$$
\mathcal{P}(B) = \{\, \underbrace{\emptyset}_{\substack{\text{zero-el.} \\ \text{subset}}},\ \underbrace{\{\heartsuit\}, \{\square\}, \{\pi\}}_{\text{1-el. subsets}},\ \underbrace{\{\heartsuit, \square\}, \{\heartsuit, \pi\}, \{\square, \pi\}}_{\text{2-el. subsets}},\ \underbrace{\{\heartsuit, \square, \pi\}}_{\text{3-el. subset}} \,\}
$$

## Exercise 2 — intervals on the real line

$$
A = \{-1\} \cup \left(-\tfrac{1}{2}, \tfrac{2}{3}\right], \qquad B = [0, 1)
$$

The teacher's sketch of $A$ on the axis $\mathbb{R}$ and $B$ on the axis $\mathbb{R}$:

![[sol01_p02_number-lines.svg]]

$$
\begin{aligned}
A \cup B &= \{-1\} \cup \left(-\tfrac{1}{2}, 1\right) \\
A \cap B &= \left[0, \tfrac{2}{3}\right] \\
A \setminus B &= \{-1\} \cup \left(-\tfrac{1}{2}, 0\right) \\
B \setminus A &= \left(\tfrac{2}{3}, 1\right) \\
A' &= (-\infty, -1) \cup \left(-1, -\tfrac{1}{2}\right] \cup \left(\tfrac{2}{3}, +\infty\right) \\
B' &= (-\infty, 0) \cup [1, +\infty)
\end{aligned}
$$

### $B \times A$ in the plane

Here on the horizontal axis "we have the set $B$" and on the vertical axis "we have the set $A$"; the product $B \times A = [0,1) \times \bigl(\{-1\} \cup (-\tfrac{1}{2}, \tfrac{2}{3}]\bigr)$ is a half-open rectangle together with a horizontal segment at height $y = -1$:

![[sol01_p02_bxa-product.svg]]

### $A \times B$ in the plane

Now the first factor $A$ lies on the horizontal axis and $B$ on the vertical one; $A \times B$ is a half-open rectangle together with a vertical segment at $x = -1$:

![[sol01_p02_axb-product.svg]]

> In the manuscript this second sketch reuses the axis annotations of the first one (set $A$ on the vertical axis, set $B$ on the horizontal), although the plotted points ($-1$, $-\tfrac{1}{2}$, $\tfrac{2}{3}$ on the horizontal axis) show that for $A \times B$ the axes are the other way around. The labels are corrected here.

## Exercise 3 — a disc and a half-plane

$$
A = \{(x, y) : x^2 + y^2 \leq 1\}, \qquad B = \{(x, y) : x \geq 0\}
$$

$A$ is the closed disc of radius $R = 1$ centered at the origin, and $B$ is the closed right half-plane:

![[sol01_p03_disc-halfplane.svg]]

> In the manuscript the boundary inequality wobbles between strict and non-strict ($x^2 + y^2 < 1$ vs. $\leq 1$) in the formulas below, and the last page swaps the axis letters. Everything here is normalized to the definition $A = \{x^2 + y^2 \leq 1\}$ (a **closed** disc): solid lines and filled dots mark included boundaries, dashed lines and open dots excluded ones.

**Union** — analytically:

$$
A \cup B = \{(x, y) : x^2 + y^2 \leq 1 \ \text{ or } \ x \geq 0\}
$$

and geometrically — the whole right half-plane together with the left half of the disc:

![[sol01_p03_union.svg]]

**Intersection** — analytically:

$$
A \cap B = \{(x, y) : x^2 + y^2 \leq 1 \ \text{ and } \ x \geq 0\}
$$

and geometrically — the closed right half-disc:

![[sol01_p03_intersection.svg]]

**Difference** — analytically:

$$
A \setminus B = \{(x, y) : x^2 + y^2 \leq 1 \ \text{ and } \ x < 0\}
$$

and geometrically — the left half-disc without the diameter on the $y$-axis:

![[sol01_p03_difference.svg]]

**The other difference** — geometrically, the right half-plane with the right half-disc cut out:

![[sol01_p04_b-minus-a.svg]]

$$
B \setminus A = \{(x, y) : x \geq 0 \ \text{ and } \ x^2 + y^2 > 1\}
$$

The teacher sketched $A \cup B$ once more on the last page:

![[sol01_p04_union-region.svg|300]]

**Final remark:** $A \times B$ cannot be visualized, because $A$ is two-dimensional and so is $B$ $\Rightarrow$ $A \times B$ is **four-dimensional**.
