package db

import (
	"GithubRepository/go_anime_api/model"
	"GithubRepository/go_anime_api/utils"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

func StartConnectionToPostgre() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println(err)
		return
	}
	Conn = conn
}

func InsertAnimeData(animeDetail *model.AnimeDetail) {
	ctx, cancel := CreateContext()
	defer cancel()

	_, err := Conn.Exec(
		ctx,
		InsertQuery,
		animeDetail.AnimeID,
		animeDetail.InternalID,
		animeDetail.ImageURL,
		animeDetail.Title,
		animeDetail.Type,
		animeDetail.Summary,
		animeDetail.Genre,
		animeDetail.ReleasedDate,
		animeDetail.AiringStatus,
		animeDetail.LatestEpisode,
		animeDetail.UpdatedTimestamp,
	)

	if err != nil {
		log.Println(err)
		return
	}

	// fmt.Println(tag.RowsAffected())
}

func UpdateTimestampAndLatestEpisode(animeDetail *model.AnimeDetail, baseTime time.Time) bool {
	timestamp := utils.CalculateTimeAfterSubstractionInMillis(baseTime)
	ctx, cancel := CreateContext()
	defer cancel()

	tag, err := Conn.Exec(ctx, UpdateTimestampQuery, timestamp, animeDetail.LatestEpisode, animeDetail.InternalID)
	if err != nil {
		log.Println(err)
		return false
	}

	return tag.RowsAffected() >= 1
}

func GetRecentlyUpdatedAnimes(offset int) ([]model.ClientRecentAnime, error) {
	result := []model.ClientRecentAnime{}

	ctx, cancel := CreateContext()
	defer cancel()

	rows, err := Conn.Query(ctx, GetRecentAnimesQuery, 20, offset)
	if err != nil {
		return result, fmt.Errorf(err.Error())
	}

	for rows.Next() {
		anime := model.ClientRecentAnime{}
		err := rows.Scan(&anime.ImageURL, &anime.Title, &anime.LatestEpisode, &anime.UpdatedTimestamp, &anime.InternalID)
		if err != nil {
			return result, fmt.Errorf("error when parsing from sql to struct")
		}
		result = append(result, anime)
	}

	return result, nil
}

func GetAnimeDetail(internalID string) model.AnimeDetail {
	var animeDetail model.AnimeDetail

	ctx, cancel := CreateContext()
	defer cancel()

	row := Conn.QueryRow(ctx, GetAnimeDetailQuery, internalID)
	row.Scan(
		&animeDetail.ImageURL,
		&animeDetail.Title,
		&animeDetail.Type,
		&animeDetail.Summary,
		&animeDetail.Genre,
		&animeDetail.ReleasedDate,
		&animeDetail.AiringStatus,
		&animeDetail.UpdatedTimestamp,
		&animeDetail.InternalID,
	)

	return animeDetail
}
