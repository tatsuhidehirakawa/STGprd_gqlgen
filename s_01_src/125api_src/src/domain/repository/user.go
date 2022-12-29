package repository

import (
	"context"

	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/domain/entity"
)

type IOfferListRepository interface {
	GetOfferList(ctx context.Context, offerlistID string) (*entity.OfferList, error)
}
