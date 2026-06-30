# Day 2 — State and Where It Lives

**Phase:** 1 of 4 — *What Even IS a System?*
**Topic:** Add a second user. Where does the data go?

---

## What I Learned

State is not one thing — it's a **spectrum of durability**.

In-memory state lives in the process. It's the fastest thing you have, but it's a lie: the moment the process exits, every key, every counter, every visit record vanishes. The mutex in today's probe protected two goroutines from stepping on each other — that's a concurrency problem. But when the OS reclaims memory on exit, the mutex can't help. It never could. Concurrency and durability are two completely different problems.

The three tiers and what each one is actually for:

| Tier | Survives restart? | Concurrent-safe across processes? | Use when... |
|---|---|---|---|
| In-memory (map + mutex) | No | No | Speed matters, loss is acceptable |
| File on disk | Yes | No — last write wins | One process, append-only, or you enjoy pain |
| Database | Yes | Yes | Multiple users, multiple processes, data that matters |

A database isn't magic. It's a process that does three things a file can't: **serializes concurrent writes**, **journals intent before committing** (so a crash mid-write doesn't corrupt data), and **answers queries** about the data. A file does none of those three. That's why databases exist.

## The Moment It Clicked

The race detector.

I removed the mutex, ran the probe normally — output looked perfectly fine. Both users wrote their five visits, final state was correct. I would have shipped that. Then I ran `go run -race main.go` and got `DATA RACE` warnings. Two goroutines were touching the same map with nothing between them, and Go's scheduler happened to be polite enough that nothing broke. The bug was real. The damage was invisible.

That's what the race detector revealed: correctness under load is not the same as correctness under observation. The mutex wasn't ceremonial.

## Still Fuzzy On

- **How does a database actually serialize concurrent writes internally?** I know it does — but does it use a lock? A queue? Something else? What's the mechanism?
- **Crash atomicity** — if a database is mid-write when power cuts, how does it know which writes committed and which didn't? What's a "journal" actually storing?
- **Two processes racing on state.json** — I saw last-write-wins. But what if both processes read first, then write? Do they overwrite each other's changes silently?

## Tomorrow

Day 3: Why databases exist — why not just write to a file?

*(Already have the question from today. Tomorrow is the answer.)*

---

**Code:** [`src/main.go`](./src/main.go) — Go probe, two concurrent users, mutex-protected in-memory store, process exit
**Notes:** [`docs/notes.md`](./docs/notes.md) — raw notes taken during the day
**Journal:** [`docs/journal.md`](./docs/journal.md) — end-of-day reflection
