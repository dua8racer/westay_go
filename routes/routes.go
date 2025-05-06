package routes

import (
	"westay-go/controllers"
	"westay-go/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Westay Rest API"})
	})

	r.POST("/login", controllers.Login)

	book := r.Group("/books")
	book.Use(middlewares.AuthMiddleware())
	{
		book.POST("/", controllers.CreateBook)
		book.GET("/", controllers.GetBooks)
		book.GET("/:id", controllers.GetBook)
		book.PUT("/:id", controllers.UpdateBook)
		book.DELETE("/:id", controllers.DeleteBook)
	}

	attendace := r.Group("attendance")
	attendace.Use(middlewares.AuthMiddleware())
	{
		attendace.POST("/", controllers.StoreAttendance)
	}
}
