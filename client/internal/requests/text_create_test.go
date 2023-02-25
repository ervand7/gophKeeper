package requests

import (
	"encoding/hex"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"gophkeeper/pkg/algorithms"
	pb "gophkeeper/proto"
)

func TestCreateTextSuccess(t *testing.T) {
	setEnv(t)
	client := NewClient()
	_ = createAuthorizedUser(t, client)
	assert.NotEmpty(t, client.token)

	result := client.CreateText(
		&pb.CreateTextRequest{
			Text: &pb.Text{
				Title:   algorithms.RandString(10),
				Content: algorithms.RandString(10),
				Meta:    algorithms.RandString(10),
			},
		},
	)
	assert.Equal(t, "created successfully", result)

	client.DeleteUser()
}

func TestCreateTextFail_NotAuthorized(t *testing.T) {
	setEnv(t)
	client := NewClient()
	assert.Empty(t, client.token)

	result := client.CreateText(
		&pb.CreateTextRequest{
			Text: &pb.Text{
				Title:   algorithms.RandString(10),
				Content: algorithms.RandString(10),
				Meta:    algorithms.RandString(10),
			},
		},
	)
	assert.Equal(t, "not authorized", result)
}

func TestCreateTextFail_UserNotFound(t *testing.T) {
	setEnv(t)
	client := NewClient()
	fakeToken := hex.EncodeToString([]byte(uuid.New().String()))
	client.token = fakeToken

	result := client.CreateText(
		&pb.CreateTextRequest{
			Text: &pb.Text{
				Title:   algorithms.RandString(10),
				Content: algorithms.RandString(10),
				Meta:    algorithms.RandString(10),
			},
		},
	)
	assert.Equal(t, "user not found", result)
}

func TestCreateTextFail_AlreadyExists(t *testing.T) {
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

	result = client.CreateText(
		&pb.CreateTextRequest{
			Text: &pb.Text{
				Title:   title,
				Content: algorithms.RandString(10),
				Meta:    algorithms.RandString(10),
			},
		},
	)
	assert.Equal(t, "entry with this title already exists", result)

	client.DeleteUser()
}
