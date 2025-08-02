# Security Policy

## Supported Versions

We take security seriously. The following versions of HITMAN are currently supported with security updates:

| Version | Supported          |
| ------- | ------------------ |
| 2.0.x   | :white_check_mark: |
| 1.x.x   | :x:                |

## Reporting a Vulnerability

If you discover a security vulnerability within HITMAN, please send an email to security@hitman-cli.dev instead of using the issue tracker. All security vulnerabilities will be promptly addressed.

Please include the following information in your report:

- Type of issue (e.g. buffer overflow, SQL injection, cross-site scripting, etc.)
- Full paths of source file(s) related to the manifestation of the issue
- The location of the affected source code (tag/branch/commit or direct URL)
- Any special configuration required to reproduce the issue
- Step-by-step instructions to reproduce the issue
- Proof-of-concept or exploit code (if possible)
- Impact of the issue, including how an attacker might exploit the issue

## Security Considerations

HITMAN is designed to terminate processes on your system, which inherently requires elevated permissions in some cases. Please be aware:

### Process Termination
- HITMAN uses system commands (`kill`, `taskkill`) to terminate processes
- On Unix systems, you may need appropriate permissions to kill processes not owned by your user
- Always verify the target before elimination

### Command Injection Prevention
- All port numbers are validated as integers
- System commands use parameterized execution to prevent injection
- User input is sanitized before being passed to system commands

### Privilege Escalation
- HITMAN does not require or request elevated privileges
- It operates within the permissions of the executing user
- Some processes may require elevated permissions to terminate

## Best Practices

When using HITMAN:

1. **Verify targets**: Always confirm you're targeting the correct port/process
2. **Use confirmation prompts**: Don't always use `--force` flag
3. **Regular updates**: Keep HITMAN updated to the latest version
4. **Principle of least privilege**: Run with minimal necessary permissions
5. **Audit usage**: Be aware of what processes you're terminating

## Responsible Disclosure

We follow responsible disclosure practices:

1. Report received and acknowledged within 24 hours
2. Initial assessment within 72 hours
3. Regular updates on progress
4. Public disclosure after fix is available (coordinated timing)

Thank you for helping keep HITMAN and its users safe! 🛡️