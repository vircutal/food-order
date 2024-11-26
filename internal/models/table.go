package models

import (
	"github.com/uptrace/bun"
)

type table struct {
	bun.BaseModel `bun:"table:table"`

	ID           string `bun:"id,pk,type:UUID"`
	Table_Number int8   `bun:"table_number,type:INTEGER"`
	Status       string `bun:"status,type:TEXT"`
}
