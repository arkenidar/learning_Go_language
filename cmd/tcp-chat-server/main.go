package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

// Client represents a connected chat client
type Client struct {
	conn     net.Conn
	nickname string
	writer   *bufio.Writer
}

// ChatServer manages all connected clients
type ChatServer struct {
	clients map[*Client]bool
	mutex   sync.RWMutex
}

// NewChatServer creates a new chat server instance
func NewChatServer() *ChatServer {
	return &ChatServer{
		clients: make(map[*Client]bool),
	}
}

// AddClient adds a new client to the server
func (s *ChatServer) AddClient(client *Client) {
	s.mutex.Lock()
	s.clients[client] = true
	s.mutex.Unlock()

	// Announce new client to all others (outside the lock to avoid deadlock)
	message := fmt.Sprintf("*** %s joined the chat", client.nickname)
	s.broadcastMessage(message, client)

	fmt.Printf("[SERVER] Client %s connected from %s\n",
		client.nickname, client.conn.RemoteAddr())
}

// RemoveClient removes a client from the server
func (s *ChatServer) RemoveClient(client *Client) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.clients[client]; exists {
		delete(s.clients, client)
		client.conn.Close()

		// Announce client left to all others
		message := fmt.Sprintf("*** %s left the chat", client.nickname)
		s.broadcastMessage(message, nil)

		fmt.Printf("[SERVER] Client %s disconnected\n", client.nickname)
	}
}

// BroadcastMessage sends a message to all clients except the sender
func (s *ChatServer) broadcastMessage(message string, sender *Client) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for client := range s.clients {
		if client != sender {
			client.writer.WriteString(message + "\n")
			client.writer.Flush()
		}
	}
}

// HandleClient manages communication with a single client
func (s *ChatServer) HandleClient(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	// Welcome message and nickname prompt
	writer.WriteString("=== Welcome to TCP Chat Server ===\n")
	writer.WriteString("Enter your nickname: ")
	writer.Flush()

	// Read nickname
	nickname, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("[ERROR] Failed to read nickname: %v\n", err)
		return
	}
	nickname = strings.TrimSpace(nickname)
	
	if nickname == "" {
		nickname = fmt.Sprintf("User_%d", time.Now().Unix()%1000)
	}

	// Create client
	client := &Client{
		conn:     conn,
		nickname: nickname,
		writer:   writer,
	}

	// Add client to server
	s.AddClient(client)
	defer s.RemoveClient(client)

	// Send instructions
	writer.WriteString(fmt.Sprintf("Hello %s! You can now chat. Type messages and press Enter.\n", nickname))
	writer.WriteString("Commands: /quit to exit, /users to see online users\n")
	writer.WriteString("=== Chat started ===\n")
	err = writer.Flush()
	if err != nil {
		fmt.Printf("[ERROR] Failed to send instructions to %s: %v\n", nickname, err)
		return
	}

	// Handle incoming messages
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("[INFO] Client %s disconnected: %v\n", client.nickname, err)
			break // Client disconnected
		}

		message = strings.TrimSpace(message)
		
		if message == "" {
			continue
		}

		// Handle commands
		if strings.HasPrefix(message, "/") {
			s.handleCommand(client, message)
			continue
		}

		// Broadcast regular message
		formattedMessage := fmt.Sprintf("[%s] %s", client.nickname, message)
		s.broadcastMessage(formattedMessage, client)
		
		// Echo the message back to the sender so they can see it was processed
		client.writer.WriteString(formattedMessage + "\n")
		client.writer.Flush()

		// Log message to server console
		fmt.Printf("%s\n", formattedMessage)
	}
}

// HandleCommand processes client commands
func (s *ChatServer) handleCommand(client *Client, command string) {
	switch {
	case command == "/quit":
		client.writer.WriteString("Goodbye!\n")
		client.writer.Flush()
		client.conn.Close()

	case command == "/users":
		s.mutex.RLock()
		users := make([]string, 0, len(s.clients))
		for c := range s.clients {
			users = append(users, c.nickname)
		}
		s.mutex.RUnlock()

		response := fmt.Sprintf("Online users (%d): %s\n",
			len(users), strings.Join(users, ", "))
		client.writer.WriteString(response)
		client.writer.Flush()

	default:
		client.writer.WriteString("Unknown command. Available: /quit, /users\n")
		client.writer.Flush()
	}
}

func main() {
	const port = "127.0.0.1:8888"

	// Create chat server
	server := NewChatServer()

	// Start TCP listener
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Printf("🚀 TCP Chat Server started on port %s\n", port)
	fmt.Printf("💡 Connect with: netcat localhost 8888\n")
	fmt.Printf("📝 Multiple clients can connect simultaneously\n")
	fmt.Println("---")

	// Accept connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Failed to accept connection: %v\n", err)
			continue
		}

		// Handle each client in a separate goroutine
		go server.HandleClient(conn)
	}
}
