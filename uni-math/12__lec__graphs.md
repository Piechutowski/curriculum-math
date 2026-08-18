# Discrete Mathematics — Lecture 12: Graphs and Trees — Introduction

> Translated from the handwritten lecture notes `pdf/lect12-grafy-MatDysk.pdf`.

## Reminder: directed graphs

**Directed graph** (**digraph**) — objects and directed lines:

- $V(G)$ — the **vertices** (or simply $V$)
- $E(G)$ — the **edges** of the graph (or simply $E$)

$$
\gamma:\ E(G) \to V(G) \times V(G)
$$

$$
\gamma(e) = (p, q)
$$

where $p$ is the **start** of the edge and $q$ is the **end** of the edge.

The teacher's Fig. 1 — a small digraph with an edge and a loop:

![[lec12_p01_digraph-example.svg]]

$$
\begin{aligned}
V(G) &= \{a, b, c\} \\
E(G) &= \{e, f\} \\
\gamma(e) &= (a, b) \\
\gamma(f) &= (c, c)
\end{aligned}
$$

A **drawing** of a directed graph $G$ is a diagram consisting of points from the set $V(G)$ and arrows corresponding to the elements of the set $E(G)$ (see Fig. 1).

If $\gamma$ is **injective**, it means that for $\tilde{e} \neq \tilde{f}$ (different edges) $\gamma(\tilde{e}) \neq \gamma(\tilde{f})$ — i.e. there can be no **multiple edges**:

![[lec12_p02_multiple-edges.svg|300]]

Here $\gamma(\tilde{f}) = (a, b)$ and $\gamma(\tilde{e}) = (a, b)$ — so multiple edges occur.

## Paths in a directed graph

A **path** (droga) in a directed graph:

- a) $e_1, e_2, \ldots, e_n$, $\quad e_i \in E(G)$
- b) $x_1, x_2, \ldots, x_n, x_{n+1}$, $\quad x_i \in V(G)$
- c) $\gamma(e_i) = (x_i, x_{i+1})$

— a sequence of edges such that the end of one edge is the start of the next.

$e_1 e_2 \ldots e_n$ — a **path** of **length $n$** **from** the vertex $x_1$ to $x_{n+1}$.

A path is **closed** if $x_1 = x_{n+1}$.

**Cycle** — a closed path where $x_1, x_2, \ldots, x_n$ are all distinct.

A directed graph having no cycles is called **acyclic**.

A **path** is **acyclic** if the directed graph containing its vertices and edges is acyclic.

### Example

**a)** A digraph with multiple edges and a loop:

![[lec12_p03_digraph-paths.svg]]

$$
V(G) = \{v, w, x, y, z\}, \qquad E(G) = \{\tilde{a}, \tilde{b}, \tilde{c}, \tilde{d}, \tilde{e}, \tilde{f}, \tilde{g}\}
$$

The vertex sequence $y\,z\,z\,z$ corresponds only to the path $\tilde{f}\tilde{g}\tilde{g}$, while the vertex sequence $y\,v\,w\,z$ corresponds to the paths $\tilde{e}\tilde{a}\tilde{c}$ and $\tilde{e}\tilde{b}\tilde{c}$. If there are no multiple edges, a path is uniquely determined by its sequence of vertices.

**b)** A digraph with no multiple edges:

![[lec12_p04_digraph-cycles.svg]]

The path $\tilde{f}\tilde{g}\tilde{a}\tilde{e}$ determines the vertex sequence $z\,y\,w\,z\,x$ — this vertex sequence itself uniquely determines the path; $\gamma$ is injective (no multiple edges):

$$
\begin{aligned}
\gamma(\tilde{a}) &= (w, z) & \gamma(\tilde{d}) &= (z, z) \\
\gamma(\tilde{b}) &= (w, x) & \gamma(\tilde{e}) &= (z, x) \\
\gamma(\tilde{c}) &= (x, z) & \gamma(\tilde{f}) &= (z, y) \\
\gamma(\tilde{h}) &= (y, x) & \gamma(\tilde{g}) &= (y, w)
\end{aligned}
$$

The path $\tilde{a}\tilde{f}\tilde{g}$ is a **cycle**, because of the vertex sequence $w\,z\,y\,w$. Similarly, $\tilde{c}\tilde{f}\tilde{h}$ and $\tilde{c}\tilde{f}\tilde{g}\tilde{b}$ are cycles, with vertex sequences $x\,z\,y\,x$ and $x\,z\,y\,w\,x$. Also $\tilde{c}\tilde{e}$ and $\tilde{d}$ are cycles. But the path $\tilde{c}\tilde{f}\tilde{g}\tilde{a}\tilde{e}$ is **not a cycle**, because in $x\,z\,y\,w\,z\,x$ the vertex $z$ repeats! $\square$

