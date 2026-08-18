# Discrete Mathematics — Lecture 5: Counting (continued) — the Pigeonhole Principle

> Translated from the handwritten lecture notes `pdf/lect5-zliczZasDirichleta-MatDysk.pdf`.

## The inclusion–exclusion principle

For $n = 2$:

$$
|A_1 \cup A_2| = |A_1| + |A_2| - |A_1 \cap A_2|
$$

It has its generalizations, e.g. for $n = 3$:

$$
\begin{aligned}
|A_1 \cup A_2 \cup A_3| \stackrel{*}{=}\ & |A_1| + |A_2| + |A_3| \\
& - |A_1 \cap A_2| - |A_1 \cap A_3| - |A_2 \cap A_3| \\
& + |A_1 \cap A_2 \cap A_3|
\end{aligned}
$$

The teacher's sketch of three sets $A$, $B$, $C$ with their elements drawn as dots:

![[lec05_p01_three-sets-count.svg]]

Reading the counts off the picture:

$$
|A \cup B \cup C| = 11, \qquad |A \cap B \cap C| = 1
$$

$$
|A| = 5, \qquad |B| = 6, \qquad |C| = 5
$$

$$
|A \cap B| = 2, \qquad |A \cap C| = 1, \qquad |B \cap C| = 3
$$

Checking against the formula $(*)$ (with $|A_1 \cap A_2 \cap A_3| = 1$):

$$
\begin{aligned}
|A_1 \cup A_2 \cup A_3| &= 11 \\
&\stackrel{*}{=} (5 + 6 + 5) + 1 - (2 + 1 + 3) \\
&= 16 + 1 - 6 = 11
\end{aligned}
$$

— the same result.

### The case $n = 4$

$$
\begin{aligned}
|A_1 \cup A_2 \cup A_3 \cup A_4| =\
&\bigl(|A_1| + |A_2| + |A_3| + |A_4|\bigr) \quad \oplus \\[2pt]
-\ &\bigl(|A_1 \cap A_2| + |A_1 \cap A_3| + |A_1 \cap A_4| \\
&\ \ + |A_2 \cap A_3| + |A_2 \cap A_4| + |A_3 \cap A_4|\bigr) \quad \ominus \\[2pt]
+\ &\bigl(|A_1 \cap A_2 \cap A_3| + |A_1 \cap A_2 \cap A_4| \\
&\ \ + |A_1 \cap A_3 \cap A_4| + |A_2 \cap A_3 \cap A_4|\bigr) \quad \oplus \\[2pt]
-\ &\;|A_1 \cap A_2 \cap A_3 \cap A_4| \quad \ominus
\end{aligned}
$$

where $\oplus$ marks **intersections of an odd number of sets** (added) and $\ominus$ marks **intersections of an even number of sets** (subtracted).

### The general formula

In general, to find the number of elements of the set

$$
A_1 \cup A_2 \cup \ldots \cup A_n
$$

**Method 1:** one has to find the set $A_1 \cup \ldots \cup A_n$ itself and count its elements.

**Method 2:**

- we find all possible intersections built from $\{A_1, A_2, \ldots, A_n\}$;
- we add the results obtained for intersections of an **odd** number of sets;
- then we subtract the results obtained for intersections of an **even** number of sets.

$$
P(n): \qquad \left| \bigcup_{i=1}^{n} A_i \right| = \sum_{\emptyset \neq J \subseteq I} (-1)^{|J|+1} \left| \bigcap_{i \in J} A_i \right|, \qquad I = \{1, 2, \ldots, n\}, \quad \boxed{n \geq 2}
$$

### Proof (by induction on $n$)

**Base step, $n = 2$:** $\ |A \cup B| = |A| + |B| - |A \cap B|$ — OK.

Indeed, for $I = \{1, 2\}$ the non-empty subsets are $J = \{1\}$, $J = \{2\}$, $J = \{1, 2\}$, so the formula gives

$$
|A_1 \cup A_2| = (-1)^{1+1}|A| + (-1)^{1+1}|B| + (-1)^{2+1}|A \cap B| = |A| + |B| - |A \cap B|.
$$

**Inductive step $P(n) \Rightarrow P(n+1)$:** with $I = \{1, 2, \ldots, n\}$, denote $B = A_1 \cup A_2 \cup \ldots \cup A_n$ and apply the $n = 2$ case to $B \cup A_{n+1}$:

