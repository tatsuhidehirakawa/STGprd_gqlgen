// O/Rマッパー
package persistence

import (
	"context"
	"database/sql"

	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/domain/entity"
	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/domain/repository"
	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/io"
)

type OfferListRepository struct {
	database *io.SQLdatabase
}

var _ repository.IOfferListRepository = (*OfferListRepository)(nil)

func NewOfferListRepository(db *io.SQLdatabase) *OfferListRepository {
	return &OfferListRepository{
		database: db,
	}
}

func (r *OfferListRepository) GetOfferList(ctx context.Context, offerlistID string) (*entity.OfferList, error) {
	var offerlist entity.OfferList
	err := r.database.SQLX.Get(&offerlist, "SELECT id, name FROM offerlists WHERE id=$1", offerlistID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, nil
		default:
			return nil, err
		}
	}
	return &offerlist, nil
}
