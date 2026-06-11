# System Design — Duke's Learning Journal

Hey, I'm Duke.

I have 2 years of software experience and one embarrassing confession:
I didn't actually understand how systems work. I could write code. I could ship features.
But ask me *why* we use a message queue here, or *what* breaks first when traffic spikes,
or *how* a URL shortener handles 10 million requests a day — and I'd wave my hands and say
something that sounded smart but meant nothing.

So I decided to fix that. From zero. In public. Day by day.

This is not a course. Not a tutorial. Not a polished reference.

**It's my story of figuring this out** — the confusion, the dead ends, the 2-hour rabbit
holes that taught me nothing, and the rare beautiful moments when something finally clicked.
Written honestly, because the resource I wish I had wasn't a perfect explanation —
it was someone who was lost in the same place I was, and showed me the way out.

Go is the language I'm using to make concepts real. You don't have to know Go.
The system design thinking is the point. Go is just how I prove I understood it.

If you're where I was when I started — overwhelmed, not sure what to learn first,
reading things that assume you already know things — you're exactly who this is for.

---

## The Journey

Two parts. The foundation comes first. No skipping.

### Part 1 — Foundation (Day 1–20)

Learning the individual concepts — what they are, why they exist,
and what breaks when you ignore them.

| Phase | Days | What I'm Figuring Out |
|---|---|---|
| Phase 1 | Day 1–5 | What even IS a system? |
| Phase 2 | Day 6–10 | When one server isn't enough |
| Phase 3 | Day 11–15 | When data gets complicated |
| Phase 4 | Day 16–20 | When systems have to talk to each other |

### Part 2 — Real System Design (Day 21+)

Only after the foundation is solid. Designing actual systems end to end —
making real decisions, defending real trade-offs, writing Go code that
makes the design falsifiable rather than just theoretical.

| System | Days | What Makes It Hard |
|---|---|---|
| URL Shortener | Day 21–24 | Read/write asymmetry, hot paths, ID generation at scale |
| Rate Limiter | Day 25–28 | Distributed state, fail open vs. closed, algorithm trade-offs |
| Twitter / Social Feed | Day 29–34 | Fan-out, celebrity problem, eventual consistency |
| YouTube / Video Pipeline | Day 35–40 | Async processing, blob storage, CDN design |
| Distributed KV Store | Day 41–47 | The final boss — building the infrastructure itself |

Full roadmap with day-by-day breakdown → [PLAN.md](./PLAN.md)

---

## How Each Day Works

Every day is a folder. Same structure, every time:

```
day-N/
├── README.md       ← what I actually learned (read this first)
├── PLAN.md         ← what I was trying to figure out (read this second)
├── docs/
│   ├── notes.md    ← raw notes as I learned — messy, honest, timestamped
│   └── journal.md  ← how it actually felt, what surprised me, what's still fuzzy
└── src/
    └── main.go     ← Go code that made the concept real
```

**Where to start:** `day-1/README.md`

**How to navigate:** each day's README links forward to what's next
and backward to what it depends on.

---

## Rules I Made for Myself

The consistency rules behind how everything is written → [docs/RULES.md](./docs/RULES.md)

The short version: every file is written like a person sharing their day,
not a textbook explaining a concept. Difficulties are part of the story,
not things to hide. The "I was confused about this for 3 days" entries
are the most valuable ones.

---

## If You Want to Follow Along

- **Star the repo** if you want to track progress
- **Open an issue** if something is wrong or you know a better way — I'm learning, corrections welcome
- **Fork it** if you want to run the same journey yourself

This repo is not trying to be famous. It's trying to be useful to the one person
who finds it at 11pm, stuck on the same thing I was stuck on, and needs to know
someone else got through it.

---

*Started: 2026-06-11*
*Status: Day 1 — in progress*

> *"The best time to understand system design was before you needed to.
> The second best time is right now, confused, starting from zero."*
