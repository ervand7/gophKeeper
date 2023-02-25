package requests

import (
	"context"

	"google.golang.org/grpc/metadata"

	pb "gophkeeper/proto"
)

func (c *Client) DeleteUser() (resultMessage string) {
	if c.token == "" {
		return "not authorized"
	}

	md := metadata.New(map[string]string{"token": c.token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	in := &pb.DeleteUserRequest{}

	_, err := c.runner.DeleteUser(ctx, in)
	if err != nil {
		return c.getGRPCError(err.Error())
	}

	c.reset()
	return "user was deleted successfully"
}
