package requests

import (
	"errors"

	pb "gophkeeper/proto"
)

func (c *Client) GetBankCard() ([]*pb.BankCard, error) {
	if c.token == "" {
		return nil, errors.New("not authorized")
	}

	var result []*pb.BankCard
	for _, val := range c.userData.BankCard {
		cardHolder, err := c.encrypter.Decrypt(val.CardHolder)
		if err != nil {
			return nil, err
		}
		cardNumber, err := c.encrypter.Decrypt(val.CardNumber)
		if err != nil {
			return nil, err
		}
		cardExpire, err := c.encrypter.Decrypt(val.CardExpire)
		if err != nil {
			return nil, err
		}
		cardCvv, err := c.encrypter.Decrypt(val.CardCvv)
		if err != nil {
			return nil, err
		}

		element := &pb.BankCard{
			Title:      val.Title,
			CardHolder: cardHolder,
			CardNumber: cardNumber,
			CardExpire: cardExpire,
			CardCvv:    cardCvv,
			Meta:       val.Meta,
			CreatedAt:  val.CreatedAt,
			UpdatedAt:  val.UpdatedAt,
		}
		result = append(result, element)
	}

	return result, nil
}
