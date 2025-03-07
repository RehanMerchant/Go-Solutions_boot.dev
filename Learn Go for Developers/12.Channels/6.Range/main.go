package main

func concurrentFib(n int) []int {

	res := []int{}
	ch := make(chan int)
	go fibonacci(n, ch)

	for item := range ch {
		res = append(res, item)
	}

	return res
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
