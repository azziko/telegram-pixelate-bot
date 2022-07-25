package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pixelate/app/client"
	"pixelate/app/client/telegram/models"
	"pixelate/app/processor"
	"pixelate/app/storage"
)

const (
	storageHost = "https://telegram-pixelate-bot.herokuapp.com/"
)

type UpdateHandler struct {
	TgClient client.Client
	Storage  storage.Storage
}

func NewUpdateHandler(c client.Client, s storage.Storage) *UpdateHandler {
	return &UpdateHandler{
		TgClient: c,
		Storage:  s,
	}
}

func (u *UpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	message := ""
	upd := &models.Update{}

	if json.NewDecoder(r.Body).Decode(upd) != nil {
		return
	}
	defer r.Body.Close()

	chatID := upd.Message.Chat.ID

	switch {
	case len(upd.Message.Photo) > 0:
		file := &models.File{}
		caption := "Here you go @pixelate_bot 💌"

		fileBytes, err := u.TgClient.GetFile(upd.Message.Photo[len(upd.Message.Photo)-1].ID)
		if err != nil {
			fmt.Println(err)
			return
		}

		if json.Unmarshal(fileBytes, file) != nil {
			return
		}

		imageBytes, err := u.TgClient.DownloadImage(file.Result.Path)
		if err != nil {
			fmt.Println(err)
			return
		}

		filename := fmt.Sprintf("output%s.png", file.Result.UniqueID)

		readerCloser := toReaderCloser(imageBytes)
		filepath := u.Storage.FilePath(filename)
		if err := processor.Pixelate(filepath, readerCloser); err != nil {
			fmt.Println(err)
		}

		serverPath := storageHost + filepath
		if err := u.TgClient.SendPhoto(chatID, serverPath, caption); err != nil {
			fmt.Println(err)
		}

		u.Storage.Delete(filename)

	case upd.Message.Document.Size > 0:
		message = "Please send me the picture as a 'Photo', not as a 'File'."
		if err := u.TgClient.SendMessage(chatID, message); err != nil {
			fmt.Println(err)
		}
	default:
		message = "Send me a picture to pixelate"
		if err := u.TgClient.SendMessage(chatID, message); err != nil {
			fmt.Println(err)
		}
	}
}

func toReaderCloser(b []byte) io.ReadCloser {
	reader := bytes.NewReader(b)

	return io.NopCloser(reader)
}
