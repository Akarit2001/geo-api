package main

import (
	"context"
	"geo-test/modules/configs"
	"geo-test/modules/handlers"
	"geo-test/modules/repositories"
	"geo-test/modules/usecases"
	"geo-test/pkg/databases"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := new(configs.Config)
	configs.LoadConfigs(cfg)

	ctx := context.TODO()
	database, err := databases.NewMongoConnection(cfg, ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Client().Disconnect(ctx)
	repo := repositories.NewGeoRepository(database)
	usecase := usecases.NewGeoUsecase(repo)
	handler := handlers.NewGeoHandler(usecase)

	// collection
	app := fiber.New()
	api := app.Group("/api")
	api.Post("/locations", handler.PostLocation)
	api.Get("/locations", handler.GetLocations)
	api.Get("/locations/:id", handler.GetLocation)
	api.Delete("/locations/:id", handler.DeleteLocation)
	api.Put("/locations/:id", handler.UpdateLocation)

	api.Post("/features", handler.CreateFeature)
	api.Get("/features", handler.GetFeatures)
	api.Get("/features/:id", handler.GetFeature)
	api.Delete("/features/:id", handler.DeleteFeature)
	api.Put("/features/:id", handler.UpdateFeature)

	app.Listen(":8080")
}
