# Discrete Mathematics — Lecture 15: Review for the Exams

> Translated from the handwritten lecture notes `pdf/lect15-Powtorzenie-MatDysk.pdf`.
> (The handwritten header reads "Lecture 17" — the teacher's own numbering.)

## Problem 1 — asymptotics (Big-O)

**Problem:** Let

$$
\begin{aligned}
f(n) &= 5n^2 + 6n - 1 \\
g(n) &= 2n\lg_2 n + n^4 \\
h(n) &= 5\sqrt[3]{n} + n - 2^n
\end{aligned}
$$

Find the asymptotics:

$$
\begin{aligned}
\text{(i)} \quad & f(n) \pm g(n) \pm h(n) = O(?) \\
\text{(ii)} \quad & f(n) \cdot g(n) = O(?) \\
\text{(iii)} \quad & f(n) \cdot h(n) = O(?)
\end{aligned}
$$

We must remember the **hierarchy**:

$$
(*) \quad 1, \ \ldots, \ \sqrt[4]{n}, \ \sqrt[3]{n}, \ \sqrt{n}, \ n, \ n\lg_2 n, \ n\sqrt{n}, \ n^2, \ n^3, \ n^4, \ \ldots, \ 2^n, \ n!, \ n^n
$$

— every sequence is of order $O$ of all the sequences to its right.

**Solution.** We use

$$
O(f_1(n)) + O(f_2(n)) = O\!\left(\max\{\,|f_1(n)|, |f_2(n)|\,\}\right)
$$

$$
\begin{aligned}
f(n) &= O(5n^2 + 6n - 1) \\
&= O(5n^2) + O(6n) + O(1) \\
&= O(n^2) + O(n) + O(1) \\
&= \underline{O(n^2)} \quad \text{by } (*)
\end{aligned}
$$

$$
g(n) = \underline{O(n^4)} \quad \text{by } (*), \qquad h(n) = \underline{O(2^n)} \quad \text{by } (*)
$$

**(i)** is of order $\underline{O(2^n)}$.

For products:

$$
O(f_1(n)) \cdot O(f_2(n)) = O(f_1(n)\, f_2(n))
$$

**(ii)** is of order $\underline{O(n^6)}$.

**(iii)** is of order $\underline{O(n^2 2^n)}$.

Review the definition of $f(n) = O(g(n))$ — the teacher's margin doodle: from some $n_0$ on, $|f(n)|$ stays below $c \cdot g(n)$:

![[lec15_p02_big-o-definition.svg|320]]

## Problem 2 — counting functions and subsets

**Problem:**

$$
X = \{1, 3, 6\}, \qquad Y = \{2, 8, 9, 10\}
$$

- (i) how many **injective** functions $f : X \to Y$ are there?
- (ii) how many functions $f : X \to Y$ are there altogether?
- (iii) how many **3-element subsets** of $Y$ are there?
- (iv) how many subsets of $X$ are there altogether?
- (v) how many **permutations** of $X$ are there?
- (vi) how many **increasing** (non-decreasing) functions from $X$ into $Y$ are there?

**Solution.**

**(i)** (variations without repetition)

$$
V_{\overline{\overline{X}}}^{\overline{\overline{Y}}} = V_3^4 = \binom{4}{3}\, 3! = \frac{4!}{3!\,(4-3)!} \cdot 3! = 4! = 24
$$

One can also "catch" them on one's fingers:

$$
\begin{aligned}
f_1(1) &= 2 & f_2(1) &= 2 \\
f_1(3) &= 8 & f_2(3) &= 9 & &\text{etc.}\ldots\\
f_1(6) &= 9 & f_2(6) &= 10
\end{aligned}
$$

**(ii)** all functions:

$$
W_{\overline{\overline{Y}}}^{\overline{\overline{X}}} = \overline{\overline{Y}}^{\,\overline{\overline{X}}} = 4^3 = \underline{64}
$$

24 of them are injective. Writing them all out is tedious.

**(iii)** $\{2,8,9\}$, $\{8,9,10\}$, $\{2,9,10\}$, $\{2,8,10\}$ — counting by hand one sees there are $4$.

$$
\binom{\overline{\overline{Y}}}{3} = \binom{4}{3} = \frac{4!}{3!\,1!} = \underline{\underline{4}}
$$

Indeed.

**(iv)**

$$
\{\, \emptyset, \{1\}, \{3\}, \{6\}, \{1,6\}, \{1,3\}, \{3,6\}, \{1,3,6\} \,\}
$$

— there are $8$ of them, found by hand.

$$
2^{\overline{\overline{X}}} = 2^3 = \underline{\underline{8}}
$$

indeed.

**(v)** permutations:

$$
\overline{\overline{X}}\,! = 3! = \underline{\underline{6}}, \qquad V_3^3 = \binom{3}{3}\, 3! = 6
$$

$$
(1\,3\,6),\ (1\,6\,3),\ (3\,1\,6),\ (3\,6\,1),\ (6\,1\,3),\ (6\,3\,1) \quad \text{— there are } 6.
$$

**(vi)** the number of increasing functions $f : X \to Y$ is

$$
\binom{\overline{\overline{Y}}}{\overline{\overline{X}}} = \binom{4}{3} = 4
$$

They can be written out explicitly:

$$
\begin{aligned}
f_1(1) &= 2 & f_2(1) &= 2 & f_3(1) &= 8 & f_4(1) &= 2 \\
f_1(3) &= 8 & f_2(3) &= 9 & f_3(3) &= 9 & f_4(3) &= 8 \\
f_1(6) &= 9 & f_2(6) &= 10 & f_3(6) &= 10 & f_4(6) &= 10
\end{aligned}
$$

**Non-decreasing** functions (i.e. $x \leq y \Rightarrow f(x) \leq f(y)$) have to be searched for by hand. E.g. $f(1) = f(3) = f(6)$ is non-decreasing. $\square$

## Problem 3 — tossing a coin three times

**Problem:** We toss a fair (symmetric) coin 3 times. Compute the probability of obtaining:

- (i) exactly 2 heads;
- (ii) at least 2 tails;
- (iii) at least 1 head.

**Solution.** (Writing $H$ for heads and $T$ for tails; the original uses $O$ — *orzeł* — and $R$ — *reszka*.) The set of elementary events:

$$
\Omega = \{ (H,H,H), (H,H,T), (H,T,H), (H,T,T), (T,H,H), (T,H,T), (T,T,H), (T,T,T) \}
$$

$$
\overline{\overline{\Omega}} = 8
$$

One can also treat an outcome as a **function** from the 3 positions into $\{H, T\}$; e.g. $(H,H,T)$ is

$$
f(\mathrm{I}) = H, \quad f(\mathrm{II}) = H, \quad f(\mathrm{III}) = T
$$

The number of all outcomes equals the number of all functions from a 3-element set into a 2-element set, i.e. $2^3 = \underline{8}$.

(The teacher solves (ii) and (iii) first, then (i).)

**(ii)**

$$
A = \{ (T,T,H), (T,H,T), (H,T,T), (T,T,T) \}
$$

Since the coin is symmetric, all elementary events are equally probable, and hence

$$
P(A) = \frac{\overline{\overline{A}}}{\overline{\overline{\Omega}}} = \frac{4}{8} = \frac{1}{2}
$$

**(iii)** Let $A'$ be the complementary event — no head at all: $A' = \{(T,T,T)\}$.

$$
P(A) = 1 - P(A') = 1 - \frac{1}{8} = \underline{\underline{\frac{7}{8}}}
\qquad \left( \frac{\overline{\overline{A'}}}{\overline{\overline{\Omega}}} = \frac{1}{8} \right)
$$

One can also determine $A$ directly and count $\overline{\overline{A}} = 7$, whence $P(A) = \frac{7}{8}$.

**(i)**

$$
A = \{ (H,H,T), (H,T,H), (T,H,H) \}, \qquad
\overline{\overline{A}} = 3 \ \Rightarrow \ P(A) = \frac{\overline{\overline{A}}}{\overline{\overline{\Omega}}} = \frac{3}{8}.
$$

## Problem 4 — proofs by induction

**Problem:** Prove by induction:

- a) $x^{2n} - y^{2n}$ is divisible by $x - y$, for $n \geq 1$;
- b) $1^2 - 2^2 + 3^2 - \ldots + (-1)^{n+1} n^2 = (-1)^{n+1}(1 + 2 + \ldots + n) = (-1)^{n+1}\dfrac{(n+1)n}{2}$, for $n \geq 1$;
- c) $2^{2^n} - 6$ is divisible by $10$, for $n \geq 2$;
- d) $\dfrac{1}{\sqrt{1}} + \dfrac{1}{\sqrt{2}} + \ldots + \dfrac{1}{\sqrt{n}} \geq \sqrt{n}$, for $n \geq 1$.

### a)

**1)** For $n = 1$:

$$
x^2 - y^2 = (x - y)(x + y)
$$

so $x - y$ divides $x^2 - y^2$.

**2)** Assume $x^{2n} - y^{2n}$ is divisible by $x - y$. Question: does $x - y$ divide $x^{2(n+1)} - y^{2(n+1)}$?

