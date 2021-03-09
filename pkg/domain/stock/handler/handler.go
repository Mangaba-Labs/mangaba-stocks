package handler

import (
	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/stock/model"
	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/stock/services"
)

// StockHandler interface
type StockHandler interface {
	CreateShare(share model.Share) error
	GetAll() ([]model.Share, error)
}

// NewStockHandler pointer to handler implementation
func NewStockHandler(s services.Service) Handler {
	return Handler{
		service: s,
	}
}
