package main

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	baseBill := calculateBaseBill(costPerMessage, numMessages)
	discount := calculateDiscountRate(numMessages) / 100.0
	newBill := baseBill - (baseBill * discount)
	return newBill
}

func calculateDiscountRate(messagesSent int) float64 {
	if messagesSent <= 1000 {
		return 0.0
	} else if messagesSent > 1000 && messagesSent <= 5000 {
		return 10.0
	} else {
		return 20.0
	}
}

// don't touch below this line

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}
