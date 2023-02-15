package core

import (
	"context"
)

type CompanyRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Currency string `json:"currency" validate:"required,iso4217"`
	Country  string `json:"country" validate:"required,iso3166_1_alpha2"`
}

type CompanyResponse struct {
	BaseResponse
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Currency  string `json:"currency"`
	Country   string `json:"country"`
	Active    bool   `json:"active"`
	CreatedAt int64  `json:"created_at"`
}

func NewResponse(m *Company) *CompanyResponse {
	return &CompanyResponse{
		BaseResponse: BaseResponse{
			Object: "company",
		},
		ID:        m.ID,
		Name:      m.Name,
		Email:     m.Email,
		Currency:  m.Currency,
		Country:   m.Country,
		Active:    m.Active,
		CreatedAt: m.CreatedAt.Time().Unix(),
	}
}

type Company struct {
	BaseModel
	Name     string `db:"name"`
	Email    string `db:"email"`
	Currency string `db:"currency"`
	Country  string `db:"country"`
	Active   bool   `db:"active"`
}

func (c *Company) TableName() string {
	return "company"
}

type CompanyRepo interface {
	Find(ctx context.Context, id string) (*Company, error)
	Create(ctx context.Context, in *Company) error
	Delete(ctx context.Context, in *Company) error
}

type CompanyService interface {
	Create(ctx context.Context, in *CompanyRequest) (*CompanyResponse, error)
	Delete(ctx context.Context, id string) error
}
