package storage

import (
	"context"
	"time"

	pb "gophkeeper/proto"
	"gophkeeper/server/internal/exceptions"
)

func (s *Storage) CreateCredentials(
	ctx context.Context, userID string, c *pb.Credentials,
) error {
	query := `
		insert into "public"."credentials" (
			"user_id", 
			"title", 
			"login", 
			"password", 
			"meta"
		)
		values ($1, $2, $3, $4, $5)
		on conflict ("user_id", "title") do nothing
		returning "title";
	`
	rows, err := s.receiveRows(ctx, query, userID, c.Title, c.Login, c.Password, c.Meta)
	if err != nil {
		return err
	}
	title, err := s.getValueFromRows(rows)
	if err != nil {
		return err
	}
	if title == "" {
		return exceptions.NewErrorAlreadyExists()
	}

	return nil
}

func (s *Storage) UpdateCredentials(
	ctx context.Context, userID string, c *pb.Credentials,
) error {
	query := `
		update "public"."credentials"
		set "login"      = $3,
			"password"   = $4,
			"meta"       = $5, 
			"updated_at" = $6
		where "user_id" = $1 and "title" = $2
		returning "title";
	`
	rows, err := s.receiveRows(ctx, query, userID, c.Title, c.Login, c.Password, c.Meta, time.Now())
	if err != nil {
		return err
	}
	title, err := s.getValueFromRows(rows)
	if err != nil {
		return err
	}
	if title == "" {
		return exceptions.NewErrorEntryNotExists()
	}

	return nil
}

func (s *Storage) DeleteCredentials(
	ctx context.Context, userID, title string,
) error {
	query := `
		delete from "public"."credentials" 
		where "user_id" = $1 and "title" = $2 
		returning "title";
	`
	rows, err := s.receiveRows(ctx, query, userID, title)
	if err != nil {
		return err
	}
	title, err = s.getValueFromRows(rows)
	if err != nil {
		return err
	}
	if title == "" {
		return exceptions.NewErrorEntryNotExists()
	}

	return nil
}
