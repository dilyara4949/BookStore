package handlers

import (
	// "fmt"
	"net/http"

	"github.com/dilyara4949/BookStore/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
  return handler{db}
}

func (h *handler) GetBooks(c *gin.Context) {
	var book []models.Book 
	h.DB.Find(&book)
	c.IndentedJSON(http.StatusOK, &book)
}

func (h *handler) Home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Helllllllo")
}

func (h *handler) GetBook(c *gin.Context) {
	id := c.Param("id")
	readProduct := &models.Book{}
  h.DB.First(&readProduct, "id = ?", id)


	if (readProduct == &models.Book{} ){
		c.IndentedJSON(http.StatusOK, "Book not found")
	} else {
		c.IndentedJSON(http.StatusOK, &readProduct)
	}
}

func (h *handler) CreateBook(c *gin.Context) {
	
}

func (h *handler) UpdateBook(c *gin.Context) {
	
}

func (h *handler) DeleteBook(c *gin.Context) {
	
}