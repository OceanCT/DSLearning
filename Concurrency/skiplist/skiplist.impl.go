package skiplist

import (
	"math/rand"
	"sync"
)

// NodeSkiplist define nodes in skiplist
type NodeSkiplist[T comparable] struct {
	prev  *NodeSkiplist[T]
	next  *NodeSkiplist[T]
	upper *NodeSkiplist[T]
	down  *NodeSkiplist[T]
	value T
}

// ImplSkiplist define the implementation of Skiplist
type ImplSkiplist[T comparable] struct {
	maxLevel int
	Less     func(T, T) bool
	layers   []struct {
		begin NodeSkiplist[T]
		end   NodeSkiplist[T]
	}
	mutex sync.Mutex
}

// Add add tar to the skiplist
func (impl *ImplSkiplist[T]) Add(tar T) {
	impl.mutex.Lock()
	defer impl.mutex.Unlock()
	startLevel := rand.Intn(impl.maxLevel)
	leftNode := &impl.layers[impl.maxLevel-1].begin
	adjustLeftNode := func() {
		if leftNode.prev == nil && leftNode.next.next != nil && impl.Less(leftNode.next.value, tar) {
			leftNode = leftNode.next
		}
		for impl.Less(leftNode.value, tar) && leftNode.next.next != nil && impl.Less(leftNode.next.value, tar) {
			leftNode = leftNode.next
		}
	}
	lis := make([]*NodeSkiplist[T], startLevel+1)
	for currentLevel := impl.maxLevel - 1; currentLevel >= 0; currentLevel-- {
		adjustLeftNode()
		if currentLevel <= startLevel {
			targetNode := &NodeSkiplist[T]{
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

// Delete delete tar from the skiplist
func (impl *ImplSkiplist[T]) Delete(tar T) (flag bool) {
	impl.mutex.Lock()
	defer impl.mutex.Unlock()
	leftNode := &impl.layers[impl.maxLevel-1].begin
	adjustLeftNode := func() {
		if leftNode.prev == nil && leftNode.next.next != nil && impl.Less(leftNode.next.value, tar) {
			leftNode = leftNode.next
		}
		for impl.Less(leftNode.value, tar) && leftNode.next.next != nil && impl.Less(leftNode.next.value, tar) {
			leftNode = leftNode.next
		}
	}
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

// Count count how many tar in the skiplist
func (impl *ImplSkiplist[T]) Count(tar T) (cnt int64) {
	leftNode := &impl.layers[impl.maxLevel-1].begin
	adjustLeftNode := func() {
		if leftNode.prev == nil && leftNode.next.next != nil && impl.Less(leftNode.next.value, tar) {
			leftNode = leftNode.next
		}
		for impl.Less(leftNode.value, tar) && leftNode.next.next != nil && impl.Less(leftNode.next.value, tar) {
			leftNode = leftNode.next
		}
	}
	for currentLevel := impl.maxLevel - 1; currentLevel >= 0; currentLevel-- {
		adjustLeftNode()
		if currentLevel != 0 {
			leftNode = leftNode.down
		}
	}
	for i := leftNode.next; i != nil; i = i.next {
		if i.value == tar {
			cnt++
		} else {
			break
		}
	}
	return cnt
}

// Empty return if the skiplist is empty or not
func (impl *ImplSkiplist[T]) Empty(tar T) bool {
	return impl.layers[0].begin.next == &impl.layers[0].end
}

// NewSkiplist return a specific type of Skiplist
func NewSkiplist[T comparable](maxLevel int,less func(T,T)bool) (Skiplist[T], error) {
	if maxLevel < 1 {
		return nil, MaxLevelError{}
	} else {
		skiplis := &ImplSkiplist[T]{
			maxLevel: maxLevel,
			layers: make([]struct {
				begin NodeSkiplist[T]
				end   NodeSkiplist[T]
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

// MaxLevelError define the error about maxLevel
type MaxLevelError struct{}

func (MaxLevelError) Error() string { return "MaxLevel should equal to or be greater than 1." }
