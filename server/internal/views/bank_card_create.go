package views

import (
	"context"

	pb "gophkeeper/proto"
	"gophkeeper/server/internal/logger"
)

func (s *GophKeeperServer) CreateBankCard(
	ctx context.Context, r *pb.CreateBankCardRequest,
) (*pb.CreateBankCardResponse, error) {
	token, err := s.getTokenFromMetadata(ctx)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.CreateBankCardResponse{}, err
	}

	userID, err := s.Storage.GetUserIDByToken(ctx, token)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.CreateBankCardResponse{}, err
	}

	err = s.Storage.CreateBankCard(ctx, userID, r.BankCard)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.CreateBankCardResponse{}, err
	}

	return &pb.CreateBankCardResponse{}, nil
}
