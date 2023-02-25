package exceptions

type ErrorAlreadyExists struct {
	Err error
}

func NewErrorAlreadyExists() error {
	return &ErrorAlreadyExists{}
}

func (e ErrorAlreadyExists) Error() string {
	return "entry with this title already exists"
}