$$
\Bigl| \underbrace{A_1 \cup A_2 \cup \ldots \cup A_n}_{B} \cup A_{n+1} \Bigr|
= |A_1 \cup A_2 \cup \ldots \cup A_n| + |A_{n+1}| - \bigl|(A_1 \cup A_2 \cup \ldots \cup A_n) \cap A_{n+1}\bigr|
$$

By $P(n)$ applied to the first term, and by distributivity in the last term:

$$
= \sum_{\emptyset \neq J \subseteq I} (-1)^{|J|+1} \Bigl| \bigcap_{i \in J} A_i \Bigr| + |A_{n+1}|
- \bigl| (A_1 \cap A_{n+1}) \cup (A_2 \cap A_{n+1}) \cup \ldots \cup (A_n \cap A_{n+1}) \bigr|
$$

$$
= \sum_{\emptyset \neq J \subseteq I} (-1)^{|J|+1} \Bigl| \bigcap_{i \in J} A_i \Bigr| + |A_{n+1}|
+ (-1)\,\bigl| B_1 \cup B_2 \cup \ldots \cup B_n \bigr|,
\qquad B_i = A_i \cap A_{n+1}, \quad i \in I
$$

We apply induction (the formula $P(n)$) once more, to the sets $B_i$:

$$
= \sum_{\emptyset \neq J \subseteq I} (-1)^{|J|+1} \Bigl| \bigcap_{i \in J} A_i \Bigr| + |A_{n+1}|
+ (-1) \sum_{\emptyset \neq J \subseteq I} (-1)^{|J|+1} \Bigl| \bigcap_{i \in J} B_i \Bigr|
$$

$$
= \sum_{\emptyset \neq J \subseteq I} (-1)^{|J|+1} \Bigl| \bigcap_{i \in J} A_i \Bigr| + |A_{n+1}|
+ \sum_{\emptyset \neq J \subseteq I} (-1)^{|J|+2} \Bigl| \bigcap_{i \in J} \bigl(A_i \cap A_{n+1}\bigr) \Bigr|
$$

where, for the last sum, $(*)$:

$$
(*) \qquad
\bigcap_{i \in J} \bigl(A_i \cap A_{n+1}\bigr)
= \Bigl(\bigcap_{i \in J} A_i\Bigr) \cap A_{n+1}
= A_{n+1} \cap \bigcap_{i \in J} A_i
$$

Now the bookkeeping of signs:

