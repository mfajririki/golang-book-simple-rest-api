package controllers

import (
	"book-simple-rest-api/database"
	"book-simple-rest-api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	err error
)

func CreateBook(ctx *gin.Context) {
	database.StartDB()
	db := database.GetDB()

	newBook := models.Book{}

	err = ctx.ShouldBind(&newBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	err = db.Create(&newBook).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_message": fmt.Sprint("Error creating book data : ", err),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Book data saved successfully.",
		"book":    newBook,
	})
}

func UpdateBook(ctx *gin.Context) {
	database.StartDB()
	db := database.GetDB()

	bookID := ctx.Param("bookID")

	book := models.Book{}

	updateBook := models.Book{}

	err = ctx.ShouldBind(&updateBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	err := db.Model(&book).Where("id = ?", bookID).Updates(updateBook).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":      fmt.Sprintf("Book with id %v has been successfully updated", bookID),
		"updated_book": updateBook,
	})
}

func GetBooks(ctx *gin.Context) {
	database.StartDB()
	db := database.GetDB()

	books := []models.Book{}

	err := db.Find(&books).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func GetBookById(ctx *gin.Context) {
	database.StartDB()
	db := database.GetDB()

	bookID := ctx.Param("bookID")

	book := models.Book{}

	err := db.First(&book, "id = ?", bookID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func DeleteBook(ctx *gin.Context) {
	database.StartDB()
	db := database.GetDB()

	bookID := ctx.Param("bookID")

	book := models.Book{}

	err := db.Where("id = ?", bookID).Delete(&book).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Error deleting book.",
			"error_message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %v has been successfully deleted", bookID),
	})
}
