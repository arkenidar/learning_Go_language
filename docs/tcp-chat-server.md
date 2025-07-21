# TCP Chat Server

A simple raw TCP socket chat server that can be used with netcat as a client.

## Features

- **Multi-client support**: Multiple users can chat simultaneously
- **Nickname system**: Each user sets a nickname when connecting
- **Real-time messaging**: Messages are broadcast to all connected clients
- **Simple commands**: `/users`, `/quit`
- **Connection notifications**: Server announces when users join/leave
- **Netcat compatible**: Designed to work with standard netcat client

## Quick Start

### Start the server

```bash
make run-tcp-chat
# or
go run ./cmd/tcp-chat-server
```

The server will start on port 8888.

### Connect clients

In separate terminals, connect using netcat:

```bash
netcat localhost 8888
# or
nc localhost 8888
```

## Usage Example

1. **Start the server:**

   ```
   🚀 TCP Chat Server started on port :8888
   💡 Connect with: netcat localhost 8888
   📝 Multiple clients can connect simultaneously
   ```

2. **Client connects:**

   ```
   === Welcome to TCP Chat Server ===
   Enter your nickname: Alice
   Hello Alice! You can now chat. Type messages and press Enter.
   Commands: /quit to exit, /users to see online users
   === Chat started ===
   ```

3. **Chat in action:**
   ```
   [Alice] Hello everyone!
   *** Bob joined the chat
   [Bob] Hi Alice!
   [Alice] Welcome Bob!
   ```

## Commands

- **Regular messages**: Just type and press Enter
- **`/users`**: Show list of online users
- **`/quit`**: Disconnect from the chat

## Technical Details

- **Protocol**: Raw TCP sockets
- **Port**: 8888
- **Concurrency**: Each client handled in a separate goroutine
- **Thread-safe**: Uses mutex for client management
- **No authentication**: Simple nickname-based identification

## Architecture

```
┌─────────────────┐    TCP :8888     ┌─────────────────┐
│   netcat        │◄─────────────────►│   ChatServer    │
│   (Client 1)    │                  │                 │
└─────────────────┘                  │  ┌─────────────┐│
                                     │  │ Client Pool ││
┌─────────────────┐                  │  └─────────────┘│
│   netcat        │◄─────────────────►│                 │
│   (Client 2)    │                  │  ┌─────────────┐│
└─────────────────┘                  │  │  Broadcast  ││
                                     │  │   Engine    ││
┌─────────────────┐                  │  └─────────────┘│
│   netcat        │◄─────────────────►│                 │
│   (Client N)    │                  └─────────────────┘
└─────────────────┘
```

## Demo

Run the interactive demo:

```bash
./demo-tcp-chat.sh
```

This will guide you through starting the server and connecting clients.
