package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"pixelate/app/client/telegram/models"
)

var TOKEN = os.Getenv("TOKEN")
var PORT = os.Getenv("PORT")

const (
	URL = "https://api.telegram.org/bot"
)

func update(w http.ResponseWriter, r *http.Request) {

	message := &models.Update{}

	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		fmt.Println(err)
	}

	msgText := ""
	chatID := message.Message.Chat.ID

	switch {
	case len(message.Message.Photo) > 0:
		msgText = "Photo received"

	case message.Message.Document.Size > 0:
		msgText = "Please send me the picture as a 'Photo', not as a 'File'."

	default:
		msgText = "Please send me a picture to begin"
	}

	respMsg := fmt.Sprintf("%s%s/sendMessage?chat_id=%d&text=%s", URL, TOKEN, chatID, msgText)

	if _, err := http.Get(respMsg); err != nil {
		fmt.Println(err)
	}
}

func main() {

	http.HandleFunc("/", update)

	fmt.Println("Listenning on port", PORT)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatal(err)
	}
}
