package apperrors

type myAppError struct {
	ErrCode
	Message string
	Err     error
}

func (myErr *myAppError) Error() string {
	return myErr.Err.Error()
}

func (myErr *myAppError) Unwrap() error {
	return myErr.Err
}
