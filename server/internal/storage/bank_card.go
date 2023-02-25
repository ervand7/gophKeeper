package storage

import (
	"context"
	"time"

	pb "gophkeeper/proto"
	"gophkeeper/server/internal/exceptions"
)

func (s *Storage) CreateBankCard(
	ctx context.Context, userID string, b *pb.BankCard,
) error {
	query := `
		insert into "public"."bank_card" (
			"user_id", 
			"title", 
			"card_holder", 
			"card_number", 
			"card_expire",
			"card_cvv",
			"meta" 
		)
		values ($1, $2, $3, $4, $5, $6, $7)
		on conflict ("user_id", "title") do nothing
		returning "title";
	`
	rows, err := s.receiveRows(
		ctx, query, userID, b.Title, b.CardHolder,
		b.CardNumber, b.CardExpire, b.CardCvv, b.Meta,
	)
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

func (s *Storage) UpdateBankCard(
	ctx context.Context, userID string, b *pb.BankCard,
) error {
	query := `
		update "public"."bank_card"
		set "card_holder" = $3,
			"card_number" = $4,
			"card_expire" = $5,
			"card_cvv"    = $6,
			"meta"        = $7, 
			"updated_at"  = $8 
		where "user_id" = $1 and "title" = $2
		returning "title";
	`
	rows, err := s.receiveRows(
		ctx, query, userID, b.Title, b.CardHolder,
		b.CardNumber, b.CardExpire, b.CardCvv, b.Meta, time.Now(),
	)
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

func (s *Storage) DeleteBankCard(
	ctx context.Context, userID, title string,
) error {
	query := `
		delete from "public"."bank_card"
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
