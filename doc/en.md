
This is an attempt to do Python-like object manipulation in Golang.  

Golang's implementation of map, slice, etc. is poor and not very user-friendly  


This library allows you to manipulate slices and maps in a Python like way to ease the manipulation of maps and lists  


For example, in Golang, to remove something from a list, you would need code like the following  


```go
list = append(list[:idx], list[idx+1:]...)
```
Furthermore, this code only allows you to change the value from index  

```py
list.remove("item")
```
It does not allow flexible List operations like python the one above.  

This is where this library comes in.   

```go
list.Remove("item")
```
In this way, you can manipulate Lists and other objects as in Python.  
