package routes

import (
	"api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "could not parse eventId"})
		return
	}

	_, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	var registration models.Registration
	registration.Event_id = eventId
	registration.User_id = userId

	err = registration.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not save the registration"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "registration added succesfully"})
}
