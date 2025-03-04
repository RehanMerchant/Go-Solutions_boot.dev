package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	result := []float64{}
	for _, c := range costs {
		if c.day == day {
			result = append(result, c.value)
		}
	}

	return result

}
