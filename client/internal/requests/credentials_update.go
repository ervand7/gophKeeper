package requests

import (
	"context"

	"google.golang.org/grpc/metadata"

	pb "gophkeeper/proto"
)

func (c *Client) UpdateCredentials(in *pb.UpdateCredentialsRequest) (resultMessage string) {
	if c.token == "" {
		return "not authorized"
	}

	md := metadata.New(map[string]string{"token": c.token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	in.Credentials.Login = c.encrypter.Encrypt(in.Credentials.Login)
	in.Credentials.Password = c.encrypter.Encrypt(in.Credentials.Password)

	_, err := c.runner.UpdateCredentials(ctx, in)
	if err != nil {
		return c.getGRPCError(err.Error())
	}
	return "updated successfully"
}
