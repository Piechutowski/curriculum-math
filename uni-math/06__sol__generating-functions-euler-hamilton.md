# Discrete Mathematics — Exercises 6: Solutions

> Translated from the handwritten lecture notes `pdf/Ćwicz6-RozwiązaniaMatDysk.pdf`.

## Problem 1 — a recurrence solved with generating functions

Solve:

$$
\begin{cases}
a_{n+2} - 5a_{n+1} + 6a_n = 2 \\
a_0 = 3, \quad a_1 = 7
\end{cases}
$$

Multiply both sides by $x^{n+2}$:

$$
a_{n+2}x^{n+2} - 5a_{n+1}x^{n+2} + 6a_n x^{n+2} = 2x^{n+2}
$$

and sum over $n$:

$$
\sum_{n=0}^{\infty}\left( a_{n+2}x^{n+2} - 5a_{n+1}x^{n+2} + 6a_n x^{n+2} \right) = \sum_{n=0}^{\infty} 2x^{n+2}
$$

$$
\sum_{n=0}^{\infty} a_{n+2}x^{n+2} - 5x\sum_{n=0}^{\infty} a_{n+1}x^{n+1} + 6x^2\sum_{n=0}^{\infty} a_n x^n = 2x^2 \sum_{n=0}^{\infty} x^n
$$

