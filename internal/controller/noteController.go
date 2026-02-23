package controller

import (
	"RESTAURANT-MANAGEMENT/internal/database"
	"RESTAURANT-MANAGEMENT/internal/model"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateNote godoc
// @Summary Create a new note
// @Description Create a new note with title and text content
// @Tags Note
// @Accept json
// @Produce json
// @Param note body model.NoteModel true "Note data (title and text required)"
// @Success 200 {object} map[string]interface{} "Note created successfully"
// @Failure 400 {object} map[string]string "Bad request - validation error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /note [post]
func CreateNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var noteCollection = database.Collection("notes")
		var note model.NoteModel
		if err := c.BindJSON(&note); err != nil {
			c.JSON(
				http.StatusBadRequest, gin.H{"error": err.Error()},
			)
		}

		validate := validator.New()
		validationErr := validate.Struct(note)
		if validationErr != nil {
			c.JSON(
				http.StatusBadRequest, gin.H{"error": validationErr.Error()},
			)
		}

		note.Created_at = time.Now()
		note.Updated_at = time.Now()
		note.ID = primitive.NewObjectID()
		note.Note_id = note.ID.Hex()
		result, insertErr := noteCollection.InsertOne(ctx, note)
		defer cancel()
		if insertErr != nil {
			msg := "Note was not created"
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": msg},
			)
		}
		c.JSON(http.StatusOK, result)
	}
}

// GetNotes godoc
// @Summary Get all notes
// @Description Retrieve a list of all notes
// @Tags Note
// @Produce json
// @Success 200 {object} []model.NoteModel "List of notes"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /notes [get]
func GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var noteCollection = database.Collection("notes")
		result, err := noteCollection.Find(ctx, bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "Error occured while listing notes"},
			)
		}
		var allNotes []bson.M
		if err = result.All(ctx, &allNotes); err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "Error occured while listing notes"},
			)
		}
		c.JSON(http.StatusOK, allNotes)
	}
}

// GetNote godoc
// @Summary Get a specific note
// @Description Retrieve note details by note ID
// @Tags Note
// @Produce json
// @Param id path string true "Note ID"
// @Success 200 {object} model.NoteModel "Note details"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /note/{id} [get]
func GetNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var noteCollection = database.Collection("notes")
		noteId := c.Param("note_id")
		var note model.NoteModel

		err := noteCollection.FindOne(ctx, bson.M{"note_id": noteId}).Decode(&note)
		defer cancel()
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "Error occured while fetching the note"},
			)
		}
		c.JSON(http.StatusOK, note)
	}
}

// UpdateNote godoc
// @Summary Update a note
// @Description Update note details by note ID
// @Tags Note
// @Accept json
// @Produce json
// @Param id path string true "Note ID"
// @Param note body model.NoteModel true "Updated note data"
// @Success 200 {object} map[string]interface{} "Note updated successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /note/{id} [put]
func UpdateNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": "Update Note",
		})
	}
}

// DeleteNote godoc
// @Summary Delete a note
// @Description Delete a note by note ID
// @Tags Note
// @Produce json
// @Param id path string true "Note ID"
// @Success 200 {object} map[string]string "Note deleted successfully"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /note/{id} [delete]
func DeleteNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": "Delete Note",
		})
	}
}
