package storage

import (
	"context"
	"time"

	pb "gophkeeper/proto"
	"gophkeeper/server/internal/exceptions"
)

func (s *Storage) CreateBinaryData(
	ctx context.Context, userID string, b *pb.BinaryData,
) error {
	query := `
		insert into "public"."binary_data" (
			"user_id", "title", "content", "meta"
		)
		values ($1, $2, $3, $4)
		on conflict ("user_id", "title") do nothing
		returning "title";
	`
	rows, err := s.receiveRows(ctx, query, userID, b.Title, b.Content, b.Meta)
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

func (s *Storage) UpdateBinaryData(
	ctx context.Context, userID string, b *pb.BinaryData,
) error {
	query := `
		update "public"."binary_data"
		set "content"    = $3,
			"meta"       = $4,
			"updated_at" = $5
		where "user_id" = $1 and "title" = $2
		returning "title";
	`
	rows, err := s.receiveRows(ctx, query, userID, b.Title, b.Content, b.Meta, time.Now())
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

func (s *Storage) DeleteBinaryData(
	ctx context.Context, userID, title string,
) error {
	query := `
		delete from "public"."binary_data"
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
