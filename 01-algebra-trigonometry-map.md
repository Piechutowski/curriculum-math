# Module 01 — Study Map: Mr. Greene's College Math Course

**Resource:** Mr. Greene's course (you are working through it now).
**How to use this file:** you already have the complete lesson list — this map tells you what each
section is *for* in terms of your five goals, and what to pay special attention to. Do the whole
course in order, keeping the Lesson → Study Notes → Practice Test rhythm. Nothing is skippable;
tagged sections are where the payoff is direct.

Tags: `[GFX]` graphics/simulation · `[CMP]` compilers · `[NET]` networking · `[DB]` databases ·
`[CRY]` cryptography. No tag = foundation that later lessons stand on.

---

## Section 1: Review of Basic Concepts
Goal here is *automaticity* — this must become effortless. Two ideas quietly return later:
interval notation (reappears in calculus limits and continuity) and absolute value as *distance*
(becomes the ε–δ definition of limits and every error bound in numerical work).

## Section 2: Review of Polynomials `[CMP]` `[CRY]` `[NET]`
Polynomial arithmetic returns in heavy disguise: polynomial rings over finite fields are how CRC
checksums `[NET]` and error-correcting codes work (Rosen 12.8), and polynomial factoring
intuition carries into crypto. Special attention: **synthetic division (lesson 55)** — it is
Horner's method, the standard way real programs evaluate polynomials.

## Section 3: Review of Rational Expressions
Feeds rational functions (Section 11) and the algebraic manipulation that computing limits in
calculus demands. Restricted values are your first taste of "domain reasoning."

## Section 4: Review of Radical Expressions
Rational exponents are used constantly in calculus differentiation. Fluency, not depth, is the
goal.

## Section 5: Solving Linear Equations and Inequalities
Foundation. Compound inequalities and interval reasoning return in convergence arguments and
error bounds.

## Section 6: Applications of Linear Equations `[GFX]`
Do not skim the word problems — translating a described situation into equations is *modeling*,
the core skill of simulation. The motion problems (lesson 168) are physics with the calculus
removed; you will meet them again with the calculus put back in.

## Section 7: Quadratic Equations and Inequalities `[GFX]` `[CRY]`
Quadratics are projectile trajectories `[GFX]`. The discriminant will later decide whether a
damped spring oscillates or settles (Calculus §10) — same algebra, new meaning. Complex numbers
(lessons 184–190) become 2D rotations in Section 25 and in Linear Algebra.

## Section 8: Radical and Rational Equations
Manipulation fluency; the equation-solving muscle the rest of the course assumes.

## Section 9: Advanced Absolute Value
Absolute value = distance on the line. This exact idea, one level up, is how calculus defines
limits and how numerical methods measure error.

## Section 10: Graphs and Functions — the most important section of the course `[GFX]` `[CMP]` `[DB]`
Functions, domain/range, and composition are the vocabulary of *everything*: program semantics
`[CMP]`, mappings and keys `[DB]`, transformations `[GFX]`. Special attention to three lessons:
graph transformations (289–303) — this is your first taste of the transform-composition thinking
that 2D graphics is built on; **the difference quotient (lesson 299)** — this IS the derivative,
one limit short; and composition of functions (301) `[CMP]`.

## Section 11: Polynomial and Rational Functions `[GFX]` `[CRY]`
Curve behavior, zeros, and factorization. The Remainder and Factor theorems return over finite
fields in cryptography and coding theory. Asymptote reasoning returns in calculus limits.

## Section 12: Inverse Functions, Exponential and Logarithmic Functions `[CRY]` `[NET]` `[DB]` `[CMP]`
The most goal-dense section in the course. Inverse functions carry the seed idea of cryptography:
a function easy to compute but hard to invert. The discrete logarithm — the hard problem under
Diffie–Hellman — is this section's "solve for the exponent," made deliberately difficult.
Logarithms are also *why* B-trees and binary search are fast `[DB]` and the language of
complexity growth `[CMP]` `[NET]`.

## Section 13: Systems of Equations and Matrices `[GFX]`
Your gateway to Linear Algebra (module 04), where all of this is redone geometrically. Special
attention: **determinant as triangle area (lessons 413–418)** — in this course it looks like a
trick; in Linear Algebra §5 that geometric meaning becomes the *definition*.

## Section 14: Conic Sections and Nonlinear Systems `[GFX]`
Ellipses, parabolas, hyperbolas: trajectories and collision shapes. Parabola vertex form is
projectile math.

## Section 15: Sequences and Series `[CMP]` `[DB]` `[NET]`
Recurrences are how algorithm costs are analyzed; geometric series are retransmission
expectations `[NET]` and amortized costs `[DB]`. Grows into Calculus §11 (infinite series).

## Section 16: Binomial Theorem, Counting, Probability, Sets `[DB]` `[CRY]` `[NET]`
Sets are the raw material of the relational model `[DB]`; counting is the substrate of all of
probability. This section is a trailer for Rosen Ch. 2 & 6 and the entire Probability module.

## Sections 17–19: Trig Functions, Right Triangles, Radians & the Unit Circle `[GFX]`
The mathematics of rotation begins here. Radians are the *only* angle unit calculus and every
graphics API use — make them native. The unit circle is the rotation machine. Special attention:
**linear & angular speed (lesson 479)** — that is rotational kinematics, verbatim from a physics
engine.

## Section 20: Graphing Trigonometric Functions `[GFX]` `[NET]`
Sinusoids are oscillation and waves: springs, orbits, camera bob `[GFX]`, and signals/modulation
`[NET]`. Amplitude, period, and phase shift are knobs you will turn in code.

## Section 21: Trigonometric Identities `[GFX]`
The algebra that keeps rotation math manageable. Special attention: **sum & difference identities
(493–495)** — they *are* the 2D rotation formulas; when you later see a rotation matrix, it is
these two identities written in a box.

## Section 22: Trig Equations and Inverse Trig Functions `[GFX]`
Angle recovery. `atan2`, which appears in essentially every game codebase, is inverse tangent
plus the quadrant care this section teaches.

## Section 23: Law of Sines and Cosines `[GFX]`
Triangle solving = distance and angle queries in geometry code. Heron's formula computes triangle
areas from side lengths — useful in collision and mesh work.

## Section 24: Vectors `[GFX]`
**The single most directly applicable section of the whole course.** Component form, magnitude,
direction angle, vector addition, scalar multiplication, dot product, angle between vectors —
this is the daily vocabulary of 2D graphics and physics. Linear Algebra §1 picks up exactly here.

## Section 25: Complex Numbers in Trigonometry `[GFX]`
Complex multiplication is rotate-and-scale — a full 2D rotation engine hiding in arithmetic.
De Moivre's theorem and roots of unity return later in the FFT (signal processing, fast
polynomial multiplication).

## Section 26: Polar/Parametric Equations and Analytic Geometry `[GFX]`
Parametric curves are motion paths — (x(t), y(t)) is how every moving object in a simulation is
described. Polar thinking powers radial effects and orbits. Calculus §12 does calculus on these.

---

## Sections that get sequels

| Greene section | Grows into |
|---|---|
| §13 Matrices & determinants | Linear Algebra (module 04), done geometrically |
| §15 Sequences & series | Calculus §11 (infinite series, Taylor) |
| §16 Counting & probability | Rosen Ch. 6 + Probability (module 05) |
| §10 Difference quotient | Calculus §3 (the derivative) |
| §24 Vectors | Linear Algebra §1, Calculus §13 |
| §25 Complex numbers | Linear Algebra §10 (rotation matrices), Calculus §10 (oscillation) |