$$
\begin{aligned}
x^{2(n+1)} - y^{2(n+1)}
&= x^{2n} x^2 - y^{2n} y^2 \\
&= x^{2n} x^2 - x^2 y^{2n} + x^2 y^{2n} - y^{2n} y^2 \\
&= \underbrace{x^2 \left( x^{2n} - y^{2n} \right)}_{\substack{\text{divisible by } x-y \\ \text{by the induction} \\ \text{hypothesis}}}
 + \underbrace{y^{2n} \left( x^2 - y^2 \right)}_{\substack{= y^{2n}(x-y)(x+y) \\ \text{divisible by } x-y}}
\end{aligned}
$$

Both summands are divisible, so the whole sum is divisible as well. By induction, a) is true for $n \geq 1$.

### b)

Base case: $(-1)^2 1^2 = (-1)^2 \cdot 1$ — true.

Assume the truth for $n$:

$$
(*) \quad 1^2 - 2^2 + 3^2 + \ldots + (-1)^{n+1} n^2 = (-1)^{n+1}(1 + 2 + \ldots + n)
$$

We add $(-1)^{n+1+1}(n+1)^2$ to both sides of $(*)$:

$$
(**) \quad 1^2 - 2^2 + 3^2 + \ldots + (-1)^{n+1} n^2 + (-1)^{n+1+1}(n+1)^2
= (-1)^{n+1}(1 + 2 + \ldots + n) + (-1)^{n+1+1}(n+1)^2
$$

