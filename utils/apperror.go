package utils

type AppError struct {
	Code    int
	Message string
}

func New(code int, message string) *AppError {
	return &AppError{
		Code: code,
		Message: message,
	}
}

func (e *AppError) Error() string {
	return e.Message
}

