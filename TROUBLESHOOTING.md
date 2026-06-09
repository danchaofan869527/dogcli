# DogCLI 构建问题解决方案

## 问题描述

执行 `go build -o dogcli` 时出现错误：
```
cmd/breed.go:7:2: missing go.sum entry for module providing package github.com/spf13/cobra
```

## 原因

项目缺少 `go.sum` 文件，这是 Go 模块的依赖校验文件。

## 解决方案

### 方法一：自动修复（推荐）

在项目目录运行以下命令：

```bash
cd /Users/fudongxin/Desktop/x-project/DogCLI

# 下载依赖并生成 go.sum
go mod tidy

# 再次构建
go build -o dogcli
```

### 方法二：使用设置脚本

```bash
cd /Users/fudongxin/Desktop/x-project/DogCLI
chmod +x setup.sh
./setup.sh
```

### 方法三：手动下载依赖

```bash
cd /Users/fudongxin/Desktop/x-project/DogCLI

# 下载依赖
go mod download

# 整理依赖
go mod tidy

# 构建
go build -o dogcli
```

## 前置条件

### 1. 安装 Go 环境

检查是否安装了 Go：
```bash
go version
```

如果未安装，请按以下方式安装：

**macOS:**
```bash
brew install go
```

**Linux (Ubuntu/Debian):**
```bash
sudo apt update
sudo apt install golang-go
```

**从官网下载:**
访问 https://golang.org/dl/ 下载安装包

### 2. 验证 Go 安装

```bash
go version
# 应该输出类似: go version go1.21.x darwin/amd64
```

## 验证构建成功

构建成功后，应该能看到：

```bash
$ ls -lh dogcli
-rwxr-xr-x  1 user  staff   2.5M Jun  9 10:15 dogcli

$ ./dogcli --help
🐕 A CLI tool to fetch dog images

Usage:
  dogcli [command]

Available Commands:
  breed        Get dog images by breed
  help         Help about any command
  list         List all available dog breeds
  random       Get random dog images

Flags:
  -h, --help   help for dogcli

Use "dogcli [command] --help" for more information about a command.
```

## 常见问题

### Q: go mod tidy 执行很慢？
A: 首次运行需要从网络下载依赖包，请耐心等待。可以设置 GOPROXY 加速：
```bash
# 使用国内镜像
export GOPROXY=https://goproxy.cn,direct
go mod tidy
```

### Q: 网络超时怎么办？
A: 使用 Go 代理镜像：
```bash
export GOPROXY=https://goproxy.cn,direct
export GO111MODULE=on
```

### Q: 权限错误？
A: 确保在项目目录有写入权限：
```bash
cd /Users/fudongxin/Desktop/x-project/DogCLI
chmod u+w .
```

## 下一步

构建成功后：

```bash
# 测试基本功能
./dogcli random --count 3

# 列出所有品种
./dogcli list

# 获取特定品种
./dogcli breed corgi -c 2

# 安装到系统（可选）
sudo cp dogcli /usr/local/bin/
```

## 需要帮助？

如果问题仍未解决，请检查：

1. Go 版本是否 >= 1.21
2. 网络连接是否正常
3. 项目目录是否正确
4. 文件权限是否正确
