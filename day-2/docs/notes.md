# Day 2 Notes — State and Where It Lives

> Raw notes taken while learning. Messy is fine — this is thinking in progress.

---

## The Core Question

Add a second user. Where does the data go?

---

## State — The Three Tiers

```
In-memory
  └── map[string]int behind a sync.Mutex
  └── lives in the process heap
  └── fastest access (nanoseconds)
  └── gone when process exits — no exceptions

File on disk
  └── os.WriteFile / os.ReadFile
  └── survives process restart
  └── one process writing = fine
  └── two processes writing simultaneously = last write wins, silent data loss

Database
  └── external process that serializes writes
  └── survives crashes (write-ahead log / journal)
  └── concurrent-safe by design
  └── queryable — you can ask questions about the data, not just read raw bytes
```

---

## The Mutex

`sync.Mutex` — one goroutine at a time inside the critical section.

```go
s.mu.Lock()       // acquire — blocks if another goroutine holds it
s.data[key] = val // critical section — only one goroutine here at a time
s.mu.Unlock()     // release — next goroutine can enter
```

What the mutex solves: concurrent goroutines stomping on the same map.
What the mutex does NOT solve: the process dying and taking the map with it.

Concurrency ≠ durability. Two separate problems.

---

## Exercise 1 — Race Detector

Remove mutex, run normally → output looks fine (scheduler is polite).
Run with `-race` flag:

```bash
go run -race main.go
```

Output: `DATA RACE` on stderr. Two goroutines touching `s.data` with no synchronization.

Lesson: absence of visible failure ≠ correctness. The race detector sees the memory access pattern, not just the output.

---

## Exercise 2 — File Persistence

Add on exit: `json.Marshal(store.data)` → write to `state.json`
Add on start: read `state.json` if exists → seed store

```bash
go run main.go && go run main.go   # state survives restart ✅
go run main.go & go run main.go    # two processes racing → last write wins ❌
```

The file race: both processes read the file (empty), both compute state, both write on exit. One overwrites the other. No error. No warning. Data just gone.

---

## Exercise 3 — The One Sentence

> "A file solves persistence but not concurrent writes, crash atomicity, or queryability — which is why a database exists."

Three problems a file can't solve:
1. **Concurrent writes** — multiple writers corrupt each other silently
2. **Crash atomicity** — power cut mid-write = partial/corrupt data, no way to recover
3. **Queryability** — you can read bytes, but you can't ask "give me all visits > 3"

A database exists specifically to solve these three. Not magic — a process with a protocol.

---

## Questions That Came Up

- How does a database serialize concurrent writes? Row-level locking? A queue? Both?
- What is a write-ahead log actually storing, and how does crash recovery use it?
- Is the two-process file race a TOCTOU (Time-of-Check-Time-of-Use) bug? Is that the class of problem databases fundamentally solve?
- What's the actual performance cost of going from in-memory to database? How many nanoseconds vs. milliseconds?
