package repositories

import (
	"geo-test/modules/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type geoRepository struct {
	db *mongo.Database
}

func NewGeoRepository(db *mongo.Database) entities.GeoRepository {
	return &geoRepository{db}
}
