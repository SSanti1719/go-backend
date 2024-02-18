package entities

type AuthError struct {
	Message string
}

func (e *AuthError) Error() string {
	return e.Message
}
