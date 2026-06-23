Absolutely — here is a more natural, human-style version of the notes.

# Notes on Log Storage

This part is about **how a database keeps track of changes** without rewriting the whole data file every time something changes. The main idea is to use a **log file** that stores changes one after another, in order, instead of trying to edit old data directly.

## What the log is

Think of the log like a diary for the database. Every time you add, update, or delete something, the database writes that change at the end of the log file. It does not go back and edit old entries. That means the log keeps a full history of what happened.

Example:

```text
SET k1 = x
SET k2 = y
SET k1 = z
DEL k2
```

If you replay those actions in order, the final result is:

```text
k1 = z
```

because the later `SET k1 = z` replaces the older value, and `DEL k2` removes `k2`.

## Why this is useful

The big advantage is that **writing to the end of a file is fast**. Also, if the database crashes, it can read the log again and rebuild the current state from scratch. So the log is really about **durability and recovery**.

## How the database uses it

The database usually does **not** wait around for some future time to apply the log. It writes the change, and then the in-memory map is updated right away so reads can return the latest value immediately. The log is there so the database can recover later if something goes wrong.

So the flow is basically:

1. A write comes in.
2. The database appends it to the log.
3. The in-memory state is updated.
4. The database responds.

Reads are usually served from memory, not by scanning the log file every time.

## What happens when the database starts

When the database starts up, it opens the log file and reads every entry from the beginning. It keeps applying each change to an in-memory map until it reaches the end of the file. That replay process recreates the latest state.

## How delete works

Deletes are stored in the log too. In this implementation, each entry has a `deleted` flag. If that flag is true, the entry means “remove this key.” If it is false, the entry means “store this key/value pair.”

## Important thing to remember

The log is not the final permanent structure forever. If it keeps growing, startup gets slower because the database has to replay more and more entries. That is why databases later compact or merge old log entries into a more efficient structure.

# Simple way to remember it

The easiest way to think about it is:

* **Memory** = current working state
* **Log** = history of changes
* **Replay log** = rebuild the current state after restart

So the log is like a backup record of everything the database has done.

If you want, I can turn this into a cleaner **study-note version** or a **super short exam revision version**.
