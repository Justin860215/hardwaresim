package main

import (
    "bufio"
    "fmt"
    "log"
	"net"
	"strings"
	hardwaresim "github.com/ntu-brizo/hardwaresim"
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
    bufferString, err := bufio.NewReader(conn).ReadString('\n')

    if err != nil {
        log.Println("client left..")
        conn.Close()
        return
    }
	original := strings.Split(bufferString, " ")
	if original[0] == "E" {
		content := []byte(original[1])
		encryptContent, key, _ := hardwaresim.Encrypt(content)
		
		hash, _ := hardwaresim.HashString(encryptContent)
		
		ID := make([]byte, 256)
		puf := hardwaresim.Puf{ID[:]}
		sig, _ := puf.Sign(hash)

		clientAddr := conn.RemoteAddr().String()
		response := fmt.Sprintf(bufferString + " from " + clientAddr + "\n")

		log.Println(response)
		conn.Write(append([]byte(encryptContent),0x0a))
		conn.Write(append([]byte(key),0x0a))
		conn.Write(append([]byte(hash),0x0a))
		conn.Write(append([]byte(sig),0x0a))
	} else if original[0] == "D"{
		//content := []byte(original[1])
		// TODO
	} else {
		fmt.Println("Wrong cmd")
	}
		
    handleConnection(conn)
}