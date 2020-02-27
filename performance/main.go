package main

import (
	"crypto/sha256"
	mathrand "math/rand"
	"reflect"
	"testing"
	"time"
	"fmt"
	hardwaresim "github.com/ntu-brizo/hardwaresim"
)

func main(){
	HashStringPerformance()
	EncryptAndDecryptPerformance()
	SignAndVerifyPerformance()
}

func HashStringPerformance() {
	content := randContentHelper(1000000)
	t1 := time.Now() // get current time
    //logic handlers
    for i := 0; i < 1000; i++ {
        hashString, _ := HashString(content)

		_ = HashStringToHash(hashString)
    }

	elapsed := time.Since(t1)
    fmt.Println("Hash elapsed: ", elapsed)
}

func EncryptAndDecryptPerformance() {
	origContent := randContentHelper(1000000)
	
	t1 := time.Now() // get current time
    //logic handlers
    for i := 0; i < 1000; i++ {
        _, _, _ = Encrypt(origContent)
    }
	t1elapsed := time.Since(t1)
	fmt.Println("Encrypt elapsed: ", t1elapsed)

	crypted, key, _ := Encrypt(origContent)

	t2 := time.Now() // get current time
    //logic handlers
    for i := 0; i < 1000; i++ {
        decrypted, _ := Decrypt(key, crypted)
    }
	t2elapsed := time.Since(t2)
	fmt.Println("Decrypt elapsed: ", t2elapsed)
}

func SignAndVerifyPerformance() {
	p := randPUF()
	content := randContentHelper(1000000)
	hashString, _ := HashString(content)

	t1 := time.Now() // get current time
    //logic handlers
    for i := 0; i < 1000; i++ {
        _, _ = p.Sign(hashString)
    }

	elapsed := time.Since(t1)
    fmt.Println("Sign elapsed: ", elapsed)

	sig, _ := p.Sign(hashString)
	public, _ := p.CreatePublicKey()

	t2 := time.Now() // get current time
    //logic handlers
    for i := 0; i < 1000; i++ {
        _, _ = Verify(hashString, sig, public)
    }
	t2elapsed := time.Since(t2)
	fmt.Println("Verify elapsed: ", t2elapsed)
	
	
}


// randContentHelper func returns a random content for testing
// 1GB: 1073741824
func randContentHelper(size int) (randContent []byte) {
	mathrand.Seed(time.Now().UnixNano())
	// 1byte to {size} random content
	randContentSlice := make([]byte, mathrand.Intn(size)+1)
	mathrand.Read(randContentSlice)
	randContent = randContentSlice[:]
	return
}

// randPUF func generates a Puf for testing
func randPUF() Puf {
	mathrand.Seed(time.Now().UnixNano())
	// suppose puf id length is 256
	randContent := make([]byte, 256)
	mathrand.Read(randContent)
	return Puf{randContent[:]}
}
