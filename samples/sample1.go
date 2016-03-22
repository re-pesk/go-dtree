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

	/* Output: 
	 File "start.conf" exists and was read.
	*/

	if err := tree.Decode(); err == nil {
			fmt.Println("tree.Value:", tree.Value)
	}

	/* Output: 
	 tree.Value: map[root:map[Map:map[a:0 b:1 c:2 i0:3] Arr:[a b c 0 1 2]]]
	*/

	// Reading json string:
	fmt.Println("\nReading json string:\n")
	tree.Content = []byte(`{"Other" : {"a" : 0, "b" : 1, "c" : 2, "i0" : 3}}`)
	if err := tree.Decode(); err == nil {
			fmt.Println("tree.Value:", tree.Value)
	}

	/* Output: 
	 tree.Value: map[Other:map[a:0 b:1 c:2 i0:3]]
	*/

	// Getting value, used path, remained path and error when path is right:
	fmt.Println("\nGetting value when path is right:\n")

	result := tree.Get("Other.a")

	fmt.Printf(`
 tree.Get("Other.a"):
 result.Value: %v
 result.UsedPath: "%v"
 result.RestPath: "%v"
 result.Error: %v
`, result.Value, result.UsedPath, result.RestPath, result.Error)

	/* Output: 
	 tree.Get("Other.a"):
	 result.Value: 0
	 result.UsedPath: "Other.a"
	 result.RestPath: ""
	 result.Error: <nil>
	*/


	// Getting value, used path, remained path and error when path is wrong:
	fmt.Println("\nGetting value when path is wrong:\n")

	result = tree.Get("Other.d")

	fmt.Printf(`
 tree.Get("Other.d"):
 result.Value: %v
 result.UsedPath: "%v"
 result.RestPath: "%v"
 result.Error: %v
`, result.Value, result.UsedPath, result.RestPath, result.Error)

	/* Output: 
	 tree.Get("Other.d"):
	 result.Value: <nil>
	 result.UsedPath: "Other"
	 result.RestPath: "d"
	 result.Error: Map has no element with key "d"!
	*/

	//Setting values:

	result = tree.Set("Other.d", tree.NewValue(`"abc"`))

	fmt.Print("\n tree.Set(\"Other.d\", tree.NewValue(`\"abc\"`))")
	fmt.Printf(`
 result.Value: %v
 result.UsedPath: "%v"
 result.RestPath: "%v"
 result.Error: %v
`, result.Value, result.UsedPath, result.RestPath, result.Error)

	/* Output: 
	 tree.Set("Other.d", tree.NewValue(`"abc"`)):
	 result.Value: abc
	 result.UsedPath: "Other.d"
	 result.RestPath: ""
	 result.Error: <nil>
	*/

	result = tree.Set("NewArr.2.1", tree.NewValue(`"new_value"`))

	fmt.Print("\n tree.Set(\"NewArr.2.1\", tree.NewValue(`\"new_value\"`))")
	fmt.Printf(`
 result.Value: %v
 result.UsedPath: "%v"
 result.RestPath: "%v"
 result.Error: %v
`, result.Value, result.UsedPath, result.RestPath, result.Error)

	/* Output: 
	 tree.Set("NewArr.2.1", tree.NewValue(`"new_value"`)):
	 result.Value: new_value
	 result.UsedPath: "NewArr.2.1"
	 result.RestPath: ""
	 result.Error: <nil>
	*/
		
	result = tree.Set("NewArr.2.+", tree.NewValue(`"plus_value"`))

	fmt.Print("\n tree.Set(\"NewArr.2.+\", tree.NewValue(`\"plus_value\"`))")
	fmt.Printf(`
 result.Value: %v
 result.UsedPath: "%v"
 result.RestPath: "%v"
 result.Error: %v
`, result.Value, result.UsedPath, result.RestPath, result.Error)

	/* Output: 
	 tree.Set("NewArr.2.+", tree.NewValue(`"plus_value"`)):
	 result.Value: plus_value
	 result.UsedPath: "NewArr.2.+"
	 result.RestPath: ""
	 result.Error: <nil>
	*/

	fmt.Printf(`
 Final tree.Value: "%v"
`, tree.Value, "\n" )

  /* Output: 
	 Final tree.Value: map[Other:map[a:0 b:1 c:2 i0:3 d:abc] NewArr:[<nil> <nil> [<nil> new_value plus_value]]]
	*/
}

