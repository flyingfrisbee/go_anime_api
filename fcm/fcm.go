package fcm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type FcmMessage struct {
	To           string           `json:"to"`
	Notification NotificationBody `json:"notification"`
}

type NotificationBody struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func A() {
	fcmMessage := FcmMessage{
		To: "sdad",
		Notification: NotificationBody{
			Title: "4 anime have new update",
			Body:  "Overlord, Kimetsu no yaiba, Violeto",
		},
	}

	jsonBytes, err := json.Marshal(fcmMessage)
	if err != nil {
		log.Println(err)
		return
	}

	r, err := http.NewRequest("POST", os.Getenv("FCMURL"), bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Println(err)
		return
	}

	r.Header.Set("Authorization", fmt.Sprintf("key=%s", os.Getenv("FCMKey")))
	r.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	_, err = client.Do(r)
	if err != nil {
		log.Println(err)
		return
	}
}
