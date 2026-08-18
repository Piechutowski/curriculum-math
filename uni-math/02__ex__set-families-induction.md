# Discrete Mathematics — Exercises 2: Set Families and Induction

> Translated from the handwritten exercise sheet `pdf/Ćwiczenia2.pdf`.

## Problem 1

Find, for $A_i = \left[0, \frac{2}{i}\right)$ and $B = \left(-1, \frac{1}{2}\right)$:

$$
\text{(i)} \quad \bigcup_{i=1}^{5} A_i = A_1 \cup A_2 \cup \ldots \cup A_5
\qquad
\text{(ii)} \quad \bigcap_{i=1}^{5} A_i = A_1 \cap A_2 \cap \ldots \cap A_5
$$

$$
\text{(iii)} \quad \bigcup_{i=1}^{\infty} A_i
\qquad
\text{(iv)} \quad \bigcap_{i=1}^{\infty} A_i
\qquad
\text{(v)} \quad \left( \bigcup_{i=1}^{\infty} A_i \right) \cap B
\qquad
\text{(vi)} \quad \left( \bigcap_{i=1}^{\infty} A_i \right) \cup B.
$$

> **Note:** on the sheet the last item is labelled "(IV)" a second time — it is clearly the sixth part, renumbered here as (vi).

## Problem 2

Using **mathematical induction** prove that:

**a)**

$$
\sum_{i=1}^{n} (2i - 1) = n^2
$$

**b)**

$$
2^n > (n+1)^2 \quad \text{for } n \geq \ ?
$$

(find the threshold value of $n$ yourself).

**c)** For the sequence defined by

$$
a_0 = 1, \quad a_1 = 2, \quad a_2 = 3, \qquad a_n = a_{n-1} + a_{n-2} + a_{n-3},
$$

show that

$$
a_n \leq 3^n \quad \forall n \in \mathbb{N}.
$$

> **Note:** the sheet writes $a_3 = 3$ and repeats the term $a_{n-2}$ in the recurrence; with three initial conditions the evident intent is $a_2 = 3$ and $a_n = a_{n-1} + a_{n-2} + a_{n-3}$, as written above.

**d)**

$$
5 \mid 2 \cdot 4^n + 3 \cdot 9^n \quad \text{for } n \geq 0.
$$

## Problem 3

Show that the sequence given recursively in **implicit form**:

$$
p_0 = 3, \quad p_1 = 7, \quad \text{and} \quad p_n = 3p_{n-1} - 2p_{n-2}
$$

has, for $n \geq 2$, the **explicit form**

$$
p_n = 2^{n+2} - 1.
$$

Use mathematical induction.
