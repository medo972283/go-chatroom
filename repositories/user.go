package repositories

import (
	"log"

	"github.com/medo972283/go-chatroom/models"
)

// Get user info by user ID
func GetUserByID(userID string) (*models.User, error) {

	// Connect to db
	db, err := models.Connet()
	if err != nil {
		log.Fatal("db connect error: ", err)
	}
	defer db.Close()

	user := new(models.User)

	err = db.QueryRow("SELECT id, account, nickname, email, created_by FROM users WHERE rec_status = 1 AND id = ?", userID).
		Scan(&user.ID, &user.Account, &user.Nickname, &user.Email, &user.CreateBy)
	if err != nil {
		return nil, err
	}

	return user, nil
}
