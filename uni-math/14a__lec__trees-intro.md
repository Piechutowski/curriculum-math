# Discrete Mathematics — Lecture 14: Trees — Introduction

> Translated from the handwritten lecture notes `pdf/lect14a-drzewawstep-a-MatDysk.pdf`.

## Definition — tree

**Tree** — a graph (**undirected**) which is:

- **connected** (every pair of vertices is joined by a path),
- **acyclic** (a graph without cycles; a **cycle** is a closed *simple path* in which the vertices $x_1, \ldots, x_n$ are distinct; a **simple path** is a path with all edges distinct).

A tree has no multiple edges and no loops. Every vertex is joined by a path with every other vertex.

**Example:** two trees, a) and b) — they are **isomorphic**:

![[lec14a_p01_two-trees.svg]]

To make their isomorphism visible, we can redraw them with labelled vertices:

![[lec14a_p02_labeled-trees.svg]]

## Spanning trees

For a connected graph we are interested in **minimal subgraphs** joining all the vertices. Such a subgraph must be acyclic (because it can always be reduced by a redundant edge while still keeping connectivity).

Such a subgraph is called a **spanning tree**.

**Definition:** $T$ (a subgraph) is a spanning tree of a **connected** graph $G$ if:

- a) $V(T) = V(G)$ — it contains every vertex of the graph $G$;
- b) it has the minimal number of edges (while preserving the connectivity of $G$).

**Remark:** there is no uniqueness of the minimal spanning tree.

**Example:** the graph below (a hexagon with a centre; edges $e_1, \ldots, e_{12}$):

![[lec14a_p03_wheel-graph.svg|340]]

It has over $300$ spanning trees. Here are 4 of them — all of them have $6$ edges:

![[lec14a_p03_spanning-trees.svg]]

## Theorem (existence of a spanning tree)

**Every finite connected graph has a spanning tree.**

Further on we will show how to construct them (i.e., we will give an algorithm).

**Example:** every tree with $7$ vertices is isomorphic to $T_1$ (a), or to $T_2$ (b), or to $T_3$ (c), or to $T_4$ (d), or to $T_5$ (e):

![[lec14a_p04_trees-T1-T5.svg]]

or to $T_6$ (f), or to $T_7$ (g), or to $T_8$ (h), or to $T_9$ (i), or to $T_{10}$ (j), or to $T_{11}$ (k):

![[lec14a_p05_trees-T6-T11.svg]]

The trees a)–k) are **not isomorphic to one another**. $\square$

> From the previous example with over $300$ spanning trees — each of them is (isomorphic to one) among a)–k).

## Example — a small graph and its spanning trees

$$
V = \{a, b, c, d\}, \qquad E = \{\, \{a,b\}, \{a,c\}, \{a,d\}, \{b,c\} \,\}, \qquad G = \{V, E\}
$$

— an **undirected graph**:

![[lec14a_p06_graph-G.svg|220]]

**This is not a tree** (it contains a cycle). But we can build spanning trees:

![[lec14a_p06_spanning-trees-abc.svg]]

They **are isomorphic** to the following subgraphs: a) to the "T-shape" (left), while b) and c) — to the path (right):

![[lec14a_p06_isomorphism-shapes.svg|350]]

The subgraph below is **not connected and is not a tree!**

![[lec14a_p06_subgraph.svg|180]]

$\square$

## Theorem (characterizations of a tree)

Let $G$ be a graph without loops and multiple edges, with

$$
\overline{\overline{V(G)}} > 1 \quad \text{(more than one vertex).}
$$

The following facts are equivalent:

- a) $G$ is a tree;
- b) every $2$ vertices are joined by exactly one simple path;
- c) $G$ is connected, but it stops being connected after the removal of any edge;
- d) $G$ is acyclic, but it stops being acyclic after the addition of any edge.

## Leaves

**Leaves** — vertices of degree $1$.

**Example** (see the previous example, the trees $T_1$–$T_{11}$):

- $T_1$ — 2 leaves;
- $T_2$, $T_3$, $T_8$ — 3 leaves;
- $T_4$, $T_5$, $T_6$, $T_7$ — 4 leaves;
- $T_9$, $T_{10}$ — 5 leaves;
- $T_{11}$ — 6 leaves. $\square$

> In the manuscript the "3 leaves" group is written as $T_2, T_3, T_4$, and $T_4$ then appears again in the "4 leaves" group; comparing with the drawings, the tree with 3 leaves is the spider $T_8$, while $T_4$ has 4 leaves.

**Lemma 1:** a tree having at least $1$ edge has at least $2$ leaves.

**Lemma 2:** a tree having $n$ vertices has exactly $n-1$ edges.

That is: for a (connected) graph with $n$ vertices, a spanning tree has $n-1$ edges.

## Theorem (trees and the number of edges)

Let $G$ be a graph without loops and multiple edges, with $\overline{\overline{G}} < \infty$ (finitely many vertices, say $n$). The following are equivalent:

- a) $G$ is a tree;
- b) $G$ is an acyclic graph with $n-1$ edges;
- c) $G$ is connected with $n-1$ edges.

An acyclic graph that is not necessarily connected is called a **forest**.

## Tree with a distinguished root

**Tree with a distinguished root** — a tree in which $1$ vertex is distinguished. We draw it so that **the root is at the top!**

Applications:

- data structures,
- visualization of relations.

**Example:** a medical decision tree with the distinguished root "temperature" (branches: high/low temperature, weak/strong pulse, skin colour flushed/normal); the leaves are the diagnosed states:

![[lec14a_p10_decision-tree.svg]]

Similarly one arranges keys for the classification of e.g. flowers, birds, mushrooms, etc.

A tree with a distinguished root can be treated in a natural way as a **directed graph**: we "grab" the tree by the root, and the force of gravity directs the edges downwards.

![[lec14a_p11_root-gravity.svg]]

**Example:** an undirected tree:

![[lec14a_p11_undirected-tree.svg|280]]

The same tree with the distinguished root $v$ (giving $T_v$) and with the distinguished root $x$ (giving $T_x$):

![[lec14a_p11_Tv-Tx.svg]]

With the distinguished root $z$ (giving $T_z$):

![[lec14a_p12_Tz.svg|280]]

The exact placement of the vertices does not matter — here is $T_v$ drawn again, differently:

![[lec14a_p12_Tv-alt.svg|360]]

Compare a) with b). $\square$
