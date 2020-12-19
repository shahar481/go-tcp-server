# Go tcp server
Simple concurrent tcp server library written in golang


## Usage example

```
func handler(conn net.Conn) {
    conn.Write("Hello world!")
    conn.Close()
}

func main() {
    t := NewTcpServer("0.0.0.0:1234", handler)
    t.StartListening()
    t.HandleConnections() // BLOCKING
}
```