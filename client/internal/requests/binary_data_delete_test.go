package requests

import (
	"encoding/hex"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"gophkeeper/pkg/algorithms"
	pb "gophkeeper/proto"
)

func TestDeleteBinaryDataSuccess(t *testing.T) {
	setEnv(t)
	client := NewClient()
	_ = createAuthorizedUser(t, client)
	assert.NotEmpty(t, client.token)

	title := algorithms.RandString(10)
	result := client.CreateBinaryData(
		&pb.CreateBinaryDataRequest{
			BinaryData: &pb.BinaryData{
				Title:   title,
				Content: []byte(algorithms.RandString(10)),
				Meta:    algorithms.RandString(10),
			},
		},
	)
	assert.Equal(t, "created successfully", result)

	result = client.DeleteBinaryData(
		&pb.DeleteBinaryDataRequest{Title: title},
	)
	assert.Equal(t, "deleted successfully", result)
	assert.Empty(t, client.userData)

	client.DeleteUser()
}

func TestDeleteBinaryDataFail_NotAuthorized(t *testing.T) {
	setEnv(t)
	client := NewClient()
	assert.Empty(t, client.token)

	result := client.DeleteBinaryData(
		&pb.DeleteBinaryDataRequest{Title: algorithms.RandString(10)},
	)
	assert.Equal(t, "not authorized", result)
}

func TestDeleteBinaryDataFail_UserNotFound(t *testing.T) {
	setEnv(t)
	client := NewClient()
	fakeToken := hex.EncodeToString([]byte(uuid.New().String()))
	client.token = fakeToken

	result := client.DeleteBinaryData(
		&pb.DeleteBinaryDataRequest{Title: algorithms.RandString(10)},
	)
	assert.Equal(t, "user not found", result)
}

func TestDeleteBinaryDataFail_EntryNotExists(t *testing.T) {
	setEnv(t)
	client := NewClient()
	_ = createAuthorizedUser(t, client)
	assert.NotEmpty(t, client.token)

	result := client.DeleteBinaryData(
		&pb.DeleteBinaryDataRequest{Title: algorithms.RandString(10)},
	)
	assert.Equal(t, "entry with this title does not exist", result)

	client.DeleteUser()
}
