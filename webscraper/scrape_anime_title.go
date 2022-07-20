package webscraper

import (
	"GithubRepository/go_anime_api/model"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

func ScrapeAnimeTitle(keyword string) []model.SearchResult {
	results := []model.SearchResult{}

	collector := colly.NewCollector()
	collector.OnHTML(".items li", func(e *colly.HTMLElement) {
		detailEndpoint := e.ChildAttr(".name a", "href")
		title := e.ChildAttr(".name a", "title")

		animeID := detailEndpoint[(strings.LastIndex(detailEndpoint, "/") + 1):]

		searchResult := model.SearchResult{
			Title:   title,
			AnimeID: animeID,
		}

		results = append(results, searchResult)
	})

	// utils.SetLoggerForColly(collector)

	URLToVisit := fmt.Sprintf("%s/search.html?keyword=%s", os.Getenv("WebsiteURL"), keyword)
	collector.Visit(URLToVisit)

	return results
}
