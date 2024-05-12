package usecases

import (
	"errors"
	"fmt"
	"geo-test/modules/entities"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (g *geoUsecase) CreateFeature(feature *entities.Feature) (*entities.Feature, error) {
	locationRes, err := g.repo.GetLocationOne(feature.Properties.CollectionID)

	if err != nil {
		return nil, fmt.Errorf("no CollectionID: %v found", feature.Properties.CollectionID.Hex())
	}

	// set feature name
	feature.Properties.Name = locationRes.Name

	createdFeature, err := g.repo.CreateFeatureOne(feature)
	if err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("collectionID not found in collection")
		}
		return nil, errors.New("something went worng")
	}
	return createdFeature, nil
}

func (g *geoUsecase) DeleteFeature(id primitive.ObjectID) error {
	err := g.repo.DeleteFeature(id)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("can not delete feature id: %v", id)
	}
	return nil
}

func (g *geoUsecase) GetFeature(id primitive.ObjectID) (*entities.Feature, error) {
	feature, err := g.repo.GetFeatureOne(id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("not found")
	}
	return feature, nil
}

func (g *geoUsecase) GetFeatures() ([]entities.Feature, error) {
	features, err := g.repo.GetFeatureMany()
	if err != nil {
		log.Println(err)
		return nil, errors.New("not found")
	}
	return features, nil
}

func (g *geoUsecase) UpdateFeature(feature *entities.Feature) (*entities.Feature, error) {

	_, err := g.repo.GetLocationOne(feature.Properties.CollectionID)

	if err != nil {
		return nil, fmt.Errorf("no CollectionID: %v found", feature.Properties.CollectionID.Hex())
	}

	updatedFeature, err := g.repo.UpdateFeature(feature)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("can not update feature id: %v", feature.ID.Hex())
	}
	return updatedFeature, nil
}
