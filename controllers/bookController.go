package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Book struct {
	BookID int    `json:"book_id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"description"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "SwordMaster"
	dbname   = "db-book"
)

var (
	db  *sql.DB
	err error
)

func ConnectToDatabase() {

}

func CreateBook(ctx *gin.Context) {
	// connect to database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("ERROR : ", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("ERROR : ", err)
		return
	}

	fmt.Println("Successfully connected to database.")

	var newBook Book

	sqlStatement := `
	INSERT INTO books (title, author, description)
	VALUES ($1, $2, $3)
	RETURNING *
	`

	err = ctx.ShouldBindJSON(&newBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	err = db.QueryRow(sqlStatement, newBook.Title, newBook.Author, newBook.Desc).
		Scan(&newBook.BookID, &newBook.Title, &newBook.Author, &newBook.Desc)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Book data saved successfully.",
		"book":    newBook,
	})
}

func UpdateBook(ctx *gin.Context) {
	// connect to database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("ERROR : ", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("ERROR : ", err)
		return
	}

	fmt.Println("Successfully connected to database.")

	bookID := ctx.Param("bookID")

	var updateBook Book

	sqlStatement := `
	UPDATE books
	SET title = $2, author = $3, description = $4
	WHERE book_id = $1`

	err = ctx.ShouldBindJSON(&updateBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	bookIDInt, _ := strconv.Atoi(bookID)
	updateBook.BookID = bookIDInt

	res, err := db.Exec(sqlStatement, &updateBook.BookID, &updateBook.Title, &updateBook.Author, &updateBook.Desc)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": err.Error(),
		})
		return
	}

	count, err := res.RowsAffected()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":         fmt.Sprintf("Book with id %v has been successfully updated", bookID),
		"Row(s) affected": count,
	})
}

func GetBooks(ctx *gin.Context) {
	// connect to database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("ERROR : ", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("ERROR : ", err)
		return
	}

	fmt.Println("Successfully connected to database.")

	var books []Book

	sqlStatement := `SELECT * FROM books`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var book Book

		err = rows.Scan(&book.BookID, &book.Title, &book.Author, &book.Desc)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}

		books = append(books, book)
	}

	ctx.JSON(http.StatusOK, books)
}

func GetBook(ctx *gin.Context) {
	// connect to database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("ERROR : ", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("ERROR : ", err)
		return
	}

	fmt.Println("Successfully connected to database.")

	bookID := ctx.Param("bookID")

	var book Book

	sqlStatement := `SELECT * FROM books
	WHERE book_id = $1`

	bookIDInt, _ := strconv.Atoi(bookID)

	err := db.QueryRow(sqlStatement, bookIDInt).Scan(&book.BookID, &book.Title, &book.Author, &book.Desc)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

func DeleteBook(ctx *gin.Context) {
	// connect to database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("ERROR : ", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("ERROR : ", err)
		return
	}

	fmt.Println("Successfully connected to database.")

	bookID := ctx.Param("bookID")

	sqlStatement := `
	DELETE from books
	WHERE book_id = $1`

	bookIDInt, _ := strconv.Atoi(bookID)

	res, err := db.Exec(sqlStatement, bookIDInt)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("Book with id %v not found", bookID),
		})
		return
	}

	count, err := res.RowsAffected()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	if count == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("Book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Row(s) affected": count,
		"message":         fmt.Sprintf("Book with id %v has been successfully deleted", bookID),
	})
}
