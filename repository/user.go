package repository

import (
	"context"

	"github.com/chirag3003/go-backend-template/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type userRepository struct {
	user *mongo.Collection
}

type UserRepository interface {
	CreateUser(context context.Context, user *models.User) error
	GetUserByID(context context.Context, id string) (*models.User, error)
	GetUserByEmail(context context.Context, email string) (*models.User, error)
	UpdateUser(context context.Context, user *models.User) error
}

func NewUserRepository() UserRepository {
	return &userRepository{
		user: conn.DB().Collection("users"),
	}
}

func (r *userRepository) CreateUser(context context.Context, user *models.User) error {
	_, err := r.user.InsertOne(context, user)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetUserByID(context context.Context, id string) (*models.User, error) {
	var user models.User
	ID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	err = r.user.FindOne(context, bson.M{"_id": ID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByEmail(context context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.user.FindOne(context, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(context context.Context, user *models.User) error {
	_, err := r.user.UpdateOne(context, bson.M{"_id": user.ID}, bson.M{"$set": user})
	if err != nil {
		return err
	}
	return nil
}
