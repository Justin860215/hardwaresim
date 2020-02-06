package main

import (
  "log"
	"net"
  "fmt"
  "bufio"
  "os"
	"encoding/gob"
  "bytes"
  "strconv"
  "strings"
)

type Block1 struct {
  E   []byte      
  K []byte
  P []byte
  H []byte 
  S []byte 
}

// DeserializeBlock deserializes a block
func DeserializeBlock(d []byte) *Block1 {
	var block Block1  

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}

func convert( b []byte ) string {
  s := make([]string,len(b))
  for i := range b {
      s[i] = strconv.Itoa(int(b[i]))
  }
  return strings.Join(s,",")
}

func main() {

  conn, _ := net.Dial("tcp", "127.0.0.1:8080")
  for { 
    // read in input from stdin
    
    fmt.Print("Send(E)> ")
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    // send to socket
    fmt.Fprintf(conn, text + "\n")
    // listen for reply
    
    message1, _ := bufio.NewReader(conn).ReadBytes('\r')
    decodeMsg := DeserializeBlock(message1)
    fmt.Println("================================")
    fmt.Printf("Encrypt content: %x\n",decodeMsg.E)
    fmt.Printf("AES key: %x\n",decodeMsg.K)
    fmt.Printf("Pub key: %x\n",decodeMsg.P)
    fmt.Printf("Hash: %x\n",decodeMsg.H)
    fmt.Printf("Sig: %x\n",decodeMsg.S)
    fmt.Println("================================")

    fmt.Print("Send(V)> ")
    reader3 := bufio.NewReader(os.Stdin)
    text3, _ := reader3.ReadString('\n')
    // send to socket
    fmt.Fprintf(conn, text3 + "\n")
    // listen for reply
    message3, _ := bufio.NewReader(conn).ReadBytes('\r')
    fmt.Println("================================")
    fmt.Printf("Valid: %s", message3)
    fmt.Println("================================")

    fmt.Print("Send(D)> ")
    reader2 := bufio.NewReader(os.Stdin)
    text2, _ := reader2.ReadString('\n')
    // send to socket
    fmt.Fprintf(conn, text2 + "\n")
    // listen for reply
    message2, _ := bufio.NewReader(conn).ReadBytes('\r')
    fmt.Println("================================")
    fmt.Printf("Original content: %s", message2)
    fmt.Println("================================")
  }
}