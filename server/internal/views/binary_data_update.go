package views

import (
	"context"

	pb "gophkeeper/proto"
	"gophkeeper/server/internal/logger"
)

func (s *GophKeeperServer) UpdateBinaryData(
	ctx context.Context, r *pb.UpdateBinaryDataRequest,
) (*pb.UpdateBinaryDataResponse, error) {
	token, err := s.getTokenFromMetadata(ctx)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.UpdateBinaryDataResponse{}, err
	}

	userID, err := s.Storage.GetUserIDByToken(ctx, token)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.UpdateBinaryDataResponse{}, err
	}

	err = s.Storage.UpdateBinaryData(ctx, userID, r.BinaryData)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.UpdateBinaryDataResponse{}, err
	}

	return &pb.UpdateBinaryDataResponse{}, nil
}
