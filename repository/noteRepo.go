package repository

import (
	"awesomeProject/models"
	"database/sql"
)

// NoteImplementation implementation for working with notes
type NoteImplementation struct {
	Db *sql.DB
}

// UpdateNote UPDATE note by identifier and DTO
func (impl NoteImplementation) UpdateNote(id int, note *models.NoteDto) {
	_, err := impl.Db.Exec("UPDATE NOTES SET Name = ?, Description = ? WHERE id = ?", (*note).Name, (*note).Description, id)
	if err != nil {
		panic(err)
	}
}

// AddNote ADD note by DTO
func (impl NoteImplementation) AddNote(note *models.NoteDto) {
	_, err := impl.Db.Exec("INSERT INTO NOTES ('Name', 'Description') VALUES ( ?, ?)", (*note).Name, (*note).Description)
	if err != nil {
		panic(err)
	}
}

// DeleteNoteById DELETE note by identifier
func (impl NoteImplementation) DeleteNoteById(id int) {
	_, err := impl.Db.Exec("DELETE FROM NOTES WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
}

// GetNoteById GET note by identifier
func (impl NoteImplementation) GetNoteById(id int) *models.Note {
	result, err := impl.Db.Query("SELECT * FROM NOTES WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	defer result.Close()

	note := models.Note{}

	for result.Next() {
		err = result.Scan(&note.Id, &note.Name, &note.Description)
		if err != nil {
			panic(err)
		}
	}
	return &note

}

// GetAllNotes GET array of notes
func (impl NoteImplementation) GetAllNotes() *[]*models.Note {

	result, err := impl.Db.Query("SELECT * FROM NOTES")
	if err != nil {
		panic(err)
	}
	defer result.Close()

	var notes []*models.Note

	for result.Next() {
		note := &models.Note{}
		err := result.Scan(&note.Id, &note.Name, &note.Description)
		if err != nil {
			panic(err)
		}
		notes = append(notes, note)
	}

	return &notes
}
