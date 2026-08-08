# Module 03 — Calculus

<!-- progress:03-calculus -->
![03 · Calculus — Thomas progress](progress/03-calculus.svg)

**0 / 89 lessons complete · 0%** — [completion log](LOG.md)
<!-- endprogress -->

**Book:** *Thomas' Calculus* (Pearson India edition). Section references are approximate across
editions — match by topic name, not number.
**University mapping:** this module is **Calculus 1 and Calculus 2, complete**, plus the opening
slice of Calculus 3 and the core of a first Differential Equations course. Precisely:
*differential calculus* = Sections 2–4 · *integral calculus* = Sections 5–8 · Calculus 2's
series/parametric material = Sections 11–12 · Differential Equations = Sections 9–10 ·
Calculus 3 (multivariable) = only its opening (Section 13); the rest of Calc 3 is deliberately
just an orientation lesson (89) because 2D work doesn't need it. If you use Shanti Narayan
instead of Thomas: his *Differential Calculus* volume covers Sections 2–4, his *Integral
Calculus* volume covers Sections 5–8.
**Prerequisite:** Greene course complete (especially §§10–12, 15, 17–26).
**Why this module exists:** physical simulation *is* calculus — velocity and acceleration are
derivatives, and simulating motion means solving differential equations. This is the biggest gap
in your current resources. The module follows the complete standard course; tags mark where each
skill lands in your goals.

**The rhythm applies to every lesson:** Lesson → Study Notes → Practice Test (work the book's
exercise set for that section until you can score ~90% cold).

Tags: `[GFX]` graphics/simulation · `[CMP]` compilers · `[NET]` networking · `[DB]` databases ·
`[CRY]` cryptography. No tag = foundation later lessons stand on.

---

## Section 1: Functions — the Bridge from Greene

- [ ] **1. Functions through the calculus lens** — domain, range, composition, and piecewise functions
   at reading speed; this is Greene §10 revisited as a warm-up, not new material. *(Thomas Ch. 1)*
- [ ] **2. Trigonometric functions for calculus** `[GFX]` — sine and cosine as functions of a real
   number (radians only, forever); amplitude, period, phase. *(Thomas Ch. 1)*
- [ ] **3. Exponentials, logarithms, and inverse functions** — the Greene §12 material restated the
   way calculus will use it. *(Thomas Ch. 1)*

## Section 2: Limits and Continuity

- [ ] **4. The idea of a limit: average vs. instantaneous rate of change** `[GFX]` — the question
   "how fast, right now?" that the whole subject exists to answer; velocity as the limit of
   average speeds. *(Thomas 2.1)*
- [ ] **5. Limit of a function and the limit laws** — evaluating limits with the algebra of limits.
   *(Thomas 2.2)*
- [ ] **6. One-sided limits** — approaching from the left and right; when the two disagree.
   *(Thomas 2.4)*
- [ ] **7. Computing limits algebraically** — factoring, rationalizing, and canceling: the Greene
   §§2–4 skills earning their keep. *(Thomas 2.2)*
- [ ] **8. Limits involving infinity; asymptotes** — end behavior, made precise; Greene §11's
   asymptotes get their real definition. *(Thomas 2.6)*
- [ ] **9. Continuity** — what it means, types of discontinuity, and why continuous functions are the
   ones you can trust. *(Thomas 2.5)*
- [ ] **10. The Intermediate Value Theorem** — a continuous function can't skip values; the theorem
    that makes root-finding algorithms work (Numerical Methods §2). *(Thomas 2.5)*
- [ ] **11. The precise definition of a limit (ε–δ)** — Greene's "absolute value as distance"
    becomes the rigorous backbone of the subject. Hard; worth it; do not skip. *(Thomas 2.3)*

## Section 3: Derivatives

- [ ] **12. The derivative at a point** — the difference quotient (Greene lesson 299) with its limit
    finally taken. *(Thomas 3.1)*
- [ ] **13. The derivative as a function; differentiability** — where derivatives fail: corners,
    cusps, vertical tangents. *(Thomas 3.2)*
- [ ] **14. Differentiation rules I: power, constant multiple, sum** — the mechanics begin.
    *(Thomas 3.3)*
