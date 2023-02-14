package repo

import (
	"github.com/cbr4yan/trepot/core"
	"github.com/cbr4yan/trepot/repo/company"
	"github.com/jmoiron/sqlx"
)

func New(
	db *sqlx.DB,
) *Provider {
	return &Provider{
		CompanyRepo: company.New(db),
	}
}

type Provider struct {
	CompanyRepo core.CompanyRepo
}
