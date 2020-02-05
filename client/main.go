package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

  conn, _ := net.Dial("tcp", "127.0.0.1:8080")
  for { 
    // read in input from stdin
    
	fmt.Print("Send> ")
	reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    // send to socket
    fmt.Fprintf(conn, text + "\n")
    // listen for reply
    message1, _ := bufio.NewReader(conn).ReadBytes('\n')
	fmt.Println(message1)
	message2, _ := bufio.NewReader(conn).ReadBytes('\n')
	fmt.Println(message2)
	message3, _ := bufio.NewReader(conn).ReadBytes('\n')
	fmt.Println(message3)
	message4, _ := bufio.NewReader(conn).ReadBytes('\n')
    fmt.Println(message4)
  }
}