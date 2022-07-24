package db

var (
	InsertAnimeQuery = `
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
		updated_timestamp, internal_id, latest_episode
	FROM stream_anime.anime
	WHERE internal_id = $1
	`

	UpdateAnimeTimestampQuery = `
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

	InsertUserTokenQuery = `
	INSERT INTO stream_anime.user (user_token, last_message_sent_timestamp)
    VALUES ($1, $2)
    ON CONFLICT (user_token) DO NOTHING
	`

	InsertUserAnimeXrefQuery = `
	INSERT INTO stream_anime.user_anime_xref (user_token, internal_id, latest_episode)
	VALUES ($1, $2, $3);
	`

	DeleteUserAnimeXrefQuery = `
	DELETE FROM stream_anime.user_anime_xref
	WHERE user_token = $1
		AND internal_id = $2;
	`

	GetAllUserDataQuery = `
	SELECT * FROM stream_anime.user;
	`

	GetUpdatedBookmarkedAnimesQuery = `
	SELECT title, a.internal_id, a.latest_episode, a.image_url
	FROM stream_anime.user_anime_xref uax
	JOIN stream_anime.anime a
		ON uax.internal_id = a.internal_id
		AND uax.user_token = $1
	WHERE uax.latest_episode < a.latest_episode;
	`

	UpdateUserMsgTimestamp = `
	UPDATE stream_anime.user 
	SET last_message_sent_timestamp = $1
	WHERE user_token = $2;
	`
)
