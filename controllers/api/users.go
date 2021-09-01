package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/medo972283/go-chatroom/models"
	"github.com/medo972283/go-chatroom/utils"
)

func IndexUser(c echo.Context) error {
	return nil
}

func ViewUser(c echo.Context) error {
	return nil
}

func CreateUser(c echo.Context) error {
	// Connect to db
	db, err := models.Connet()
	if err != nil {
		log.Fatal("db connect error: ", err)
	}
	defer db.Close()

	user := new(models.User)

	// Bind body params to data model
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Produce a v1 uuid
	user.ID = utils.ProduceV1UUID()

	// Encrypt the password
	user.Password = utils.BcryptString(user.Password)

	// Prepare sql statement
	stat, err := db.Prepare("INSERT INTO users(id, account, password, nickname, email, rec_status) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer stat.Close()

	// Execute sql statement
	result, err := stat.Exec(user.ID, user.Account, user.Password, user.Nickname, user.Email, 1)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Get number of rows just affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSONPretty(
		http.StatusOK,
		echo.Map{
			"code": 200,
			"msg":  fmt.Sprintf("新增 %d 筆資料", rowsAffected),
			"data": user,
		}, "\t")
}

func UpdateUser(c echo.Context) error {
	return nil
}

func DeleteUser(c echo.Context) error {
	return nil
}
