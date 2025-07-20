# Learning Go Language

This repository contains three Go programs demonstrating different aspects of Go development.

## 🏗️ Project Structure

```
learning_Go_language/
├── go.mod                         # Go module with dependencies
├── go.sum                         # Dependency checksums
├── Makefile                       # Build automation
├── .gitignore                     # Git ignore rules
├── cmd/                           # Executable commands
│   ├── read-line/                # 📖 Input reading program
│   │   └── main.go
│   ├── http-server/              # 🌐 REST API (in-memory)
│   │   └── main.go
│   └── http-server-gorm/         # 🗄️ REST API (persistent)
│       └── main.go
├── bin/                          # Built executables (gitignored)
├── .vscode/                      # VS Code configuration
│   ├── launch.json              # Debug configurations
│   ├── tasks.json               # Build tasks
│   └── settings.json            # Go language settings
└── docs/                         # Documentation
    ├── api-testing.md           # API testing guide
    └── learn_from.md
```

## 🚀 Three Projects Overview

### 1. 📖 Read-Line Program (`cmd/read-line/`)

- **Purpose**: Learn basic input/output and string handling
- **Features**: Interactive console input reader
- **Dependencies**: Standard library only

### 2. 🌐 HTTP Server (`cmd/http-server/`)

- **Purpose**: Learn REST API development with gorilla/mux
- **Features**: In-memory CRUD operations, JSON API
- **Dependencies**: `github.com/gorilla/mux`

### 3. 🗄️ HTTP Server with GORM (`cmd/http-server-gorm/`)

- **Purpose**: Learn database ORM and persistent storage
- **Features**: SQLite database, auto-migration, persistent data
- **Dependencies**: `github.com/gorilla/mux`, `gorm.io/gorm`, `gorm.io/driver/sqlite`

## 🎯 Quick Start

### Using Make (Recommended)

```bash
# Build all three programs
make build

# Run the read-line program
make run-readline

# Run the HTTP server (in-memory)
make run-server

# Run the HTTP server with database
make run-server-gorm

# Show all available commands
make help
```

### Manual Build and Run

```bash
# Build all programs
go build -o bin/read-line ./cmd/read-line
go build -o bin/http-server ./cmd/http-server
go build -o bin/http-server-gorm ./cmd/http-server-gorm

# Run programs
go run ./cmd/read-line                # Interactive input reader
go run ./cmd/http-server             # REST API on :8080
go run ./cmd/http-server-gorm        # REST API with database on :8080
```

```bash
# From project root
go build ./cmd/http-server
go build ./cmd/read-line
```

## VS Code Development & Debugging

This project is configured for VS Code with proper debugging support.

### Debugging Configurations Available

1. **Debug HTTP Server** - Debug the HTTP server program
2. **Debug Read-Line Program** - Debug the input reading program
3. **Debug Current File** - Debug whatever Go file is currently open
4. **Attach to HTTP Server** - Attach to a running server for debugging

### How to Debug

1. Open VS Code in this project folder
2. Set breakpoints in your Go code
3. Press `F5` or go to **Run and Debug** panel
4. Select the appropriate debug configuration
5. Click the green play button

### Available VS Code Tasks

Press `Ctrl+Shift+P` and type "Tasks: Run Task" to access:

- **Build All Programs** - Builds both executables
- **Build HTTP Server** - Builds only the HTTP server
- **Build Read-Line Program** - Builds only the read-line program
- **Run HTTP Server** - Runs the HTTP server
- **Run Read-Line Program** - Runs the read-line program
- **Clean Build Artifacts** - Removes built binaries
- **Go: Test All** - Runs all tests

### Command Line Debugging

For command-line debugging with `dlv`:

```bash
# Install delve if not already installed
go install github.com/go-delve/delve/cmd/dlv@latest

# Debug HTTP server
make debug-server

# Or manually
dlv debug ./cmd/http-server --headless --listen=:2345 --api-version=2
```

## Why This Structure?

This follows Go best practices for projects with multiple executables:

1. **Separation of Concerns**: Each executable has its own directory
2. **Clear Organization**: The `cmd/` directory clearly indicates these are executable commands
3. **No Conflicts**: No multiple `main` functions in the same package
4. **Scalable**: Easy to add more programs without conflicts

## Alternative Approaches

### 1. Build Tags (for simple cases)

You could use build tags to conditionally compile different main functions:

```go
//go:build server
// +build server

package main
// server code here
```

```go
//go:build readline
// +build readline

package main
// readline code here
```

Then build with: `go build -tags server` or `go build -tags readline`

### 2. Single main.go with flags

Use command-line flags to determine which program to run:

```go
package main

import (
    "flag"
    "fmt"
)

func main() {
    var mode = flag.String("mode", "", "Mode: server or readline")
    flag.Parse()

    switch *mode {
    case "server":
        runServer()
    case "readline":
        runReadline()
    default:
        fmt.Println("Use -mode=server or -mode=readline")
    }
}
```

But the **cmd/** structure is the most professional and recommended approach.
