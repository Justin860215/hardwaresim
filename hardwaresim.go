package hardwaresim

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
	mathrand "math/rand"
	"time"
)

// Puf is the physically unclonable function provided by FPGA
type Puf struct {
	id []byte
}

// wallet stores private and public keys
type wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

// Encrypt func encrypts a slice of bytes and returns the encrypted data.
func Encrypt(content []byte) (encryptedContent []byte, AESKey []byte, err error) {
	mathrand.Seed(time.Now().UnixNano())
	AESKey = make([]byte, 32)
	mathrand.Read(AESKey)

	block, err := aes.NewCipher(AESKey)
	if err != nil {
		return nil, nil, err
	}

	blockSize := block.BlockSize()

	origData := PKCS5Padding(content, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, AESKey[:blockSize])
	encryptedContent = make([]byte, len(origData))
	// crypted := origData
	blockMode.CryptBlocks(encryptedContent, origData)

	return
}

// Decrypt func decrypts an encrypted slice of bytes and returns the decrypted data.
func Decrypt(key, encryptedContent []byte) (content []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	content = make([]byte, len(encryptedContent))
	blockMode.CryptBlocks(content, encryptedContent)
	content = PKCS5UnPadding(content)
	return
}

// HashString func computes hash value of a slice of bytes.
func HashString(content []byte) (hashstring string, err error) {
	hash := sha256.Sum256(content)
	hashstring = fmt.Sprintf("%x", hash[:])
	return
}

// Sign func generates a signature on a string.
func (p Puf) Sign(hashstring string) (signature []byte, err error) {
	w := createWallet(p)
	hash := HashStringToHash(hashstring)
	r, s, err := ecdsa.Sign(rand.Reader, &w.PrivateKey, hash)
	if err != nil {
		return nil, err
	}
	signature = append(r.Bytes(), s.Bytes()...)
	return
}

// Verify func verifies the signature and returns a boolean.
func Verify(hashstring string, signature, publicKey []byte) (bool, error) {
	hash := HashStringToHash(hashstring)
	curve := elliptic.P256()

	r := big.Int{}
	s := big.Int{}
	sigLen := len(signature)
	r.SetBytes(signature[:(sigLen / 2)])
	s.SetBytes(signature[(sigLen / 2):])

	publicX := big.Int{}
	publicY := big.Int{}
	keyLen := len(publicKey)
	publicX.SetBytes(publicKey[:(keyLen / 2)])
	publicY.SetBytes(publicKey[(keyLen / 2):])

	rawPubKey := ecdsa.PublicKey{Curve: curve, X: &publicX, Y: &publicY}
	if ecdsa.Verify(&rawPubKey, hash, &r, &s) == false {
		return false, nil
	}
	return true, nil
}

// CreatePublicKey func generates a public key .
func (p Puf) CreatePublicKey() (publicKey []byte, err error) {
	w := createWallet(p)
	publicKey = w.PublicKey
	return
}

func createWallet(p Puf) *wallet {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, bytes.NewReader(p.id))
	if err != nil {
		log.Panic(err)
	}
	public := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	w := wallet{*private, public}

	return &w
}

//  CreateTestPuf func generates a Puf for testing
func CreateTestPuf() *Puf {
	// suppose puf id length is 256
	randContent := make([]byte, 256)
	return Puf{randContent[:]}
}