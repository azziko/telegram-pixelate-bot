package clients

import "net/url"

type Client interface {
	GetFile(fileID int) ([]byte, error)
	SendMessage(chatID int, text string) error
	SendPhoto(chatID int, fileID int) error
	DownloadImage(chatID int, filePath string) ([]byte, error)
	doRequest(method string, query url.Values) error
}
