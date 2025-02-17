package main

import "fmt"

func main() {
	var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful."
	var accountBalance float64 = 100.0
	var bulkMessageCost float64 = 75.0
	var isPremiumUser bool = true
	var discountRate float64 = 0.10
	var finalCost float64

	// don't edit above this line

	finalCost = bulkMessageCost
	discount := 0.0
	if isPremiumUser == true {
		discount = finalCost * discountRate
	}

	finalCost = finalCost - discount

	if finalCost < accountBalance {
		accountBalance = accountBalance - finalCost
		fmt.Print(purchaseSuccessMessage)
	} else {
		fmt.Print(insufficientFundMessage)
	}

	// don't edit below this line

	fmt.Println("Account balance:", accountBalance)
}
