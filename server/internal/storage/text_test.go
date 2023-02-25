package storage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"gophkeeper/pkg/algorithms"
	pb "gophkeeper/proto"
)

func TestCreateTextSuccess(t *testing.T) {
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
	err = storage.CreateText(
		ctx,
		userID,
		&pb.Text{
			Title:   title,
			Content: algorithms.RandString(10),
			Meta:    algorithms.RandString(10),
		},
	)
	assert.NoError(t, err)

	rows, err := storage.receiveRows(ctx, `select "title" from "public"."text"`)
	assert.NoError(t, err)
	result, err := storage.getValueFromRows(rows)
	assert.NoError(t, err)
	assert.Equal(t, title, result)
}

func TestCreateTextFail_AlreadyExists(t *testing.T) {
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
	err = storage.CreateText(
		ctx,
		userID,
		&pb.Text{
			Title:   title,
			Content: algorithms.RandString(10),
			Meta:    algorithms.RandString(10),
		},
	)
	assert.NoError(t, err)

	err = storage.CreateText(
		ctx,
		userID,
		&pb.Text{
			Title:   title,
			Content: algorithms.RandString(10),
			Meta:    algorithms.RandString(10),
		},
	)
	assert.Equal(t, "entry with this title already exists", err.Error())
}

func TestUpdateTextSuccess(t *testing.T) {
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
	err = storage.CreateText(
		ctx,
		userID,
		&pb.Text{
			Title:   title,
			Content: algorithms.RandString(10),
			Meta:    algorithms.RandString(10),
		},
	)
	assert.NoError(t, err)

	rows, err := storage.receiveRows(ctx, `select "updated_at" from "public"."text"`)
	assert.NoError(t, err)
	updatedAtBefore, err := storage.getValueFromRows(rows)
	assert.NoError(t, err)

	err = storage.UpdateText(
		ctx,
		userID,
		&pb.Text{
			Title:   title,
			Content: algorithms.RandString(10),
			Meta:    algorithms.RandString(10),
		},
	)
	assert.NoError(t, err)

	rows, err = storage.receiveRows(ctx, `select "updated_at" from "public"."text"`)
	assert.NoError(t, err)
	updatedAtAfter, err := storage.getValueFromRows(rows)
	assert.NoError(t, err)

	assert.NotEqual(t, updatedAtBefore, updatedAtAfter)
}

func TestUpdateTextFail_NotExists(t *testing.T) {
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

	err = storage.UpdateText(
		ctx,
		userID,
		&pb.Text{
			Title:   algorithms.RandString(10),
			Content: algorithms.RandString(10),
			Meta:    algorithms.RandString(10),
		},
	)
	assert.Equal(t, "entry with this title does not exist", err.Error())
}

func TestDeleteTextSuccess(t *testing.T) {
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
	err = storage.CreateText(
		ctx,
		userID,
		&pb.Text{
			Title:   title,
			Content: algorithms.RandString(10),
			Meta:    algorithms.RandString(10),
		},
	)
	assert.NoError(t, err)

	rows, err := storage.receiveRows(ctx, `select "title" from "public"."text"`)
	assert.NoError(t, err)
	result, err := storage.getValueFromRows(rows)
	assert.NoError(t, err)
	assert.Equal(t, title, result)

	err = storage.DeleteText(ctx, userID, title)
	assert.NoError(t, err)

	rows, err = storage.receiveRows(ctx, `select "title" from "public"."text"`)
	assert.NoError(t, err)
	title, err = storage.getValueFromRows(rows)
	assert.NoError(t, err)
	assert.Empty(t, title)
}

func TestDeleteTextFail_NotExists(t *testing.T) {
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

	err = storage.DeleteText(ctx, userID, algorithms.RandString(10))
	assert.Equal(t, "entry with this title does not exist", err.Error())
}
