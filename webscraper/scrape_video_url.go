package webscraper

import (
	"GithubRepository/go_anime_api/model"
	"fmt"
	"os"

	"github.com/gocolly/colly/v2"
)

func ScrapeVideoURL(endpoint string) model.WatchInfo {
	watchInfo := model.WatchInfo{}

	collector := colly.NewCollector()
	collector.OnHTML(".play-video", func(e *colly.HTMLElement) {
		videoURL := "https:" + e.ChildAttr("iframe", "src")
		watchInfo.VideoURL = videoURL
	})

	// utils.SetLoggerForColly(collector)

	URLToVisit := fmt.Sprintf("%s%s", os.Getenv("WebsiteURL"), endpoint)
	collector.Visit(URLToVisit)

	return watchInfo
}