The left side of $(**)$ is what we need. We transform the right side:

$$
\begin{aligned}
& (-1)^{n+1}(1 + 2 + \ldots + n) + (-1)^{n+2}(n+1)^2 \\
&= (-1)^{n+2}(-1)(1 + 2 + \ldots + n) + (-1)^{n+2}(n+1)^2 \\
&= (-1)^{n+2}\left[ (n+1)^2 - (1 + 2 + \ldots + n) \right] \\
&= (-1)^{n+2}\left[ (n+1)^2 - \frac{n(n+1)}{2} \right] \quad \text{(from the lecture)} \\
&= (-1)^{n+2}\left( \frac{2n^2 + 4n + 2 - n^2 - n}{2} \right) \\
&= (-1)^{n+2}\left( \frac{n^2 + 3n + 2}{2} \right) \\
&= (-1)^{n+2}\left( \frac{(n+1)(n+2)}{2} \right) \\
&= (-1)^{n+2}\left( 1 + 2 + \ldots + n + (n+1) \right)
\end{aligned}
$$

OK. By induction, b) is true for $n \geq 1$.

### c)

**1)** For $n = 2$:

$$
2^{2^2} - 6 = 2^4 - 6 = 16 - 6 = 10
$$

— divisible by $10$.

**2)** Assume $2^{2^n} - 6$ is divisible by $10$. Question: is $2^{2^{n+1}} - 6$ also divisible by $10$?

