package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }} // use default options

var (
	host string
	port int
	path string
	addr string
)

func init() {
	flag.StringVar(&host, "host", "0.0.0.0", "the address to listen for new connections")
	flag.IntVar(&port, "port", 9191, "the local listening port")
	flag.StringVar(&path, "path", "socket", "URL path to listen for websockets")
	flag.StringVar(&addr, "addr", "localhost:7777", "target address after socket tunnel")
}

func pipeConnections(target net.Conn, client net.Conn) {
	var wg sync.WaitGroup
	go func() {
		io.Copy(target, client)
		defer wg.Done()
	}()
	go func() {
		io.Copy(client, target)
		defer wg.Done()
	}()
	wg.Add(2)
	wg.Wait()
}

func socket(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	log.Print("Upgraded connection")
	defer c.Close()

	target, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal("could not connect to target", err)
	}
	defer target.Close()
	fmt.Printf("connection to server %v established!\n", target.RemoteAddr())

	client := c.UnderlyingConn()
	defer client.Close()

	pipeConnections(target, client)
}

func main() {
	flag.Parse()
	http.HandleFunc(fmt.Sprintf("/%s", path), socket)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil))
}
