package company

import (
	"context"

	"github.com/cbr4yan/trepot/core"
	"github.com/cbr4yan/trepot/pkg/tools"
)

func New(companyRepo core.CompanyRepo) *service {
	return &service{
		companyRepo: companyRepo,
	}
}

type service struct {
	companyRepo core.CompanyRepo
}

func (s *service) Create(ctx context.Context, in *core.CompanyRequest) (*core.CompanyResponse, error) {
	if err := tools.ValidateStruct(in); err != nil {
		return nil, err
	}

	model := &core.Company{
		Name:     in.Name,
		Email:    in.Email,
		Currency: in.Currency,
		Country:  in.Country,
		Active:   true,
	}

	if err := s.companyRepo.Create(ctx, model); err != nil {
		return nil, err
	}

	return core.NewResponse(model), nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	model, err := s.companyRepo.Find(ctx, id)
	if err != nil {
		return err
	}
	if err := s.companyRepo.Delete(ctx, model); err != nil {
		return err
	}
	return nil
}
