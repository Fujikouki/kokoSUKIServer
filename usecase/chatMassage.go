package usecase

import (
	"awesomeProject1/domain/object"
	"awesomeProject1/domain/repository"
	"context"
	"github.com/jmoiron/sqlx"
)

type ChatMessageU interface {
	Save(ctx context.Context, accountId, roomId int, message string) error
}

type ChatMessageR struct {
	db              *sqlx.DB
	chatMessageRepo repository.ChatMessageRepository
}

var _ ChatMessageU = (*ChatMessageR)(nil)

func NewChatMessageU(db *sqlx.DB, chatMessageRepo repository.ChatMessageRepository) *ChatMessageR {
	return &ChatMessageR{
		db:              db,
		chatMessageRepo: chatMessageRepo,
	}
}

func (c *ChatMessageR) Save(ctx context.Context, accountId, roomId int, message string) error {

	newMessage, err := object.NewChatMessage(accountId, roomId, message)

	if err != nil {
		return err
	}

	tx, err := c.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}()

	if err := c.chatMessageRepo.Save(ctx, tx, newMessage); err != nil {
		return err
	}

	return nil
}
