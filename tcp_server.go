package main

import (
	"net"
	"fmt"
	"bufio"
	"strings"
)

// create an err check function with err variable
func check(err error, message string) {
	if err != nil {
		panic(err) // panic used if program has reach an unrecoverable state
	}
	fmt.Printf("%s\n", message)
}

func main() {

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, err := net.Listen("tcp", "10.112.5.13:8080")
	check(err, "Server is ready")

	// defer ln to close? >> defer ln.Close

	// run loop forever (or until ctrl-c)
	for {
		// accept multiple connections for each new client connectionon port
		conn, err := ln.Accept()
		check(err, "Accepted connection")

		// will listen for message to process ending in newline (\n)
		// each connection would have its own goroutine, but only one goroutine is actively running at a time
		// when goroutine blocked / waiting for client, Go's scheduler moves to a different routine that isn't blocked
		go func() {
			buf := bufio.NewReader(conn)

			// iterate communication
			for {
				message, err := buf.ReadString('\n')

				// if client disconnects, err != nil
				if err != nil {
					fmt.Printf("Client disconnected.\n")
					break
				}

				// output message received
				fmt.Print("Message Received:", string(message))

				// sample process for string received
				newmessage := strings.ToUpper(message)

				// send new string back to client
				conn.Write([]byte(newmessage + "\n"))
			}
		}() // "select" statement, similar to switch case, used more like a timeout
		// run the function that's supposed to be concurrent with the goroutine
		// client.Close()
	}
}
