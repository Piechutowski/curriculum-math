# Discrete Mathematics — Exercises 6: Recurrences and Graphs

> Translated from the handwritten lecture notes `pdf/Ćwiczenia6.pdf`.

## Exercise 1 — explicit form of a recurrence via a generating function

Solve the implicit (recursive) form

$$
\begin{cases}
a_{n+2} - 5a_{n+1} + 6a_n = 2 \\
a_3 = 3, \quad a_1 = 7
\end{cases}
$$

with the help of a **generating function** (i.e. find the explicit formula for $a_n$).

> The handwritten subscripts of the initial conditions read $a_3 = 3$ and $a_1 = 7$. Note that with *consecutive* initial terms $a_0 = 3$, $a_1 = 7$ the problem has the clean solution $a_n = 2 \cdot 3^n + 1$, so the sheet most likely intends two consecutive initial values.

## Exercise 2 — drawing a multigraph from its edge table

Draw the picture of the (**undirected**) graph $G$, where

$$
V(G) = \{x, y, z, w\}, \qquad E(G) = \{\bar{a}, \bar{b}, \bar{c}, \bar{d}, \bar{f}, \bar{g}, \bar{h}\},
$$

and the (incidence) function $\gamma$ is given in the table:

| edge | $\bar{a}$ | $\bar{b}$ | $\bar{c}$ | $\bar{d}$ | $\bar{f}$ | $\bar{g}$ | $\bar{h}$ |
|---|---|---|---|---|---|---|---|
| $f(\text{edge})$ | $\{x,y\}$ | $\{x,y\}$ | $\{w,x\}$ | $\{w,y\}$ | $\{y,z\}$ | $\{y,z\}$ | $\{w,z\}$ |

Check the relation between the **number of edges** and the **sum of the degrees of the vertices**.

> The last three table cells are squeezed together in the handwriting; they read $\{y,z\}$, $\{y,z\}$, $\{w,z\}$, which is consistent with the handshaking check ($\sum_v \deg v = 14 = 2 \cdot 7$).

## Exercise 3 — Euler cycles and Euler paths

### a) From a picture to an algebraic description

From the graph expressed by the drawing, deduce (**"guess"**) its algebraic description. The teacher's sketch — vertices $x, y, z$; parallel edges $\bar{d}, \bar{e}$ between $y$ and $z$; and a loop $\bar{c}$ at $z$:

![[ex06_p01_multigraph-xyz.svg|350]]

> The sketch carries small arrow flicks on the strokes of $\bar{e}$ and at $y$; the surrounding exercises deal with undirected graphs, so these appear to be pen flourishes indicating how the edges were traced rather than edge directions.

### b) Which graph is Eulerian?

Which of the graphs has an **Euler cycle**, an **Euler path** which is not an Euler cycle, or has neither one nor the other? The three graphs (i) $G_1$, (ii) $G_2$, (iii) $G_3$:

![[ex06_p01_euler-graphs.svg]]

If it does, draw the corresponding Euler cycle / Euler path.

> The sheet labels the third graph "(iv)" — an evident slip for "(iii)".

## Exercise 4 — a Hamiltonian graph

Justify why the graph $G$ is **Hamiltonian**. Use **3 criteria** giving **sufficient conditions**. The graph $G$ on vertices $v_1, \ldots, v_5$:

![[ex06_p01_hamiltonian-graph.svg|380]]
