#!/bin/bash

echo "🌱 Testing seed compilation..."
cd /Users/asdfghjkl/project/betaTasker/backer/godotask

# コンパイルチェックのみ
echo "Checking compilation..."
go build -o /tmp/test_seed cmd/seed/main.go

if [ $? -eq 0 ]; then
    echo "✅ Compilation successful!"
    echo "You can now run the seed with:"
    echo "  go run cmd/seed/main.go"
    echo "  go run cmd/seed/main.go -clean"
else
    echo "❌ Compilation failed!"
    exit 1
fi