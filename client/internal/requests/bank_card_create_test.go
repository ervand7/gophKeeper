package requests

import (
	"encoding/hex"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"gophkeeper/pkg/algorithms"
	pb "gophkeeper/proto"
)

func TestCreateBankCardSuccess(t *testing.T) {
	setEnv(t)
	client := NewClient()
	_ = createAuthorizedUser(t, client)
	assert.NotEmpty(t, client.token)

	result := client.CreateBankCard(
		&pb.CreateBankCardRequest{
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
	assert.Equal(t, "created successfully", result)

	time.Sleep(1 * time.Second)
	client.DeleteUser()
}

func TestCreateBankCardFail_NotAuthorized(t *testing.T) {
	setEnv(t)
	client := NewClient()
	assert.Empty(t, client.token)

	result := client.CreateBankCard(
		&pb.CreateBankCardRequest{
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

func TestCreateBankCardFail_UserNotFound(t *testing.T) {
	setEnv(t)
	client := NewClient()
	fakeToken := hex.EncodeToString([]byte(uuid.New().String()))
	client.token = fakeToken

	result := client.CreateBankCard(
		&pb.CreateBankCardRequest{
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

func TestCreateBankCardFail_AlreadyExists(t *testing.T) {
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

	result = client.CreateBankCard(
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
	assert.Equal(t, "entry with this title already exists", result)

	client.DeleteUser()
}
