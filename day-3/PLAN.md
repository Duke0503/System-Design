# Day 3 Plan — Why Databases Exist

> Written before I open anything. This is my intention for today.

---

## What I'm Trying to Figure Out

Yesterday I ended with a question:

> "A file solves persistence but not concurrent writes, crash atomicity,
> or queryability — which is why a database exists."

I wrote that sentence. I believe it. But I don't actually *feel* it yet.
Today I'm going to feel it. I'm going to prove it with numbers.

The question isn't "what is a database?" I know what one is.
The question is: **what goes wrong when you don't have one?**

---

## Goals for Today

- [x] Write down what I think "concurrent write corruption" actually looks like — before running code
- [x] Build a probe: N goroutines, a counter, two strategies — file vs. SQLite
- [x] Run the file version and watch the count be wrong
- [x] Run the SQLite version and watch it be exactly right
- [x] Explain in one paragraph what SQLite does internally that the file doesn't

---

## What I Already Think I Know (Before Learning)

*Writing this before I start so I can see how wrong I am later.*

A database... persists data? It has tables. You write SQL.
It's concurrent-safe somehow — I think it uses locks?
There's something called a transaction. `BEGIN`, `COMMIT`, `ROLLBACK`.

Why not just write to a file? Because... two writers corrupt each other.
I saw that yesterday with `state.json` and two processes.

But I don't know *how* a database avoids that. Internal locking? A queue?
Does every write go through a single point? Is that why it's "slower" than in-memory?

That's genuinely all I've got. Let's go prove it.

---

## Time Plan

- Read & take raw notes → `docs/notes.md`
- Build the Go probe → `src/main.go`
- Reflect → `docs/journal.md`
- Summarize → `README.md` (written last, after I actually know what happened)

---

## How I'll Know Today Was Worth It

If I can run `go run main.go` and see:

```
File-based counter:   743 / 1000  ✗ CORRUPTED
SQLite counter:       1000 / 1000 ✓ CORRECT
```

And then explain — in plain English, no hand-waving — *why* those two numbers differ.
