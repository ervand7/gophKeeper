package requests

import (
	"errors"

	pb "gophkeeper/proto"
)

func (c *Client) GetBinaryData() ([]*pb.BinaryData, error) {
	if c.token == "" {
		return nil, errors.New("not authorized")
	}

	var result []*pb.BinaryData
	for _, val := range c.userData.BinaryData {
		element := &pb.BinaryData{
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