- [ ] **15. Differentiation rules II: product and quotient rules** *(Thomas 3.3)*
- [ ] **16. Derivatives of trigonometric functions** `[GFX]` — why d/dx sin = cos, and what that says
    about oscillation. *(Thomas 3.5)*
- [ ] **17. The chain rule** — differentiating compositions; the most-used rule in all of applied
    mathematics. *(Thomas 3.6)*
- [ ] **18. Implicit differentiation** — differentiating curves that aren't function graphs (circles,
    conics — Greene §14). *(Thomas 3.7)*
- [ ] **19. Derivatives of exponentials and logarithms; logarithmic differentiation** `[NET]` — eˣ as
    the function that is its own derivative: the reason exponentials model growth and decay.
    *(Thomas 3.8)*
- [ ] **20. Derivatives of inverse trigonometric functions** `[GFX]` — differentiating the
    angle-recovery functions. *(Thomas 3.9)*
- [ ] **21. The derivative as rate of change: velocity and acceleration** `[GFX]` — position →
    velocity → acceleration by differentiation; the kinematics dictionary. *(Thomas 3.4)*
- [ ] **22. Related rates** — when two changing quantities are linked, relate their rates; classic
    modeling practice. *(Thomas 3.10)*
- [ ] **23. Linearization and differentials** `[GFX]` — approximating a function by its tangent line;
    the idea behind every "small step" a simulation takes. *(Thomas 3.11)*

## Section 4: Applications of Derivatives

- [ ] **24. Absolute and local extrema; critical points** — where functions peak and bottom out.
    *(Thomas 4.1)*
- [ ] **25. The Mean Value Theorem** — the honest average-vs-instantaneous theorem underneath
    everything in this section. *(Thomas 4.2)*
- [ ] **26. Monotonicity and the first-derivative test** *(Thomas 4.3)*
- [ ] **27. Concavity, inflection points, and the second-derivative test** *(Thomas 4.4)*
- [ ] **28. Curve sketching: the full function autopsy** — unify Greene §11 with everything since
    lesson 12; sketch anything from its formula. *(Thomas 4.4)*
- [ ] **29. Indeterminate forms and L'Hôpital's rule** — resolving 0/0 and ∞/∞. *(Thomas 4.5)*
- [ ] **30. Applied optimization** — best/cheapest/largest problems; the calculus behind every tuning
    knob. *(Thomas 4.6)*
- [ ] **31. Newton's method** `[GFX]` — using tangent lines to solve equations numerically; your first
    real algorithm from calculus (returns in Numerical Methods §2). *(Thomas 4.7)*
- [ ] **32. Antiderivatives** — differentiation run backwards; setting the stage for integration.
    *(Thomas 4.8)*

## Section 5: Integration

- [ ] **33. Estimating with finite sums; Riemann sums** — area under a curve as a limit of rectangles.
    *(Thomas 5.1)*
- [ ] **34. Sigma notation and limits of sums** — Greene §15's notation doing real work.
    *(Thomas 5.2)*
- [ ] **35. The definite integral and its properties** *(Thomas 5.3)*
- [ ] **36. The Fundamental Theorem of Calculus, Part 1** — accumulation functions and their
    derivatives. *(Thomas 5.4)*
- [ ] **37. The Fundamental Theorem of Calculus, Part 2** — derivatives and integrals are inverse
    operations: the payoff of the course so far. *(Thomas 5.4)*
- [ ] **38. Indefinite integrals and the substitution method** — the chain rule in reverse.
    *(Thomas 5.5–5.6)*
- [ ] **39. Substitution in definite integrals; area between curves** *(Thomas 5.6)*

## Section 6: Applications of Integration

- [ ] **40. Volumes by slicing and rotation** — integration as "add up the pieces"; geometry muscle.
    *(Thomas 6.1–6.2)*
- [ ] **41. Arc length** `[GFX]` — the true length of a curve; needed whenever an object must move
    along a path at constant speed. *(Thomas 6.3)*
