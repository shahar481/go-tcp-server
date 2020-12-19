package main

import (
	"github.com/go-tcp-server"
	"net"
)

func handler(conn net.Conn) {
	conn.Write([]byte("Hey world!"))
	conn.Close()
}

func main() {
	t := tcp_server.NewTcpServer("0.0.0.0:12345", handler)
	err := t.StartListening()
	if err != nil {
		panic(err)
	}
	t.HandleConnections()
}
