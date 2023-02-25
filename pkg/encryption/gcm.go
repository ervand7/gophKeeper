// Package encryption for encrypt sensitive data.
package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"log"
	"os"
	"strconv"
	"strings"
)

var nonce = []byte{3, 108, 138, 34, 220, 31, 144, 154, 107, 125, 33, 4}

type Encrypter struct {
	gcm cipher.AEAD
}

func NewEncrypter() *Encrypter {
	keyRaw := os.Getenv("encryptionKey")
	if keyRaw == "" {
		log.Fatal("encryption key must be passed in env")
	}
	key := sha256.Sum256([]byte(keyRaw))

	block, err := aes.NewCipher(key[:])
	if err != nil {
		log.Fatal(err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &Encrypter{
		gcm: gcm,
	}
}

func (e *Encrypter) Encrypt(src string) string {
	encrypted := e.gcm.Seal(nil, nonce, []byte(src), nil)
	return e.toStr(encrypted)
}

func (e *Encrypter) Decrypt(data string) (string, error) {
	byteArray, err := e.toByteArray(data)
	if err != nil {
		return "", nil
	}
	decrypted, err := e.gcm.Open(nil, nonce, byteArray, nil)
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

func (e *Encrypter) toStr(byteArray []byte) string {
	builder := strings.Builder{}
	for _, val := range byteArray {
		elem := strconv.Itoa(int(val))
		builder.WriteString(elem + " ")
	}
	return builder.String()
}

func (e *Encrypter) toByteArray(src string) ([]byte, error) {
	split := strings.Split(strings.TrimSpace(src), " ")
	var byteArray []byte
	for _, val := range split {
		elem, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		byteArray = append(byteArray, byte(elem))
	}
	return byteArray, nil
}
