package main

import (
	"github.com/Mangaba-Labs/mangaba-stocks.git/database"
	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/cli"
	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/config"
)

func main() {
	config.SetupEnvVars()
	database.SetupDatabase()
	cli.StartApp()
	defer database.Instance.Close()
}
