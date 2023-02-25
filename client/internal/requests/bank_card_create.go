package requests

import (
	"context"

	"google.golang.org/grpc/metadata"

	pb "gophkeeper/proto"
)

func (c *Client) CreateBankCard(in *pb.CreateBankCardRequest) (resultMessage string) {
	if c.token == "" {
		return "not authorized"
	}

	md := metadata.New(map[string]string{"token": c.token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	in.BankCard.CardHolder = c.encrypter.Encrypt(in.BankCard.CardHolder)
	in.BankCard.CardNumber = c.encrypter.Encrypt(in.BankCard.CardNumber)
	in.BankCard.CardExpire = c.encrypter.Encrypt(in.BankCard.CardExpire)
	in.BankCard.CardCvv = c.encrypter.Encrypt(in.BankCard.CardCvv)

	_, err := c.runner.CreateBankCard(ctx, in)
	if err != nil {
		return c.getGRPCError(err.Error())
	}
	return "created successfully"
}
