package company

import (
	"context"

	"github.com/cbr4yan/trepot/core"
	"github.com/cbr4yan/trepot/repo/common"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *repo {
	return &repo{
		dao: common.New(db),
	}
}

type repo struct {
	dao *common.Dao
}

func (r *repo) Create(ctx context.Context, in *core.Company) error {
	query := `INSERT INTO company (id, name, email, currency, country, active, created_at, updated_at) VALUES (:id, :name, :email, :currency, :country, :active, :created_at, :updated_at)`
	return r.dao.Create(ctx, query, in)
}

func (r *repo) Update(ctx context.Context, in *core.Company) error {
	query := `UPDATE company SET name=:name, email=:email, currency=:currency, country=:country, active=:active, updated_at=:updated_at WHERE id=:id`
	return r.dao.Update(ctx, query, in)
}

func (r *repo) Delete(ctx context.Context, in *core.Company) error {
	return r.dao.Delete(ctx, in)
}

func (r *repo) Find(ctx context.Context, id string) (*core.Company, error) {
	out := &core.Company{}
	if err := r.dao.FindByID(ctx, out, id); err != nil {
		return nil, err
	}
	return out, nil
}
