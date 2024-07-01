package routes

import (
	"api/models"
	"api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid eventID"})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events"})
		return
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	utils.BindOrRespondBadRequest(&event, context)

	event.UserID = context.GetInt64("userId")

	err := event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not create event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "added succesfully"})
}

func updateEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	eventFound, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	loggedUserId := context.GetInt64("userId")
	if loggedUserId != eventFound.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorised"})
		return
	}

	var event models.Event
	utils.BindOrRespondBadRequest(&event, context)

	event.ID = eventID
	err = event.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event updated succesfully"})
}

func deleteEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	eventFound, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	loggedUserId := context.GetInt64("userId")
	if loggedUserId != eventFound.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorised"})
		return
	}

	err = models.DeleteEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete the event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event deleted succesfully"})
}
