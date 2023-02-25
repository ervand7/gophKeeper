package requests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCredentialsFail_NotAuthorized(t *testing.T) {
	setEnv(t)
	client := NewClient()
	assert.Empty(t, client.token)

	_, err := client.GetCredentials()
	assert.Equal(t, "not authorized", err.Error())
}
