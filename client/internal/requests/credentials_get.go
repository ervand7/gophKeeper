package requests

import (
	"errors"

	pb "gophkeeper/proto"
)

func (c *Client) GetCredentials() ([]*pb.Credentials, error) {
	if c.token == "" {
		return nil, errors.New("not authorized")
	}

	var result []*pb.Credentials
	for _, val := range c.userData.Credentials {
		login, err := c.encrypter.Decrypt(val.Login)
		if err != nil {
			return nil, err
		}
		password, err := c.encrypter.Decrypt(val.Password)
		if err != nil {
			return nil, err
		}

		element := &pb.Credentials{
			Title:     val.Title,
			Login:     login,
			Password:  password,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
			Meta:      val.Meta,
		}
		result = append(result, element)
	}

	return result, nil
}
