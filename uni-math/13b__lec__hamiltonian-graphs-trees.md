# Discrete Mathematics — Lecture 13b: Hamiltonian Graphs and Trees

> Translated from the handwritten lecture notes `pdf/lect13b-GrafyHamil-Drzewa-MatDysk.pdf`.

> **Note:** the header of every scanned page reads "Wykład 15 — Drzewa" ("Lecture 15 — Trees"); the file is catalogued as lecture 13b. Page 14 of the scan appears twice — the content below is complete without repetition.

## Rooted trees

**A tree with a distinguished root:**

- **Leaf** — a vertex of degree $1$ (with the exception of the root).
- In the **directed-graph representation**: the **root** is the only vertex that is *not the end of any edge*; the **leaves** are the terminal nodes; the remaining vertices are the **internal nodes**.
- $(v, w)$ — such a pair, i.e. an edge of a tree with a distinguished root, means:
  - $v$ is the **parent** of $w$,
  - $w$ is a **child** of $v$.
- A parent may have several children.
- $w$ is a **descendant** of $v$ (where $w \neq v$) if $v$ is a vertex of the unique simple path from the root $r$ to the vertex $w$.

The unique simple path from the root $r$ down through $v$ to $w$ — so $w$ is a descendant of $v$:

![[lec13b_p02_descendant-path.svg|180]]

For any vertex $v$, the **subtree** rooted at $v$ is the tree $T_v$ consisting of the vertex $v$, all of its descendants, and the directed edges joining these vertices.

The teacher's example tree and three of its subtrees $T_v$, $T_w$, $T_s$:

![[lec13b_p02_tree-subtrees.svg]]

## Binary and $m$-ary trees

A tree with a distinguished root is a **binary tree** when every node has at most **two** children. We then have a *left* and a *right* child (this suggests an ordering).

**Tree with $m$-branching** ($m$-ary tree): every vertex (parent) has at most $m$ children.

If ($m \geq 2$) every vertex has **degree $m$ or $0$** (i.e. exactly $m$ children or none), then such a tree is called **regular**.

A regular tree with $m = 3$:

![[lec13b_p03_regular-tree.svg|300]]

## Level number and height

**Level number** of a vertex $v$: the length of the unique simple path from the root to $v$. **The root has level $0$.**

**Height of a tree:** the largest level number of a vertex.

A tree of height $4$:

![[lec13b_p04_tree-height4.svg|340]]

A **regular** tree with $m$-branching is **full** if all leaves have level number equal to the height of the tree.

Left: a binary tree that is *not full* (one leaf has level $1$); right: a *full* one:

![[lec13b_p04_full-vs-notfull.svg|420]]

## Ordered rooted trees

We can order the children, e.g. according to a linear order:

$$
ch_1 \preceq ch_2 \preceq ch_3 \preceq \ldots \preceq ch_n
$$

![[lec13b_p05_ordered-children.svg|450]]

We obtain an **ordered tree with a distinguished root**. The order runs from left to right.

**Example:**

**a)** Let $\Sigma$ be an (ordered) alphabet. We build a rooted tree with the empty word as the root. For a word $w \in \Sigma^*$ we create the children:

$$
\{\, wx : x \in \Sigma \,\}
$$

Since $\Sigma$ is ordered, we can order every set of children:

$$
wx \preceq wy \iff x \preceq y
$$

E.g. for $\Sigma = \{a, b\}$ — this is a **binary tree**:

![[lec13b_p06_word-tree-ab.svg]]

**b)** Similarly for $\Sigma = \{0, 1\}$ — a **binary tree**:

![[lec13b_p06_word-tree-01.svg]]

$\square$

## Hamiltonian paths and cycles

In the problem of the bridges of Königsberg (Euler cycles) the **edges** may be "used" **once**, but the **vertices** may be visited **multiple times**.

**Hamiltonian path:** a path that passes **exactly once** through **every vertex** of the graph.

**Hamiltonian cycle** — a closed Hamiltonian path (the last vertex has visit multiplicity $2$).

A graph with a Hamiltonian cycle is a **Hamiltonian graph**.

