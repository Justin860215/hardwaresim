package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

  conn, _ := net.Dial("tcp", "127.0.0.1:8080")
  for { 
    // read in input from stdin
    
	fmt.Print("CMD to send: ")
	reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    // send to socket
    fmt.Fprintf(conn, text + "\n")
    // listen for reply
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message from server: "+message)
  }
}