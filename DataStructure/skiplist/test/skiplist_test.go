package skiplist

import (
	// "fmt"
	"DSLearning/DataStructure/skiplist"
	"math/rand"
	"testing"
)

func TestSkiplist(t *testing.T) {
	skiplis, _ := skiplist.NewSkiplist(200)
	mp := make(map[float64]int)
	for i := float64(0); i <= 100; i++ {
		cnt := rand.Intn(300)
		for j := 0; j <= cnt; j++ {
			mp[i] = mp[i] + 1
			skiplis.Add(i)
		}
	}
	for i := float64(0); i <= 100; i++ {
		if mp[i] != skiplis.Count(i) {
			t.Error("The implementation of skiplist is not correct.")
			break
		}
	}
}
