Listオブジェクトを作るためには

```go
list := NewListFrom("test", 1, "hoge") //list = ["test", 1, "hoge"]
list2 := NewList() //list=[]
list3 := NewListFromClice([]string{"test", "1", "hoge"}) //list=["test", "1", "hoge"]
```

のようにして生成できます。

```go
list.Index("hoge") //list.index("hoge")
// Return 1

list.Append("fuga") //list.append("fuga")
// list: ["test", 1, "hoge", "fuga"]

list.Remove(1) //list.remove(1)
// list: ["test", "hoge", "fuga"]

list.Length() //len(list)
// Return 3

list.Insert(1, "hage")  //list.insert(1, "hage")
// list: ["test", "hage", "hoge", "fuga"]

list.Pop() //list.pop()
// Return: fuga, list: ["test", "hage", "hoge"]

list.Pop(2) //list.pop(2)
// Return: hoge, list: ["test", "hage"]

list.Clear() //list.clear()
// list: []


list3.Copy() //list.copy()
// Return: ["test", "1", "hoge"]

list3.RandomChoice() //random.choice(list3)
// Return: "test" | "1" | "hoge"

list3.RandomChoices(2) // random.choices(list3, k=2)
// Return: [random 2 value]

list3.Shuffle()
// list3: ["1", "test" "hoge"]
```
