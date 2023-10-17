package cache

import (
	"context"
	"net/http"

	"github.com/rl404/fairy/cache"
	"github.com/rl404/fairy/errors/stack"
	"github.com/rl404/mal-cover/internal/domain/mal/entity"
	"github.com/rl404/mal-cover/internal/domain/mal/repository"
	"github.com/rl404/mal-cover/internal/errors"
	"github.com/rl404/mal-cover/internal/utils"
)

// Cache contains functions for mal cache.
type Cache struct {
	cacher cache.Cacher
	repo   repository.Repository
}

// New to create new mal cache.
func New(cacher cache.Cacher, repo repository.Repository) *Cache {
	return &Cache{
		cacher: cacher,
		repo:   repo,
	}
}

// GetList to get anime/manga list from cache.
func (c *Cache) GetList(ctx context.Context, username, mainType string) (data []entity.Entry, code int, err error) {
	key := utils.GetKey("list", username, mainType)
	if c.cacher.Get(ctx, key, &data) == nil {
		return data, http.StatusOK, nil
	}

	data, code, err = c.repo.GetList(ctx, username, mainType)
	if err != nil {
		return nil, code, stack.Wrap(ctx, err)
	}

	if err := c.cacher.Set(ctx, key, data); err != nil {
		return nil, http.StatusInternalServerError, stack.Wrap(ctx, err, errors.ErrInternalCache)
	}

	return data, code, nil
}
