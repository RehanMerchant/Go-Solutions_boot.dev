package main

import (
	"errors"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {

		return 0.0, errors.New("no dividing by 0")

	}
	return x / y, nil
}
