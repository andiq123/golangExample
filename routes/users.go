package routes

import (
	"api/models"
	"api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User
	utils.BindOrRespondBadRequest(&user, context)

	err := user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not signup"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"error": "sign up success!"})
}

func signIn(context *gin.Context) {
	var user models.User
	utils.BindOrRespondBadRequest(&user, context)

	err := user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
}
