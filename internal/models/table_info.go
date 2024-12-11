package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type TableInfo struct {
	bun.BaseModel `bun:"table:table_info"`

	ID          uuid.UUID `bun:"id,pk,type:UUID"`
	TableNumber int16     `bun:"table_number,notnull"`
	Status      string    `bun:"status,type:text,notnull"`
}
