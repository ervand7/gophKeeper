package requests

import (
	"encoding/hex"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"gophkeeper/pkg/algorithms"
	pb "gophkeeper/proto"
)

func TestUpdateCredentialsSuccess(t *testing.T) {
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

	result = client.UpdateCredentials(
		&pb.UpdateCredentialsRequest{
			Credentials: &pb.Credentials{
				Title:    title,
				Login:    algorithms.RandString(10),
				Password: algorithms.RandString(10),
				Meta:     algorithms.RandString(10),
			},
		},
	)
	assert.Equal(t, "updated successfully", result)

	client.DeleteUser()
}

func TestUpdateCredentialsFail_NotAuthorized(t *testing.T) {
	setEnv(t)
	client := NewClient()
	assert.Empty(t, client.token)

	result := client.UpdateCredentials(
		&pb.UpdateCredentialsRequest{
			Credentials: &pb.Credentials{
				Title:    algorithms.RandString(10),
				Login:    algorithms.RandString(10),
				Password: algorithms.RandString(10),
				Meta:     algorithms.RandString(10),
			},
		},
	)
	assert.Equal(t, "not authorized", result)
}

func TestUpdateCredentialsFail_UserNotFound(t *testing.T) {
	setEnv(t)
	client := NewClient()
	fakeToken := hex.EncodeToString([]byte(uuid.New().String()))
	client.token = fakeToken

	result := client.UpdateCredentials(
		&pb.UpdateCredentialsRequest{
			Credentials: &pb.Credentials{
				Title:    algorithms.RandString(10),
				Login:    algorithms.RandString(10),
				Password: algorithms.RandString(10),
				Meta:     algorithms.RandString(10),
			},
		},
	)
	assert.Equal(t, "user not found", result)
}

func TestUpdateCredentialsFail_EntryNotExists(t *testing.T) {
	setEnv(t)
	client := NewClient()
	_ = createAuthorizedUser(t, client)
	assert.NotEmpty(t, client.token)

	result := client.UpdateCredentials(
		&pb.UpdateCredentialsRequest{
			Credentials: &pb.Credentials{
				Title:    algorithms.RandString(10),
				Login:    algorithms.RandString(10),
				Password: algorithms.RandString(10),
				Meta:     algorithms.RandString(10),
			},
		},
	)
	assert.Equal(t, "entry with this title does not exist", result)

	client.DeleteUser()
}
