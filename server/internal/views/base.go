package views

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"google.golang.org/grpc/metadata"

	pb "gophkeeper/proto"
	"gophkeeper/server/internal/exceptions"
	"gophkeeper/server/internal/logger"
	"gophkeeper/server/internal/storage"
)

const dataUpdateTimeoutSecond = 1

type GophKeeperServer struct {
	pb.UnimplementedGophKeeperServer
	Storage *storage.Storage
}

func (s *GophKeeperServer) getTokenFromMetadata(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", exceptions.NewErrorMetadataNotFound()
	}

	values := md.Get("token")
	if len(values) == 0 {
		return "", exceptions.NewErrorTokenAbsenceInMetadata()
	}

	encodedToken := values[0]
	return s.decode(encodedToken)
}

func (s *GophKeeperServer) decode(encoded string) (string, error) {
	decoded, err := hex.DecodeString(encoded)
	if err != nil {
		logger.Logger.Error(err.Error())
		return "", err
	}
	return string(decoded), nil
}

func (s *GophKeeperServer) hash256(src string) (dst string) {
	dst = fmt.Sprintf("%x", sha256.Sum256([]byte(src)))
	return dst
}
