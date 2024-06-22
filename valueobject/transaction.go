package valueobject

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Amount    int
	From      uuid.UUID
	to        uuid.UUID
	CreatedAt time.Time
}
