package main

import (
	"fmt"
	"net"
)

func connectServer(address string) (*net.UDPConn, error) {
	s, err := net.ResolveUDPAddr(NETWORK, address)
	if checkError(err) { return nil, err }
	return net.ListenUDP(NETWORK, s)
}

func runServerLoop(connection *net.UDPConn) {
	buffer := make([]byte, MAX_BUFFER_SIZE)
	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		printReceived(buffer, n)
		text := string(buffer[0:n])

		if isStopCommand(text) { return }

		// Reply by sending the same string back
		err = sendString(text, connection, addr)
		if checkError(err) { return	}		
	}
}

func main() {
	// Connect
	port := ":" + getFromArgument(1, DEFAULT_PORT)
	connection, err := connectServer(port)
	if checkError(err) { return	}

	// Other setup and print feedback
	defer connection.Close()
	fmt.Printf("\n\n\nServer running on port %s and listening for clients...\n", port)
	fmt.Println("Server will echo everything it receives.")
	printHeader()
	
	// Run the main loop
	runServerLoop(connection)
	fmt.Println("Stopping UDP server\n")
}
