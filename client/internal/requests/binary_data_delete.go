package requests

import (
	"context"

	"google.golang.org/grpc/metadata"

	pb "gophkeeper/proto"
)

func (c *Client) DeleteBinaryData(in *pb.DeleteBinaryDataRequest) (resultMessage string) {
	if c.token == "" {
		return "not authorized"
	}

	md := metadata.New(map[string]string{"token": c.token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	_, err := c.runner.DeleteBinaryData(ctx, in)
	if err != nil {
		return c.getGRPCError(err.Error())
	}
	return "deleted successfully"
}
