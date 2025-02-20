package main

import (
	"fmt"
)

func printPrimes(max int) {

	for i := 2; i <= max; i++ {

		if i == 2 {
			fmt.Println(i)
			continue
		}
		if i%2 == 0 {
			continue
		}
		isPrime := true
		for j := 3; j*j < i+1; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if !isPrime {
			continue
		}

		fmt.Println(i)

	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
