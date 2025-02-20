package main

func bulkSend(numMessages int) float64 {

	cost := 0.0
	for i := 0; i < numMessages; i++ {
		cost += 1.0 + (0.01 * float64(i))
	}

	return cost
}
