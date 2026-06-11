# Learning Journal — Rules & Style Guide

> These are the rules Duke and the crew agreed on before starting.
> Read this on day 1. Re-read it on day 30 when you've forgotten why you made these decisions.
> This is your contract with yourself — and with anyone else following along.

---

## 0. The Philosophy — This Is a Story, Not a Tutorial

This repo is not a reference. It's not a course. It's not documentation.

**It's a story.**

The story of one person — Duke — who decided to learn system design from zero,
day by day, in public. And like every good story, it has:

- **A hero** — you, at the start, knowing almost nothing
- **Obstacles** — concepts that broke your brain, days where nothing clicked,
  code that didn't work for reasons you couldn't explain
- **Turning points** — the moment CAP theorem finally made sense, the day you
  understood WHY databases exist, the first time you built something and felt
  it in your hands
- **Growth** — not just knowledge accumulated, but a different way of thinking
  about problems

When someone reads Day 1, they should feel your confusion and your curiosity.
When they read Day 15, they should feel how far you've come. When they reach
the end of a phase, they should feel the satisfaction of watching someone climb
a wall they're about to face themselves.

**Every file you write is a chapter.** Not a note. Not a log. A chapter.

This means:
- **Difficulties are not embarrassing — they are the plot.** If something
  broke your brain for two days, write those two days. That IS the story.
- **The struggle is the content.** A reader who is stuck on the same thing
  you were stuck on will find your "I was lost here for 3 days" more valuable
  than any polished explanation.
- **What you GAINED matters as much as what you LEARNED.** After every phase,
  write: what can I do now that I couldn't do before? How do I think
  differently? That transformation is what makes someone keep reading.

The person who finds this repo at day 45 should feel: *"this person was
exactly where I am, and they made it through — so can I."*

That feeling is the product. Everything else is just structure.

---

## 1. Repo Structure

Every single day follows the same layout. No exceptions, no creativity here —
consistency is what makes it easy to navigate six months from now.

```
learning-journal/
├── README.md              ← what this repo is (2 sentences, max)
├── PLAN.md                ← your full curriculum roadmap across all phases
│
├── day-1/
│   ├── README.md          ← what you actually learned that day (written AFTER)
│   ├── PLAN.md            ← your agenda/goals for that day (written BEFORE)
│   ├── docs/
│   │   ├── notes.md       ← raw notes while learning (written DURING)
│   │   └── journal.md     ← honest reflection (written AFTER)
│   └── src/
│       └── main.go        ← code for that day's concept
│
├── day-2/
│   ├── README.md
│   ├── PLAN.md
│   ├── docs/
│   │   ├── notes.md
│   │   └── journal.md
│   └── src/
│       └── main.go
```

**`src/` is part of every day folder — not optional.** Even if it's just 10 lines of Go
that illustrates the concept. Writing the code makes the concept stick in a way that
reading alone never does. No code = the day isn't really done.

---

## 2. What Each File Does

These must never overlap. When you wonder "where do I write this?" — this table answers.

| File | Written | Job | One-Line Description |
|---|---|---|---|
| `day-N/PLAN.md` | **Before** you start | Roadmap | What am I setting out to do today? |
| `docs/notes.md` | **During** learning | Working notes | Raw thoughts, snippets, diagrams — messy is fine |
| `src/main.go` | **During / after** | Code | Implement the concept, even if it's 10 lines |
| `docs/journal.md` | **Right after** | Reflection | What stuck, what confused, what surprised me |
| `day-N/README.md` | **Last** of the day | Snapshot | What did I actually learn? One paragraph summary |

**Writing order for every day:**
1. `day-N/PLAN.md` — before you open any book or video
2. `docs/notes.md` — while learning, stream of consciousness
3. `src/main.go` — code the concept while it's still fresh
4. `docs/journal.md` — right after, honest
5. `day-N/README.md` — final summary, what you'd tell a friend

