package views

import (
	"context"

	pb "gophkeeper/proto"
	"gophkeeper/server/internal/logger"
)

func (s *GophKeeperServer) CreateText(
	ctx context.Context, r *pb.CreateTextRequest,
) (*pb.CreateTextResponse, error) {
	token, err := s.getTokenFromMetadata(ctx)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.CreateTextResponse{}, err
	}

	userID, err := s.Storage.GetUserIDByToken(ctx, token)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.CreateTextResponse{}, err
	}

	err = s.Storage.CreateText(ctx, userID, r.Text)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.CreateTextResponse{}, err
	}

	return &pb.CreateTextResponse{}, nil
}
