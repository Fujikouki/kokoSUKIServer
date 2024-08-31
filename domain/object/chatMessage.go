package object

import "time"

type ChatMessage struct {
	ID        int       `db:"id"`
	AccountId int       `db:"account_id"`
	RoomId    int       `db:"room_id"`
	Message   string    `db:"message"`
	Time      time.Time `db:"time"`
}

func NewChatMessage(accountId, roomId int, message string) (*ChatMessage, error) {

	newMassage := &ChatMessage{
		AccountId: accountId,
		RoomId:    roomId,
		Message:   message,
		Time:      time.Now(),
	}

	return newMassage, nil

}
