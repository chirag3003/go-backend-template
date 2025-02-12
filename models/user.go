package models

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID      bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string        `json:"name" bson:"name"`
	Email   string        `json:"email" bson:"email"`
	PhoneNo string        `json:"phoneNo" bson:"phoneNo"`
	Hash    string        `json:"hash" bson:"hash"`
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

func (user *User) GenerateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      user.ID,
		"name":    user.Name,
		"email":   user.Email,
		"phoneNo": user.PhoneNo,
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (user *User) ParseJWT(tokenString string) (*User, error) {
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
	id, err := bson.ObjectIDFromHex(claims["id"].(string))
  fmt.Println(id, "IDDD")
	if err != nil {
		return nil, err
	}
	user = &User{
		ID:      id,
		Name:    claims["name"].(string),
		Email:   claims["email"].(string),
		PhoneNo: claims["phoneNo"].(string),
	}
	return user, nil
}

type LoginUserBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
