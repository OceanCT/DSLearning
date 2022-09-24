# Skiplist

Implement skiplist and try to pass the test.

## Interface

```go
package skiplist
type Skiplist interface {
	/*
	*   Add a float64 number into skiplist 
	 */
	Add(float64)
	/*
	*   Delete a float64 number from skiplist
	*   if there exists more than one same number in the skiplist, 
	*       delete only one
	*   else if there exists no such number in the skiplist, 
	*       return false
	 */
	Delete(float64) bool
	/*
	*   Return the number of a float64 in the skiplist  
	 */
	Count(float64) bool
	/*
	*   Return if the skiplist is empty
	 */
	Empty() bool
}
```



