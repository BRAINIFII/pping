#!/bin/bash

OUTPUT_DIR="dist"

echo "🧹 Cleaning and creating $OUTPUT_DIR directory..."
rm -rf "$OUTPUT_DIR"
mkdir -p "$OUTPUT_DIR"

echo "🔥 Building pong for all major platforms..."

# Windows (64-bit)
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o "$OUTPUT_DIR/pong.exe" pong.go

# Linux (64-bit)
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o "$OUTPUT_DIR/pong-linux" pong.go

# macOS (Intel)
GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o "$OUTPUT_DIR/pong-mac" pong.go

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o "$OUTPUT_DIR/pong-mac-arm" pong.go

# Raspberry Pi / ARMv7
GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=0 go build -o "$OUTPUT_DIR/pong-armv7" pong.go

echo "✅ All builds done! Check the '$OUTPUT_DIR/' directory 🗂️"
