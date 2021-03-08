package services

import (
	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/stock/model"
	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/stock/repository"
)

// StockService interface
type StockService interface {
	AddShare(share model.Share) (model.Share, error)
	UpdateShares(shares []model.Share) ([]model.Share, error)
}

// NewStockService repository sqlite implementation
func NewStockService(repository repository.Repository) (service StockService) {
	service = &Service{
		Repository: repository,
	}
	return
}