- the middle term corresponds to $\tilde{J} = \{n+1\}$, $|\tilde{J}| = 1$: since $|\tilde{J}| + 1 = 2$ and $(-1)^2 = 1$, we may write $|A_{n+1}| = (-1)^{|\tilde{J}|+1} |A_{n+1}|$;
- a subset $J' = J \cup \{n+1\}$ with $\emptyset \neq J \subseteq I$ has $|J'| = |J| + 1$, so $(-1)^{|J'|+1} = (-1)^{|J|+2}$ — exactly the sign appearing in the last sum.

So the three parts together run over all non-empty subsets of $I \cup \{n+1\}$:

$$
= \underbrace{\sum_{\emptyset \neq J \subseteq I \cup \{n+1\}} (-1)^{|J|+1} \Bigl| \bigcap_{i \in J} A_i \Bigr|}_{P(n+1)}
$$

By mathematical induction, $P(n)$ is true for $n \geq 2$. $\square$

## Example (inclusion–exclusion)

$$
S = \{1, 2, 3, \ldots, 2000\}
$$

How many numbers are divisible by 9, or by 11, or by 13, or by 15?

$$
D_k = \{\, n \in S : n \text{ is divisible by } k \,\}
$$

We are looking for $|D_9 \cup D_{11} \cup D_{13} \cup D_{15}|$, where

$$
|D_k| = \left\lfloor \frac{2000}{k} \right\rfloor
$$

The single sets ($\oplus$ — an odd number of sets, added):

$$
|D_9| = 222, \qquad |D_{11}| = 181, \qquad |D_{13}| = 153, \qquad |D_{15}| = 133
$$

The pairwise intersections ($\ominus$ — an even number of sets, subtracted); note that $D_a \cap D_b = D_{\operatorname{lcm}(a,b)}$ — e.g. $117$ is the least common multiple of $9$ and $13$:

$$
\begin{aligned}
|D_9 \cap D_{11}| &= |D_{99}| = 20 \\
|D_9 \cap D_{13}| &= |D_{117}| = 17 \\
|D_9 \cap D_{15}| &= |D_{45}| = 44 \\
|D_{11} \cap D_{13}| &= |D_{143}| = 13 \\
|D_{11} \cap D_{15}| &= |D_{165}| = 12 \\
|D_{13} \cap D_{15}| &= |D_{195}| = 10
\end{aligned}
$$

The triple intersections ($\oplus$ — odd, added):

$$
\begin{aligned}
|D_9 \cap D_{11} \cap D_{13}| &= |D_{1287}| = 1 \\
|D_9 \cap D_{11} \cap D_{15}| &= |D_{495}| = 4 \\
|D_9 \cap D_{13} \cap D_{15}| &= |D_{585}| = 3 \\
|D_{11} \cap D_{13} \cap D_{15}| &= |D_{2145}| = 0
\end{aligned}
$$

The quadruple intersection ($\ominus$ — even, subtracted):

$$
|D_9 \cap D_{11} \cap D_{13} \cap D_{15}| = |D_{6435}| = 0
$$

So, by inclusion–exclusion with $n = 4$:

$$
\begin{aligned}
|D_9 \cup D_{11} \cup D_{13} \cup D_{15}|
&= 222 + 181 + 153 + 133 \\
&\quad - (20 + 17 + 44 + 13 + 12 + 10) \\
&\quad + (1 + 4 + 3 + 0) - 0 = \underline{581}
\end{aligned}
$$

## Towards the pigeonhole principle

Consider the situation: we have $m$ objects (a set $X$ with $\overline{\overline{X}} = m$) which must be placed into $n$ boxes (a set $Y$ with $\overline{\overline{Y}} = n$), where

$$
\overline{\overline{X}} = m > n = \overline{\overline{Y}}
$$

The teacher's margin doodle of the $m$ objects going into the $n$ boxes:

![[lec05_p07_objects-boxes.svg|300]]

Then into some box more than one object must go! — there is **no injective** ("różnowartościowa") function

$$
f : X \to Y
$$

> Margin note (looking ahead to Theorem 1): $X = f^{-1}\{1\} \cup f^{-1}\{2\} \cup \ldots \cup f^{-1}\{n\}$, so $\exists\, j \ |A_j| \geq \frac{|X|}{n} = \frac{m}{n} > 1$ — at least 2!

## Generalization: Dirichlet's pigeonhole (drawer) principle

### Theorem 1

> **Theorem 1.** If a finite set $S$ is partitioned into $k$ sets, then at least one of these sets has $|S|/k$ or more elements.
>
> Here $S = A_1 \cup A_2 \cup \ldots \cup A_k$ with $A_i \cap A_j = \emptyset$ for $i \neq j$, $\ i, j \in \{1, 2, \ldots, k\}$.

In other words, $\exists\, j \in \{1, 2, \ldots, k\}$:

$$
\boxed{\ |A_j| \geq \frac{|S|}{k}\ }
$$

**Proof.** Let $\hat{S}$ be the average number of elements in each set:

$$
\hat{S} = \frac{|A_1| + |A_2| + \ldots + |A_k|}{k}
$$

But $\{A_i\}_{i=1}^{k}$ form a partition of $S$, so $\sum_{i=1}^{k} |A_i| = |S|$ and

$$
\hat{S} = \frac{|S|}{k}.
$$

$\Rightarrow$ the most numerous set must have at least as many elements as the average: $\exists\, k_0 \in \{1, 2, \ldots, k\}$ with

$$
|A_{k_0}| \geq \hat{S} = \frac{|S|}{k},
$$

i.e.

$$
|A_{k_0}| \geq \frac{|S|}{k}. \qquad \square
$$

Often the partition is determined by a function:

### Theorem 2

> **Theorem 2.** Let a function $f : S \to T$ be given, where
> $$|S| < \infty \quad \text{and} \quad |T| < \infty, \qquad \text{and} \qquad |S| > r \cdot |T| \quad (*)$$
> Then at least one of the sets
> $$f^{-1}(t_1), \ f^{-1}(t_2), \ \ldots, \ f^{-1}(t_{|T|})$$
> has **more than $r$** elements.

**Proof.**

$$
S = \bigcup_{t \in T} f^{-1}(t)
$$

— a partition of the set $S$ into $K$ sets, where $K \leq |T|$. By Theorem 1, there exists $k_0$ such that $f^{-1}(t_{k_0})$ has at least $\frac{|S|}{K}$ elements:

$$
\bigl| f^{-1}(t_{k_0}) \bigr| \geq \frac{|S|}{K} \geq \frac{|S|}{|T|} > r \quad \text{by } (*). \qquad \blacksquare
$$

### Remark: the case $r = 1$

For $\boxed{r = 1}$ the principle states: if

- a) $f : S \to T$,
- b) $|S| > |T| \cdot \underset{r}{1}$,

