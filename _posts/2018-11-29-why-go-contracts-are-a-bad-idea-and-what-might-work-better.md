---
layout: post
---

title:  "Why Go Contracts Are A Bad Idea And What Might Work Better"

1. Intro: contracts/generics is a fundamental and long term decision.
1. 'contracts' is still the official proposal. All criticism so far didn't change that.
1. The Go community is changing as more and more pragmatic people enter it.
1. these people want to get a job done and aren't interested in every detail of the tool they use (Go).
1. They like to use pragmatic programming techniques like copy/paste to get the job done.

* contracts lead to very bad copy/paste programming.
* Interfaces are inherently easier to read.
* Operators could be used in interfaces as special methods (e.g. `__Less__(a, b int) bool`.
* This easily leads towards operator overloading: REALLY BAD unless you figure out how to limit use cases.



### Phrases

* Pragmatic programmers can very well be better programmers.
  They often know the problem domain very well and it is better to solve
  the right problem in a simple, straight forward way than elegantly solving the wrong problem.
  Or even solving the (right or wrong) problem in an overengineered way.
* Something else...
