package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// StartClient connects to the server and handles user input and messages.
func StartClient(serverAddress string) {
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to the chat server.")
	fmt.Println("Commands:")
	fmt.Println("  /nick YourNickname  - Set your nickname.")
	fmt.Println("  /msg Nickname Message  - Send a private message to a specific user.")

	go readMessages(conn)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Fprintln(conn, scanner.Text())
	}
}

func readMessages(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
