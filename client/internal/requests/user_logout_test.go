package requests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	pb "gophkeeper/proto"
)

func TestLogoutSuccess(t *testing.T) {
	setEnv(t)
	client := NewClient()
	user := createAuthorizedUser(t, client)
	assert.NotEmpty(t, client.token)

	result := client.Logout()
	assert.Equal(t, "success logout", result)
	assert.Empty(t, client.token)
	assert.Empty(t, client.userData)

	client.Auth(
		&pb.AuthRequest{Login: user.login, Password: user.password},
	)
	client.DeleteUser()
}

func TestLogoutFail_NotAuthorized(t *testing.T) {
	setEnv(t)
	client := NewClient()
	assert.Empty(t, client.token)

	result := client.Logout()
	assert.Equal(t, "not authorized", result)
	assert.Empty(t, client.token)
	assert.Empty(t, client.userData)
}
