package repository

import (
	"awesomeProject/models"
)

// NoteRepository interface for working with notes
type NoteRepository interface {
	GetAllNotes() *[]*models.Note
	GetNoteById(id int) *models.Note
	DeleteNoteById(id int)
	AddNote(note *models.NoteDto)
	UpdateNote(id int, note *models.NoteDto)
}
