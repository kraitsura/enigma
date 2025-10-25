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

	// Lets create buffer and read the curl req
	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)

	if err != nil {
		log.Println("Error while reading buffer:", err)
		conn.Close()
	}

	fmt.Println(string(buffer[:n]))

	lines := strings.Split(string(buffer[:n]), "\r\n")

	firstLine := lines[0]

	route := strings.Split(firstLine, " ")[1]
	var responseBody string
	var statusLine string
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

	fmt.Println("Connected Accepted from:", conn.RemoteAddr())
	contentLength := len(responseBody)
	response := fmt.Sprintf("%s\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", statusLine, contentLength, responseBody)
	conn.Write([]byte(response))

	conn.Close()
}
