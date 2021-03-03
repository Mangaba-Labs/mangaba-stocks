package main

import (
	"fmt"

	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/storage"
)

func main() {
	shares := []storage.Share{}
	storage.SetupDatabase()
	shares = storage.GetAllShares()
	// cli.StartApp()
	fmt.Println(shares)

	defer storage.Instance.Close()
}
