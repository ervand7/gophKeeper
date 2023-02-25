package views

import (
	"context"

	pb "gophkeeper/proto"
	"gophkeeper/server/internal/logger"
)

func (s *GophKeeperServer) UpdateBankCard(
	ctx context.Context, r *pb.UpdateBankCardRequest,
) (*pb.UpdateBankCardResponse, error) {
	token, err := s.getTokenFromMetadata(ctx)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.UpdateBankCardResponse{}, err
	}

	userID, err := s.Storage.GetUserIDByToken(ctx, token)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.UpdateBankCardResponse{}, err
	}

	err = s.Storage.UpdateBankCard(ctx, userID, r.BankCard)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.UpdateBankCardResponse{}, err
	}

	return &pb.UpdateBankCardResponse{}, nil
}
