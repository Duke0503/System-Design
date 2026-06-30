# Day 2 Plan — State and Where It Lives

> Written before I open anything. This is my intention for today.

---

## What I'm Trying to Figure Out

Yesterday I had one user hitting one server and calling it a system.
Today I add a second user.

That sounds trivial. It isn't.

The moment a second user shows up, the question changes from
*"what happens?"* to *"what happens to WHOM, and does the other person see it?"*

That's state. And I don't actually know where it lives.

---

## Goals for Today

- [ ] Write down MY answer to "where does data go?" before reading anything
- [ ] Build a Go probe with two concurrent "users" and watch state behave
- [ ] Break the in-memory store on purpose — kill the process, lose the data
- [ ] Understand why "just write to a file" isn't the full answer either
- [ ] Write at least one sentence that explains WHY state is hard — not just what it is

---

## What I Already Think I Know (Before Learning)

*Writing this before I start so I can see how wrong I am later.*

When two users hit the same server... the server handles them both?
Data probably goes in memory somewhere. Or a database. The database "persists" it.

But if I'm being honest — I don't know what "state" actually means in a system.
Like, is the session a state? Is the database state? Is the in-memory variable state?
Is everything state?

That's genuinely all I've got. Let's go.

---

## Time Plan

- Read & take raw notes → `docs/notes.md`
- Build the Go probe → `src/main.go`
- Reflect → `docs/journal.md`
- Summarize → `README.md` (written last, after I actually know what happened)

---

## How I'll Know Today Was Worth It

If I can explain — clearly, without hand-waving — what "state" is,
where it lives in a running system, and what happens to it when the process dies.

And if my Go probe makes that visible: two users, shared state, process exit, gone.
That moment of "oh — it's just gone" is the lesson.