then at least one of the sets $f^{-1}(t_1), f^{-1}(t_2), \ldots, f^{-1}(t_{|T|})$ has **more than 1** element!

The teacher's illustration of a function $f: S \to T$ and its preimages ($|S| = 7$, $|T| = 2$):

![[lec05_p10_function-preimages.svg]]

$$
\begin{cases}
f(s_i) = t_1, & 1 \leq i \leq 4 \\
f(s_i) = t_2, & 5 \leq i \leq 7
\end{cases}
\qquad\qquad
\underset{7}{|S|} > \underset{2}{|T|} \cdot \underset{r}{3}
$$

$$
\begin{aligned}
S_1 = f^{-1}(t_1) &= \{s_1, s_2, s_3, s_4\} \\
S_2 = f^{-1}(t_2) &= \{s_5, s_6, s_7\}
\end{aligned}
\qquad\qquad
|S| = 7, \quad |T| = 2
$$

$$
|S_1| = 4 > r = 3, \qquad |S_2| = 3 \not> r = 3
$$

Theorem 2 holds (here checked with $r = 3$).

## Example (pigeonhole principle)

- a) $A \subseteq \{1, 2, 3, \ldots, 50\}$,
- b) $|A| = 10$ (a ten-element subset).

We will show that $A$ has **2 different five-element subsets** such that the sums of all the elements of each of them are identical.

$$
S = \{\, B \subseteq A : |B| = 5 \,\}
$$

— the five-element subsets of $A$. To every $B \in S$ assign the sum of its elements:

$$
\forall\, B \in S \ \xrightarrow{\ f\ } \ \sum_{i=1}^{5} e_i, \qquad e_i \in B
$$

Obviously

$$
f(B) \geq 1 + 2 + 3 + 4 + 5 = 15
$$

$$
f(B) \leq 50 + 49 + 48 + 47 + 46 = 240
$$

So

$$
f : S \to T = \{15, 16, 17, \ldots, 239, 240\}, \qquad |T| = 226, \qquad |S| = \binom{10}{5} = 252
$$

($|S|$ counts the five-element subsets of a ten-element set.) Since

$$
|S| > \underset{r}{1} \cdot |T|,
$$

by Dirichlet's pigeonhole principle there exists $k_0 \in \{15, 16, \ldots, 240\} = T$ such that $f^{-1}(k_0)$ has **2 elements or more**: two subsets $B_1, B_2$ with

$$
\operatorname{sum}(B_1) = f(B_1) = f(B_2) = \operatorname{sum}(B_2). \qquad \square
$$

## The Josephus problem

Josephus Flavius and 41 Jews were surrounded in a cave by the Romans. Rather than surrender, they decided to commit suicide. They formed a circle and every 3rd person was killed. Flavius together with a friend wanted to escape death: they positioned themselves so as to remain last (assuming the counting starts from a chosen position).

**Simplification:**

- $n$ people,
- every **second** person dies,
- we start counting from number 1.

The teacher's sketch for $n = 10$ (the survivor, number 5, is circled):

![[lec05_p12_josephus-circle.svg|300]]

The elimination chain: $2, 4, 6, 8, 10, 3, 7, 1, 9$ $\ \Rightarrow\ $ number **5** remains!

$$
\boxed{J(n)} \ \longrightarrow \ \text{the number of the last one?}
$$

### First observations

| $n$ | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 |
|---|---|---|---|---|---|---|---|---|---|---|
| $J(n)$ | 1 | 1 | 3 | 1 | 3 | 5 | 7 | 1 | 3 | 5 |

(the columns $n = 1, 2, 4, 8$ are $2^0, 2^1, 2^2, 2^3$)

> Obviously $J(n)$ is **not** an increasing sequence!

- $J(n)$ seems to be **odd** (the first cycle eliminates the even numbers).

### a) The even case: $2n = n_0$ people in circle 1

Suppose $2n = n_0$ people stand in "circle 1". The first elimination cycle removes all the even numbers and leaves "circle 2" with $n$ elements — the counting then starts again from person 1:

