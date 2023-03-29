package handlers

import (
	"errors"
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
	

	var ord models.Read
	if err := c.BindJSON(&ord); err != nil {
		c.IndentedJSON(http.StatusOK, "Input is not correct")
		panic(err)
	} else {
		var book []models.Book
		if ord.Ord != "" && ord.Ord == "desc"{ 
			h.DB.Order("id desc").Find(&book)
		} else { 
			h.DB.Order("id asc").Find(&book)
		}
		c.IndentedJSON(http.StatusOK, &book)
		// c.IndentedJSON(http.StatusOK, ord)
	}
}

func (h *handler) Home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Helllllllo")
}

func (h *handler) GetBook(c *gin.Context) {
	id := c.Param("id")
	readProduct := models.Book{}

	dbRresult := h.DB.Where("title = ?", id).First(&readProduct)
	if errors.Is(dbRresult.Error, gorm.ErrRecordNotFound) {
		if dbRresult = h.DB.Where("id = ?", id).First(&readProduct); dbRresult.Error != nil {
			c.IndentedJSON(http.StatusOK, "Book not found")
		} else {
			c.IndentedJSON(http.StatusOK, &readProduct)
		}
	} else {
		c.IndentedJSON(http.StatusOK, &readProduct)
	}

}

func (h *handler) CreateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusOK, "Input is not correct")
		panic(err)
	} else {
		h.DB.Create(&newBook)
		c.IndentedJSON(http.StatusOK, newBook)
	}
}

func (h *handler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	readProduct := &models.Book{}

	if dbRresult := h.DB.Where("id = ?", id).First(&readProduct); dbRresult.Error != nil {
		c.IndentedJSON(http.StatusOK, "Book not found")
	} else {

		var newBook models.Book
		if err := c.BindJSON(&newBook); err != nil {
			c.IndentedJSON(http.StatusOK, "Input is not correct")
			panic(err)
		} else {
			if newBook.Cost != 0{
				readProduct.Cost = newBook.Cost
			}
			if newBook.Title != ""{
				readProduct.Title = newBook.Title
			}
			if newBook.Description != ""{
				readProduct.Description = newBook.Description
			}
			h.DB.Save(readProduct)
			c.IndentedJSON(http.StatusOK, readProduct)
		}
	}

}

func (h *handler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var deleteBook models.Book

	if dbRresult := h.DB.Where("id = ?", id).First(&deleteBook); dbRresult.Error != nil {
		c.IndentedJSON(http.StatusOK, "Book not found")
	} else {
		h.DB.Where("id = ?", id).Delete(&deleteBook)
		c.IndentedJSON(http.StatusOK, "Book deleted")
	}
}