> A Hamiltonian path must be a *simple* path (i.e. each edge used only once!).

**Example:**

**a)** The pentagon graph: $v\,w\,x\,y\,z\,v$ is a Hamiltonian cycle. **b)** We add edges — the complete graph is also Hamiltonian:

![[lec13b_p08_hamilton-examples.svg]]

Every complete graph with $n \geq 3$ is Hamiltonian.

**(c)** $v\,w\,x\,y\,z$ is a Hamiltonian path, but there is no Hamiltonian cycle (the cycle $w\,x\,y\,z\,w$ does not pass through $v$). **(d)** This graph has no Hamiltonian path at all:

![[lec13b_p08_path-vs-none.svg|480]]

$\square$

## What the theory of Hamiltonian cycles can say

(Compare: Euler's theorem handles Euler cycles neatly.)

- A good **characterization** of connected graphs having Hamiltonian cycles **is not known**.
- Moreover, even if a Hamiltonian cycle exists in a given graph, **no efficient algorithm** for finding a Hamiltonian cycle is known.

The above problem is a special case of the **travelling salesman problem**:

$$
\underbrace{\text{Graph}}_{\substack{\text{cities} \\ \text{(vertices)}}} + \underbrace{\text{weights added to the edges}}_{\substack{\text{they denote distance, cost, working time} \\ \text{or another quantity to be minimized}}}
$$

**Goal:** to find the "shortest route" that visits each city exactly once.

Automatically, a solution of the travelling salesman problem gives a solution of the search for a Hamiltonian cycle (we put all weights $= 1$).

- Obviously, a Hamiltonian graph with $n$ vertices must have at least $n$ edges.

## Theorem (Dirac)

> If $G$ has no loops and no multiple edges, $|V(G)| = n \geq 3$, and
>
> $$
> \deg(v) \geq \frac{n}{2} \quad \forall\, v \in V(G),
> $$
>
> then $G$ is a Hamiltonian graph.

**Example:**

**a)** $G_0 = K_5$: here $|V(G)| = 5$ and $\deg(v) = 4$ for all $v \in V(G)$.

![[lec13b_p10_k5.svg|420]]

$$
4 \geq \frac{5}{2}
$$

There are no multiple edges and no loops $\Rightarrow$ the assumptions of the theorem are satisfied $\Rightarrow$ a Hamiltonian cycle exists (sketched on the right).

**b)** The graphs $G_1$ and $G_2$:

![[lec13b_p11_g1-g2.svg]]

Here $|V(G)| = 5$, but there exist vertices with $\deg(x) = 2$, and

$$
2 \geq \frac{5}{2} \quad \text{is false.}
$$

The assumptions of the theorem are **not** satisfied! Nevertheless, in both cases $G_1$ and $G_2$ **are** Hamiltonian graphs — their Hamiltonian cycles:

![[lec13b_p11_hamilton-cycles.svg|520]]

$\square$

The previous theorem imposes uniform conditions on all vertices. The next theorem requires that there be sufficiently many edges in total.

## Theorem (edge count)

> If $G$ has $n$ vertices, no loops and no multiple edges, and has at least
>
> $$
> \frac{1}{2}(n-1)(n-2) + 2
> $$
>
> edges $\Rightarrow$ $G$ is Hamiltonian.

**Example:**

**a)** For $G_1$ (see the previous example): $n = 5$,

$$
\frac{1}{2}(5-1)(5-2) + 2 = 8.
$$

Our graph must have at least $8$ edges (and it has exactly $8$) $\Rightarrow$ it is Hamiltonian.

**b)** For $G_2$ we have $7$ edges. The theorem cannot be applied — but $G_2$ is nevertheless Hamiltonian. $\square$

## Theorem (Ore)

> Let $|V(G)| = n \geq 3$, with no loops and no multiple edges. If
>
> $$
> \deg(v) + \deg(w) \geq n
> $$
>
> for **every pair of vertices $v$ and $w$ not joined by an edge** $\Rightarrow$ $G$ is Hamiltonian.

**Example:** For $G_2$:

