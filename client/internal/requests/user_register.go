package requests

import (
	"context"
	"fmt"

	pb "gophkeeper/proto"
)

func (c *Client) Register(in *pb.RegisterRequest) (resultMessage string) {
	resp, err := c.runner.Register(context.Background(), in)
	if err != nil {
		return c.getGRPCError(err.Error())
	}

	c.token = resp.Token
	go c.StreamDB()
	return fmt.Sprintf("success register for %s", in.Login)
}
