package views

import (
	"context"

	pb "gophkeeper/proto"
	"gophkeeper/server/internal/logger"
)

func (s *GophKeeperServer) CreateBinaryData(
	ctx context.Context, r *pb.CreateBinaryDataRequest,
) (*pb.CreateBinaryDataResponse, error) {
	token, err := s.getTokenFromMetadata(ctx)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.CreateBinaryDataResponse{}, err
	}

	userID, err := s.Storage.GetUserIDByToken(ctx, token)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.CreateBinaryDataResponse{}, err
	}

	err = s.Storage.CreateBinaryData(ctx, userID, r.BinaryData)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.CreateBinaryDataResponse{}, err
	}

	return &pb.CreateBinaryDataResponse{}, nil
}
