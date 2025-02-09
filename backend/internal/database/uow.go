package database

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
)

const tx_context = "tx"

func UowSelect[TModel any](ctx context.Context, db *bun.DB) (*TModel, *bun.SelectQuery) {
	model := new(TModel)
	query := db.NewSelect()
	if tx, ok := ctx.Value(tx_context).(*bun.Tx); ok {
		query = tx.NewSelect()
	}

	return model, query.Model(model)
}

func UowUpdate[TModel any](ctx context.Context, db *bun.DB, model *TModel) *bun.UpdateQuery {
	query := db.NewUpdate()
	if tx, ok := ctx.Value(tx_context).(*bun.Tx); ok {
		query = tx.NewUpdate()
	}

	return query.Model(model)
}

func UowInsert[TModel any](ctx context.Context, db *bun.DB, model *TModel) *bun.InsertQuery {
	query := db.NewInsert()
	if tx, ok := ctx.Value(tx_context).(*bun.Tx); ok {
		query = tx.NewInsert()
	}

	return query.Model(model)
}

func UowDelete[TModel any](ctx context.Context, db *bun.DB) (*TModel, *bun.DeleteQuery) {
	model := new(TModel)
	query := db.NewDelete()
	if tx, ok := ctx.Value(tx_context).(*bun.Tx); ok {
		query = tx.NewDelete()
	}

	return model, query.Model(model)
}

func Uow(ctx context.Context, db *bun.DB, callback func(context.Context, *bun.Tx) error) error {
	return db.RunInTx(ctx, &sql.TxOptions{}, func(callbackContext context.Context, tx bun.Tx) error {
		txContext := context.WithValue(callbackContext, tx_context, &tx)
		return callback(txContext, &tx)
	})
}
