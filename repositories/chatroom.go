package repositories

import (
	"log"

	"github.com/medo972283/go-chatroom/models"
)

// Get chatroom info by room ID
func GetChatroomByID(roomID string) (*models.Chatroom, error) {

	// Connect to db
	db, err := models.Connet()
	if err != nil {
		log.Fatal("db connect error: ", err)
	}
	defer db.Close()

	chatroom := new(models.Chatroom)

	err = db.QueryRow("SELECT id, name, created_by FROM chatrooms WHERE rec_status = 1 AND id = ?", roomID).Scan(&chatroom.ID, &chatroom.Name, &chatroom.CreateBy)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}
