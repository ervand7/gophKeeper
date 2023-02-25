package exceptions

type ErrorMetadataNotFound struct {
	Err error
}

func NewErrorMetadataNotFound() error {
	return &ErrorMetadataNotFound{}
}

func (e ErrorMetadataNotFound) Error() string {
	return "metadata not found"
}

type ErrorTokenAbsenceInMetadata struct {
	Err error
}

func NewErrorTokenAbsenceInMetadata() error {
	return &ErrorTokenAbsenceInMetadata{}
}

func (e ErrorTokenAbsenceInMetadata) Error() string {
	return "expected absence in metadata"
}
