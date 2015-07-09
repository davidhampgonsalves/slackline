package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
  "golang.org/x/crypto/sha3"
	"errors"
	"io"
	"os"
	"runtime"
	"strings"
)

func Decrypt(str string) (string, error) {
	if len(str) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	bytes := []byte(str)
	iv := bytes[:aes.BlockSize]
	bytes = bytes[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(generateCipherBlock(), iv)
	cfb.XORKeyStream(bytes, bytes)

	return string(bytes), nil
}

func Encrypt(str string) string {
	cipherData := make([]byte, aes.BlockSize+len(str))
	iv := []byte(cipherData[:aes.BlockSize])
	io.ReadFull(rand.Reader, iv)

	cfb := cipher.NewCFBEncrypter(generateCipherBlock(), iv)

	cfb.XORKeyStream(cipherData[aes.BlockSize:], []byte(str))

	return string(cipherData)
}
func generateCipherBlock() cipher.Block {
	host, _ := os.Hostname()
	cpuCount := string(runtime.NumCPU())
	keyString := strings.Join([]string{cpuCount, runtime.GOARCH, runtime.GOOS, host}, "-")

	key := make([]byte, 32)
  sha3.ShakeSum256(key, []byte(keyString))

  block, _ := aes.NewCipher(key)

	return block
}
