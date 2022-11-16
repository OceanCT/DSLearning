package skiplist

import (
	"math/rand"
	"time"
)

func init() {
	// generate rand seed
	rand.Seed(time.Now().UnixNano())
}

type Skiplist interface {
	Add(float64)
	Delete(float64) bool
	Count(float64) int
	Empty() bool
}
