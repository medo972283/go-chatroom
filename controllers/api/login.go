package api

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/medo972283/go-chatroom/models"
	"golang.org/x/crypto/bcrypt"
)

// Log in action
func Login(c echo.Context) error {
	// Connect to db
	db, err := models.Connet()
	if err != nil {
		log.Fatal("db connect error: ", err)
	}
	defer db.Close()

	user := new(models.User)
	storedUser := new(models.User)

	// Bind body params to data model
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = db.QueryRow("SELECT id, account, password, email, nickname from users WHERE rec_status = 1 AND (email = ? OR account = ?)", user.Email, user.Account).Scan(&storedUser.ID, &storedUser.Account, &storedUser.Password, &storedUser.Email, &storedUser.Nickname)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(storedUser.Password),
		[]byte(user.Password))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	session, _ := session.Get("session", c)
	session.Options = &sessions.Options{
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   3600 * 8,
		HttpOnly: true,
	}

	session.Values["userID"] = storedUser.ID.String()
	session.Values["userAccount"] = storedUser.Account
	session.Values["userEmail"] = storedUser.Email
	session.Values["userNickname"] = storedUser.Nickname

	err = session.Save(c.Request(), c.Response())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSONPretty(
		http.StatusOK,
		echo.Map{
			"code": 200,
			"msg":  "Login success",
		}, "\t")
}

// Log out action
func Logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}

	sess.Values["userID"] = ""
	sess.Values["userAccount"] = ""
	sess.Values["userEmail"] = ""
	sess.Values["userNickname"] = ""

	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusSeeOther, "/")
}
