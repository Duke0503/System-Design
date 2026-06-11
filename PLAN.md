# The Learning Plan

> Written before I know what I'm doing — which is kind of the point.
> This is the map. The territory will be messier. That's fine.

---

## The Big Idea

Every phase in this plan introduces a new problem *caused by solving the previous one*.
That's not a flaw in the curriculum — that's the actual shape of distributed systems.
You solve one thing, two new problems appear. Your job is to keep solving.

By the end, I want to be able to look at any system and ask the right questions:
- What breaks first?
- What trade-off was made here?
- What would I do differently?

---

## Phase 1 — What Even IS a System? (Day 1–5)

**After this phase I should be able to:**
- Trace a web request end-to-end and explain every hop
- Explain why databases exist (not just "they store data" — WHY)
- Distinguish stateful from stateless and know why it matters
- Feel what "failure" actually means at the most basic level

| Day | Topic | The Core Question |
|---|---|---|
| Day 1 | What is a system? | One user, one machine. What is actually happening? |
| Day 2 | State and where it lives | Add a second user. Where does the data go? |
| Day 3 | Why databases exist | Why not just write to a file? |
| Day 4 | Read vs. write asymmetry | Most systems read way more than they write. So what? |
| Day 5 | Failure — the first taste | What happens when the database goes down? |

---

## Phase 2 — When Scale Starts Hurting (Day 6–10)

**After this phase I should be able to:**
- Identify which of the 4 bottlenecks (CPU, memory, disk, network) a system is hitting
- Explain caching and — more importantly — explain what new problems it creates
- Describe what a load balancer does and why sticky sessions feel like cheating
- Understand why adding a second server is not just "run the same thing twice"

| Day | Topic | The Core Question |
|---|---|---|
| Day 6 | Bottlenecks | One server is slow. Which part of it? |
| Day 7 | Caching | What does caching actually buy you, and what does it cost? |
| Day 8 | Cache consistency | Write-through, write-behind, cache-aside — which do you choose? |
| Day 9 | Horizontal scaling | Two servers now. Where's the session? Where's the cache? |
| Day 10 | Load balancing | How do you split traffic? And why does "sticky" feel wrong? |

---

## Phase 3 — Data Gets Complicated (Day 11–15)

**After this phase I should be able to:**
- Choose between SQL and NoSQL with an actual reason, not a vibe
- Explain why queries slow down as tables grow and what indexes do about it
- Describe replication lag and why two nodes can disagree on what's "true"
- Explain CAP theorem without hand-waving — and feel why the choice is real

| Day | Topic | The Core Question |
|---|---|---|
| Day 11 | SQL vs. NoSQL | Not "which is better" — what are you optimizing FOR? |
| Day 12 | Indexes | Why does adding an index help reads but hurt writes? |
| Day 13 | The N+1 problem | Write it, feel the pain, fix it, never forget it |
| Day 14 | Database replication | Reads from replica, writes to primary — what could go wrong? |
| Day 15 | CAP theorem | If two nodes can't talk, what do you do? Pick. |

---

## Phase 4 — When Systems Talk to Each Other (Day 16–20)

**After this phase I should be able to:**
- Design a simple async flow and explain why you'd want one
- Explain the difference between at-least-once and exactly-once delivery
- Know when to reach for a message queue vs. a direct HTTP call
- Understand rate limiting as self-protection, not just "slowing users down"

| Day | Topic | The Core Question |
|---|---|---|
| Day 16 | Sync vs. async | Why would you ever NOT wait for an answer? |
| Day 17 | Message queues | What does "decoupled" actually mean in practice? |
| Day 18 | Delivery guarantees | At-least-once, exactly-once — why is the second one so hard? |
| Day 19 | API design | REST conventions — why they exist and where they break |
| Day 20 | Rate limiting | Protecting the system from itself |

---

## Phase 5 — Real System Design (Day 21 onward)

> **This phase only starts after Phases 1–4 are fully complete.**
> Not "mostly done." Not "I kind of get it." Done.
>
> The foundation has to be solid before the building goes up.
> If Phase 4 feels shaky, go back. Phase 5 will make no sense without it.

**After completing Phases 1–4, I should be ready to:**
- Hold every trade-off from the foundation in my head at once
- Look at a real-world system and ask: what breaks first? what was the designer's constraint? what would I do differently?
- Make a defensible design decision and explain it — not just describe the system

**What "designing a system" means here:**
This is not building a production system. It's sitting at a whiteboard (a markdown file)
and making real decisions with real reasoning. The Go code is a *probe* — a small binary
that makes one design decision falsifiable, not a full clone.

Every system gets:
- A design doc with requirements, constraints, and capacity estimates
- Key decisions with trade-offs written out (the "why", not just the "what")
- A Go probe that tests the one hardest design decision
- An honest journal entry: "I assumed X. The code showed me Y."

---

### System 1 — URL Shortener (Day 21–24)

The first real system. Small surface area, deep decisions.
You've used bit.ly. You have human intuition here. Use it.

**The questions you'll actually have to answer:**

| Day | Focus | The Hard Question |
|---|---|---|
| Day 21 | Requirements + capacity | What are you actually building? How much traffic, storage, latency? |
| Day 22 | Core design | How do you generate short codes without collisions? Where does data live? |
| Day 23 | Scale | Hot URLs — one tweet breaks you. How do you survive it? |
| Day 24 | Hard edges | Analytics, failure modes, what breaks at 10x and what you'd fix |

**The Go probe (Day 21–22):**
Generate 1,000,000 short codes using two strategies — `base62(sha256[:7])` vs `crypto/rand`.
Measure collision rate. Measure speed. Let the numbers change your assumption.

**Why this system first:**
Every large system you design after this — Twitter feed, YouTube, distributed KV store —
uses the same mental tools forged here: read/write asymmetry, hot path isolation,
ID generation under scale, the real cost of strong consistency.
You're not starting easy. You're starting right.

---

### System 2 — Rate Limiter (Day 25–28)
*Unlocks after System 1 is complete*

---

### System 3 — Twitter / Social Feed (Day 29–34)
*Unlocks after System 2 is complete*

---

### System 4 — YouTube / Video Pipeline (Day 35–40)
*Unlocks after System 3 is complete*

---

### System 5 — Distributed Key-Value Store (Day 41–47)
*Unlocks after System 4 is complete — the final boss*

---

## Progress

- [ ] Phase 1 — What Even IS a System?
- [ ] Phase 2 — When Scale Starts Hurting
- [ ] Phase 3 — Data Gets Complicated
- [ ] Phase 4 — When Systems Talk to Each Other
- [ ] Phase 5 — Real System Design
  - [ ] System 1: URL Shortener
  - [ ] System 2: Rate Limiter
  - [ ] System 3: Twitter / Social Feed
  - [ ] System 4: YouTube / Video Pipeline
  - [ ] System 5: Distributed Key-Value Store (final boss)

---

*"A plan is just a list of things that will change. Make it anyway."*
