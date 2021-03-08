package pythonlikego

type ListIter struct {
	List *List
	Idx  int
}

func (li ListIter) HasNext() bool {
	if li.Idx < li.List.getLastIndex() {
		return true
	}
	return false
}
