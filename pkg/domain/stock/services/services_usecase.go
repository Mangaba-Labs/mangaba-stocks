package services

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/stock/model"
	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/stock/repository"
)

// Service struct
type Service struct {
	Repository repository.StockRepository
}

// AddShare take the actual value from the Share and send a share to repo
func (s *Service) AddShare(share model.Share) error {
	// s.Repository.InsertShare(share.Name)
	hgURI := fmt.Sprintf("https://api.hgbrasil.com/finance/stock_price?key=%s&symbol=%s", os.Getenv("HG_KEY"), share.Name)

	response, err := http.Get(hgURI)

	if err != nil {
		return err
	}

	fmt.Println(hgURI)
	return nil
}
