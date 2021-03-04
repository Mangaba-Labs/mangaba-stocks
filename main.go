package main

import (
	"github.com/Mangaba-Labs/mangaba-stocks.git/database"
	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/cli"
)

func main() {
	database.SetupDatabase()
	cli.StartApp()
	defer database.Instance.Close()
}
