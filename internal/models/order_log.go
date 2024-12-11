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
	FoodID           uuid.UUID `bun:"food_id,type:UUID"`
	FoodPrice        float64   `bun:"food_price,type:REAL,notnull"`
	Quantity         int64     `bun:"quantity,type:INTEGER,notnull"`
	OrderDescription string    `bun:"order_description,type:TEXT"`
	OrderedTime      time.Time `bun:"ordered_time,type:TIMESTAMP,notnull"`
}
