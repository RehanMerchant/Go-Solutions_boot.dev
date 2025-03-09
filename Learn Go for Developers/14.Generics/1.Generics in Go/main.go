package main

func getLast[T any](s []T) T {

	var zero T
	len := len(s)

	if len > 0 {
		return s[len-1]
	}

	return zero

}
