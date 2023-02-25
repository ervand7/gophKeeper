package requests

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"gophkeeper/pkg/algorithms"
	pb "gophkeeper/proto"
)

type userCredentials struct {
	login    string
	password string
}

func createAuthorizedUser(t *testing.T, client *Client) userCredentials {
	login := algorithms.RandString(10)
	password := algorithms.RandString(10)
	assert.Empty(t, client.token)

	result := client.Register(&pb.RegisterRequest{Login: login, Password: password})
	assert.NotEmpty(t, client.token)
	assert.Equal(t, fmt.Sprintf("success register for %s", login), result)

	return userCredentials{login: login, password: password}
}

func createNotAuthorizedUser(t *testing.T, client *Client) userCredentials {
	login := "hello"
	password := "world"
	assert.Empty(t, client.token)

	result := client.Register(&pb.RegisterRequest{Login: login, Password: password})
	assert.NotEmpty(t, client.token)
	assert.Equal(t, fmt.Sprintf("success register for %s", login), result)

	client.Logout()

	return userCredentials{login: login, password: password}
}

func setEnv(t *testing.T) {
	err := os.Setenv("encryptionKey", "qwerty")
	assert.NoError(t, err)
}
