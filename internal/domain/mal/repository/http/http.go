package http

import (
	"context"
	"encoding/json"
	_errors "errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rl404/fairy/errors/stack"
	"github.com/rl404/mal-cover/internal/domain/mal/entity"
	"github.com/rl404/mal-cover/internal/errors"
	"github.com/rl404/mal-cover/internal/utils"
	"github.com/rl404/mal-cover/pkg/log"
)

// Client contains functions for mal http client.
type Client struct {
	client http.Client
	malURL string
}

// New to create new http client.
func New(client http.Client) *Client {
	return &Client{
		client: client,
		malURL: "https://myanimelist.net",
	}
}

// GetList to get anime/manga list.
func (c *Client) GetList(ctx context.Context, username, mainType string) ([]entity.Entry, int, error) {
	// User's url.
	url := fmt.Sprintf("%s/%slist/%s/load.json?status=7", c.malURL, mainType, username)
	offset := 0

	// Loop them all.
	var list []entity.Entry
	for {
		// Get raw list.
		tmp, code, err := c.getRaw(ctx, fmt.Sprintf("%s&offset=%d", url, offset))
		if err != nil {
			return nil, code, stack.Wrap(ctx, err)
		}

		// Clean image url.
		for _, l := range tmp {
			switch mainType {
			case "anime":
				list = append(list, entity.Entry{
					ID:    l.AnimeID,
					Image: utils.ImageURLCleaner(l.AnimeImage),
				})
			case "manga":
				list = append(list, entity.Entry{
					ID:    l.MangaID,
					Image: utils.ImageURLCleaner(l.MangaImage),
				})
			}
		}

		// Done.
		if len(tmp) < 300 {
			return list, http.StatusOK, nil
		}

		// Next batch.
		offset += 300
	}
}

func (c *Client) getRaw(ctx context.Context, url string) ([]rawList, int, error) {
	var code int
	now := time.Now()
	defer func() { c.log(code, url, time.Since(now)) }()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, http.StatusInternalServerError, stack.Wrap(ctx, err, errors.ErrInternalServer)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, stack.Wrap(ctx, err, errors.ErrInternalServer)
	}
	defer resp.Body.Close()

	code = resp.StatusCode

	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, stack.Wrap(ctx, _errors.New(http.StatusText(resp.StatusCode)))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, stack.Wrap(ctx, err, errors.ErrInternalServer)
	}

	var list []rawList
	if err := json.Unmarshal(body, &list); err != nil {
		return nil, http.StatusInternalServerError, stack.Wrap(ctx, err, errors.ErrInternalServer)
	}

	return list, http.StatusOK, nil
}

func (c *Client) log(code int, url string, t time.Duration) {
	utils.Log(map[string]interface{}{
		"level":    log.DebugLevel,
		"code":     code,
		"url":      url,
		"duration": t.String(),
	})
}
