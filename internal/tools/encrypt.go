package tools

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

var key []byte

func init() {
	if err := godotenv.Load(); err != nil {
		log.Error("No .env file found")
	}

	key = []byte(os.Getenv("AES_KEY"))
}

func EncryptAESG(val string) (encVal string, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Error("Failed to create cipher block: ", err)
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(val))
	iv := cipherText[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(val))

	return hex.EncodeToString(cipherText), nil
}
