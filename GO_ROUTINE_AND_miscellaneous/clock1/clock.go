// TCP server that periodically writes time
package main

import (
  "net"
  "time"
  "log"
  "io"
  "os"
)

// must provide first arguement as address:port
// for eg localhost:8000 or 0.0.0.0:8000
func main() {
  listener, err := net.Listen("tcp", os.Args[1])
  if err != nil {
    log.Fatal(err)
  }
  for {
    conn, err  := listener.Accept()
    if err != nil {
      log.Print(err)                                // connection aborted
      continue
    }
    go handleConn(conn)                               // handle one connection at time
  }
}

func handleConn(c net.Conn) {
  defer c.Close()
  for {
    _, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
    if err != nil {
      return                                       // client disconnected
    }
    time.Sleep(1* time.Second)
  }
}
