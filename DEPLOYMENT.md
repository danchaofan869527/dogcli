# DogCLI 部署和发布指南

## 开发环境设置

### 1. 安装 Go

确保安装了 Go 1.21 或更高版本：

```bash
# macOS
brew install go

# Linux (Ubuntu/Debian)
sudo apt update
sudo apt install golang-go

# 验证安装
go version
```

### 2. 项目结构

```
dogcli/
├── cmd/                  # 命令行命令定义
│   ├── root.go          # 根命令
│   ├── random.go        # 随机图片命令
│   ├── breed.go         # 品种图片命令
│   └── list.go          # 列出品种命令
├── api/                  # API 客户端
│   └── dog_api.go       # Dog API 集成
├── main.go              # 程序入口
├── go.mod               # Go 模块依赖
├── install.sh           # 安装脚本
├── Makefile             # 构建配置
└── README.md            # 项目文档
```

## 本地开发

### 安装依赖

```bash
go mod download
```

### 构建项目

```bash
# 简单构建
go build -o dogcli

# 带版本信息构建
VERSION=v1.0.0 go build -ldflags "-X main.version=${VERSION}" -o dogcli

# 使用 Makefile
make build
```

### 测试

```bash
# 本地测试
./dogcli --help
./dogcli list
./dogcli random --count 3
./dogcli breed corgi --count 2
```

## 发布流程

### 1. 多平台构建

```bash
# 构建所有平台
make build-all

# 手动构建特定平台
GOOS=linux GOARCH=amd64 go build -o dogcli-linux-amd64 .
GOOS=darwin GOARCH=arm64 go build -o dogcli-darwin-arm64 .
GOOS=windows GOARCH=amd64 go build -o dogcli-windows-amd64.exe .
```

### 2. 创建 GitHub Release

1. 在 GitHub 创建新的 Release
2. 上传编译好的二进制文件到 Release
3. 更新 install.sh 中的 VERSION 变量

### 3. 发布到 GitHub Pages

将 install.sh 放在 GitHub 仓库中，用户可以直接运行：

```bash
curl -fsSL https://raw.githubusercontent.com/yourusername/dogcli/main/install.sh | bash
```

## CDN 部署选项

### 使用 GitHub Releases 作为 CDN

```bash
# 在 install.sh 中设置正确的下载 URL
DOWNLOAD_URL="https://github.com/dogcli/dogcli/releases/download/${VERSION}/dogcli-${OS}-${ARCH}"
```

### 使用对象存储 (阿里云 OSS/AWS S3)

```bash
# 构建完成后上传到 OSS
# 示例：使用阿里云 OSS
ossutil cp dogcli-linux-amd64 oss://your-bucket/dogcli/v1.0.0/
```

### 修改 install.sh 使用 CDN

```bash
DOWNLOAD_URL="https://cdn.yourdomain.com/dogcli/${VERSION}/dogcli-${OS}-${ARCH}"
```

## 安全建议

1. **代码签名**：对二进制文件进行签名
   ```bash
   # macOS
   codesign -s "Developer ID Application: Your Name" dogcli
   ```

2. **校验和验证**：提供 SHA256 校验和
   ```bash
   shasum -a 256 dogcli-linux-amd64 > dogcli-linux-amd64.sha256
   ```

3. **HTTPS**：确保所有下载使用 HTTPS
4. **最小权限**：避免需要 root 权限安装

## 监控和日志

### 添加使用统计（可选）

可以在 API 调用时添加匿名统计：

```go
// 在 api/dog_api.go 中
func trackUsage(command string) {
    // 发送匿名统计数据
    go func() {
        http.Get("https://your-analytics.com/track?cmd=" + command)
    }()
}
```

## 故障排查

### 常见问题

1. **网络问题**
   - API 可能响应慢，添加重试机制
   - 提供离线模式选项

2. **权限问题**
   - 某些系统需要 sudo 安装到 /usr/local/bin
   - 建议安装到 ~/.local/bin

3. **平台兼容性**
   - 测试不同 Linux 发行版
   - 处理不同的 musl/glibc 版本

### 调试模式

```bash
# 添加调试标志
./dogcli random --debug
```

## 版本管理

使用语义化版本：

- **v1.0.0** - 初始发布
- **v1.0.1** - 补丁修复
- **v1.1.0** - 新功能
- **v2.0.0** - 破坏性变更

## 持续集成 (CI/CD)

### GitHub Actions 示例

```yaml
name: Build and Release

on:
  push:
    tags:
      - 'v*'

jobs:
  build:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        goos: [linux, darwin, windows]
        goarch: [amd64, arm64]
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      - run: go build -o dogcli-${{ matrix.goos }}-${{ matrix.goarch }}
      - uses: actions/upload-artifact@v3
        with:
          name: dogcli-${{ matrix.goos }}-${{ matrix.goarch }}
```

## 推广和使用

### 1. 文档完善

- 详细的 README
- 使用示例
- FAQ 部分

### 2. 社区建设

- GitHub Discussions
- Issues 模板
- 贡献指南

### 3. 分发渠道

- Homebrew (macOS/Linux)
- Scoop (Windows)
- Snap (Linux)
- AUR (Arch Linux)

## 下一步扩展功能

1. **图片过滤**：按大小、颜色过滤
2. **收藏功能**：保存喜欢的图片 URL
3. **幻灯片模式**：连续显示图片
4. **API 缓存**：减少网络请求
5. **多语言支持**：i18n 国际化

## 许可证

MIT License - 允许自由使用和修改。
