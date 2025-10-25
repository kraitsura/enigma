package main

import (
	"fmt"
	"log"
	"net"
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

		fmt.Println("Connected Accepted from:", conn.RemoteAddr())

		conn.Close()

	}
}
