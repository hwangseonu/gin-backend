package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/hwangseonu/gin-backend/models"
	"github.com/satori/go.uuid"
	"time"
)

type CustomClaims struct {
	jwt.StandardClaims
	Identity string `json:"identity"`
}

func (c *CustomClaims) Valid() error {
	if err := c.StandardClaims.Valid(); err != nil {
		return err
	}
	u, err := models.GetUser(c.Identity)

	if err != nil {
		return err
	}
	if u == nil {
		return errors.New("could not find user")
	}
	return nil
}

func GenerateToken(t, id string) (string, error) {
	claims := jwt.StandardClaims{
		Audience:  "",
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
		Id:        uuid.Must(uuid.NewV4(), nil).String(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "",
		Subject: t,
		NotBefore: time.Now().Unix(),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS512, CustomClaims{claims, id}).SignedString([]byte("secret"))
}