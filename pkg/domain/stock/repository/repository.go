package repository

import (
	"database/sql"

	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/stock/model"
)

// StockRepository Contract
type StockRepository interface {
	GetAll() ([]model.Share, error)
	InsertShare(share model.Share) (model.Share, error)
	RemoveShare(id int) error
	UpdateValue(share model.Share) error
}

// NewStockRepository repository sqlite implementation
func NewStockRepository(db *sql.DB) StockRepository {
	return &Repository{}
}
