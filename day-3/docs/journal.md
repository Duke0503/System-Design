# Day 3 Journal

> Written after the day is done. The campfire reflection.
> Be honest. Nobody's grading this.

---

**What did I study today?**

Today I studied why databases exist by comparing a plain JSON file against SQLite under concurrent writes.

The setup was simple: 10 goroutines, 100 increments each, 1000 expected increments total. The file-based counter finished at `9 / 1000`. SQLite finished at `1000 / 1000`.

That result made the lesson hard to ignore. A file can persist bytes, but it does not protect the meaning of an operation.

---

**What finally clicked?**

The thing that clicked was this: "increment" is not automatically one safe operation.

With a file, increment means read the file, parse the number, add one, and write the whole file back. Other goroutines can interrupt between any of those steps. Two writers can read the same old value and write the same new value. One writer can read while another writer is halfway through replacing the file.

So the file did not fail because persistence is bad. It failed because persistence alone does not coordinate mutation.

SQLite changed the boundary. I did not manually read, modify, and write bytes. I gave SQLite one statement:

```sql
UPDATE counter SET val = val + 1 WHERE id = 1;
```

That statement became the safe operation.

---

**What still confuses me?**

I understand that SQLite serializes writes, but I do not fully understand the internal mechanism yet.

I want to understand the difference between rollback journal mode and WAL mode. I also want to understand what happens when many readers are active while a writer is trying to commit.

The high-level answer is clear now. The internals are still fuzzy.

---

**What did I gain today?**

I gained a cleaner mental model:

```text
File persistence solves "will the bytes survive?"
Database transactions solve "will the change happen safely?"
```

Before today, I would have said a database is where data lives. That is true, but incomplete. A database is where safe changes to important data happen.

That distinction matters.

---

**Vibe check**

This one felt concrete. The output was brutal enough that I did not have to convince myself.

Seeing `9 / 1000` made the bug obvious. The file was not "a little unsafe". It was completely untrustworthy under concurrent writes.

SQLite getting `1000 / 1000` made the reason for databases feel less abstract. The point is not tables or SQL syntax. The point is controlled mutation of durable state.
