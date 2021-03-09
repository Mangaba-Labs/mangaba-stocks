package handler

import (
	"fmt"

	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/stock/model"
	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/stock/services"
)

// Handler for user service
type Handler struct {
	service services.Service
}

// CreateShare handle the create share option in CLI
func (h *Handler) CreateShare(share model.Share) error {
	_, err := h.service.AddShare(share)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// GetAll get all registers in database
func (h *Handler) GetAll() ([]model.Share, error) {
	shares, err := h.service.GetAll()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return shares, nil
}