**Change of indices** ($n+2 = n'$ in the first sum, $n+1 = n''$ in the second):

$$
\sum_{n=2}^{\infty} a_n x^n - 5x\sum_{n=1}^{\infty} a_n x^n + 6x^2\sum_{n=0}^{\infty} a_n x^n = 2x^2 \sum_{n=0}^{\infty} x^n
$$

Now complete each sum down to $n=0$ (in the first sum add and subtract $a_0x^0$ and $a_1x^1$; in the second add and subtract $a_0x^0$, times $5x$). Denoting $f(x) = \sum_{n=0}^{\infty} a_n x^n$ and using $\sum_{n=0}^{\infty} x^n = \frac{1}{1-x}$:

$$
\underbrace{\sum_{n=0}^{\infty} a_n x^n}_{f(x)} - (\underbrace{a_1}_{7}x + \underbrace{a_0}_{3}) - 5x\underbrace{\sum_{n=0}^{\infty} a_n x^n}_{f(x)} + 5x\underbrace{a_0}_{3} + 6x^2\underbrace{\sum_{n=0}^{\infty} a_n x^n}_{f(x)} = 2x^2\underbrace{\sum_{n=0}^{\infty} x^n}_{\frac{1}{1-x}}
$$

$$
f(x) - (7x+3) - 5x f(x) + 15x + 6x^2 f(x) = \frac{2x^2}{1-x}
$$

$$
f(x)\left[1 - 5x + 6x^2\right] + 15x - 7x - 3 = \frac{2x^2}{1-x}
$$

$$
f(x)\left[1 - 5x + 6x^2\right] = 3 - 8x + \frac{2x^2}{1-x}
$$

$$
\Rightarrow \quad f(x) = \frac{3-8x}{1-5x+6x^2} + \frac{2x^2}{(1-x)(1-5x+6x^2)}
= \frac{(3-8x)(1-x) + 2x^2}{(1-5x+6x^2)(1-x)}
= \frac{10x^2 - 11x + 3}{(1-5x+6x^2)(1-x)}
$$

Factor the quadratic:

$$
\Delta = 25 - 4\cdot 6\cdot 1 = 1, \qquad x_1 = \frac{5-1}{12} = \frac{1}{3}, \qquad x_2 = \frac{5+1}{12} = \frac{1}{2}
$$

$$
\begin{aligned}
1 - 5x + 6x^2 &= 6\left(x - \tfrac{1}{2}\right)\left(x - \tfrac{1}{3}\right) \\
&= 3\cdot 2\left(x - \tfrac{1}{2}\right)\left(x - \tfrac{1}{3}\right) \\
&= (2x-1)(3x-1) \\
&= (1-2x)(1-3x)
\end{aligned}
$$

$$
f(x) = \frac{10x^2 - 11x + 3}{(1-2x)(1-3x)(1-x)}
$$

Now we do a **partial fraction decomposition**:

$$
\frac{10x^2 - 11x + 3}{(1-x)(1-2x)(1-3x)} = \frac{A}{1-x} + \frac{B}{1-2x} + \frac{C}{1-3x}
$$

$$
= \frac{A(1-2x)(1-3x)}{(1-x)(1-2x)(1-3x)}
+ \frac{B(1-x)(1-3x)}{(1-x)(1-2x)(1-3x)}
+ \frac{C(1-x)(1-2x)}{(1-x)(1-2x)(1-3x)}
$$

Expanding the numerators:

$$
\begin{aligned}
A\left[1 - 3x - 2x + 6x^2\right] &= A\left[1 - 5x + 6x^2\right] \\
B\left[1 - 3x - x + 3x^2\right] &= B\left[1 - 4x + 3x^2\right] \\
C\left[1 - 2x - x + 2x^2\right] &= C\left[1 - 3x + 2x^2\right]
\end{aligned}
$$

Adding them up:

$$
A + B + C + x\left[-5A - 4B - 3C\right] + x^2\left[6A + 3B + 2C\right]
$$

This must equal the numerator of the left-hand side, $3 + x\cdot(-11) + x^2\cdot 10$, hence:

$$
\Rightarrow \quad
\begin{cases}
A + B + C = 3 \\
-5A - 4B - 3C = -11 \\
6A + 3B + 2C = 10
\end{cases}
$$

with solution:

$$
A = 1, \qquad B = 0, \qquad C = 2
$$

$$
\Rightarrow \quad f(x) = \frac{2}{1-3x} + \frac{1}{1-x} = \sum_{n=0}^{\infty} a_n x^n \qquad (\bullet)
$$

Each of the two fractions also has a series expansion — the **geometric series**:

$$
\frac{a_1}{1-q} = a_1 + a_1 q + a_1 q^2 + \ldots \qquad |q| < 1
$$

So:

$$
2\cdot\frac{1}{1-3x} = 2\left(1 + 3x + 3^2x^2 + 3^3x^3 + \ldots\right) = 2\sum_{n=0}^{\infty} 3^n x^n
$$

$$
\frac{1}{1-x} = \sum_{n=0}^{\infty} x^n
$$

Substituting into $(\bullet)$:

$$
2\sum_{n=0}^{\infty} 3^n x^n + \sum_{n=0}^{\infty} x^n = \sum_{n=0}^{\infty} a_n x^n
$$

$$
\sum_{n=0}^{\infty} \left(2\cdot 3^n + 1\right)x^n = \sum_{n=0}^{\infty} a_n x^n
\quad\Rightarrow\quad
\boxed{a_n = 2\cdot 3^n + 1}
$$

> Margin note: $a_n = \Theta(3^n)$.

## Problem 2 — checking the handshaking lemma

Vertices of an **undirected graph**:

$$
V(G) = \{x, y, z, w\}
$$

and its edges:

$$
E(G) = \{\bar a, \bar b, \bar c, \bar d, \bar f, \bar g, \bar h\}
$$

with the incidence function:

$$
\begin{aligned}
\gamma(\bar a) &= \{x,y\}, & \gamma(\bar b) &= \{x,y\}, & \gamma(\bar c) &= \{w,x\}, & \gamma(\bar d) &= \{w,y\}, \\
\gamma(\bar f) &= \{y,z\}, & \gamma(\bar g) &= \{y,z\}, & \gamma(\bar h) &= \{w,z\}. &&
\end{aligned}
$$

The graph redrawn from the teacher's sketch:

![[sol06_p05_multigraph-degrees.svg]]

Check:

$$
\sum_{v \in V(G)} \deg(v) \stackrel{?}{=} 2\cdot|E(G)|
$$

$$
\underbrace{(3_x + 5_y + 3_w + 3_z)}_{14} = \underbrace{2\cdot 7}_{14}, \qquad 14 = 14 \quad \text{YES}
$$

> There are multiple edges, there is a cycle, there are no loops.

## Problem 3a — reading off the algebraic description of a digraph

From the drawing of the graph, recover its algebraic description.

The teacher's directed graph (arrows, one loop):

![[sol06_p05_digraph-with-loop.svg]]

$$
V = \{x, y, z\}, \qquad E(G) = \{\bar a, \bar b, \bar c, \bar d, \bar e\}
$$

$$
\begin{aligned}
\gamma(\bar a) &= (x, y) & \gamma(\bar b) &= (x, z) \\
\gamma(\bar c) &= (z, z) \ \text{— a loop} & \gamma(\bar d) &= (y, z) \\
\gamma(\bar e) &= (y, z) &&
\end{aligned}
$$

We have multiple edges and there is no cycle. But there is a loop.

## Problem 4 — Euler cycles and Euler paths

> **Euler cycle:** a closed walk passing exactly once through each edge!

**(i)** $G_1$ is a **connected** graph:

![[sol06_p06_euler-graph-g1.svg|340]]

We have exactly $2$ vertices of odd degree (the rest have even degree)
$\Rightarrow$ there exists an **Euler path which is not an Euler cycle**.

The order in which the edges can be traversed (dotted sketch):

![[sol06_p06_euler-path-order.svg|300]]

We start at one vertex of odd degree and finish at the other vertex of odd degree.

**(ii)** The graph $G_2$:

![[sol06_p06_graph-g2.svg|340]]

Here there are $2$ vertices with $\deg() = 3$ and $1$ vertex with $\deg() = 1$ — more than two vertices of odd degree. So there is **neither an Euler cycle nor an Euler path** which is not an Euler cycle.

**(iii)** The small multigraph, with the Euler cycle traced on the right:

![[sol06_p06_euler-cycle-multigraph.svg|420]]

$2$ vertices of even degree $= 2$, and $1$ vertex of even degree $= 4$ $\Rightarrow$ an **Euler cycle** exists.

## Problem 4′ — Hamiltonian criteria

> In the original notes this problem is numbered ④ again (a duplicate number).

**(i)** **Dirac's criterion.** If $G$ is without loops and multiple edges, $|V(G)| \geq 3$, and

$$
\deg(v) \geq \frac{n}{2} \quad \forall\, v \in V(G)
\quad\Rightarrow\quad G \ \text{is Hamiltonian}
$$

(we visit every vertex exactly once).

The graph $G_1$:

![[sol06_p07_hamilton-graph-g1.svg|340]]

$$
|V(G)| = 5 \geq 3
$$

no multiple edges, no loops, and for every vertex:

$$
\forall v \quad \deg(v) \geq \frac{5}{2}
$$

since $\deg(v)$ is either $3$ or $4$:

$$
\deg(v_i) = 3 \ \text{ for } i = 1,2,3,4, \qquad \deg(v_5) = 4
$$

— so $G_1$ is **Hamiltonian**.

**(ii)** **Edge-count criterion.** The same assumptions as above, and $G$ has at least

$$
\frac{1}{2}(n-1)(n-2) + 2
$$

edges $\Rightarrow$ $G$ is Hamiltonian.

$$
\frac{1}{2}(5-1)\cdot(5-2) + 2 = 8
$$

We have $8$ edges, so also by the 2nd criterion $G_1$ is **Hamiltonian**.

**(iii)** **Ore's criterion.** Assumptions as above; for all pairs $v, w$ of vertices **not joined by an edge**:

$$
\deg(v) + \deg(w) \geq n
$$

Here the non-adjacent pairs are $(v_1, v_3)$ and $(v_2, v_4)$, and both times:

$$
6 \geq 5 \quad \text{YES}
$$

We get "yes" twice $\Rightarrow$ $G_1$ is **Hamiltonian**.

## Problem 5 — extra problem: the house, the doors, and the locks

Notation:

- $\bar d_j$ — the $j$-th door, $j = 1, \ldots, 6$ (the **edges**),
- $P_i$ — the $i$-th room, $z$ — the outside (the **vertices**).

Question: can one walk through the house (passing exactly once through each door) and fit a lock in each of them?

**a)** The floor plan of the house:

