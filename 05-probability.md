# Module 05 — Probability

<!-- progress:05-probability -->
![05 · Probability — Veerarajan progress](progress/05-probability.svg?v=2-0-46)

**0 / 46 lessons complete · 0%** — 0 of 9 sections finished · [completion log](LOG.md)
<!-- endprogress -->

**Book:** T. Veerarajan, *Probability, Statistics and Random Processes* (McGraw-Hill India).
Alternative: Sheldon Ross, *A First Course in Probability* (Pearson India). Later, when
networking gets serious: Kishor Trivedi, *Probability and Statistics with Reliability, Queuing
and Computer Science Applications*.
**Prerequisites:** Greene §16 and Rosen Ch. 6 (counting) before starting; Calculus §8
(improper integrals) before Section 5 of this module.
**Why this module exists:** Greene has one probability lesson and your edition of Rosen has no
probability chapter — but networking (reliability, queues), databases (hashing, cost
estimation), and cryptography (why attacks work) all run on it.

**The rhythm applies to every lesson:** Lesson → Study Notes → Practice Test.

Tags: `[GFX]` graphics/simulation · `[CMP]` compilers · `[NET]` networking · `[DB]` databases ·
`[CRY]` cryptography. No tag = foundation later lessons stand on.

---

## Section 1: Counting — the Bridge

- [ ] **1. The multiplication principle, permutations, combinations** — Greene §16 and Rosen
   6.1–6.3 consolidated into working fluency; this module assumes them cold.
- [ ] **2. Binomial coefficients in probability** — Pascal's triangle and the binomial theorem
   (Greene lesson 439) as counting tools.
- [ ] **3. Counting with repetition; balls into bins** `[DB]` — the counting model behind hash
   tables: n keys into m buckets.

## Section 2: Probability Foundations

- [ ] **4. Sample spaces and events** — outcomes as sets; the algebra of events is Greene §16's set
   operations doing real work.
- [ ] **5. The axioms of probability** — three rules, and everything else derived from them.
- [ ] **6. Equally likely outcomes** `[CRY]` — counting → probability; keyspace arithmetic: what
   "2¹²⁸ possible keys" actually buys you.
- [ ] **7. Inclusion–exclusion and the complement trick** — P(at least one) via 1 − P(none): the
   most-used computation pattern in the subject.
- [ ] **8. The birthday problem** `[CRY]` — how many people before two share a birthday: the single
   lesson that explains why hash outputs must be *twice* as long as the security you want.

## Section 3: Conditional Probability and Independence

- [ ] **9. Conditional probability** — probability, given what you already know.
- [ ] **10. The multiplication rule; tree diagrams** — chaining conditionals.
- [ ] **11. The law of total probability** — averaging over the ways something can happen.
- [ ] **12. Bayes' theorem** `[NET]` — reasoning backwards from effect to cause: was that packet loss
    congestion or corruption? (Also the engine of spam filters.)
- [ ] **13. Independence** `[NET]` `[CRY]` — when knowing one thing tells you nothing about another;
    independent trials as the model for repeated transmissions and repeated guesses.

## Section 4: Discrete Random Variables

- [ ] **14. Random variables and probability mass functions** — attaching numbers to outcomes.
- [ ] **15. Cumulative distribution functions** — P(X ≤ x) and why it determines everything.
- [ ] **16. Expectation** — the long-run average, and the center of the whole subject.
- [ ] **17. Linearity of expectation** `[DB]` `[CMP]` — expectations add even when nothing is
    independent: the workhorse of average-case algorithm analysis.
- [ ] **18. Variance and standard deviation** — how spread out; the second number every distribution
    gets.
- [ ] **19. Bernoulli and binomial distributions** `[NET]` — k successes in n independent trials:
    exactly k lost packets out of n sent.
- [ ] **20. The geometric distribution and memorylessness** `[NET]` — trials until first success:
    how many retransmissions until a packet gets through.
- [ ] **21. The Poisson distribution** `[NET]` — counts of rare events per interval: packet arrivals
    per second, requests per minute.
- [ ] **22. Applications: hashing and collisions** `[DB]` `[CRY]` — expected collisions, expected
    bucket sizes, and when a hash table degrades; balls-in-bins with expectations attached.

