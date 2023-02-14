package common

import (
	"context"
	"errors"
	"fmt"

	"github.com/cbr4yan/trepot/core"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *Dao {
	return &Dao{
		db: db,
	}
}

type Dao struct {
	db *sqlx.DB
}

func (d *Dao) Create(ctx context.Context, query string, in core.Model) error {
	if in.GetID() == "" {
		in.RefreshID()
	}

	if in.GetCreatedAt().IsZero() {
		in.RefreshCreatedAt()
	}

	if in.GetUpdatedAt().IsZero() {
		in.RefreshUpdatedAt()
	}

	_, err := d.db.NamedExecContext(ctx, query, in)
	return err
}

func (d *Dao) Update(ctx context.Context, query string, in core.Model) error {
	if in.GetID() == "" {
		return errors.New("id is not set")
	}

	if in.GetCreatedAt().IsZero() {
		in.RefreshCreatedAt()
	}

	in.RefreshUpdatedAt()

	_, err := d.db.NamedExecContext(ctx, query, in)
	return err
}

func (d *Dao) Delete(ctx context.Context, in core.Model) error {
	if in.GetID() == "" {
		return errors.New("id is not set")
	}

	in.RefreshDeletedAt()

	query := fmt.Sprintf("UPDATE %s SET deleted_at = $1 WHERE id = $2", in.TableName())
	_, err := d.db.ExecContext(ctx, query, in.GetDeletedAt(), in.GetID())
	return err
}

func (d *Dao) FindByID(ctx context.Context, m core.Model, id string) error {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1 AND deleted_at IS NULL", m.TableName())
	return d.db.GetContext(ctx, m, query, id)
}
