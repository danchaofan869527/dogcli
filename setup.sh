#!/bin/bash

# DogCLI Setup Script
# 此脚本用于初始化项目依赖

set -e

echo "🐕 DogCLI 项目设置"
echo "=================="

# 检查 Go 是否安装
if ! command -v go &> /dev/null; then
    echo "❌ 错误: 未找到 Go 环境"
    echo ""
    echo "请先安装 Go:"
    echo "  macOS:   brew install go"
    echo "  Ubuntu:  sudo apt install golang-go"
    echo "  访问:    https://golang.org/dl/"
    exit 1
fi

echo "✅ 检测到 Go 版本: $(go version)"
echo ""

# 进入项目目录
cd "$(dirname "$0")"

echo "📦 正在下载依赖..."
go mod tidy

echo ""
echo "🔨 正在构建项目..."
go build -o dogcli

echo ""
echo "✅ 设置完成！"
echo ""
echo "现在可以运行:"
echo "  ./dogcli --help"
echo "  ./dogcli random"
echo "  ./dogcli list"
echo ""
