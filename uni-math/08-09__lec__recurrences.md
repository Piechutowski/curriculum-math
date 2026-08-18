# Discrete Mathematics — Lectures 8–9: Recurrences

> Translated from the handwritten lecture notes `pdf/Lect8-9-Rekurencje-MatDysk.pdf`.

## Motivating example — searching a sorted list

**Example:** a list of (English) words in lexicographic order, e.g.:

| $i$ | $w_i$ |
|:---:|:------|
| 1 | Are |
| 2 | Body |
| 3 | Career |
| 4 | Computer |
| 5 | Dam |
| 6 | Mathematics |
| 7 | Science |
| 8 | Ugh |
| 9 | Why |

For a new word $w$ we want to determine **whether it is a word on the list, and at which position**.

**Input:**

- $n$ — the number of elements of the list, $n > 0$;
- $w_i$, $1 \leq i \leq n$ — the list;
- $w$ — the word searched for.

## The algorithm SEQ SEARCH

```
i := 1
repeat
    if  w = w_i  then print i; stop
        w < w_i  then print "failure"; stop
        w > w_i  then i := i + 1
    endif
    if i > n then print "failure"; stop
endrepeat

output:  i  or  "failure"
```

- **Best case:** $1$ comparison\*
- **Worst case:** $n$ comparisons

The number of comparisons **"on average"**:

$$
E_n(X) \ \longleftarrow \ \text{one computes the expected value}
$$

($E_n$ — the mean; $X$ — the number of comparisons.)

> \* De facto the 3 comparisons in the loop are 2 comparisons on a computer — but these 2 we treat as one.

## The expected number of comparisons

We assume that $w$ **is** on the list. The events $w = w_i$ are equally probable:

$$
P(\{w = w_i\}) = \frac{1}{n}
$$

The random variable $X_i$ — the number of comparisons. Then

$$
E_n(X) = \sum_{i=1}^{n} P(\{w = w_i\}) \cdot X_i
$$

— the number of comparisons (treated here as a random variable). Hence

$$
E_n(X) = \frac{1}{n}\,\bigl(\underbrace{1}_{\text{best case}} + 2 + \ldots + \underbrace{n}_{\text{worst case}}\bigr)
= \frac{1}{n} \cdot \frac{n(n+1)}{2}
$$

$$
\boxed{\ E_n(X) = \frac{n+1}{2}\ } \quad (**) \ \text{(the mean)}
$$

$E_n$ — the expected number of comparisons.

## A recurrence relation for $E_n$

One can also find a **recurrence relation**. Take any $k$ with $1 \leq k \leq n-1$ and split the list $L$ into a front part $L_1$ and a back part $L_2$:

![[lec0809_p03_list-split.svg]]

**SEQ SEARCH as a two-stage recursive process:**

- **STAGE 1** — SEQ SEARCH on $L_1$;
- **STAGE 2** — if failure on $L_1$, then SEQ SEARCH on $L_2$.

### Stage 1

$$
\begin{aligned}
\text{if } w \in L_1 &\ \Rightarrow\ E_k \\
\text{if } w \notin L_1,\ w \in L_2 &\ \Rightarrow\ \text{the number of comparisons is } k
\end{aligned}
$$

(it is surely in $L_2$, since the word is not in $L_1$ but the word **is** in $L$). The probabilities:

$$
P(w \in L_1) = \frac{k}{n}, \qquad P(w \in L_2) = \frac{n-k}{n}
$$

The (mean) number of comparisons for stage 1:

$$
= \frac{k}{n} E_k + \frac{n-k}{n} \cdot k
$$

### Stage 2

The number of comparisons is $E_{n-k}$; we reach it with probability $\frac{n-k}{n}$.

The (mean) number of comparisons for stage 2:

$$
= \frac{k}{n} \cdot 0 + \frac{n-k}{n} E_{n-k}
$$

(with probability $\frac{k}{n}$ the word was already found in $L_1$ — we do not compare, so the expected number of comparisons is $0$).

**Adding** the two stages:

$$
E_n = \frac{k}{n} E_k + \frac{n-k}{n}\bigl(k + E_{n-k}\bigr)
$$

This formula holds for all $k$, $1 \leq k \leq n-1$.

### Taking $k = 1$

$$
\begin{aligned}
E_n &= \frac{1}{n} E_1 + \frac{n-1}{n}\,(1 + E_{n-1}) \\
E_n &= \frac{1}{n} E_1 + \frac{n-1}{n} E_{n-1} + \frac{n-1}{n}
\end{aligned}
$$

But for $n = 1$ (the number of comparisons) $E_1 = 1$, so

$$
E_n = \frac{1}{n} + \frac{n-1}{n} E_{n-1} + 1 - \frac{1}{n}
$$

$$
\boxed{\ E_n = E_{n-1}\,\frac{n-1}{n} + 1, \qquad E_1 = 1\ }
$$

— a **linear difference equation of order 1** with non-constant coefficients (inhomogeneous).

We showed that this recurrence is satisfied by

$$
E_n = \frac{n+1}{2} \quad \longleftarrow \ \text{the explicit (closed) form},
$$

