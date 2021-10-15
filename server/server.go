package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strings"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	// TODO: all
	if err != nil {
		fmt.Println("honestly i don't know what went wrong")
	}
	// Deal with an error event.
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// TODO: all
	// Continuously accept a network connection from the Listener
	for {
		conn, err := ln.Accept()
		handleError(err)
		conns <- conn
	}
	// and add it to the channel for handling connections.
}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	// TODO: all
	// So long as this connection is alive:
	reader := bufio.NewReader(client) //you don't need that many readers
	for {
		text, err := reader.ReadString('\n')
		handleError(err)
		msg := Message{
			sender:  clientid,
			message: text}
		msgs <- msg
	}
	// Read in new messages as delimited by '\n's
	// Tidy up each message and add it to the messages channel,
	// recording which client it came from.
}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

	//TODO Create a Listener for TCP connections on the port given above.
	ln, err := net.Listen("tcp", *portPtr)
	handleError(err)

	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)

	names := make(map[int]string)

	//Start accepting connections
	go acceptConns(ln, conns)
	for {
		select {
		case conn := <-conns:
			//TODO Deal with a new connection
			// - assign a client ID
			clientid := len(clients)
			// - add the client to the clients channel
			clients[clientid] = conn
			//take a name
			fmt.Fprintln(conn, "Enter name: ")
			nameTaker := bufio.NewReader(conn)
			name, _ := nameTaker.ReadString('\n')
			name = strings.TrimSpace(name)
			names[clientid] = name
			// - start to asynchronously handle messages from this client
			go handleClient(conn, clientid, msgs)
		case msg := <-msgs:
			//TODO Deal with a new message
			// Send the message to all clients that aren't the sender
			for i, client := range clients {
				if msg.sender != i { //client is a connection, so we can't use it to compare
					fmt.Fprintf(client, "\n%s :: %s", names[msg.sender], msg.message) //does msg have a print method?
				}
			}
		}
	}
}
