package repository

import (
	"database/sql"

	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/stock/model"
)

// Repository concrete type
type Repository struct {
	db *sql.DB
}

// GetAll get all shares in db
func (r Repository) GetAll() []model.Share {
	rows, _ := r.db.Query("SELECT id, company_name, name, buy_value, now_value FROM shares")
	shares := []model.Share{}
	tempShare := model.Share{}
	for rows.Next() {
		rows.Scan(&tempShare.ID, &tempShare.CompanyName, &tempShare.Name, &tempShare.BuyValue, &tempShare.NowValue)
		shares = append(shares, tempShare)
	}
	return shares
}

// InsertShare add a new share
func (r Repository) InsertShare(name string) error {
	// statement, err := database.Instance.Prepare("INSERT INTO shares (company_name, name, buy_value, now_value) VALUES (?, ?, ?, ?)")
	// if err != nil {
	// 	return err
	// }
	// _, err = statement.Exec(share.CompanyName, share.Name, share.BuyValue, share.NowValue)
	// if err != nil {
	// 	return err
	// }
	return nil
}

// RemoveShare delete share from database
func (r Repository) RemoveShare(id int) error {
	statement, err := r.db.Prepare("DELETE FROM shares WHERE id = (?)")

	if err != nil {
		return err
	}

	_, err = statement.Exec(id)

	if err != nil {
		return err
	}
	return nil
}
