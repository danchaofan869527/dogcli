# DogCLI - 快速开始指南 🐕

## 📋 项目概述

DogCLI 是一个用 Go 语言开发的命令行工具，可以快速获取狗狗图片。支持随机图片、按品种获取、批量下载等功能。

**特点：**
- 🎲 随机获取狗狗图片
- 🐕 按品种搜索（支持 100+ 品种）
- 📥 批量下载（最多 50 张）
- 🌐 跨平台支持
- ⚡ 快速轻量

## 🚀 5 分钟快速开始

### 方式一：使用 Go 环境（推荐开发者）

```bash
# 1. 进入项目目录
cd dogcli

# 2. 安装依赖
go mod download

# 3. 构建
go build -o dogcli

# 4. 运行
./dogcli random --count 3
```

### 方式二：直接使用（无需安装 Go）

```bash
# 使用现有的安装脚本（需要先发布二进制文件）
curl -fsSL https://your-domain.com/install.sh | bash

# 然后直接使用
dogcli list
dogcli random
```

## 💡 使用示例

### 1. 查看所有可用命令
```bash
./dogcli --help
```

### 2. 获取随机狗狗图片
```bash
# 获取 1 张（默认）
./dogcli random

# 获取 5 张随机图片
./dogcli random --count 5

# 保存到指定目录
./dogcli random -c 10 -o ~/Pictures/dogs
```

### 3. 按品种获取图片
```bash
# 获取金毛图片
./dogcli breed golden-retriever

# 获取哈士奇图片（支持连字符）
./dogcli breed husky -c 3

# 获取柯基图片
./dogcli breed corgi
```

### 4. 列出所有品种
```bash
./dogcli list
```

### 5. 查看子品种
```bash
# 查看寻回犬的子品种
./dogcli breed retriever/golden
```

## 🛠️ 开发命令

```bash
# 构建
make build

# 清理
make clean

# 运行测试
make test

# 构建所有平台
make build-all

# 创建发布版本
make release
```

## 📂 项目文件说明

```
dogcli/
├── main.go           # 程序入口点
├── go.mod            # Go 模块依赖管理
├── cmd/              # 命令行命令定义
│   ├── root.go       # 根命令和基础设置
│   ├── random.go     # 随机图片命令
│   ├── breed.go      # 品种图片命令
│   └── list.go       # 列出品种命令
├── api/              # API 客户端
│   └── dog_api.go    # Dog API 集成和图片下载
├── install.sh        # 一键安装脚本
├── Makefile          # 构建和发布配置
├── README.md         # 完整文档
├── DEPLOYMENT.md     # 部署发布指南
└── QUICKSTART.md     # 本文件
```

## 🔧 技术栈

- **语言**: Go 1.21+
- **框架**: Cobra (命令行框架)
- **API**: Dog CEO API (https://dog.ceo/dog-api/)
- **构建**: Make + Go Build
- **安装**: Bash 脚本

## 🎯 下一步

1. **修改配置**: 编辑 `install.sh` 设置你的下载 URL
2. **构建发布**: 运行 `make build-all` 生成多平台二进制文件
3. **上传到 GitHub**: 创建 Release 并上传二进制文件
4. **测试安装**: `curl -fsSL https://your-repo/raw/main/install.sh | bash`
5. **分享给用户**: 让用户使用你的 CLI 工具

## ❓ 常见问题

### Q: 如何发布到公网？
A: 将二进制文件上传到 GitHub Releases 或 CDN，然后修改 `install.sh` 中的 `DOWNLOAD_URL`。

### Q: 支持哪些平台？
A: 支持 Linux、macOS 和 Windows 的 x86_64 和 ARM64 架构。

### Q: 如何添加新功能？
A: 在 `cmd/` 目录添加新的命令文件，在 `api/` 目录添加 API 调用逻辑。

### Q: 图片保存在哪里？
A: 默认保存在当前目录，可以使用 `--output` 参数指定目录。

## 📞 获取帮助

- 查看完整文档: `cat README.md`
- 查看部署指南: `cat DEPLOYMENT.md`
- 命令帮助: `./dogcli --help`
- 特定命令帮助: `./dogcli random --help`

## 🎉 开始使用

现在你已经准备好了！试试运行：

```bash
# 构建项目
go build -o dogcli

# 获取你的第一张狗狗图片
./dogcli random

# 查看所有支持的品种
./dogcli list
```

祝你使用愉快！🐕
