package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rl404/fairy/log"
	"github.com/rl404/fairy/monitoring/newrelic/middleware"
	"github.com/rl404/mal-cover/internal/service"
	"github.com/rl404/mal-cover/internal/utils"
)

// API contains all functions for api endpoints.
type API struct {
	service service.Service
}

// New to create new api endpoints.
func New(service service.Service) *API {
	return &API{
		service: service,
	}
}

// Register to register api routes.
func (api *API) Register(r chi.Router, nrApp *newrelic.Application) {
	r.Route("/", func(r chi.Router) {
		r.Use(middleware.NewHTTP(nrApp))
		r.Use(log.HTTPMiddlewareWithLog(utils.GetLogger(0), log.APIMiddlewareConfig{Error: true}))
		r.Use(log.HTTPMiddlewareWithLog(utils.GetLogger(1), log.APIMiddlewareConfig{
			RequestHeader:  true,
			RequestBody:    true,
			ResponseHeader: true,
			ResponseBody:   true,
			RawPath:        true,
			Error:          true,
		}))
		r.Use(utils.Recoverer)

		r.Get("/{user}/{type}", api.handleGetCover)
	})
}

func (api *API) handleGetCover(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "user")
	mainType := chi.URLParam(r, "type")
	style := r.URL.Query().Get("style")

	css, code, err := api.service.GenerateCover(r.Context(), service.GenerateCoverRequest{
		Username: username,
		Type:     mainType,
		Style:    style,
	})

	utils.RespondWithCSS(w, code, css, err)
}
