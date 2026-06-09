#!/bin/bash

# DogCLI Installation Script
# This script installs the dogcli binary to your system

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Configuration
REPO="dogcli/dogcli"
VERSION="v1.0.0"
BINARY_NAME="dogcli"

# Detect platform
OS="$(uname -s)"
ARCH="$(uname -m)"

case "$OS" in
    Linux*)
        OS=linux
        ;;
    Darwin*)
        OS=darwin
        ;;
    MINGW*|MSYS*|CYGWIN*)
        echo "❌ Windows not supported by this installer. Please download from releases page."
        exit 1
        ;;
    *)
        echo "❌ Unsupported OS: $OS"
        exit 1
        ;;
esac

case "$ARCH" in
    x86_64|amd64)
        ARCH=amd64
        ;;
    i386|i686)
        ARCH=386
        ;;
    arm64|aarch64)
        ARCH=arm64
        ;;
    *)
        echo "❌ Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

# Download URL (placeholder - replace with your actual download URL)
DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${VERSION}/dogcli-${OS}-${ARCH}"

# Installation directory
INSTALL_DIR="/usr/local/bin"

if [ ! -w "$INSTALL_DIR" ]; then
    INSTALL_DIR="$HOME/.local/bin"
fi

echo "🐕 Installing DogCLI..."
echo "Platform: ${OS}-${ARCH}"
echo "Install directory: ${INSTALL_DIR}"
echo ""

# Create installation directory if needed
mkdir -p "$INSTALL_DIR"

# Download binary
echo "⏳ Downloading binary..."
if command -v curl >/dev/null 2>&1; then
    curl -fsSL "$DOWNLOAD_URL" -o "${INSTALL_DIR}/${BINARY_NAME}"
elif command -v wget >/dev/null 2>&1; then
    wget -qO "${INSTALL_DIR}/${BINARY_NAME}" "$DOWNLOAD_URL"
else
    echo "❌ Neither curl nor wget is installed"
    exit 1
fi

# Make binary executable
chmod +x "${INSTALL_DIR}/${BINARY_NAME}"

echo ""
echo "✅ DogCLI installed successfully!"
echo ""
echo "To get started:"
echo "  dogcli --help"
echo "  dogcli random"
echo "  dogcli list"
echo ""
echo "Installation location: ${INSTALL_DIR}/${BINARY_NAME}"
echo ""

# Check if installation directory is in PATH
if [[ ":$PATH:" != *":${INSTALL_DIR}:"* ]]; then
    echo "⚠️  ${INSTALL_DIR} is not in your PATH"
    echo "Add the following to your ~/.bashrc or ~/.zshrc:"
    echo "  export PATH=\"${INSTALL_DIR}:\$PATH\""
fi
