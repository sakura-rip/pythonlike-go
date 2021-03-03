package pythonlike_go

import (
	"fmt"
	"testing"
)

func TestList_Insert(t *testing.T) {
	list := List{"hoge", "hage"}
	err := list.Insert(2, "hi")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(list)
}

func TestList_Insert2(t *testing.T) {
	list := List{"hoge", "hage"}
	err := list.Insert(-1, "hi")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(list)
}

func TestList_Insert3(t *testing.T) {
	list := List{"hoge", "hage"}
	err := list.Insert(0, "hi")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(list)
}

func TestList_Pop(t *testing.T) {
	list := List{"hoge", "hage"}
	v, err := list.Pop()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
	fmt.Println(list)
}

func TestList_Pop2(t *testing.T) {
	list := List{"hoge", "hage"}
	v, err := list.Pop(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
	fmt.Println(list)
}

func TestList_Clear(t *testing.T) {
	list := List{"hoge", "hage"}
	list.Clear()
	fmt.Println(list)
}

func TestList_Copy(t *testing.T) {
	list := List{"hoge", "hage"}
	fmt.Println(list)
	newList := list.Copy()
	fmt.Println(newList)
	newList.Clear()
	fmt.Println(newList)
}
