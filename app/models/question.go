package models

import (
	"math/rand"
	"time"
)

func Numerics() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}