## Undirected graphs

We now turn to undirected graphs.

**Undirected graph:**

- $V(G)$ — the set of vertices
- $E(G)$ — the set of edges

$$
\gamma:\ E(G) \to \{\, \{u, v\} : u, v \in V(G) \,\}
$$

— one-element ($u = v$) or two-element ($u \neq v$) subsets of $V(G)$.

The order $\{u, v\} = \{v, u\}$ plays no role (hence the graph is *undirected*).

The elements of $\gamma(e)$ are the **vertices** of $e$ (or its **ends**).

$e \neq f$ and $\gamma(e) = \gamma(f)$ — **multiple edges**.

**Loop** — an edge with a single end.

If there are no multiple edges ($\gamma$ injective), then $\gamma(e)$ determines the edge uniquely: $\{u, v\}$, or $\{u, u\} = \{u\}$ for a loop.

## Drawing of an undirected graph

A drawing of an (undirected) graph $G$ consists of points corresponding to the vertices of $G$ and arcs corresponding to the edges $\gamma(e) = \{u, v\}$.

### Example

A drawing of a graph with a pair of multiple edges and a loop:

![[lec12_p06_undirected-example.svg]]

$$
V(G) = \{x, y, w, z\}, \qquad E(G) = \{a, b, c, d, e, f\}
$$

$$
\begin{aligned}
\gamma(\tilde{a}) &= \{x, y\} \\
\gamma(\tilde{b}) &= \{y, w\} &&\leftarrow\ \text{no injectivity} \\
\gamma(\tilde{c}) &= \{y, w\} &&\leftarrow\ \text{(multiple edges)} \\
\gamma(\tilde{d}) &= \{y, z\} \\
\gamma(\tilde{e}) &= \{z, w\} \\
\gamma(\tilde{f}) &= \{w\} &&\leftarrow\ \text{loop}
\end{aligned}
$$

## Paths in an undirected graph

A **path of length $n$ from vertex $u$ to $v$** is a sequence:

- $e_1 e_2 \ldots e_n$ (edges)
- $x_1 x_2 \ldots x_{n+1}$ (vertices)
- $\gamma(e_i) = \{x_i, x_{i+1}\}$
- $u = x_1$, $v = x_{n+1}$

A path **between** $u$ and $v$ is a path from $u$ to $v$ or from $v$ to $u$.

If $u = v$, the path is **closed**.

**Simple path:** all edges are distinct.

**Cycle:**

- a simple path,
- a closed path,
- the vertices $x_1, x_2, \ldots, x_n$ are distinct.

A graph containing no cycles is called **acyclic**.

A **path** is **acyclic** if the subgraph consisting of its vertices and edges is acyclic.

### Example

**a)** A single edge:

![[lec12_p08_single-edge.svg|240]]

$$
V(G) = \{x_1, x_2\}, \qquad E(G) = \{e\}, \qquad \gamma(e) = \{x_1, x_2\}
$$

The path $ee$ together with the vertex sequence $x_1 x_2 x_1$ is a closed path — but it is **not a cycle**! It is not a simple path (we have $e$ twice). Similarly $ee$ with $x_2 x_1 x_2$ is not a cycle. $G$ is an acyclic graph.

**b)** Two parallel edges:

![[lec12_p08_double-edge.svg|240]]

$$
V(G) = \{x_1, x_2\}, \qquad E(G) = \{e, f\}, \qquad \gamma(e) = \gamma(f) = \{x_1, x_2\}
$$

The path $ef$ with $x_1 x_2 x_1$ **is** now a cycle (because $e \neq f$), and similarly $ef$ with $x_2 x_1 x_2$ is also a cycle.

> **Note:** the teacher wrote $V(E) = \{e, f\}$ — an evident notation slip for $E(G) = \{e, f\}$.

**c)** A larger graph:

![[lec12_p09_undirected-cyclic.svg]]

$$
V(G) = \{v, u, w, y, x, z\}, \qquad E(G) = \{\tilde{e}, \tilde{f}, \tilde{g}, \tilde{h}, \tilde{m}, \tilde{i}, \tilde{l}, \tilde{j}, \tilde{k}\}
$$

$$
\begin{aligned}
\gamma(\tilde{e}) &= \{v, u\} & \gamma(\tilde{k}) &= \{w, y\} \\
\gamma(\tilde{g}) &= \{u, w\} & \gamma(\tilde{m}) &= \{w, z\} \\
\gamma(\tilde{f}) &= \{v, w\} & \gamma(\tilde{i}) &= \{y, x\} \\
\gamma(\tilde{h}) &= \{w, x\} & \gamma(\tilde{j}) &= \{x, z\} \\
& & \gamma(\tilde{l}) &= \{y, z\}
\end{aligned}
$$

