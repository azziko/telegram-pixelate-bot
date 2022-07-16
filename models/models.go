package models

type Client interface {
	GetFile(fileID int) (File, error)
	SendMessage(chatID int, text string) error
	SendPhoto(chatID int, fileID int) error
}

type Update struct {
	Message Message `json:"message"`
}

type Message struct {
	Chat     Chat              `json:"chat"`
	Text     string            `json:"text"`
	Photo    []PhotoSize       `json:"photo"`
	Entities []MessageEntities `json:"entities"`
}

type Chat struct {
	ID int `json:"id"`
}

type MessageEntities struct {
	Type string `json:"bot_command"`
}

type File struct {
	ID   int `json:"file_id"`
	Size int `json:"file_size"`
	Path int `json:"file_path"`
}

type PhotoSize struct {
	ID     string `json:"file_id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
