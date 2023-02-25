package storage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"gophkeeper/pkg/algorithms"
	pb "gophkeeper/proto"
)

func TestCreateBinaryDataSuccess(t *testing.T) {
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
	err = storage.CreateBinaryData(
		ctx,
		userID,
		&pb.BinaryData{
			Title:   title,
			Content: []byte(algorithms.RandString(10)),
			Meta:    algorithms.RandString(10),
		},
	)
	assert.NoError(t, err)

	rows, err := storage.receiveRows(ctx, `select "title" from "public"."binary_data"`)
	assert.NoError(t, err)
	result, err := storage.getValueFromRows(rows)
	assert.NoError(t, err)
	assert.Equal(t, title, result)
}

func TestCreateBinaryDataFail_AlreadyExists(t *testing.T) {
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
	err = storage.CreateBinaryData(
		ctx,
		userID,
		&pb.BinaryData{
			Title:   title,
			Content: []byte(algorithms.RandString(10)),
			Meta:    algorithms.RandString(10),
		},
	)
	assert.NoError(t, err)

	err = storage.CreateBinaryData(
		ctx,
		userID,
		&pb.BinaryData{
			Title:   title,
			Content: []byte(algorithms.RandString(10)),
			Meta:    algorithms.RandString(10),
		},
	)
	assert.Equal(t, "entry with this title already exists", err.Error())
}

func TestUpdateBinaryDataSuccess(t *testing.T) {
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
	err = storage.CreateBinaryData(
		ctx,
		userID,
		&pb.BinaryData{
			Title:   title,
			Content: []byte(algorithms.RandString(10)),
			Meta:    algorithms.RandString(10),
		},
	)
	assert.NoError(t, err)

	rows, err := storage.receiveRows(ctx, `select "updated_at" from "public"."binary_data"`)
	assert.NoError(t, err)
	updatedAtBefore, err := storage.getValueFromRows(rows)
	assert.NoError(t, err)

	err = storage.UpdateBinaryData(
		ctx,
		userID,
		&pb.BinaryData{
			Title:   title,
			Content: []byte(algorithms.RandString(10)),
			Meta:    algorithms.RandString(10),
		},
	)
	assert.NoError(t, err)

	rows, err = storage.receiveRows(ctx, `select "updated_at" from "public"."binary_data"`)
	assert.NoError(t, err)
	updatedAtAfter, err := storage.getValueFromRows(rows)
	assert.NoError(t, err)

	assert.NotEqual(t, updatedAtBefore, updatedAtAfter)
}

func TestUpdateBinaryDataFail_NotExists(t *testing.T) {
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

	err = storage.UpdateBinaryData(
		ctx,
		userID,
		&pb.BinaryData{
			Title:   algorithms.RandString(10),
			Content: []byte(algorithms.RandString(10)),
			Meta:    algorithms.RandString(10),
		},
	)
	assert.Equal(t, "entry with this title does not exist", err.Error())
}

func TestDeleteBinaryDataSuccess(t *testing.T) {
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
	err = storage.CreateBinaryData(
		ctx,
		userID,
		&pb.BinaryData{
			Title:   title,
			Content: []byte(algorithms.RandString(10)),
			Meta:    algorithms.RandString(10),
		},
	)
	assert.NoError(t, err)

	rows, err := storage.receiveRows(ctx, `select "title" from "public"."binary_data"`)
	assert.NoError(t, err)
	result, err := storage.getValueFromRows(rows)
	assert.NoError(t, err)
	assert.Equal(t, title, result)

	err = storage.DeleteBinaryData(ctx, userID, title)
	assert.NoError(t, err)

	rows, err = storage.receiveRows(ctx, `select "title" from "public"."binary_data"`)
	assert.NoError(t, err)
	title, err = storage.getValueFromRows(rows)
	assert.NoError(t, err)
	assert.Empty(t, title)
}

func TestDeleteBinaryDataFail_NotExists(t *testing.T) {
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

	err = storage.DeleteBinaryData(ctx, userID, algorithms.RandString(10))
	assert.Equal(t, "entry with this title does not exist", err.Error())
}
