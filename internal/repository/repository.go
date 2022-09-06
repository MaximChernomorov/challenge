package repository

import "context"

// Repository hides particular db implementation from client code
type Repository interface {
	// AddGeolocationSlice stores geolocation slice into db
	AddGeolocationSlice(ctx context.Context, geolocationSlice GeolocationSlice) error
	// LocateIP finds IP address in db and returns geolocation
	LocateIP(ctx context.Context, IP string) (Geolocation, error)

	// Close db
	Close() error
}
