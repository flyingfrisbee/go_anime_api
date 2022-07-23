package webscraper

import (
	"GithubRepository/go_anime_api/db"
	"GithubRepository/go_anime_api/model"
	"GithubRepository/go_anime_api/utils"
	"fmt"
	"log"
	"strings"
	"time"
)

var (
	repeaterCh = make(chan struct{})
)

func AutomateScrapingRepeatedly() {
	defer close(repeaterCh)

	go scrapeAnimesData()
	channelListener()
}

func channelListener() {
	for {
		select {
		case <-repeaterCh:
			go scrapeAnimesData()
		}
	}
}

func scrapeAnimesData() {
	baseTime := time.Now()
	pageCounter := 1

	for pageCounter <= 10 {
		recentAnimes := ScrapeRecentAnimes(pageCounter)
		for _, v := range recentAnimes {
			animeDetail, isError := ScrapeAnimeDetail(v.FirstID)
			if isError {
				animeDetail, isError = ScrapeAnimeDetail(v.SecondID)
				if isError {
					log.Println("Error when scrape anime detail")
					continue
				}
			}

			addEpisodesAndTimestamp(animeDetail, baseTime)
			db.InsertAnimeData(animeDetail)
			episodesLength := len(animeDetail.Episodes)
			newDataCount := db.CheckIfContainNewData(animeDetail.InternalID, episodesLength)
			switch newDataCount {
			case -1:
			case -2:
				db.InsertBulkEpisodes(animeDetail.InternalID, animeDetail.Episodes)
			default:
				ok := db.UpdateTimestampAndLatestEpisode(animeDetail, baseTime)
				if ok {
					db.InsertBulkEpisodes(animeDetail.InternalID, animeDetail.Episodes[episodesLength-newDataCount:])
				}
			}
		}
		pageCounter++
	}
	repeaterCh <- struct{}{}
}

func TestScrapeAccuracy() {
	baseTime := time.Now()
	errorRecord := map[int]string{}
	startPage := 1
	lastPage := startPage

	for lastPage <= 2 {
		recentAnimes := ScrapeRecentAnimes(lastPage)
		for _, v := range recentAnimes {
			animeDetail, isError := ScrapeAnimeDetail(v.FirstID)
			if !isError {
				addEpisodesAndTimestamp(animeDetail, baseTime)
				continue
			}

			animeDetail, isError = ScrapeAnimeDetail(v.SecondID)
			if isError {
				errorRecord[lastPage] += fmt.Sprintf("$%s %s$", v.FirstID, v.SecondID)
				continue
			}
			addEpisodesAndTimestamp(animeDetail, baseTime)
		}
		lastPage++
	}

	fmt.Printf("error data: %v\n", errorRecord)

	errorCounter := 0
	for _, v := range errorRecord {
		errorCounter += strings.Count(v, "$") / 2
	}
	var percentage float32 = 100
	if errorCounter != 0 {
		percentage = float32(100*((lastPage-startPage+1)*20-errorCounter)) / float32((lastPage-startPage+1)*20)
	}

	fmt.Printf("percent accuracy: %f percent", percentage)
}

func addEpisodesAndTimestamp(animeDetail *model.AnimeDetail, baseTime time.Time) {
	episodes := ScrapeEpisodes(animeDetail.InternalID)
	animeDetail.Episodes = episodes
	animeDetail.LatestEpisode = episodes[len(episodes)-1].ForUI

	animeDetail.UpdatedTimestamp = utils.CalculateTimeAfterSubstractionInMillis(baseTime)
	// uncomment below to inspect data
	// fmt.Println(animeDetail)
}

// fmt.Println(webscraper.ScrapeRecentAnimes(1))
// detail, isError := webscraper.ScrapeAnimeDetail("86")
// if isError {
// 	log.Println("Error occured")
// }
// fmt.Println(detail)
// fmt.Println(webscraper.ScrapeEpisodes("171"))
// fmt.Println(webscraper.ScrapeAnimeTitle("naruto"))
// fmt.Println(webscraper.ScrapeVideoURL("/86-2nd-season-episode-9"))