---

## 3. Tone Rules

This is a personal learning journal, NOT a technical document.
Every file — README, PLAN, notes, journal — must sound like a real person
sharing their day, not a textbook author explaining a concept.

**The 6 rules:**

**Rule 1 — First-person, present tense**
- ❌ "System design principles should be understood before coding"
- ✅ "I'm learning that system design is really about trade-offs, not tools"

**Rule 2 — Name the struggle, not just the solution**
- ❌ "Horizontal scaling improves throughput"
- ✅ "I hit a wall when one server couldn't handle the traffic. That's when I started learning about horizontal scaling"

**Rule 3 — Use honest anchors**
Say "I realized...", "I didn't expect...", "I was confused about...", "this broke my brain...".
These ground everything in your journey, not some abstract truth.

**Rule 4 — Analogy first, jargon second**
- ❌ "CAP theorem dictates trade-offs between consistency and availability"
- ✅ "CAP theorem is like choosing between a library that's always perfectly up-to-date vs. one that's always open — you genuinely can't have both. Pick your pain."

**Rule 5 — One genuine reaction per section**
Every section earns one real human reaction:
- "Wait, that's not what I expected"
- "Oh — THAT'S why people do it that way"
- "This clicked immediately" / "I'm still lost, moving on"

**Rule 6 — Celebrate wins AND document dead ends**
- "Finally understood X when I tried Y"
- "Spent 2 hours going the wrong direction. Learned Z from it. Worth it."

**The meta-rule:** If a friend would fall asleep reading it, rewrite it.

---

## 4. Daily Journal — The 4 Questions

Every `docs/journal.md` answers exactly these four questions. No more, no less.

```markdown
**What did I study today?**
[One to three sentences. The topic, the resource, the time spent.]

**What finally clicked?**
[The one thing that went from confusing to clear. Be specific.]

**What still confuses me?**
[Honest. Don't pretend you understood everything. Future-you needs to know.]

**Where can someone follow along?**
[Link to the day folder, or the specific file/code.]
```

---

## 5. The Curriculum Spine

This is what you're learning, in order. Don't skip ahead.
Each phase introduces a new problem caused by solving the previous one — that's
not a bug, that's the actual shape of distributed systems.

### Phase 1 — What Is a System? (Day 1–5)
> Goal: trace a web request end-to-end; explain why databases exist; distinguish stateful from stateless

- Day 1: One user, one machine. What does request/response look like? Draw it on paper before writing any code.
- Day 2: Add a second user. Where does the data live? Plant the seed of "state."
- Day 3: Introduce a database. Not the theory — the WHY. Why not just write to a file?
- Day 4: Read vs. write asymmetry. Most systems read far more than they write. Why does that matter?
- Day 5: First failure scenario. What happens when the database goes down? Hello, durability.

### Phase 2 — When Scale Hurts (Day 6–10)
> Goal: identify a bottleneck; explain caching trade-offs; describe what a load balancer actually does

- Day 6: One server is slow. Is it CPU? Memory? Disk I/O? Network? Ask "which one?" before reaching for a solution.
- Day 7: Caching — the most impactful idea in distributed systems. And the new problems it creates (stale data, invalidation).
- Day 8: Cache consistency. Write-through vs. write-behind vs. cache-aside — real trade-offs, not theory.
- Day 9: Horizontal scaling. Add a second server. Now where's the session? Where's the cache?
- Day 10: Load balancing. Round robin, least connections, sticky sessions. Why does sticky sessions feel like cheating?

### Phase 3 — Data Gets Complicated (Day 11–15)
> Goal: choose SQL vs NoSQL with reasoning; explain indexes; describe CAP theorem without hand-waving

