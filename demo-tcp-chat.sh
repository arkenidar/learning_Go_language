#!/bin/bash
# demo-tcp-chat.sh - Demonstration script for the TCP chat server

echo "🚀 TCP Chat Server Demo"
echo "======================="
echo ""

# Check if netcat is available
if ! command -v nc &> /dev/null && ! command -v netcat &> /dev/null; then
    echo "❌ netcat (nc) is not installed. Please install it first:"
    echo "   Ubuntu/Debian: sudo apt install netcat"
    echo "   macOS: brew install netcat"
    echo "   Or use: sudo apt install netcat-openbsd"
    exit 1
fi

echo "📋 This demo will:"
echo "   1. Start the TCP chat server on port 8888"
echo "   2. Show you how to connect with netcat"
echo "   3. Demonstrate multi-client chat functionality"
echo ""

echo "💡 How to test the chat server:"
echo ""
echo "1️⃣ Start the server:"
echo "   make run-tcp-chat"
echo "   (or: go run ./cmd/tcp-chat-server)"
echo ""
echo "2️⃣ Connect clients (in separate terminals):"
echo "   netcat localhost 8888"
echo "   (or: nc localhost 8888)"
echo ""
echo "3️⃣ Chat features:"
echo "   • Enter your nickname when prompted"
echo "   • Type messages and press Enter to send"
echo "   • Use /users to see who's online"
echo "   • Use /quit to disconnect"
echo "   • Multiple clients can chat simultaneously"
echo ""

read -p "🤔 Do you want to start the server now? (y/n): " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo ""
    echo "🚀 Starting TCP chat server..."
    echo "📝 Open another terminal and run: netcat localhost 8888"
    echo "🛑 Press Ctrl+C to stop the server"
    echo ""
    make run-tcp-chat
else
    echo ""
    echo "👍 You can start the server manually with:"
    echo "   make run-tcp-chat"
    echo ""
    echo "🔗 Then connect clients with:"
    echo "   netcat localhost 8888"
fi