![[sol06_p08_house-floorplan.svg|420]]

The corresponding graph:

![[sol06_p08_house-graph.svg|440]]

$G$ is **connected**; we look for an **Euler path which is not an Euler cycle**.

$$
\deg(P_2) = \deg(P_3) = \deg(z) = 2, \qquad \deg(P_1) = 3, \qquad \deg(P_4) = 3
$$

There are $2$ vertices of odd degree, so an Euler path (not being a cycle) **exists**, and it must start at one of the odd-degree vertices and finish at the other odd-degree vertex (so e.g. the inspection goes from room $P_1$ to $P_4$).

The route (dotted sketch — start in $P_1$, end in $P_4$):

![[sol06_p09_euler-walk-route.svg|440]]

Through the room $P_1$ we pass **twice**, but through each door only **once**! (There is **no Euler cycle** here!)

**b)** Now, if we want to clean the rooms (we have heavy equipment), then we want to pass **only once** through each room — and once through the yard — a **Hamiltonian cycle/path** (but we are looking for a cycle!). Here $n = 5$.

**(i)** Dirac:

$$
\deg(v_i) = 2 \ \text{or} \ 3, \qquad 2 \geq \frac{5}{2} \ \text{— no!}
$$

So criterion (i) is **not satisfied**, and we cannot conclude that a Hamiltonian cycle exists.

**(ii)**

$$
\frac{1}{2}(5-1)(5-2) + 2 = 8
$$

We would need at least $8$ edges — but we have $6$, so again we **cannot conclude** that a Hamiltonian cycle exists.

**(iii)** Ore (e.g. for the non-adjacent pair $P_3$, $P_2$):

$$
\deg(P_3) + \deg(P_2) \stackrel{?}{\geq} 5, \qquad 4 \geq 5 \ \text{— no}
$$

So once more we **cannot conclude** that a Hamiltonian cycle exists.

**But a Hamiltonian path exists:**

$$
z \; P_1 \; P_2 \; P_4 \; P_3
$$

A **cycle does not exist** if we assume that we clean only the rooms and the yard ($z$): then one cannot start from $z$, because it is impossible to pass only once through the rooms $P_i$ and return to $z$. And if we start from some $P_i$, then it is impossible to return to $P_i$ without passing twice through some other vertex.

> **Exercise:** check what would happen if we did not clean the yard.
