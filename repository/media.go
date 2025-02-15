package repository

import (
	"context"

	"github.com/chirag3003/go-backend-template/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type mediaRepository struct {
	media *mongo.Collection
}

type MediaRepository interface {
	CreateMedia(context context.Context, media *models.Media) error
	GetMediaByID(context context.Context, id string) (*models.Media, error)
	GetMediaByKey(context context.Context, key string) (*models.Media, error)
	UpdateMedia(context context.Context, media *models.Media) error
}

func NewMediaRepository() MediaRepository {
	return &mediaRepository{
		media: conn.DB().Collection("media"),
	}
}

func (r *mediaRepository) CreateMedia(context context.Context, media *models.Media) error {
	_, err := r.media.InsertOne(context, media)
	if err != nil {
		return err
	}
	return nil
}

func (r *mediaRepository) GetMediaByID(context context.Context, id string) (*models.Media, error) {
	var media models.Media
	ID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	err = r.media.FindOne(context, bson.M{"_id": ID}).Decode(&media)
	if err != nil {
		return nil, err
	}
	return &media, nil
}

func (r *mediaRepository) GetMediaByKey(context context.Context, key string) (*models.Media, error) {
	var media models.Media
	err := r.media.FindOne(context, bson.M{"key": key}).Decode(&media)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &media, nil
}

func (r *mediaRepository) UpdateMedia(context context.Context, media *models.Media) error {
	_, err := r.media.UpdateOne(context, bson.M{"_id": media.ID}, bson.M{"$set": media})
	if err != nil {
		return err
	}
	return nil
}


