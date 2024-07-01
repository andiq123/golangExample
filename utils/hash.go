package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func BindOrRespondBadRequest(entity any, c *gin.Context) any {
	err := c.ShouldBindJSON(&entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse json"})
		return nil
	}
	return entity
}
