package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Menu struct {
	bun.BaseModel `bun:"table:menu"`

	ID              uuid.UUID `bun:"id,pk,type:UUID"`
	FoodName        string    `bun:"food_name,type:TEXT,notnull"`
	FoodPrice       float64   `bun:"food_price,type:REAL,notnull"`
	FoodDescription *string   `bun:"food_description,type:TEXT"`
	FoodImageURL    *string   `bun:"food_image_url,type:TEXT"`
}
