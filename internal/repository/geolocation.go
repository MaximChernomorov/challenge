package repository

import (
	"context"
	"database/sql"

	"github.com/MaximChernomorov/challenge-test/internal/repository/model"
	"github.com/friendsofgo/errors"
	"github.com/lib/pq"
)

// Geolocation to hide implementation we should operate with this struct outside of repository instead of
// underlying model
type Geolocation struct {
	model.Geolocation
}

// GeolocationSlice to hide implementation we should operate with this struct outside of repository instead
// of underlying model
type GeolocationSlice struct {
	model.GeolocationSlice
}

func (s *GeolocationSlice) GetLength() int {
	return len(s.GeolocationSlice)
}

func (s *GeolocationSlice) insertAll(ctx context.Context, conn *sql.DB) error {
	if len(s.GeolocationSlice) == 0 {
		return nil
	}
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "failed to start transaction")
	}
	statement, err := tx.Prepare(pq.CopyIn(
		model.TableNames.Geolocations,
		model.GeolocationColumns.City,
		model.GeolocationColumns.Country,
		model.GeolocationColumns.CountryCode,
		model.GeolocationColumns.IPAddress,
		model.GeolocationColumns.Coordinates,
		model.GeolocationColumns.MysteryValue))
	if err != nil {
		return errors.Wrap(err, "failed to prepare transaction")
	}

	for _, geolocation := range s.GeolocationSlice {
		_, err = statement.Exec(
			geolocation.City,
			geolocation.Country,
			geolocation.CountryCode,
			geolocation.IPAddress,
			geolocation.Coordinates,
			geolocation.MysteryValue)
		if err != nil {
			return errors.Wrap(err, "failed to execute statement")
		}
	}

	_, err = statement.Exec()
	if err != nil {
		return errors.Wrap(err, "failed to execute statement")
	}

	err = statement.Close()
	if err != nil {
		return errors.Wrap(err, "failed to close statement")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}

	return nil
}
