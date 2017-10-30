package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func main() {

	// prompt user for sign in
	// send data to the server (save it as csvfile or something? hashtable?)

	// connect to this socket
	conn, _ := net.Dial("tcp", "10.112.5.13:8080")

	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')

		// send to socket
		fmt.Fprintf(conn, text + "\n")
		//listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
