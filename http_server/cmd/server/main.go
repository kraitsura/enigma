package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

type Request struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    []byte
}

type Response struct {
	StatusCode int
	StatusText string
	Headers    map[string]string
	Body       string
}

func parseHeaders(lines []string) map[string]string { //Parse Headers from array

	headers := make(map[string]string)
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			break
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			headers[parts[0]] = strings.TrimSpace(parts[1])
		}
	}

	return headers
}

func parseRequest(buffer []byte, n int) (*Request, error) { //Parse raw bytes into Request struct

	fmt.Println(string(buffer[:n]))
	lines := strings.Split(string(buffer[:n]), "\r\n")

	firstLine := lines[0]

	headers := parseHeaders(lines)

	parts := strings.Split(firstLine, " ")

	if len(parts) < 3 {
		return nil, fmt.Errorf("malformed request line: %s", firstLine)
	}

	method := parts[0]
	path := parts[1]
	ver := parts[2]

	req := &Request{
		Method:  method,
		Path:    path,
		Version: ver,
		Headers: headers,
	}

	return req, nil
}

func handleRequest(req *Request) *Response { //Take a Request, return a Response (routing logic)
	method := req.Method
	route := req.Path

	var statusCode int
	var statusText string
	var responseBody string

	switch method {
	case "GET":
		switch route {
		case "/":
			statusCode = 200
			statusText = "OK"
			responseBody = "Home Route"
		case "/hello":
			statusCode = 200
			statusText = "OK"
			responseBody = "Hello Route"
		default:
			statusCode = 404
			statusText = "Not Found"
			responseBody = "404! Route not found"
		}
	case "PUT":
		statusCode = 405
		statusText = "Method Not Allowed"
		responseBody = "PUT method not supported"
	case "POST":
		statusCode = 405
		statusText = "Method Not Allowed"
		responseBody = "POST method not supported"
	case "DELETE":
		statusCode = 405
		statusText = "Method Not Allowed"
		responseBody = "DELETE method not supported"
	default:
		statusCode = 405
		statusText = "Method Not Allowed"
		responseBody = "Unknown method not supported"
	}

	resp := &Response{
		StatusCode: statusCode,
		StatusText: statusText,
		Headers:    make(map[string]string), // empty for now
		Body:       responseBody,
	}

	return resp
}

func sendResponse(conn net.Conn, resp *Response) error { //Write Response to connection
	contentLength := len(resp.Body)
	statusLine := fmt.Sprintf("HTTP/1.1 %d %s", resp.StatusCode, resp.StatusText)
	response := fmt.Sprintf("%s\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", statusLine, contentLength, resp.Body)
	_, err := conn.Write([]byte(response))
	return err
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	shouldClose := false
	buffer := make([]byte, 1024)

	for { // read the curl req

		n, err := conn.Read(buffer)

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("Error while reading buffer:", err)
			break
		}

		req, respErr := parseRequest(buffer, n)

		if respErr != nil {
			log.Println(respErr)
			resp := &Response{
				StatusCode: 400,
				StatusText: "Bad Request",
				Body:       "Invalid request format",
				Headers:    make(map[string]string), // empty for now
			}
			sendResponse(conn, resp)
			continue
		}

		//Validate Host header
		if req.Headers["Host"] == "" {
			resp := &Response{
				StatusCode: 400,
				StatusText: "Bad Request",
				Body:       "Missing Host header",
				Headers:    make(map[string]string),
			}
			sendResponse(conn, resp)
			continue
		}

		resp := handleRequest(req)

		if req.Headers["Connection"] == "close" {
			shouldClose = true
			resp.Headers["Connection"] = "close"
		}
		fmt.Println("Connected! Accepted from:", conn.RemoteAddr())
		if err := sendResponse(conn, resp); err != nil {
			log.Println("Error sending a response:", err)
			break
		}

		if shouldClose {
			break
		}
	}
}

// ---------------------------------------------------------------------//

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		listener, err = net.Listen("tcp", "localhost:0")
		if err != nil {
			log.Fatal("Failed to listen on any port:", err)
		}
	}

	fmt.Println("Succesfully Connected, listening on", listener.Addr())

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println("Could not accept connection", err)
			continue
		}

		go handleConnection(conn)
	}
}

// ---------------------------------------------------------------------//
