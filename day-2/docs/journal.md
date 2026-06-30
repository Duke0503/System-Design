# Day 2 Journal

> Written after the day is done. The campfire reflection.
> Be honest. Nobody's grading this.

---

**What did I study today?**

"State and where it lives" — what happens when a second user shows up and the system needs to remember something for both of them.

I ran three exercises against the probe:
1. Removed the mutex, ran with `-race` — race detector caught what normal execution hid
2. Added file persistence, ran the process twice — state survived restart; then ran two instances simultaneously and watched last-write-wins corruption
3. Articulated in one sentence why databases exist: concurrent writes + crash atomicity + queryability

---

**What finally clicked?**

The race detector exercise. I expected to see broken output when I removed the mutex. Instead — nothing. Output looked correct. Both users recorded five visits, final state was right. I would have called that working code.

Then `go run -race main.go` printed `DATA RACE`. The bug was there the whole time. The scheduler just happened to be nice.

That's the lesson I'll carry: the absence of observable failure is not the same as correctness. The mutex existed for a real reason, and I would have deleted it based on a test that couldn't see the problem.

---

**What still confuses me?**

- How does a database actually serialize concurrent writes? There must be a lock somewhere — but is it a single global lock? Row-level? Something smarter?
- What is a write-ahead log actually doing? I understand the concept (journal intent first, then commit) but not the mechanics of how you recover from a crash mid-journal.
- The two-process file race: if both processes read the file, compute new state, then write — they each overwrite the other's work silently. Is this a TOCTOU bug? Is that what databases fundamentally solve?

---

**What did I gain today?**

A concrete three-tier mental model for state: in-memory (fast, volatile), file (durable, not concurrent-safe), database (durable, concurrent-safe, queryable). The choice between them isn't about preference — it's about what you need to survive: process restarts, concurrent writers, or queries across the data. Those requirements, not the technology, make the decision.

---

**Vibe check**

The race detector moment was genuinely surprising. I didn't expect "correct-looking output" to be the dangerous case. I expected broken output to be obvious. It wasn't. That's unsettling in a useful way — it means I need to treat "seems fine" differently than "is correct."

Tomorrow is Day 3: why databases exist. I'm going into it with the actual question now, not just the topic.