which is consistent with $(**)$ computed at the beginning of this lecture!

**Fibonacci:**

$$
a_0 = a_1 = 1, \qquad a_{n+1} = a_n + a_{n-1}
$$

- a linear difference equation of order 2,
- constant coefficients,
- homogeneous.

## Definition 1 — order and types of difference equations

**Order (degree) of a difference equation** — the difference between the largest and the smallest index:

$$
a_n = f(a_{n-1}, a_{n-2}, \ldots, a_{n-k}, n) \qquad \text{— order } k.
$$

**Linear** difference equation (homogeneous):

$$
a_n = f(a_{n-1}, a_{n-2}, \ldots, a_{n-k}, n) = \sum_{i=1}^{k} f_i(n)\, a_{n-i}
$$

Linear homogeneous **with constant coefficients**:

$$
a_n = f(a_{n-1}, a_{n-2}, \ldots, a_{n-k}) = \sum_{i=1}^{k} \alpha_i\, a_{n-i}
$$

($\alpha_i$ constant — it does not depend on $n$).

Linear (**inhomogeneous**):

$$
a_n = \sum_{i=1}^{k} f_i(n)\, a_{n-i} + g(n)
$$

**Examples:**

$$
\begin{array}{lll}
(0) & E_n = E_{n-1}\dfrac{n-1}{n} + 1 & \text{order 1, linear, inhomogeneous, non-constant coeff.} \\[2mm]
(i) & a_{n+1} = \dfrac{1}{2}\left(a_n + \dfrac{4}{a_n}\right) & \text{order 1, nonlinear} \\[2mm]
(ii) & a_n = a_{n-1}\, a_{n-2} & \text{order 2, nonlinear} \\[2mm]
(iii) & a_n = a_{n-10} & \text{order 10, linear, homogeneous, constant coeff.} \\[2mm]
(iv) & a_n = 1 + a_{\lfloor n/2 \rfloor} & \text{of non-constant order; linear, inhomogeneous, constant coeff.} \\[2mm]
(v) & a_{n+1} = 3a_n + n\, a_{n-1} + 2^n & \text{linear, order 2, inhomogeneous, non-constant coeff.} \\[2mm]
(vi) & a_n = a_{n-1} + n^3 - 2n^2 + 3n + 2 & \text{order 1, linear, inhomogeneous, constant coeff.} \\[2mm]
(vii) & a_{n+2} = a_{n+1} + a_n \ \text{(Fibonacci)} & \text{linear, constant coeff., homogeneous, order 2}
\end{array}
$$

Difference (recurrence) equations of order $k$ are usually supplemented with **initial conditions**

$$
a_0, a_1, \ldots, a_{k-1}
$$

in order to **guarantee uniqueness** of the solution of the difference equation.

## Theorem — existence and uniqueness

**Theorem.** Let a function

$$
f: \mathbb{R}^k \times \mathbb{N} \longrightarrow \mathbb{R}
$$

be given. Then there exists exactly one sequence $\{a_n\}_{n \geq 0}$ satisfying:

- a) $a_n = f(a_{n-1}, a_{n-2}, \ldots, a_{n-k}, n)$,
- b) $a_i = c_i$ for $0 \leq i \leq k-1$ (given numbers).

**Proof.** Let $P(n)$: there is only one $\{a_n\}_{n \geq 0}$ satisfying a) and b) for $n \in \mathbb{N}$ (a specific $n$).

**Base step:** $P(0), P(1), \ldots, P(k-1)$ are true (because b) is satisfied).

**Induction step:** $P(j)$ OK for all $j < n$ $\ \stackrel{?}{\Rightarrow}\ $ $P(n)$ true.

$$
a_n = f(a_{n-1}, a_{n-2}, \ldots, a_{n-k}, n)
$$

By the induction hypothesis the $a_i$ for $n-k \leq i \leq n-1$ are uniquely determined; $f$ is a function $\Rightarrow$ $a_n$ is also uniquely determined! By induction, $P(n)$ OK for $n \geq 0$. $\square$

## Special cases

### (I) Second-order linear homogeneous with constant coefficients

$$
\boxed{\ a_n = a\, a_{n-1} + b\, a_{n-2} \quad (*), \qquad \oplus\ a_0, a_1 \ \text{given}\ }
$$

— linear, homogeneous, order 2, constant coefficients.

**Case $b = 0$** ($a_1$ is superfluous):

$$
a_n = a\, a_{n-1} \ \Rightarrow\ \frac{a_n}{a_{n-1}} = a \qquad (a_0 \ \text{necessary})
$$

$a_n$ is a geometric sequence:

$$
\boxed{\ a_n = a^n a_0\ }
$$

**Case $a = 0$** ($a_0$, $a_1$ necessary):

$$
a_n = b\, a_{n-2} \ \Rightarrow\ \frac{a_n}{a_{n-2}} = b
$$

$$
\Rightarrow \quad a_n =
\begin{cases}
a_{2k} = b^k a_0 & \text{here } 2k = n \\
a_{2k+1} = b^k a_1 & \text{here } 2k+1 = n
\end{cases}
$$

