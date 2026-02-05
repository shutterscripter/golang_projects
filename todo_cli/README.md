# Todo CLI

A simple command-line todo list application written in Go that allows you to manage your tasks directly from the terminal.

## Features

- **Add todos**: Create new tasks with a title
- **List todos**: View all your tasks with their status
- **Toggle completion**: Mark tasks as complete or incomplete
- **Edit todos**: Update task titles
- **Delete todos**: Remove tasks from your list
- **Persistent storage**: All todos are saved to a JSON file

## Installation

```bash
go build -o todo
```

## Usage

### List all todos
```bash
go run . -list
```

### Add a new todo
```bash
go run . -add "Buy groceries"
```

### Toggle todo completion status
```bash
go run . -toggle 1
```

### Edit a todo
```bash
go run . -edit "1:Buy groceries and cook dinner"
```

### Delete a todo
```bash
go run . -delete 1
```

## Data Storage

Todos are stored in `todos.json` in the same directory as the application.

## Requirements

- Go 1.16 or higher
