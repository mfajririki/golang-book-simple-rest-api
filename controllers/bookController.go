package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID int    `json:"car_id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"description"`
}

var BookDatas = []Book{}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	err := ctx.ShouldBindJSON(&newBook)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.BookID = len(BookDatas) + 1
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"book": newBook,
	})
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false

	var updateBook Book

	err := ctx.ShouldBindJSON(&updateBook)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookDatas {
		if bookID == string(rune(book.BookID)) {
			condition = true
			BookDatas[i] = updateBook
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("Book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %v has been successfully updated", bookID),
	})
}

func GetBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false

	var bookData Book

	for i, book := range BookDatas {
		if bookID == string(rune(book.BookID)) {
			condition = true
			bookData = BookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("Book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": bookData,
	})
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false

	var bookIndex int

	for i, book := range BookDatas {
		if bookID == string(rune(book.BookID)) {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("Book with id %v not found", bookID),
		})
		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %v has been successfully deleted", bookID),
	})
}
