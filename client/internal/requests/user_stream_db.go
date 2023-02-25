package requests

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/metadata"

	pb "gophkeeper/proto"
)

// StreamDB invokes stream from server. Stream can be stopped if
// Client.stopStream chan receives value.
func (c *Client) StreamDB() {
	md := metadata.New(map[string]string{"token": c.token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	in := &pb.StreamDBRequest{}

	stream, err := c.runner.StreamDB(ctx, in)
	if err != nil {
		log.Fatal(fmt.Sprintf("openinig stream error: %v", err.Error()))
	}

	for {
		select {
		case <-c.stopStream:
			return
		default:
			resp, err := stream.Recv()
			if err != nil {
				log.Fatalf("receiving stream data error: %v", err.Error())
			} else {
				c.userData = resp.UserData
			}
		}
	}
}
