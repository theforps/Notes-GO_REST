package models

type Note struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type NoteDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
