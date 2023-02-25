package requests

import (
	"encoding/hex"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"gophkeeper/pkg/algorithms"
	pb "gophkeeper/proto"
)

func TestDeleteCredentialsSuccess(t *testing.T) {
	setEnv(t)
	client := NewClient()
	_ = createAuthorizedUser(t, client)
	assert.NotEmpty(t, client.token)

	title := algorithms.RandString(10)
	result := client.CreateCredentials(
		&pb.CreateCredentialsRequest{
			Credentials: &pb.Credentials{
				Title:    title,
				Login:    algorithms.RandString(10),
				Password: algorithms.RandString(10),
				Meta:     algorithms.RandString(10),
			},
		},
	)
	assert.Equal(t, "created successfully", result)

	result = client.DeleteCredentials(
		&pb.DeleteCredentialsRequest{Title: title},
	)
	assert.Equal(t, "deleted successfully", result)
	assert.Empty(t, client.userData)

	client.DeleteUser()
}

func TestDeleteCredentialsFail_NotAuthorized(t *testing.T) {
	setEnv(t)
	client := NewClient()
	assert.Empty(t, client.token)

	result := client.DeleteCredentials(
		&pb.DeleteCredentialsRequest{Title: algorithms.RandString(10)},
	)
	assert.Equal(t, "not authorized", result)
}

func TestDeleteCredentialsFail_UserNotFound(t *testing.T) {
	setEnv(t)
	client := NewClient()
	fakeToken := hex.EncodeToString([]byte(uuid.New().String()))
	client.token = fakeToken

	result := client.DeleteCredentials(
		&pb.DeleteCredentialsRequest{Title: algorithms.RandString(10)},
	)
	assert.Equal(t, "user not found", result)
}

func TestDeleteCredentialsFail_EntryNotExists(t *testing.T) {
	setEnv(t)
	client := NewClient()
	_ = createAuthorizedUser(t, client)
	assert.NotEmpty(t, client.token)

	result := client.DeleteCredentials(
		&pb.DeleteCredentialsRequest{Title: algorithms.RandString(10)},
	)
	assert.Equal(t, "entry with this title does not exist", result)

	client.DeleteUser()
}
