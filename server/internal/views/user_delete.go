package views

import (
	"context"

	"gophkeeper/server/internal/logger"

	pb "gophkeeper/proto"
)

func (s *GophKeeperServer) DeleteUser(
	ctx context.Context, _ *pb.DeleteUserRequest,
) (*pb.DeleteUserResponse, error) {
	token, err := s.getTokenFromMetadata(ctx)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.DeleteUserResponse{}, err
	}

	err = s.Storage.DeleteUser(ctx, token)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.DeleteUserResponse{}, err
	}

	return &pb.DeleteUserResponse{}, nil
}
