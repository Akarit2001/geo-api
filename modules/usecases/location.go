package usecases

import (
	"errors"
	"geo-test/modules/entities"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (g *geoUsecase) CreateLocation(location *entities.Location) (*entities.Location, error) {
	if location.Name == "" {
		return nil, errors.New("name is empty")
	}
	res, err := g.repo.CreateLocationOne(location)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("create point id:", res.ID)
	return res, nil
}

// GetLocations implements entities.GeoUseCase.
func (g *geoUsecase) GetAllLocations() ([]entities.Location, error) {

	res, err := g.repo.GetLocationMany()
	if err != nil {
		log.Println(err)
		return nil, errors.New("some thing went worng")
	}
	return res, nil

}

// GetLocation implements entities.GeoUseCase.
func (g *geoUsecase) GetLocation(id primitive.ObjectID) (*entities.Location, error) {
	res, err := g.repo.GetLocationOne(id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("not found")
	}
	return res, nil
}

// DeleteLocation implements entities.GeoUseCase.
func (g *geoUsecase) DeleteLocation(id primitive.ObjectID) error {
	err := g.repo.DeleteLocation(id)
	if err != nil {
		log.Println(err)
		return errors.New("some thing went worng")
	}
	return nil
}

// UpdateLocation implements entities.GeoUseCase.
func (g *geoUsecase) UpdateLocation(location *entities.Location) (*entities.Location, error) {
	res, err := g.repo.UpdateLocation(location)
	if err != nil {
		log.Println(err)
		return nil, errors.New("some thing went worng")
	}
	return res, nil
}
