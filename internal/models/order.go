package models

import "github.com/uptrace/bun"

type order_logs struct {
	bun.BaseModel `bun:"table:order_logs"`

	ID       string `bun:"id,pk,type:UUID"`
	Food     string `bun:"food,type:TEXT"`
	Table_Id string `bun:"table_id,type:UUID"`
	Price    int64  `bun:"price,type:INTEGER"`
}