**Example:**

(i) $S_n = 3 S_{n-1}$ (so $a = 3$), $S_0 = 5$:

$$
\Rightarrow \quad \boxed{\ S_n = 5 \cdot 3^n\ }
$$

(ii) $S_n = 3 S_{n-2}$ (so $b = 3$), $S_0 = 5$, $S_1 = 2$:

$$
\boxed{\ S_{2n} = 5 \cdot 3^n, \qquad S_{2n+1} = 2 \cdot 3^n\ } \qquad \square
$$

### Case $a \neq 0$ and $b \neq 0$ — the characteristic equation

We look for a solution of the form

$$
S_n = c\, r^n
\qquad
\left(c \neq 0, \ \text{since otherwise } S_n \equiv 0 \ \text{— that case is easy to check}\right)
$$

Substituting into $S_n = a S_{n-1} + b S_{n-2}$:

$$
c\, r^n = a\, c\, r^{n-1} + b\, c\, r^{n-2} \qquad /\!\!: c
$$

$$
\boxed{\ r^2 - a r - b = 0\ } \quad (**)
$$

— the **characteristic equation** of the recurrence relation $(*)$.

> Note: as $b \neq 0$, $r = 0$ is impossible.

We assume that $\boxed{\Delta \geq 0}$ — real roots.

### a) $\Delta > 0$

(There exist 2 solutions $r_1$ and $r_2$ of $(**)$.) The sequences

$$
\{x_n\}_{n \geq 0} = \{\widetilde{C}_1\, r_1^n\}_{n \geq 0}, \qquad
\{y_n\}_{n \geq 0} = \{\widetilde{C}_2\, r_2^n\}_{n \geq 0}
\qquad (\widetilde{C}_1, \widetilde{C}_2 \ \text{arbitrary})
$$

satisfy the difference equation (without initial conditions):

$$
\boxed{\ S_n = a\, S_{n-1} + b\, S_{n-2}\ } \quad (\blacksquare)
$$

One sees immediately that

$$
\{z_n\}_{n \geq 0} = \{\alpha x_n + \beta y_n\}_{n \geq 0} \quad \text{also satisfies } (\blacksquare).
$$

> Margin note: $r_1^n \neq r_2^n \cdot \lambda$, since $\left(\frac{r_1}{r_2}\right)^n = \lambda$ would tend to $0$ or $\infty$ when $r_1 \neq r_2$ — so the two solutions are independent.

Indeed:

$$
\begin{aligned}
\alpha\bigl(a \widetilde{C}_1 r_1^{n-1} + b \widetilde{C}_1 r_1^{n-2}\bigr) &= \alpha\bigl(\widetilde{C}_1 r_1^{n}\bigr) \\
+\ \beta\bigl(a \widetilde{C}_2 r_2^{n-1} + b \widetilde{C}_2 r_2^{n-2}\bigr) &= \beta\bigl(\widetilde{C}_2 r_2^{n}\bigr)
\end{aligned}
$$

Summing the two rows:

$$
a\bigl(\alpha \widetilde{C}_1 r_1^{n-1} + \beta \widetilde{C}_2 r_2^{n-1}\bigr)
+ b\bigl(\alpha \widetilde{C}_1 r_1^{n-2} + \beta \widetilde{C}_2 r_2^{n-2}\bigr)
= \alpha \widetilde{C}_1 r_1^{n} + \beta \widetilde{C}_2 r_2^{n}
$$

$$
a\bigl(\alpha x_{n-1} + \beta y_{n-1}\bigr) + b\bigl(\alpha x_{n-2} + \beta y_{n-2}\bigr) = \bigl(\alpha x_n + \beta y_n\bigr)
$$

$$
\boxed{\ a\, z_{n-1} + b\, z_{n-2} = z_n\ }
$$

The constants $\widetilde{C}_1$ and $\widetilde{C}_2$ are computed from the conditions $S_0$ and $S_1$ (the given initial conditions).

### b) $\Delta = 0$

$$
x_n = C_1 r_1^n, \quad \widetilde{x}_n = r_1^n \qquad \text{(one solution — we know it satisfies the equation)}
$$

$$
y_n = C_2\, n\, r_1^n, \quad \widetilde{y}_n = n\, r_1^n \qquad \text{(the second solution)}, \qquad r_1 \neq 0
$$

Here $r_1$ is a **double root** of $r^2 - a r - b = 0$:

$$
(r - r_1)^2 = 0
$$

$$
\begin{cases}
r^2 - 2 r r_1 + r_1^2 = 0 \\
r^2 - a r - b = 0
\end{cases}
\ \Rightarrow\
\begin{cases}
b = -r_1^2 \\
a = 2 r_1
\end{cases}
\quad (*)
$$

We check whether $a_n = a\, a_{n-1} + b\, a_{n-2}$ holds for $a_n = C_2\, n\, r_1^n$:

