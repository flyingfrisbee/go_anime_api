package db

import (
	"GithubRepository/go_anime_api/model"
	"GithubRepository/go_anime_api/utils"
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func StartConnectionToPostgre() {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println(err)
		return
	}
	Conn = pool
}

func InsertAnimeData(animeDetail *model.AnimeDetail) {
	ctx, cancel := CreateContext()
	defer cancel()

	_, err := Conn.Exec(
		ctx,
		InsertAnimeQuery,
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

	tag, err := Conn.Exec(ctx, UpdateAnimeTimestampQuery, timestamp, animeDetail.LatestEpisode, animeDetail.InternalID)
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
		return result, err
	}

	for rows.Next() {
		anime := model.ClientRecentAnime{}
		err := rows.Scan(&anime.ImageURL, &anime.Title, &anime.LatestEpisode, &anime.UpdatedTimestamp, &anime.InternalID)
		if err != nil {
			return result, err
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

func SaveUserToken(hashedUserToken string) (int, error) {
	ctx, cancel := CreateContext()
	defer cancel()

	tag, err := Conn.Exec(ctx, InsertUserTokenQuery, hashedUserToken, 0)
	if err != nil {
		return 0, err
	}

	return int(tag.RowsAffected()), nil
}

func SaveBookmarkedAnime(data model.AddBookmarkRequest) (int, error) {
	ctx, cancel := CreateContext()
	defer cancel()

	tag, err := Conn.Exec(ctx, InsertUserAnimeXrefQuery, data.UserToken, data.InternalID, data.LatestEpisode)
	if err != nil {
		return 0, err
	}

	return int(tag.RowsAffected()), nil
}

func DeleteBookmarkedAnime(data model.DeleteBookmarkRequest) (int, error) {
	ctx, cancel := CreateContext()
	defer cancel()

	tag, err := Conn.Exec(ctx, DeleteUserAnimeXrefQuery, data.UserToken, data.InternalID)
	if err != nil {
		return 0, err
	}

	return int(tag.RowsAffected()), nil
}

func GetAllUsersData() ([]model.Token, error) {
	ctx, cancel := CreateContext()
	defer cancel()

	result := []model.Token{}

	rows, err := Conn.Query(ctx, GetAllUserDataQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		token := model.Token{}
		err := rows.Scan(&token.UserToken, &token.LastMessageSentTimestamp)
		if err != nil {
			return nil, err
		}
		result = append(result, token)
	}

	return result, nil
}

func GetUpdatedBookmarkedAnimes(userToken string) ([]string, error) {
	ctx, cancel := CreateContext()
	defer cancel()

	result := []string{}

	rows, err := Conn.Query(ctx, GetUpdatedBookmarkedAnimesQuery, userToken)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var title string
		err := rows.Scan(&title)
		if err != nil {
			return nil, err
		}
		result = append(result, title)
	}

	return result, nil
}

func UpdateUserTimestamp(userToken string) error {
	ctx, cancel := CreateContext()
	defer cancel()

	_, err := Conn.Exec(ctx, UpdateUserMsgTimestamp, time.Now().UnixMilli(), userToken)
	if err != nil {
		return err
	}

	return nil
}
