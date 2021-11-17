package http

type rawList struct {
	AnimeID    int    `json:"anime_id"`
	AnimeImage string `json:"anime_image_path"`
	MangaID    int    `json:"manga_id"`
	MangaImage string `json:"manga_image_path"`
}
