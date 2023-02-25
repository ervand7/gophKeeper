package tui

import (
	"fmt"
	"os"

	pb "gophkeeper/proto"
)

const incorrectInput = "incorrect input data, see prompt"

// HandleCommands receives commands from keyboard and passes them to client.
func (e *Executor) HandleCommands(raw string) {
	data := e.parseRawData(raw)
	command := data[0]
	args := data[1:]

	switch command {
	case "register":
		if len(args) != 2 {
			fmt.Println(incorrectInput)
			return
		}
		result := e.Client.Register(
			&pb.RegisterRequest{Login: args[0], Password: args[1]},
		)
		fmt.Println(result)
		return

	case "auth":
		if len(args) != 2 {
			fmt.Println(incorrectInput)
			return
		}
		result := e.Client.Auth(
			&pb.AuthRequest{Login: args[0], Password: args[1]},
		)
		fmt.Println(result)
		return

	case "logout":
		result := e.Client.Logout()
		fmt.Println(result)
		return

	case "delete-user":
		result := e.Client.DeleteUser()
		fmt.Println(result)
		return

	case "create-bank-card":
		if len(args) < 5 {
			fmt.Println(incorrectInput)
			return
		}
		meta := ""
		if len(args) == 6 {
			meta = args[5]
		}
		result := e.Client.CreateBankCard(
			&pb.CreateBankCardRequest{
				BankCard: &pb.BankCard{
					Title:      args[0],
					CardHolder: args[1],
					CardNumber: args[2],
					CardExpire: args[3],
					CardCvv:    args[4],
					Meta:       meta,
				},
			},
		)
		fmt.Println(result)
		return

	case "create-binary-data":
		if len(args) < 2 {
			fmt.Println(incorrectInput)
			return
		}
		meta := ""
		if len(args) == 3 {
			meta = args[2]
		}
		result := e.Client.CreateBinaryData(
			&pb.CreateBinaryDataRequest{
				BinaryData: &pb.BinaryData{
					Title:   args[0],
					Content: []byte(args[1]),
					Meta:    meta,
				},
			},
		)
		fmt.Println(result)
		return

	case "create-credentials":
		if len(args) < 3 {
			fmt.Println(incorrectInput)
			return
		}
		meta := ""
		if len(args) == 4 {
			meta = args[3]
		}
		result := e.Client.CreateCredentials(
			&pb.CreateCredentialsRequest{
				Credentials: &pb.Credentials{
					Title:    args[0],
					Login:    args[1],
					Password: args[2],
					Meta:     meta,
				},
			},
		)
		fmt.Println(result)
		return

	case "create-text":
		if len(args) < 2 {
			fmt.Println(incorrectInput)
			return
		}
		meta := ""
		if len(args) == 3 {
			meta = args[2]
		}
		result := e.Client.CreateText(
			&pb.CreateTextRequest{
				Text: &pb.Text{
					Title:   args[0],
					Content: args[1],
					Meta:    meta,
				},
			},
		)
		fmt.Println(result)
		return

	case "get-bank-card":
		result, err := e.Client.GetBankCard()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, val := range result {
			fmt.Println(val)
		}
		return

	case "get-binary-data":
		result, err := e.Client.GetBinaryData()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, val := range result {
			fmt.Println(val)
		}
		return

	case "get-credentials":
		result, err := e.Client.GetCredentials()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, val := range result {
			fmt.Println(val)
		}
		return

	case "get-text":
		result, err := e.Client.GetText()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, val := range result {
			fmt.Println(val)
		}
		return

	case "update-bank-card":
		if len(args) < 5 {
			fmt.Println(incorrectInput)
			return
		}
		meta := ""
		if len(args) == 6 {
			meta = args[5]
		}
		result := e.Client.UpdateBankCard(
			&pb.UpdateBankCardRequest{
				BankCard: &pb.BankCard{
					Title:      args[0],
					CardHolder: args[1],
					CardNumber: args[2],
					CardExpire: args[3],
					CardCvv:    args[4],
					Meta:       meta,
				},
			},
		)
		fmt.Println(result)
		return

	case "update-binary-data":
		if len(args) < 2 {
			fmt.Println(incorrectInput)
			return
		}
		meta := ""
		if len(args) == 3 {
			meta = args[2]
		}
		result := e.Client.UpdateBinaryData(
			&pb.UpdateBinaryDataRequest{
				BinaryData: &pb.BinaryData{
					Title:   args[0],
					Content: []byte(args[1]),
					Meta:    meta,
				},
			},
		)
		fmt.Println(result)
		return

	case "update-credentials":
		if len(args) < 3 {
			fmt.Println(incorrectInput)
			return
		}
		meta := ""
		if len(args) == 4 {
			meta = args[3]
		}
		result := e.Client.UpdateCredentials(
			&pb.UpdateCredentialsRequest{
				Credentials: &pb.Credentials{
					Title:    args[0],
					Login:    args[1],
					Password: args[2],
					Meta:     meta,
				},
			},
		)
		fmt.Println(result)
		return

	case "update-text":
		if len(args) < 2 {
			fmt.Println(incorrectInput)
			return
		}
		meta := ""
		if len(args) == 3 {
			meta = args[2]
		}
		result := e.Client.UpdateText(
			&pb.UpdateTextRequest{
				Text: &pb.Text{
					Title:   args[0],
					Content: args[1],
					Meta:    meta,
				},
			},
		)
		fmt.Println(result)
		return

	case "delete-bank-card":
		if len(args) == 0 {
			fmt.Println(incorrectInput)
			return
		}
		result := e.Client.DeleteBankCard(
			&pb.DeleteBankCardRequest{Title: args[0]},
		)
		fmt.Println(result)
		return

	case "delete-binary-data":
		if len(args) == 0 {
			fmt.Println(incorrectInput)
			return
		}
		result := e.Client.DeleteBinaryData(
			&pb.DeleteBinaryDataRequest{Title: args[0]},
		)
		fmt.Println(result)
		return

	case "delete-credentials":
		if len(args) == 0 {
			fmt.Println(incorrectInput)
			return
		}
		result := e.Client.DeleteCredentials(
			&pb.DeleteCredentialsRequest{Title: args[0]},
		)
		fmt.Println(result)
		return

	case "delete-text":
		if len(args) == 0 {
			fmt.Println(incorrectInput)
			return
		}
		result := e.Client.DeleteText(
			&pb.DeleteTextRequest{Title: args[0]},
		)
		fmt.Println(result)
		return

	case "exit":
		fmt.Println("Good bye")
		os.Exit(1)
		return
	}
}
