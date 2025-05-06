package controllers

import (
	"net/http"
	"westay-go/config"
	"westay-go/models"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&book)
	c.JSON(http.StatusOK, book)
}

func GetBooks(c *gin.Context) {
	var books []models.Book
	var response models.Response
	// userName, _ := c.Get("userName")
	userAll, _ := c.Get("user")
	// log.Println(userAll.(models.UserResponse).Username)
	config.DB.Find(&books)
	response.Status = http.StatusOK
	response.Data = books
	response.Message = userAll.(models.UserResponse).Username
	c.JSON(http.StatusOK, response)
}

func GetBook(c *gin.Context) {
	var book models.Book
	var response models.Response
	if err := config.DB.First(&book, "id = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	response.Status = http.StatusOK
	response.Data = book
	c.JSON(http.StatusOK, response)
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	if err := config.DB.First(&book, "id = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.ShouldBindJSON(&book)
	config.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := config.DB.First(&book, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found  "})
		return
	}
	config.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
