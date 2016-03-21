package main

import (
	"github.com/re-pe/dtree"
	"fmt"
)

func main(){

	// Reading json file:
	fmt.Println("\nReading json file:\n")
	var tree dtree.JsonHandler
	if err := tree.ReadFile("start.conf"); err != nil {
			panic(err)
	}
	fmt.Printf("File \"%v\" exists and was read.\n\n", tree.FileName)
	if err := tree.Decode(); err == nil {
			fmt.Println("tree.Value:", tree.Value)
	}

	// Reading json string:
	fmt.Println("\nReading json string:\n")
	tree.Content = []byte(`{"Other" : {"a" : 0, "b" : 1, "c" : 2, "i0" : 3}}`)
	if err := tree.Decode(); err == nil {
			fmt.Println("tree.Value:", tree.Value)
	}

	// Getting value, used path, remained path and error when path is right:
	fmt.Println("\nGetting value when path is right:\n")

	result := tree.Get("Other.a")

	fmt.Printf(`
 tree.Get("Other.a"):
 result.Value: "%v"
 result.Used path: "%v"
 result.Remaining path: "%v"
 result.Error: %v
`, result.Value, result.UsedPath, result.RestPath, result.Error)

	// Getting value, used path, remained path and error when path is wrong:
	fmt.Println("\nGetting value when path is wrong:\n")

	result = tree.Get("Other.d")

	fmt.Printf(`
 tree.Get("Other.d"):
 result.Value: "%v"
 result.Used path: "%v"
 result.Remaining path: "%v"
 result.Error: %v
`, result.Value, result.UsedPath, result.RestPath, result.Error)

	//Setting value:

	result = tree.Set("Other.d", tree.NewValue(`"abc"`))

	fmt.Print("\n tree.Set(\"Other.d\", tree.NewValue(`\"abc\"`))")
	fmt.Printf(`
 result.Value: "%v"
 result.Used path: "%v"
 result.Remaining path: "%v"
 result.Error: %v
`, result.Value, result.UsedPath, result.RestPath, result.Error)
		
	result = tree.Set("NewMap.2.1", tree.NewValue(`"new_value"`))

	fmt.Print("\n tree.Set(\"NewMap.2.1\", tree.NewValue(`\"new_value\"`))")
	fmt.Printf(`
 result.Value: "%v"
 result.Used path: "%v"
 result.Remaining path: "%v"
 result.Error: %v
`, result.Value, result.UsedPath, result.RestPath, result.Error)
		
	result = tree.Set("NewMap.2.+", tree.NewValue(`"plus_value"`))

	fmt.Print("\n tree.Set(\"NewMap.2.+\", tree.NewValue(`\"plus_value\"`))")
	fmt.Printf(`
 result.Value: "%v"
 result.Used path: "%v"
 result.Remaining path: "%v"
 result.Error: %v
`, result.Value, result.UsedPath, result.RestPath, result.Error)

	fmt.Printf(`
 tree.Value: "%v"%v`, tree.Value, "\n" )

}

