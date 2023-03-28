package controllers

import (
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
