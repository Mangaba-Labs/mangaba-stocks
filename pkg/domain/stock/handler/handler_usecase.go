package handler

import "github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/stock/services"

// Handler for user service
type Handler struct {
	service services.StockService
}
