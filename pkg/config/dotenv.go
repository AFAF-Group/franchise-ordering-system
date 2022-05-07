package config

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

const defaultConfigFile = ".env"

var (
	ErrDecryptDataToShort = errors.New("decrypt: plaint text to short")
)

func DecryptData(plainText, key string) string {
	if plainText == "" {
		return plainText
	}
	text, errDecodeString := base64.StdEncoding.DecodeString(plainText)
	if errDecodeString != nil {
		return errDecodeString.Error()
	}
	block, errNewCipher := aes.NewCipher([]byte(key))
	if errNewCipher != nil {
		return errNewCipher.Error()
	}
	if len(text) < aes.BlockSize {
		return ErrDecryptDataToShort.Error()
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)

	return string(text)
}

func Env(key, defaultValue string, decrypt bool) string {
	if strVal, ok := os.LookupEnv(key); ok {
		return strVal
	}
	return defaultValue
}

func EnvAsInt(key string, defaultVal int, decrypt bool) int {
	strVal := Env(key, "", decrypt)
	if val, err := strconv.Atoi(strVal); err == nil {
		return val
	}
	return defaultVal
}

func LoadEnv(file string) {
	once.Do(func() {
		if file == "" {
			file = defaultConfigFile
		}
		if err := godotenv.Load(file); err != nil {
			log.Fatalf("Error loading %s file", file)
		}
	})
}
