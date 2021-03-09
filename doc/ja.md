
# WHAT IS THIS??

これはGolangでPythonのようなオブジェクト操作をしようという試みです。

Golangのmapやsliceなどの実装は、貧相で、使い勝手がよくないです

このライブラリは、mapやsliceなどの操作を楽にするためにPythonのようにListやMapを操作できるものです

例えば、Golangではlistからものを削除するために下のようなコードが必要です

```go
list = append(list[:idx], list[idx+1:]...)
```
さらに、このコードではindexからの値の変更しかできず、
```py
list.remove("item")
```
といったPythonのような柔軟なList操作ができません。

そこで、このライブラリの出番です。
```go
list.Remove("item")
```

このように、PythonのようにListなどを操作することができます。

[Usage](list_ja.md)