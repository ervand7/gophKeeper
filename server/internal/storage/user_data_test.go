package storage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"gophkeeper/pkg/algorithms"
	pb "gophkeeper/proto"
)

func TestFindUserDataSuccess(t *testing.T) {

	defer Downgrade()

	storage := NewStorage()
	ctx := context.Background()
	login := algorithms.RandString(10)
	password := algorithms.RandString(10)

	rows, err := storage.receiveRows(ctx, `select * from "public"."user"`)
	assert.NoError(t, err)
	users, err := storage.getValueFromRows(rows)
	assert.NoError(t, err)
	assert.Empty(t, users)

	token, err := storage.CreateUser(ctx, login, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	userID, err := storage.GetUserIDByToken(ctx, token)
	assert.NoError(t, err)
	assert.NotEmpty(t, userID)

	count := 10
	for i := 0; i < count; i++ {
		err = storage.CreateBankCard(ctx, userID,
			&pb.BankCard{
				Title:      algorithms.RandString(10),
				CardHolder: algorithms.RandString(10),
				CardNumber: algorithms.RandString(10),
				CardExpire: algorithms.RandString(10),
				CardCvv:    algorithms.RandString(10),
				Meta:       algorithms.RandString(10),
			},
		)
		assert.NoError(t, err)

		err = storage.CreateBinaryData(ctx, userID,
			&pb.BinaryData{
				Title:   algorithms.RandString(10),
				Content: []byte(algorithms.RandString(10)),
				Meta:    algorithms.RandString(10),
			},
		)
		assert.NoError(t, err)

		err = storage.CreateCredentials(ctx, userID,
			&pb.Credentials{
				Title:    algorithms.RandString(10),
				Login:    algorithms.RandString(10),
				Password: algorithms.RandString(10),
				Meta:     algorithms.RandString(10),
			},
		)
		assert.NoError(t, err)

		err = storage.CreateText(ctx, userID,
			&pb.Text{
				Title:   algorithms.RandString(10),
				Content: algorithms.RandString(10),
				Meta:    algorithms.RandString(10),
			},
		)
		assert.NoError(t, err)
	}

	result, err := storage.FindUserData(ctx, userID)
	assert.NoError(t, err)
	assert.Len(t, result.BankCard, count)
	assert.Len(t, result.BinaryData, count)
	assert.Len(t, result.Credentials, count)
	assert.Len(t, result.Text, count)
}
