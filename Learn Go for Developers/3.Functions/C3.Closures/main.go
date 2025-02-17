package main

func adder() func(int) int {
	sum := 0
	return func(a int) int {
		sum += a
		return sum
	}

}