The path $\tilde{e}\tilde{f}\tilde{h}\tilde{i}\tilde{k}\tilde{g}$ (of length 6) with the vertex sequence $u\,v\,w\,x\,y\,w\,u$ is closed and simple, but it is **not** a cycle ($w$ repeats!). But $u\,v\,w\,u$ ($\tilde{e}\tilde{f}\tilde{g}$) **is** a cycle. The graph is cyclic. $\square$

## Theorems on paths and cycles

Example a) showed the exceptional situation (here $n = 2$) where:

- the vertices were distinct,
- the path was closed,

but it was **not simple**!

**Theorem:** Every closed path $e_1, \ldots, e_n$ of length at least 3 with distinct vertices $x_1, x_2, \ldots, x_n$ is a **cycle**.

**Theorem:** A path has all vertices distinct $\iff$ it is simple and acyclic.

**Theorem:** If $u \neq v$, $u, v \in V(G)$ (vertices), and there exists a path from $u$ to $v$, then there exists a **simple and acyclic** path from $u$ to $v$.

> The condition "there exists a path from $u$ to $v$" was added by the teacher in pen to both of the last two theorems.

**Theorem:** If $u \neq v$ in an **acyclic** graph ($u, v \in V(G)$ — vertices), then there exists **at most one** simple path in $G$ from $u$ to $v$.

**Example:** the theorem is false for cyclic graphs. In the graph from example c):

![[lec12_p11_two-simple-paths.svg]]

$$
gh \quad \text{and} \quad gki \quad \longleftarrow \ \text{simple paths from } u \text{ to } x
$$

## Isomorphism of graphs

Many graphs are similar to one another. The notion of **isomorphism**:

**(i)** $G$ and $H$ graphs (**without multiple edges**).

An **isomorphism** $\alpha$ of the graph $G$ onto $H$:

$$
\alpha:\ V(G) \to V(H)
$$

a **bijection** (an injection and a surjection) such that if $\{u, v\}$ is an edge in $G$, then $\{\alpha(u), \alpha(v)\}$ is an edge in $H$ (and conversely).

$$
G \simeq H
$$

— if such an $\alpha$ exists, the two graphs are **isomorphic** (this is the notation for graph isomorphism!).

$\alpha$ an isomorphism $\;\Rightarrow\;$ $\alpha^{-1}$ is one too!

### Example

**a)** Two isomorphic graphs $G$ (left) and $H$ (right):

![[lec12_p12_isomorphic-graphs.svg]]

Both graphs are isomorphic:

$$
\begin{aligned}
\alpha(t) &= t' & \alpha(y) &= y' & \alpha(z) &= z' \\
\alpha(u) &= u' & \alpha(w) &= w' & \alpha(x) &= x'
\end{aligned}
$$

$\alpha: V(G) \to V(H)$ is a bijection (1‑1), and moreover $\{\alpha(p), \alpha(q)\}$ is an edge in $E(H)$ $\iff$ $\{p, q\} \in E(G)$. So $G \simeq H$.

**b)** Two isomorphic multigraphs (loops, multiple edges and pendant edges drawn differently):

![[lec12_p13_multigraph-isomorphism.svg]]

but $G_1 \not\simeq G$, because $G$ has no loops (see the previous example).

**(ii)** For graphs **with multiple edges**, isomorphism requires **2** maps:

$$
\begin{aligned}
\alpha&:\ V(G) \xrightarrow{\ 1\text{-}1^{*}\ } V(H) \\
\beta&:\ E(G) \xrightarrow{\ 1\text{-}1\ } E(H)
\end{aligned}
$$

$$
e \text{ joins } u \text{ with } v \quad (e \in E(G),\ u, v \in V(G))
$$

$$
\Updownarrow
$$

$$
\beta(e) \text{ joins } \alpha(u) \text{ with } \alpha(v) \quad (\beta(e) \in E(H),\ \alpha(u), \alpha(v) \in V(H))
$$

$^{*}$ 1‑1 means a bijection ("mutually one-to-one").

## Invariants of isomorphism

Thus 2 graphs are isomorphic $\iff$ they have the same drawing (except for the labels of the edges and vertices).

Obviously, an **invariant** of an isomorphism of 2 graphs is:

- the same number of vertices ($\alpha$ is 1‑1),
- the same number of edges ($\beta$ is 1‑1),

but also:

- the number of loops,
- the number of simple paths of a given length.

Often it is useful to **count** the edges meeting at a chosen vertex.

## Degree of a vertex

**Definition:** the **degree of a vertex** $v$ ($\deg(v)$) is the number of two-vertex edges with $v$ as one of the endpoints **plus twice** the number of loops at the vertex $v$.

