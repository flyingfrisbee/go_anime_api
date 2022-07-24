package router

import (
	"GithubRepository/go_anime_api/db"
	"GithubRepository/go_anime_api/model"
	"GithubRepository/go_anime_api/utils"
	"net/http"
)

func addBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	var body model.AddBookmarkRequest
	err := utils.ParseRequestBody(r, &body)
	if err != nil {
		utils.WriteResponse(w, nil, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	rowsCount, err := db.SaveBookmarkedAnime(body)
	if err != nil {
		utils.WriteResponse(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsCount < 1 {
		utils.WriteResponse(w, nil, "failed bookmarking anime", http.StatusInternalServerError)
		return
	}

	utils.WriteResponse(w, nil, "Success bookmark anime", http.StatusOK)
}

func deleteBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	var body model.DeleteBookmarkRequest
	err := utils.ParseRequestBody(r, &body)
	if err != nil {
		utils.WriteResponse(w, nil, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	rowsCount, err := db.DeleteBookmarkedAnime(body)
	if err != nil {
		utils.WriteResponse(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsCount < 1 {
		utils.WriteResponse(w, nil, "Cannot find bookmarked anime", http.StatusNotFound)
		return
	}

	utils.WriteResponse(w, nil, "Success unbookmark anime", http.StatusOK)
}

func newUpdateBookmarkedAnimeHandler(w http.ResponseWriter, r *http.Request) {
	userToken := r.URL.Query().Get("token")
	bookmarkedAnimes, err := db.GetUpdatedBookmarkedAnimes(userToken)
	if err != nil {
		utils.WriteResponse(w, nil, err.Error(), http.StatusNotFound)
		return
	}

	utils.WriteResponse(w, bookmarkedAnimes, "Success retrieving bookmarked anime with updates", http.StatusOK)
}
