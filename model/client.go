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
	IsInternalID *bool `json:"is_internal_id" validate:"required"`
}

type Token struct {
	UserToken                string `json:"user_token" validate:"required"`
	LastMessageSentTimestamp int64  `json:"last_msg_sent_timestamp,omitempty"`
}

type AddBookmarkRequest struct {
	UserToken     string `json:"user_token" validate:"required"`
	InternalID    string `json:"internal_id" validate:"required"`
	LatestEpisode string `json:"latest_episode" validate:"required"`
}

type DeleteBookmarkRequest struct {
	UserToken  string `json:"user_token" validate:"required"`
	InternalID string `json:"internal_id" validate:"required"`
}
