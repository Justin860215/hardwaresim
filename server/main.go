package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
)

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatal("tcp server listener error:", err)
    }

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatal("tcp server accept error", err)
        }

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    bufferBytes, err := bufio.NewReader(conn).ReadString('\n')

    if err != nil {
        log.Println("client left..")
        conn.Close()
        return
    }

    message := string(bufferBytes)
    clientAddr := conn.RemoteAddr().String()
    response := fmt.Sprintf(message + " from " + clientAddr + "\n")

    log.Println(response)

    conn.Write([]byte("you sent: " + response))

    handleConnection(conn)
}