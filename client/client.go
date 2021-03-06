package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func read(conn net.Conn) { //i removed the pointer thingy because i hate pointers
	//TODO In a continuous loop, read a message from the server and display it.
	reader := bufio.NewReader(conn) //does it need to be in the loop?
	for {
		msg, _ := reader.ReadString('\n')
		fmt.Println(msg)
	}
}

func write(conn net.Conn) { //did i mention that i hate pointers?
	//TODO Continually get input from the user and send messages to the server.
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter message: ")
		text, _ := stdin.ReadString('\n')
		fmt.Fprintf(conn, text) //flag returns pointers
	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	//TODO Try to connect to the server
	conn, _ := net.Dial("tcp", *addrPtr) //can't be helped here i suppose
	//TODO Start asynchronously reading and displaying messages
	go read(conn)
	//TODO Start getting and sending user messages.
	write(conn)
}
