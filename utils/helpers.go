package utils

import (
	"GithubRepository/go_anime_api/model"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gocolly/colly/v2"
)

func SetLoggerForColly(collector *colly.Collector) {
	collector.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})
}

func CalculateTimeAfterSubstractionInMillis(baseTime time.Time) int64 {
	elapsedTime := time.Since(baseTime)
	return baseTime.Add(-elapsedTime).UnixMilli()
}

func WriteResponse(w http.ResponseWriter, data interface{}, message string, statusCode int) {
	response := model.GenericResponse{
		Data:       data,
		Message:    message,
		StatusCode: statusCode,
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonBytes)
}
