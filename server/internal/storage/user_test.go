package storage

import (
	"context"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"gophkeeper/pkg/algorithms"
)

func TestCreateUserSuccess(t *testing.T) {

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
}

func TestCreateUserFail_AlreadyExists(t *testing.T) {

	defer Downgrade()

	storage := NewStorage()
	ctx := context.Background()
	login := algorithms.RandString(10)
	password := algorithms.RandString(10)

	token, err := storage.CreateUser(ctx, login, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	token, err = storage.CreateUser(ctx, login, password)
	assert.Empty(t, token)
	assert.Equal(t, fmt.Sprintf("%s already exists", login), err.Error())
}

func TestGetUserIDByTokenSuccess(t *testing.T) {

	defer Downgrade()

	storage := NewStorage()
	ctx := context.Background()
	login := algorithms.RandString(10)
	password := algorithms.RandString(10)

	token, err := storage.CreateUser(ctx, login, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	userID, err := storage.GetUserIDByToken(ctx, token)
	assert.NoError(t, err)
	assert.NotEmpty(t, userID)
}

func TestGetUserIDByTokenFail_UserNotFound(t *testing.T) {

	defer Downgrade()

	storage := NewStorage()
	fakeToken := hex.EncodeToString([]byte(uuid.New().String()))

	userID, err := storage.GetUserIDByToken(context.Background(), fakeToken)
	assert.Empty(t, userID)
	assert.Equal(t, "user not found", err.Error())
}

func TestGetTokenByCredentialsSuccess(t *testing.T) {

	defer Downgrade()

	storage := NewStorage()
	ctx := context.Background()
	login := algorithms.RandString(10)
	password := algorithms.RandString(10)

	token, err := storage.CreateUser(ctx, login, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	result, err := storage.GetTokenByCredentials(ctx, login, password)
	assert.NoError(t, err)
	assert.Equal(t, token, result)
}

func TestGetTokenByCredentialsFail_UserNotFound(t *testing.T) {

	defer Downgrade()

	storage := NewStorage()
	ctx := context.Background()

	token, err := storage.GetTokenByCredentials(
		ctx, algorithms.RandString(10), algorithms.RandString(10),
	)
	assert.Empty(t, token)
	assert.Equal(t, "user not found", err.Error())
}

func TestDeleteUserSuccess(t *testing.T) {

	defer Downgrade()

	storage := NewStorage()
	ctx := context.Background()
	login := algorithms.RandString(10)
	password := algorithms.RandString(10)

	token, err := storage.CreateUser(ctx, login, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	err = storage.DeleteUser(ctx, token)
	assert.NoError(t, err)

	userID, err := storage.GetUserIDByToken(context.Background(), token)
	assert.Empty(t, userID)
	assert.Equal(t, "user not found", err.Error())
}

func TestDeleteUserFail_UserNotFound(t *testing.T) {

	defer Downgrade()

	storage := NewStorage()
	ctx := context.Background()
	fakeToken := hex.EncodeToString([]byte(uuid.New().String()))

	err := storage.DeleteUser(ctx, fakeToken)
	assert.Equal(t, "user not found", err.Error())
}