$$
\begin{aligned}
C_2\, n\, r_1^n &\stackrel{?}{=} a\, C_2 (n-1) r_1^{n-1} + b\, C_2 (n-2) r_1^{n-2} \\
&\stackrel{(*)}{=} 2 r_1 (n-1) r_1^{n-1} - (n-2)\, r_1^2\, r_1^{n-2} \\
&= 2 n r_1^n - n r_1^n - 2 r_1^n + 2 r_1^n \\
&= n\, r_1^n \qquad \text{O.K.}
\end{aligned}
$$

> Margin note: $n r_1^n$ and $r_1^n$ are linearly independent, for if $n r_1^n = \lambda r_1^n$ ($r_1 \neq 0$) then $n = \lambda$ — a contradiction for all $n \in \mathbb{N}$.

**Remark:** $\widetilde{x}_n$ and $\widetilde{y}_n$ are linearly independent!

$$
\alpha \widetilde{x}_n + \beta \widetilde{y}_n = 0 \quad \forall n \iff \alpha = \beta = 0
$$

## The space of sequences

Let $V_\infty$ denote the space of all sequences $\{a_n\}_{n \in \mathbb{N}}$ (here $a_n \in \mathbb{R}$ or $\mathbb{C}$), with operations

$$
\begin{aligned}
\{a_n\} \oplus \{b_n\} &= \{a_n + b_n\} \\
\lambda \odot \{a_n\} &= \{\lambda \cdot a_n\}
\end{aligned}
$$

$(V_\infty, \oplus, \odot)$ is a **linear space**.

Consider the subspace

$$
V_R = \bigl\{\, \{a_n\} \in V_\infty : \ a_n = a\, a_{n-1} + b\, a_{n-2}, \quad a_0 \text{ and } a_1 \text{ arbitrary} \,\bigr\}
$$

— a **linear subspace**, i.e.

$$
\{\alpha a_n + \beta b_n\} \in V_R \quad \text{as long as } \{a_n\}, \{b_n\} \in V_R \quad (*)
$$

Indeed:

$$
\begin{aligned}
c_n = \alpha a_n + \beta b_n
&\stackrel{(*)}{=} \alpha\,(a\, a_{n-1} + b\, a_{n-2}) + \beta\,(a\, b_{n-1} + b\, b_{n-2}) \\
&= (\alpha a_{n-1} + \beta b_{n-1}) \cdot a + (\alpha a_{n-2} + \beta b_{n-2}) \cdot b \\
&= a\, c_{n-1} + b\, c_{n-2}
\end{aligned}
$$

$$
\Rightarrow \quad c_n \in V_R.
$$

### $\dim(V_R) = 2$

$$
\boxed{\ \dim(V_R) = 2\ } \quad \longleftarrow \ \text{we will show this}
$$

> Margin note: $\binom{a_0}{a_1} = \lambda \binom{b_0}{b_1} \Rightarrow \{a_n\} = \lambda\{b_n\}$ — linearly dependent.

For

$$
\begin{pmatrix} a_0 \\ a_1 \end{pmatrix} \ \stackrel{(\blacktriangle)}{\neq}\ \lambda \begin{pmatrix} b_0 \\ b_1 \end{pmatrix}
\quad \text{we have} \quad
\det \begin{pmatrix} a_0 & b_0 \\ a_1 & b_1 \end{pmatrix} \ \stackrel{(\bullet)}{\neq}\ 0
$$

The sequences

$$
V_1 = \begin{pmatrix} a_0 \\ a_1 \\ a_2 \\ \vdots \\ a_n \\ \vdots \end{pmatrix}
\quad \text{and} \quad
V_2 = \begin{pmatrix} b_0 \\ b_1 \\ b_2 \\ \vdots \\ b_n \\ \vdots \end{pmatrix}
$$

are linearly independent, $V_1, V_2 \in V_\infty$. Consider

$$
(*) \qquad
\alpha \begin{pmatrix} a_0 \\ a_1 \\ \vdots \\ a_n \\ \vdots \end{pmatrix}
+ \beta \begin{pmatrix} b_0 \\ b_1 \\ \vdots \\ b_n \\ \vdots \end{pmatrix}
= \begin{pmatrix} 0 \\ 0 \\ \vdots \\ 0 \\ \vdots \end{pmatrix}
\qquad \text{is } \alpha = \beta = 0\,?
$$

But $(*)$ gives

$$
\alpha \begin{pmatrix} a_0 \\ a_1 \end{pmatrix} + \beta \begin{pmatrix} b_0 \\ b_1 \end{pmatrix} = \begin{pmatrix} 0 \\ 0 \end{pmatrix},
\qquad
\begin{bmatrix} a_0 & b_0 \\ a_1 & b_1 \end{bmatrix} \begin{bmatrix} \alpha \\ \beta \end{bmatrix} = \begin{pmatrix} 0 \\ 0 \end{pmatrix}
$$

Because of $(\bullet)$:

$$
\begin{pmatrix} \alpha \\ \beta \end{pmatrix}
= \begin{bmatrix} a_0 & b_0 \\ a_1 & b_1 \end{bmatrix}^{-1} \begin{bmatrix} 0 \\ 0 \end{bmatrix}
= \begin{bmatrix} 0 \\ 0 \end{bmatrix},
\qquad \alpha = \beta = 0
$$