$$
\begin{aligned}
2^{2^{n+1}} - 6 &= \left(2^{2^n}\right)^2 - 6
= \left(2^{2^n}\right)^2 - 6 \cdot 2^{2^n} + 6 \cdot 2^{2^n} - 6 \\
&= 2^{2^n}\left(2^{2^n} - 6\right) + 6 \cdot 2^{2^n} - 6 \\
&= 2^{2^n}\left(2^{2^n} - 6\right) + 6 \cdot 2^{2^n} - 36 + 30 \\
&= \underbrace{2^{2^n}\left(2^{2^n} - 6\right)}_{\text{divisible by } 10}
 + \underbrace{6\left(2^{2^n} - 6\right)}_{\text{divisible by } 10}
 + \underbrace{30}_{\text{divisible by } 10}
\end{aligned}
$$

(here $2^{2^n \cdot 2} = \left(2^{2^n}\right)^2$). Each summand is divisible by $10$, so the whole sum is divisible by $10$. By induction, c) is true for $n \geq 2$.

### d)

Base case: $\dfrac{1}{\sqrt{1}} \geq \sqrt{1}$ — true.

Assume

$$
(*) \quad \sum_{i=1}^{n} \frac{1}{\sqrt{i}} \geq \sqrt{n}
\qquad \stackrel{?}{\Longrightarrow} \qquad
\sum_{i=1}^{n+1} \frac{1}{\sqrt{i}} \geq \sqrt{n+1}
$$

From $(*)$, adding $\dfrac{1}{\sqrt{n+1}}$ to both sides:

$$
\underbrace{\sum_{i=1}^{n} \frac{1}{\sqrt{i}} + \frac{1}{\sqrt{n+1}}}_{\displaystyle \sum_{i=1}^{n+1} \frac{1}{\sqrt{i}} \ \text{— what we want}}
\ \geq\ \sqrt{n} + \frac{1}{\sqrt{n+1}}
$$

So it suffices to show that

$$
(\star) \quad \sqrt{n} + \frac{1}{\sqrt{n+1}} \geq \sqrt{n+1}
$$

Then by transitivity we get

$$
\sum_{i=1}^{n+1} \frac{1}{\sqrt{i}} \geq \sqrt{n+1}
$$

which is what needed to be shown. To prove $(\star)$, multiply both sides by $\sqrt{n+1} \geq 0$:

$$
\sqrt{n}\sqrt{n+1} + 1 \geq n + 1
$$

$$
(**) \quad \sqrt{n^2 + n} + 1 \geq n + 1
\qquad \Longleftrightarrow \qquad \sqrt{n^2 + n} \geq n
$$

Both sides of $(**)$ are non-negative, so we may square and preserve the sign of the inequality:

$$
n^2 + n \geq n^2
$$

$$
\underline{n \geq 0} \quad \text{OK} \quad \Rightarrow (\star) \text{ is true.}
$$

By mathematical induction, our inequality is true. $\square$

## Problem 5 — a divide-and-conquer recurrence

**Problem:** Let

$$
T(1) = 5, \qquad T(n) = 8\, T\!\left(\left\lfloor \tfrac{n}{2} \right\rfloor\right)
$$

define a recurrence. Then $T(n) = O(?)$.

**Solution.** Consider first $n = 2^k$:

$$
\begin{aligned}
T(n) &= 8 \cdot T\!\left(\left\lfloor \tfrac{n = 2^k}{2} \right\rfloor\right) \\
&= 8 \cdot T\!\left(2^{k-1}\right) \\
&= 8 \cdot 8\, T\!\left(\left\lfloor \tfrac{2^{k-1}}{2} \right\rfloor\right) \\
&= 8^2\, T\!\left(2^{k-2}\right) \\
&= \ldots \\
&= 8^k\, T\!\left(2^{k-k}\right) = 8^k\, T(1) = 5 \cdot 8^k
\end{aligned}
$$

