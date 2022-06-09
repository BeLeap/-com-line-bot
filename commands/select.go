package commands

import (
	"com-line-bot/utils"
	"math/rand"
	"time"
)

func Select(slice []string) string {
	rand.Seed(time.Now().UnixNano())
	return slice[utils.Random(0, len(slice))]
}