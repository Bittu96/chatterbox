package cryptos

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	utils "projects/chatterbox/server/pkgs/utilities"
)

var (
	AES_256_MASTER_KEY = utils.GetEnv("AES_256_MASTER_KEY").(string)
)

func AES() {
	secret := []byte(AES_256_MASTER_KEY)
	data := []byte("Test")
	encResult, err := Encrypt(data, secret)
	fmt.Println("encResult", encResult)
	fmt.Println("enc err", err)
	//  encrypted := "U2FsdGVkX1+LU7rE47VtIDwGIOsJa05BzOmAzQfdbVk="

	result, err := Decrypt(encResult, secret)
	fmt.Println("decrypted result", result)
	fmt.Println("decryption err", err)
}

/*CBC encryption Follow the example code of the golang standard library
But there is no padding inside, so make up
*/

// Use PKCS7 to fill, IOS is also 7
func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// aes encryption, filling the 16 bits of the key key, 24, 32 respectively corresponding to AES-128, AES-192, or AES-256.
func AesCBCEncrypt(rawData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	//fill the original
	blockSize := block.BlockSize()
	rawData = PKCS7Padding(rawData, blockSize)
	// Initial vector IV must be unique, but does not need to be kept secret
	cipherText := make([]byte, blockSize+len(rawData))
	//block size 16
	iv := cipherText[:blockSize]
	fmt.Println("iv", iv)

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	//block size and initial vector size must be the same
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[blockSize:], rawData)

	return cipherText, nil
}

func AesCBCDecrypt(encryptData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	blockSize := block.BlockSize()

	if len(encryptData) < blockSize {
		panic("cipherText too short")
	}
	iv := encryptData[:blockSize]
	encryptData = encryptData[blockSize:]

	// CBC mode always works in whole blocks.
	if len(encryptData)%blockSize != 0 {
		panic("cipherText is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(encryptData, encryptData)
	// Un fill
	encryptData = PKCS7UnPadding(encryptData)
	return encryptData, nil
}

func Encrypt(rawData, key []byte) (string, error) {
	data, err := AesCBCEncrypt(rawData, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func Decrypt(rawData string, key []byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		return "", err
	}
	dnData, err := AesCBCDecrypt(data, key)
	if err != nil {
		return "", err
	}
	return string(dnData), nil
}
