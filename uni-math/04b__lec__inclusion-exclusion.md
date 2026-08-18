# Discrete Mathematics — Lecture 4B: Counting — Sum Rule, Product Rule, Inclusion–Exclusion

> Translated from the handwritten lecture notes `pdf/lect4B-ZliczWlaczWyLacz-MatDysk.pdf`.

## The sum rule

**a)** If $A \cap B = \emptyset$:

$$
(*) \quad |A \cup B| = |A| + |B|
$$

**b)** If $A \cap B \neq \emptyset$:

$$
(**) \quad |A \cup B| = |A| + |B| - |A \cap B|
$$

Formula $(*)$ is a special case of formula $(**)$.

The teacher's doodle — two overlapping blobs with $A = \{e_1, e_2, \cdot\}$ and $B = \{\cdot, f_1\}$ sharing one element:

![[lec04b_p01_sum-rule-blobs.svg|280]]

$$
|A \cup B| = 3 + 2 - 1 = 4
$$

### Example

How many natural numbers are there in the set $S = \{1, 2, 3, \ldots, 1000\}$ divisible by 3 or by 5 or by both?

$$
D_3 = \{ n \in S : 3 \mid n \}, \qquad D_5 = \{ n \in S : 5 \mid n \}
$$

$$
|D_3 \cup D_5| = |D_3| + |D_5| - |D_3 \cap D_5|
$$

$$
|D_3| = \lfloor 1000/3 \rfloor = 333, \qquad |D_5| = \lfloor 1000/5 \rfloor = 200
$$

$$
D_3 \cap D_5 = \{ n \in S : 3 \mid n \ \text{and} \ 5 \mid n \} = \{ n \in S : 15 \mid n \}
$$

$$
|D_3 \cap D_5| = \lfloor 1000/15 \rfloor = 66
$$

$$
|D_3 \cup D_5| \stackrel{(**)}{=} 333 + 200 - 66 = \underline{\underline{467}} \qquad \square
$$

The formula $(**)$ has its generalization (to come later).

## The product rule

For a finite number of sets $S_1, S_2, \ldots, S_k$, the set of all tuples satisfies:

$$
(\Pi) \quad |S_1 \times S_2 \times \ldots \times S_k| = \prod_{i=1}^{k} |S_i|
$$

### Example — registration plates

A car registration plate consists of 2 letters + 3 digits, e.g. RD357 or SK092. How many different cars can we register in city X (registered with the full plate of city X)?

We have at our disposal 26 letters and 10 digits. Since letters and digits may repeat:

$$
10 \cdot 10 \cdot 10 \cdot 26 \cdot 26 \stackrel{(\Pi)}{=} 676{,}000 \ \text{cars}
$$

### Example — multiplying matrices

How many multiplications and additions must be performed when multiplying $n \times n$ matrices?

$$
A \in M_{n \times n}(K), \qquad B \in M_{n \times n}(K)
$$

where $K = \mathbb{R}$ or e.g. $\mathbb{C}$ (an arbitrary field). Let $C = A \circ B \in M_{n \times n}(K)$, $C = \{c_{ij}\}_{i,j=1}^{n}$:

$$
(\ast) \quad c_{ij} = \sum_{k=1}^{n} a_{ik} \, b_{kj}
$$

The teacher's sketch — the $i$-th row of $A$ times the $j$-th column of $B$:

![[lec04b_p03_matrix-row-column.svg]]

The formula $(\ast)$ contains $\underline{n}$ multiplications and $\underline{n-1}$ additions. But there are $n \times n = n^2$ coefficients $c_{ij}$. So we have:

$$
\begin{aligned}
n^2 \cdot n &= n^3 \quad \text{multiplications} \quad (\odot) \\
n^2 \cdot (n-1) &= n^3 - n \quad \text{additions} \quad (\oplus)
\end{aligned}
$$

Altogether ($\oplus$ and $\odot$) we have $2n^3 - n$ operations.

## Counting the number of functions

### Example

**a)** Let us determine all the functions (their number)

