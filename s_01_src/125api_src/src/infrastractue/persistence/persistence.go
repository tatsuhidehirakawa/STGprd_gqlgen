package persistence

import (
	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/domain/repository"
	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/io"
)

type Repositories struct {
	db   *io.SQLdatabase
	OfferList repository.IOfferListRepository
}

func NewRepositories(db *io.SQLdatabase) (*Repositories, error) {
	return &Repositories{
		db:   db,
		OfferList: NewOfferListRepository(db),
	}, nil
}
