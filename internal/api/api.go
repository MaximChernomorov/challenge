package api

import (
	"github.com/MaximChernomorov/challenge-test/internal/repository"
)

type API struct {
	repo repository.Repository
}

type ErrorResponse struct {
	Error string `json:"error,omitempty"`
}

func (api *API) SetRepo(repo repository.Repository) {
	api.repo = repo
}
