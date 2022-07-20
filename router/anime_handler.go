package router

import (
	"GithubRepository/go_anime_api/db"
	"GithubRepository/go_anime_api/model"
	"GithubRepository/go_anime_api/utils"
	"GithubRepository/go_anime_api/webscraper"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func recentAnimeHandler(w http.ResponseWriter, r *http.Request) {
	page := 1
	requestedPage := r.URL.Query().Get("page")
	if requestedPage != "" {
		newPage, err := strconv.Atoi(requestedPage)
		if err != nil {
			utils.WriteResponse(w, nil, err.Error(), http.StatusInternalServerError)
			return
		}
		page = newPage
	}

	if page < 1 {
		utils.WriteResponse(w, nil, "Requested page does not exist", http.StatusBadRequest)
		return
	}

	offset := (page - 1) * 20
	animes, err := db.GetRecentlyUpdatedAnimes(offset)
	if err != nil {
		utils.WriteResponse(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteResponse(w, animes, "Success retrieving recently updated anime", http.StatusOK)
}

func searchTitleHandler(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	if keyword == "" {
		utils.WriteResponse(w, nil, "Cannot find requested anime title", http.StatusBadRequest)
		return
	}

	searchResults := webscraper.ScrapeAnimeTitle(keyword)
	utils.WriteResponse(
		w,
		searchResults,
		fmt.Sprintf(
			"Found %d results for keyword: %s",
			len(searchResults),
			keyword,
		),
		http.StatusOK,
	)
}

func animeDetailHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["anime_id"]

	jsonBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteResponse(w, nil, "Error when reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var body model.ClientBodyAnimeDetail
	err = json.Unmarshal(jsonBytes, &body)
	if err != nil {
		utils.WriteResponse(w, nil, "Error when parsing JSON to struct", http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err = validate.Struct(body)
	if err != nil {
		utils.WriteResponse(w, nil, err.Error(), http.StatusBadRequest)
		return
	}

	// get from db
	if *body.IsInternalID {
		animeDetail := db.GetAnimeDetail(ID)
		animeDetail.Episodes = db.GetEpisodes(animeDetail.InternalID)
		utils.WriteResponse(w, animeDetail, "Success fetch anime detail", http.StatusOK)
		return
	}

	// scrape
	animeDetail, isError := webscraper.ScrapeAnimeDetail(ID)
	if isError {
		utils.WriteResponse(w, nil, "Something bad happened", http.StatusInternalServerError)
		return
	}
	animeDetail.Episodes = webscraper.ScrapeEpisodes(animeDetail.InternalID)
	utils.WriteResponse(w, *animeDetail, "Success fetch anime detail", http.StatusOK)
}

func videoURLHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	endpoint := params["endpoint"]
	watchInfo := webscraper.ScrapeVideoURL("/" + endpoint)
	if watchInfo.VideoURL == "" {
		utils.WriteResponse(w, nil, "Cannot find video URL", http.StatusInternalServerError)
		return
	}

	utils.WriteResponse(w, watchInfo, "Success fetch video URL", http.StatusOK)
}