Hence for arbitrary data $(\blacktriangle)$ we have 2 linearly independent vectors in $V_\infty$. So

$$
\boxed{\ \dim(V_R) \geq 2\ }
$$

But also $\dim(V_R) < 3$ (and therefore $\dim(V_R) = 2$).

> The manuscript writes $\dim(V_\infty) < 3$ here — an evident slip: the bound being proved is for the subspace $V_R$.

Indeed, consider 3 arbitrary sequences (in $V_R$):

$$
\begin{pmatrix} a_0 \\ a_1 \\ \vdots \\ a_n \\ \vdots \end{pmatrix}
\qquad
\begin{pmatrix} b_0 \\ b_1 \\ \vdots \\ b_n \\ \vdots \end{pmatrix}
\qquad
\begin{pmatrix} c_0 \\ c_1 \\ \vdots \\ c_n \\ \vdots \end{pmatrix}
$$

We will show that there exist $\alpha$ and $\beta$ such that

$$
\{c_n\} = \alpha \{a_n\} + \beta \{b_n\}
$$

— every 3rd vector is linearly dependent on the 2 remaining ones, as long as $\{a_n\}$ and $\{b_n\}$ are linearly independent.

For $n = 0$, $n = 1$: as $\binom{a_0}{a_1}, \binom{b_0}{b_1}, \binom{c_0}{c_1} \in \mathbb{R}^2$,

$$
\exists\, \widetilde{\alpha} \text{ and } \widetilde{\beta}: \qquad
\widetilde{\alpha} \begin{pmatrix} a_0 \\ a_1 \end{pmatrix} + \widetilde{\beta} \begin{pmatrix} b_0 \\ b_1 \end{pmatrix} = \begin{pmatrix} c_0 \\ c_1 \end{pmatrix}
\qquad \text{— this is } P(0), P(1).
$$

It is easy to show by induction that on every coordinate

$$
\widetilde{\alpha}\, a_n + \widetilde{\beta}\, b_n = c_n \qquad P(n)
$$

Assume $P(n-1)$, $P(n-2)$, $n \geq 2$. Since

$$
\begin{aligned}
a_n &= a\, a_{n-1} + b\, a_{n-2} \\
b_n &= a\, b_{n-1} + b\, b_{n-2}
\end{aligned}
$$

we get

$$
\begin{aligned}
\widetilde{\alpha}\, a_n + \widetilde{\beta}\, b_n
&= \widetilde{\alpha}\,(a\, a_{n-1} + b\, a_{n-2}) + \widetilde{\beta}\,(a\, b_{n-1} + b\, b_{n-2}) \\
&= a\,\underbrace{(\widetilde{\alpha} a_{n-1} + \widetilde{\beta} b_{n-1})}_{P(n-1)} + b\,\underbrace{(\widetilde{\alpha} a_{n-2} + \widetilde{\beta} b_{n-2})}_{P(n-2)} \\
&= a\, c_{n-1} + b\, c_{n-2} = c_n \qquad P(n)
\end{aligned}
$$

From $P(n)$ it follows that

$$
\widetilde{\alpha}\{a_n\} + \widetilde{\beta}\{b_n\} = \{c_n\}, \qquad \widetilde{\alpha}, \widetilde{\beta} \in \mathbb{R}
$$

Hence the linear dependence. $\square$

### Recipe

For the equation

$$
\begin{cases}
a_n = a\, a_{n-1} + b\, a_{n-2} \\
a_0, a_1 \ \text{arbitrary}
\end{cases}
$$

1. We find 2 linearly independent solutions $\{\psi_1, \psi_2\}$ $\Rightarrow$ $\psi = C_1 \psi_1 + C_2 \psi_2$.
2. If $S_0$ and $S_1$ are given, then the constants $C_1$ and $C_2$ are found from these constraints.

## Examples

### a) Distinct roots

$$
\boxed{\ a_n = a_{n-1} + 2 a_{n-2}\ } \qquad n > 1
$$

$$
W(r) = r^2 - r - 2 = 0 \qquad \text{— characteristic equation}
$$

$$
r_1 = 2, \qquad r_2 = -1
$$

So we have 2 linearly independent solutions:

$$
a_n = (-1)^n, \qquad b_n = 2^n
$$

The general one:

$$
\boxed{\ c_n = C_1 (-1)^n + C_2\, 2^n\ }
$$

Now let us add $\boxed{a_0 = a_1 = 3}$:

$$
\begin{cases}
3 = C_1 (-1)^0 + C_2\, 2^0 \\
3 = C_1 (-1)^1 + C_2\, 2^1
\end{cases}
\qquad
\begin{cases}
3 = C_1 + C_2 \\
3 = -C_1 + 2 C_2
\end{cases}
\ \Rightarrow\
\begin{aligned}
C_1 &= 1 \\
C_2 &= 2
\end{aligned}
$$

$$
\boxed{\ a_n = 2^{n+1} + (-1)^n\ }
$$

### b) A double root ($\Delta = 0$)

$$
\boxed{\ S_n = 6 S_{n-1} - 9 S_{n-2}\ }
$$

