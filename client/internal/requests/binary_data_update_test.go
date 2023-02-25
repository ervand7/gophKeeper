package requests

import (
	"encoding/hex"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"gophkeeper/pkg/algorithms"
	pb "gophkeeper/proto"
)

func TestUpdateBinaryDataSuccess(t *testing.T) {
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

	result = client.UpdateBinaryData(
		&pb.UpdateBinaryDataRequest{
			BinaryData: &pb.BinaryData{
				Title:   title,
				Content: []byte(algorithms.RandString(10)),
				Meta:    algorithms.RandString(10),
			},
		},
	)
	assert.Equal(t, "updated successfully", result)

	client.DeleteUser()
}

func TestUpdateBinaryDataFail_NotAuthorized(t *testing.T) {
	setEnv(t)
	client := NewClient()
	assert.Empty(t, client.token)

	result := client.UpdateBinaryData(
		&pb.UpdateBinaryDataRequest{
			BinaryData: &pb.BinaryData{
				Title:   algorithms.RandString(10),
				Content: []byte(algorithms.RandString(10)),
				Meta:    algorithms.RandString(10),
			},
		},
	)
	assert.Equal(t, "not authorized", result)
}

func TestUpdateBinaryDataFail_UserNotFound(t *testing.T) {
	setEnv(t)
	client := NewClient()
	fakeToken := hex.EncodeToString([]byte(uuid.New().String()))
	client.token = fakeToken

	result := client.UpdateBinaryData(
		&pb.UpdateBinaryDataRequest{
			BinaryData: &pb.BinaryData{
				Title:   algorithms.RandString(10),
				Content: []byte(algorithms.RandString(10)),
				Meta:    algorithms.RandString(10),
			},
		},
	)
	assert.Equal(t, "user not found", result)
}

func TestUpdateBinaryDataFail_EntryNotExists(t *testing.T) {
	setEnv(t)
	client := NewClient()
	_ = createAuthorizedUser(t, client)
	assert.NotEmpty(t, client.token)

	result := client.UpdateBinaryData(
		&pb.UpdateBinaryDataRequest{
			BinaryData: &pb.BinaryData{
				Title:   algorithms.RandString(10),
				Content: []byte(algorithms.RandString(10)),
				Meta:    algorithms.RandString(10),
			},
		},
	)
	assert.Equal(t, "entry with this title does not exist", result)

	client.DeleteUser()
}
