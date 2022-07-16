package models

type Update struct {
	Message Message `json:"message"`
}

type Message struct {
	Chat     Chat        `json:"chat"`
	Photo    []PhotoSize `json:"photo"`
	Document Document    `json:"document"`
}

type Chat struct {
	ID int `json:"id"`
}

type File struct {
	Size int    `json:"file_size"`
	Path string `json:"file_path"`
}

type PhotoSize struct {
	ID string `json:"file_id"`
}

type Document struct {
	Size int `json:"file_size"`
}
