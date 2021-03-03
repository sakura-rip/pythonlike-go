package pythonlike_go

import "fmt"

type List []interface{}

func NewList() List {
	return List{}
}

func (l List) Index(element interface{}) int {
	for k, v := range l {
		if element == v {
			return k
		}
	}
	return -1
}

func (l *List) Append(element interface{}) {
	*l = append(*l, element)
}

func (l *List) Remove(element interface{}) error {
	idx := l.Index(element)
	if idx == -1 {
		return fmt.Errorf("element not found")
	}
	return l.removeByIndex(idx)
}

func (l *List) removeByIndex(idx int) error {
	list := []interface{}(*l)
	*l = append(list[:idx], list[idx+1:]...)
	return nil
}

func (l List) Length() int {
	return len(l)
}

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

func (l *List) Pop(idx ...int) (interface{}, error) {
	index := l.getLastIndex()
	if ll := len(idx); ll > 1 {
		return nil, fmt.Errorf("only 1 or 0 arguments are allowed")
	} else if ll == 0 {
		//	下の条件分岐でINDEX ERRORが起きないように
	} else if idx[0] > index {
		return nil, fmt.Errorf("index out of range")
	} else {
		index = idx[0]
	}
	element := []interface{}(*l)[index]
	return element, l.removeByIndex(index)
}

func (l *List) Clear() {
	*l = NewList()
}

func (l List) Count(element interface{}) int {
	count := 0
	for _, v := range l {
		if v == element {
			count++
		}
	}
	return count
}
