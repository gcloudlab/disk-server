package helper

import (
	"math/rand"
	"strconv"
	"time"
)

func Random() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(30) + 1
	return strconv.Itoa(r)
}
