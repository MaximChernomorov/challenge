package repository

import (
	"context"
	"database/sql"

	"github.com/MaximChernomorov/challenge-test/internal/repository/model"
	"github.com/friendsofgo/errors"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type PostgresRepo struct {
	conn *sql.DB
}

func NewPostgresRepo(psqlURL string) (Repository, error) {
	repo := &PostgresRepo{}
	conn, err := sql.Open("postgres", psqlURL)
	if err != nil {
		return repo, errors.Wrap(err, "failed to connect to db")
	}

	boil.SetDB(conn) // global set
	repo.conn = conn

	return repo, nil
}

func (repo *PostgresRepo) AddGeolocationSlice(ctx context.Context, geolocationSlice GeolocationSlice) error {
	return geolocationSlice.insertAll(ctx, repo.conn)
}

func (repo *PostgresRepo) LocateIP(ctx context.Context, IP string) (Geolocation, error) {
	location := Geolocation{}
	locationDB, err := model.Geolocations(qm.Where(model.GeolocationColumns.IPAddress+"=?", IP)).OneG(ctx)
	if err != nil {
		return location, errors.Wrap(err, "failed to get geo location from db")
	}

	return Geolocation{*locationDB}, nil
}

func (repo *PostgresRepo) Close() error {
	if repo.conn != nil {
		return repo.conn.Close()
	}

	return nil
}