- [ ] **42. Areas of surfaces of revolution** — more setup practice. *(Thomas 6.4)*
- [ ] **43. Work and energy** `[GFX]` — force integrated over distance: springs (Hooke's law),
    lifting, compression; the physics-engine quantities. *(Thomas 6.5)*
- [ ] **44. Moments and centers of mass** `[GFX]` — where an object balances; the center of mass is
    the point your rigid-body simulation moves. *(Thomas 6.6)*
- [ ] **45. Fluid forces** — pressure integrated over area; final round of set-up-the-integral
    practice. *(Thomas 6.5)*

## Section 7: Transcendental Functions and Growth

- [ ] **46. The natural logarithm defined by an integral** — the honest construction of ln x, and eˣ
    as its inverse. *(Thomas 7.1–7.2)*
- [ ] **47. Exponential change: growth and decay** `[NET]` `[GFX]` — the differential equation
    y′ = ky and its solution; half-life, cooling, damping, exponential backoff. *(Thomas 7.2)*
- [ ] **48. Relative rates of growth** `[CMP]` `[DB]` `[NET]` — proving log ≪ polynomial ≪
    exponential; Big-O from Rosen Ch. 3, now with proofs. *(Thomas 7.4)*
- [ ] **49. Hyperbolic functions** `[GFX]` — sinh and cosh: hanging chains, smooth ramps, and the
    trig of certain damping problems. *(Thomas 7.3)*

## Section 8: Techniques of Integration

- [ ] **50. Integration by parts** — the product rule in reverse. *(Thomas 8.2)*
- [ ] **51. Trigonometric integrals** *(Thomas 8.3)*
- [ ] **52. Trigonometric substitution** — integrals with √(a²−x²): circles entering through algebra.
    *(Thomas 8.4)*
- [ ] **53. Integration by partial fractions** — Greene lessons 377–379 were preparation for exactly
    this. *(Thomas 8.5)*
- [ ] **54. Integration strategy** — choosing the technique; recognizing when no elementary
    antiderivative exists. *(Thomas 8.6)*
- [ ] **55. Numerical integration: trapezoidal and Simpson's rules** `[GFX]` — when formulas fail,
    compute; first contact with Numerical Methods. *(Thomas 8.7)*
- [ ] **56. Improper integrals** `[NET]` `[CRY]` — integrals to infinity; the gateway to continuous
    probability (Probability §5 depends on this lesson). *(Thomas 8.8)*

## Section 9: First-Order Differential Equations — the Physics Gateway

- [ ] **57. What a differential equation is** `[GFX]` — equations whose unknowns are functions;
    solutions, initial conditions, and why physics is written in ODEs. *(Thomas 9.1)*
- [ ] **58. Slope fields** `[GFX]` — seeing every solution of an ODE at once, before solving
    anything. *(Thomas 9.1)*
- [ ] **59. Euler's method** `[GFX]` — follow the slope in small steps: your first physics
    integrator, and the beating heart of every game loop. Do this one by hand, repeatedly.
    *(Thomas 9.1)*
- [ ] **60. Separable equations** *(Thomas 9.1)*
- [ ] **61. First-order linear equations; integrating factors** *(Thomas 9.2)*
- [ ] **62. Exponential models revisited** `[GFX]` `[NET]` — mixing, Newton's law of cooling,
    terminal velocity with drag: real modeling, start to finish. *(Thomas 9.3)*
- [ ] **63. Autonomous equations, equilibria, and stability** `[GFX]` — where systems settle and why;
    the concepts behind "why does my simulation explode or freeze?" *(Thomas 9.4)*
- [ ] **64. Systems of ODEs and the phase plane** `[GFX]` — two coupled equations, pictured as a flow
    in the plane; predator–prey as the classic. *(Thomas 9.5)*

## Section 10: Second-Order ODEs and Oscillation

*(In some Thomas editions this is a later/online chapter; any ODE text covers it identically.)*

- [ ] **65. Homogeneous linear second-order equations with constant coefficients** — the equation
    class that all of mechanical oscillation lives in. 
- [ ] **66. The characteristic equation: real, repeated, and complex roots** — Greene's quadratic
    formula and complex numbers (§7, §25) cash in completely.
- [ ] **67. The harmonic oscillator** `[GFX]` — springs and pendulums: x″ = −ω²x and why the answer
    is sine; the equation under every soft-body and spring system you will ever code.
- [ ] **68. Damped oscillation** `[GFX]` — under-, over-, and critical damping, decided by the
    discriminant; critically-damped springs are how polished games smooth cameras and UI.
- [ ] **69. Forced oscillation and resonance** `[GFX]` `[NET]` — driving frequencies, resonance
    peaks; where mechanical intuition meets signals.

## Section 11: Infinite Sequences and Series

- [ ] **70. Sequences and convergence** — Greene §15 grown up: limits of sequences, monotone
    convergence. *(Thomas 10.1)*
- [ ] **71. Infinite series; geometric and telescoping series** `[NET]` — summing forever; the
    geometric series computes expected retransmissions and amortized costs. *(Thomas 10.2)*
- [ ] **72. The integral test; p-series** *(Thomas 10.3)*
- [ ] **73. Comparison tests** *(Thomas 10.4)*
- [ ] **74. Ratio and root tests** *(Thomas 10.5)*
- [ ] **75. Alternating series; absolute vs. conditional convergence** *(Thomas 10.6)*
- [ ] **76. Power series and radius of convergence** *(Thomas 10.7)*
- [ ] **77. Taylor and Maclaurin series** — functions rewritten as infinite polynomials; the deepest
    idea in the module. *(Thomas 10.8)*
- [ ] **78. Taylor polynomials and error bounds** `[GFX]` — cut the series off and bound the damage:
    this is literally how your CPU computes sin(x), and how integrator error is analyzed.
    *(Thomas 10.9)*
- [ ] **79. Working with power series** — substitute, differentiate, and integrate series to
    manufacture new ones. *(Thomas 10.10)*

## Section 12: Parametric Curves and Polar Calculus

- [ ] **80. Parametric curves** `[GFX]` — motion as (x(t), y(t)); Greene §26 becomes the standard
    description of every moving object. *(Thomas 11.1)*
- [ ] **81. Calculus of parametric curves** `[GFX]` — tangents, speed, and arc length along a path.
    *(Thomas 11.2)*
- [ ] **82. Calculus in polar coordinates** — slopes and areas in r and θ. *(Thomas 11.4–11.5)*

## Section 13: Vectors and Motion

- [ ] **83. Vectors, calculus edition** `[GFX]` — Greene §24 restated in calculus notation; bridge to
    Linear Algebra §1. *(Thomas 12.1–12.3)*
- [ ] **84. Vector-valued functions** `[GFX]` — curves as vector functions r(t); limits and
    derivatives componentwise. *(Thomas 13.1)*
- [ ] **85. Velocity and acceleration vectors; projectile motion** `[GFX]` — differentiate the
    position vector and you have the whole kinematics of a 2D game. *(Thomas 13.2)*
- [ ] **86. Arc length, the unit tangent, and curvature** `[GFX]` — how sharply a path bends;
    steering behaviors and path smoothing. *(Thomas 13.3–13.4)*
- [ ] **87. Functions of several variables; partial derivatives** `[GFX]` — f(x, y) and its rates of
    change in each direction. *(Thomas 14.1–14.3)*
- [ ] **88. The gradient** `[GFX]` — the direction of steepest change; potential fields, force from
    energy, and the mathematical heart of collision resolution (and, someday, machine learning).
    *(Thomas 14.5)*
- [ ] **89. Orientation: what lies beyond** — multiple integrals, line integrals, vector fields:
    what they are, which goals would ever need them (3D graphics, fluids), and why it's safe to
    stop here for now. *(Thomas Ch. 15–16, survey only)*

---

## Exit criteria

You are done with this module when you can, cold:

- Differentiate and integrate any standard function combination without hesitation.
- State and use both parts of the Fundamental Theorem of Calculus.
- Set up and solve an applied optimization problem from a word description.
- Solve separable and first-order linear ODEs by hand — and simulate *any* first-order ODE for
  three steps of Euler's method with pencil and paper.
- Derive the motion of a damped spring and explain how the discriminant of the characteristic
  equation decides whether it oscillates, creeps, or snaps back.
- Write the Taylor polynomial of sin x to 5 terms and bound its error on [−π/2, π/2].
