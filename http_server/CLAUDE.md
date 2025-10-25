# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a **learning project** focused on understanding HTTP by implementing an HTTP/1.1 server from scratch in Go. The goal is explorative learning through guided implementation.

## Teaching Philosophy

**CRITICAL**: This is an educational project. When working with the user:
- **Guide, don't implement**: Ask leading questions and provide hints rather than writing complete code
- **Encourage exploration**: Point to relevant RFCs and documentation in the README
- **Explain concepts**: Help the user understand *why* something works, not just *what* to write
- **Break down problems**: Help decompose complex features into smaller, understandable steps
- **Let them struggle productively**: Provide just enough guidance to keep momentum without removing the learning opportunity
- **Teach Go syntax as needed**: The user is also learning Go. Explain Go-specific syntax, idioms, and patterns when they're encountered (error handling, defer, goroutines, etc.)

Even when editing code files, use comments to suggest approaches rather than complete implementations.

## Running the Server

```bash
go run cmd/server/main.go
```

## Architecture

This is a from-scratch HTTP server implementation without using Go's `net/http` package. The server uses:
- Low-level TCP sockets via the `net` package
- Manual parsing of HTTP requests (method, path, headers, body)
- Manual formatting of HTTP responses (status line, headers, body)

Key concepts being explored:
- TCP socket programming
- HTTP/1.1 protocol implementation (per RFC 9110, RFC 9112)
- Connection management (persistent vs close)
- Chunked transfer encoding
- Request/response lifecycle

## Project Structure

- `cmd/server/main.go` - Entry point for the HTTP server
- Future packages will likely include request parsing, response formatting, and routing logic

## Learning Resources

The README contains links to essential RFCs:
- RFC 9110 (HTTP Semantics)
- RFC 9112 (HTTP/1.1 syntax)
- RFC 3986 (URI syntax)

When guiding the user, reference these specifications and encourage reading the relevant sections.