$$
r^2 - 6r + 9 = 0, \qquad r = 3 \ \text{— a double root}
$$

$$
a_n = 3^n, \qquad b_n = n\, 3^n
$$

$$
\boxed{\ S_n = \widetilde{a}_1\, 3^n + \widetilde{a}_2\, n\, 3^n\ }
$$

Adding the conditions $S_0 = 1$, $S_1 = -3$:

$$
\begin{cases}
S_0 = \widetilde{a}_1\, 3^0 + 0 = \widetilde{a}_1 \\
S_1 = \widetilde{a}_1\, 3^1 + \widetilde{a}_2 \cdot 1 \cdot 3^1 = 3\widetilde{a}_1 + 3\widetilde{a}_2
\end{cases}
\qquad
\left.
\begin{aligned}
\widetilde{a}_1 &= 1 \\
3\widetilde{a}_1 + 3\widetilde{a}_2 &= -3
\end{aligned}
\right\}
\ \Rightarrow\ \widetilde{a}_1 = 1 \ \text{and}\ \widetilde{a}_2 = -2
$$

$$
\boxed{\ S_n = 3^n - 2 \cdot n \cdot 3^n\ }
$$

## Higher-order difference equations

**Example:**

a)

$$
\boxed{\ a_n = 2 a_{n-1} + 5 a_{n-2} - 6 a_{n-3}\ }
$$

again expecting $x_n = c\, r^n$:

$$
r^3 - 2r^2 - 5r + 6 = (r-1)(r+2)(r-3) = 0
$$

$$
a_n = A \cdot 1^n + B(-2)^n + C \cdot 3^n
$$

b)

$$
\begin{cases}
b_{n+1} = 5 b_n - 3 b_{n-1} - 9 b_{n-2} \\
b_0 = 6, \quad b_1 = -8, \quad b_2 = -22
\end{cases}
$$

$$
r^3 - 5r^2 + 3r + 9 = 0 = (r+1)(r-3)^2
$$

$$
b_n = A(-1)^n + B\, 3^n + C\, n\, 3^n
$$

$$
\begin{cases}
A + B = 6 \\
-A + 3B + 3C = -8 \\
A + 9B + 18C = -22
\end{cases}
\qquad
A = 5, \quad B = 1, \quad C = -2
$$

$$
\boxed{\ b_n = 5(-1)^n + 3^n - 2n\, 3^n\ }
$$

### Definition — characteristic equation of order $k$

**Def.** For constants $c_1, c_2, \ldots, c_k$, the **characteristic equation** of the difference equation of order $k$ (homogeneous)

$$
a_n = \sum_{i=1}^{k} c_i\, a_{n-i} \quad (*) \qquad (a_0, a_1, \ldots, a_{k-1} \ \text{arbitrary})
$$

is expressed by the formula

$$
W_k(r) = r^k - \sum_{i=1}^{k} c_i\, r^{k-i} = 0
$$

Let us denote by $V_r^k$ the space of all sequences $\{a_n\}$ satisfying $(*)$ and having arbitrary values of $a_0, a_1, \ldots, a_{k-1}$.

### Lemma 1

- a) $V_r^k$ — a linear space.
- b) $\dim(V_r^k) = k$.

(The proof is similar to $k = 2$ — induction.)

### Theorem 1

Let the characteristic polynomial of $(*)$ have $r_1, r_2, \ldots, r_{l_k}$ distinct roots ($\in \mathbb{R}$, $\in \mathbb{C}$) with multiplicities $m_1, m_2, \ldots, m_{l_k}$, such that $m_1 + m_2 + \ldots + m_{l_k} = k$. That is:

$$
W_k(r) = \prod_{i=1}^{l_k} (r - r_i)^{m_i}
$$

Then every solution of $(*)$ is a linear combination of

$$
\begin{matrix}
r_1^n, & n\, r_1^n, & \ldots, & n^{m_1 - 1} r_1^n \\
r_2^n, & n\, r_2^n, & \ldots, & n^{m_2 - 1} r_2^n \\
\vdots \\
r_{l_k}^n, & n\, r_{l_k}^n, & \ldots, & n^{m_{l_k} - 1} r_{l_k}^n
\end{matrix}
\qquad (\blacktriangle)
$$

**Proof** is analogous to $k = 2$:

- **STEP 1:** by Lemma 1, $V_r^k$ is a linear space of dimension $k$.
- **STEP 2:** we look for $k$ linearly independent solutions of $(*)$, proving by induction that the forms $(\blacktriangle)$ satisfy $(*)$ and that they are linearly independent.

**Similarly:** the specification of the initial conditions

$$
c_0 = a_0, \ c_1 = a_1, \ \ldots, \ c_{k-1} = a_{k-1} \qquad + \quad (*)
$$

gives a unique solution.

**Remark:** for $k = 2$ we omitted $\Delta < 0$ (complex roots). Theorem 1 admits the case when $r_i \in \mathbb{C}$. Then the solutions are

