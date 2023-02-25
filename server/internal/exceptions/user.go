package exceptions

type ErrorUserNotFound struct {
	Err error
}

func NewErrorUserNotFound() error {
	return &ErrorUserNotFound{}
}

func (e ErrorUserNotFound) Error() string {
	return "user not found"
}
