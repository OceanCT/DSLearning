package skiplist



type Skiplist[T comparable] interface {
	Add(T)
	Delete(T) bool
	Count(T) int64
	Empty(T) bool
}


