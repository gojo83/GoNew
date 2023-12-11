package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Book struct {
	Sr    int    `json:"sr"`
	Title string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{Sr: 1, Title: "Ulysses", Author: "James Joyce"},
	{Sr: 2, Title: "The Great Gatsby", Author: "F.Scott Fitzgerald"},
	{Sr: 3, Title: "David Copperfield", Author: "Charles Dickens"},
	{Sr: 4, Title: "Hamlet", Author: "William Shakespeare"},
}

func main() {

	e := echo.New()
	e.GET("/books", GetBooks)
	e.POST("/books", CreateBook)
	e.PUT("/books/:sr", UpdateBook)
	e.DELETE("/books/:sr", DeleteBook)
	e.Start(":8080")
}

func GetBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, books)
}

func CreateBook(c echo.Context) error {
	var newBook Book
	if err := c.Bind(&newBook); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	newBook.Sr = len(books) + 1
	books = append(books, newBook)

	return c.JSON(http.StatusCreated, newBook)
}

func UpdateBook(c echo.Context) error {
	sr := c.Param("sr")
	for i, book := range books {
		if bookSr := c.Param("sr"); bookSr == sr {
			var updatedBook Book
			if err := c.Bind(&updatedBook); err != nil {
				return c.JSON(http.StatusBadRequest, err.Error())
			}

			books[i] = updatedBook

			return c.JSON(http.StatusOK, updatedBook)
		}
	}
	return c.JSON(http.StatusNotFound, nil)
}

func DeleteBook(c echo.Context) error {
	sr := c.Param("sr")
	for i, book := range books {
		if bookSr := c.Param("sr"); bookSr == sr {
			books = append(books[:i], books[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, nil)
}