One sees, then, that for **any** $n$ (not necessarily of the form $n = 2^k$):

$$
(\blacktriangle) \quad \boxed{\ T(n) = 5 \cdot 8^k\ }
\qquad \text{where } k \text{ is the maximal number such that } \boxed{2^k \leq n} \ (*)
$$

— this follows from the division $\left\lfloor \frac{n}{2} \right\rfloor$. The inductive proof of $(\blacktriangle)$ comes afterwards. Thus

$$
\begin{aligned}
|T(n)| &= 5 \cdot 8^k = 5 \cdot (2 \cdot 2 \cdot 2)^k \\
&= 5 \cdot 2^k \cdot 2^k \cdot 2^k \\
&\underset{(*)}{\leq} 5 \cdot n \cdot n \cdot n = 5n^3 = \underline{O(n^3)}
\end{aligned}
$$

Therefore

$$
\boxed{\ T(n) = O(n^3)\ }
$$

> $(\blacktriangle)$ may be omitted, but here we will do the inductive proof.

### Inductive proof of $(\blacktriangle)$

**Base case** — for $n = 1$: $T(1) = 5$ (from the definition). Also, the maximal $k$ such that $2^k \leq 1$ is $k = 0$, so

$$
T(1) = 5 \cdot 8^0 = \underline{\underline{5}}
$$

— the same thing.

**Inductive step:** assume the truth of $(\blacktriangle)$ for all $n < n_0$. Question: is $(\blacktriangle)$ true for $n_0$?

**1. $n_0$ even** (case 1):

$$
T(n_0) = 8\, T\!\left(\left\lfloor n_0/2 \right\rfloor\right) = 8 \cdot T\!\left(\tfrac{n_0}{2}\right)
$$

(here $\left\lfloor \frac{n_0}{2} \right\rfloor = \frac{n_0}{2}$ because $n_0$ is even). By the induction hypothesis, since $\frac{n_0}{2} < n_0$:

$$
= 8 \cdot 8^{k_0} \cdot 5 = \underline{5 \cdot 8^{k_0 + 1}}
$$

where $k_0$ is maximal such that $2^{k_0} \leq \frac{n_0}{2}$. We know that

$$
2^{k_0} \leq \frac{n_0}{2} \quad \Longrightarrow \quad 2^{k_0 + 1} \leq n_0
$$

and of course $k_0 + 1$ **is maximal**: if it were $\boxed{k_0 + 1 < l}$ and

$$
(**) \quad 2^l \leq n_0
$$

then $2^{l-1} \leq \frac{n_0}{2}$, and since $k_0$ was maximal such that $2^{k_0} \leq \frac{n_0}{2}$, we get $l - 1 \leq k_0$, i.e. $\boxed{l \leq k_0 + 1}$ — a contradiction with $(**)$.

We have shown, then, that when $n_0$ is even,

$$
T(n_0) = 5 \cdot 8^{\tilde{k}}, \qquad \tilde{k} = k_0 + 1
$$

where $\tilde{k}$ is maximal with $2^{\tilde{k}} \leq n_0$ — which was to be shown.

**2. Case $n_0$ odd:** then $n_0 - 1$ is even, and hence

$$
(*) \quad \left\lfloor \frac{n_0}{2} \right\rfloor = \left\lfloor \frac{n_0 - 1}{2} \right\rfloor = \frac{n_0 - 1}{2}
$$

So

$$
T(n_0) = 8 \cdot T\!\left(\left\lfloor \tfrac{n_0}{2} \right\rfloor\right)
\underset{(*)}{=} 8 \cdot T\!\left(\tfrac{n_0 - 1}{2}\right)
$$

But since $\frac{n_0 - 1}{2} < n_0$, the induction hypothesis yields

$$
= 8 \cdot 8^{k_0} \cdot 5 = 8^{k_0 + 1} \cdot 5 = \underline{\underline{5 \cdot 8^{k_0 + 1}}}
$$

