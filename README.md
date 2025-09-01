# Path Size Utility

A command-line utility written in Go that calculates and displays the size of files and directories with various formatting options.

[![Go Report Card](https://goreportcard.com/badge/github.com/Konstantin-Gromakovskiy/go-project-242)](https://goreportcard.com/report/github.com/Konstantin-Gromakovskiy/go-project-242)
[![Actions Status](https://github.com/Konstantin-Gromakovskiy/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/Konstantin-Gromakovskiy/go-project-242/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/Konstantin-Gromakovskiy/go-project-242.svg)](https://pkg.go.dev/github.com/Konstantin-Gromakovskiy/go-project-242)

## Features

- Calculate file and directory sizes
- Recursive directory size calculation
- Human-readable output formatting
- Support for hidden files and directories
- Cross-platform compatibility

## Installation

### Prerequisites

- Go 1.24 or higher
- Git (for installation from source)

### From Source

```bash
git clone https://github.com/Konstantin-Gromakovskiy/go-project-242.git
cd go-project-242
go build -o path-size ./cmd/hexlet-path-size
```

### Using Go Install

```bash
go install github.com/Konstantin-Gromakovskiy/go-project-242/cmd/hexlet-path-size@latest
```

## Usage

```
path-size [options] <path>

Options:
  -a, --all         Include hidden files and directories
  -h, --help        Show help
  -H, --human       Print sizes in human-readable format (e.g., 1K 234M 2G)
  -r, --recursive   Recurse into directories
```

### Examples

Basic usage:
```bash
path-size /path/to/file.txt
```

Human-readable output:
```bash
path-size -H /path/to/directory
```

Include hidden files:
```bash
path-size -a /path/to/directory
```

Recursive directory size:
```bash
path-size -r /path/to/directory
```

## Dependencies

- [urfave/cli/v3](https://github.com/urfave/cli) - Command line interface library
- [stretchr/testify](https://github.com/stretchr/testify) - Testing toolkit

## Development

### Building

```bash
make build
```

### Testing

Run all tests:
```bash
make test
```

Run tests with coverage:
```bash
make test-coverage
```

### Linting

```bash
make lint
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.