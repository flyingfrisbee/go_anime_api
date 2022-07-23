package main

import (
	"GithubRepository/go_anime_api/db"
	"GithubRepository/go_anime_api/fcm"
	"GithubRepository/go_anime_api/router"
	"GithubRepository/go_anime_api/utils"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// uncomment below code to test in local env
	utils.ProvideEnv()

	// webscraper.TestScrapeAccuracy()

	db.StartConnectionToDB()
	defer db.Conn.Close(context.Background())

	// go webscraper.AutomateScrapingRepeatedly()
	go fcm.StartFCMService()

	r := mux.NewRouter()
	router.OpenRoutes(r)

	address := ":8080"
	if port := os.Getenv("PORT"); port != "" {
		address = fmt.Sprintf(":%s", port)
	}

	err := http.ListenAndServe(address, r)
	if err != nil {
		log.Println(err)
		return
	}
}
