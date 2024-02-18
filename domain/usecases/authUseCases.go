package usecases

import (
	"backend-go/mod/domain/entities"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthUseCases struct {
}

func NewAuthUseCases() AuthUseCases {
	return AuthUseCases{}
}

func (usecases AuthUseCases) GenerateJwt(user string, pass string) (*string, error) {

	if user != os.Getenv("EXTERNAL_USER") || pass != os.Getenv("EXTERNAL_PASS") {
		err := entities.AuthError{
			Message: "Usuario inválido",
		}
		return nil, &err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["iat"] = time.Now().Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_AUTH_SECRET")))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func (usecases AuthUseCases) ValidateJwt(token string) (bool, error) {
	_, err := jwt.Parse(token, func(tokenJwt *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_AUTH_SECRET")), nil
	})
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (usecases AuthUseCases) ValidateBasic(user string, pass string) (bool, error) {
	return (user == os.Getenv("BASIC_AUTH_USER") && pass == os.Getenv("BASIC_AUTH_PASS")), nil
}