- Day 11: SQL vs. NoSQL — not "which is better" but "what are you optimizing for?"
- Day 12: Indexes — why queries slow as tables grow, and why indexes fix it.
- Day 13: The N+1 query problem. Write it, feel the pain, fix it.
- Day 14: Database replication. Primary/replica. Replication lag — reads and writes can disagree on what's "true."
- Day 15: CAP theorem. Don't memorize it — sit with the question. If two nodes can't talk, what do you do?

### Phase 4 — The System Talks to Itself (Day 16–20)
> Goal: design a simple async flow; explain delivery guarantees; know when rate limiting matters

- Day 16: Synchronous vs. asynchronous. Why would you ever NOT wait for an answer?
- Day 17: Message queues — the post office analogy. Why this decouples services from each other's availability.
- Day 18: At-least-once vs. exactly-once delivery. Idempotency as the practical solution.
- Day 19: API design — REST conventions, why they exist, where they fall short at scale.
- Day 20: Rate limiting — protecting your system from itself as much as from bad actors.

---

## 6. Commit Message Rules

Short, present tense, scoped to the day or file changed:

```
day-1: add what-is-a-system notes
day-3: add database lesson + code example
plan: update Phase 2 milestones
day-7: add caching journal — cache invalidation broke my brain
```

Never: "update files", "fix stuff", "day 7 done". Give future-you something to scan.

---

## 7. Writing the Story — What Each File Carries

Every file has a narrative job inside the larger story:

| File | Story Role | The Question It Answers |
|---|---|---|
| `PLAN.md` (day) | *Setting the scene* | "What was Duke trying to figure out today?" |
| `docs/notes.md` | *The adventure* | "What happened while he was learning?" |
| `src/main.go` | *Proof of the journey* | "What did he build to make it real?" |
| `docs/journal.md` | *The campfire reflection* | "How did it actually feel? What changed?" |
| `README.md` (day) | *Chapter summary* | "What did Duke walk away with?" |

**How to write a difficulty:**
Don't hide it. Name it. Give it a full scene:
> "I spent the whole morning convinced I understood caching. Then I tried to
> implement it and realized I had no idea what 'invalidation' actually meant
> in practice. Spent 2 hours going in circles. The thing that finally unlocked
> it was..."

**How to write a turning point:**
Make the reader feel the click:
> "And then it hit me. The reason we need a message queue isn't about speed —
> it's about *not caring* if the other service is even running right now.
> That's the whole point. The producer doesn't wait. The consumer catches up
> when it can. I drew this on paper and just stared at it for a minute."

**How to write what you gained:**
At the end of each phase, don't just list topics covered. Write the
transformation:
> "Four weeks ago I thought 'the database' was just... a place where data
> lives. Now I know it's a set of trade-offs I'm making every time I choose
> one. Read replicas, eventual consistency, index write amplification — these
> aren't features, they're prices. I'm now the kind of person who asks
> 'what do we give up to get this?' That's new."

---

## 8. The Golden Rules (Read These When You Want to Quit)

1. **Write `PLAN.md` before you open any resource.** Intention before consumption.
2. **`notes.md` is allowed to be messy.** That's the point. Don't clean it up.
3. **Never skip `journal.md`.** The reflection is where the learning actually locks in.
4. **Don't pre-create empty folders.** Only `day-N/` exists when you sit down to actually learn.
5. **Don't system-hop before depth.** Three days on one concept beats one day on three concepts.
6. **The order of the curriculum matters.** Each week introduces problems caused by the previous week's solutions.
7. **Difficulties are the story, not the shame.** Write the days where nothing clicked. They are the most valuable chapters.
8. **Honest footprints over polished content.** The "I was confused about X for 3 days" is more valuable than any perfect explanation.
9. **Write what you GAINED, not just what you LEARNED.** Knowledge is information. Gain is transformation. The reader wants to feel the second one.
10. **Cadence over perfection.** A 10-minute entry on a bad day beats zero entries. Just show up.

---

*Last updated: 2026-06-11*
*These rules were built in conversation — they can be updated when something stops working.*
