package requests

import (
	"context"

	"google.golang.org/grpc/metadata"

	pb "gophkeeper/proto"
)

func (c *Client) UpdateBinaryData(in *pb.UpdateBinaryDataRequest) (resultMessage string) {
	if c.token == "" {
		return "not authorized"
	}

	md := metadata.New(map[string]string{"token": c.token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	_, err := c.runner.UpdateBinaryData(ctx, in)
	if err != nil {
		return c.getGRPCError(err.Error())
	}
	return "updated successfully"
}