$$
(a_i + i b_i)^n, \quad n (a_i + i b_i)^n, \quad \ldots, \quad n^{k_i - 1} (a_i + i b_i)^n
$$

— but then in $\sum c_i r_i^n$ the constants are $\in \mathbb{C}$.

## Inhomogeneous linear recurrences

Because of the applications it is important to consider:

- linear difference schemes with constant coefficients but **inhomogeneous**, e.g.

$$
T_n = 2 T_{n-1} + 1 \qquad \longleftarrow \ \text{for the Towers of Hanoi},
$$

$$
\boxed{\ a_n = \sum_{i=1}^{k} c_i\, a_{n-i} + f(n)\ } \quad (**)
$$

### Theorem 2

Assume that $\{a_n\}$ is a **particular solution** of the inhomogeneous linear difference equation $(**)$ with constant coefficients and of order $k$. Then $\{b_n\}$ is a **general solution** of the inhomogeneous equation $\iff$

$$
b_n = a_n + h_n
$$

where $h_n$ is the general solution of the corresponding homogeneous equation

$$
a_n = \sum_{i=1}^{k} c_i\, a_{n-i}.
$$

**Proof:** it is not hard — see the section "Proof of Theorem 2" below.

**Remark:** again $s_0 = a_0, \ldots, s_{k-1} = a_{k-1}$ $\Rightarrow$ a unique solution.

### Solution scheme

$$
\begin{cases}
a_n = \displaystyle\sum_{i=1}^{k} c_i\, a_{n-i} + f(n) \\
a_0, a_1, \ldots, a_{k-1} \ \text{[OPTIONALLY]}
\end{cases}
$$

**Solution:**

**I.** Solve the homogeneous equation $a_n = \sum_{i=1}^{k} c_i a_{n-i}$ by the standard method $\to h_n$ (it depends on $k$ parameters).

**II.** Find a **particular solution** of the inhomogeneous equation $a_n = \sum_{i=1}^{k} c_i a_{n-i} + f(n)$ $\to b_n$ $\longleftarrow$ **the only difficulty**.

**III.** $a_n = h_n + b_n$; we compute the constants (if $a_0, a_1, \ldots, a_{k-1}$ are given) $\Rightarrow$ the unique solution.

**Example:**

$$
\begin{cases}
a_n - 2 a_{n-1} = 3^n \\
a_0 = 4
\end{cases}
$$

We look for a particular solution of the form $b_n = A \cdot 3^n$ (the guessing method — because the right-hand side has the form $3^n$):

$$
\begin{aligned}
A\, 3^n - 2\,(A\, 3^{n-1}) &= 3^n \\
A\, 3^n - \tfrac{2}{3} A\, 3^n &= 3^n \\
\tfrac{A}{3}\, 3^n &= 3^n \ \Rightarrow\ \boxed{A = 3}
\end{aligned}
$$

$$
b_n = 3^{n+1} \qquad \text{(particular solution of the difference equation)}
$$

The characteristic polynomial for $a_n - 2 a_{n-1} = 0$:

$$
r - 2 = 0, \qquad r = 2, \qquad h_n = B\, 2^n \quad \text{(general form of the solution of the homogeneous equation)}
$$

Hence

$$
\boxed{\ a_n = 3^{n+1} + B\, 2^n\ }
$$

— the general form of the solution of the inhomogeneous equation. But

$$
a_0 = 4 = 3^1 + B\, 2^0, \qquad 4 = 3 + B \ \Rightarrow\ \boxed{B = 1}
$$

$$
\boxed{\ a_n = 3^{n+1} + 2^n\ }
$$

### How to find a special solution? (for the inhomogeneous case)

**Trick** (for the scheme $(\circledast)$):

$$
(\circledast) \quad
\begin{cases}
a_n - a_{n-1} = f(n) \\
a_0
\end{cases}
$$

$$
\begin{aligned}
a_1 &= a_0 + f(1), \\
a_2 &= a_1 + f(2) = a_0 + f(1) + f(2), \\
&\ \vdots \\
a_n &= a_0 + \sum_{i=1}^{n} f(i).
\end{aligned}
$$

We know a formula for $a_n$ if we know a formula for $\sum_{i=1}^{n} f(i)$.

> Conversely: $\sum_{i=1}^{n} f(i)$ can be computed by solving $(\circledast)$ for $a_n$ by other methods.

## The general method — finding a particular solution by example

**Example ($\blacktriangle$).** (The general method — shown on an example: how to find the particular solution.)

$$
\boxed{\ a_n + a_{n-1} - 6 a_{n-2} = 2^n - 1\ } \quad (\blacksquare)
$$

There are no initial conditions (if there are some, we take them into account in step 9).

**STEP 1:** $f(n) = 2^n - 1$.

**STEP 2:** the characteristic polynomial for $a_n + a_{n-1} - 6 a_{n-2} = 0$:

$$
p(r) = r^2 + r - 6 = (r+3)(r-2)
$$

**STEP 3:** $f(n) = 2^n - 1$ is itself a solution of the recurrence $(\odot)$: $a_n - 3 a_{n-1} + 2 a_{n-2} = 0$, whose characteristic polynomial is

