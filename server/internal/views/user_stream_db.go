package views

import (
	"time"

	pb "gophkeeper/proto"
	"gophkeeper/server/internal/logger"
)

// StreamDB opens stream between server and client.
func (s *GophKeeperServer) StreamDB(
	_ *pb.StreamDBRequest, streamServer pb.GophKeeper_StreamDBServer,
) error {
	ctx := streamServer.Context()
	token, err := s.getTokenFromMetadata(ctx)
	if err != nil {
		logger.Logger.Error(err.Error())
		return err
	}

	userID, err := s.Storage.GetUserIDByToken(ctx, token)
	if err != nil {
		logger.Logger.Error(err.Error())
		return err
	}

	for {
		userData, err := s.Storage.FindUserData(ctx, userID)
		if err != nil {
			logger.Logger.Error(err.Error())
			return err
		}

		err = streamServer.Send(
			&pb.StreamDBResponse{UserData: userData},
		)
		if err != nil {
			logger.Logger.Error(err.Error())
			return err
		}

		time.Sleep(time.Second * dataUpdateTimeoutSecond)
	}
}
