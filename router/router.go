package router

import (
	"GithubRepository/go_anime_api/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func OpenRoutes(r *mux.Router) {
	r.MethodNotAllowedHandler = methodNotAllowedHandler()

	r.HandleFunc("/", recentAnimeHandler).Methods("GET")                  // /?page=1
	r.HandleFunc("/title", searchTitleHandler).Methods("GET")             // /title?keyword=kimetsu
	r.HandleFunc("/anime/{anime_id}", animeDetailHandler).Methods("POST") // /anime/86
	r.HandleFunc("/video/{endpoint}", videoURLHandler).Methods("GET")     // /video//86-episode-1

	r.HandleFunc("/user/create", insertUserTokenHandler).Methods("POST")

	r.HandleFunc("/bookmark/create", addBookmarkHandler).Methods("POST")
	r.HandleFunc("/bookmark/delete", deleteBookmarkHandler).Methods("POST")
	r.HandleFunc("/bookmark/update", newUpdateBookmarkedAnimeHandler).Methods("GET") // /bookmark/update?token='abeceda'
}

func methodNotAllowedHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.WriteResponse(w, nil, "Method not allowed", http.StatusMethodNotAllowed)
	})
}
