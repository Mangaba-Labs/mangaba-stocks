package cli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/stock/handler"
	"github.com/Mangaba-Labs/mangaba-stocks.git/pkg/domain/stock/model"
	"github.com/gookit/color"
)

var shares []model.Share

// StartApp entry point to the application
func StartApp() {
	var stockHandler = new(handler.Handler)

	shares, err := stockHandler.GetAll()

	if err != nil {
		fmt.Println(err)
		return
	}
	displayDashboard(shares)
	menu()
}

func displayDashboard(shares []model.Share) {
	clear()
	baseColor := "#bf5af2" // ID            Company Name             Share Name   Price   Price Actual
	color.HEX(baseColor).Println("##############################################################################")
	color.HEX(baseColor).Println("#        #                                  #          #          #          #")
	color.HEX(baseColor).Println("#   ID   #           Company Name           #   Name   #  Price   #    Now   #")
	color.HEX(baseColor).Println("#        #                                  #          #          #          #")
	color.HEX(baseColor).Println("##############################################################################")

	for i := 0; i < len(shares); i++ {
		color.HEX(baseColor).Println("#        #                                  #          #          #          #")
		companySides := getSides(34, len(shares[i].CompanyName))
		nameSides := getSides(10, len(shares[i].Symbol))
		buyValueString := fmt.Sprintf("R$%.2f", shares[i].BuyValue)
		buyValueSides := getSides(10, len(buyValueString))
		nowValueString := fmt.Sprintf("R$%.2f", shares[i].NowValue)
		nowValueSides := getSides(10, len(nowValueString))

		color.HEX(baseColor).Printf("#")
		if shares[i].ID > 9 {
			fmt.Printf("   %d   ", shares[i].ID)
		} else {
			fmt.Printf("   %d    ", shares[i].ID)
		}
		color.HEX(baseColor).Printf("#")
		fmt.Printf("%s%s%s", companySides, shares[i].CompanyName, companySides)
		color.HEX(baseColor).Printf("#")
		if len(shares[i].Symbol)%2 == 0 {
			fmt.Printf("%s%s%s", nameSides, shares[i].Symbol, nameSides)
		} else {
			fmt.Printf("%s%s %s", nameSides, shares[i].Symbol, nameSides)
		}
		color.HEX(baseColor).Printf("#")
		if len(buyValueString)%2 == 0 {
			fmt.Printf("%s%s%s", buyValueSides, buyValueString, buyValueSides)
		} else {
			fmt.Printf("%s%s %s", buyValueSides, buyValueString, buyValueSides)
		}
		color.HEX(baseColor).Printf("#")
		if len(nowValueString)%2 == 0 {
			fmt.Printf("%s%s%s", nowValueSides, nowValueString, nowValueSides)
		} else {
			fmt.Printf("%s%s %s", nowValueSides, nowValueString, nowValueSides)
		}
		color.HEX(baseColor).Printf("#\n")
		color.HEX(baseColor).Println("#        #                                  #          #          #          #")
		color.HEX(baseColor).Println("##############################################################################")
	}
}

func getSides(fieldSpace int, itemLength int) string {
	fieldSides := fieldSpace - itemLength
	return strings.Repeat(" ", fieldSides/2)
}

func menu() {
	fmt.Printf("\n")
	for {
		fmt.Println("1- Add share    \t2- Remove Share")
		fmt.Println("3- Quit")
		var cmd int
		fmt.Scanf("%d", &cmd)

		if cmd == 1 {
			addShare()
		} else if cmd == 2 {
			removeShare()
			// shares = stock.GetAllShares()
			displayDashboard(shares)
		} else if cmd == 3 {
			quit()
		} else {
			fmt.Println("Huh? Invalid option")
		}
	}
}

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
