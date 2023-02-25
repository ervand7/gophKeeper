// Package runner - for running of server.
package runner

import (
	"crypto/tls"
	"fmt"

	"google.golang.org/grpc"

	c "gophkeeper/pkg/cert"
	pb "gophkeeper/proto"
	"gophkeeper/server/internal/config"
	"gophkeeper/server/internal/logger"
	"gophkeeper/server/internal/storage"
	"gophkeeper/server/internal/views"
)

func Run(server *grpc.Server) {
	address := config.GetServerAddress()
	listen, err := tls.Listen("tcp", address, c.LoadServerCertificate())
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}

	pb.RegisterGophKeeperServer(
		server,
		&views.GophKeeperServer{
			Storage: storage.NewStorage(),
		},
	)

	logger.Logger.Info(
		fmt.Sprintf(
			"===== RPC server started on address %s =====", address,
		),
	)
	if err = server.Serve(listen); err != nil {
		logger.Logger.Fatal(err.Error())
	}
}
