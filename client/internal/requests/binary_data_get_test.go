package requests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBinaryDataFail_NotAuthorized(t *testing.T) {
	setEnv(t)
	client := NewClient()
	assert.Empty(t, client.token)

	_, err := client.GetBinaryData()
	assert.Equal(t, "not authorized", err.Error())
}
