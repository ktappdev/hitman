# Contributing to HITMAN

First off, thanks for taking the time to contribute! 🎯

The following is a set of guidelines for contributing to HITMAN, the elite process terminator. These are mostly guidelines, not rules. Use your best judgment, and feel free to propose changes to this document in a pull request.

## Code of Conduct

This project and everyone participating in it is governed by our commitment to creating a welcoming and inclusive environment. By participating, you are expected to uphold professional standards.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check the existing issues as you might find out that you don't need to create one. When you are creating a bug report, please include as many details as possible:

- **Use a clear and descriptive title**
- **Describe the exact steps to reproduce the problem**
- **Provide specific examples to demonstrate the steps**
- **Describe the behavior you observed and what behavior you expected**
- **Include your operating system and Go version**

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion, please include:

- **Use a clear and descriptive title**
- **Provide a step-by-step description of the suggested enhancement**
- **Provide specific examples to demonstrate the steps**
- **Describe the current behavior and explain the behavior you expected**
- **Explain why this enhancement would be useful**

### Pull Requests

- Fill in the required template
- Do not include issue numbers in the PR title
- Include screenshots and animated GIFs in your pull request whenever possible
- Follow the Go coding standards
- Include thoughtfully-worded, well-structured tests
- Document new code based on the Documentation Styleguide
- End all files with a newline

## Development Setup

1. Fork the repo
2. Clone your fork: `git clone https://github.com/ktappdev/hitman.git`
3. Create a feature branch: `git checkout -b feature/amazing-feature`
4. Install dependencies: `go mod tidy`
5. Make your changes
6. Test your changes: `go test ./...`
7. Build and test: `go build -o hitman && ./hitman --help`
8. Commit your changes: `git commit -m 'Add amazing feature'`
9. Push to the branch: `git push origin feature/amazing-feature`
10. Open a Pull Request

## Styleguides

### Git Commit Messages

- Use the present tense ("Add feature" not "Added feature")
- Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
- Limit the first line to 72 characters or less
- Reference issues and pull requests liberally after the first line

### Go Styleguide

- Follow the official [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `gofmt` to format your code
- Use meaningful variable and function names
- Add comments for exported functions and types
- Keep functions small and focused

### Documentation Styleguide

- Use [Markdown](https://daringfireball.net/projects/markdown/)
- Keep line length to 80 characters when possible
- Use clear, concise language
- Include code examples where appropriate

## Testing

- Write tests for new functionality
- Ensure all tests pass before submitting PR
- Test on multiple platforms when possible (macOS, Linux, Windows)
- Include both unit tests and integration tests where appropriate

## Questions?

Don't hesitate to ask questions by opening an issue with the "question" label.

Thanks for contributing! 🎯💥