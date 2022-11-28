package skiplist

import (
	"math/rand"
)

// NodeSkiplist define nodes in skiplist
type NodeSkiplist[T comparable] struct {
	prev  *NodeSkiplist[T]
	next  *NodeSkiplist[T]
	upper *NodeSkiplist[T]
	down  *NodeSkiplist[T]
	value T
}

// Tall tell whether a node is the tail
func (node *NodeSkiplist[T]) IsTail() bool {
	return node.next.IsHead()
}

// Head tell whether a head is the head
func (node *NodeSkiplist[T]) IsHead() bool {
	return node.prev == nil
}

// ImplSkiplist define the implementation of Skiplist
type ImplSkiplist[T comparable] struct {
	// maxLevel to stop the level of skiplist from growing unlimitedly
	maxLevel int
	// Less self-defined Less function between T-type variables
	less func(T, T) bool
	// layers use dulNode to define a layer
	layerHeaders []*NodeSkiplist[T]
}

// NewSkiplist generate an specific type of skiplist
func NewSkiplist[T comparable](maxLevel int, less func(T, T) bool) (Skiplist[T], error) {
	// if maxLevel does not statisfy the requirement
	if maxLevel < 1 {
		return nil, MaxLevelError{}
	} else {
		skiplist := &ImplSkiplist[T]{
			maxLevel:     maxLevel,
			less:         less,
			layerHeaders: make([]*NodeSkiplist[T], maxLevel),
		}
		for i := 0; i < maxLevel; i++ {
			skiplist.layerHeaders[i] = &NodeSkiplist[T]{}
			skiplist.layerHeaders[i].next = skiplist.layerHeaders[i]
			if i != 0 {
				skiplist.layerHeaders[i].down = skiplist.layerHeaders[i-1]
				skiplist.layerHeaders[i-1].upper = skiplist.layerHeaders[i]
			}
		}
		return skiplist, nil
	}
}

// Add add tar to the skiplist
func (impl *ImplSkiplist[T]) Add(tar T) {
	// impl.mutex.Lock()
	// defer impl.mutex.Unlock()
	// startLevel to define the highest level tar can be in
	startLevel := rand.Intn(impl.maxLevel)
	// leftNode is the expected leftNode after successfullt inserting tar into the skiplist
	leftNode := impl.layerHeaders[impl.maxLevel-1]
	// adjustLeftNode to adjust the leftNode in one layer to approach the exact place tar should be
	adjustLeftNode := func() {
		// while 1. current node is the head node or the value of current node is smaller than tar
		//       2. the next node it is not the head node and its value is smaller than tar
		//       move the current node to its next node
		for (leftNode.IsHead() || impl.less(leftNode.value, tar)) && !leftNode.next.IsHead() && impl.less(leftNode.next.value, tar) {
			leftNode = leftNode.next
		}
	}
	// lis to store the nodes added during the process
	lis := make([]*NodeSkiplist[T], startLevel+1)
	// from the highest layer to lower layers
	for currentLevel := impl.maxLevel - 1; currentLevel >= 0; currentLevel-- {
		adjustLeftNode()
		// if the layer should be inserted tar into
		if currentLevel <= startLevel {
			targetNode := &NodeSkiplist[T]{
				prev:  leftNode,
				next:  leftNode.next,
				value: tar,
			}
			leftNode.next = targetNode
			if !targetNode.IsTail() {
				targetNode.next.prev = targetNode
			}
			lis[currentLevel] = targetNode
		}
		leftNode = leftNode.down
	}
	// for the inserted node, init their upper and down
	for i := 0; i < startLevel; i++ {
		lis[i].upper = lis[i+1]
	}
	for i := startLevel; i > 0; i-- {
		lis[i].down = lis[i-1]
	}
}

// Delete delete tar once from the skiplist and return false if there exists no tar in the skiplist
func (impl *ImplSkiplist[T]) Delete(tar T) (flag bool) {
	//	impl.mutex.Lock()
	//	defer impl.mutex.Unlock()
	// leftNode is the expected leftNode after successfullt inserting tar into the skiplist
	leftNode := impl.layerHeaders[impl.maxLevel-1]
	// adjustLeftNode to adjust the leftNode in one layer to approach the exact place tar should be
	adjustLeftNode := func() {
		// while 1. current node is the head node or the value of current node is smaller than tar
		//       2. the next node it is not the head node and its value is smaller than tar
		//       move the current node to its next node
		for (leftNode.IsHead() || impl.less(leftNode.value, tar)) && !leftNode.next.IsHead() && impl.less(leftNode.next.value, tar) {
			leftNode = leftNode.next
		}
	}
	for currentLevel := impl.maxLevel - 1; currentLevel >= 0; currentLevel-- {
		adjustLeftNode()
		if !leftNode.IsTail() && leftNode.next.value == tar {
			flag = true
			targetNode := leftNode.next
			for targetNode != nil {
				targetNode.prev.next = targetNode.next
				if !targetNode.IsTail() {
					targetNode.next.prev = targetNode.prev
				}
				targetNode = targetNode.down
			}
			break
		}
		leftNode = leftNode.down
	}
	return flag
}

// Count count how many tar in the skiplist
func (impl *ImplSkiplist[T]) Count(tar T) (cnt int64) {
	//	impl.mutex.Lock()
	//	defer impl.mutex.Unlock()
	// leftNode is the expected leftNode after successfullt inserting tar into the skiplist
	leftNode := impl.layerHeaders[impl.maxLevel-1]
	// adjustLeftNode to adjust the leftNode in one layer to approach the exact place tar should be
	adjustLeftNode := func() {
		// while 1. current node is the head node or the value of current node is smaller than tar
		//       2. the next node it is not the head node and its value is smaller than tar
		//       move the current node to its next node
		for (leftNode.IsHead() || impl.less(leftNode.value, tar)) && !leftNode.next.IsHead() && impl.less(leftNode.next.value, tar) {
			leftNode = leftNode.next
		}
	}
	for currentLevel := impl.maxLevel - 1; currentLevel > 0; currentLevel-- {
		adjustLeftNode()
		leftNode = leftNode.down
	}
	for i := leftNode.next; !i.IsHead(); i = i.next {
		if i.value == tar {
			cnt++
		} else {
			break
		}
	}
	return cnt
}

// Empty return if the skiplist is empty or not
func (impl *ImplSkiplist[T]) Empty() bool {
	return impl.layerHeaders[0].IsHead()
}

// ToString convert skiplist to string, which shows the inner structure of skiplist
func (impl *ImplSkiplist[T]) ToString(toString func(T) string) (res string) {
	for _, i := range impl.layerHeaders {
		res += "\nLayer:"
		for i = i.next; !i.IsHead(); i = i.next {
			res += toString(i.value) + ","
		}
	}
	res += "\n"
	return res
}

// ToMap convert skiplist to map
func (impl *ImplSkiplist[T]) ToMap() (result map[T]int64) {
	result = make(map[T]int64)
	for i := impl.layerHeaders[0].next; !i.IsHead(); i = i.next {
		result[i.value] = result[i.value] + 1
	}
	return result
}

// MaxLevelError define the error about maxLevel
type MaxLevelError struct{}

func (MaxLevelError) Error() string { return "MaxLevel should equal to or be greater than 1." }
