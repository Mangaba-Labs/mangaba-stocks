package repository

import (
	"log"

	"github.com/Mangaba-Labs/mangaba-stocks.git/database"
	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/stock/model"
)

// Repository concrete type
type Repository struct{}

// GetAll get all shares in db
func (r Repository) GetAll() ([]model.Share, error) {
	rows, err := database.Instance.Query("SELECT id, company_name, name, buy_value, now_value FROM shares")
	if err != nil {
		return nil, err
	}
	shares := []model.Share{}
	tempShare := model.Share{}
	for rows.Next() {
		rows.Scan(&tempShare.ID, &tempShare.CompanyName, &tempShare.Symbol, &tempShare.BuyValue, &tempShare.NowValue)
		shares = append(shares, tempShare)
	}
	return shares, nil
}

// InsertShare add a new share
func (r Repository) InsertShare(share model.Share) (model.Share, error) {
	statement, err := database.Instance.Prepare("INSERT INTO shares (company_name, name, buy_value, now_value) VALUES (?, ?, ?, ?)")
	if err != nil {
		return share, err
	}
	_, err = statement.Exec(share.CompanyName, share.Symbol, share.BuyValue, share.NowValue)
	if err != nil {
		return share, err
	}
	return share, nil
}

// RemoveShare delete share from database
func (r Repository) RemoveShare(id int) error {
	statement, err := database.Instance.Prepare("DELETE FROM shares WHERE id = (?)")

	if err != nil {
		return err
	}

	_, err = statement.Exec(id)

	if err != nil {
		return err
	}
	return nil
}

// UpdateValue update de now_value of a share
func (r Repository) UpdateValue(share model.Share) error {
	statement, err := database.Instance.Prepare("UPDATE shares SET now_value = ? WHERE id = ?")

	if err != nil {
		log.Fatalln(err)
		return err
	}

	_, err = statement.Exec(share.NowValue, share.ID)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}
