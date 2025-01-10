package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Restaurant struct {
	bun.BaseModel `bun:"table:restaurant"`

	ID             uuid.UUID `bun:"id,pk,type:UUID"`
	RestaurantName string    `bun:"restaurant_name,notnull"`
	Branch         string    `bun:"branch,notnull"`
}