$$
f : \underbrace{\{a, b\}}_{X} \longrightarrow \underbrace{\{1, 2, 3\}}_{Y}
$$

$$
\boxed{\,W = |Y|^{|X|}\,} \ \longleftarrow \ \text{proof by mathematical induction}
$$

The sought number is $W = 3^2 = 9$ — **variations with repetition**. Indeed:

$$
\begin{aligned}
&\text{(1)} \begin{cases} f_1(a) = 1 \\ f_1(b) = 1 \end{cases}
\quad \text{(2)} \begin{cases} f_2(a) = 2 \\ f_2(b) = 2 \end{cases}
\quad \text{(3)} \begin{cases} f_3(a) = 3 \\ f_3(b) = 3 \end{cases} \\[4pt]
&\text{(4)} \begin{cases} f_4(a) = 1 \\ f_4(b) = 2 \end{cases}
\quad \text{(5)} \begin{cases} f_5(a) = 2 \\ f_5(b) = 1 \end{cases}
\quad \text{(6)} \begin{cases} f_6(a) = 1 \\ f_6(b) = 3 \end{cases} \\[4pt]
&\text{(7)} \begin{cases} f_7(a) = 3 \\ f_7(b) = 1 \end{cases}
\quad \text{(8)} \begin{cases} f_8(a) = 2 \\ f_8(b) = 3 \end{cases}
\quad \text{(9)} \begin{cases} f_9(a) = 3 \\ f_9(b) = 2 \end{cases}
\end{aligned}
$$

Indeed — 9 functions.

**b)** How many of these functions are **injective** (one-to-one)?

> Margin note: for $|X| = k$, $|Y| = n$:
> $$V_k^n = \binom{n}{k} \cdot k! = n(n-1)\cdots(n-k+1), \qquad n \geq k$$
> (proof — induction).

As many as there are **variations without repetition**, i.e.

$$
V_2^3 = \binom{3}{2} \cdot 2! = 3 \cdot 2 = \underline{\underline{6}}
$$

Indeed, $f_4, f_5, f_6, f_7, f_8, f_9$ are the only injective functions (there are 6 of them).

**c)** Another question: among the injective functions, how many are **strictly increasing**? (We assume that $X$ and $Y$ carry an order relation $<$ or $>$; e.g. $x < y \Rightarrow f(x) < f(y)$.)

Let us index the (potential) values of the function by position, $|Y| = 3$:

$$
\{ y_{\mathrm{I}}, y_{\mathrm{II}}, y_{\mathrm{III}} \}, \qquad y_{\mathrm{I}} < y_{\mathrm{II}} < y_{\mathrm{III}}
$$

The possible layouts for $|X| = 2$:

$$
\{ y_{\mathrm{I}}, y_{\mathrm{II}} \}, \qquad \{ y_{\mathrm{I}}, y_{\mathrm{III}} \}, \qquad \{ y_{\mathrm{II}}, y_{\mathrm{III}} \}
$$

That is: all two-element subsets of a 3-element set (**combinations**):

$$
C_2^3 = \binom{3}{2} = 3
$$

Indeed, $f_4, f_6, f_8$ are the only increasing functions.

## Example — diagonals, cards

**a)** How many diagonals are there in a convex polygon? We have a polygon with $n$ vertices (convexity — every pair of vertices gives a diagonal or a side).

![[lec04b_p06_polygon-diagonals.svg]]

**Method 1.** The number of connections: each pair of vertices forms a pair $\{w_i, w_j\}$.

$$
\binom{n}{2} \ \longleftarrow \ \text{the number of 2-element subsets}
$$

But the pairs of *neighbouring* vertices, $\{w_i, w_{i+1}\}$ and $\{w_{i+1}, w_{i+2}\}$, do not give diagonals (they are the sides). There are

$$
\{w_1, w_2\}, \{w_2, w_3\}, \{w_3, w_4\}, \ldots, \{w_{n-1}, w_n\}, \{w_n, w_1\}
$$

— that is, $n$ of such 2-element subsets not representing a diagonal. Hence the answer:

$$
\binom{n}{2} - n
$$

**Method 2.** From each vertex we have $(n-1) - 2 = n - 3$ diagonals (we discard the vertex itself and its two neighbours: $(w_i, w_i)$, $(w_i, w_{i+1})$, $(w_i, w_{i-1})$ do not count as diagonals; for now the edges are *directed*):

![[lec04b_p06_vertex-neighbors.svg|280]]

But the directed edges

$$
I_h = \underbrace{n}_{\text{number of vertices}} \cdot \underbrace{(n-3)}_{\text{number of edges from one vertex}}
$$

are counted twice! So:

$$
\frac{n(n-3)}{2} \ \longleftarrow \ \text{question: does it equal } \binom{n}{2} - n \, ?
$$

$$
\binom{n}{2} - n = \frac{n!}{2!\,(n-2)!} - n = \frac{n(n-1)}{2} - \frac{2n}{2} = \frac{n^2 - 3n}{2} = \frac{n(n-3)}{2} \quad \checkmark
$$

**b)** How many ways are there to draw 5 cards (**with replacement**) from a deck of 52 cards?

**Product rule:** $D = \{ 52 \text{ cards} \}$:

$$
|D \times D \times D \times D \times D| = 52 \cdot 52 \cdot 52 \cdot 52 \cdot 52 = 52^5
$$

**Variations with repetition:** each drawing is a function

$$
f : \overbrace{\{ \mathrm{I}, \mathrm{II}, \mathrm{III}, \mathrm{IV}, \mathrm{V} \}}^{X} \longrightarrow D
$$

($f(\mathrm{II}) =$ the card in the 2nd draw), so

$$
W_X^D = |D|^{|X|} = 52^5.
$$

**c)** Similarly as in b), but now we do **not** return the card. Then $f : X \to D$ must be injective (**variations without repetition**):

$$
\binom{52}{5} \cdot 5! = \text{Result}
$$

Or from the product rule:

$$
\begin{aligned}
\text{1st draw} &\ - \ 52 \ \text{ways} \\
\text{2nd draw} &\ - \ 51 \ \text{ways} \\
\text{3rd draw} &\ - \ 50 \ \text{ways} \\
\text{4th draw} &\ - \ 49 \ \text{ways} \\
\text{5th draw} &\ - \ 48 \ \text{ways}
\end{aligned}
$$

$$
\text{Result} = 52 \cdot 51 \cdot 50 \cdot 49 \cdot 48 = \binom{52}{5} \, 5!
$$

> The teacher's heading for c) literally repeats "we return the card", but the computation that follows (injective functions, $\binom{52}{5} \cdot 5!$) makes clear that drawing is *without* replacement.

## Example — two-digit numbers

How many two-digit numbers can be built from the digits $\{0, 1, 2, \ldots, 9\}$?

We know from school: there are $90$ of them, since $10, 11, 12, \ldots, 99$.

Another way: each number $\overline{xy}$ corresponds to a function $f$:

$$
f : \{ \mathrm{I}, \mathrm{II} \} \longrightarrow \{ 0, 1, 2, \ldots, 9 \}
$$

$$
f(\mathrm{I}) = x \quad (\text{tens position}), \qquad f(\mathrm{II}) = y \quad (\text{units position})
$$

So we have $10^2$ numbers, and

$$
10^2 - 10 = 90
$$

— we discard $\{00, 01, \ldots, 09\}$.

## Example — subsets of a 3-element set

**b)** Let $\Omega = \{e_1, e_2, e_3\}$. How many subsets does $\Omega$ have?

$$
|\mathcal{P}(\Omega)| = 2^{|\Omega|} = 8
$$

$$
\begin{array}{llc}
\emptyset \subseteq \Omega & \text{one zero-element subset} & \binom{3}{0} = 1 \\[6pt]
\{e_1\},\ \{e_2\},\ \{e_3\} \subseteq \Omega & \text{three one-element subsets} & \binom{3}{1} = 3 \\[6pt]
\{e_1, e_2\},\ \{e_2, e_3\},\ \{e_1, e_3\} \subseteq \Omega & \text{three two-element subsets} & \binom{3}{2} = 3 \\[6pt]
\{e_1, e_2, e_3\} \subseteq \Omega & \text{one three-element subset} & \binom{3}{3} = 1
\end{array}
$$