$$
q(r) = (r-2)(r-1) = r^2 - 3r + 2
$$

**STEP 4:**

$$
P(r) = p(r) \cdot q(r) = (r+3)(r-2)^2 (r-1)
$$

$$
H(n) = A(-3)^n + B\, 2^n + C\, n\, 2^n + D\, (1)^n
$$

$\uparrow$ the general solution of the difference equation corresponding to $P(r)$.

**STEP 5:** we remove from $H(n)$ the part

$$
D(n) = A(-3)^n + B\, 2^n
$$

($\uparrow$ the solution of $a_n + a_{n-1} - 6 a_{n-2} = 0$; roots $r_1 = 2$, $r_2 = -3$). The remaining part

$$
\boxed{\ R(n) = C\, n\, 2^n + D \cdot 1^n\ }
$$

is the **guess for the particular solution** of the inhomogeneous equation.

**STEP 6:** we substitute $R(n)$ into $(\blacksquare)$:

$$
\bigl(C n\, 2^n + D\bigr) + \bigl(C(n-1)\, 2^{n-1} + D\bigr) - 6\bigl(C(n-2)\, 2^{n-2} + D\bigr) = 2^n - 1
$$

$$
\Downarrow
$$

$$
\boxed{\ \frac{5}{2}\, C\, 2^n - 4D = 2^n - 1\ } \qquad \forall n \in \mathbb{N}
$$

$$
\frac{5}{2} C = 1, \qquad -4D = -1
$$

$$
\Downarrow
$$

$$
C = \frac{2}{5}, \qquad D = \frac{1}{4}
$$

**STEP 7:**

$$
\boxed{\ a_n = \frac{2}{5}\, n\, 2^n + \frac{1}{4}\ } \qquad \text{(the particular solution of } \blacksquare\text{)}
$$

**STEP 8:** $h_n = A(-3)^n + B\, 2^n$, so

$$
\Rightarrow \quad b_n = A(-3)^n + B\, 2^n + \frac{2n}{5}\, 2^n + \frac{1}{4}
$$

**STEP 9:** there are no initial conditions $a_0$, $a_1$. If there were, the constants $A$ and $B$ would get computed.

## Proof of Theorem 2

$$
\begin{aligned}
R_n &= \sum_{i=1}^{k} c_i\, R_{n-i} + f(n) \qquad \text{(IH)} \\
R_n &= \sum_{i=1}^{k} c_i\, R_{n-i} \qquad\qquad\ \ \text{(H)}
\end{aligned}
$$

$\{a_n\}$ — a special solution of (IH); $\{h_n\}$ — the general solution of (H).

$(\Rightarrow)$

$$
\begin{aligned}
a_n + h_n &= \sum_{i=1}^{k} c_i\, a_{n-i} + f(n) + \sum_{i=1}^{k} c_i\, h_{n-i} \\
&= \sum_{i=1}^{k} c_i\,(a_{n-i} + h_{n-i}) + f(n)
\end{aligned}
$$

$$
b_n = (a_n + h_n) = \sum_{i=1}^{k} c_i\,\underbrace{(a_{n-i} + h_{n-i})}_{b_{n-i}} + f(n)
\quad \Rightarrow \quad b_n \ \text{satisfies (IH)}.
$$

$(\Leftarrow)$ If $b_n$ is the general solution of (IH) and $a_n$ is a particular solution of (IH):

$$
\begin{aligned}
h_n = b_n - a_n &= \sum_{i=1}^{k} c_i\, b_{n-i} + f(n) - \sum_{i=1}^{k} c_i\, a_{n-i} - f(n) \\
&= \sum_{i=1}^{k} c_i\,\underbrace{(b_{n-i} - a_{n-i})}_{h_{n-i}}
\end{aligned}
$$

so $h_n$ satisfies

$$
\boxed{\ R_n = \sum_{i=1}^{k} c_i\, R_{n-i}\ } \qquad \square
$$

## Table of predicted particular solutions

> The last two pages come from the companion notes "Algorithms & Complexity, Lecture 6" (pages 14–15 there).

The teacher's table predicting the form of the particular solution $b_n$ from the form of $f(n)$:

![[lec0809_p28_particular-solution-table.svg]]

Things get trickier if the summand $f_1(n)$ of $f(n)$ is a constant multiple of a solution of the associated homogeneous system. This happens e.g. when $f(n)$ contains a summand such as $c\, r^n$ or $(c_1 + c_2 n)\, r^n$ and $r$ is a root of the characteristic polynomial.

Then we multiply such a solution by the smallest power of $n$, say $n^s$, for which no summand $n^s f_1(n)$ is a solution of the associated homogeneous relation.

## Advantages and disadvantages of the method

The method based on I–IV has its:

**advantages:**

- the algorithmic set of steps to transform the implicit form into the explicit form of the recursive formula, permitting e.g. to find $a_n = O(f(n))$;

**disadvantages:**

- it applies only to the schemes

$$
a_n = \sum_{i=1}^{k} c_i\, a_{n-i} + f(n);
$$

- it relies on finding the roots of the characteristic polynomial (which can be hard).
