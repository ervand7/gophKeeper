package views

import (
	"context"

	pb "gophkeeper/proto"
	"gophkeeper/server/internal/logger"
)

func (s *GophKeeperServer) CreateCredentials(
	ctx context.Context, r *pb.CreateCredentialsRequest,
) (*pb.CreateCredentialsResponse, error) {
	token, err := s.getTokenFromMetadata(ctx)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.CreateCredentialsResponse{}, err
	}

	userID, err := s.Storage.GetUserIDByToken(ctx, token)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.CreateCredentialsResponse{}, err
	}

	err = s.Storage.CreateCredentials(ctx, userID, r.Credentials)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.CreateCredentialsResponse{}, err
	}

	return &pb.CreateCredentialsResponse{}, nil
}
