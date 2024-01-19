package router

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/events-app/models"
	"example.com/events-app/utils"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {

	events, err := models.GetAllEvents()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get all events",
			"errors":  []string{err.Error()},
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Fetched all events successfully",
		"data": gin.H{
			"events": events,
		},
	})
}

func createEvent(context *gin.Context) {

	var event *models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	user_claims, _ := context.Keys["user_claims"].(*utils.UserClaims)

	event.UserId = user_claims.UserId
	err = event.Save()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"message": "Event created successfully",
		"event":   event,
	})
}

func getEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"message": "Event fetched successfully",
		"event":   event,
	})
}

func updateEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	user_claims, _ := context.Keys["user_claims"].(*utils.UserClaims)
	user_id := user_claims.UserId

	existing_event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "No such event exists",
		})
		return
	}

	if existing_event.UserId != user_id {
		context.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Not authorized to update the event",
		})
		return
	}

	var event *models.Event
	err = context.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "unable to parse the data",
		})
		return
	}

	event.Id = eventId
	err = models.UpdateEventById(eventId, user_id, event)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "event updated successfully",
	})

}

func deleteEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	user_claims, _ := context.Keys["user_claims"].(*utils.UserClaims)
	user_id := user_claims.UserId

	existing_event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "No such event exists",
		})
		return
	}

	if existing_event.UserId != user_id {
		context.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Not authorized to delete the event",
		})
		return
	}

	err = models.DeleteEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "event deleted successfully",
	})

}

func getAllRegistrationsForEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	user_claims, _ := context.Keys["user_claims"].(*utils.UserClaims)
	user_id := user_claims.UserId

	existing_event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "No such event exists",
		})
		return
	}

	if existing_event.UserId != user_id {
		context.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Not authorized to fetch all the event registrations",
		})
		return
	}

	err = existing_event.GetAllRegistrations()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":    401,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Fetched all registartions successfully",
		"event":   existing_event,
	})
}

func registerForEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	user_claims, _ := context.Keys["user_claims"].(*utils.UserClaims)
	user_id := user_claims.UserId

	existing_event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "No such event exists",
		})
		return
	}

	registration_id, err := existing_event.RegisterUser(user_id)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"code":            201,
		"message":         "Registered successfully",
		"registration_id": registration_id,
	})

}

func removeRegistrationForEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	user_claims, _ := context.Keys["user_claims"].(*utils.UserClaims)
	user_id := user_claims.UserId

	existing_event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "No such event exists",
		})
		return
	}

	err = existing_event.DeleteUserRegistration(user_id)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Registration deleted successfully",
	})
}
