package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

// StartHTTP to start serving HTTP.
func StartHTTP() error {
	r := chi.NewRouter()

	// Set default recommended go-chi router middlewares.
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(cors.AllowAll().Handler)

	// Register base routes.
	registerBaseRoutes(r)

	// Register main routes.
	registerRoutes(r)

	log.Println("server listen at", Cfg.Port)
	return http.ListenAndServe(Cfg.Port, r)
}

// respondWithCSS to write response as CSS.
func respondWithCSS(w http.ResponseWriter, statusCode int, data string) {
	w.Header().Set("Content-Type", "text/css")
	w.WriteHeader(statusCode)
	w.Write([]byte(data))
}

// registerBaseRoutes to register common routes.
func registerBaseRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		respondWithCSS(w, http.StatusOK, "it's working\n\nfor more info: https://github.com/rl404/mal-cover")
	})
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		respondWithCSS(w, http.StatusOK, "pong")
	})
	r.NotFound(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respondWithCSS(w, http.StatusNotFound, "page not found\njust like your future")
	}))
}

// registerRoutes to register main routes.
func registerRoutes(r *chi.Mux) {
	r.Get("/auto", getAutoCover)
	r.Get("/{user}/{type}", getCover)
}

// getCover to get user's anime/manga list cover.
func getCover(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "user")
	mainType := chi.URLParam(r, "type")
	style := r.URL.Query().Get("style")

	css, code, err := GenerateCover(username, mainType, style)
	if err != nil {
		respondWithCSS(w, code, err.Error())
	} else {
		respondWithCSS(w, code, css)
	}
}

// getAutoCover to get user's anime/manga list cover automatically.
func getAutoCover(w http.ResponseWriter, r *http.Request) {
	style := r.URL.Query().Get("style")

	userURL := r.Header.Get("Referer")
	userURL = strings.Replace(userURL, "https://myanimelist.net", "", -1)

	reg := regexp.MustCompile(`\/.+(list)\/`)
	mainType := reg.FindString(userURL)
	mainType = strings.Replace(mainType, "/", "", -1)
	mainType = strings.Replace(mainType, "list", "", -1)

	username := strings.Replace(userURL, "/animelist/", "", -1)
	username = strings.Replace(username, "/mangalist/", "", -1)
	split := strings.Split(username, "?")

	if split[0] == "" {
		respondWithCSS(w, http.StatusBadRequest, "call this URL inside your MyAnimeList")
		return
	}

	css, code, err := GenerateCover(split[0], mainType, style)
	if err != nil {
		respondWithCSS(w, code, err.Error())
	} else {
		respondWithCSS(w, code, css)
	}
}

// getRawData to get JSON from MyAnimeList page.
func getRawData(URL string) (list []RawData, err error) {
	log.Println("parsing", URL)
	resp, err := http.Get(URL)
	if err != nil {
		return list, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return list, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return list, err
	}

	err = json.Unmarshal(body, &list)
	if err != nil {
		return list, err
	}

	return list, nil
}
