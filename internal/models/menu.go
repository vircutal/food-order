package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Menu struct {
	bun.BaseModel `bun:"table:menu"`

	ID           uuid.UUID `bun:"id,pk,type:UUID"`
	MenuName     string    `bun:"menu_name,type:text"`
	RestaurantID uuid.UUID `bun:"restaurant_id,type:UUID"`
}
