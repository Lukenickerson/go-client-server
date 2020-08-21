package main

import (
	"fmt"
	"os"
	"strings"
	"net"
	"math"
	"time"
)

const DEFAULT_HOST = "127.0.0.1"
const DEFAULT_PORT = "40000"
const MAX_BUFFER_SIZE = 1024
const NETWORK = "udp4"
const STOP_COMMAND = "STOP"
const SEND_COMMAND = "SEND"
const INDENT = "                                        "
const SAMPLE_PAYLOAD = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris sit amet accumsan nibh. 123456789ABC"
const ONE_TENTH_SECOND time.Duration = 100000000
const ONE_MINUTE time.Duration = 60000000000 // 1s (1000000000 nan-seconds) * 60 = 1 minute
const TEN_MINUTES time.Duration = 600000000000 // 1s (1000000000 nan-seconds) * 600 = 10 minutes
const SAMPLE_SEND_DURATION time.Duration = TEN_MINUTES // try ONE_MINUTE for testing
const SAMPLE_SEND_PAUSE = ONE_TENTH_SECOND
const PRINT_STRING_CUTOFF = 20

var sendIndex int = 0

func getFromArgument(index int, defaultValue string) string {
	arguments := os.Args
	if len(arguments) <= index {
		return defaultValue
	}
	return arguments[index]
}

func isStopCommand(text string) bool {
	return strings.TrimSpace(text) == STOP_COMMAND
}

func isSendCommand(text string) bool {
	return strings.TrimSpace(text) == SEND_COMMAND
}

func checkError(err error) bool {
	isError := (err != nil)
	if isError {
		fmt.Println(err)
	}
	return isError
}

func truncate(text string, cutoff int) string {
	length := len(text)
	end := int(math.Min(float64(cutoff), float64(length)))
	ellipsis := ""
	if (length > cutoff) {
		ellipsis = "..."
	}
	return text[0:end] + ellipsis
}

func sendString(text string, connection *net.UDPConn, address *net.UDPAddr) error {
	data := []byte(text)
	sendIndex += 1
	fmt.Printf(INDENT + ">> #%d (%d) %s\n", sendIndex, len(data), truncate(string(data), PRINT_STRING_CUTOFF))
	var err error
	if address == nil {
		_, err = connection.Write(data)
	} else {
		_, err = connection.WriteToUDP(data, address)
	}
	return err
}

func printReceived(buffer []byte, n int) {
	fmt.Printf("-> (%d) %s\n", n, truncate(string(buffer[0:n]), PRINT_STRING_CUTOFF))
}

func printHeader() {
	const header1 = "|-----------RECEIVED-----------|"
	const header2 = "|-----------SENDING------------|"
	spacingLength := len(INDENT) - len(header1)
	spacing := INDENT[0:spacingLength]
	fmt.Printf("%s%s%s\n", header1, spacing, header2)
}
