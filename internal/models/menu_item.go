package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MenuItem struct {
	bun.BaseModel `bun:"table:menu_item"`

	ID                  uuid.UUID `bun:"id,pk,type:UUID"`
	MenuID              uuid.UUID `bun:"menu_id,type:UUID"`
	MenuItemName        string    `bun:"menu_item_name,type:TEXT,notnull"`
	MenuItemPrice       float64   `bun:"menu_item_price,type:REAL,notnull"`
	MenuItemDescription *string   `bun:"menu_item_description,type:TEXT"`
	MenuItemImageKey    *string   `bun:"menu_item_image_url,type:TEXT"`
}
