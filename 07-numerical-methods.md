# Module 07 — Numerical Methods (Bridging Module)

<!-- progress:07-numerical-methods -->
![07 · Numerical Methods — Sastry progress](progress/07-numerical-methods.svg?v=2-0-18)

**0 / 18 lessons complete · 0%** — 0 of 5 sections finished · [completion log](LOG.md)
<!-- endprogress -->

**Book:** S.S. Sastry, *Introductory Methods of Numerical Analysis* (PHI).
**Prerequisites:** Calculus §§2–5 and §9; Linear Algebra §2. Best taken last, alongside your
first real simulation code.
**Why this module exists:** this is where the math meets the machine. Everything before this
module assumes exact real numbers; computers don't have them. Short module, outsized payoff —
most "physics bugs" in games are lessons 7–12 of this file.

**The rhythm applies to every lesson**, with one addition: every lesson here should also be
*programmed* — a 10–30 line implementation counts as part of the practice test.

Tags: `[GFX]` graphics/simulation · `[CMP]` compilers · `[NET]` networking · `[DB]` databases ·
`[CRY]` cryptography.

---

## Section 1: How Computers Do Real Numbers

- [ ] **1. Floating-point numbers (IEEE 754)** `[GFX]` `[CMP]` `[DB]` — what a `float`/`double` can
   and cannot represent; why 0.1 + 0.2 ≠ 0.3; the number system all your simulations actually
   run in.
- [ ] **2. Rounding error, machine epsilon, catastrophic cancellation** `[GFX]` — how error enters,
   accumulates, and occasionally detonates (subtracting nearly equal numbers); why physics
   objects jitter at rest.
- [ ] **3. Conditioning and stability** — sensitive problems vs. unstable algorithms: two different
   diseases with different cures.

## Section 2: Solving Equations Numerically

- [ ] **4. The bisection method** — the Intermediate Value Theorem (Calculus lesson 10) turned into
   an unbreakable, slow algorithm.
- [ ] **5. Newton's method, revisited** `[GFX]` — Calculus lesson 31 as production code: quadratic
   convergence, and the ways it fails (bad starts, flat spots); the core of constraint solvers.
- [ ] **6. Fixed-point iteration** — solving x = g(x) by repetition, and when repetition converges
   (the eigenvalue/derivative condition — Linear Algebra §9 echo).

## Section 3: Simulating Motion — Numerical ODEs `[GFX]` core

- [ ] **7. Euler's method and its error** `[GFX]` — Calculus lesson 59 measured honestly: first-order
   accuracy and what that costs.
- [ ] **8. Why Euler explodes: stability and the timestep** `[GFX]` — energy gain, the stability
   region, and the classic bug: springs that fling objects into orbit when the frame rate drops.
- [ ] **9. Semi-implicit (symplectic) Euler** `[GFX]` — one-line change, dramatically better energy
   behavior: the default integrator of game physics engines.
- [ ] **10. Verlet integration** `[GFX]` — position-based integration: cloth, ropes, and particle
    systems; why it pairs beautifully with constraints.
- [ ] **11. Runge–Kutta 4** `[GFX]` — the accuracy workhorse: four slope samples per step; when RK4
    over symplectic Euler and vice versa.
- [ ] **12. Choosing an integrator; fixed vs. variable timestep** `[GFX]` — accuracy vs. stability
    vs. cost; why "fix your timestep" is the most-quoted game-physics advice ever written.

## Section 4: Curves and Interpolation

- [ ] **13. Polynomial interpolation and its dangers** — one polynomial through n points; the Runge
    phenomenon, or why high-degree fits wiggle.
- [ ] **14. Splines** `[GFX]` — piecewise low-degree polynomials: smooth paths without the wiggle.
- [ ] **15. Bézier curves and de Casteljau's algorithm** `[GFX]` — the curves under every font,
    vector graphic, and animation-easing function; evaluated by repeated linear interpolation.

## Section 5: Odds and Ends

- [ ] **16. Numerical integration in practice** — trapezoid and Simpson (Calculus lesson 55) as
    code; adaptive step sizing in one idea.
- [ ] **17. Solving linear systems numerically** `[GFX]` — why naive Gaussian elimination fails on a
    computer, and how pivoting fixes it (Linear Algebra §2 meets lesson 2 of this module).
- [ ] **18. Random number generation** `[CRY]` `[GFX]` — from linear congruential generators to
    cryptographic RNGs: why the fast ones are predictable, why the safe ones are slower, and why
    games and cryptography must never swap theirs.

---

## Exit criteria

You are done with this module when you can, cold:

- Explain why 0.1 + 0.2 ≠ 0.3 in binary floating point, to another programmer, in two minutes.
- Implement Euler, semi-implicit Euler, and RK4 for a spring, plot the three energies over
  time, and explain the plot.
- State what happens to each integrator when the timestep doubles, before running it.
- Evaluate a cubic Bézier at t = 0.5 by hand with de Casteljau's algorithm.
- Say which random number generator you'd use for particle effects vs. for key generation, and
  defend the choice.