$$
1 + 3 + 3 + 1 = \underline{\underline{8}} \quad \checkmark
$$

**Note:** the order is *not* important in subsets.

## Example — groups, full houses, coins, round tables

**a)** In how many ways can we choose, from a group of 17 boys and 8 girls, a subgroup of 2 boys and 3 girls?

$$
\binom{17}{2} \cdot \binom{8}{3}
$$

— to each pair of boys we attach a triple of girls.

**b)** How many poker hands are **full houses**?

$$
(J\clubsuit,\ J\diamondsuit,\ J\heartsuit,\ 8\clubsuit,\ 8\heartsuit)
\ \neq \
(J\spadesuit,\ J\diamondsuit,\ J\heartsuit,\ 8\clubsuit,\ 8\heartsuit)
$$

— different full houses. But

$$
(J\clubsuit,\ J\diamondsuit,\ J\heartsuit,\ 8\clubsuit,\ 8\heartsuit)
\quad \text{and} \quad
(J\clubsuit,\ J\diamondsuit,\ J\heartsuit,\ 8\heartsuit,\ 8\clubsuit)
$$

— the **same** full house.

We denote the jack–eight full house by its **type**:

$$
\text{type of full house} \
\begin{cases}
(J, 8) & \text{(3 jacks, 2 eights)} \\
(8, J) & \text{(3 eights, 2 jacks)}
\end{cases}
\qquad (J,8) \neq (8,J)
$$

The type of a full house is an *injective* function:

$$
f : \{1, 2\} \longrightarrow \{2, 3, 4, 5, 6, 7, 8, 9, 10, J, Q, K, A\}
$$

in our case a): $f(1) = J$, $f(2) = 8$.

The number of types, i.e. of such functions $f$:

$$
V_2^{13} = 13 \cdot 12
$$

The number of full houses of each type (since we have 4 suits):

$$
\binom{4}{3} \cdot \binom{4}{2} = 4 \cdot 6 = 24
$$

The number of all full houses:

$$
24 \cdot 13 \cdot 12 = \underline{\underline{3744}}
$$

**c)** How many possibilities are there that in 10 coin tosses heads comes up either 3 times, or 4 times, or 9 times?

A sequence such as

$$
(H, T, T, T, H, T, T, H, T, T)
$$

can be treated as $\{e_1, e_5, e_8\}$ — a subset of the abstract set $\{e_1, e_2, e_3, e_4, e_5, e_6, e_7, e_8, e_9, e_{10}\}$. Hence:

$$
\begin{aligned}
\text{3 heads} &\ - \ \binom{10}{3} \\
\text{4 heads} &\ - \ \binom{10}{4} \\
\text{9 heads} &\ - \ \binom{10}{9}
\end{aligned}
\qquad \Longrightarrow \qquad
\text{Result} = \binom{10}{3} + \binom{10}{4} + \binom{10}{9}
$$

**d)** In how many ways can $n$ people be seated

- **a)** in a row of $n$ chairs (for $n = 10$):

$$
P_{n=10} = 10! \qquad \text{— permutations}
$$

- **b)** at a round table with $n$ chairs (arrangements equivalent up to rotation):

$$
\frac{10!}{10} = 9!
$$

For $n = 3$ — the three row-permutations map to circular arrangements which are rotations of one another, and we consider rotations equivalent arrangements (there are $n = 3$ rotations):

![[lec04b_p12_round-tables.svg]]

## Example — cutting a round cake

We have a round cake with a given number of points on its circumference. The knife cuts the cake along "**chords**" joining every pair of points. How many pieces of cake do we obtain if there are $N$ points?

> We assume that any intersection point inside the cake comes from exactly **2** chords!

