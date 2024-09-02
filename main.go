package main

import (
	"ConvoGo/client"
	"ConvoGo/server"
	"flag"
	"fmt"
	"os"
)

func main() {
	mode := flag.String("mode", "", "Start in 'server' or 'client' mode")
	port := flag.String("port", "8080", "Port to connect to or listen on")
	serverAddress := flag.String("address", "localhost:8080", "Address of the chat server")
	flag.Parse()

	if *mode == "server" {
		server.StartServer(":" + *port)
	} else if *mode == "client" {
		client.StartClient(*serverAddress)
	} else {
		fmt.Println("Please specify a mode: -mode=server or -mode=client")
		os.Exit(1)
	}
}
