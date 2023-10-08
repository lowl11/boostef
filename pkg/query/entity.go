package query

import (
	"github.com/google/uuid"
	"time"
)

type Entity struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
