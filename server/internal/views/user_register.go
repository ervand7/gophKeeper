package views

import (
	"context"
	"encoding/hex"
	"fmt"

	"gophkeeper/server/internal/logger"

	pb "gophkeeper/proto"
)

func (s *GophKeeperServer) Register(
	ctx context.Context, r *pb.RegisterRequest,
) (*pb.RegisterResponse, error) {
	token, err := s.Storage.CreateUser(ctx, r.Login, s.hash256(r.Password))
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.RegisterResponse{Token: ""}, err
	}

	encodedToken := hex.EncodeToString([]byte(token))
	logger.Logger.Info(
		fmt.Sprintf("success GophKeeperServer.Register %s", r.Login),
	)
	return &pb.RegisterResponse{Token: encodedToken}, nil
}
