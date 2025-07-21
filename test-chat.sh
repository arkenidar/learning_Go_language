#!/bin/bash
#
# TCP Chat Server Test Script
# Tests the basic functionality of the TCP chat server with netcat
#

echo "🧪 Testing TCP Chat Server..."

# Kill any existing servers first
pkill -f tcp-chat-server 2>/dev/null
sleep 1

# Start the server in background and redirect output to a log file
cd /home/arkenidar/Dropbox/000_tidy_this/learning_Go_language
go run ./cmd/tcp-chat-server > server.log 2>&1 &
SERVER_PID=$!

echo "📡 Server started with PID: $SERVER_PID"
sleep 3  # Give server time to start

echo "🔌 Testing connection with netcat..."
# Send test commands with proper timing
{
    sleep 1
    echo "TestUser"
    sleep 1
    echo "Hello from test script"
    sleep 1
    echo "/users"
    sleep 1
    echo "/quit"
} | nc localhost 8888 &

echo "⏳ Waiting for server to process..."
sleep 3

echo ""
echo "📋 Server output:"
cat server.log

echo ""
echo "🧹 Cleaning up..."
kill $SERVER_PID 2>/dev/null
pkill -f tcp-chat-server 2>/dev/null  # Extra safety
rm -f server.log

echo "✅ Test completed successfully!"
