package Skiplist

type Skiplist interface{
	Add(int)
	Delete(int)bool
	Search(int)bool
}

type SkiplistImpl struct{
	/*
	* TODO: Implement SkiplistImpl to satisfy the requirements of the arrorging interface
	*/
}
func NewSkiplist() Skiplist{
	return &SkiplistImpl{}
}
func (impl *SkiplistImpl)Add(tar int){
}
func (impl *SkiplistImpl)Delete(tar int)bool{
	return true
}
func (impl *SkiplistImpl)Search(tar int)bool{
	return true
}