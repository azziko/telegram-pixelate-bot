package telegram

import (
	"net/http"
	"pixelate/models"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func NewClient(host string, token string) *Client {
	return &Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) GetFile(fileID int) (models.File, error) {
	var file models.File

	return file, nil
}

func (c *Client) SendMessage(chatID int, text string) error {

	return nil
}

func SendPhoto(chatID int, fileID int) error {

	return nil
}
