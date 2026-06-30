# Day 4 Plan - Read vs. Write Asymmetry

> Written before I open anything. This is my intention for today.

---

## What I'm Trying to Figure Out

Yesterday I learned why databases exist: a file can persist bytes, but a database controls safe changes to data.

Today the question changes.

If a database is the source of truth, do all requests stress it the same way? Probably not. In most real systems, people read far more than they write. They open profiles, refresh feeds, view product pages, check order status, and search catalogs many more times than they update them.

The core question:

**If reads happen 100x more than writes, what should the system care about first?**

---

## Goals for Today

- [ ] Predict what changes when a workload is read-heavy
- [ ] Build a Go probe with many reads and few writes
- [ ] Compare direct source-of-truth reads against a read-optimized path
- [ ] Notice the new responsibility writes get: keeping read state correct
- [ ] Explain why read/write ratio changes system design decisions

---

## What I Already Think I Know (Before Learning)

*Writing this before I start so I can see how wrong I am later.*

Reads are probably cheaper than writes because reads do not change data.
Writes are probably more dangerous because they change the source of truth.

But if a system has way more reads than writes, maybe reads become the bigger bottleneck even if each individual read is simple.

I think this is why systems add caches, replicas, indexes, and materialized views.

But I do not want to jump to those tools yet. Today I want to feel the shape of the workload first.

---

## Time Plan

- Read & take raw notes -> `docs/notes.md`
- Build the Go probe -> `src/main.go`
- Reflect -> `docs/journal.md`
- Summarize -> `README.md` (written last, after I actually know what happened)

---

## How I'll Know Today Was Worth It

If I can look at output like this:

```text
1000 reads, 10 writes = 100:1 read/write ratio

Direct reads: source read 1000 times
Read-optimized path: source read about 10 times
```

And explain why optimizing reads helps so much, but also why every write now has an extra job.
