# Code Cloner

![CI](https://github.com/Qyroxen/Code-Cloner/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Code-Cloner/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Code-Cloner?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Code-Cloner)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Code-Cloner)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Code-Cloner?style=social)](https://github.com/Qyroxen/Code-Cloner/stargazers)

## What is it?

Code Cloner is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Code-Cloner.git
cd Code-Cloner
go build -o codecloner .

# Run
./codecloner --help
```

## CLI Usage

```bash
# Basic usage
./codecloner

# With flags
./codecloner --verbose --output json

# Get help
./codecloner --help
```

## Examples

```bash
# Example 1
./codecloner example1

# Example 2
./codecloner example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o codecloner .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Code-Cloner/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Code-Cloner?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Code-Cloner/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Code-Cloner?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Code-Cloner/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Code-Cloner" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Code-Cloner/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Code-Cloner" alt="Pull Requests">
  </a>
</p>
