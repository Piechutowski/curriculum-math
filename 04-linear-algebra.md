# Module 04 — Linear Algebra (Geometric)

<!-- progress:04-linear-algebra -->
![04 · Linear Algebra — Kumaresan progress](progress/04-linear-algebra.svg?v=0-52)

**0 / 52 lessons complete · 0%** — [completion log](LOG.md)
<!-- endprogress -->

**Book:** S. Kumaresan, *Linear Algebra: A Geometric Approach* (PHI). Alternative that pairs
well with video: Gilbert Strang, *Introduction to Linear Algebra* + the free MIT OCW 18.06
lectures. Lessons below reference topics, not section numbers — both books cover them; find by
name.
**Prerequisites:** Greene §13 (matrices), §24 (vectors). Independent of Calculus — run the two
modules in parallel.
**Why this module exists:** Greene teaches matrices as equation-solving machinery. Graphics
needs the other view — matrices as *transformations of space* — and that view is this module's
spine.

**The rhythm applies to every lesson:** Lesson → Study Notes → Practice Test. For this subject,
practice tests should always include a *drawing*: sketch what the transformation does to the
unit square.

Tags: `[GFX]` graphics/simulation · `[CMP]` compilers · `[NET]` networking · `[DB]` databases ·
`[CRY]` cryptography. No tag = foundation later lessons stand on.

---

## Section 1: Vector Geometry

- [ ] **1. Vectors as arrows and as coordinates** `[GFX]` — addition, scalar multiplication, and the
   two pictures (tip-to-tail, components) that must become one picture. Greene §24, deepened.
- [ ] **2. Length and the dot product; the angle between vectors** `[GFX]` — the dot product as
   "how aligned are these?"; the single most-used operation in graphics code.
- [ ] **3. Orthogonality and projection** `[GFX]` — projecting one vector onto another: sliding along
   walls, shadows on axes, decomposing forces.
- [ ] **4. The cross product and the 2D perp-dot product** `[GFX]` — "which side of the line?" —
   winding order, turn direction, and torque all come from this one scalar.
- [ ] **5. Lines and planes in vector form** `[GFX]` — parametric lines p + t·d are rays; the
   foundation of every raycast.
- [ ] **6. Distances** `[GFX]` — point-to-line and point-to-plane distance: the geometry of collision
   detection.

## Section 2: Solving Linear Systems

