package pythonlike_go

import (
	"fmt"
	"testing"
)

func TestList_Insert(t *testing.T) {
	list := List{"hoge", "hage"}
	err := list.Insert(3, "hi")
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
