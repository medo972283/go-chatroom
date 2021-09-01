// package controllers
package controllers

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/medo972283/go-chatroom/service"
)

func Homepage(c echo.Context) error {

	if !service.CheckSignin(c) {
		return c.Redirect(http.StatusFound, "/login")
	}

	sess, _ := session.Get("session", c)

	return c.Render(http.StatusOK, "homepage", echo.Map{
		"UserNickname": sess.Values["userNickname"],
	})
}
