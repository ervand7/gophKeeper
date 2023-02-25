package requests

import (
	"context"
	"fmt"

	pb "gophkeeper/proto"
)

func (c *Client) Auth(in *pb.AuthRequest) (resultMessage string) {
	if c.token != "" {
		return "already authorized"
	}
	resp, err := c.runner.Auth(context.Background(), in)
	if err != nil {
		return c.getGRPCError(err.Error())
	}

	c.token = resp.Token
	go c.StreamDB()
	return fmt.Sprintf("success auth for %s", in.Login)
}
