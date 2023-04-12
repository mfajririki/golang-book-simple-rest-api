package routers

import (
	"book-simple-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/books", controllers.CreateBook)
	router.GET("/books", controllers.GetBooks)
	router.PUT("/books/:bookID", controllers.UpdateBook)
	router.GET("/books/:bookID", controllers.GetBookById)
	router.DELETE("/books/:bookID", controllers.DeleteBook)

	return router
}
