package db

var (
	InsertQuery = `
    INSERT INTO stream_anime.anime (
		anime_id, internal_id, image_url, 
		title, type, summary, genre, 
		released_date, airing_status, latest_episode, 
		updated_timestamp
	)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
    ON CONFLICT (internal_id) DO NOTHING
    `

	GetAnimeDetailQuery = `
	SELECT image_url, title, type, summary, 
		genre, released_date, airing_status, 
		updated_timestamp, internal_id
	FROM stream_anime.anime
	WHERE internal_id = $1
	`

	UpdateTimestampQuery = `
	UPDATE stream_anime.anime 
	SET updated_timestamp = $1, 
		latest_episode = $2
	WHERE internal_id = $3;
	`

	GetRecentAnimesQuery = `
	SELECT image_url, title, latest_episode, updated_timestamp, internal_id
	FROM stream_anime.anime
	ORDER BY updated_timestamp DESC
	LIMIT $1 OFFSET $2
	`
)
