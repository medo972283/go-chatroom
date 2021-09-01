package repositories

import (
	"log"

	"github.com/medo972283/go-chatroom/models"
)

type UserMessage struct {
	UserName string
	CreateAt string
	Content  string
}

func GetMessageByID(id string) (*UserMessage, error) {

	// Connect to db
	db, err := models.Connet()
	if err != nil {
		log.Fatal("db connect error: ", err)
	}
	defer db.Close()

	message := new(UserMessage)

	err = db.QueryRow(`
		SELECT u.nickname, m.content, m.created_at
		FROM messages as m
		LEFT JOIN users as u
		ON m.user_id = u.id
		WHERE m.rec_status = 1 AND u.rec_status = 1 AND m.id = ?
		`, id).Scan(&message.UserName, &message.Content, &message.CreateAt)
	if err != nil {
		return &UserMessage{}, err
	}

	return message, nil
}

func GetMessagesByRoomID(roomID string) ([]UserMessage, error) {

	// Connect to db
	db, err := models.Connet()
	if err != nil {
		log.Fatal("db connect error: ", err)
	}
	defer db.Close()

	var messages []UserMessage

	// Prepare sql statement
	stat, err := db.Prepare(`
		SELECT m.content, u.nickname, m.created_at
		FROM messages as m
		LEFT JOIN users as u
		ON m.user_id = u.id
		WHERE m.rec_status = 1 AND u.rec_status = 1 AND m.room_id = ?
		ORDER BY m.created_at`)
	if err != nil {
		return nil, err
	}
	defer stat.Close()

	// Execute query
	rows, err := stat.Query(roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		message := new(UserMessage)

		err = rows.Scan(&message.Content, &message.UserName, &message.CreateAt)
		if err != nil {
			return nil, err
		}

		messages = append(messages, *message)
	}

	return messages, nil
}

func CreateMessage(message *models.Message) error {

	// Connect to db
	db, err := models.Connet()
	if err != nil {
		log.Fatal("db connect error: ", err)
	}
	defer db.Close()

	// Prepare sql statement
	stat, err := db.Prepare("INSERT IGNORE INTO messages(id, content, user_id, room_id, created_by, rec_status) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stat.Close()

	// Execute sql statement
	_, err = stat.Exec(message.ID, message.Content, message.UserID, message.RoomID, message.UserID, 1)
	if err != nil {
		return err
	}

	return nil
}
