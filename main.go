package main

import (
	"GithubRepository/go_anime_api/db"
	"GithubRepository/go_anime_api/router"
	"GithubRepository/go_anime_api/utils"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// uncomment below code to test in local env
	utils.ProvideEnv()

	// webscraper.TestScrapeAccuracy()

	db.StartConnectionToDB()
	defer func() {
		db.Conn.Close(context.Background())
	}()

	// go func() {
	// 	webscraper.AutomateScrapingRepeatedly()
	// }()

	r := mux.NewRouter()
	router.OpenRoutes(r)

	address := "localhost:8080"
	if port := os.Getenv("PORT"); port != "" {
		address = fmt.Sprintf(":%s", port)
	}

	http.ListenAndServe(address, r)
}
