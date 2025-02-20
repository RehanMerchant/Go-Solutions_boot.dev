package main

func maxMessages(thresh int) int {

	cost := 0
	for i := 0; ; i++ {
		cost += 100 + i
		if cost > thresh {
			return i
		}
	}

}
