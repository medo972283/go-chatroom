package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/medo972283/go-chatroom/models"
	"github.com/medo972283/go-chatroom/repositories"
)

func IndexChatrooms(c echo.Context) error {
	// Connect to db
	db, err := models.Connet()
	if err != nil {
		log.Fatal("db connect error: ", err)
	}
	defer db.Close()

	var chatrooms []models.Chatroom

	// Prepare sql statement
	stat, err := db.Prepare("SELECT * from chatrooms where rec_status = 1")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer stat.Close()

	// Do query
	rows, err := stat.Query()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	// Traverse all results
	for rows.Next() {
		chatroom := &models.Chatroom{}
		err = rows.Scan(&chatroom.ID, &chatroom.Name, &chatroom.CreateAt, &chatroom.CreateBy, &chatroom.UpdateAt, &chatroom.UpdateBy, &chatroom.DeleteAt, &chatroom.DeleteBy, &chatroom.RecStatus)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// get the room host's nickname
		user := &models.User{}
		err := db.QueryRow("SELECT nickname from users WHERE rec_status = 1 AND id = ?", (*chatroom).CreateBy).Scan(&user.Nickname)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		chatroom.CreateBy = user.Nickname

		chatrooms = append(chatrooms, *chatroom)
	}

	return c.JSONPretty(
		http.StatusOK,
		echo.Map{
			"code": 200,
			"data": chatrooms,
		}, "\t")
}

func ViewChatroom(c echo.Context) error {

	// Connect to db
	db, err := models.Connet()
	if err != nil {
		log.Fatal("db connect error: ", err)
	}
	defer db.Close()

	// Get path parameter
	roomID := c.Param("id")

	chatroom, err := repositories.GetChatroomByID(roomID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	user, err := repositories.GetUserByID(chatroom.CreateBy)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	participants, err := repositories.GetParticipantsByRoomID(roomID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	messages, err := repositories.GetMessagesByRoomID(roomID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSONPretty(
		http.StatusOK,
		echo.Map{
			"code": 200,
			"msg":  "",
			"data": echo.Map{
				"chatroom":     chatroom,
				"user":         user,
				"participatns": participants,
				"messages":     messages,
			},
		}, "\t")
}

func CreateChatroom(c echo.Context) error {

	// Connect to db
	db, err := models.Connet()
	if err != nil {
		log.Fatal("db connect error: ", err)
	}
	defer db.Close()

	// New a data model
	chatroom := new(models.Chatroom)

	// Bind body params to data model
	if err := c.Bind(chatroom); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Produce a v1 uuid
	uuid1, err := uuid.NewUUID()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Assign it
	chatroom.ID = uuid1

	// Prepare sql statement
	stat, err := db.Prepare("INSERT INTO chatrooms(id, name, created_by, rec_status) VALUES (?, ?, ?, ?)")

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer stat.Close()

	session, _ := session.Get("session", c)
	// Execute sql statement
	result, err := stat.Exec(chatroom.ID, chatroom.Name, session.Values["userID"], 1)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Get number of rows just affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSONPretty(
		http.StatusOK,
		echo.Map{
			"code": 200,
			"msg":  fmt.Sprintf("新增 %d 筆資料", rowsAffected),
		}, "\t")
}

func UpdateChatroom(c echo.Context) error {
	// Connect to db
	db, err := models.Connet()
	if err != nil {
		log.Fatal("db connect error: ", err)
	}
	defer db.Close()

	return nil
}

func DeleteChatroom(c echo.Context) error { // Connect to db
	db, err := models.Connet()
	if err != nil {
		log.Fatal("db connect error: ", err)
	}
	defer db.Close()

	return nil
}
