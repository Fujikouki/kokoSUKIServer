package object

import "time"

type ChatMessage struct {
	ID      int       `db:"id"`
	Name    string    `db:"name"`
	Message string    `db:"message"`
	Time    time.Time `db:"time"`
}

func NewChatMessage(name, message string) (*ChatMessage, error) {

	newMassage := &ChatMessage{
		Name:    name,
		Message: message,
		Time:    time.Now(),
	}

	return newMassage, nil

}
