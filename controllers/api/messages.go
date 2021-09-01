package api

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/medo972283/go-chatroom/models"
	"github.com/medo972283/go-chatroom/repositories"
	"github.com/medo972283/go-chatroom/utils"
)

func IndexMessages(c echo.Context) error {
	return nil
}

func ViewMessage(c echo.Context) error {
	return nil
}

func CreateMessage(c echo.Context) error {

	session, _ := session.Get("session", c)

	message := new(models.Message)

	if err := c.Bind(message); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	message.ID = utils.ProduceV1UUID()
	message.UserID = utils.ConverseToUUID(session.Values["userID"])

	if err := repositories.CreateMessage(message); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSONPretty(
		http.StatusOK,
		echo.Map{
			"code": 201,
			"msg":  "Insert successfully",
			"data": message,
		}, "\t")
}

func UpdateMessage(c echo.Context) error {
	return nil
}

func DeleteMessage(c echo.Context) error {
	return nil
}
