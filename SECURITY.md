# Security Policy

## Supported Versions

We currently support the latest stable release of the RPC Service. Security updates are only provided for the most recent version.

| Version | Supported |
|---------|-----------|
| Latest  | ✅ Supported |
| Older   | ❌ Not supported |

## Reporting a Vulnerability

We take the security of RPC Service seriously. If you believe you have found a security vulnerability, please report it responsibly.

### How to Report

1. **Do NOT** open a public issue
2. Email the security team at: wangdw2012@gmail.com
3. Include:
   - Description of the vulnerability
   - Steps to reproduce
   - Potential impact
   - Any exploit code (if available)

### What to Expect

- We will acknowledge your report within 24-48 hours
- We will keep you updated on the progress
- Once the vulnerability is fixed, you will be credited (if you wish)

## Security Best Practices

- Keep your Go installation updated
- Review dependencies regularly
- Use environment variables for sensitive configuration
- Follow the principle of least privilege

## Dependencies

We use `govulncheck` to scan for vulnerabilities in dependencies. Run:
```bash
go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck ./...
```
