package storage

// GetAllShares in database
func GetAllShares() []Share {
	rows, _ := Instance.Query("SELECT id, company_name, name, buy_value, now_value FROM shares")
	shares := []Share{}
	tempShare := Share{}
	for rows.Next() {
		rows.Scan(&tempShare.id, &tempShare.companyName, &tempShare.name, &tempShare.buyValue, &tempShare.nowValue)
		shares = append(shares, tempShare)
	}
	return shares
}

func insertShare(share Share) error {
	statement, err := Instance.Prepare("INSERT INTO shares (company_name, name, buy_value, now_value) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(share.companyName, share.name, share.buyValue, share.nowValue)
	if err != nil {
		return err
	}
	return nil
}
