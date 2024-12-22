package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type CustomerHistory struct {
	bun.BaseModel `bun:"table:customer_history"`

	ID          uuid.UUID  `bun:"id,pk,type:UUID"`
	TableNumber int        `bun:"table_number,type:INTEGER,notnull"`
	Status      string     `bun:"status,type:TEXT,notnull"`
	TimeIn      time.Time  `bun:"time_in,type:TIMESTAMP,notnull"`
	TimeOut     *time.Time `bun:"time_out,type:TIMESTAMP"`
	PaymentTime *time.Time `bun:"payment_time,type:TIMESTAMP"`
	TotalPrice  *float64   `bun:"total_price,type:REAL"`
}
