package usecases

import (
	"geo-test/modules/entities"
)

type geoUsecase struct {
	repo entities.GeoRepository
}

func NewGeoUsecase(repo entities.GeoRepository) entities.GeoUseCase {
	return &geoUsecase{repo: repo}
}
