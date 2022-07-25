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
	Ok     bool       `json:"ok"`
	Result FileResult `json:"result"`
}

type FileResult struct {
	UniqueID string `json:"file_unique_id"`
	Path     string `json:"file_path"`
}

type PhotoSize struct {
	ID string `json:"file_id"`
}

type Document struct {
	Size int `json:"file_size"`
}
