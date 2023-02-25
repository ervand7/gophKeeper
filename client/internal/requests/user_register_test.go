package requests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"gophkeeper/pkg/algorithms"
	pb "gophkeeper/proto"
)

func TestRegisterSuccess(t *testing.T) {
	setEnv(t)
	client := NewClient()
	assert.Empty(t, client.token)

	login := "hello"
	result := client.Register(
		&pb.RegisterRequest{
			Login:    login,
			Password: algorithms.RandString(10),
		},
	)
	assert.NotEmpty(t, client.token)
	assert.Equal(t, fmt.Sprintf("success register for %s", login), result)

	client.DeleteUser()
}

func TestRegisterFail_AlreadyExists(t *testing.T) {
	setEnv(t)
	client := NewClient()
	assert.Empty(t, client.token)

	login := "hello"
	result := client.Register(
		&pb.RegisterRequest{
			Login:    login,
			Password: algorithms.RandString(10),
		},
	)
	assert.NotEmpty(t, client.token)
	assert.Equal(t, fmt.Sprintf("success register for %s", login), result)

	result = client.Register(
		&pb.RegisterRequest{
			Login:    login,
			Password: algorithms.RandString(10),
		},
	)
	assert.Equal(t, fmt.Sprintf("%s already exists", login), result)

	client.DeleteUser()
}