- [ ] **7. Systems as geometry** — two lines meet in a point (or don't): what solving *means* before
   how.
- [ ] **8. Gaussian elimination and row echelon form** — Greene §13's Gauss–Jordan, now as a
   systematic algorithm you could program.
- [ ] **9. Solution sets: none, one, infinitely many** — and what each case looks like geometrically.
- [ ] **10. The matrix–vector product Ax** — as a combination of A's columns: the reading that makes
    everything later obvious.
- [ ] **11. Applications: fitting and balancing** `[GFX]` `[NET]` — fit a curve through given points
    (Greene lesson 373 generalized); balance flows in a network.

## Section 3: Matrices as Transformations — the heart of the module

- [ ] **12. A matrix is a function** `[GFX]` — matrices as linear transformations of the plane; a
    2×2 matrix is fully known by where it sends the two basis vectors.
- [ ] **13. The transformation zoo** `[GFX]` — rotation, scaling, reflection, shear: recognize each
    matrix on sight and sketch its effect on the unit square. (The rotation matrix is Greene's
    sum identities, lessons 493–495, written in a box.)
- [ ] **14. Composition = matrix multiplication** `[GFX]` — why matrix multiplication is defined the
    way it is, and why order matters: rotate-then-translate ≠ translate-then-rotate.
- [ ] **15. Matrix algebra** — multiplication properties, transpose, and the algebra rules that do
    and don't carry over from numbers.
- [ ] **16. Inverse matrices** `[GFX]` — undoing a transformation; singular matrices as
    transformations that destroy information (flatten the plane).
- [ ] **17. Kernel and image, geometrically** `[CMP]` — what a transformation flattens to zero and
    what it can reach.

## Section 4: Affine Transformations and Homogeneous Coordinates

- [ ] **18. Translation is not linear** `[GFX]` — the problem: sliding the plane moves the origin, so
    no matrix can do it. The stage is set.
- [ ] **19. Homogeneous coordinates** `[GFX]` — add a third coordinate and 3×3 matrices do
    rotation + scaling + translation in one object: THE tool of 2D graphics.
- [ ] **20. The 2D transform pipeline** `[GFX]` — local → world → camera/view → screen as a product
    of matrices; the architecture of every renderer.
- [ ] **21. Transform hierarchies** `[GFX]` — parent–child composition: a turret on a tank on a
    moving map; scene graphs and articulated bodies.
- [ ] **22. Inverting affine transforms** `[GFX]` — screen-to-world: how mouse picking works.

## Section 5: Determinants

- [ ] **23. The determinant as signed area** `[GFX]` — Greene lessons 413–418 promoted from trick to
    definition; 3×3 as volume.
- [ ] **24. Orientation** `[GFX]` — what a negative determinant means: reflections, flipped winding
    order, inside-out triangles.
- [ ] **25. Properties and computation** — determinants via elimination; why the cofactor formula is
    for theory and elimination is for practice.
- [ ] **26. Determinants and invertibility; Cramer's rule revisited** — det = 0 exactly when the
    transformation flattens; Greene's Cramer's rule, now with meaning.
- [ ] **27. How transformations scale area** `[GFX]` — |det| as the area-scaling factor; the bridge
    to change of variables and to conservation in physics.

## Section 6: Vector Spaces and the Big Picture

- [ ] **28. Vector spaces and subspaces** — the abstraction, and why polynomials and matrices being
    "vectors" too is a feature, not pedantry.
- [ ] **29. Span and linear independence** — what a set of vectors can build; redundancy made precise.
- [ ] **30. Basis and dimension** — minimal building sets; every plane through the origin is "2D" in
    exactly this sense.
- [ ] **31. Rank and the rank–nullity theorem** — the conservation law of linear maps: dimensions in
    = dimensions flattened + dimensions out.
- [ ] **32. The four fundamental subspaces** — column space, null space, row space, left null space:
    Strang's map of everything a matrix does.

## Section 7: Coordinates and Change of Basis

- [ ] **33. Coordinates relative to a basis** `[GFX]` — the same point, described in different
    languages; coordinates are a *choice*.
- [ ] **34. The change-of-basis matrix** `[GFX]` — translating between languages; local coordinates,
    world coordinates, and camera coordinates are just three bases.
- [ ] **35. Orthonormal bases; rotation as change of basis** `[GFX]` — the cleanest coordinate
    systems, and why rotation matrices have orthonormal columns.
- [ ] **36. The matrix of a map in different bases; similarity** — the same transformation can look
    simple or ugly depending on the basis: the motivation for eigenvectors.

## Section 8: Orthogonality and Least Squares

- [ ] **37. Orthogonal projection onto a subspace** — the nearest point in a subspace; error
    perpendicular to the target.
- [ ] **38. Gram–Schmidt orthogonalization** — manufacturing orthonormal bases from any basis.
- [ ] **39. Orthogonal matrices** `[GFX]` — transformations preserving lengths and angles: the rigid
    motions; why rotations compose and invert so cheaply (inverse = transpose).
- [ ] **40. Least squares** `[NET]` — the best solution when no exact one exists: fitting lines to
    noisy data, calibration, measurement.

## Section 9: Eigenvalues and Eigenvectors

- [ ] **41. Eigenvalues and eigenvectors** — the directions a transformation only stretches; the
    transformation's true axes.
- [ ] **42. The characteristic polynomial** — finding eigenvalues; Greene §11 root-finding cashes in.
- [ ] **43. Diagonalization; powers of a matrix** — in the eigenbasis, a transformation is just
    scaling; Aⁿ becomes trivial.
- [ ] **44. Solving recurrences with eigenvalues** `[CMP]` — Fibonacci in closed form via matrix
    powers; the linear-algebra view of Rosen Ch. 7's recurrences.
- [ ] **45. Symmetric matrices and principal axes** `[GFX]` — real eigenvalues, perpendicular
    eigenvectors: moments of inertia and best-fit bounding ellipses.
- [ ] **46. Markov matrices and steady states** `[NET]` `[DB]` — column sums 1, eigenvalue 1: where
    random processes settle; the idea under PageRank. (Bridge to Probability §8.)
- [ ] **47. Eigenvalues and stability of iteration** `[GFX]` — repeat a transformation forever:
    eigenvalues inside the unit circle decay, outside explode — the spectral reason simulations
    blow up.

## Section 10: Where Linear Algebra Meets Your Goals (capstone)

- [ ] **48. Complex numbers as 2×2 matrices** `[GFX]` — a + bi as a rotation-scaling matrix; Greene
    §25 and Section 3 of this module turn out to be the same subject.
- [ ] **49. Matrices over ℤₚ: the Hill cipher** `[CRY]` — linear algebra with modular arithmetic
    (Rosen Ch. 4); encrypt with a matrix, and see why linearity is a cryptographic weakness.
- [ ] **50. Adjacency matrices** `[NET]` `[DB]` — graphs as matrices; counting walks with matrix
    powers (Rosen Ch. 9 in matrix form).
- [ ] **51. The SVD in one lesson** — every matrix is rotate–stretch–rotate: what the singular value
    decomposition is and what it's for (compression, pseudo-inverse). Orientation, not mastery.
- [ ] **52. Exit project: the full 2D transform stack, on paper** `[GFX]` — a camera, two objects,
    one parented to the other; build every matrix, compose them, transform a point all the way
    to screen coordinates, then invert back for a "mouse click." Verify numerically.

---

## Exit criteria

You are done with this module when you can, cold:

- Read a 2×2 matrix and sketch what it does to the unit square, without computing.
- Build the 3×3 homogeneous matrix for "rotate 30° about the point (4, 2)" as a product of
  simpler matrices, and know the order.
- Explain what det A = 0 means three ways: algebraically, geometrically, and "as information."
- Compute eigenvalues/eigenvectors of a 2×2 matrix and use them to predict what Aⁿv does as
  n grows.
- Take a basis of ℝ², run Gram–Schmidt, and state why the result's matrix has inverse equal to
  its transpose.
