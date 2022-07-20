package utils

import (
	"GithubRepository/go_anime_api/model"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gocolly/colly/v2"
	"golang.org/x/crypto/bcrypt"
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

func HashPassword(input string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func IsHashAndPasswordSame(input string, hashedPwd string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(input))
	if err != nil {
		return false, err
	}

	return true, nil
}
