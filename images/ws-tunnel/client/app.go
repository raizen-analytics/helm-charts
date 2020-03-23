package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/gorilla/websocket"
)

var (
	host string
	port int
)

func init() {
	flag.StringVar(&host, "host", "0.0.0.0", "the local address to listen for connections")
	flag.IntVar(&port, "port", 9192, "the local listening port")
}

func main() {
	flag.Parse()
	wsAddr := flag.Arg(0)
	fmt.Println(wsAddr)
	if wsAddr == "" {
		fmt.Println("Websocket address is required. Ex: ws://address, wss://address")
		return
	}
	incomingConn, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("could not start server on %d: %v", fmt.Sprintf("%s:%p", host, port), err)
	}
	fmt.Printf("server running on %d\n", port)
	defer incomingConn.Close()

	for {
		client, err := incomingConn.Accept()
		if err != nil {
			log.Fatal("could not accept client connection", err)
		}
		defer client.Close()
		fmt.Printf("client '%v' connected!\n", client.RemoteAddr())

		wsConn, _, err := websocket.DefaultDialer.Dial(wsAddr, nil)
		if err != nil {
			log.Fatal("dial:", err)
		}
		log.Print("Connected")
		defer wsConn.Close()
		target := wsConn.UnderlyingConn()

		go func() { io.Copy(target, client) }()
		go func() { io.Copy(client, target) }()
	}
}
