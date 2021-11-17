package service

import (
	"context"

	"github.com/rl404/mal-cover/internal/domain/mal/repository"
)

// Service contains functions for service.
type Service interface {
	GenerateCover(ctx context.Context, data GenerateCoverRequest) (css string, code int, err error)
}

type service struct {
	mal repository.Repository
}

// New to create new service.
func New(mal repository.Repository) Service {
	return &service{
		mal: mal,
	}
}
