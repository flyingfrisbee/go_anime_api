package fcm

import (
	"GithubRepository/go_anime_api/db"
	"GithubRepository/go_anime_api/model"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type FcmMessage struct {
	To           string           `json:"to"`
	Notification NotificationBody `json:"notification"`
}

type NotificationBody struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

var (
	repeaterCh         = make(chan struct{})
	twentyHourInMillis = 20 * time.Hour.Milliseconds()
)

func StartFCMService() {
	defer close(repeaterCh)

	go sendFCMToElligibleUsers()
	channelListener()
}

func channelListener() {
	for {
		select {
		case <-repeaterCh:
			time.Sleep(10 * time.Minute)
			go sendFCMToElligibleUsers()
		}
	}
}

func sendFCMToElligibleUsers() {
	defer func() {
		repeaterCh <- struct{}{}
	}()

	users, err := retrieveUsersData()
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range users {
		timeDifference := time.Now().UnixMilli() - v.LastMessageSentTimestamp
		shouldSendMessage := timeDifference > twentyHourInMillis
		if shouldSendMessage {
			animeTitlesWithUpdate, err := db.GetUpdatedBookmarkedAnimes(v.UserToken)
			if err != nil {
				log.Println(err)
				return
			}

			if len(animeTitlesWithUpdate) < 1 {
				continue
			}

			err = sendFCMToUser(v.UserToken, animeTitlesWithUpdate)
			if err != nil {
				log.Println(err)
				return
			}

			err = db.UpdateUserTimestamp(v.UserToken)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func retrieveUsersData() ([]model.Token, error) {
	users, err := db.GetAllUsersData()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func sendFCMToUser(userToken string, animeTitlesWithUpdate []string) error {
	var msgBody string
	for _, v := range animeTitlesWithUpdate {
		msgBody += v + ", "
	}
	msgBody = msgBody[:len(msgBody)-2]

	fcmMessage := FcmMessage{
		To: userToken,
		Notification: NotificationBody{
			Title: fmt.Sprintf("%d bookmarked anime have new update", len(animeTitlesWithUpdate)),
			Body:  msgBody,
		},
	}

	jsonBytes, err := json.Marshal(fcmMessage)
	if err != nil {
		log.Println(err)
		return err
	}

	r, err := http.NewRequest("POST", os.Getenv("FCMURL"), bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Println(err)
		return err
	}

	r.Header.Set("Authorization", fmt.Sprintf("key=%s", os.Getenv("FCMKey")))
	r.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	_, err = client.Do(r)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
