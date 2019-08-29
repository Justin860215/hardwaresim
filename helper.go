package hardwarewallet

import "bytes"

// PKCS5Padding is a padding helper for aes encryption
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5UnPadding is an unpadding helper for aes dencryption
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// HashStringToHash convert hash from string type to []byte type
func HashStringToHash(hashString string) []byte {
	hashStringByte := []byte(hashString)
	var hash []byte
	var a1 byte
	i := 0
	for _, e := range hashStringByte {
		i++
		if e > 57 {
			e -= 87
		} else {
			e -= 48
		}
		if i == 1 {
			a1 = e
		} else if i == 2 {
			hash = append(hash, a1*16+e)
			i = 0
		}
	}
	return hash
}
