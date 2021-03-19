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
	"github.com/rs/cors"
)

func startHTTP() error {
	r := chi.NewRouter()
	r.Use(cors.AllowAll().Handler)

	// Register base routes.
	registerBaseRoutes(r)

	// Register main routes.
	registerRoutes(r)

	log.Println("server listen at", cfg.Port)
	return http.ListenAndServe(cfg.Port, r)
}

func respondWithCSS(w http.ResponseWriter, statusCode int, data string) {
	w.Header().Set("Content-Type", "text/css")
	w.WriteHeader(statusCode)
	_, _ = w.Write([]byte(data))
}

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

func registerRoutes(r *chi.Mux) {
	r.Get("/auto", getAutoCover)
	r.Get("/{user}/{type}", getCover)
}

func getCover(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "user")
	mainType := chi.URLParam(r, "type")
	style := r.URL.Query().Get("style")

	css, code, err := generateCover(username, mainType, style)
	if err != nil {
		respondWithCSS(w, code, err.Error())
	} else {
		respondWithCSS(w, code, css)
	}
}

func getAutoCover(w http.ResponseWriter, r *http.Request) {
	style := r.URL.Query().Get("style")

	userURL := r.Header.Get("Referer")
	userURL = strings.Replace(userURL, myAnimeListURL, "", -1)

	reg := regexp.MustCompile(`\/.+(list)\/`)
	mainType := reg.FindString(userURL)
	mainType = strings.Replace(mainType, "/", "", -1)
	mainType = strings.Replace(mainType, "list", "", -1)

	username := strings.Replace(userURL, "/animelist/", "", -1)
	username = strings.Replace(username, "/mangalist/", "", -1)
	split := strings.Split(username, "?")

	if split[0] == "" {
		respondWithCSS(w, http.StatusBadRequest, "call this URL inside your MyAnimeList CSS file")
		return
	}

	css, code, err := generateCover(split[0], mainType, style)
	if err != nil {
		respondWithCSS(w, code, err.Error())
	} else {
		respondWithCSS(w, code, css)
	}
}

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

	if err = json.Unmarshal(body, &list); err != nil {
		return list, err
	}

	return list, nil
}
