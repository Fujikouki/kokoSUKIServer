package dao

import (
	"awesomeProject1/domain/object"
	"awesomeProject1/domain/repository"
	"context"
	"github.com/jmoiron/sqlx"
)

type (
	ChatMessage struct {
		db *sqlx.DB
	}
)

var _ repository.ChatMessageRepository = (*ChatMessage)(nil)

func NewChatMessage(db *sqlx.DB) *ChatMessage {
	return &ChatMessage{
		db: db,
	}
}

func (c *ChatMessage) Save(ctx context.Context, tx *sqlx.Tx, me *object.ChatMessage) error {

	_, err := c.db.Exec("INSERT INTO chat_messages (name, message, time) VALUES ($1, $2, $3)", me.Name, me.Message, me.Time)
	if err != nil {
		return err
	}

	return nil
}
