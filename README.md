# DTree

Tool for working with semi-structured data like json or xml

### Usage

Reading json file:
```
var tree dtree.JsonHandler
if err = tree.ReadFile(confName); err != nil {
	panic(err)
}
fmt.Printf("File \"%v\" exists and was read.\n\n", appConf.FileName)
if err = tree.Decode(); err == nil {
    fmt.Println(tree.Value)
}
```

Decoding string:
```
var tree dtree.JsonHandler
tree.FileCOntent = `{"Map" : {"a" : 0, "b" : 1, "c" : 2, "i0" : 3}}``
if err = tree.Decode(); err == nil {
    fmt.Println(tree.Value)
}

```	

Getting value, used path, rest path and error:
```
result := tree.Get("Map.a")

fmt.Println("Value:", tree.Value)
fmt.Println("Used path", tree.UsedPath)
fmt.Println("Remain path:", tree.RestPath)
fmt.Println("Error:", tree.Error)

result := tree.Get("Map.d")

fmt.Println("Value:", tree.Value)
fmt.Println("Used path", tree.UsedPath)
fmt.Println("Remain path:", tree.RestPath)
fmt.Println("Error:", tree.Error)
```

Setting value:
```
result := tree.Set("Map.d", tree.NewValue(`"abc"`))

fmt.Println("Value:", tree.Value)
fmt.Println("Used path", tree.UsedPath)
fmt.Println("Remain path:", tree.RestPath)
fmt.Println("Error:", tree.Error)
```