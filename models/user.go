package models

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID      string `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string `json:"name" bson:"name"`
	Email   string `json:"email" bson:"email"`
	PhoneNo string `json:"phoneNo" bson:"phoneNo"`
	Hash    string `json:"hash" bson:"hash"`
}

func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Hash = string(hash)
	return nil
}

func (u *User) VerifyPassword(password string) bool {
	// return HashPassword(password) == u.Hash
	return bcrypt.CompareHashAndPassword([]byte(u.Hash), []byte(password)) == nil
}

func GenerateJWT(user *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      user.ID,
		"name":    user.Name,
		"email":   user.Email,
		"phoneNo": user.PhoneNo,
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ParseJWT(tokenString string) (*User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}
	user := &User{
		ID:      claims["id"].(string),
		Name:    claims["name"].(string),
		Email:   claims["email"].(string),
		PhoneNo: claims["phoneNo"].(string),
	}
	return user, nil
}