## Section 5: Continuous Random Variables

*(Requires Calculus §8 — improper integrals.)*

- [ ] **23. Probability density functions** — why single points have probability zero and areas are
    everything.
- [ ] **24. CDFs, expectation, and variance, continuous edition** — the same machinery, integrals
    replacing sums.
- [ ] **25. The uniform distribution** `[CRY]` — what "pick a random number" and "uniformly random
    key" formally mean; the distribution every generator tries to produce.
- [ ] **26. The exponential distribution** `[NET]` — time between arrivals; memorylessness again, and
    why it pairs with the Poisson distribution.
- [ ] **27. The normal distribution** — the bell curve: parameters, the 68–95–99.7 rule, and
    standardization.
- [ ] **28. Functions of a random variable; sampling by inverse CDF** `[GFX]` — turning uniform
    randomness into any distribution you want: how simulations and renderers sample.

## Section 6: Joint Distributions

- [ ] **29. Joint and marginal distributions** — two random quantities at once.
- [ ] **30. Independence of random variables** — the formal version, and how to check it.
- [ ] **31. Covariance and correlation** — do they move together; what correlation does and does not
    mean.
- [ ] **32. Sums of independent random variables** — distributions of sums; the convolution idea
    (why sums smooth out).

## Section 7: Inequalities and Limit Theorems

- [ ] **33. Markov and Chebyshev inequalities** `[DB]` `[CMP]` — guarantees from almost no
    information: the tail bounds used to analyze randomized algorithms.
- [ ] **34. The law of large numbers** — why averages settle down; the license behind estimating by
    simulation.
- [ ] **35. The central limit theorem** `[NET]` — why sums of anything look normal: the shape of
    aggregate traffic, aggregate error, aggregate everything.
- [ ] **36. Estimation basics** — sample mean, sample variance, and confidence intuition: how to
    benchmark a program and mean it.

## Section 8: Random Processes — the Networking Payoff

*(Veerarajan's home territory; this section is why that book was chosen.)*

- [ ] **37. The Bernoulli and Poisson processes** `[NET]` — randomness unfolding in time; the
    standard model of packet arrivals.
- [ ] **38. Markov chains** `[NET]` `[DB]` — state machines with random transitions; transition
    matrices (Linear Algebra §9 cashes in).
- [ ] **39. Steady-state distributions** `[NET]` — where a chain settles; solving πP = π.
- [ ] **40. Birth–death chains and the M/M/1 queue** `[NET]` — arrivals meet a server: your first
    queueing model, and the equation behind "why does latency explode as load nears 100%?"
- [ ] **41. Random walks and PageRank** `[NET]` `[DB]` — a walker on a graph (Rosen Ch. 9 + adjacency
    matrices); the steady state that ranked the web.

## Section 9: Applications Capstone

- [ ] **42. The birthday attack, quantified** `[CRY]` — redo lesson 8 as an attack: expected work to
    find a hash collision, and the 2ⁿ/² rule.
- [ ] **43. Entropy in one lesson** `[CRY]` `[NET]` — measuring uncertainty in bits: password
    strength, key strength, and the limit of compression. Orientation, not mastery.
- [ ] **44. Reliability: series and parallel systems** `[NET]` — when components fail independently:
    why redundancy multiplies nines.
- [ ] **45. Monte Carlo methods** `[GFX]` — estimate by sampling: π from random darts, integrals from
    random points; the idea at the bottom of modern rendering.
- [ ] **46. Exit project: a lossy link, on paper** `[NET]` — packet loss probability p, retransmit
    until success: derive expected transmissions and expected delay; then write a 20-line
    simulation and watch the law of large numbers agree with you.

---

## Exit criteria

You are done with this module when you can, cold:

- Compute P(at least one collision) among n random values in a space of size m — and invert it
  to size a hash.
- Use Bayes' theorem on a two-cause problem without a formula sheet.
- State the PMF/PDF, expectation, and variance of: Bernoulli, binomial, geometric, Poisson,
  uniform, exponential, normal — and say what each *models*.
- Explain memorylessness and name the two distributions that have it.
- Set up and solve the steady state of a 3-state Markov chain.
- Derive E[transmissions] = 1/(1−p) for a lossy link two different ways (geometric series, and
  first-step conditioning).
