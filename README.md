# TODO CLI App

This repository contains a simple Command-Line Interface (CLI) application written in Go for managing your TODO tasks. The application allows users to add, organize, complete, list, and remove tasks. Tasks can be categorized for better organization.

## Commands

### Add

Add a new task to your TODO list.

**Usage:**

```bash
./todo add [flags]
```

**Aliases:** `add`, `a`

**Flags:**

- `-n, --name string`          Name of the task (required).
- `-d, --description string`   Description of the task.
- `-c, --category string`      Category to which the task belongs.

### Complete

Mark tasks as completed by their IDs.

**Usage:**

```bash
./todo complete [flags]
```

**Aliases:** `complete`, `c`

**Flags:**

- `-a, --all`          Mark all tasks as completed.

### List

List all tasks, with optional filters for category or completion status.

**Usage:**

```bash
./todo list [flags]
```

**Aliases:** `list`, `l`

**Flags:**

- `-c, --category string`   Filter tasks by category.
- `-x, --completed`         Show only completed tasks.

### Remove

Remove a task from your TODO list by its ID.

**Usage:**

```bash
./todo remove <id>
```

**Aliases:** `remove`, `r`, `rm`

## Example

Add a task named "Buy groceries" to the category "Shopping":

```bash
./todo add -n "Buy groceries" -c "Shopping" -d "Milk, eggs, and bread"
```

Mark the task with ID `1` as completed:

```bash
./todo complete 1
```

List all tasks in the "Shopping" category:

```bash
./todo list -c "Shopping"
```

Remove the task with ID `1`:

```bash
./todo remove 1
```

