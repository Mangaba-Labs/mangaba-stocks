package cli

import (
	"fmt"
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
	stockHandler.CreateShare(share)
}

func removeShare() {
	var id int
	fmt.Println("Which share do u want remove? (ID):")
	fmt.Scanf("%d", &id)
	// stock.RemoveShare(id)
}

func quit() {
	os.Exit(0)
}
