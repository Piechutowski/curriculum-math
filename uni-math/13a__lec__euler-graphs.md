# Discrete Mathematics — Lecture 13a: Euler Graphs

> Translated from the handwritten lecture notes `pdf/lect13a-grafyEulera-MatDysk.pdf`.
> Every page carries the header *"Lecture 13 — Graphs and Trees"*; this first part covers Euler graphs (trees follow in the next part).

## The Königsberg bridge problem

**The Königsberg bridge problem** (moving along the edges of an *undirected* graph).

The teacher's plan of the town: the river with its two banks (town), the islands, and the seven bridges — problem $(*)$:

![[lec13a_p01_konigsberg-map.svg]]

$(*)$ Can one take a walk through the town shown above on the plan,

- crossing every bridge **exactly once**,
- and return home?

**Leonhard Euler** solved this problem in **1736**. He converted problem $(*)$ into a problem of **walking along an undirected graph**.

Each piece of land becomes a vertex, each bridge an edge — the labeled multigraph of the town:

![[lec13a_p02_konigsberg-multigraph.svg]]

That is, an **undirected graph**:

![[lec13a_p02_multigraph-plain.svg|300]]

Does there exist here a **closed walk** passing through every edge **exactly once**?

— an **Euler cycle** (the vertices are *not* necessarily visited only once).

More generally, an **Euler path** in $G$:

- a **simple path** (each edge used once),
- which contains **all** edges of $G$.

If an Euler path is **closed**, it is called an **Euler cycle** in $G$.

## Theorem 1 (Euler)

> **Theorem (Euler).** A graph which has an Euler cycle must have all vertices of **even degree**.

**Proof:**

**a)** Starting from an arbitrary vertex in the Euler cycle and passing to the next one, we *erase* every edge along which we move.

The doodle: the edge we have just traveled is erased — we will not walk along it again:

![[lec13a_p03_erase-edge.svg|320]]

**b)** Passing *through* a given vertex, we erase one edge arriving at it and one edge leaving it:

![[lec13a_p03_erase-through.svg|360]]

- In each case the erasing (passing through a vertex) lowers the degree of that vertex (i.e. the number of incident edges) by $2$.
- Finally, in order to have passed through **all** the edges, the degree of every vertex must end up equal to $0$.
- Hence we must have started from **even** degrees (because we finished at $0$, and each time we lowered the starting degree by $2$!). $\square$

> **Corollary:** For the Königsberg bridge graph **all** vertices have **odd** degree, namely $3$ or $5$ — there is **no Euler cycle**!

## Examples

### a) A graph given by vertex and edge sets

$$
V = \{a, b, c, d, e, f, g\}
$$

$$
E = \{\{a,b\}, \{a,c\}, \{a,d\}, \{a,e\}, \{b,c\}, \{d,e\}, \{a,f\}, \{f,g\}, \{a,g\}\}
$$

The edges are single (no multiple edges), so we may give them as pairs of points!

The graph $G$:

![[lec13a_p05_example-graph.svg|380]]

- degree $2$ — vertices $e, d, c, b, f, g$;
- degree $6$ — vertex $a$.

*(The teacher writes "index" for the degree of a vertex.)*

The **necessary condition** for the existence of an Euler cycle is satisfied, because the degrees are even numbers.

- The path $\{a,e\}, \{e,d\}, \{d,a\}$ is **not** an Euler cycle (it is simple and closed, but it does not use all the edges).
- The path $\{a,e\}, \{e,d\}$ is **not** an Euler cycle, because it is not closed.

The above simple paths are not Euler paths either.

> Later annotation in pen: **it has an Euler cycle!** E.g.
> $$\{a,e\}, \{e,d\}, \{d,a\}, \{a,c\}, \{c,b\}, \{b,a\}, \{a,f\}, \{f,g\}, \{g,a\}$$

### b) A graph with two odd vertices

The graph on the vertices $u, v, w, x$ with edges $a, \ldots, f$ and the loop $g$ at $x$:

![[lec13a_p06_graph-uvwx.svg|380]]

This graph has **no Euler cycle**, because $u$ and $v$ have odd degree (namely $3$).

But the path $b\,a\,c\,d\,g\,f\,e$ **is an Euler path**!

— it is not a cycle, because it is not a closed path!

### c) The square

![[lec13a_p06_square-cycle.svg|220]]

The graph has all vertices of degree $2$ (even) and obviously has the Euler cycles

$$
abcd, \quad bcda, \quad cdab \quad \text{and} \quad dabc.
$$

### d) Two nested squares

![[lec13a_p07_two-squares.svg|260]]

$G$ has all vertices of degree $2$, but it has **no Euler cycle**, because it consists of $2$ graphs $G_1$ and $G_2$ **not connected** with each other! $\square$

> **Corollary:** all vertices having even degree is a **necessary** but **not sufficient** condition for the existence of an Euler cycle in a graph!

Before introducing the sufficient condition (**CONNECTIVITY** of the graph), we state one more consequence of Euler's theorem.

## Theorem 2 and connectivity

> **Theorem 2.** A graph $G$ having an Euler path has either exactly $2$ vertices of odd degree, or no vertices of odd degree at all.

