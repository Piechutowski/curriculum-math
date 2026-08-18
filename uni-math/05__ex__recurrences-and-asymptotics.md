# Discrete Mathematics — Exercises 5: Recurrences and Asymptotics

> Translated from the handwritten lecture notes `pdf/Ćwiczenia5.pdf`.

## Exercise 1

For the sequences:

- **(i)**

$$
1^2 + 3^2 + \ldots + (2n-1)^2 = n(2n-1)(2n+1)\,\frac{1}{3}
$$

> The handwritten sheet reads $1^2 + 2^2 + \ldots + (2n-1)^2$, but the stated formula $\frac{1}{3}n(2n-1)(2n+1)$ is the sum of the squares of the first $n$ **odd** numbers, so the left-hand side should be $1^2 + 3^2 + \ldots + (2n-1)^2$.

- **(ii)**

$$
\sum_{i=1}^{n} 2^{i-1} = 2^n - 1
$$

- **(iii)**

$$
a_1 = 1, \quad a_2 = 2, \qquad a_n = a_{n-1} + a_{n-2}
$$

determine $a_n = O(?)$ and $a_n = \Theta(?)$.

For (iii) prove that

$$
0 \leq a_n \leq \left(\frac{7}{4}\right)^n
$$

## Exercise 2

Find the closed form for

$$
a_{n+2} = 4a_{n+1} - 4a_n
$$

with $a_0 = 1$ and $a_1 = 3$. Determine $a_n = O(?)$.

## Exercise 3

For the scheme

$$
2a_{n+3} = a_{n+2} + 2a_{n+1} - a_n
$$

$$
a_0 = 0, \quad a_1 = 1 \ \text{ and } \ a_2 = 2
$$

find the closed form and determine $a_n = O(?)$.

## Exercise 4

For

$$
a_{n+2} - 10a_{n+1} + 21a_n = f(n)
$$

find the form of the **homogeneous equation** and its solution. For $f_1(n) = 5$ find the general form. Trace through the procedure for $f_2(n) = 3n^2 - 2$.

## Exercise 5

For

$$
a_{n+1} - a_n = n \quad \text{ and } \quad a_2 = 1
$$

find the general solution.

## Exercise 6

For

$$
a_n = a_{n-1} + a_{n-2} + 1, \qquad a_0 = 0 \ \text{ and } \ a_1 = 1
$$

find the closed form of $a_n$.
