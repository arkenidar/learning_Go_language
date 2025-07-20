# VS Code Configuration for Go Development

This document explains the VS Code configuration for this Go project with multiple executables.

## Files Created/Updated

### `.vscode/launch.json`

Debug configurations for:

- HTTP Server debugging
- Read-Line program debugging
- Current file debugging
- Remote debugging (attach mode)

### `.vscode/tasks.json`

Build and run tasks for:

- Building individual programs
- Building all programs (default build task)
- Running programs
- Cleaning build artifacts
- Running tests

### `.vscode/settings.json`

Go-specific settings for:

- Code formatting and linting
- Language server configuration
- Build and test settings
- Import organization

## Quick Start

1. **Install Go Extension**: Make sure you have the official Go extension installed
2. **Install Tools**: The extension will prompt you to install Go tools (accept all)
3. **Open Project**: Open this folder in VS Code
4. **Set Breakpoints**: Click in the gutter next to line numbers
5. **Start Debugging**: Press `F5` and select a debug configuration

## Debug Configurations Explained

### Debug HTTP Server

- **Program**: `./cmd/http-server`
- **Console**: Integrated Terminal
- **Use Case**: Debug the web server while making HTTP requests

### Debug Read-Line Program

- **Program**: `./cmd/read-line`
- **Console**: External Terminal
- **Use Case**: Debug interactive input program (needs external terminal for proper input)

### Debug Current File

- **Program**: `${file}` (current open file)
- **Console**: Integrated Terminal
- **Use Case**: Quick debugging of any Go file

### Attach to HTTP Server

- **Mode**: Remote attach
- **Port**: 2345
- **Use Case**: Attach to a server started with `dlv debug --headless`

## Keyboard Shortcuts

- `F5` - Start debugging
- `Ctrl+F5` - Run without debugging
- `F9` - Toggle breakpoint
- `F10` - Step over
- `F11` - Step into
- `Shift+F11` - Step out
- `Ctrl+Shift+F5` - Restart debugging
- `Shift+F5` - Stop debugging

## Tasks Shortcuts

- `Ctrl+Shift+P` → "Tasks: Run Task" - Show all tasks
- `Ctrl+Shift+P` → "Tasks: Run Build Task" - Run default build task

## Troubleshooting

### Go Tools Not Installed

1. Press `Ctrl+Shift+P`
2. Type "Go: Install/Update Tools"
3. Select all tools and install

### Debugger Not Working

1. Make sure you're in a Go workspace with `go.mod`
2. Check that the Go extension is enabled
3. Verify Go is installed: `go version`
4. Install Delve: `go install github.com/go-delve/delve/cmd/dlv@latest`

### Build Errors

1. Run `go mod tidy` to fix module dependencies
2. Check that all imports are correct
3. Use `go build ./cmd/...` to build all programs

## Advanced Configuration

### Custom Build Tags

Edit tasks.json to add build tags:

```json
"args": ["build", "-tags", "debug", "./cmd/http-server"]
```

### Environment Variables

Add to launch.json configurations:

```json
"env": {
    "GO_ENV": "development",
    "DEBUG": "true"
}
```

### Working Directory

Change working directory in launch.json:

```json
"cwd": "${workspaceFolder}/cmd/http-server"
```
