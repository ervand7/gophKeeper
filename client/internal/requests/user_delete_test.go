package requests

import (
	"encoding/hex"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUserSuccess(t *testing.T) {
	setEnv(t)
	client := NewClient()
	_ = createAuthorizedUser(t, client)
	assert.NotEmpty(t, client.token)

	result := client.DeleteUser()
	assert.Equal(t, "user was deleted successfully", result)
	assert.Empty(t, client.token)
	assert.Empty(t, client.userData)

	client.DeleteUser()
}

func TestDeleteUserFail_NotAuthorized(t *testing.T) {
	setEnv(t)
	client := NewClient()
	assert.Empty(t, client.token)

	result := client.DeleteUser()
	assert.Equal(t, "not authorized", result)
	assert.Empty(t, client.token)
	assert.Empty(t, client.userData)

	client.DeleteUser()
}

func TestDeleteUserFail_UserNotFound(t *testing.T) {
	setEnv(t)
	client := NewClient()
	fakeToken := hex.EncodeToString([]byte(uuid.New().String()))
	client.token = fakeToken

	result := client.DeleteUser()
	assert.Equal(t, "user not found", result)
}
