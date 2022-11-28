package skiplist

type Skiplist[T comparable] interface {
	// Add add one to skiplist, allow dipulicate elements
	Add(T)
	// Delete one from skiplist, and return if there is one element deleted
	// 	Precisely, return false if no such element in this skiplist  
	Delete(T) bool
	// Count return the count of target element in skiplist 
	Count(T) int64
	// Empty return if there is no element in the skiplist
	Empty() bool
	// ToString() return a string that show the inner structure pf the skiplist
	ToString(toString func(T)string)string
	// ToMap() convert this skiplist to a map
	ToMap() map[T]int64
}