![[lec05_p13_relabel-circles.svg]]

The two circles contain the same people standing in the same physical places (with $n$ people in between removed) — only the indices are different: giving circle 2 new indices $\tilde{\imath}$, the old numbers are recovered as

$$
1 = 2 \cdot \tilde{1} - 1, \qquad 3 = 2 \cdot \tilde{2} - 1, \qquad \ldots, \qquad (*)\ \boxed{2n - 1 = 2\tilde{n} - 1}
$$

The last person for circle 1 **is the same** as the last person for circle 2 ("the same" meaning: standing in the same physical place, and with an odd index). The number of this person differs only through the index change $(*)$:

$$
\boxed{J(2n) = 2\,J(n) - 1}, \qquad n \geq 1
$$

(in the big circle the survivor has the odd index $2J(n) - 1$, where $J(n)$ is his index in the small, re-indexed circle of $n$ people).

### b) The odd case: $n_0 = 2n + 1$

Similarly, for $n_0 = 2n + 1$:

$$
\boxed{J(2n+1) = 2\,J(n) + 1}, \qquad n \geq 1
$$

(here the first cycle eliminates $2, 4, \ldots, 2n$ and then, right after $2n$, also person 1; the $n$ survivors $3, 5, \ldots, 2n+1$ get new indices $\tilde{\imath}$ with old number $= 2\tilde{\imath} + 1$ — see the Remark at the end.)

### The recurrence scheme

We therefore have the following recurrence scheme:

$$
(\blacksquare) \quad
\begin{cases}
J(1) = 1 \\
J(2n) = 2\,J(n) - 1 \\
J(2n+1) = 2\,J(n) + 1
\end{cases}
\qquad (*), \quad n \geq 1
$$

> Margin remarks: this is **not** a recurrence of constant order, so the theory from the previous lectures does not apply directly; $(*)$ resembles a "divide and conquer" scheme, but we want a formula for all $n$, not only for $n = 2^k$; and $J(n)$ is not monotone!

One can build a longer "base of data" (**Table 1**):

| $n$ | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 | 15 | 16 |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| $J(n)$ | 1 | 1 | 3 | 1 | 3 | 5 | 7 | 1 | 3 | 5 | 7 | 9 | 11 | 13 | 15 | 1 |

(the groups are $2^0 \mid 2^1 \mid 2^2 \mid 2^3 \mid 2^4$)

One can see that at the beginning of each group $2^k$ we have $J(2^k) = 1$, and then the value **grows by 2**.

**Induction hypothesis** (read off from Table 1 — we do not harness any previous theory): writing $n = 2^m + p$,

$$
P(n): \quad (*)\ \boxed{\ J(\underbrace{2^m + p}_{n}) = 2p + 1\ }, \qquad m \geq 0, \quad 0 \leq p < 2^m
$$

### Proof by induction (strong version)

**I. Base step:** $n = 1$, i.e. $m = 0$ $\Rightarrow$ $2^0 + p = 1$ and $0 \leq p < 2^0 = 1$ $\Rightarrow$ $p = 0$:

$$
J(1) = 2 \cdot 0 + 1 = 1 \quad \text{OK}
$$

**II. Inductive step:** we assume the truth of $(*)$ for all $k < n_0$.

**(i)** Let $n_0 = 2n$ (**even**). Write $2n = 2^m + p$ (every number can be represented in this way; here $m > 0$), so $n = \frac{2^m + p}{2}$. Since

$$
2^m - 2n = -p
$$

is an even number, $p$ is **even**. From the recurrence relation $(\blacksquare)$:

$$
\begin{aligned}
J(\underbrace{2^m + p}_{2n}) &= 2\,J(n) - 1 \\
&= 2\,J\!\left( \frac{2^m + p}{2} \right) - 1 \\
&= 2\,J\!\left( 2^{m-1} + \frac{p}{2} \right) - 1 \qquad \left(\tfrac{p}{2} \text{ is an integer, since } p \text{ is even}\right)
\end{aligned}
$$

By the inductive assumption (applied to $k = n < 2n = n_0$, noting $0 \leq \frac{p}{2} < 2^{m-1}$):

$$
J(2^m + p) = 2\left[ 2 \cdot \frac{p}{2} + 1 \right] - 1 = 2p + 2 - 1 = \underline{2p + 1}
$$

For $n_0 = 2n$, $P(2n)$ is satisfied.

**(ii)** Let $2^m + p = 2n + 1 = n_0$ (**odd**). Then

