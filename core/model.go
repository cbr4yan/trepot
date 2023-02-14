package core

import (
	"github.com/cbr4yan/trepot/pkg/types"
	"github.com/rs/xid"
)

type Model interface {
	TableName() string
	GetID() string
	SetID(string)
	GetCreatedAt() types.DateTime
	GetUpdatedAt() types.DateTime
	GetDeletedAt() types.DateTime
	RefreshID()
	RefreshCreatedAt()
	RefreshUpdatedAt()
	RefreshDeletedAt()
}

type BaseModel struct {
	ID        string         `db:"id"`
	CreatedAt types.DateTime `db:"created_at"`
	UpdatedAt types.DateTime `db:"updated_at"`
	DeletedAt types.DateTime `db:"deleted_at"`
}

func (m *BaseModel) GetID() string {
	return m.ID
}

func (m *BaseModel) SetID(id string) {
	m.ID = id
}

func (m *BaseModel) GetCreatedAt() types.DateTime {
	return m.CreatedAt
}

func (m *BaseModel) GetUpdatedAt() types.DateTime {
	return m.UpdatedAt
}

func (m *BaseModel) GetDeletedAt() types.DateTime {
	return m.DeletedAt
}

func (m *BaseModel) RefreshID() {
	m.ID = xid.New().String()
}

func (m *BaseModel) RefreshCreatedAt() {
	m.CreatedAt = types.NowDateTime()
}

func (m *BaseModel) RefreshUpdatedAt() {
	m.UpdatedAt = types.NowDateTime()
}

func (m *BaseModel) RefreshDeletedAt() {
	m.DeletedAt = types.NowDateTime()
}
