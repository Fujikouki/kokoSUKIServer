package repository

import (
	"awesomeProject1/domain/object"
	"context"
	"github.com/jmoiron/sqlx"
)

type ChatMessageRepository interface {
	Save(ctx context.Context, tx *sqlx.Tx, message *object.ChatMessage) error
}