where $k_0$ is maximal such that

$$
(\blacksquare) \quad 2^{k_0} \leq \frac{n_0 - 1}{2}
\quad \Longrightarrow \quad 2^{k_0 + 1} \leq n_0 - 1 \leq n_0
\quad \Longrightarrow \quad \boxed{2^{k_0 + 1} \leq n_0}
$$

It suffices to show that $k_0 + 1$ is **maximal** such that $2^{k_0 + 1} \leq n_0$. Suppose there exists $l_0 > k_0 + 1$ such that

$$
\underbrace{2^{l_0}}_{\text{even}} \leq \underbrace{n_0}_{\text{odd}}
\quad \Longrightarrow \quad 2^{l_0} \leq n_0 - 1
\quad \Longrightarrow \quad 2^{l_0 - 1} \leq \frac{n_0 - 1}{2}
$$

$$
\overset{(\blacksquare)}{\Longrightarrow} \quad l_0 - 1 \leq k_0 \quad \Longrightarrow \quad l_0 \leq k_0 + 1
$$

— a contradiction with $l_0 > k_0 + 1$. We have shown, then, that for $n_0$ odd

$$
T(n_0) = 5 \cdot 8^k, \qquad \text{where } k \text{ is maximal with } \boxed{2^k \leq n_0}
$$

Q.E.D.

## Problem 6 — even products and even sums

**Problem:** Let $A = \{1, 2, 3, 4, 5\}$.

- a) in how many ways can one choose (**without replacement**) 2 numbers so that their **product** is even?
- b) in how many ways can one choose (**without replacement**) 2 numbers so that their **sum** is even?

**Remark:** for a) we assume that the **order** of choosing the 2 numbers is **irrelevant**. For b) we assume that the order of drawing the 2 numbers **is relevant**.

**a)** The product of 2 numbers is even when one is even and the other is even or odd. Write

$$
N = \{1, 3, 5\} \ \text{(odd)}, \qquad P = \{2, 4\} \ \text{(even)}
$$

Since the order is irrelevant (two-element subsets), we have

$$
P \cdot P + P \cdot N \qquad \longrightarrow \qquad
\binom{2}{2} + \binom{2}{1}\binom{3}{1} = 1 + 2 \cdot 3 = \underline{\underline{7}} \ \text{ ways}
$$

($\{P, P\}$ — 1 choice, plus 2 possibilities $\cdot$ 3 possibilities $= 7$). Here they are:

$$
\begin{aligned}
&1.\ \{2,4\} && \text{(both even)} \\
&2.\ \{2,1\}, \quad 3.\ \{2,3\}, \quad 4.\ \{2,5\}, \quad 5.\ \{4,1\}, \quad 6.\ \{4,3\}, \quad 7.\ \{4,5\} && \text{(even and odd)}
\end{aligned}
$$

**b)** The sum of 2 numbers is even iff **both are even or both are odd**. Here the order of drawing the numbers is relevant, so we consider **ordered pairs**:

$$
(p_1, p_2) + (np_1, np_2)
\qquad \longrightarrow \qquad
V_2^2 + V_2^3 = \binom{2}{2} \cdot 2! + \binom{3}{2} \cdot 2! = 2 + 6 = \underline{\underline{8}}
$$

> (Marginal note: "without replacement" $\Rightarrow$ **variations without repetition** $V_k^n$.)

Here they are:

$$
\begin{aligned}
&1.\ (2,4), \quad 2.\ (4,2) && \text{— ordered pairs of even numbers} \\
&3.\ (1,3), \quad 4.\ (1,5), \quad 5.\ (3,1), \quad 6.\ (5,1), \quad 7.\ (3,5), \quad 8.\ (5,3) && \text{— ordered pairs of 2 odd numbers}
\end{aligned}
$$

One can also count it by the counting (multiplication) method:

$$
(p^1, p^2) + (np^1, np^2) \qquad \longrightarrow \qquad 2 \cdot 1 + 3 \cdot 2 = \underline{8}
$$

