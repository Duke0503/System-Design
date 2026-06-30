# Day 3 Notes — Why Databases Exist

> Raw notes taken while learning. Messy is fine — this is thinking in progress.

---

## The Core Question

Why not just write to a file?

---

## My Notes (fill this in as you go)

A file gives me persistence, but persistence is only one part of the problem.

Today's probe expected this:

```text
10 goroutines x 100 increments = 1000 expected
```

But the file result was:

```text
File-based counter:     9 / 1000  CORRUPTED
SQLite counter:      1000 / 1000  CORRECT
```

The file version is unsafe because one logical operation, "increment the counter once", is actually split into multiple steps:

```text
read file
parse count
count++
write file
```

That is a read-modify-write sequence. The file does not know those steps belong together.

If two goroutines both read `41`, both increment to `42`, and both write `42`, one increment disappears. Both writers did work, but the final state only shows one of them.

That is lost update.

It can get worse than simple lost updates. `os.WriteFile` rewrites the whole file. While one goroutine is writing, another goroutine may read an empty, partial, or invalid file. Because the probe ignores read/unmarshal/write errors, a bad read can silently turn into `0`, then the next write stores a low number again. That explains why the file result can be extremely wrong, like `9 / 1000`, not just `843 / 1000`.

SQLite is different because I send the database the intended mutation:

```sql
UPDATE counter SET val = val + 1 WHERE id = 1;
```

That statement is the unit of work. SQLite coordinates concurrent writers so each update gets applied safely. The important thing is not that SQLite stores data on disk. The important thing is that SQLite owns mutation of the data.

Mental model:

```text
File:     stores bytes
Database: stores data and controls safe changes to that data
```

What SQLite gives that a plain file does not:

- Atomic mutation: the increment either happens as a whole operation or it does not.
- Write serialization: conflicting writers do not freely overwrite each other.
- Transaction machinery: the database has rules for commit/rollback instead of random partial state.
- Queryability: I can ask for structured data with SQL instead of manually parsing bytes.

This connects directly to Day 2. A mutex protected in-memory state from concurrent goroutines, but memory disappeared on process exit. A file survived restart, but did not safely handle concurrent mutation. A database exists because real systems need both durability and safe coordination.

---

## Questions That Came Up

- How exactly does SQLite implement its locks internally?
- What exactly is written into the journal or WAL before a commit?
- If SQLite allows many readers but only one writer, how does that affect performance under write-heavy load?
- In a bigger database like Postgres, how different is the locking model from SQLite?

---

## Things That Clicked

The file did not fail because the disk forgot the data. It failed because the program had no safe boundary around the operation.

The operation I cared about was not "write bytes". The operation I cared about was "increment exactly once".

SQLite let me express that operation directly. The file made me manually assemble it from unsafe steps.

That is the core difference.
