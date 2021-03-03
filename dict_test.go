package pythonlike_go

import (
	"fmt"
	"os"
	"testing"
)

var (
	ma Dict
)

func init() {
	ma = NewDict()
	ma["hoge"] = "hage"
	ma["fifi"] = "fasjfa"
}

func TestMain(m *testing.M) {
	code := m.Run()
	fmt.Println(ma)
	os.Exit(code)
}

func TestMapLength(t *testing.T) {
	length := ma.Length()
	fmt.Println(length)
}
func TestMapKeys(t *testing.T) {
	keys := ma.Keys()
	fmt.Println(keys)
}
func TestMapIsKeyExist(t *testing.T) {
	exist := ma.IsKeyExist("hoge")
	fmt.Println(exist)
}
func TestMapGet(t *testing.T) {
	get := ma.Get("fasdfa", nil)
	fmt.Println(get)
}
func TestMapPop(t *testing.T) {
	pop := ma.Pop("hoge", nil)
	fmt.Println(pop)
}
func TestMapPopItem(t *testing.T) {
	item, i := ma.PopItem()
	fmt.Println(item, i)
}
func TestMapSetDefault(t *testing.T) {
	ma.SetDefault("kasu", Dict{"hi": "f"})
}

func TestMapUpdate(t *testing.T) {
	ma.Update(Dict{"test": "gogo"})
}

func TestMapValues(t *testing.T) {
	ma.Values()
}
