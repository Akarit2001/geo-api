package repositories

import (
	"context"
	"errors"
	"geo-test/modules/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *geoRepository) CreateLocationOne(location *entities.Location) (*entities.Location, error) {
	result, err := r.db.Collection("locations").InsertOne(context.TODO(), location)
	if err != nil {
		return nil, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("inserted ID not found")
	}

	location.ID = insertedID

	return location, nil
}

func (r *geoRepository) GetLocationMany() ([]entities.Location, error) {
	filter := bson.D{{}}

	cursor, err := r.db.Collection("locations").Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var locations []entities.Location

	for cursor.Next(context.TODO()) {
		var location entities.Location
		if err := cursor.Decode(&location); err != nil {
			return nil, err
		}
		locations = append(locations, location)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return locations, nil
}

func (r *geoRepository) GetLocationOne(ID primitive.ObjectID) (*entities.Location, error) {
	filter := bson.D{{Key: "_id", Value: ID}}

	var location entities.Location
	err := r.db.Collection("locations").FindOne(context.TODO(), filter).Decode(&location)
	if err != nil {
		return nil, err
	}

	return &location, nil
}

// DeleteLocation implements entities.GeoRepository.
func (r *geoRepository) DeleteLocation(ID primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: ID}}
	_, err := r.db.Collection("locations").DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

// UpdateLocation implements entities.GeoRepository.
func (r *geoRepository) UpdateLocation(location *entities.Location) (*entities.Location, error) {
	filter := bson.D{{Key: "_id", Value: location.ID}}

	update := bson.D{{Key: "$set", Value: location}}

	result, err := r.db.Collection("locations").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, errors.New("no matching document found for update")
	}
	return location, nil
}
