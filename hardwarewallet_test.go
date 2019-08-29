package hardwarewallet

import (
	"crypto/sha256"
	mathrand "math/rand"
	"reflect"
	"testing"
	"time"
)

func TestHashStringAndReverse(t *testing.T) {
	content := randContentHelper(1000000)

	hashString, err := HashString(content)
	// fmt.Printf("hashString: %s\n", hashString)
	if err != nil {
		t.Fatal("got errors:", err)
	}

	hash := HashStringToHash(hashString)
	// fmt.Printf("hash: %x\n", hash)

	sha := sha256.Sum256(content)
	// fmt.Printf("Sha256: %x\n", sha[:])

	success := (reflect.DeepEqual(hash, sha[:]))
	testLogHelper(t, success, "hash result is not correct", "HashString OK")
}

func TestEncryptAndDecrypt(t *testing.T) {
	origContent := randContentHelper(1000000)
	crypted, key, err := Encrypt(origContent)
	if err != nil {
		t.Fatal("got errors: ", err)
	}
	decrypted, err := Decrypt(key, crypted)
	if err != nil {
		t.Fatal("got errors: ", err)
	}

	success := (reflect.DeepEqual(origContent, decrypted))
	testLogHelper(t, success, "decrypted data is not equal to original data", "Encrypt, Decrypt OK")
}

func TestSignAndVerify(t *testing.T) {
	p := randPUF()
	content := randContentHelper(1000000)
	hashString, _ := HashString(content)
	sig, err := p.Sign(hashString)
	if err != nil {
		t.Fatal("got errors: ", err)
	}
	public, err := p.CreatePublicKey()
	if err != nil {
		t.Fatal("got errors: ", err)
	}
	success, err := Verify(hashString, sig, public)
	if err != nil {
		t.Fatal("got errors: ", err)
	}
	testLogHelper(t, success, "Not verified", "Sign, Verify OK")
}

// testLogHelper func usage: testHelper(t, true, "err", "good")
func testLogHelper(t *testing.T, success bool, errMsg, successMsg string) {
	if success {
		t.Log(successMsg)
	} else {
		t.Error(errMsg)
	}
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
