package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

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

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Lets create buffer and read the curl req
	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)

	if err != nil {
		log.Println("Error while reading buffer:", err)
		return
	}

	fmt.Println(string(buffer[:n]))

	lines := strings.Split(string(buffer[:n]), "\r\n")

	firstLine := lines[0]

	headers := parseHeaders(lines)
	//Validate Host header
	if headers["Host"] == "" {
		statusLine := "HTTP/1.1 400 Bad Request"
		responseBody := "Missing Host header"
		contentLength := len(responseBody)
		response := fmt.Sprintf("%s\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", statusLine, contentLength, responseBody)
		conn.Write([]byte(response))
		return
	}

	parts := strings.Split(firstLine, " ")

	if len(parts) < 3 {
		log.Println("Malformed request line:", firstLine)
		return
	}
	method := parts[0]
	route := parts[1]

	var responseBody string
	var statusLine string

	switch method {
	case "GET":
		switch route {
		case "/":
			statusLine = "HTTP/1.1 200 OK"
			responseBody = "Home Route"
		case "/hello":
			statusLine = "HTTP/1.1 200 OK"
			responseBody = "Hello Route"
		default:
			statusLine = "HTTP/1.1 404 Not Found"
			responseBody = "404! Route not found"
		}
	case "PUT":
		statusLine = "HTTP/1.1 405 Method Not Allowed"
		responseBody = "PUT method not supported"
	case "POST":
		statusLine = "HTTP/1.1 405 Method Not Allowed"
		responseBody = "POST method not supported"
	case "DELETE":
		statusLine = "HTTP/1.1 405 Method Not Allowed"
		responseBody = "DELETE method not supported"
	default:
		statusLine = "HTTP/1.1 405 Method Not Allowed"
		responseBody = "Unknown method not supported"
	}

	fmt.Println("Connected Accepted from:", conn.RemoteAddr())
	contentLength := len(responseBody)
	response := fmt.Sprintf("%s\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", statusLine, contentLength, responseBody)
	conn.Write([]byte(response))

}

func parseHeaders(lines []string) map[string]string {

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
