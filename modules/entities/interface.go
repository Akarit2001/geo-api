package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type GeoUseCase interface {
	CreateLocation(location *Location) (*Location, error)
	GetAllLocations() ([]Location, error)
	GetLocation(id primitive.ObjectID) (*Location, error)
	DeleteLocation(id primitive.ObjectID) error
	UpdateLocation(*Location) (*Location, error)

	GetFeature(id primitive.ObjectID) (*Feature, error)
	GetFeatures() ([]Feature, error)
	CreateFeature(*Feature) (*Feature, error)
	UpdateFeature(*Feature) (*Feature, error)
	DeleteFeature(id primitive.ObjectID) error
}

type GeoRepository interface {
	GetLocationOne(id primitive.ObjectID) (*Location, error)
	GetLocationMany() ([]Location, error)
	CreateLocationOne(*Location) (*Location, error)
	UpdateLocation(*Location) (*Location, error)
	DeleteLocation(id primitive.ObjectID) error

	GetFeatureOne(id primitive.ObjectID) (*Feature, error)
	GetFeatureMany() ([]Feature, error)
	CreateFeatureOne(*Feature) (*Feature, error)
	UpdateFeature(*Feature) (*Feature, error)
	DeleteFeature(id primitive.ObjectID) error
}
