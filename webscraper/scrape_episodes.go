package webscraper

import (
	"GithubRepository/go_anime_api/model"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

func ScrapeEpisodes(internalID string) []model.Episode {
	episodes := []model.Episode{}

	collector := colly.NewCollector()
	collector.OnHTML("li", func(e *colly.HTMLElement) {
		episodeEndpoint := e.ChildAttr("a", "href")
		episodeUI := strings.Replace(e.ChildText("a .name"), "EP ", "", 1)

		episode := model.Episode{
			ForUI:       episodeUI,
			ForEndpoint: episodeEndpoint,
		}
		episodes = append(episodes, episode)
	})

	// utils.SetLoggerForColly(collector)

	URLToVisit := fmt.Sprintf("%s?ep_start=0&ep_end=100000&id=%s", os.Getenv("EpisodeURL"), internalID)
	collector.Visit(URLToVisit)

	for i, j := 0, (len(episodes) - 1); i <= j; i, j = i+1, j-1 {
		temp := episodes[i]
		episodes[i] = episodes[j]
		episodes[j] = temp
	}

	return episodes
}
