package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type OrderLog struct {
	bun.BaseModel `bun:"table:order_log"`

	ID               uuid.UUID `bun:"id,pk,type:UUID"`
	CustomerID       uuid.UUID `bun:"customer_id,type:UUID"`
	MenuItemID       uuid.UUID `bun:"menu_item_id,type:UUID"`
	MenuItemPrice    float64   `bun:"menu_item_price,type:REAL,notnull"`
	Quantity         int       `bun:"quantity,type:INTEGER,notnull"`
	OrderDescription *string   `bun:"order_description,type:TEXT"`
	OrderedTime      time.Time `bun:"ordered_time,type:TIMESTAMP,notnull"`
}
