package main

func getMessageCosts(messages []string) []float64 {

	costs := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		message := messages[i]
		cost := float64(len(message)) * 0.01
		costs[i] = cost

	}
	return costs

}
