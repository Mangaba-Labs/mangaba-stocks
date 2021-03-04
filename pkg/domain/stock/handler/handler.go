package handler

import "github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/stock/services"

// StockHandler interface
type StockHandler interface{}

// NewStockHandler pointer to handler implementation
func NewStockHandler(s services.StockService) Handler {
	return Handler{
		service: s,
	}
}
