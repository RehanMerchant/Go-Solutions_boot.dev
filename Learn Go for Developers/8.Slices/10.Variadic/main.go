package main

func sum(nums ...int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum
}
