# Learning Go Language

This repository contains multiple Go programs for learning purposes.

## Structure

```
learning_Go_language/
├── go.mod                    # Go module file
├── cmd/                      # Executable commands
│   ├── http-server/
│   │   └── main.go          # HTTP server example
│   └── read-line/
│       └── main.go          # Input reading example
└── docs/                    # Documentation
    └── learn_from.md
```

## Running the Programs

### HTTP Server

```bash
go run ./cmd/http-server
```

Then visit:

- http://localhost:8080 for homepage
- http://localhost:8080/about for about page

### Read Line Program

```bash
go run ./cmd/read-line
```

## Building Executables

### Build all programs

```bash
# Build HTTP server
go build -o bin/http-server ./cmd/http-server

# Build read-line program
go build -o bin/read-line ./cmd/read-line
```

### Build specific program

```bash
# From project root
go build ./cmd/http-server
go build ./cmd/read-line
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
