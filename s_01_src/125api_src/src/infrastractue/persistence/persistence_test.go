package persistence

import (
	"context"
	"os"
	"testing"

	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/config"
	"github.com/gqlgensamples/golang-gqlgen-postgresql-example/io"
)

var (
	offerlistRepo *OfferListRepository
)

func TestMain(m *testing.M) {
	cfg, _ := config.LoadConfig(context.Background())
	db, _ := io.NewSQLdatabase(cfg)
	offerlistRepo = NewOfferListRepository(db)

	res := m.Run()

	os.Exit(res)

}
