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
	"go.mongodb.org/mongo-driver/mongo"
)

var noteCollection *mongo.Collection = database.OpenCollection(database.Client, "note")

func CreateNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
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

func GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
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

func GetNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
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

func UpdateNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": "Update Note",
		})
	}
}

func DeleteNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": "Delete Note",
		})
	}
}
