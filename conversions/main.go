package main

import (
	"fmt"
)

func main() {
	myForexPortfolio := ForexPortfolio{
		{"INR", 1123},
		{"USD", 344},
		{"AUD", 500},
	}
	myForexPortfolio[0].Convert("USD")
	fmt.Println(myForexPortfolio)

	myRoomHeight := Distance{"M", 4.5}
	myRoomHeight.Convert("CM")
	fmt.Println(myRoomHeight)

	myForexPortfolio.BulkConvert("BHD")

	fmt.Println(myForexPortfolio)
}
