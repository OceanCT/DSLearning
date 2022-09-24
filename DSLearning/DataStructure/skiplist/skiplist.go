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
type NodeSkiplist struct {
	prev  *NodeSkiplist
	next  *NodeSkiplist
	upper *NodeSkiplist
	down  *NodeSkiplist
	value float64
}
type ImplSkiplist struct {
	maxLevel int
	layers   []struct {
		begin NodeSkiplist
		end   NodeSkiplist
	}
}
type MaxLevelError struct{}

func (MaxLevelError) Error() string { return "MaxLevel should equal to or be greater than 1." }
func NewSkiplist(maxLevel int) (Skiplist, error) {
	if maxLevel < 1 {
		errorReturnFunc := func() error {
			return MaxLevelError{}
		}
		return nil, errorReturnFunc()
	} else {
		skiplis := &ImplSkiplist{
			maxLevel: maxLevel,
			layers: make([]struct {
				begin NodeSkiplist
				end   NodeSkiplist
			}, maxLevel),
		}
		for i := 0; i < maxLevel-1; i++ {
			skiplis.layers[i].begin.upper = &skiplis.layers[i+1].begin
			skiplis.layers[i].end.upper = &skiplis.layers[i+1].end
		}
		for i := maxLevel - 1; i > 0; i-- {
			skiplis.layers[i].begin.down = &skiplis.layers[i-1].begin
			skiplis.layers[i].end.down = &skiplis.layers[i-1].end
		}
		for i := 0; i < maxLevel; i++ {
			skiplis.layers[i].begin.next = &skiplis.layers[i].end
			skiplis.layers[i].end.prev = &skiplis.layers[i].begin
		}
		return skiplis, nil
	}
}
func (impl *ImplSkiplist) Add(tar float64) {
	startLevel := rand.Intn(impl.maxLevel)
	leftNode := &impl.layers[impl.maxLevel-1].begin
	adjustLeftNode := func() {
		if leftNode.prev == nil && leftNode.next.next != nil && leftNode.next.value < tar {
			leftNode = leftNode.next
		}
		for leftNode.value < tar && leftNode.next.next != nil && leftNode.next.value < tar {
			leftNode = leftNode.next
		}
	}
	lis := make([]*NodeSkiplist, startLevel+1)
	for currentLevel := impl.maxLevel - 1; currentLevel >= 0; currentLevel-- {
		adjustLeftNode()
		if currentLevel <= startLevel {
			targetNode := &NodeSkiplist{
				prev:  leftNode,
				next:  leftNode.next,
				value: tar,
			}
			leftNode.next = targetNode
			targetNode.next.prev = targetNode
			lis[currentLevel] = targetNode
		}
		leftNode = leftNode.down
	}
	for i := 0; i < startLevel; i++ {
		lis[i].upper = lis[i+1]
	}
	for i := startLevel; i > 0; i-- {
		lis[i].down = lis[i-1]
	}
}
func (impl *ImplSkiplist) Delete(tar float64) bool {
	leftNode := &impl.layers[impl.maxLevel-1].begin
	adjustLeftNode := func() {
		if leftNode.prev == nil && leftNode.next.next != nil && leftNode.next.value < tar {
			leftNode = leftNode.next
		}
		for leftNode.value < tar && leftNode.next.next != nil && leftNode.next.value < tar {
			leftNode = leftNode.next
		}
	}
	flag := false
	for currentLevel := impl.maxLevel - 1; currentLevel >= 0; currentLevel-- {
		adjustLeftNode()
		if leftNode.next != nil && leftNode.next.value == tar {
			flag = true
			for i := leftNode.next; i != nil; i = i.down {
				i.prev.next = i.next
				i.next.prev = i.prev
			}
			break
		}
	}
	return flag
}
func (impl *ImplSkiplist) Count(tar float64) int {
	leftNode := &impl.layers[impl.maxLevel-1].begin
	adjustLeftNode := func() {
		if leftNode.prev == nil && leftNode.next.next != nil && leftNode.next.value < tar {
			leftNode = leftNode.next
		}
		for leftNode.value < tar && leftNode.next.next != nil && leftNode.next.value < tar {
			leftNode = leftNode.next
		}
	}
	for currentLevel := impl.maxLevel - 1; currentLevel >= 0; currentLevel-- {
		adjustLeftNode()
		if currentLevel != 0 {
			leftNode = leftNode.down
		}
	}
	cnt := 0
	for i := leftNode.next; i != nil; i = i.next {
		if i.value == tar {
			cnt++
		} else {
			break
		}
	}
	return cnt
}
func (impl *ImplSkiplist) Empty() bool {
	return impl.layers[0].begin.next == &impl.layers[0].end
}
