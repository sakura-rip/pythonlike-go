package pythonlikego

import (
	"fmt"
	"math/rand"
	"time"
)

//List is python like list object in Golang
type List []interface{}

//NewList create new list
func NewList() List {
	return List{}
}

func NewListFrom(val ...interface{}) List {
	list := NewList()
	for _, v := range val {
		list.Append(v)
	}
	return list
}

func NewListFromSlice(slice []interface{}) List {
	return NewListFrom(slice...)
}

//Index return zero-based index in the list of the first item whose value is equal to x.
//return -1 if there is no such item.
func (list List) Index(element interface{}) int {
	for k, v := range list {
		if element == v {
			return k
		}
	}
	return -1
}

// Add an item to the end of the list.
func (list *List) Append(element interface{}) {
	*list = append(*list, element)
}

//Remove the first item from the list whose value is equal to x. It returns an error if there is no such item.
func (list *List) Remove(element interface{}) error {
	idx := list.Index(element)
	if idx == -1 {
		return fmt.Errorf("element not found")
	}
	return list.removeByIndex(idx)
}

//removeByIndex remove selected index from list
func (list *List) removeByIndex(idx int) error {
	list_ := []interface{}(*list)
	*list = append(list_[:idx], list_[idx+1:]...)
	return nil
}

func (list List) getByIndex(idx int) interface{} {
	if list.getLastIndex() < idx {
		return nil
	}
	return list[idx]
}

//Length get the length of list
func (list List) Length() int {
	return len(list)
}

//Insert an item at a given position.
//The first argument is the index of the element before which to insert,
//so a.Insert(0, x) inserts at the front of the list,
//and a.Insert(len(a)-1, x) is equivalent to a.Append(x).
func (list *List) Insert(idx int, element interface{}) error {
	if list.Length() < idx || idx < 0 {
		return fmt.Errorf("index out of range")
	}
	list_ := []interface{}(*list)
	*list = append(list_[:idx], append([]interface{}{element}, list_[idx:]...)...)
	return nil
}

func (list List) getLastIndex() int {
	return list.Length() - 1
}

//Pop remove the item at the given position in the list, and return it.
//If no index is specified, a.Pop() removes and returns the last item in the list.
func (list *List) Pop(idx ...int) (interface{}, error) {
	index := list.getLastIndex()
	ll := len(idx)

	if ll > 1 {
		return nil, fmt.Errorf("only 1 or 0 arguments are allowed")
	}

	// in case of `list.Pop()`
	element := []interface{}(*list)[index]
	if ll == 0 {
		return element, list.removeByIndex(index)
	}

	if idx[0] > index {
		return nil, fmt.Errorf("index out of range")
	}

	index = idx[0]
	return element, list.removeByIndex(index)
}

//Clear remove all items from the list
func (list *List) Clear() {
	*list = NewList()
}

//Count Return the number of times x appears in the list.
func (list List) Count(element interface{}) int {
	count := 0
	for _, v := range list {
		if v == element {
			count++
		}
	}
	return count
}

//Copy Return a shallow copy of the list
func (list List) Copy() List {
	return list
}

//RandomChoice Return a random element from the list
func (list List) RandomChoice() interface{} {
	rand.Seed(time.Now().Unix())
	return list[rand.Intn(len(list))]
}

//RandomChoices Return a length sized list of elements chosen from the population with replacement.
func (list List) RandomChoices(length int) []interface{} {
	var result []interface{}
	for i := 0; i < length; i++ {
		result = append(result, list.RandomChoice())
	}
	return result
}

//Shuffle shuffle the list
func (list List) Shuffle() {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
}

//RandomSample Return a count length list of unique elements chosen from the list. Used for random sampling without replacement.
func (list List) RandomSample(count int) List {
	newList := make(List, list.Length())
	copy(newList, list)
	for i := 0; i < count; i++ {
		j := rand.Intn(i + 1)
		newList[i], newList[j] = newList[j], newList[i]
	}
	return newList
}

//Extend extend the list by appending all the items from lst.
func (list *List) Extend(lst List) {
	list_ := []interface{}(*list)
	*list = append(list_, lst...)
}

//Sum start and the items of a list from left to right and returns the total.
func (list List) Sum() int {
	var sum int
	for _, val := range list {
		intVal, ok := val.(int)
		if !ok {
			continue
		}
		sum += intVal
	}
	return sum
}

//ToString return string of list
func (list List) ToString() string {
	str := "List("
	for _, val := range list {
		str += fmt.Sprintf("%v", val)
	}
	return str + ")"
}

//Equal check if other have same element
func (list List) Equal(other List) bool {
	if other.Length() != list.Length() {
		return false
	}
	for idx, val := range list {
		if other[idx] != val {
			return false
		}
	}
	return true
}

//Iter return ListIter struct for iterate list
func (list List) Iter() ListIter {
	return ListIter{
		List: &list,
	}
}

//Contains Check if the item exists
func (list List) Contains(item interface{}) bool {
	if list.Count(item) == 0 {
		return false
	}
	return true
}
