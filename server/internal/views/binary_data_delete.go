package views

import (
	"context"

	pb "gophkeeper/proto"
	"gophkeeper/server/internal/logger"
)

func (s *GophKeeperServer) DeleteBinaryData(
	ctx context.Context, r *pb.DeleteBinaryDataRequest,
) (*pb.DeleteBinaryDataResponse, error) {
	token, err := s.getTokenFromMetadata(ctx)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.DeleteBinaryDataResponse{}, err
	}

	userID, err := s.Storage.GetUserIDByToken(ctx, token)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.DeleteBinaryDataResponse{}, err
	}

	err = s.Storage.DeleteBinaryData(ctx, userID, r.Title)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.DeleteBinaryDataResponse{}, err
	}

	return &pb.DeleteBinaryDataResponse{}, nil
}
