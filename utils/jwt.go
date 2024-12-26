package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// IMPORTANT: This should be in a .env file,
// i will not do this since in this project is not important
var Key = []byte("jlsdjfalskdjflskdjfkey")

func JwtNew(id int64, email string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"sub": email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	s, err := t.SignedString(Key)
	if err != nil {
		return "", err
	}
	return s, nil
}

type claims struct {
	Id    int64
	Email string
}

func JwtParse(token string) (*claims, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method ")
		}
		return Key, nil
	})

	if err != nil {
		return nil, errors.New("could not parse token")
	}

	if t.Valid {
		return &claims{
			Id:    int64(t.Claims.(jwt.MapClaims)["id"].(float64)),
			Email: t.Claims.(jwt.MapClaims)["sub"].(string),
		}, nil
	}
	return nil, err
}
