package handlers

import (
	"geo-test/modules/entities"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GeoHandler struct {
	service entities.GeoUseCase
}

func NewGeoHandler(usecase entities.GeoUseCase) *GeoHandler {
	return &GeoHandler{usecase}
}

func (h *GeoHandler) GetLocations(c *fiber.Ctx) error {

	locations, err := h.service.GetAllLocations()
	if err != nil {
		return err
	}

	return c.JSON(locations)
}

func (h *GeoHandler) GetLocation(c *fiber.Ctx) error {
	locationID := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(locationID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid location ID")
	}

	location, err := h.service.GetLocation(objID)
	if err != nil {
		return err
	}
	return c.JSON(location)
}

func (h *GeoHandler) PostLocation(c *fiber.Ctx) error {
	var location entities.Location
	if err := c.BodyParser(&location); err != nil {
		return err
	}

	createdLocation, err := h.service.CreateLocation(&location)
	if err != nil {
		return err
	}

	return c.JSON(createdLocation)
}

func (h *GeoHandler) DeleteLocation(c *fiber.Ctx) error {

	locationID := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(locationID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid location ID")
	}

	if err := h.service.DeleteLocation(objID); err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}
func (h *GeoHandler) UpdateLocation(c *fiber.Ctx) error {

	locationID := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(locationID)
	if err != nil {

		return c.Status(fiber.StatusBadRequest).SendString("Invalid location ID")
	}
	var updateData entities.Location
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}
	updateData.ID = objID

	updatedLocation, err := h.service.UpdateLocation(&updateData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update location")
	}
	return c.JSON(updatedLocation)
}

func (h *GeoHandler) CreateFeature(c *fiber.Ctx) error {
	var feature entities.Feature
	if err := c.BodyParser(&feature); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	createdFeature, err := h.service.CreateFeature(&feature)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(createdFeature)
}

func (h *GeoHandler) DeleteFeature(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID format")
	}

	err = h.service.DeleteFeature(objectID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *GeoHandler) GetFeature(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID format")
	}

	feature, err := h.service.GetFeature(objectID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	return c.JSON(feature)
}

func (h *GeoHandler) GetFeatures(c *fiber.Ctx) error {
	features, err := h.service.GetFeatures()
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	return c.JSON(features)
}

func (h *GeoHandler) UpdateFeature(c *fiber.Ctx) error {

	locationID := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(locationID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid location ID")
	}
	var feature entities.Feature
	if err := c.BodyParser(&feature); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	feature.ID = objID

	updatedFeature, err := h.service.UpdateFeature(&feature)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	return c.JSON(updatedFeature)
}
