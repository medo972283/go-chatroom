package repositories

import (
	"log"

	"github.com/medo972283/go-chatroom/models"
)

type ChatroomParticipant struct {
	ID     string
	UserID string
	RoomID string
	Name   string
}

func GetParticipantsByRoomID(roomID string) ([]ChatroomParticipant, error) {

	// Connect to db
	db, err := models.Connet()
	if err != nil {
		log.Fatal("db connect error: ", err)
	}
	defer db.Close()

	var participants []ChatroomParticipant

	// Prepare sql statement
	stat, err := db.Prepare(`SELECT p.id, p.user_id, p.room_id, u.nickname 
		FROM participants as p 
		LEFT JOIN users as u 
		ON p.user_id = u.id 
		WHERE u.rec_status = 1 AND p.rec_status = 1 AND room_id = ?
		ORDER BY p.created_at`)
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

	// Traverse all results
	for rows.Next() {
		participant := new(ChatroomParticipant)

		err = rows.Scan(&participant.ID, &participant.UserID, &participant.RoomID, &participant.Name)
		if err != nil {
			return nil, err
		}

		participants = append(participants, *participant)
	}

	return participants, nil
}

func CreateParticipant(participant *models.Participant) error {

	// Connect to db
	db, err := models.Connet()
	if err != nil {
		log.Fatal("db connect error: ", err)
	}
	defer db.Close()

	// Prepare sql statement
	stat, err := db.Prepare("INSERT IGNORE INTO participants(id, user_id, room_id, created_by, rec_status) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stat.Close()

	// Execute sql statement
	_, err = stat.Exec(participant.ID, participant.UserID, participant.RoomID, participant.UserID, 1)
	if err != nil {
		return err
	}

	return nil
}
