package repositories

import (
	"context"
	"food-order/internal/utils"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type BaseDB[T any] struct {
	db *bun.DB
}

func GetBaseDB[T any]() *BaseDB[T] {
	return &BaseDB[T]{
		db: utils.InitDB(),
	}
}

// func (b *BaseDB[T]) BeginTransaction() (*bun.Tx, error) {
// 	tx, err := b.db.Begin()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to begin transaction: %w", err)
// 	}
// 	return tx, nil
// }

func (b *BaseDB[T]) AddOne(ctx context.Context, model *T) error {
	_, err := b.db.NewInsert().Model(model).Exec(ctx)
	return err
}

func (b *BaseDB[T]) BulkAdd(ctx context.Context, models []*T) error {
	_, err := b.db.NewInsert().Model(models).Exec(ctx)
	return err
}

func (b *BaseDB[T]) UpdateOne(ctx context.Context, model *T) error {
	_, err := b.db.NewUpdate().Model(model).WherePK().Exec(ctx)
	return err
}

func (b *BaseDB[T]) FindOneById(ctx context.Context, id uuid.UUID) (*T, error) {
	var model T
	err := b.db.NewSelect().Model(&model).Where("id = ?", id).Scan(ctx)
	return &model, err
}

func (b *BaseDB[T]) DeleteOneById(ctx context.Context, id uuid.UUID) error {
	var model T
	_, err := b.db.NewDelete().Model(&model).Where("id = ?", id).Exec(ctx)
	return err
}

func (b *BaseDB[T]) CheckExistByID(ctx context.Context, id uuid.UUID) bool {
	var model T
	err := b.db.NewSelect().Model(&model).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return false
	}
	return true
}
