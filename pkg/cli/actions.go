package cli

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/stock/handler"
	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/stock/model"
)

func addShare() {
	var share model.Share
	var stockHandler = new(handler.Handler)

	fmt.Println("Which share do u want track?(ex: PETR4):")
	fmt.Scanf("%s", &share.Symbol)
	fmt.Println("Buy price(float):")
	fmt.Scanf("%f", &share.BuyValue)
	share.Symbol = strings.ToUpper(share.Symbol)
	err := stockHandler.CreateShare(share)

	if err != nil {
		log.Fatalln(err)
	}
	clear()
	StartApp()
}

func removeShare() {
	var stockHandler = new(handler.Handler)
	var id int
	fmt.Println("Which share do u want remove? (ID):")
	fmt.Scanf("%d", &id)
	stockHandler.Remove(id)
}

func quit() {
	os.Exit(0)
}
