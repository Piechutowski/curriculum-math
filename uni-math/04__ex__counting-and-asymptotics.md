# Discrete Mathematics — Exercises 4: Counting and Asymptotics

> Translated from the handwritten lecture notes `pdf/Ćwiczenia4.pdf`.

## Exercise 1

Among the numbers $\{1, 2, \ldots, 100\}$, find the number of numbers which are **not divisible by 2 or 3 or 5**. Apply the **inclusion–exclusion principle**.

## Exercise 2

In a class: 4 people play football, 5 people play basketball, and 2 people play both football and basketball.

- **a)** How many students play at least one of the above games?
- **b)** How many students play no game at all (the class has 30 people)?

## Exercise 3

**Dirichlet's pigeonhole principle:** if $n + 1$ or more pigeons occupy $n$ holes, then at least 2 pigeons are in the same hole.

We have 5 pairs of socks in 5 different colors. After mixing them we have 10 socks. How many socks must we draw at random to be certain that we have 1 pair of the same color?

## Exercise 4

Determine the asymptotics of the sequence defined recursively in "**divide and conquer**" form:

$$
(*) \quad
\begin{cases}
a_n = 16\, a_{\lfloor n/2 \rfloor} \\
a_1 = 3
\end{cases}
\qquad \text{i.e.} \qquad a_n = O(?)
$$

First transform $(*)$ into explicit (closed) form.

## Exercise 5

Examine the asymptotics of:

$$
\text{(i)} \quad a_n = 5n^4 + 6\ln n
\qquad\qquad
\text{(ii)} \quad b_n = n^2 - \binom{n}{2} \cdot 2
$$