$$
\begin{aligned}
N = 1: &\quad S_1 = 1 = 2^0 \ \text{(the whole cake)} \\
N = 2: &\quad S_2 = 2 = 2^1 \\
N = 3: &\quad S_3 = 4 = 2^2 \\
N = 4: &\quad S_4 = 8 = 2^3
\end{aligned}
$$

![[lec04b_p12_cake-small-n.svg]]

**Hypothesis:**

$$
\boxed{\,S_N = 2^{N-1}\,} \ : \ P(N) \ ?
$$

For $N = 5$: $S_5 = 16 = 2^4$ — it would seem that $S(N) = 2^{N-1}$ (hypothesis $P(N)$). But for $N = 6$:

![[lec04b_p13_cake-n5-n6.svg]]

$$
S_6 = 31 \neq 2^5 \qquad P(6) \ \text{is NOT satisfied! (a surprise)}
$$

$$
S_7 = 57 \neq 2^6 \qquad P(7) \ \text{is not satisfied either!}
$$

**What next???**

— It is clear that the number of cake pieces grows as $N$ grows. We add one point $a$, i.e. one extra cut (1 chord); each **segment** of the new chord divides an existing piece into 2 sub-pieces:

![[lec04b_p13_add-point.svg|340]]

### Adding one secant

- it divides a piece into 2 parts;
- the number of added pieces depends on the number of **segments**, and these are delimited by the interior intersection points (with the other secants). Each intersection point corresponds to exactly 2 secants (i.e. to 4 points on the circumference).

Let:

- $N_S$ — the number of segments on the secant $=$ the number of added pieces,
- $N_I$ — the number of intersection points on the secant (interior crossing points).

$$
\boxed{\,N_S = N_I + 1\,}
$$

**a)** we add a secant crossing no other chord — the number of pieces grows by

$$
N_S = 0 + 1 = 1 \qquad (\text{indeed: } 6 - 5)
$$

**b)** we add a secant crossing one chord — the number of pieces grows by

$$
N_S = 1 + 1 = 2 \qquad (\text{indeed: } 6 - 4)
$$

![[lec04b_p14_add-secant.svg]]

### Adding $c$ cuts

What happens when we add $c$ cuts?

$$
f_{\text{added}} = c \cdot 1 + \ \text{the sum of all intersection points lying on these } c \text{ chords}
$$

![[lec04b_p15_two-chords.svg]]

$$
f_{\text{added}} = 2 \cdot 1 + 2 = 4 \qquad (\text{indeed: } 8 - 4 = 4).
$$

### The general formula

Assuming $N \geq 4$, we treat the problem as an extension of the case $N = 3$:

- $f_N$ — the number of pieces for $N$ points on the circumference;
- $f_3$ — the number of pieces for 3 points on the circumference;
- $f_{\text{added}}^{N-3}$ — the number of pieces added by adjoining $N-3$ new points to the case $N = 3$ and drawing all the possible cuts.

$$
\boxed{\,f_N = f_3 + f_{\text{added}}^{N-3}\,} \qquad \text{but } f_3 = 4
$$

$$
\boxed{\,f_N = 4 + f_{\text{added}}^{N-3}\,}
$$

The case $N = 5$ — from the triangle configuration to all cuts drawn:

![[lec04b_p16_n5-sequence.svg]]

Using the previous analysis:

$$
f_{\text{added}}^{N-3} = \underbrace{\text{the number of all added cuts}}_{C^{N-3}} + \underbrace{\text{the number of all intersection points inside the cake}}_{No_I^{N-3}} = C^{N-3} + No_I^{N-3}
$$

$$
C^{N-3} = \underbrace{\text{the number of all cuts}}_{\binom{N}{2}} - \ \text{the number of cuts for } N=3 \ = \binom{N}{2} - \binom{3}{2}
$$

$$
No_I^{N-3} = \binom{N}{4} \ \longleftarrow \ \text{every 4 points on the circumference give exactly one interior intersection point}
$$

Hence:

$$
f_N = \binom{N}{4} + 4 + \binom{N}{2} - \binom{3}{2}
$$

