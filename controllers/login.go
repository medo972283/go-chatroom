package controllers

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	session, _ := session.Get("session", c)

	if session.Values["userID"] == nil {
		return c.Render(http.StatusFound, "login", "")
	} else {
		return c.Redirect(http.StatusFound, "/homepage")
	}

	// if val, ok := session.Values["email"].(string); ok {
	// 	// if val is a string
	// 	switch val {
	// 	case "":
	// 		http.Redirect(res, req, "html/login.html", http.StatusFound)
	// 	default:
	// 		http.Redirect(res, req, "html/home.html", http.StatusFound)
	// 	}
	// } else {
	// 	// if val is not a string type
	// 	http.Redirect(res, req, "html/login.html", http.StatusFound)
	// }
}
