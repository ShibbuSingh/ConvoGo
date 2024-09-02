# ConvoGo - TCP Chat Server

A Simple TCP Chat Server built in Go that allows multiple clients to connect, chat with each other, and send private messages. The server supports nickname setting and message broadcasting. This project demonstrates basic networking and concurrency concepts in Go.

## Key Features

- **Multi-Client Chat**: Allows multiple clients to connect and chat in real-time.
- **Nicknames**: Clients can set and update their custom nicknames using the `/nick` command.
- **Private Messaging**: Clients can send private messages to other users using the `/msg username message` command.
- **Message Broadcasting**: All messages are broadcasted to all connected clients, except private messages which are sent directly.

## Project Structure

- `server/`: Contains the server code responsible for managing client connections, handling messages, and broadcasting updates.
  - `server.go`: Main server implementation.
- `client/`: Contains the client code that connects to the server and sends/receives messages.
  - `client.go`: Main client implementation.
- `main.go`: Entry point for starting the server or client based on the specified mode.

## How to Run the Project

1. Open a terminal and navigate to the project directory.
2. Run the server with the following command:
   ```bash
   go run main.go -mode=server -port=8080
3. Run the client with the following command:
   ```bash
    go run main.go -mode=client -address=localhost:8080

