# Go Client-Server Test App

## What is This?

This application is a learning exercise for making a very simple UDP client and server in GoLang.

## How to Use

1. Install Go
1. Run the **server** with `./server.sh` (if you have bash) or `go run core.go udp-server.go`
   * Optionally add a param for port, e.g. `./server.sh 40000`
   * The server will start up and wait for data from the client(s).
   * Server is very dumb, and just echoes data that it receives.
1. Run the **client** with `./client.sh` or `go run core.go udp-client.go` on a separate machine or separate terminal.
   * Optionally add a param for host and port like `./client.sh example.com:40000` (default is localhost `127.0.0.1:40000`)
   * The client will start up and await for your commands.
1. Enter any text, e.g., `Hello world`, on the client and hit enter. You should see the text sent from client to server, then echoed back from server to client.
1. Enter `STOP` to stop both the client and the server.
1. Enter `SEND` to begin continually sending data from the client to the server for 10 minutes.
    * Purpose of this is to test the packet drop-rate.
	* After the cycle stops the number of expected and received packets and bytes should be shown, and a `STOP` command will be issued.
	* (Note: This is untested in a real environment.)

## Good Resources
* https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/
* https://ops.tips/blog/udp-client-and-server-in-go/
* https://golang.org/pkg/net/