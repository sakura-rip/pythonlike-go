package pythonlike_go

import "fmt"

//List is python like list object in Golang
type List []interface{}

//NewList create new list
func NewList() List {
	return List{}
}

//Index return zero-based index in the list of the first item whose value is equal to x.
//return -1 if there is no such item.
func (l List) Index(element interface{}) int {
	for k, v := range l {
		if element == v {
			return k
		}
	}
	return -1
}

// Add an item to the end of the list.
func (l *List) Append(element interface{}) {
	*l = append(*l, element)
}

//Remove the first item from the list whose value is equal to x. It returns an error if there is no such item.
func (l *List) Remove(element interface{}) error {
	idx := l.Index(element)
	if idx == -1 {
		return fmt.Errorf("element not found")
	}
	return l.removeByIndex(idx)
}

//removeByIndex remove selected index from list
func (l *List) removeByIndex(idx int) error {
	list := []interface{}(*l)
	*l = append(list[:idx], list[idx+1:]...)
	return nil
}

//Length get the length of list
func (l List) Length() int {
	return len(l)
}

//Insert an item at a given position.
//The first argument is the index of the element before which to insert,
//so a.Insert(0, x) inserts at the front of the list,
//and a.Insert(len(a)-1, x) is equivalent to a.Append(x).
func (l *List) Insert(idx int, element interface{}) error {
	if l.Length() < idx || idx < 0 {
		return fmt.Errorf("index out of range")
	}
	list := []interface{}(*l)
	*l = append(list[:idx], append([]interface{}{element}, list[idx:]...)...)
	return nil
}

func (l List) getLastIndex() int {
	return l.Length() - 1
}

//Pop remove the item at the given position in the list, and return it.
//If no index is specified, a.Pop() removes and returns the last item in the list.
func (l *List) Pop(idx ...int) (interface{}, error) {
	index := l.getLastIndex()
	ll := len(idx)

	if ll > 1 {
		return nil, fmt.Errorf("only 1 or 0 arguments are allowed")
	}

	// in case of `l.Pop()`
	if ll == 0 {
		element := []interface{}(*l)[index]
		return element, l.removeByIndex(index)
	}

	if idx[0] > index {
		return nil, fmt.Errorf("index out of range")
	}

	index = idx[0]
	element := []interface{}(*l)[index]
	return element, l.removeByIndex(index)
}

//Clear remove all items from the list
func (l *List) Clear() {
	*l = NewList()
}

//Count Return the number of times x appears in the list.
func (l List) Count(element interface{}) int {
	count := 0
	for _, v := range l {
		if v == element {
			count++
		}
	}
	return count
}

//Copy Return a shallow copy of the list
func (l List) Copy() List {
	return l
}
