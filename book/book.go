package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rasad2k/go_fiber/database"
)

type Book struct {
	gorm.Model
	Title  string  `json:"name"`
	Year   int     `json:"year"`
	Author string  `json:"author"`
	Rating float64 `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book

	if err := db.Find(&book, id); err != nil {
		c.Status(404).Send("Book not found")
		return
	}

	c.JSON(book)
}

func NewBook(c *fiber.Ctx) {
	db := database.DBConn
	book := new(Book)

	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&book)
	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)

	if book.Title == "" {
		c.Status(500).Send("No Book Found with ID")
		return
	}

	db.Delete(&book)
	c.Send("Book Successfully deleted")
}
