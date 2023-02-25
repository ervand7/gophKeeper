package views

import (
	"context"
	"encoding/hex"
	"fmt"

	pb "gophkeeper/proto"
	"gophkeeper/server/internal/logger"
)

func (s *GophKeeperServer) Auth(
	ctx context.Context, r *pb.AuthRequest,
) (*pb.AuthResponse, error) {
	token, err := s.Storage.GetTokenByCredentials(
		ctx, r.Login, s.hash256(r.Password),
	)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &pb.AuthResponse{Token: ""}, err
	}

	encodedToken := hex.EncodeToString([]byte(token))
	logger.Logger.Info(
		fmt.Sprintf("success GophKeeperServer.Login %s", r.Login),
	)
	return &pb.AuthResponse{Token: encodedToken}, nil
}
