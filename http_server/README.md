# HTTP Server from Scratch

A learning project to understand HTTP by implementing a server from scratch.

## Running the server

```bash
go run cmd/server/main.go
```

## HTTP Specifications & Resources

### Core HTTP/1.1 RFCs
- [RFC 9110 - HTTP Semantics](https://www.rfc-editor.org/rfc/rfc9110.html) - Core HTTP semantics
- [RFC 9112 - HTTP/1.1](https://www.rfc-editor.org/rfc/rfc9112.html) - Message syntax and routing
- [RFC 3986 - URI Generic Syntax](https://www.rfc-editor.org/rfc/rfc3986.html) - Understanding URIs

### Key Concepts to Explore
- TCP socket programming in Go (`net` package)
- HTTP request parsing (method, path, version, headers, body)
- HTTP response formatting (status line, headers, body)
- Connection management (persistent vs close)
- Chunked transfer encoding
- Request methods (GET, POST, PUT, DELETE, etc.)
- Status codes (2xx, 3xx, 4xx, 5xx)
- Common headers (Content-Type, Content-Length, Host, etc.)
