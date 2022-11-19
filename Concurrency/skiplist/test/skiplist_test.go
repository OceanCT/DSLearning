package skiplist

import (
	"DSLearning/Concurrency/skiplist"
	// "math/rand"
	"testing"
)

func TestSkiplist(t *testing.T) {
	skiplis,err := skiplist.NewSkiplist(200,func(i float64,j float64)bool{return i<j})
	if err!=nil {
		t.Error("The initialization of skiplist breaks down.")
	} 
	skiplis.Add(1)
	// mp := make(map[float64]int64)
	// for i := float64(0); i <= 100; i++ {
	// 	cnt := rand.Intn(300)
	// 	for j := 0; j <= cnt; j++ {
	// 		mp[i] = mp[i] + 1
	// 		skiplis.Add(i)
	// 	}
	// }
	// t.Log("Have add elements to skiplist.")
	// for i := float64(0); i <= 100; i++ {
	// 	if mp[i] != skiplis.Count(i) {
	// 		t.Error("The implementation of skiplist is not correct.")
	// 	}
	// }
}