package model

type GenericResponse struct {
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
}

type ClientRecentAnime struct {
	ImageURL         string `json:"image_url"`
	Title            string `json:"title"`
	LatestEpisode    string `json:"latest_episode"`
	UpdatedTimestamp int64  `json:"updated_timestamp"`
	InternalID       string `json:"internal_id"`
}

type ClientBodyAnimeDetail struct {
	IsInternalID bool `json:"is_internal_id"`
}