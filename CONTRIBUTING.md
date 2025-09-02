# Contributing to Juice WRLD API Wrapper (Go)

Thank you for your interest in contributing to the Juice WRLD API Wrapper! This document provides guidelines and information for contributors.

## Code of Conduct

This project follows the [Contributor Covenant Code of Conduct](https://www.contributor-covenant.org/). By participating, you agree to uphold this code.

## Getting Started

### Prerequisites

- Go 1.22 or later
- Git
- Basic understanding of Go and REST APIs

### Setting Up Development Environment

1. Fork the repository on GitHub
2. Clone your fork locally:
   ```bash
   git clone https://github.com/your-username/juicewrld-api-wrapper.git
   cd juicewrld-api-wrapper/go
   ```
3. Ensure you have the latest version:
   ```bash
   git remote add upstream https://github.com/hackinhood/juicewrld-api-wrapper-go.git
   git fetch upstream
   git checkout main
   git merge upstream/main
   ```

## Development Guidelines

### Code Style

- Follow Go's standard formatting: `go fmt`
- Use `go vet` to check for common mistakes
- Follow Go naming conventions
- Write clear, self-documenting code
- Add comments for exported functions and types

### Testing

- All new features must include tests
- Run the CLI test suite to verify functionality:
  ```bash
  go run cmd/main.go all
  ```
- Test individual functions as needed:
  ```bash
  go run cmd/main.go get-artist 1
  go run cmd/main.go search-songs "test query"
  ```

### Error Handling

- Use the existing error types when appropriate
- Provide meaningful error messages
- Handle context cancellation properly
- Follow Go's error handling conventions

### Dependencies

- **No external dependencies allowed** - use only Go standard library
- If you need functionality not in the standard library, consider if it's truly necessary
- Document any new standard library imports

## Pull Request Process

### Before Submitting

1. **Test your changes thoroughly**
   ```bash
   go run cmd/main.go all
   ```

2. **Check code quality**
   ```bash
   go fmt ./...
   go vet ./...
   ```

3. **Update documentation** if you've added new features or changed existing behavior

### Submitting a Pull Request

1. Create a feature branch from `main`:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. Make your changes and commit them:
   ```bash
   git add .
   git commit -m "Add: brief description of your changes"
   ```

3. Push your branch:
   ```bash
   git push origin feature/your-feature-name
   ```

4. Open a Pull Request on GitHub

### Pull Request Guidelines

- **Title**: Use a clear, descriptive title
- **Description**: Explain what your PR does and why
- **Testing**: Describe how you tested your changes
- **Breaking Changes**: Clearly mark any breaking changes
- **Documentation**: Update relevant documentation

### Commit Message Format

Use clear, descriptive commit messages:

```
Add: new feature description
Fix: bug description
Update: change description
Remove: removal description
Docs: documentation update
Test: test addition/update
```

## Types of Contributions

### Bug Reports

When reporting bugs, please include:

- Go version
- Operating system
- Steps to reproduce
- Expected vs actual behavior
- Error messages (if any)

### Feature Requests

For new features, please:

- Check existing issues first
- Provide a clear use case
- Explain why this feature would be valuable
- Consider if it fits the project's scope

### Code Contributions

We welcome contributions for:

- Bug fixes
- Performance improvements
- New API endpoint support
- Documentation improvements
- Test coverage improvements

## API Compatibility

- Maintain backward compatibility when possible
- If breaking changes are necessary, document them clearly
- Follow semantic versioning principles
- Update version numbers appropriately

## Documentation

- Update README.md for user-facing changes
- Add/update code comments for new functions
- Update API documentation
- Include examples for new features

## Release Process

Releases are managed by maintainers. When your PR is merged:

1. Maintainers will update version numbers
2. A new release will be created on GitHub
3. The release will be tagged appropriately

## Getting Help

- **Issues**: Use GitHub Issues for bug reports and feature requests
- **Discussions**: Use GitHub Discussions for questions and general discussion
- **API Documentation**: Check [juicewrldapi.com](https://juicewrldapi.com) for API details

## Recognition

Contributors will be recognized in:
- GitHub contributors list
- Release notes (for significant contributions)

Thank you for contributing to the Juice WRLD community! 🎵
