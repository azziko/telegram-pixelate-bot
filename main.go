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

const (
	URL  = "https://api.telegram.org/bot"
	PORT = "80"
)

func update(w http.ResponseWriter, r *http.Request) {

	message := &models.ReceiveMessage{}

	chatID := 0
	msgText := ""

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		fmt.Println(err)
	}

	if message.Message.Chat.ID != 0 {
		fmt.Println(message.Message.Chat.ID, message.Message.Text)
		chatID = message.Message.Chat.ID
		msgText = message.Message.Text
	} else {
		fmt.Println(message.ChannelPost.Chat.ID, message.ChannelPost.Text)
		chatID = message.ChannelPost.Chat.ID
		msgText = message.ChannelPost.Text
	}

	respMsg := fmt.Sprintf("%s%s/sendMessage?chat_id=%d&text=Received: %s", URL, TOKEN, chatID, msgText)

	_, err = http.Get(respMsg)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	http.HandleFunc("/", update)

	fmt.Println("Listenning on port", PORT, ".")
	if err := http.ListenAndServe("/", nil); err != nil {
		log.Fatal(err)
	}
}