> **Definition.** A graph is **connected** if every pair of distinct vertices is joined by a path in this graph.
>
> A connected subgraph $H \subseteq G$ of a graph $G$ is called a **component** of $G$ if it is not contained in any larger connected subgraph of $G$.

**Example:** three sketches — a) a connected graph; b) a graph that is not connected; c) a graph $G$ that is not connected: $G_1$ is one component of $G$, $G_2$ is the other component of $G$. $\square$

![[lec13a_p09_connectivity-examples.svg]]

> **Theorem (Euler's, part 2).** A finite **connected** graph in which every vertex has **even degree** has an **Euler cycle**.

## Corollary 1

> **Corollary 1:** A finite connected graph having **exactly two** vertices of odd degree has an **Euler path**.

**Example:**

**a)** The graph $G$ with edges $l_1, \ldots, l_8$:

![[lec13a_p10_euler-cycle-example.svg|380]]

- degrees of $a, b$ — equal to $2$,
- degrees of $c, d, e$ — equal to $4$,

all **even**, **and** the graph is connected

$$
\Downarrow
$$

it **has an Euler cycle**! E.g. $l_1\, l_2\, l_3\, l_4\, l_5\, l_6\, l_7\, l_8$.

**b)** The graph with edges $l_1, \ldots, l_6$:

![[lec13a_p10_euler-path-example.svg|360]]

It has **no Euler cycle**:

- $\deg(a) = 1$ — odd,
- $\deg(b) = 3$ — odd,
- $\deg(c) = \deg(d) = 4$,

so there are $2$ vertices of odd degree $\;\Rightarrow\;$ $l_1\, l_2\, l_3\, l_4\, l_5\, l_6$ is an **Euler path**.

## How to find an Euler path or cycle — Fleury's algorithm

> Later annotation in pen: *for a connected graph.*

**Step 1:** Choose an arbitrary vertex $v$ of **odd** degree, if one exists. Otherwise choose an arbitrary vertex $v$. Let $VS = v$ and let $ES = \emptyset$.

**Step 2:** If no edge leaves the vertex $v$ — **stop**.

**Step 3:** If **exactly one** edge leaving the vertex $v$ remains (say $e$, from $v$ to $w$), remove $e$ from $E(G)$ and $v$ from $V(G)$, and go to step 5.

**Step 4:** If **more than one** edge leaving the vertex $v$ remains, choose such an edge (let $e$ go from $v$ to $w$) that after its removal the graph **stays connected**. Then remove $e$ from $E(G)$.

**Step 5:** Append $w$ at the end of the sequence $VS$, append $e$ at the end of the sequence $ES$, replace $v$ by the vertex $w$, and go to step 2.

One can show that the above algorithm works correctly.

See p. 348 of **K. Ross & C. Wright, *Discrete Mathematics***.

## Worked example — Fleury's algorithm

The graph $(a)$, with edges $a, b, c, d, f, h, i$ and the loop $g$ at $y$:

![[lec13a_p13_fleury-graph.svg|460]]

$G$ is connected, and since e.g. $\deg(z) = 1$ (odd), $G$ has **no Euler cycle**.

Since $\deg(y) = 5$ (a loop counts **twice**) and all the remaining vertices have even degrees, by Corollary 1 it **has an Euler path**! (Because there are $2$ vertices of odd degree, and the rest have even degrees.)

We apply Fleury's algorithm:

- We choose one of the vertices of odd degree, e.g. $v = z$ (step 1):

$$
VS = z, \qquad ES = \emptyset
$$

- The only edge from $z$ is the edge $i$.
- We go to step 3 — we choose $e = i$ and $w = y$.
- We remove $i$ from $E(G)$ and $z$ from $V(G)$.
- We pass to step 5; then $VS = zy$ and $ES = i$.

The new graph after removing $i$ and $z$ (now $v = y$):

![[lec13a_p14_fleury-step1.svg|460]]

- We take (step 2) $v = y$.
- From the vertex $y$ three edges leave, so from step 2 we go to step 4.
- We may choose as the edge $f$ or $g$ or $h$. Let us choose $e = f$; $w = x$. We remove $f$ from $E(G)$.
- In step 5:

$$
VS = zyx, \qquad ES = if, \qquad v = x
$$

The new graph:

![[lec13a_p15_fleury-step2.svg|460]]

- We return to step 2.
- From the vertex $x$ three edges leave, so again $\rightarrow$ step 4.
- Now $e$ may be $a$ or $d$ (it **cannot** be the edge $h$, since the graph would not stay connected!).
- We choose $e = a$ (we remove $a$ from $E(G)$).
- Step 5:

$$
VS = zyxr, \qquad ES = ifa, \qquad v = r
$$

- The next $3$ moves are **forced**; each leads to step 3 and to the removal of a further edge together with its vertex.
- We obtain:

$$
VS = zyxrstx, \qquad ES = ifabcd, \qquad v = x
$$

with the graph:

![[lec13a_p16_fleury-step3.svg|460]]

- The last $2$ moves are also forced:

![[lec13a_p16_fleury-final.svg|560]]

Finally:

$$
VS = zyxrstxyy, \qquad ES = ifabcdhg \qquad \square
$$

## Trees (preview)

**Trees** — **acyclic** and **connected** graphs.
