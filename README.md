# pythonlike-go
Touch the python-like Go

# WHAT IS THIS??

This is an attempt to do Python-like object manipulation in Golang.
これはGolangでPythonのようなオブジェクト操作をしようという試みです。

Golang's implementation of map, slice, etc. is poor and not very user-friendly
Golangのmapやsliceなどの実装は、貧相で、使い勝手がよくないです

This library allows you to manipulate slices and maps in a Python like way to ease the manipulation of maps and lists
このライブラリは、mapやsliceなどの操作を楽にするためにPythonのようにListやMapを操作できるものです

For example, in Golang, to remove something from a list, you would need code like the following
例えば、Golangではlistからものを削除するために下のようなコードが必要です

```go
list = append(list[:idx], list[idx+1:]...)
```
Furthermore, this code only allows you to change the value from index
さらに、このコードではindexからの値の変更しかできず、
```py
list.remove("item")
```
It does not allow flexible List operations like python the one above.
といった柔軟なList操作ができません。

This is where this library comes in.
そこで、このライブラリの出番です。
```go
list.Remove("item")
```
このように、PythonのようにListなどを操作することができます。
In this way, you can manipulate Lists and other objects as in Python.
