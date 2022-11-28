package skiplist

import (
	"DSLearning/Concurrency/skiplist"
	"math/rand"
	"testing"
)

func TestSkiplist(t *testing.T) {
	skiplis, err := skiplist.NewSkiplist(200, func(i, j float64) bool { return i < j })
	if err != nil {
		t.Error("The initialization of skiplist breaks down.")
	}
	ceil := 3.0
	mp := make(map[float64]int64, int64(ceil))
	// test Add()
	for i := float64(0); i <= ceil; i++ {
		mp[i] = int64(rand.Intn(3))
		for j := int64(1); j <= mp[i]; j++ {
			skiplis.Add(i)
		}
	}
	// test Count()
	for i := float64(0); i <= ceil; i++ {
		if mp[i] != skiplis.Count(i) {
			t.Error("The implementation of skiplist is not correct. Add() or Count() is not correct.")
			break
		}
	}
	t.Log(mp)
	// test Delete()
	for i := float64(0); i <= ceil; i++ {
		if mp[i] > 0 {
			cnt := int64(rand.Intn(int(mp[i])))
			for j := int64(1); j <= cnt; j++ {
				if !skiplis.Delete(i) {
					t.Error("The implementation of skiplist is not correct. Delete() is not correct.")
				}
			}
			mp[i] = mp[i] - cnt
		}
	}
	for i := float64(0); i <= ceil; i++ {
		if mp[i] != skiplis.Count(i) {
			t.Log(mp)
			t.Log(i, skiplis.Count(i))
			t.Error("The implementation of skiplist is not correct. Delete() or Count() is not correct.")
			break
		}
	}
}
