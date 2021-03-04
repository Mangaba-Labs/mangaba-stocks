package services

import (
	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/stock/repository"
)

// Service struct
type Service struct {
	Repository repository.StockRepository
}
