package repositories

import (
	"context"
	"errors"
	"geo-test/modules/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *geoRepository) CreateFeatureOne(feature *entities.Feature) (*entities.Feature, error) {

	result, err := r.db.Collection("features").InsertOne(context.TODO(), feature)
	if err != nil {
		return nil, err
	}
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("inserted ID not found")
	}
	feature.ID = insertedID
	return feature, nil
}

func (r *geoRepository) DeleteFeature(id primitive.ObjectID) error {

	filter := bson.D{{Key: "_id", Value: id}}

	_, err := r.db.Collection("features").DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (r *geoRepository) GetFeatureMany() ([]entities.Feature, error) {

	cursor, err := r.db.Collection("features").Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var features []entities.Feature

	for cursor.Next(context.TODO()) {
		var feature entities.Feature
		if err := cursor.Decode(&feature); err != nil {
			return nil, err
		}
		features = append(features, feature)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return features, nil
}

func (r *geoRepository) GetFeatureOne(id primitive.ObjectID) (*entities.Feature, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	var feature entities.Feature
	err := r.db.Collection("features").FindOne(context.TODO(), filter).Decode(&feature)
	if err != nil {
		return nil, err
	}

	return &feature, nil
}

// UpdateFeature implements entities.GeoRepository.
func (r *geoRepository) UpdateFeature(feature *entities.Feature) (*entities.Feature, error) {
	filter := bson.D{{Key: "_id", Value: feature.ID}}

	featureBytes, err := bson.Marshal(feature)
	if err != nil {
		return nil, err
	}

	result, err := r.db.Collection("features").ReplaceOne(context.TODO(), filter, featureBytes)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, errors.New("no matching document found for update")
	}

	return feature, nil
}
