package api

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/medo972283/go-chatroom/models"
	"github.com/medo972283/go-chatroom/repositories"
	"github.com/medo972283/go-chatroom/utils"
)

func IndexParticipants(c echo.Context) error {
	return nil
}

func ViewParticipant(c echo.Context) error {
	return nil
}

func CreateParticipant(c echo.Context) error {

	session, _ := session.Get("session", c)

	participant := new(models.Participant)

	if err := c.Bind(participant); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Produce a v1 uuid
	participant.ID = utils.ProduceV1UUID()
	// Fill in the current user ID
	participant.UserID, _ = uuid.Parse((fmt.Sprintf("%v", session.Values["userID"])))

	err := repositories.CreateParticipant(participant)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSONPretty(
		http.StatusOK,
		echo.Map{
			"code": 201,
			"msg":  "Insert successfully",
			"data": participant,
		}, "\t")
}

func UpdateParticipant(c echo.Context) error {
	return nil
}

func DeleteParticipant(c echo.Context) error {
	return nil
}
