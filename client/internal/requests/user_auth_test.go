package requests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"gophkeeper/pkg/algorithms"
	pb "gophkeeper/proto"
)

func TestAuthSuccess(t *testing.T) {
	setEnv(t)
	client := NewClient()
	user := createNotAuthorizedUser(t, client)
	assert.Empty(t, client.token)

	result := client.Auth(
		&pb.AuthRequest{Login: user.login, Password: user.password},
	)
	assert.Equal(t, fmt.Sprintf("success auth for %s", user.login), result)
	assert.NotEmpty(t, client.token)

	client.DeleteUser()
}

func TestAuthFail_AlreadyAuthorized(t *testing.T) {
	setEnv(t)
	client := NewClient()
	user := createAuthorizedUser(t, client)
	assert.NotEmpty(t, client.token)

	result := client.Auth(
		&pb.AuthRequest{Login: user.login, Password: user.password},
	)
	assert.Equal(t, "already authorized", result)

	client.DeleteUser()
}

func TestAuthFail_UserNotFound(t *testing.T) {
	setEnv(t)
	client := NewClient()
	assert.Empty(t, client.token)

	result := client.Auth(
		&pb.AuthRequest{
			Login:    algorithms.RandString(10),
			Password: algorithms.RandString(10),
		},
	)
	assert.Equal(t, "user not found", result)
	assert.Empty(t, client.token)

	client.DeleteUser()
}
