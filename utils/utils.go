package utils

import (
	"math/rand"
	"time"
)

func Random(min, max int) int {
	return rand.Intn(max-min) + min
}

func Select(slice []string) string {
	rand.Seed(time.Now().UnixNano())
	return slice[Random(0, len(slice))]
}
