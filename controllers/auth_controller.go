package controllers

import (
	"log"
	"net/http"
	"os"
	"time"
	"westay-go/config"
	"westay-go/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey []byte

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variable.")
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Println("JWT_SECRET not set in environment")
	}

	jwtKey = []byte(secret)
}

func Login(c *gin.Context) {
	var input models.User
	var user models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.UserError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	config.DB.Where("username = ?", input.Username).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized,
			models.UserError{
				Status:  http.StatusUnauthorized,
				Message: "Invalid credentials :p",
			})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid credentials :p"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"user_name": user.Username,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString(jwtKey)
	c.JSON(http.StatusOK, models.UserLogin{
		Status:  http.StatusOK,
		Message: "Login Success",
		Token:   tokenString,
		User: models.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
