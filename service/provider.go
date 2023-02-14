package service

import (
	"github.com/cbr4yan/trepot/core"
	"github.com/cbr4yan/trepot/repo"
	"github.com/cbr4yan/trepot/service/company"
)

func New(
	repoProvider *repo.Provider,
) *Provider {
	return &Provider{
		CompanyService: company.New(repoProvider.CompanyRepo),
	}
}

type Provider struct {
	CompanyService core.CompanyService
}
