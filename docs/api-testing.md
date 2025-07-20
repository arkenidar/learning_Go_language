# API Testing Examples

Now that the HTTP server is running with gorilla/mux, you can test the REST API endpoints:

## Test the API endpoints

### 1. Get all todos

```bash
curl http://localhost:8080/todos
```

### 2. Get a specific todo

```bash
curl http://localhost:8080/todos/1
```

### 3. Create a new todo

```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"New todo item","done":false}'
```

### 4. Update a todo

```bash
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated todo","done":true}'
```

### 5. Delete a todo

```bash
curl -X DELETE http://localhost:8080/todos/1
```

## Using Make commands

### Start the server

```bash
make run-server
```

### Build all programs

```bash
make build
```

### Install dependencies

```bash
make deps
```

### Clean up module dependencies

```bash
make tidy
```

## VS Code Integration

The project is now configured for VS Code debugging:

1. Open the project in VS Code
2. Go to Run and Debug (Ctrl+Shift+D)
3. Select "Debug HTTP Server" from the dropdown
4. Press F5 to start debugging

You can set breakpoints in your code and debug step by step.
