# DogCLI 🐕

A command-line tool to fetch random dog images or images by breed from the Dog API.

## Features

- 🎲 Get random dog images
- 🐕 Fetch images by breed
- 📋 List all available breeds
- 📥 Download images with custom output directory
- ⚡ Fast and lightweight
- 🌟 Cross-platform support

## Installation

### Quick Install (Linux & macOS)

```bash
curl -fsSL https://raw.githubusercontent.com/danchaofan869527/dogcli/main/install.sh | bash
```

### Manual Install

1. Download the binary for your platform from [Releases](https://github.com/danchaofan869527/dogcli/releases)
2. Make it executable: `chmod +x dogcli`
3. Move to your PATH: `mv dogcli /usr/local/bin/`

## Usage

### Get Random Dog Images

```bash
# Get 1 random dog image
dogcli random

# Get 5 random dog images
dogcli random --count 5

# Save to custom directory
dogcli random --output ~/Pictures/dogs
```

### Get Images by Breed

```bash
# Get images for a specific breed
dogcli breed golden-retriever

# Get multiple images
dogcli breed husky --count 3

# List sub-breeds
dogcli breed "retriever/golden"
```

### List Available Breeds

```bash
dogcli list
```

## Options

- `--count, -c`: Number of images to fetch (1-50, default: 1)
- `--output, -o`: Output directory for downloaded images (default: current directory)

## Examples

```bash
# Get 10 random dog images to a specific folder
dogcli random --count 10 --output ~/my-dogs

# Get corgi images
dogcli breed corgi --count 5

# See all available breeds
dogcli list
```

## Building from Source

```bash
# Clone the repository
git clone https://github.com/danchaofan869527/dogcli.git
cd dogcli

# Install dependencies
go mod download

# Build
go build -o dogcli

# Install
go install
```

## Development

### Project Structure

```
dogcli/
├── cmd/              # Command definitions
│   ├── root.go       # Root command
│   ├── random.go     # Random command
│   ├── breed.go      # Breed command
│   └── list.go       # List command
├── api/              # API client
│   └── dog_api.go    # Dog API integration
├── main.go           # Entry point
└── go.mod            # Go module file
```

### Running Tests

```bash
go test ./...
```

## API

This tool uses the [Dog API](https://dog.ceo/dog-api/) - a free, open API for dog images.

## License

MIT License - feel free to use this project for any purpose.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Support

If you encounter any issues or have questions, please [open an issue](https://github.com/danchaofan869527/dogcli/issues).
