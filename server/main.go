package main

import (
    "bufio"
    "fmt"
    "log"
	"net"
	"strings"
	"strconv"
	"encoding/gob"
	"bytes"
	hardwaresim "github.com/ntu-brizo/hardwaresim"
)

type Block1 struct {
    E   []byte      
	K []byte 
	H []byte
	P []byte
	S []byte 
}

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
	puf := hardwaresim.CreateTestPuf()

	if original[0] == "E" {
		content := []byte(original[1])
		encryptContent, key, _ := hardwaresim.Encrypt(content)
		
		hashString, _ := hardwaresim.HashString(encryptContent)
		hash := hardwaresim.HashStringToHash(hashString)
		
		
		sig, pubKey, _ := puf.Sign(hashString)

		clientAddr := conn.RemoteAddr().String()
		response := fmt.Sprintf(bufferString + " from " + clientAddr + "\n")

		log.Println(response)
		// todo: ENCODE 4 of them together
		res := &Block1{
			E: encryptContent,
			K: key,
			H: hash,
			P: pubKey,
			S: sig,
		}
		encode := res.Serialize()
		conn.Write(append(encode,0x0d))

	} else if original[0] == "V"{
		hashString := original[1]
		sig := hardwaresim.HashStringToHash(original[2])
		pubKey := hardwaresim.HashStringToHash(original[3])
		valid, _ := hardwaresim.Verify(hashString, sig, pubKey)

		clientAddr := conn.RemoteAddr().String()
		response := fmt.Sprintf(bufferString + " from " + clientAddr + "\n")
		log.Println(response)
		fmt.Println(strconv.FormatBool(valid))
		conn.Write(append([]byte(strconv.FormatBool(valid)),0x0d))

	} else if original[0] == "D"{
		encryptContent := hardwaresim.HashStringToHash(original[1])
		key := hardwaresim.HashStringToHash(original[2])
		content, _ := hardwaresim.Decrypt(key, encryptContent)

		clientAddr := conn.RemoteAddr().String()
		response := fmt.Sprintf(bufferString + " from " + clientAddr + "\n")
		log.Println(response)

		conn.Write(append(content,0x0d))

	} else {
		fmt.Println("Wrong cmd")
	}
		
    handleConnection(conn)
}

// Serialize serializes the block
func (b *Block1) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

