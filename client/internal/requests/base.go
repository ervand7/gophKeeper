package requests

import (
	"log"
	"strings"

	"google.golang.org/grpc"

	"gophkeeper/client/internal/config"
	c "gophkeeper/pkg/cert"
	"gophkeeper/pkg/encryption"
	pb "gophkeeper/proto"
)

type Client struct {
	runner     pb.GophKeeperClient
	userData   *pb.UserData
	token      string
	stopStream chan bool
	encrypter  *encryption.Encrypter
}

func NewClient() *Client {
	conn, err := grpc.Dial(
		config.GetServerAddress(),
		grpc.WithTransportCredentials(c.LoadClientCertificate()),
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	runner := pb.NewGophKeeperClient(conn)
	return &Client{
		runner:     runner,
		userData:   &pb.UserData{},
		stopStream: make(chan bool),
		encrypter:  encryption.NewEncrypter(),
	}
}

func (c *Client) getGRPCError(raw string) string {
	separator := "desc = "
	return strings.Split(raw, separator)[1]
}

func (c *Client) reset() {
	c.userData = &pb.UserData{}
	c.token = ""
	c.stopStream <- true
}
