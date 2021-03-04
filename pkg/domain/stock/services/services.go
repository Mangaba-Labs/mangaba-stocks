package services

import (
	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/stock/repository"
)

// StockService interface
type StockService interface {
}

// NewStockService repository sqlite implementation
func NewStockService(repository repository.Repository) (service StockService) {
	service = &Service{
		Repository: repository,
	}

	return
}
