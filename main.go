package main

import (
	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/cli"
	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/storage"
)

func main() {
	storage.SetupDatabase()
	cli.StartApp()
	defer storage.Instance.Close()
}
