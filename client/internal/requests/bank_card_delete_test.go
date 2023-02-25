package requests

import (
	"encoding/hex"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"gophkeeper/pkg/algorithms"
	pb "gophkeeper/proto"
)

func TestDeleteBankCardSuccess(t *testing.T) {
	setEnv(t)
	client := NewClient()
	_ = createAuthorizedUser(t, client)
	assert.NotEmpty(t, client.token)

	title := algorithms.RandString(10)
	result := client.CreateBankCard(
		&pb.CreateBankCardRequest{
			BankCard: &pb.BankCard{
				Title:      title,
				CardHolder: algorithms.RandString(10),
				CardNumber: algorithms.RandString(10),
				CardExpire: algorithms.RandString(10),
				CardCvv:    algorithms.RandString(10),
				Meta:       algorithms.RandString(10),
			},
		},
	)
	assert.Equal(t, "created successfully", result)

	result = client.DeleteBankCard(
		&pb.DeleteBankCardRequest{Title: title},
	)
	assert.Equal(t, "deleted successfully", result)
	assert.Empty(t, client.userData)

	client.DeleteUser()
}

func TestDeleteBankCardFail_NotAuthorized(t *testing.T) {
	setEnv(t)
	client := NewClient()
	assert.Empty(t, client.token)

	result := client.DeleteBankCard(
		&pb.DeleteBankCardRequest{Title: algorithms.RandString(10)},
	)
	assert.Equal(t, "not authorized", result)
}

func TestDeleteBankCardFail_UserNotFound(t *testing.T) {
	setEnv(t)
	client := NewClient()
	fakeToken := hex.EncodeToString([]byte(uuid.New().String()))
	client.token = fakeToken

	result := client.DeleteBankCard(
		&pb.DeleteBankCardRequest{Title: algorithms.RandString(10)},
	)
	assert.Equal(t, "user not found", result)
}

func TestDeleteBankCardFail_EntryNotExists(t *testing.T) {
	setEnv(t)
	client := NewClient()
	_ = createAuthorizedUser(t, client)
	assert.NotEmpty(t, client.token)

	result := client.DeleteBankCard(
		&pb.DeleteBankCardRequest{Title: algorithms.RandString(10)},
	)
	assert.Equal(t, "entry with this title does not exist", result)

	client.DeleteUser()
}
