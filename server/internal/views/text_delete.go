package views

import (
	"context"

	pb "gophkeeper/proto"
	"gophkeeper/server/internal/logger"
)

func (s *GophKeeperServer) DeleteText(
	ctx context.Context, r *pb.DeleteTextRequest,
) (*pb.DeleteTextResponse, error) {
	token, err := s.getTokenFromMetadata(ctx)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.DeleteTextResponse{}, err
	}

	userID, err := s.Storage.GetUserIDByToken(ctx, token)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.DeleteTextResponse{}, err
	}

	err = s.Storage.DeleteText(ctx, userID, r.Title)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.DeleteTextResponse{}, err
	}

	return &pb.DeleteTextResponse{}, nil
}