$$
\underbrace{2^m - 2n - 1}_{\text{odd}} = -p \qquad (m > 0)
$$

so $p$ is **odd**; $p$ is of the form $\boxed{p = 2l + 1}$ $(\blacktriangle)$.

Side computation:

$$
\left\lfloor \frac{2n+1}{2} \right\rfloor = \left\lfloor \frac{2n}{2} + \frac{1}{2} \right\rfloor = \left\lfloor n + \frac{1}{2} \right\rfloor = \lfloor n \rfloor = n,
$$

hence the odd rule of $(\blacksquare)$ reads $J(n_0) = 2\,J(\lfloor n_0/2 \rfloor) + 1$. Using the recurrence relation $(\blacksquare)$:

$$
\begin{aligned}
J(2^m + p) &= 2\,J\!\left( \left\lfloor \frac{2^m + p}{2} \right\rfloor \right) + 1 \\
&= 2\,J\!\left( \left\lfloor 2^{m-1} + \frac{p}{2} \right\rfloor \right) + 1 \\
&= 2\,J\!\left( 2^{m-1} + \left\lfloor \frac{p}{2} \right\rfloor \right) + 1
\qquad \left( \lfloor k + x \rfloor = k + \lfloor x \rfloor \text{ for } k \in \mathbb{Z} \right)
\end{aligned}
$$

Since

$$
\underbrace{2^{m-1} + \left\lfloor \frac{p}{2} \right\rfloor}_{k} < \underbrace{2^m + p}_{n_0},
$$

we may use the inductive assumption:

$$
J(2^m + p) = 2\left( 2\left\lfloor \frac{p}{2} \right\rfloor + 1 \right) + 1
$$

But by $(\blacktriangle)$, $p = 2l + 1 \Rightarrow \frac{p}{2} = l + \frac{1}{2}$, and $\left\lfloor \frac{p}{2} \right\rfloor = \left\lfloor l + \frac{1}{2} \right\rfloor = l$, so

$$
J(2^m + p) = 2(2l + 1) + 1 = 4l + 2 + 1 = 2(2l + 1) + 1 = \underline{2p + 1}
$$

### Conclusion

Hence, by mathematical induction (the **strong** version),

$$
\boxed{\ J(\underbrace{2^m + p}_{n}) = 2p + 1\ }, \qquad m \geq 0, \quad 0 \leq p < 2^m
$$

is true for every $n \geq 1$.

The theorem is easy to generalize to "cuts every three".

If we assume $n = 42$ and cut every second, Flavius should position himself at:

$$
J(42 = 2^5 + 10) = 2 \cdot 10 + 1 = \underline{\underline{21}}
$$

**Check.** The teacher verifies this by crossing out the numbers $1, \ldots, 42$ cycle by cycle; after the 1st cycle the evens are gone, after the 2nd cycle $3, 7, 11, \ldots, 39$ are gone:

![[lec05_p17_elimination-cycles-1-2.svg]]

The remaining cycles eliminate $1, 9, 17, 25, 33, 41$, then $13, 29$, and finally $5$ and $37$:

![[lec05_p18_elimination-cycles-3-5.svg]]

Indeed **21** wins!

### $(*)$ Remark — where the two recurrences come from

Small sketches of the first elimination cycle in both cases (crossed = eliminated, $\tilde{\imath}$ = new index of a survivor):

![[lec05_p18_recurrence-circles.svg]]

**a)** $n_0 = 2n$: the first cycle crosses out the even numbers; the survivors $1, 3, 5, \ldots$ get new indices with

$$
1 = 2 \cdot \tilde{1} - 1, \qquad 3 = 2 \cdot \tilde{2} - 1, \qquad 5 = 2 \cdot \tilde{3} - 1, \quad \ldots
$$

$$
\boxed{J(2n) = 2\,J(n) - 1}
$$

**b)** $n_0 = 2n + 1$: the first cycle crosses out $2, 4, 6, \ldots$ and then also person 1; the survivors $3, 5, 7, \ldots$ satisfy (note $\left\lfloor \frac{2n+1}{2} \right\rfloor = n$)

$$
3 = 2 \cdot \tilde{1} + 1, \qquad 5 = 2 \cdot \tilde{2} + 1, \qquad 7 = 2 \cdot \tilde{3} + 1, \quad \ldots
$$

$$
\boxed{J(2n+1) = 2\,J(n) + 1}
$$
