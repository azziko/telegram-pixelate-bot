package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"pixelate/models"
)

var TOKEN = os.Getenv("TOKEN")
var PORT = os.Getenv("PORT")

const (
	URL = "https://api.telegram.org/bot"
)

func update(w http.ResponseWriter, r *http.Request) {

	message := &models.Update{}

	chatID := 0
	msgText := ""

	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		fmt.Println(err)
	}

	if message.Message.Chat.ID != 0 {
		fmt.Println(message.Message.Chat.ID, message.Message.Text)
		chatID = message.Message.Chat.ID
		msgText = message.Message.Text
	}

	respMsg := fmt.Sprintf("%s%s/sendMessage?chat_id=%d&text=Received: %s", URL, TOKEN, chatID, msgText)

	if _, err := http.Get(respMsg); err != nil {
		fmt.Println(err)
	}
}

func start(w http.ResponseWriter, r *http.Request) {
	message := &models.Update{}

	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		fmt.Println(err)
	}

	msgText := `Hi there! Send a picture to begin`
	chatID := message.Message.Chat.ID

	respMsg := fmt.Sprintf("%s%s/sendMessage?chat_id=%d&text=%s", URL, TOKEN, chatID, msgText)

	if _, err := http.Get(respMsg); err != nil {
		fmt.Println(err)
	}
}

func help(w http.ResponseWriter, r *http.Request) {
	message := &models.Update{}

	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		fmt.Println(err)
	}

	msgText := `Simply send me a picture to pixelate and wait until it's done. 
	p.s It might take some time to process a picture. Thanks for your patience`
	chatID := message.Message.Chat.ID

	respMsg := fmt.Sprintf("%s%s/sendMessage?chat_id=%d&text=%s", URL, TOKEN, chatID, msgText)

	if _, err := http.Get(respMsg); err != nil {
		fmt.Println(err)
	}
}

func main() {

	http.HandleFunc("/", update)
	http.HandleFunc("/start", start)
	http.HandleFunc("/help", help)

	fmt.Println("Listenning on port", PORT, ".")
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatal(err)
	}
}
