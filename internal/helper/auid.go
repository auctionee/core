package helper

import (
	"math/rand"

	"github.com/auctionee/core/internal/core/data"
)

var tokens string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GetAUID() data.AUID {
	rand.Seed(rand.Int63())
	str := ""
	for i := 0; i < 6; i++ {
		str += string(tokens[rand.Intn(len(tokens))])
	}
	return data.AUID(str)
}
