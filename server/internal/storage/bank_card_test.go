package storage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"gophkeeper/pkg/algorithms"
	pb "gophkeeper/proto"
)

func TestCreateBankCardSuccess(t *testing.T) {
	defer Downgrade()

	storage := NewStorage()
	ctx := context.Background()
	login := algorithms.RandString(10)
	password := algorithms.RandString(10)

	token, err := storage.CreateUser(ctx, login, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	userID, err := storage.GetUserIDByToken(context.Background(), token)
	assert.NoError(t, err)
	assert.NotEmpty(t, userID)

	title := algorithms.RandString(10)
	err = storage.CreateBankCard(
		ctx,
		userID,
		&pb.BankCard{
			Title:      title,
			CardHolder: algorithms.RandString(10),
			CardNumber: algorithms.RandString(10),
			CardExpire: algorithms.RandString(10),
			CardCvv:    algorithms.RandString(10),
			Meta:       algorithms.RandString(10),
		},
	)
	assert.NoError(t, err)

	rows, err := storage.receiveRows(ctx, `select "title" from "public"."bank_card"`)
	assert.NoError(t, err)
	result, err := storage.getValueFromRows(rows)
	assert.NoError(t, err)
	assert.Equal(t, title, result)
}

func TestCreateBankCardFail_AlreadyExists(t *testing.T) {
	defer Downgrade()

	storage := NewStorage()
	ctx := context.Background()
	login := algorithms.RandString(10)
	password := algorithms.RandString(10)

	token, err := storage.CreateUser(ctx, login, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	userID, err := storage.GetUserIDByToken(context.Background(), token)
	assert.NoError(t, err)
	assert.NotEmpty(t, userID)

	title := algorithms.RandString(10)
	err = storage.CreateBankCard(
		ctx,
		userID,
		&pb.BankCard{
			Title:      title,
			CardHolder: algorithms.RandString(10),
			CardNumber: algorithms.RandString(10),
			CardExpire: algorithms.RandString(10),
			CardCvv:    algorithms.RandString(10),
			Meta:       algorithms.RandString(10),
		},
	)
	assert.NoError(t, err)

	err = storage.CreateBankCard(
		ctx,
		userID,
		&pb.BankCard{
			Title:      title,
			CardHolder: algorithms.RandString(10),
			CardNumber: algorithms.RandString(10),
			CardExpire: algorithms.RandString(10),
			CardCvv:    algorithms.RandString(10),
			Meta:       algorithms.RandString(10),
		},
	)
	assert.Equal(t, "entry with this title already exists", err.Error())
}

func TestUpdateBankCardSuccess(t *testing.T) {
	defer Downgrade()

	storage := NewStorage()
	ctx := context.Background()
	login := algorithms.RandString(10)
	password := algorithms.RandString(10)

	token, err := storage.CreateUser(ctx, login, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	userID, err := storage.GetUserIDByToken(context.Background(), token)
	assert.NoError(t, err)
	assert.NotEmpty(t, userID)

	title := algorithms.RandString(10)
	err = storage.CreateBankCard(
		ctx,
		userID,
		&pb.BankCard{
			Title:      title,
			CardHolder: algorithms.RandString(10),
			CardNumber: algorithms.RandString(10),
			CardExpire: algorithms.RandString(10),
			CardCvv:    algorithms.RandString(10),
			Meta:       algorithms.RandString(10),
		},
	)
	assert.NoError(t, err)

	rows, err := storage.receiveRows(ctx, `select "updated_at" from "public"."bank_card"`)
	assert.NoError(t, err)
	updatedAtBefore, err := storage.getValueFromRows(rows)
	assert.NoError(t, err)

	err = storage.UpdateBankCard(
		ctx,
		userID,
		&pb.BankCard{
			Title:      title,
			CardHolder: algorithms.RandString(10),
			CardNumber: algorithms.RandString(10),
			CardExpire: algorithms.RandString(10),
			CardCvv:    algorithms.RandString(10),
			Meta:       algorithms.RandString(10),
		},
	)
	assert.NoError(t, err)

	rows, err = storage.receiveRows(ctx, `select "updated_at" from "public"."bank_card"`)
	assert.NoError(t, err)
	updatedAtAfter, err := storage.getValueFromRows(rows)
	assert.NoError(t, err)

	assert.NotEqual(t, updatedAtBefore, updatedAtAfter)
}

func TestUpdateBankCardFail_NotExists(t *testing.T) {
	defer Downgrade()

	storage := NewStorage()
	ctx := context.Background()
	login := algorithms.RandString(10)
	password := algorithms.RandString(10)

	token, err := storage.CreateUser(ctx, login, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	userID, err := storage.GetUserIDByToken(context.Background(), token)
	assert.NoError(t, err)
	assert.NotEmpty(t, userID)

	err = storage.UpdateBankCard(
		ctx,
		userID,
		&pb.BankCard{
			Title:      algorithms.RandString(10),
			CardHolder: algorithms.RandString(10),
			CardNumber: algorithms.RandString(10),
			CardExpire: algorithms.RandString(10),
			CardCvv:    algorithms.RandString(10),
			Meta:       algorithms.RandString(10),
		},
	)
	assert.Equal(t, "entry with this title does not exist", err.Error())
}

func TestDeleteBankCardSuccess(t *testing.T) {
	defer Downgrade()

	storage := NewStorage()
	ctx := context.Background()
	login := algorithms.RandString(10)
	password := algorithms.RandString(10)

	token, err := storage.CreateUser(ctx, login, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	userID, err := storage.GetUserIDByToken(context.Background(), token)
	assert.NoError(t, err)
	assert.NotEmpty(t, userID)

	title := algorithms.RandString(10)
	err = storage.CreateBankCard(
		ctx,
		userID,
		&pb.BankCard{
			Title:      title,
			CardHolder: algorithms.RandString(10),
			CardNumber: algorithms.RandString(10),
			CardExpire: algorithms.RandString(10),
			CardCvv:    algorithms.RandString(10),
			Meta:       algorithms.RandString(10),
		},
	)
	assert.NoError(t, err)

	rows, err := storage.receiveRows(ctx, `select "title" from "public"."bank_card"`)
	assert.NoError(t, err)
	result, err := storage.getValueFromRows(rows)
	assert.NoError(t, err)
	assert.Equal(t, title, result)

	err = storage.DeleteBankCard(ctx, userID, title)
	assert.NoError(t, err)

	rows, err = storage.receiveRows(ctx, `select "title" from "public"."bank_card"`)
	assert.NoError(t, err)
	title, err = storage.getValueFromRows(rows)
	assert.NoError(t, err)
	assert.Empty(t, title)
}

func TestDeleteBankCardFail_NotExists(t *testing.T) {
	defer Downgrade()

	storage := NewStorage()
	ctx := context.Background()
	login := algorithms.RandString(10)
	password := algorithms.RandString(10)

	token, err := storage.CreateUser(ctx, login, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	userID, err := storage.GetUserIDByToken(context.Background(), token)
	assert.NoError(t, err)
	assert.NotEmpty(t, userID)

	err = storage.DeleteBankCard(ctx, userID, algorithms.RandString(10))
	assert.Equal(t, "entry with this title does not exist", err.Error())
}
