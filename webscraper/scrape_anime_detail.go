package webscraper

import (
	"GithubRepository/go_anime_api/model"
	"GithubRepository/go_anime_api/utils"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

func ScrapeAnimeDetail(animeID string) (*model.AnimeDetail, bool) {
	var animeDetail *model.AnimeDetail
	var isError = false

	collector := colly.NewCollector()
	collector.OnHTML(".anime_info_body .anime_info_body_bg", func(e *colly.HTMLElement) {
		animeTitle := e.ChildText("h1")
		imgURL := e.ChildAttr("img", "src")

		animeInfoUnparsed := e.ChildTexts("p")
		animeType := strings.TrimSpace(animeInfoUnparsed[1][len(os.Getenv("AnimeInfoType")):])
		animeSummary := strings.TrimSpace(animeInfoUnparsed[2][len(os.Getenv("AnimeInfoSummary")):])
		animeGenre := strings.TrimSpace(animeInfoUnparsed[3][len(os.Getenv("AnimeInfoGenre")):])
		animeReleasedDate := strings.TrimSpace(animeInfoUnparsed[4][len(os.Getenv("AnimeInfoReleased")):])
		animeAiringStatus := strings.TrimSpace(animeInfoUnparsed[5][len(os.Getenv("AnimeInfoStatus")):])

		animeDetail = &model.AnimeDetail{
			ImageURL:     imgURL,
			Title:        animeTitle,
			Type:         animeType,
			Summary:      animeSummary,
			Genre:        animeGenre,
			ReleasedDate: animeReleasedDate,
			AiringStatus: animeAiringStatus,
		}
	})

	collector.OnHTML(".anime_info_episodes .anime_info_episodes_next", func(e *colly.HTMLElement) {
		internalID := e.ChildAttr("#movie_id", "value")
		if internalID == "" {
			isError = true
			return
		}

		animeDetail.AnimeID = animeID
		animeDetail.InternalID = internalID
	})

	utils.SetLoggerForColly(collector)

	URLToVisit := fmt.Sprintf("%s/category/%s", os.Getenv("WebsiteURL"), animeID)
	collector.Visit(URLToVisit)

	if animeDetail == nil {
		isError = true
	}

	return animeDetail, isError
}
