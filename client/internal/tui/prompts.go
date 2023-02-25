package tui

import "github.com/c-bata/go-prompt"

// ShowPrompts shows prompts to user.
func (e *Executor) ShowPrompts(d prompt.Document) []prompt.Suggest {
	var completions []prompt.Suggest

	if d.FindStartOfPreviousWord() == 0 {
		completions = []prompt.Suggest{
			{Text: "register", Description: "Example: register;login;password"},
			{Text: "auth", Description: "Example: auth;login;password"},
			{Text: "logout", Description: "Example: logout"},
			{Text: "delete-user", Description: "Example: delete-user"},

			{Text: "create-bank-card", Description: "Example: create-bank-card;title;card_holder;card_number;card_expire;card_cvv;meta_info"},
			{Text: "create-binary-data", Description: "Example: create-binary-data;title;content;meta_info"},
			{Text: "create-credentials", Description: "Example: create-credentials;title;login;password;meta_info"},
			{Text: "create-text", Description: "Example: create-text;title;content;meta_info"},

			{Text: "get-bank-card", Description: "Example: get-bank-card"},
			{Text: "get-binary-data", Description: "Example: get-binary-data"},
			{Text: "get-credentials", Description: "Example: get-credentials"},
			{Text: "get-text", Description: "Example: get-text"},

			{Text: "update-bank-card", Description: "Example: update-bank-card;title;card_holder;card_number;card_expire;card_cvv;meta_info"},
			{Text: "update-binary-data", Description: "Example: update-binary-data;title;content;meta_info"},
			{Text: "update-credentials", Description: "Example: update-credentials;title;login;password;meta_info"},
			{Text: "update-text", Description: "Example: update-text;title;content;meta_info"},

			{Text: "delete-bank-card", Description: "Example: delete-bank-card;title"},
			{Text: "delete-binary-data", Description: "Example: delete-binary-data;title"},
			{Text: "delete-credentials", Description: "Example: delete-credentials;title"},
			{Text: "delete-text", Description: "Example: delete-text;title"},

			{Text: "exit", Description: "Example: exit"},
		}
	}

	return prompt.FilterHasPrefix(completions, d.GetWordBeforeCursor(), false)
}
