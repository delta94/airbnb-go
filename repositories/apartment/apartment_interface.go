package apartment

import (
	"context"

	entities "github.com/CezarGarrido/airbnb-go/entities"
)

// ApartmentRepo :
type ApartmentRepo interface {
	Create(ctx context.Context, apartment *entities.Apartment) (int64, error)
	Update(ctx context.Context, apartment *entities.Apartment) (*entities.Apartment, error)
	Delete(ctx context.Context, apartment *entities.Apartment) error
	FindAll(ctx context.Context) ([]*entities.Apartment, error)
}