$$
(**) \quad \boxed{\ f_N = 1 + \binom{N}{2} + \binom{N}{4}\ } \qquad N \geq 4
$$

$$
f_N = 1 + \frac{N(N-1)}{2} + \frac{N(N-1)(N-2)(N-3)}{4!}
$$

$$
(*) \quad \boxed{\ f_N = \frac{N^4 - 6N^3 + 23N^2 - 18N + 24}{24}\ }
$$

From this it is easy to check that:

$$
f_N = \Theta(N^4)
$$

Moreover

$$
f_1 = 1, \qquad f_2 = 2, \qquad f_3 = 4
$$

agree with $(*)$. Hence $(*)$ is true for all $N \geq 1$.

**Remark.** Problems of this type are usually **tricky**!!! Formula $(*)$ is not something one could simply guess. The combinatorial approach, however, yields $(**)$ — which is equivalent to $(*)$.

## Example — digit sums and balls in boxes

**a)** How many numbers from the set $\{1, 2, 3, \ldots, 100000\}$ have the property that the sum of their digits equals 7?

We may omit the last number and allow $0$ at the front instead (the number $100000$ does not satisfy the condition, and $0$ does not satisfy the condition that the digit sum $= 7$ either). Now we deal with numbers of one up to five digits ($0$ – $99999$).

The teacher's sketch — digits as balls placed in boxes labelled by the decimal places:

![[lec04b_p18_digit-boxes.svg]]

$$
\underline{0\ 0\ 1\ 4\ 2} \ \text{— the number } 142, \qquad \underline{3\ 0\ 1\ 2\ 1} \ \text{— the number } 30121
$$

Our problem is therefore equivalent to placing **7 identical balls in 5 boxes**. The number of such placements is:

$$
\binom{7+5-1}{5-1} = \binom{11}{4} = \underline{\underline{330}}
$$

> Margin note ($n$ balls, $k$ boxes): each placement corresponds to a $0/1$ string of length $n + k - 1$ with $k-1$ ones (the box separators), e.g. $11010000100$ for a) and $00011010010$ for b). There are $\binom{n+k-1}{k-1}$ of them (choosing the positions of the 1s), equivalently $\binom{n+k-1}{n}$ (choosing the positions of the 0s). Explanation — see the theoretical appendix below.

**b)** In how many ways can 20 people be assigned to 5 teams:

$$
\mathrm{I} - 5 \ \text{people}, \quad \mathrm{II} - 5 \ \text{people}, \quad \mathrm{III} - 3 \ \text{people}, \quad \mathrm{IV} - 4 \ \text{people}, \quad \mathrm{V} - 3 \ \text{people} \ ?
$$

The order *inside* a team is not important. This is an arrangement of the 20 people in a row where the order within each of $\mathrm{I}, \mathrm{II}, \mathrm{III}, \mathrm{IV}, \mathrm{V}$ plays no role:

![[lec04b_p19_teams.svg]]

**Permutations with repetitions:**

$$
\frac{20!}{5! \; 5! \; 3! \; 4! \; 3!}
$$

## Theoretical appendix — stars and bars explained

(For $n = 5$ balls and $k = 3$ boxes.) A placement of balls in boxes corresponds to a $0/1$ string of length $n + k - 1$: the $0$s are the balls, the $1$s are the $k - 1$ separators between boxes:

![[lec04b_p19_stars-bars.svg]]

For a) and b): there are as many such strings as there are ways to place the $1$s (the $1$s mark a subset of the positions $\{e_1, e_2, \ldots, e_7\}$ — namely, where the $1$s stand). So there are

$$
\binom{n+k-1}{k-1}
$$

of them.

One can, however, ask alternatively: how many strings are there when counted by the **zeros**? As many as there are subsets of $\{e_1, e_2, \ldots, e_7\}$ marking where the $0$s stand:

$$
\binom{n+k-1}{n}
$$

**Note:** but

$$
\binom{n+k-1}{k-1} = \binom{n+k-1}{n}
$$

— both formulas are OK.
