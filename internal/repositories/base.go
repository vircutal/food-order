package repositories

import (
	"context"

	"github.com/uptrace/bun"
)

type BaseDB[T any] struct {
	db *bun.DB
}

func (b *BaseDB[T]) AddOne(ctx context.Context, model *T) error {
	_, err := b.db.NewInsert().Model(model).Exec(ctx)
	return err
}

func (b *BaseDB[T]) BulkAdd(ctx context.Context, models []*T) error {
	_, err := b.db.NewInsert().Model(models).Exec(ctx)
	return err
}
