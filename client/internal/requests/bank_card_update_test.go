package requests

import (
	"encoding/hex"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"gophkeeper/pkg/algorithms"
	pb "gophkeeper/proto"
)

func TestUpdateBankCardSuccess(t *testing.T) {
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

	result = client.UpdateBankCard(
		&pb.UpdateBankCardRequest{
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
	assert.Equal(t, "updated successfully", result)

	client.DeleteUser()
}

func TestUpdateBankCardFail_NotAuthorized(t *testing.T) {
	setEnv(t)
	client := NewClient()
	assert.Empty(t, client.token)

	result := client.UpdateBankCard(
		&pb.UpdateBankCardRequest{
			BankCard: &pb.BankCard{
				Title:      algorithms.RandString(10),
				CardHolder: algorithms.RandString(10),
				CardNumber: algorithms.RandString(10),
				CardExpire: algorithms.RandString(10),
				CardCvv:    algorithms.RandString(10),
				Meta:       algorithms.RandString(10),
			},
		},
	)
	assert.Equal(t, "not authorized", result)
}

func TestUpdateBankCardFail_UserNotFound(t *testing.T) {
	setEnv(t)
	client := NewClient()
	fakeToken := hex.EncodeToString([]byte(uuid.New().String()))
	client.token = fakeToken

	result := client.UpdateBankCard(
		&pb.UpdateBankCardRequest{
			BankCard: &pb.BankCard{
				Title:      algorithms.RandString(10),
				CardHolder: algorithms.RandString(10),
				CardNumber: algorithms.RandString(10),
				CardExpire: algorithms.RandString(10),
				CardCvv:    algorithms.RandString(10),
				Meta:       algorithms.RandString(10),
			},
		},
	)
	assert.Equal(t, "user not found", result)
}

func TestUpdateBankCardFail_EntryNotExists(t *testing.T) {
	setEnv(t)
	client := NewClient()
	_ = createAuthorizedUser(t, client)
	assert.NotEmpty(t, client.token)

	result := client.UpdateBankCard(
		&pb.UpdateBankCardRequest{
			BankCard: &pb.BankCard{
				Title:      algorithms.RandString(10),
				CardHolder: algorithms.RandString(10),
				CardNumber: algorithms.RandString(10),
				CardExpire: algorithms.RandString(10),
				CardCvv:    algorithms.RandString(10),
				Meta:       algorithms.RandString(10),
			},
		},
	)
	assert.Equal(t, "entry with this title does not exist", result)

	client.DeleteUser()
}
