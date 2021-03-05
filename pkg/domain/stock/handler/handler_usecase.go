package handler

import (
	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/stock/model"
	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/stock/services"
)

// Handler for user service
type Handler struct {
	service services.Service
}

// CreateShare handle the create share option in CLI
func (h *Handler) CreateShare(share model.Share) error {
	h.service.AddShare(share)
	return nil
}
