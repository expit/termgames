package fastReader

import (
	"math"
	. "math/rand"
	"time"
)

func NumberGenerator(l float64 )int {
	Seed(time.Now().UnixNano())
	min := math.Pow(10,l-1)
	max := math.Pow(10,l)-1
	return Intn(int(max-min+1)) + int(min)
}
