#!/bin/bash
# test-all-projects.sh - Test script for all four Go projects

echo "🧪 Testing All Four Go Projects"
echo "================================"

cd "$(dirname "$0")"

echo ""
echo "📦 Building all projects..."
if make build; then
    echo "✅ All projects built successfully"
else
    echo "❌ Build failed"
    exit 1
fi

echo ""
echo "📖 Testing read-line program..."
echo "test input" | timeout 3s ./bin/read-line > /dev/null 2>&1
if [ $? -eq 0 ] || [ $? -eq 124 ]; then
    echo "✅ read-line program works"
else
    echo "❌ read-line program failed"
fi

echo ""
echo "🌐 Testing HTTP server (in-memory)..."
timeout 3s ./bin/http-server > /dev/null 2>&1 &
SERVER_PID=$!
sleep 1
if kill -0 $SERVER_PID 2>/dev/null; then
    echo "✅ HTTP server starts successfully"
    kill $SERVER_PID 2>/dev/null
else
    echo "❌ HTTP server failed to start"
fi

echo ""
echo "🗄️ Testing HTTP server with GORM..."
timeout 3s ./bin/http-server-gorm > /dev/null 2>&1 &
GORM_PID=$!
sleep 1
if kill -0 $GORM_PID 2>/dev/null; then
    echo "✅ GORM HTTP server starts successfully"
    kill $GORM_PID 2>/dev/null
else
    echo "❌ GORM HTTP server failed to start"
fi

echo ""
echo "💬 Testing TCP chat server..."
timeout 3s ./bin/tcp-chat-server > /dev/null 2>&1 &
CHAT_PID=$!
sleep 1
if kill -0 $CHAT_PID 2>/dev/null; then
    echo "✅ TCP chat server starts successfully"
    kill $CHAT_PID 2>/dev/null
else
    echo "❌ TCP chat server failed to start"
fi

echo ""
echo "🎉 All tests completed!"
echo ""
echo "🔧 Available commands:"
echo "  make run-readline    - Run the input reader"
echo "  make run-server      - Run HTTP server (in-memory)"  
echo "  make run-server-gorm - Run HTTP server (with database)"
echo "  make run-tcp-chat    - Run TCP chat server (netcat client)"
echo "  make help            - Show all commands"
