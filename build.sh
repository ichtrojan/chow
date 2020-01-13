#!/bin/bash

# Build for 32-bit macOS
GOOS=darwin GOARCH=386 go build -o chow_Darwin_386 main.go

checksum=$(shasum -a 256 chow_Darwin_386)

echo "macOS 32-bit ✅"
echo $checksum

# Build for 64-bit macOS
GOOS=darwin GOARCH=amd64 go build -o chow_Darwin_x86_64 main.go

checksum=$(shasum -a 256 chow_Darwin_x86_64)

echo "macOS 64-bit ✅"
echo $checksum

# Build for 32-bit Linux
GOOS=linux GOARCH=386 go build -o chow_Linux_386 main.go

checksum=$(shasum -a 256 chow_Linux_386)

echo "Linux 32-bit ✅"
echo $checksum

# Build for 64-bit Linux
GOOS=linux GOARCH=amd64 go build -o chow_Linux_x86_64 main.go

checksum=$(shasum -a 256 chow_Linux_x86_64)

echo "Linux 64-bit ✅"
echo $checksum

echo "Chow built succesfully 🔥"