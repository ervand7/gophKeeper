// Package main is an entry point to the client app.
package main

import (
	"fmt"
	"log"

	"github.com/c-bata/go-prompt"

	"gophkeeper/client/internal/tui"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
)

func logBuildInfo() {
	log.Println(fmt.Sprintf("Build version: %s", buildVersion))
	log.Println(fmt.Sprintf("Build date: %s", buildDate))
}

func main() {
	logBuildInfo()

	executor := tui.NewExecutor()
	serverTUI := prompt.New(
		executor.HandleCommands,
		executor.ShowPrompts,
		prompt.OptionShowCompletionAtStart(),
		prompt.OptionPrefix("> "),
		prompt.OptionMaxSuggestion(21),
		prompt.OptionInputTextColor(prompt.Green),
	)
	serverTUI.Run()
}
