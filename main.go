package main

import (
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {

	err := CreateDb()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/get-notes", getNotes)
	r.GET("/get-note/:id", getNote)
	r.DELETE("/delete-note/:id", deleteNote)
	r.PUT("/update-note/:id", updateNote)
	r.POST("/add-note", createNote)

	r.SetTrustedProxies(nil)

	err = r.Run("localhost:5050")
	if err != nil {
		panic(err)
	}
}

// CREATE NOTE
func createNote(c *gin.Context) {
	noteRepo, db := DbInit()
	defer DbClose(db)

	var createdNote *models.NoteDto

	if err := c.BindJSON(&createdNote); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	(*noteRepo).AddNote(createdNote)

	c.JSON(http.StatusOK, "Success")
}

// UPDATE NOTE
func updateNote(c *gin.Context) {
	noteRepo, db := DbInit()
	defer DbClose(db)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	var updatedNote *models.NoteDto

	if err = c.BindJSON(&updatedNote); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	(*noteRepo).UpdateNote(id, updatedNote)

	c.JSON(http.StatusOK, "Success")
}

// DELETE NOTE
func deleteNote(c *gin.Context) {
	noteRepo, db := DbInit()
	defer DbClose(db)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	(*noteRepo).DeleteNoteById(id)

	c.JSON(http.StatusOK, "Success")
}

// GET NOTES
func getNotes(c *gin.Context) {
	noteRepo, db := DbInit()
	defer DbClose(db)

	notes := (*noteRepo).GetAllNotes()

	c.JSON(http.StatusOK, gin.H{
		"notes": notes,
	})

}

// GET NOTE
func getNote(c *gin.Context) {
	noteRepo, db := DbInit()
	defer DbClose(db)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	note := (*noteRepo).GetNoteById(id)

	c.JSON(http.StatusOK, gin.H{
		"note": note,
	})

}
