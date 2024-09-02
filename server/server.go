package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

type Client struct {
	Conn     net.Conn
	Nickname string
	Address  string
}

var (
	clients   = make(map[net.Conn]*Client)
	broadcast = make(chan string)
	mutex     = sync.Mutex{}
)

// StartServer starts the TCP server and listens for incoming connections.
func StartServer(port string) {
	server, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("Error starting TCP server: %s\n", err)
		return
	}
	defer server.Close()

	fmt.Printf("Server started on port %s\n", port)
	go broadcastMessages()

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Add client to the map with default nickname as their address
	mutex.Lock()
	clients[conn] = &Client{Conn: conn, Nickname: conn.RemoteAddr().String(), Address: conn.RemoteAddr().String()}
	mutex.Unlock()

	// Notify others
	broadcast <- fmt.Sprintf("%s has joined the chat!", clients[conn].Nickname)

	// Handle messages from client
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		if strings.HasPrefix(message, "/nick ") {
			newNickname := strings.TrimSpace(strings.TrimPrefix(message, "/nick "))
			if newNickname == "" {
				fmt.Fprintln(conn, "Nickname cannot be empty.")
				continue
			}
			mutex.Lock()
			oldNickname := clients[conn].Nickname
			clients[conn].Nickname = newNickname
			mutex.Unlock()
			broadcast <- fmt.Sprintf("%s is now known as %s", oldNickname, newNickname)
		} else if strings.HasPrefix(message, "/msg ") {
			parts := strings.SplitN(message, " ", 3)
			if len(parts) < 3 {
				fmt.Fprintln(conn, "Usage: /msg <nickname> <message>")
				continue
			}
			recipientName := parts[1]
			privateMessage := parts[2]
			sendPrivateMessage(conn, recipientName, privateMessage)
		} else {
			broadcast <- fmt.Sprintf("%s: %s", clients[conn].Nickname, message)
		}
	}

	// Client disconnects
	mutex.Lock()
	name := clients[conn].Nickname
	delete(clients, conn)
	mutex.Unlock()
	broadcast <- fmt.Sprintf("%s has left the chat!", name)
}

func broadcastMessages() {
	for {
		msg := <-broadcast
		mutex.Lock()
		for _, client := range clients {
			fmt.Fprintln(client.Conn, msg)
		}
		mutex.Unlock()
	}
}

func sendPrivateMessage(sender net.Conn, recipientName, message string) {
	mutex.Lock()
	defer mutex.Unlock()

	var recipient *Client
	for _, client := range clients {
		if client.Nickname == recipientName {
			recipient = client
			break
		}
	}

	if recipient != nil {
		fmt.Fprintf(recipient.Conn, "[Private] %s: %s\n", clients[sender].Nickname, message)
		fmt.Fprintf(sender, "[Private] to %s: %s\n", recipient.Nickname, message)
	} else {
		fmt.Fprintln(sender, "User not found.")
	}
}
