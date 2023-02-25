package tui

import (
	"strings"

	"gophkeeper/client/internal/requests"
)

// Executor for implementing Terminal User Interface.
type Executor struct {
	Client *requests.Client
}

func NewExecutor() *Executor {
	return &Executor{
		Client: requests.NewClient(),
	}
}

func (e *Executor) parseRawData(raw string) []string {
	raw = strings.TrimSpace(raw)
	return strings.Split(raw, ";")
}
