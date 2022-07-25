package client

type Client interface {
	GetFile(fileID string) ([]byte, error)
	SendMessage(chatID int, text string) error
	SendPhoto(chatID int, filepath string, caption string) error
	DownloadImage(filePath string) ([]byte, error)
}
