package tcp_server

import(
	"net"
)


type TcpServer struct {
	listenAddr string
	listener net.Listener
	handler func(conn net.Conn)
}

// Returns a TcpServer struct with the given parameters
func NewTcpServer(listenAddr string, handler func(conn net.Conn)) *TcpServer {
	t := new(TcpServer)
	t.listenAddr = listenAddr
	t.handler = handler
	return t
}

// Starts listening and initializes a socket in the TcpServer struct
func (t *TcpServer) StartListening() error {
	var err error
	t.listener, err = net.Listen("tcp4", t.listenAddr)
	if err != nil {
		return err
	}
	return nil
}

// Blocking function, accepts connections and calls the handler
// function with a goroutine
func (t TcpServer) HandleConnections() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			return
		}
		go t.handler(conn)
	}
}