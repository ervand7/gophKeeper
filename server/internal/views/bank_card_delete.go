package views

import (
	"context"

	pb "gophkeeper/proto"
	"gophkeeper/server/internal/logger"
)

func (s *GophKeeperServer) DeleteBankCard(
	ctx context.Context, r *pb.DeleteBankCardRequest,
) (*pb.DeleteBankCardResponse, error) {
	token, err := s.getTokenFromMetadata(ctx)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.DeleteBankCardResponse{}, err
	}

	userID, err := s.Storage.GetUserIDByToken(ctx, token)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.DeleteBankCardResponse{}, err
	}

	err = s.Storage.DeleteBankCard(ctx, userID, r.Title)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.DeleteBankCardResponse{}, err
	}

	return &pb.DeleteBankCardResponse{}, nil
}
