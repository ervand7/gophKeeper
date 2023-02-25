package views

import (
	"context"

	pb "gophkeeper/proto"
	"gophkeeper/server/internal/logger"
)

func (s *GophKeeperServer) DeleteCredentials(
	ctx context.Context, r *pb.DeleteCredentialsRequest,
) (*pb.DeleteCredentialsResponse, error) {
	token, err := s.getTokenFromMetadata(ctx)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.DeleteCredentialsResponse{}, err
	}

	userID, err := s.Storage.GetUserIDByToken(ctx, token)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.DeleteCredentialsResponse{}, err
	}

	err = s.Storage.DeleteCredentials(ctx, userID, r.Title)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.DeleteCredentialsResponse{}, err
	}

	return &pb.DeleteCredentialsResponse{}, nil
}