## Problem 7 — properties of a graph

**Problem:** Let $G = (V, E)$ be the undirected graph defined as follows:

$$
V = \{a, b, c, d, e, f\}
$$

$$
E = \{\, \{a,b\}, \{b,c\}, \{c,f\}, \{a,d\}, \{d,e\}, \{a,e\} \,\}
$$

The drawing of $G$ (Fig. 1):

![[lec15_p22_graph-G.svg]]

- a) is the graph $G$ connected?
- b) does $G$ have an Euler cycle?
- c) is $G$ a tree?
- d) is $G$ acyclic?
- e) is $G$ a Hamiltonian graph?
- f) is $G$ a complete graph, or a regular one?
- g) compute $\mathrm{Deg}(d)$ and $D_2(G)$;
- h) does $G$ have a spanning tree? (if so, find at least one);
- i) does the graph $G$ have loops or multiple edges?

**ad a)** $G$ is **connected** if every pair of distinct vertices is joined by a path in this graph.

**YES** (see Fig. 1).

**ad b)**

> **Thm 1.** A finite connected graph in which every vertex has even degree has an Euler cycle.
>
> **Thm 2.** A graph which has an Euler cycle must have all vertices of even degree.
>
> **Thm 1 + Thm 2.** A connected graph has an Euler cycle $\iff$ every vertex has even degree.

An **Euler path** is a simple path (each edge once) which contains all the edges of $G$; an **Euler cycle** is a closed Euler path.

In our case $G$ is connected, but the vertices $a$ and $f$ do **not** have even degree:

$$
\mathrm{Deg}(a) = 3, \qquad \mathrm{Deg}(f) = 1
$$

$\Rightarrow$ **$G$ has no Euler cycle!**

**c)** $G$ is **not** a tree (trees are the acyclic **and** connected graphs). $G$ has the cycle $\{a, e, d\}$ — a closed simple path (single edges) with distinct vertices.

**d)** $G$ is **not** acyclic; the cycle is

$$
\{a, e, d\}
$$

— a closed **simple** path (each edge once) with distinct vertices.

**e)** A graph having a Hamilton cycle is a **Hamiltonian graph**. A **Hamilton cycle** is a closed Hamilton path — a path which passes through every vertex of $G$ exactly once. One can see "with the naked eye" that there is no such cycle here. The sufficient conditions (Theorems 1, 2, 3 of the lecture on Hamiltonian graphs) are not satisfied either.

**f)** A **regular** graph:

$$
\forall v \in V(G) \quad \deg(v) = k \quad \text{(the same degree)}
$$

In our case $\deg(f) = 1$, $\deg(e) = 2$, so **$G$ is not regular**.

A **complete** graph: no loops, no multiple edges, and $\deg(v) = n - 1$ for every $v \in V(G)$, where $|V(G)| = n$ (a graph with $n$ vertices). Our graph is **not complete**, because it is not regular.

**g)**

$$
\deg(d) = 2 \quad \text{(the number of edges at the vertex } d\text{)}, \qquad D_2(G) = \{e, d, b, c\}
$$

— the vertices of index (degree) $2$.

**(i)** $G$ has **no loops and no multiple edges**, i.e. there is none of:

![[lec15_p27_loop-multiedge.svg|380]]

**h)** A **spanning tree**:

- a) a subgraph of $G$;
- b) contains all the vertices $V(G)$;
- c) has the minimal number of edges;
- d) is a tree.

**Every finite connected graph has a spanning tree.** The teacher's three examples (the only cycle of $G$ is $\{a, d, e\}$, and deleting one of its edges gives a tree):

![[lec15_p27_spanning-trees.svg]]

There are many such trees!

## Problem 8 — isomorphism invariants

**Problem:** Let $V = \{1, 4, 3, 7\}$ plus certain sets of edges $E_1$ and $E_2$. The undirected graphs

$$
G_1 = (V, E_1) \ \simeq \ G_2 = (V, E_2)
$$

(the **same vertices**) are **isomorphic**. Which of the following statements are then true?

