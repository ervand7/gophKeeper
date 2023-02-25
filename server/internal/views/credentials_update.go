package views

import (
	"context"

	pb "gophkeeper/proto"
	"gophkeeper/server/internal/logger"
)

func (s *GophKeeperServer) UpdateCredentials(
	ctx context.Context, r *pb.UpdateCredentialsRequest,
) (*pb.UpdateCredentialsResponse, error) {
	token, err := s.getTokenFromMetadata(ctx)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.UpdateCredentialsResponse{}, err
	}

	userID, err := s.Storage.GetUserIDByToken(ctx, token)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.UpdateCredentialsResponse{}, err
	}

	err = s.Storage.UpdateCredentials(ctx, userID, r.Credentials)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.UpdateCredentialsResponse{}, err
	}

	return &pb.UpdateCredentialsResponse{}, nil
}
