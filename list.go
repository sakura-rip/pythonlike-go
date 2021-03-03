package pythonlike_go

import "fmt"

type List []interface{}

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
