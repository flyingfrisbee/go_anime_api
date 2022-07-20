package webscraper

import (
	"GithubRepository/go_anime_api/model"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

func ScrapeRecentAnimes(page int) []model.RecentAnime {
	animes := []model.RecentAnime{}

	collector := colly.NewCollector()
	collector.OnHTML(".items li", func(e *colly.HTMLElement) {
		latestEpisode := e.ChildText(".episode")
		title := e.ChildAttr(".name a", "title")

		streamEndpoint := e.ChildAttr(".img a", "href")
		imgURL := e.ChildAttr(".img a img", "src")

		var firstID string
		if strings.Contains(streamEndpoint, "-episode-") {
			firstID = strings.ToLower(streamEndpoint[1:strings.Index(streamEndpoint, "-episode-")])
		} else {
			firstID = strings.ToLower(streamEndpoint[1:])
		}

		secondID := strings.ToLower(imgURL[(strings.LastIndex(imgURL, "/") + 1):(len(imgURL) - 4)])

		anime := model.RecentAnime{
			Title:         title,
			LatestEpisode: latestEpisode,
			ImageURL:      imgURL,
			FirstID:       firstID,
			SecondID:      secondID,
		}

		animes = append(animes, anime)
	})

	// utils.SetLoggerForColly(collector)

	URLToVisit := fmt.Sprintf("%s?page=%d", os.Getenv("WebsiteURL"), page)
	collector.Visit(URLToVisit)

	return animes
}
