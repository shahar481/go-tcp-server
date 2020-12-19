# Go tcp server
Simple concurrent tcp server library written in golang. 
Working example file in the example directory.

## Usage example

```
func handler(conn net.Conn) {
    conn.Write([]byte("Hey world!"))
    conn.Close()
}

func main() {
    t := NewTcpServer("0.0.0.0:1234", handler)
    t.StartListening()
    t.HandleConnections() // BLOCKING
}
```