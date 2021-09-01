package service

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func CheckSignin(c echo.Context) bool {
	session, _ := session.Get("session", c)

	if session.Values["userID"] == nil {
		return false
	}

	return true
}
