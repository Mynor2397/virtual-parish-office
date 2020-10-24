package service

import (
	"context"
	"github.com/google/uuid"

	"github.com/Mynor2397/virtual-parish-office/internal/models"
)

func (*personService) GetPlaces(ctx context.Context) ([]models.Place, error) {
	return myPersonStorage.GetPlaces(ctx)
}

func (*personService) CreatePlace(ctx context.Context, place models.Place) (int, error) {
	place.ID = uuid.New().String()
	return myPersonStorage.CreatePlace(ctx, place)
}

func (*personService) GetCountPlace(ctx context.Context) (int, error) {
	return myPersonStorage.GetLastPlace(ctx)
}
