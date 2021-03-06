package main

import (
	mathrand "math/rand"
	"time"
	"fmt"
	hardwaresim "github.com/ntu-brizo/hardwaresim"
)

func main(){
	content := randContentHelper(1000000)
	HashStringPerformance(content)
	EncryptAndDecryptPerformance(content)
	SignAndVerifyPerformance(content)
}

func HashStringPerformance(content []byte) {
	t1 := time.Now() // get current time
    //logic handlers
    for i := 0; i < 1000; i++ {
        hashString, _ := hardwaresim.HashString(content)

		_ = hardwaresim.HashStringToHash(hashString)
    }

	elapsed := time.Since(t1)
    fmt.Println("Hash elapsed: ", elapsed)
}

func EncryptAndDecryptPerformance(origContent []byte) {
	
	t1 := time.Now() // get current time
    //logic handlers
    for i := 0; i < 1000; i++ {
        _, _, _ = hardwaresim.Encrypt(origContent)
    }
	t1elapsed := time.Since(t1)
	fmt.Println("Encrypt elapsed: ", t1elapsed)

	crypted, key, _ := hardwaresim.Encrypt(origContent)

	t2 := time.Now() // get current time
    //logic handlers
    for i := 0; i < 1000; i++ {
        _, _ = hardwaresim.Decrypt(key, crypted)
    }
	t2elapsed := time.Since(t2)
	fmt.Println("Decrypt elapsed: ", t2elapsed)
}

func SignAndVerifyPerformance(content []byte) {
	p := hardwaresim.CreateTestPuf()
	hashString, _ := hardwaresim.HashString(content)

	t1 := time.Now() // get current time
    //logic handlers
    for i := 0; i < 1000; i++ {
        _, _, _ = p.Sign(hashString)
    }

	elapsed := time.Since(t1)
    fmt.Println("Sign elapsed: ", elapsed)

	sig, _, _ := p.Sign(hashString)
	public, _ := p.CreatePublicKey()

	t2 := time.Now() // get current time
    //logic handlers
    for i := 0; i < 1000; i++ {
        _, _ = hardwaresim.Verify(hashString, sig, public)
    }
	t2elapsed := time.Since(t2)
	fmt.Println("Verify elapsed: ", t2elapsed)
	
	
}


// randContentHelper func returns a random content for testing
// 1GB: 1073741824
func randContentHelper(size int) (randContent []byte) {
	mathrand.Seed(time.Now().UnixNano())
	// 1byte to {size} random content
	randContentSlice := make([]byte, size)
	mathrand.Read(randContentSlice)
	randContent = randContentSlice[:]
	return
}
