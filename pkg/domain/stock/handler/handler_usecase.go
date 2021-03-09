package handler

import (
	"log"

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
		return err
	}
	return nil
}

// GetAll get all registers in database
func (h *Handler) GetAll() ([]model.Share, error) {
	shares, err := h.service.GetAll()

	if err != nil {
		return nil, err
	}

	return shares, nil
}

// UpdateAll used to handle the application startup
func (h *Handler) UpdateAll() error {
	shares, err := h.GetAll()

	if err != nil {
		log.Fatalln(err)
		return err
	}

	shares, err = h.service.UpdateShares(shares)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

// Remove a share from database
func (h *Handler) Remove(id int) error {
	err := h.service.Remove(id)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}
