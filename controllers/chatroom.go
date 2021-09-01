package controllers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/medo972283/go-chatroom/models"
	"github.com/medo972283/go-chatroom/repositories"
	"github.com/medo972283/go-chatroom/service"
	"github.com/medo972283/go-chatroom/utils"
)

func Chatroom(c echo.Context) error {
	if !service.CheckSignin(c) {
		return c.Redirect(http.StatusFound, "/login")
	}

	session, _ := session.Get("session", c)

	// New a participant
	participant := new(models.Participant)
	// Produce a v1 uuid
	participant.ID = utils.ProduceV1UUID()
	// Fill in the current user ID & target room ID
	participant.UserID, _ = uuid.Parse((fmt.Sprintf("%v", session.Values["userID"])))
	participant.RoomID, _ = uuid.Parse(c.FormValue("RoomID"))

	// Insert current user into the room participant
	err := repositories.CreateParticipant(participant)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.Render(http.StatusOK, "chatroom", echo.Map{
		"UserNickname": session.Values["userNickname"],
		"RoomID":       c.FormValue("RoomID"),
	})
}
