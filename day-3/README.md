# Day 3 — Why Databases Exist

**Phase:** 1 of 4 — *What Even IS a System?*
**Topic:** Why not just write to a file?

---

> Written after the day is complete. Fill this in when you're done.

## What I Learned

A file is persistence, not coordination.

Today's probe made that visible. I expected 1000 increments. The file-based counter finished at `9 / 1000`. SQLite finished at `1000 / 1000`.

The file version failed because "increment once" was not one safe operation. It was a sequence: read the file, parse the number, increment in memory, then write the whole file back. Multiple goroutines could step on each other between those steps. Some increments were overwritten. Some reads may have happened while another goroutine was rewriting the file.

SQLite succeeded because the mutation was expressed as one database statement:

```sql
UPDATE counter SET val = val + 1 WHERE id = 1;
```

SQLite owns the coordination around that statement. It serializes conflicting writes and applies the update atomically. That is the difference: a file stores bytes, but a database controls safe changes to data.

## The Moment It Clicked

The moment was seeing this output:

```text
File-based counter:     9 / 1000  CORRUPTED
SQLite counter:      1000 / 1000  CORRECT
```

That made the problem impossible to hand-wave. The file did not lose a tiny number of updates. It lost almost all of them.

The real lesson: the operation I cared about was "increment exactly once", but the file only understood "read bytes" and "write bytes". SQLite understood the mutation as a single operation.

## Still Fuzzy On

- How SQLite locking works internally.
- What exactly gets written to a rollback journal or WAL during commit.
- How SQLite behaves when there are many readers and one writer.
- How this compares to a server database like Postgres.

## Tomorrow

Day 4: Read vs. write asymmetry — most systems read way more than they write. So what?

---

**Code:** [`src/main.go`](./src/main.go) — file vs. SQLite counter, 10 goroutines × 100 increments
**Notes:** [`docs/notes.md`](./docs/notes.md) — raw notes taken during the day
**Journal:** [`docs/journal.md`](./docs/journal.md) — end-of-day reflection
