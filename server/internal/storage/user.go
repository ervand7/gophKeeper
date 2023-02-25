package storage

import (
	"context"

	"github.com/google/uuid"

	"gophkeeper/server/internal/exceptions"
)

func (s *Storage) CreateUser(
	ctx context.Context, login, password string,
) (string, error) {
	query := `
		insert into "public"."user" ("login", "password", "token")
		values ($1, $2, $3)
		on conflict ("login") do nothing
		returning "token";
	`
	token := uuid.New().String()
	rows, err := s.receiveRows(ctx, query, login, password, token)
	if err != nil {
		return "", err
	}
	token, err = s.getValueFromRows(rows)
	if err != nil {
		return "", err
	}
	if token == "" {
		return "", exceptions.NewErrorLoginAlreadyExists(login)
	}

	return token, nil
}

func (s *Storage) GetUserIDByToken(
	ctx context.Context, token string,
) (userID string, err error) {
	query := `
		select "id" from "public"."user"
		where "token" = $1;
	`
	rows, err := s.receiveRows(ctx, query, token)
	if err != nil {
		return "", err
	}
	userID, err = s.getValueFromRows(rows)
	if err != nil {
		return "", err
	}
	if userID == "" {
		return "", exceptions.NewErrorUserNotFound()
	}

	return userID, nil
}

func (s *Storage) GetTokenByCredentials(
	ctx context.Context, login, password string,
) (string, error) {
	query := `
		select "token" from "public"."user" 
		where "login" = $1 and "password" = $2;
	`
	rows, err := s.receiveRows(ctx, query, login, password)
	if err != nil {
		return "", err
	}
	token, err := s.getValueFromRows(rows)
	if err != nil {
		return "", err
	}
	if token == "" {
		return "", exceptions.NewErrorUserNotFound()
	}

	return token, nil
}

func (s *Storage) DeleteUser(
	ctx context.Context, token string,
) error {
	query := `
		delete from "public"."user" where token = $1 
		returning "token";
	`
	rows, err := s.receiveRows(ctx, query, token)
	if err != nil {
		return err
	}
	token, err = s.getValueFromRows(rows)
	if err != nil {
		return err
	}
	if token == "" {
		return exceptions.NewErrorUserNotFound()
	}

	return nil
}
