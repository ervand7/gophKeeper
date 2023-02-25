package encryption

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"gophkeeper/pkg/algorithms"
)

func TestEncryptAndDecryptSuccess(t *testing.T) {
	err := os.Setenv("encryptionKey", "qwerty")
	assert.NoError(t, err)

	encrypter := NewEncrypter()
	src := algorithms.RandString(10)

	encrypted := encrypter.Encrypt(src)
	assert.NotEqual(t, src, encrypted)

	decrypted, err := encrypter.Decrypt(encrypted)
	assert.NoError(t, err)

	assert.Equal(t, decrypted, src)
}
