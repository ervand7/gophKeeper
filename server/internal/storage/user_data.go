package storage

import (
	"context"
	"database/sql"

	pb "gophkeeper/proto"
	"gophkeeper/server/internal/logger"
)

func (s *Storage) FindUserData(ctx context.Context, userID string) (*pb.UserData, error) {
	userData := new(pb.UserData)
	userData, err := s.bankCard(ctx, userID, userData)
	if err != nil {
		return nil, err
	}
	userData, err = s.binaryData(ctx, userID, userData)
	if err != nil {
		return nil, err
	}
	userData, err = s.credentials(ctx, userID, userData)
	if err != nil {
		return nil, err
	}
	userData, err = s.text(ctx, userID, userData)
	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (s *Storage) bankCard(
	ctx context.Context, userID string, userData *pb.UserData,
) (*pb.UserData, error) {
	query := `
		select "title", "card_holder", "card_number", 
		"card_expire", "card_cvv", "created_at", "updated_at", "meta"  
		from "public"."bank_card" where "user_id" = $1
	`
	var (
		title, cardHolder, cardNumber, cardExpire,
		cardCvv, createdAt, updatedAt string
		metaRaw sql.NullString
	)
	rows, err := s.receiveRows(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer s.db.closeRows(rows)

	for rows.Next() {
		err = rows.Scan(
			&title, &cardHolder, &cardNumber, &cardExpire,
			&cardCvv, &createdAt, &updatedAt, &metaRaw,
		)
		if err != nil {
			return nil, err
		}
		entry := pb.BankCard{
			Title: title, CardHolder: cardHolder, CardNumber: cardNumber,
			CardExpire: cardExpire, CardCvv: cardCvv, CreatedAt: createdAt,
			UpdatedAt: updatedAt, Meta: s.getMeta(metaRaw),
		}
		userData.BankCard = append(userData.BankCard, &entry)
	}

	return userData, nil
}

func (s *Storage) binaryData(
	ctx context.Context, userID string, userData *pb.UserData,
) (*pb.UserData, error) {
	query := `
		select "title", "content", "created_at", "updated_at", "meta"  
		from "public"."binary_data" where "user_id" = $1
	`
	var (
		title, createdAt, updatedAt string
		metaRaw                     sql.NullString
		content                     []byte
	)
	rows, err := s.receiveRows(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer s.db.closeRows(rows)

	for rows.Next() {
		err = rows.Scan(
			&title, &content, &createdAt, &updatedAt, &metaRaw,
		)
		if err != nil {
			return nil, err
		}
		entry := pb.BinaryData{
			Title: title, Content: content, CreatedAt: createdAt,
			UpdatedAt: updatedAt, Meta: s.getMeta(metaRaw),
		}
		userData.BinaryData = append(userData.BinaryData, &entry)
	}

	return userData, nil
}

func (s *Storage) credentials(
	ctx context.Context, userID string, userData *pb.UserData,
) (*pb.UserData, error) {
	query := `
		select "title", "login", "password", "created_at", "updated_at", "meta"  
		from "public"."credentials" where "user_id" = $1
	`
	var (
		title, login, password, createdAt, updatedAt string
		metaRaw                                      sql.NullString
	)
	rows, err := s.receiveRows(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer s.db.closeRows(rows)

	for rows.Next() {
		err = rows.Scan(
			&title, &login, &password, &createdAt, &updatedAt, &metaRaw,
		)
		if err != nil {
			return nil, err
		}
		entry := pb.Credentials{
			Title: title, Login: login, Password: password,
			CreatedAt: createdAt, UpdatedAt: updatedAt, Meta: s.getMeta(metaRaw),
		}
		userData.Credentials = append(userData.Credentials, &entry)
	}

	return userData, nil
}

func (s *Storage) text(
	ctx context.Context, userID string, userData *pb.UserData,
) (*pb.UserData, error) {
	query := `
		select "title", "content", "created_at", "updated_at", "meta"  
		from "public"."text" where "user_id" = $1
	`
	var (
		title, content, createdAt, updatedAt string
		metaRaw                              sql.NullString
	)
	rows, err := s.receiveRows(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer s.db.closeRows(rows)

	for rows.Next() {
		err = rows.Scan(
			&title, &content, &createdAt, &updatedAt, &metaRaw,
		)
		if err != nil {
			logger.Logger.Fatal(err.Error())
		}
		entry := pb.Text{
			Title: title, Content: content, CreatedAt: createdAt,
			UpdatedAt: updatedAt, Meta: s.getMeta(metaRaw),
		}
		userData.Text = append(userData.Text, &entry)
	}

	return userData, nil
}