The number of vertices of degree $k$ in the graph $G$ (denoted $D_k(G)$):

- is an **invariant** of graph isomorphism;
- another invariant is the sequence of numbers of vertices of consecutive degrees $(D_0(G), D_1(G), D_2(G), \ldots)$.

### Example

**a)** For the graph $G$ from the isomorphism example:

![[lec12_p15_degree-sequence.svg|300]]

$$
(D_0(G), D_1(G), D_2(G), D_3(G), \ldots) = (0,\ 0,\ 2,\ 4,\ 0,\ 0, \ldots)
$$

$D_2(G) = 2$ (vertices $z, u$), $D_3(G) = 4$ (vertices $t, y, w, x$) — an edge reaches every vertex (at least 2 and no more than 3). The same holds for the isomorphic graph $H$ from the previous example.

**b)** But consider the three graphs:

![[lec12_p16_cubic-graphs.svg]]

$$
H_1 \simeq H_2 \not\simeq H_3
$$

Each $H_i$, $i = 1, 2, 3$, has 8 vertices of degree 3, so:

$$
(0, 0, 0, 8, 0, \ldots)
$$

but, as we can see, it does **not** follow from this that they must be isomorphic! $\square$

> The teacher's pen note: $H_1$ and $H_2$ contain a cycle $\triangle$ (a triangle), while $H_3$ has no $\triangle$ cycles — which is why $H_3$ cannot be isomorphic to them.

Graphs all of whose vertices have the same degree (not necessarily isomorphic graphs) are called **REGULAR**; here the degree was 3.

## Complete graphs

Graphs which are:

- without loops,
- without multiple edges,
- with every vertex joined to every other vertex by an edge,

are called **COMPLETE** graphs.

A complete graph on $n$ vertices has all vertices of degree $n - 1$ (so it is regular). All complete graphs on $n$ vertices are isomorphic — they are denoted by the symbol $K_n$.

### Example

The complete graphs $K_1, \ldots, K_5$, and a regular graph $G$ that is not complete:

![[lec12_p17_complete-graphs.svg]]

Here (in $G$) every vertex has degree 3, but $G$ is **not** complete (loops!). $\square$

## The handshaking theorem

**Theorem:**

**a)** The sum of the degrees of the vertices of a graph is 2 times the number of edges:

$$
\sum_{v \in V(G)} \deg(v) = 2 \cdot \overline{\overline{E(G)}}
$$

**b)** $D_1(G) + 2 D_2(G) + 3 D_3(G) + \ldots = 2\,\overline{\overline{E(G)}}$, i.e.:

$$
\sum_{i=0}^{\infty} i \, D_i(G) = 2\, \overline{\overline{E(G)}}
$$

### Example

**a)** The **complete** graph has $n$ vertices, each of degree $n-1$ (i.e. $D_{n-1}(G) = n$ and $D_i(G) = 0$ for $i \neq n-1$), and the number of edges equals the number of two-element subsets $\{v, w\}$ (since order does not matter for an undirected graph), and there are $\binom{n}{2}$ of these (see counting — combinatorics):

$$
\binom{n}{2} = \frac{n(n-1)}{2}
$$

**Check:**

**(a)**

$$
\underbrace{n}_{\substack{\text{number of} \\ \text{vertices}}} \cdot \underbrace{(n-1)}_{\substack{\text{degree} \\ \text{of each}}} \ \stackrel{?}{=}\ 2 \cdot \underbrace{\frac{n(n-1)}{2}}_{\substack{\text{number} \\ \text{of edges}}} \qquad \text{yes}
$$

**(b)**

$$
\sum_{i=0}^{\infty} i\, D_i(G) \ \stackrel{?}{=}\ 2\,\frac{n(n-1)}{2}
$$

$$
0 + 0 + \ldots + \underbrace{D_{n-1}(G)}_{n} \cdot (n-1) + 0 + \ldots \ \stackrel{?}{=}\ n(n-1)
$$

$$
n(n-1) = n(n-1) \qquad \text{yes.}
$$

**b)** For the multigraph below (the loop vertex has degree 4!):

![[lec12_p19_degree-example.svg|340]]

the sequence of numbers of vertices of consecutive degrees is:

$$
(0,\ 2,\ 0,\ 6,\ 1,\ 0,\ 0,\ \ldots)
$$

**(b)** $1 \cdot 2 + 3 \cdot 6 + 4 \cdot 1 = 2 \cdot 12 \ \Rightarrow\ 24 = 24$ — YES

**(a)** $1 + 1 + 3 + 3 + 3 + 3 + 3 + 3 + 4 = 24$, $\quad 24 = 24$ $\square$