![[lec13b_p13_g2.svg|280]]

For the pair $(v, z)$:

$$
\deg(v) + \deg(z) = 3 + 3 \geq 5
$$

For the pair $(w, x)$:

$$
\deg(w) + \deg(x) = 3 + 2 \geq 5
$$

For the pair $(x, y)$:

$$
\deg(x) + \deg(y) = 2 + 3 \geq 5
$$

$\Rightarrow$ by the last theorem, $G_2$ is Hamiltonian. $\square$

## Spanning trees

An algorithm that, for a **connected** graph, constructs a **spanning tree** starting from a chosen vertex $v$.

Otherwise (when the graph is **disconnected**), it constructs a spanning tree of the (connected) **component** of the graph $G$.

**ALGORITHM TREE($v$):**

1. Let $V := \{v\}$ and $E := \emptyset$.
2. While there exist edges in the graph $G$ joining vertices from the set $V$ with vertices that are $\notin V$, do:
   - choose such an edge $\{u, w\}$ joining a vertex $u \in V$ with a vertex $w \notin V$;
   - add the vertex $w$ to $V$ and the edge $\{u, w\}$ to $E$.

**Example:** the graph $G$ below is **not connected**; its components $G_1$ and $G_2$ are connected graphs:

![[lec13b_p15_graph-components.svg]]

**TREE(1)** — builds $T_1$, starting from $v = \{1\}$:

$$
V := \{1\}; \qquad E := \emptyset.
$$

Choose the edge $\{1, 6\}$:

$$
V = \{1, 6\}; \qquad E := \{\{1,6\}\}.
$$

Choose the edges $\{6,3\}, \{6,7\}, \{6,10\}$ (the order does not matter here):

$$
V = \{1, 6, 3, 7, 10\}; \qquad E = \{\{1,6\}, \{6,3\}, \{6,7\}, \{6,10\}\}.
$$

Choose the edge $\{7, 12\}$ (and **not** $\{3,7\}$, since both $3$ and $7 \in V$):

$$
V = \{1, 6, 3, 7, 10, 12\}; \qquad E = \{\{1,6\}, \{6,3\}, \{6,7\}, \{6,10\}, \{7,12\}\}
$$

for the component $G_1$. The spanning tree $T_1$:

![[lec13b_p15_spanning-tree-t1.svg|300]]

**TREE(2)** — builds $T_2$, starting from $v = \{2\}$:

$$
V := \{2\}; \qquad E := \emptyset.
$$

Choose the edges $\{2,5\}, \{2,8\}$:

$$
V := \{2, 5, 8\}; \qquad E := \{\{2,5\}, \{2,8\}\}.
$$

Choose the edges $\{5,9\}, \{5,11\}$:

$$
V := \{2, 5, 8, 9, 11\}; \qquad E := \{\{2,5\}, \{2,8\}, \{5,9\}, \{5,11\}\}.
$$

Choose the edge $\{8,4\}$:

$$
V := \{2, 5, 8, 9, 11, 4\}; \qquad E := \{\{2,5\}, \{2,8\}, \{5,9\}, \{5,11\}, \{8,4\}\}.
$$

The spanning tree $T_2$ of the component $G_2$:

![[lec13b_p16_spanning-tree-t2.svg|320]]

$\square$

## Theorem (correctness of TREE)

> The procedure TREE($v$) constructs a spanning tree of the component of the graph $G$ containing the vertex $v$.

## Minimum spanning trees

The problem of finding a spanning tree is especially important when the edges have assigned **weights** $w_{ij}$ (with $w_{ij} \geq 0$).

The problem consists in finding a spanning tree whose weight is less than or equal to that of any other spanning tree. Such a tree is a **MINIMUM SPANNING TREE**.

**Example:** a weighted triangle graph $G$ and its three spanning trees with total weights $25$, $30$ and $15$ — the last one is the minimum spanning tree:

![[lec13b_p17_mst-example.svg]]

There are $2$ algorithms:

- **Kruskal's algorithm**,
- **Prim's algorithm**,

together with $2$ theorems justifying their correctness. We omit them in this lecture.