- a) the number of loops in $G_1$ is different from the number of loops in $G_2$;
- b) the number of edges in $E_1$ and $E_2$ is the same;
- c) the sequences $(D_0(G_1), D_1(G_1), D_2(G_1), \ldots)$ and $(D_0(G_2), D_1(G_2), D_2(G_2), \ldots)$ are identical;
- d) the sets $E_1 = E_2$ (here $V_1 = V_2$);
- e) if $\{1,3\} \in E_1 \Rightarrow \{1,3\} \in E_2$;
- f) if $\{1,3\} \notin E_1 \Rightarrow \{1,3\} \notin E_2$;
- g) connectedness of $G_1 \Longrightarrow$ connectedness of $G_2$;
- h) $\deg(v) = \deg(v)$ for every vertex of the graph, where $V(G_1) = V(G_2)$ — i.e. every vertex has the same degree in $G_1$ as in $G_2$ ($\alpha$ — the isomorphism).

**Answers:**

**a)** The number of loops is an **isomorphism invariant** of graphs, so it must be the **same**. — **NO.**

**b)** The number of edges and vertices is also an isomorphism invariant of graphs (here $V_1 = V_2 = V$, so the number of vertices is automatically the same). The number of edges in $E_1$ and $E_2$ must be identical. — **YES.**

**c)** The sequences $(D_0(G_i), D_1(G_i), D_2(G_i), \ldots)$ are an isomorphism invariant — they must be identical. — **YES.**

**d)** The sets $E_1$ and $E_2$ **need not** be identical (even if the vertices are the same).

**Example:**

$$
G_1 = \{\, \underbrace{V(G_1)}_{\{a,b,c\}},\ \underbrace{E(G_1)}_{\{\{a,b\},\{b,c\}\}} \,\}
\qquad
G_2 = \{\, \underbrace{V(G_2)}_{\{a,b,c\}},\ \underbrace{E(G_2)}_{\{\{a,c\},\{b,c\}\}} \,\}
$$

> (In our problem $G = \{V, E\}$ with $V = \{1,4,3,7\}$ — this is an example for a different graph, but it is the same thing!)

![[lec15_p30_isomorphism-example.svg|420]]

$$
\alpha(a) = b, \quad \alpha(b) = c, \quad \alpha(c) = a
$$

is an isomorphism:

$$
\begin{aligned}
\alpha(\{a,b\}) &= \{\alpha(a), \alpha(b)\} = \{b, c\} \\
\alpha(\{b,c\}) &= \{\alpha(b), \alpha(c)\} = \{c, a\}
\end{aligned}
$$

So here $V(G_1) = V(G_2)$ but $E(G_1) \neq E(G_2)$. — **NO.**

**e)** Is it true that $\{1,3\} \in E_1 \Rightarrow \{1,3\} \in E_2$? The previous example showed

$$
\{a,b\} \in E_1 \quad \not\Rightarrow \quad \{a,b\} \in E_2
$$

— **NO.**

**f)** Similarly, the previous example showed that

$$
\{a,c\} \notin E_1 \quad \not\Rightarrow \quad \{a,c\} \notin E_2
$$

— **NO.**

**h)** Only $\deg(v) = \deg(\alpha(v))$ for all $v \in G$ holds. **No** — the previous example:

$$
\underset{\text{in } G_1}{\deg(b) = 2}, \qquad \underset{\text{in } G_2}{\deg(b) = 1}
$$

— **NO.**

**g)** **YES** — connectedness of graphs is an isomorphism invariant.

> (In the handwritten answers the teacher's labels drift: the last two answers are marked h) and (i); they answer questions h) and g) above, respectively.)

## Final task — what to master for the exam

**Master:**

- the theorems, definitions, lemmas;
- the examples;
- redo the exercises.

(Added in pen:)

- $a_n = a\, a_{n-1} + b\, a_{n-2}$ (characteristic functions);
- **sums!!!**
- $a_n = 2a_{n-1} + f(n)$, "**divide and conquer**".
