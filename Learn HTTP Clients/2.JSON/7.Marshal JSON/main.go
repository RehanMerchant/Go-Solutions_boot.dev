package main

import (
	"encoding/json"
)

func marshalAll[T any](items []T) ([][]byte, error) {

	var res [][]byte

	for _, val := range items {
		data, err := json.Marshal(val)
		if err != nil {
			return nil, err
		}
		res = append(res, data)
	}

	return res, nil

}
