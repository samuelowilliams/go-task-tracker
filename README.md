# My Go Task CLI

A simple CLI tool written in Go to manage tasks.

## Requirements

- Go 1.24 or higher must be installed.  
- A Linux/Mac/Windows terminal.  

Check Go version:

```bash
go version
````

## Setup

```bash
git clone https://github.com/samuelowilliams/go-task-tracker
cd go-task-tracker
go build -o task_tracker
```

## Usage

Add a new task:

```bash
./task_tracker add "Learn Go"
```

List all tasks:

```bash
./task_tracker list
```

List tasks by status:

```bash
./task_tracker list done
```

```bash
./task_tracker list todo
```

```bash
./task_tracker list in-progress
```

Update task status:

```bash
./task_tracker done 1
```

```bash
./task_tracker in-progress 1
```

Tasks are stored in `tasks.json` in the project directory.
