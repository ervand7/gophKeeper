package exceptions

import (
	"errors"
	"fmt"
)

type ErrorLoginAlreadyExists struct {
	Err error
}

func NewErrorLoginAlreadyExists(login string) error {
	return &ErrorLoginAlreadyExists{
		Err: fmt.Errorf(`%w`, errors.New(login)),
	}
}

func (e ErrorLoginAlreadyExists) Unwrap() error {
	return e.Err
}

func (e ErrorLoginAlreadyExists) Error() string {
	return fmt.Sprintf("%s already exists", e.Err)
}

type ErrorEntryNotExists struct {
	Err error
}

func NewErrorEntryNotExists() error {
	return &ErrorEntryNotExists{}
}

func (e ErrorEntryNotExists) Error() string {
	return "entry with this title does not exist"
}
