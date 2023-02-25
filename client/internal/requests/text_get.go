package requests

import (
	"errors"

	pb "gophkeeper/proto"
)

func (c *Client) GetText() ([]*pb.Text, error) {
	if c.token == "" {
		return nil, errors.New("not authorized")
	}

	var result []*pb.Text
	for _, val := range c.userData.Text {
		element := &pb.Text{
			Title:     val.Title,
			Content:   val.Content,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
			Meta:      val.Meta,
		}
		result = append(result, element)
	}

	return result, nil
}
