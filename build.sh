#!/bin/bash

# macOS (Apple Silicon)
echo "Building for macOS (Apple Silicon)..."
GOOS=darwin GOARCH=arm64 go build -o KeyEvent-arm

# Windows
echo "Building for Windows..."
GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 go build -o KeyEvent.exe
