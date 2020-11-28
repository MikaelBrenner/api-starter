package entities

import (
	"api-starter/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/goonode/mogo"
)

type UserInfo struct {
	KeyA string `json:"key-a"`
	KeyB string `json:"key-b"`
	KeyC string `json:"key-c"`
}

type User struct {
	mogo.DocumentModel `bson:",inline" coll:"users"`
	Email              string   `idx:"{email},unique" json:"email" binding:"required"`
	Password           string   `json:"password" binding:"required"`
	Name               string   `json:"name"`
	Info               UserInfo `json:"user-info"`
}

func (user *User) GetJwtToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
	})

	secretKey := utils.EnvVar("TOKEN_KEY")
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}

func init() {
	mogo.ModelRegistry.Register(User{})
}
