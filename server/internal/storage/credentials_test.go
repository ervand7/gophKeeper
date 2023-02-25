package storage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"gophkeeper/pkg/algorithms"
	pb "gophkeeper/proto"
)

func TestCreateCredentialsSuccess(t *testing.T) {
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
	err = storage.CreateCredentials(
		ctx,
		userID,
		&pb.Credentials{
			Title:    title,
			Login:    algorithms.RandString(10),
			Password: algorithms.RandString(10),
			Meta:     algorithms.RandString(10),
		},
	)
	assert.NoError(t, err)

	rows, err := storage.receiveRows(ctx, `select "title" from "public"."credentials"`)
	assert.NoError(t, err)
	result, err := storage.getValueFromRows(rows)
	assert.NoError(t, err)
	assert.Equal(t, title, result)
}

func TestCreateCredentialsFail_AlreadyExists(t *testing.T) {
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
	err = storage.CreateCredentials(
		ctx,
		userID,
		&pb.Credentials{
			Title:    title,
			Login:    algorithms.RandString(10),
			Password: algorithms.RandString(10),
			Meta:     algorithms.RandString(10),
		},
	)
	assert.NoError(t, err)

	err = storage.CreateCredentials(
		ctx,
		userID,
		&pb.Credentials{
			Title:    title,
			Login:    algorithms.RandString(10),
			Password: algorithms.RandString(10),
			Meta:     algorithms.RandString(10),
		},
	)
	assert.Equal(t, "entry with this title already exists", err.Error())
}

func TestUpdateCredentialsSuccess(t *testing.T) {
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
	err = storage.CreateCredentials(
		ctx,
		userID,
		&pb.Credentials{
			Title:    title,
			Login:    algorithms.RandString(10),
			Password: algorithms.RandString(10),
			Meta:     algorithms.RandString(10),
		},
	)
	assert.NoError(t, err)

	rows, err := storage.receiveRows(ctx, `select "updated_at" from "public"."credentials"`)
	assert.NoError(t, err)
	updatedAtBefore, err := storage.getValueFromRows(rows)
	assert.NoError(t, err)

	err = storage.UpdateCredentials(
		ctx,
		userID,
		&pb.Credentials{
			Title:    title,
			Login:    algorithms.RandString(10),
			Password: algorithms.RandString(10),
			Meta:     algorithms.RandString(10),
		},
	)
	assert.NoError(t, err)

	rows, err = storage.receiveRows(ctx, `select "updated_at" from "public"."credentials"`)
	assert.NoError(t, err)
	updatedAtAfter, err := storage.getValueFromRows(rows)
	assert.NoError(t, err)

	assert.NotEqual(t, updatedAtBefore, updatedAtAfter)
}

func TestUpdateCredentialsFail_NotExists(t *testing.T) {
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

	err = storage.UpdateCredentials(
		ctx,
		userID,
		&pb.Credentials{
			Title:    algorithms.RandString(10),
			Login:    algorithms.RandString(10),
			Password: algorithms.RandString(10),
			Meta:     algorithms.RandString(10),
		},
	)
	assert.Equal(t, "entry with this title does not exist", err.Error())
}

func TestDeleteCredentialsSuccess(t *testing.T) {
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
	err = storage.CreateCredentials(
		ctx,
		userID,
		&pb.Credentials{
			Title:    title,
			Login:    algorithms.RandString(10),
			Password: algorithms.RandString(10),
			Meta:     algorithms.RandString(10),
		},
	)
	assert.NoError(t, err)

	rows, err := storage.receiveRows(ctx, `select "title" from "public"."credentials"`)
	assert.NoError(t, err)
	result, err := storage.getValueFromRows(rows)
	assert.NoError(t, err)
	assert.Equal(t, title, result)

	err = storage.DeleteCredentials(ctx, userID, title)
	assert.NoError(t, err)

	rows, err = storage.receiveRows(ctx, `select "title" from "public"."credentials"`)
	assert.NoError(t, err)
	title, err = storage.getValueFromRows(rows)
	assert.NoError(t, err)
	assert.Empty(t, title)
}

func TestDeleteCredentialsFail_NotExists(t *testing.T) {
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

	err = storage.DeleteCredentials(ctx, userID, algorithms.RandString(10))
	assert.Equal(t, "entry with this title does not exist", err.Error())
}
