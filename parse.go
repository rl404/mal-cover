package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const myAnimeListURL = "https://myanimelist.net"
const animeType = "anime"
const mangaType = "manga"

// RawData is raw JSON response from MyAnimeList.
type RawData struct {
	AnimeID    int    `json:"anime_id"`
	AnimeImage string `json:"anime_image_path"`
	MangaID    int    `json:"manga_id"`
	MangaImage string `json:"manga_image_path"`
}

func getList(username, mainType string) (list []RawData, err error) {
	// Search from cache.
	cacheKey := getCacheKey(username, mainType)
	if getCache(cacheKey, &list) == nil {
		log.Println("from cache", cacheKey)
		return list, nil
	}

	// User URL.
	listURL := fmt.Sprintf("%s/%slist/%s/load.json?status=7", myAnimeListURL, mainType, username)
	offset := 0

	// Loop them all.
	for {
		// Get raw list.
		tmp, err := getRawData(fmt.Sprintf("%s&offset=%v", listURL, offset))
		if err != nil {
			return list, err
		}

		// Clean the image URL.
		for _, l := range tmp {
			l.AnimeImage = imageURLCleaner(l.AnimeImage)
			l.MangaImage = imageURLCleaner(l.MangaImage)
			list = append(list, l)
		}

		if len(tmp) < 300 {
			// Return and save to cache.
			return list, setCache(cacheKey, list)
		}

		// Next batch.
		offset += 300
	}
}

func generateCover(username, mainType, style string) (css string, code int, err error) {
	// Empty style.
	if style == "" {
		return css, http.StatusBadRequest, errors.New("empty style param\nplease check your list style\n\ntry this example:\n\n.animetitle[href*='/{id}/']:before{background-image:url({url})}")
	}

	// Get all anime/manga list.
	list, err := getList(username, mainType)
	if err != nil {
		return css, http.StatusInternalServerError, err
	}

	// Empty list.
	if len(list) == 0 {
		return css, http.StatusNotFound, errors.New("empty list\ngo add them in MyAnimeList")
	}

	// Convert list to css.
	var cssRows []string
	style, _ = url.QueryUnescape(style)
	for _, l := range list {
		if mainType == animeType {
			cssRow := strings.Replace(style, "{id}", strconv.Itoa(l.AnimeID), -1)
			cssRow = strings.Replace(cssRow, "{url}", l.AnimeImage, -1)
			cssRows = append(cssRows, cssRow)
		} else if mainType == mangaType {
			cssRow := strings.Replace(style, "{id}", strconv.Itoa(l.MangaID), -1)
			cssRow = strings.Replace(cssRow, "{url}", l.MangaImage, -1)
			cssRows = append(cssRows, cssRow)
		}
	}

	return strings.Join(cssRows, "\n"), http.StatusOK, nil
}
