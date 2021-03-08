package pythonlikego

type ListIter struct {
	List *List
	Idx  int
}

//HasNext check if ListIter has next value
func (li ListIter) HasNext() bool {
	if li.Idx < li.List.getLastIndex() {
		return true
	}
	return false
}

//Next Retrieve the next item from the iterator.
//if the iterator is exhausted, nil returned,
func (li ListIter) Next() interface{} {
	return li.List.getByIndex(li.Idx)
}
