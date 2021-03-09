package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/stock/model"
	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/stock/repository"
)

// Service struct
type Service struct {
	Repository repository.Repository
}

// AddShare take the actual value from the Share and send a share to repo
func (s *Service) AddShare(share model.Share) (model.Share, error) {
	hgURI := fmt.Sprintf("https://api.hgbrasil.com/finance/stock_price?key=%s&symbol=%s", os.Getenv("HG_KEY"), share.Symbol)

	response, err := http.Get(hgURI)

	if err != nil {
		return share, err
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return share, err
	}

	r := &model.ResponseHG{}

	if err = json.Unmarshal(bodyBytes, &r); err != nil {
		return share, err
	}

	share.ParseResponse(r, r.Results[share.Symbol].Symbol)
	s.Repository.InsertShare(share)
	return share, nil
}

// UpdateShares of all shares registered in our database
func (s *Service) UpdateShares(shares []model.Share) ([]model.Share, error) {
	for i := 0; i < len(shares); i++ {
		hgURI := fmt.Sprintf("https://api.hgbrasil.com/finance/stock_price?key=%s&symbol=%s", os.Getenv("HG_KEY"), shares[i].Symbol)
		response, err := http.Get(hgURI)
		if err != nil {
			return nil, err
		}

		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		r := &model.ResponseHG{}

		if err = json.Unmarshal(bodyBytes, &r); err != nil {
			return nil, err
		}

		shares[i].ParseResponse(r, r.Results[shares[i].Symbol].Symbol)
		err = s.Repository.UpdateValue(shares[i])
		if err != nil {
			return nil, err
		}
	}
	return shares, nil
}

// GetAll registers in database
func (s *Service) GetAll() ([]model.Share, error) {
	shares, err := s.Repository.GetAll()

	if err != nil {
		return nil, err
	}

	return shares, nil
}

// Remove service to delete register
func (s *Service) Remove(id int) error {
	err := s.Repository.RemoveShare(id)

	if err != nil {
		return err
	}
	return nil
}
