# Contributing to RPC Service

Thank you for your interest in contributing to RPC Service! This document provides guidelines and instructions for contributing.

## How to Contribute

1. **Fork the repository** - Click the "Fork" button at the top right of the page
2. **Clone your fork**:
   ```bash
   git clone https://github.com/YOUR_USERNAME/rpc-service.git
   cd rpc-service
   ```
3. **Create a feature branch**:
   ```bash
   git checkout -b feature/amazing-new-feature
   ```
4. **Make your changes** - Follow the coding standards below
5. **Run tests and checks**:
   ```bash
   make check
   ```
6. **Commit your changes**:
   ```bash
   git commit -m "Add amazing new feature"
   ```
7. **Push to your fork**:
   ```bash
   git push origin feature/amazing-new-feature
   ```
8. **Create a Pull Request** - Go to the original repository and click "New Pull Request"

## Coding Standards

- Follow Go coding conventions
- Run `make format` before committing
- Ensure all tests pass with `make test`
- Add tests for new functionality
- Update documentation as needed

## Code Style

- Use `gofumpt` for formatting (run `make format`)
- Use `golangci-lint` for linting (run `make lint`)
- Write clear, descriptive commit messages
- Keep pull requests focused and small

## Reporting Issues

When reporting issues, please include:
- A clear description of the problem
- Steps to reproduce the issue
- Expected behavior vs actual behavior
- Go version and operating system
- Any relevant error messages or logs

## Questions?

If you have questions, feel free to open an issue for discussion.
