package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func connectClient(address string) (*net.UDPConn, error) {
	s, err := net.ResolveUDPAddr(NETWORK, address)
	if checkError(err) { return nil, err }
	return net.DialUDP(NETWORK, nil, s)
}

func receiveData(connection *net.UDPConn) ([]byte, int, error) {
	buffer := make([]byte, MAX_BUFFER_SIZE)
	n, _, err := connection.ReadFromUDP(buffer)
	return buffer, n, err
}

func runSendLoop(connection *net.UDPConn) error {
	text := SAMPLE_PAYLOAD
	expectedSize := len( []byte(SAMPLE_PAYLOAD) )
	packetsSent := 0
	packetsReceived := 0
	totalBytesSent := 0
	totalBytesReceived := 0
	var err error
	timeStart := time.Now()
	timeEnd := timeStart.Add(SAMPLE_SEND_DURATION)
	for t := timeStart; t.Before(timeEnd); t = time.Now() {
		err = sendString(text, connection, nil)
		packetsSent += 1
		totalBytesSent += expectedSize
		if err != nil { break }
		buffer, n, receiveErr := receiveData(connection)
		err = receiveErr
		if err != nil { break }
		packetsReceived += 1
		printReceived(buffer, n)
		totalBytesReceived += n
		time.Sleep(SAMPLE_SEND_PAUSE)
	}
	fmt.Printf("\nPackets Sent: %d, Packets Received: %d", packetsSent, packetsReceived)
	fmt.Printf("\nBytes Sent: %d, Bytes Received: %d\n", totalBytesSent, totalBytesReceived)
	return err
}

func readStringFromTerminal() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(INDENT + ">> ")
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(strings.TrimSuffix(text, "\n"))
}

func runClientLoop(connection *net.UDPConn) {
	for {
		text := readStringFromTerminal()
		if isSendCommand(text) {
			err := runSendLoop(connection)
			if checkError(err) { return }
			sendString(STOP_COMMAND, connection, nil)
			return
		}
		err := sendString(text, connection, nil)

		if checkError(err) { return }
		if isStopCommand(text) { return }

		buffer, n, err := receiveData(connection)
		if checkError(err) { return }
		printReceived(buffer, n)
	}
}

func main() {
	// Connect
	address := getFromArgument(1, DEFAULT_HOST + ":" + DEFAULT_PORT)
	connection, err := connectClient(address)
	if checkError(err) { return	}
	
	// Other setup and print feedback
	fmt.Printf("\n\n\nClient connected to UDP server at %s\n", connection.RemoteAddr().String())
	fmt.Printf(
		"Type text to echo, type %s to shut down client and server, type %s to send test data.\n",
		STOP_COMMAND,
		SEND_COMMAND,
	)
	printHeader()
	defer connection.Close()

	// Run the main loop
	runClientLoop(connection)
	fmt.Println("Stopping UDP client\n")
}
