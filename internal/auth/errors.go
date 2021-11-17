package auth

type WrongPasswordError struct {
	message string
}

func (e *WrongPasswordError) Error() string {
	return e.message
}

func NewWrongPasswordError(message string) *WrongPasswordError {
	return &WrongPasswordError{
		message: message,
	}
}

type InvalidUserNameError struct {
	message string
}

func (e *InvalidUserNameError) Error() string {
	return e.message
}

func NewInvalidUserNameError(message string) *InvalidUserNameError {
	return &InvalidUserNameError{
		message: message,
	}
}
