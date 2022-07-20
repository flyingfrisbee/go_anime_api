package model

type RecentAnime struct {
	Title         string `json:"title"`
	LatestEpisode string `json:"latest_episode"`
	ImageURL      string `json:"image_url"`
	FirstID       string `json:"first_id"`
	SecondID      string `json:"second_id"`
}

type AnimeDetail struct {
	AnimeID          string    `json:"anime_id,omitempty"`
	InternalID       string    `json:"internal_id"`
	ImageURL         string    `json:"image_url"`
	Title            string    `json:"title"`
	Type             string    `json:"type"`
	Summary          string    `json:"summary"`
	Genre            string    `json:"genre"`
	ReleasedDate     string    `json:"released_date"`
	AiringStatus     string    `json:"airing_status"`
	Episodes         []Episode `json:"episode_list"`
	LatestEpisode    string    `json:"latest_episode,omitempty"`
	UpdatedTimestamp int64     `json:"updated_timestamp"`
}

type Episode struct {
	ForUI       string `json:"episode_for_ui" bson:"episode_for_ui"`
	ForEndpoint string `json:"episode_for_endpoint" bson:"episode_for_endpoint"`
}

type SearchResult struct {
	Title   string `json:"title"`
	AnimeID string `json:"anime_id"`
}

type WatchInfo struct {
	VideoURL string `json:"video_url"`
}
