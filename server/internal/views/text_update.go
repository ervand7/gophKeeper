package views

import (
	"context"

	pb "gophkeeper/proto"
	"gophkeeper/server/internal/logger"
)

func (s *GophKeeperServer) UpdateText(
	ctx context.Context, r *pb.UpdateTextRequest,
) (*pb.UpdateTextResponse, error) {
	token, err := s.getTokenFromMetadata(ctx)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.UpdateTextResponse{}, err
	}

	userID, err := s.Storage.GetUserIDByToken(ctx, token)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.UpdateTextResponse{}, err
	}

	err = s.Storage.UpdateText(ctx, userID, r.Text)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.UpdateTextResponse{}, err
	}

	return &pb.UpdateTextResponse{}, nil
}
