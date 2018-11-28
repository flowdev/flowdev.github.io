---
title:  "Why Go Generics With Contracts Are A Bad Idea And What Might Be Better"
---

## Why Go Generics With Contracts Are A Bad Idea And What Might Be Better

* contracts lead to very bad copy/paste programming.
* Interfaces are inherently easier to read.
* Operators could be used in interfaces as special methods (e.g. `__Less__(a, b int) bool`.
* This easily leads towards operator overloading: REALLY BAD unless you figure out how to limit use cases.
