# Module 01 — Mr. Greene's College Math Course · Checklist & Map

<!-- progress:01-algebra-trigonometry -->
![01 · Algebra &amp; Trigonometry — Greene progress](progress/01-algebra-trigonometry.svg)

**0 / 220 lessons complete · 0%**
<!-- endprogress -->

**Resource:** Mr. Greene's course (you are working through it now).
**How to tick:** one checkbox per skill, using the course's own lesson numbers. Tick a box only
when all three parts of the rhythm are done — Lesson watched, Study Notes written, Practice Test
passed. (The course's separate Study Notes / Practice Test entries are folded into their
lesson's box.) After editing checkboxes, run `go run ./tools/progress` to refresh the bars.

Tags: `[GFX]` graphics/simulation · `[CMP]` compilers · `[NET]` networking · `[DB]` databases ·
`[CRY]` cryptography. No tag = foundation that later lessons stand on.

---

## Section 1: Review of Basic Concepts
Goal is *automaticity*. Two ideas quietly return later: interval notation (calculus limits and
continuity) and absolute value as *distance* (becomes ε–δ limits and every numerical error
bound).

- [ ] 1. Welcome to the Course
- [ ] 2. Classifying Real Numbers
- [ ] 7. Exponents and the Order of Operations (PEMDAS)
- [ ] 12. Ratios, Rates, and Proportions
- [ ] 17. Properties of Real Numbers
- [ ] 23. Inequalities, the Number Line, and Interval Notation
- [ ] 28. Absolute Value and the Distance Between Two Points on a Number Line

## Section 2: Review of Polynomials `[CMP]` `[CRY]` `[NET]`
Polynomial arithmetic returns in heavy disguise: polynomial rings over finite fields are how CRC
checksums `[NET]` and error-correcting codes work (Rosen 12.8). Special attention: **synthetic
division (55)** — it is Horner's method, the standard way programs evaluate polynomials.

- [ ] 33. The Product and Power Rules for Exponents
- [ ] 37. What is a Polynomial?
- [ ] 41. Adding and Subtracting Polynomials
- [ ] 45. Multiplying Polynomials
- [ ] 49. Special Polynomial Products
- [ ] 52. Dividing Polynomials
- [ ] 55. Synthetic Division
- [ ] 59. Factoring Out the Greatest Common Factor (GCF)
- [ ] 62. Factoring by Grouping
- [ ] 65. Factoring Trinomials when the Leading Coefficient is 1
- [ ] 68. Factoring Trinomials using the AC Method
- [ ] 71. Factoring Trinomials using Reverse FOIL (Trial and Error)
- [ ] 74. Special Factoring
- [ ] 77. Factoring by Substitution

## Section 3: Review of Rational Expressions
Feeds rational functions (§11) and the algebra of calculus limits. Restricted values are your
first taste of domain reasoning.

- [ ] 81. Finding the Restricted Values for a Rational Expression
- [ ] 84. Simplifying Rational Expressions
- [ ] 87. Multiplying and Dividing Rational Expressions
- [ ] 90. Finding the LCD for Rational Expressions
- [ ] 93. Adding and Subtracting Rational Expressions
- [ ] 96. Complex Rational Expressions

## Section 4: Review of Radical Expressions
Rational exponents are used constantly in calculus differentiation. Fluency, not depth.

- [ ] 100. Negative Exponents and the Quotient Rule
- [ ] 103. Rational Exponents (Fractional Exponents)
- [ ] 106. Radical Expressions
- [ ] 109. Simplifying Radical Expressions
- [ ] 112. Operations with Radical Expressions
- [ ] 115. Rationalizing Denominators

## Section 5: Solving Linear Equations and Inequalities
Foundation. Compound inequalities and interval reasoning return in convergence arguments and
error bounds.

- [ ] 119. The Addition Property of Equality
- [ ] 122. The Multiplication Property of Equality
- [ ] 125. Solving Linear Equations in One Variable
- [ ] 128. Solving Linear Equations with Fractions or Decimals
- [ ] 131. Repeating Decimal to Fraction
- [ ] 134. Identifying Types of Equations
- [ ] 137. Solving Proportion Equations
- [ ] 140. Solving Literal Equations
- [ ] 143. Solving Linear Inequalities in One Variable
- [ ] 146. Solving a Three-Part Inequality in One Variable
- [ ] 149. Solving Compound Inequalities with "and" or "or"
- [ ] 152. Solving Absolute Value Equations I
- [ ] 155. Solving Absolute Value Inequalities I

## Section 6: Applications of Linear Equations `[GFX]`
Translating a described situation into equations is *modeling* — the core skill of simulation.
The motion problems (168) are physics with the calculus removed.

- [ ] 159. Translating Phrases into Algebraic Expressions and Equations
- [ ] 162. Consecutive Integer Word Problems
- [ ] 165. Age Word Problems
- [ ] 168. Motion Word Problems
- [ ] 171. Mixture Word Problems
- [ ] 174. Money Word Problems
- [ ] 177. Simple Interest Word Problems
- [ ] 180. Percent Word Problems

## Section 7: Quadratic Equations and Inequalities `[GFX]` `[CRY]`
Quadratics are projectile trajectories; the discriminant later decides whether a damped spring
oscillates or settles (Calculus §10). Complex numbers become 2D rotations (§25, Linear Algebra).

- [ ] 184. The Imaginary Unit i
- [ ] 187. Operations with Complex Numbers
- [ ] 190. Simplifying Powers of i
- [ ] 193. Solving Quadratic Equations by Factoring
- [ ] 196. The Square Root Property
- [ ] 199. Completing the Square
- [ ] 202. The Quadratic Formula
- [ ] 205. Quadratic in Form
- [ ] 208. Applications of Quadratic Equations
- [ ] 211. Solving Quadratic Inequalities

## Section 8: Radical and Rational Equations
Manipulation fluency; the equation-solving muscle the rest of the course assumes.

- [ ] 215. Solving Radical Equations I
- [ ] 218. Solving Radical Equations II
- [ ] 221. Solving Radical Inequalities I
- [ ] 224. Solving Radical Inequalities II
- [ ] 227. Solving Rational Equations
- [ ] 230. Solving Rational Inequalities
- [ ] 233. Applications of Rational Equations

## Section 9: Advanced Absolute Value
Absolute value = distance on the line — the exact idea calculus uses to define limits and
numerics uses to measure error.

- [ ] 237. Absolute Value Equations II
- [ ] 240. Absolute Value Equations III
- [ ] 243. Absolute Value Inequalities II

## Section 10: Graphs and Functions — the most important section `[GFX]` `[CMP]` `[DB]`
Functions, domain/range, composition are the vocabulary of everything. Special attention: graph
transformations (289–303) — first taste of the transform thinking 2D graphics is built on;
**the difference quotient (299)** — this IS the derivative, one limit short; composition (301)
`[CMP]`.

- [ ] 247. The Rectangular Coordinate System
- [ ] 249. Distance between Two Points (Distance Formula)
- [ ] 251. Determining Whether Three Points are the Vertices of a Right Triangle
- [ ] 253. Determining Whether Three Points are Collinear
- [ ] 255. The Midpoint Formula
- [ ] 257. Plotting Complex Numbers
- [ ] 259. Absolute Value of a Complex Number
- [ ] 261. Distance and Midpoint Formulas in the Complex Plane
- [ ] 263. Graphing Circles
- [ ] 266. Relations and Functions
- [ ] 268. Domain and Range
- [ ] 270. Vertical Line Test
- [ ] 272. Function Notation
- [ ] 274. Graphing Linear Functions
- [ ] 276. Finding the Slope of a Line
- [ ] 278. Equations of Lines
- [ ] 280. Parallel and Perpendicular Lines
- [ ] 283. Graphs of Basic Functions
- [ ] 285. Piecewise Functions
- [ ] 287. Increasing, Decreasing, and Constant Functions
- [ ] 289. Graphing Techniques: Shrinking and Stretching
- [ ] 291. Graphing Techniques: Reflecting
- [ ] 293. Even and Odd Functions
- [ ] 295. Graphing Techniques: Translations
- [ ] 297. Operations on Functions
- [ ] 299. The Difference Quotient
- [ ] 301. Composition of Functions
- [ ] 303. Combination of Graphing Transformations

## Section 11: Polynomial and Rational Functions `[GFX]` `[CRY]`
Curve behavior, zeros, factorization. The Remainder and Factor theorems return over finite
fields in cryptography and coding theory; asymptote reasoning returns in calculus limits.

- [ ] 306. Finding the Vertex Form of a Parabola
- [ ] 308. Graphing Parabolas
- [ ] 310. The Remainder Theorem
- [ ] 312. The Factor Theorem
- [ ] 314. The Rational Zeros Theorem
- [ ] 316. Fundamental Theorem of Algebra / Writing Polynomial Functions
- [ ] 318. Conjugate Zeros Theorem
- [ ] 320. Descartes' Rule of Signs
- [ ] 322. Intermediate Value and Boundedness Theorems
- [ ] 324. Finding the Zeros of a Polynomial Function
- [ ] 326. Polynomial Functions and Their Graphs
- [ ] 328. Solving Polynomial Inequalities
- [ ] 330. Graphing Rational Functions
- [ ] 332. Solving Variation Word Problems

## Section 12: Inverse, Exponential and Logarithmic Functions `[CRY]` `[NET]` `[DB]` `[CMP]`
The most goal-dense section of the course. Inverse functions carry the seed idea of crypto — a
function easy to compute, hard to invert; the discrete logarithm is this section's "solve for
the exponent," made deliberately hard. Logarithms are why B-trees and binary search are fast.

- [ ] 334. One-to-One Functions and the Horizontal Line Test
- [ ] 336. One-to-One Functions — Algebraic Method
- [ ] 338. How to Find the Inverse of a One-to-One Function
- [ ] 340. How to Determine if Two Functions are Inverses Using Composition
- [ ] 342. How to Find the Inverse of a Domain Restricted Function
- [ ] 344. Graphing Exponential Functions
- [ ] 346. How to Solve Exponential Equations with Like Bases
- [ ] 348. Compound Interest Word Problems
- [ ] 350. Continuous Compound Interest Word Problems
- [ ] 352. Introduction to Logarithms
- [ ] 354. Properties of Logarithms: Condensing and Expanding
- [ ] 356. Special Properties of Logarithms
- [ ] 358. Change of Base Formula
- [ ] 360. Solving Exponential Equations Using Logarithms
- [ ] 362. Solving Logarithmic Equations
- [ ] 364. Solving Exponential and Logarithmic Inequalities
- [ ] 366. Applications of Exponential and Logarithmic Functions

## Section 13: Systems of Equations and Matrices `[GFX]`
Gateway to Linear Algebra (module 04), where all of this is redone geometrically. Special
attention: **determinant as triangle area (413–418)** — here it looks like a trick; in Linear
Algebra §5 that geometric meaning becomes the *definition*.

- [ ] 369. Solving Linear Systems in Two Variables
- [ ] 371. Solving Linear Systems in Three Variables
- [ ] 373. Finding the Equation of a Parabola Given Three Points
- [ ] 375. Applications of Linear Systems
- [ ] 377. Partial Fraction Decomposition with Linear Factors
- [ ] 379. Partial Fraction Decomposition with Quadratic Factors
- [ ] 381. Gauss-Jordan Elimination: Two-Variable System
- [ ] 383. Gauss-Jordan Elimination: Three-Variable System
- [ ] 385. Gauss-Jordan Elimination: Four-Variable System
- [ ] 387. Introduction to Matrices
- [ ] 389. Adding and Subtracting Matrices
- [ ] 391. Multiplying a Matrix by a Scalar
- [ ] 393. Multiplying Matrices
- [ ] 395. Finding the Inverse of a Matrix
- [ ] 397. Solving Linear Systems using Inverse Matrices
- [ ] 399. Finding the Determinant of an n x n Matrix
- [ ] 401. Upper Triangular Form Determinant
- [ ] 403. Finding the Transpose of a Matrix
- [ ] 405. Finding the Adjoint of a Matrix
- [ ] 407. Finding the Inverse of a Matrix using the Determinant and Adjoint
- [ ] 409. Cramer's Rule: 2 x 2 System
- [ ] 411. Cramer's Rule: 3 x 3 System
- [ ] 413. Determinants — Area of a Triangle
- [ ] 415. Determinants — Test for Collinearity
- [ ] 417. Determinants — Equation of a Line

## Section 14: Conic Sections and Nonlinear Systems `[GFX]`
Ellipses, parabolas, hyperbolas: trajectories and collision shapes.

- [ ] 419. Graphing Sideways Parabolas
- [ ] 421. Conic Sections: The Parabola
- [ ] 423. Conic Sections: The Ellipse
- [ ] 425. Conic Sections: The Hyperbola
- [ ] 427. Solving Nonlinear Systems of Equations
- [ ] 429. Solving Systems of Linear and Nonlinear Inequalities

## Section 15: Sequences and Series `[CMP]` `[DB]` `[NET]`
Recurrences are algorithm analysis; geometric series are retransmission expectations and
amortized costs. Grows into Calculus §11.

- [ ] 431. Introduction to Sequences
- [ ] 433. Introduction to Series
- [ ] 435. Arithmetic Sequences and Series
- [ ] 437. Geometric Sequences and Series

## Section 16: Binomial Theorem, Counting, Probability, Sets `[DB]` `[CRY]` `[NET]`
Sets are the raw material of the relational model; counting is the substrate of probability.
A trailer for Rosen Ch. 2 & 6 and the entire Probability module.

- [ ] 439. The Binomial Theorem
- [ ] 441. Counting Theory
- [ ] 443. The Basics of Probability
- [ ] 445. Sets Part 1
- [ ] 447. Sets Part 2

## Section 17: Trigonometric Functions `[GFX]`
The mathematics of rotation begins here.

- [ ] 449. Angles in Trigonometry
- [ ] 451. Degrees, Minutes, and Seconds
- [ ] 453. Coterminal Angles
- [ ] 455. Angle Relationships and Similar Triangles
- [ ] 457. Finding Trigonometric Function Values of an Angle
- [ ] 459. Reciprocal Identities and finding Missing Trigonometric Function Values
- [ ] 461. Pythagorean and Quotient Identities

## Section 18: Acute Angles and Right Triangles `[GFX]`

- [ ] 464. Trigonometric Functions of Acute Angles
- [ ] 466. Trigonometric Functions of Non-Acute Angles
- [ ] 468. Finding Trigonometric Function Values Using a Calculator
- [ ] 470. Solving Right Triangles

## Section 19: Radian Measure and the Unit Circle `[GFX]`
Radians are the *only* angle unit calculus and every graphics API use — make them native.
Special attention: **linear & angular speed (479)** — rotational kinematics, verbatim from a
physics engine.

- [ ] 473. Radian Measure
- [ ] 475. Finding the Arc Length on a Circle
- [ ] 477. The Unit Circle: Circular Functions
- [ ] 479. Linear and Angular Speed

## Section 20: Graphing Trigonometric Functions `[GFX]` `[NET]`
Sinusoids are oscillation and waves: springs, orbits, camera bob, signals. Amplitude, period,
phase are knobs you will turn in code.

- [ ] 482. Graphing Sine and Cosine
- [ ] 484. Graphing Tangent and Cotangent
- [ ] 486. Graphing Secant and Cosecant

## Section 21: Trigonometric Identities `[GFX]`
Special attention: **sum & difference identities (493–495)** — they *are* the 2D rotation
formulas; a rotation matrix is these two identities written in a box.

- [ ] 489. Fundamental Identities
- [ ] 491. Verifying Trigonometric Identities
- [ ] 493. Sum and Difference Identities for Cosine
- [ ] 495. Sum and Difference Identities for Sine and Tangent
- [ ] 497. Double-Angle, Product-to-Sum, Sum-to-Product Identities
- [ ] 499. Half-Angle Identities

## Section 22: Trig Equations and Inverse Trig Functions `[GFX]`
Angle recovery: `atan2` in every game codebase is inverse tangent plus the quadrant care this
section teaches.

- [ ] 502. Inverse Circular Functions
- [ ] 504. Solving Trigonometric Equations Using Linear Methods
- [ ] 506. Solving Trigonometric Equations Using Factoring
- [ ] 508. Solving Trigonometric Equations Using Square Roots, Squaring, Identities
- [ ] 510. Solving Trigonometric Equations with Half-Angles
- [ ] 512. Solving Trigonometric Equations with Multiple Angles
- [ ] 514. Solving Trigonometric Equations Using Inverse Trigonometric Functions
- [ ] 516. Solving Trigonometric Inequalities

## Section 23: Law of Sines and Cosines `[GFX]`
Triangle solving = distance and angle queries in geometry code; Heron's formula computes
triangle areas from side lengths.

- [ ] 519. Law of Sines Part 1: (SAA) and (ASA)
- [ ] 521. Law of Sines Part 2: (SSA) Ambiguous Case
- [ ] 523. Law of Cosines: (SAS) and (SSS), Heron's Formula

## Section 24: Vectors `[GFX]`
**The single most directly applicable section of the whole course.** The daily vocabulary of 2D
graphics and physics; Linear Algebra §1 picks up exactly here.

- [ ] 526. Vectors Part 1: Component Form, Magnitude, Direction Angle
- [ ] 528. Vectors Part 2: Adding Vectors, Multiplying by a Scalar
- [ ] 530. Vectors Part 3: The Dot Product and the Angle Between Vectors

## Section 25: Complex Numbers in Trigonometry `[GFX]`
Complex multiplication is rotate-and-scale — a 2D rotation engine hiding in arithmetic. Roots of
unity return in the FFT.

- [ ] 533. Review of Complex Numbers
- [ ] 535. Polar Form of Complex Numbers
- [ ] 537. Product and Quotient Theorems
- [ ] 539. De Moivre's Theorem: Powers and Roots of Complex Numbers

## Section 26: Polar/Parametric Equations and Analytic Geometry `[GFX]`
Parametric curves are motion paths — (x(t), y(t)) is how every moving object is described.
Calculus §12 does calculus on these.

- [ ] 542. Polar Coordinates
- [ ] 544. Distance Between Polar Points
- [ ] 546. Polar Form of a Line
- [ ] 548. Circles in Polar Form
- [ ] 550. Analytic Geometry: Lines

---

## Sections that get sequels

| Greene section | Grows into |
|---|---|
| §13 Matrices & determinants | Linear Algebra (module 04), done geometrically |
| §15 Sequences & series | Calculus §11 (infinite series, Taylor) |
| §16 Counting & probability | Rosen Ch. 6 + Probability (module 05) |
| §10 Difference quotient | Calculus §3 (the derivative) |
| §24 Vectors | Linear Algebra §1, Calculus §13 |
| §25 Complex numbers | Linear Algebra §10, Calculus §10 (oscillation) |
