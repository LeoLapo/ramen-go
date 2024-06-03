package utils

import "math/rand"

func GenerateOrderID() (int, error) {
	return rand.Intn(1000000), nil
}
