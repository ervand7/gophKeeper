package requests

import (
	"encoding/hex"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"gophkeeper/pkg/algorithms"
	pb "gophkeeper/proto"
)

func TestUpdateTextSuccess(t *testing.T) {
	setEnv(t)
	client := NewClient()
	_ = createAuthorizedUser(t, client)
	assert.NotEmpty(t, client.token)

	title := algorithms.RandString(10)
	result := client.CreateText(
		&pb.CreateTextRequest{
			Text: &pb.Text{
				Title:   title,
				Content: algorithms.RandString(10),
				Meta:    algorithms.RandString(10),
			},
		},
	)
	assert.Equal(t, "created successfully", result)

	result = client.UpdateText(
		&pb.UpdateTextRequest{
			Text: &pb.Text{
				Title:   title,
				Content: algorithms.RandString(10),
				Meta:    algorithms.RandString(10),
			},
		},
	)
	assert.Equal(t, "updated successfully", result)

	client.DeleteUser()
}

func TestUpdateTextFail_NotAuthorized(t *testing.T) {
	setEnv(t)
	client := NewClient()
	assert.Empty(t, client.token)

	result := client.UpdateText(
		&pb.UpdateTextRequest{
			Text: &pb.Text{
				Title:   algorithms.RandString(10),
				Content: algorithms.RandString(10),
				Meta:    algorithms.RandString(10),
			},
		},
	)
	assert.Equal(t, "not authorized", result)
}

func TestUpdateTextFail_UserNotFound(t *testing.T) {
	setEnv(t)
	client := NewClient()
	fakeToken := hex.EncodeToString([]byte(uuid.New().String()))
	client.token = fakeToken

	result := client.UpdateText(
		&pb.UpdateTextRequest{
			Text: &pb.Text{
				Title:   algorithms.RandString(10),
				Content: algorithms.RandString(10),
				Meta:    algorithms.RandString(10),
			},
		},
	)
	assert.Equal(t, "user not found", result)
}

func TestUpdateTextFail_EntryNotExists(t *testing.T) {
	setEnv(t)
	client := NewClient()
	_ = createAuthorizedUser(t, client)
	assert.NotEmpty(t, client.token)

	result := client.UpdateText(
		&pb.UpdateTextRequest{
			Text: &pb.Text{
				Title:   algorithms.RandString(10),
				Content: algorithms.RandString(10),
				Meta:    algorithms.RandString(10),
			},
		},
	)
	assert.Equal(t, "entry with this title does not exist", result)

	client.DeleteUser()
}
