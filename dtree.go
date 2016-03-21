package dtree

import (
	"fmt"
	"strconv"
	"strings"
)

type NTree interface{}

type DTree struct{
	Value interface{}
	UsedPath string
	RestPath string
	Error error
}

func (tree *DTree) Get(path string) (result DTree) {
	path = strings.Trim(path, " ") 
	if path == "" {
		result.Value = tree.Value
		return
	}
	result.RestPath = path
	switch typedVal := tree.Value.(type){
	case map[string]interface{}:
		temp := DMap{typedVal}
		result = temp.Get(path)
	case []interface{}:
		temp := DArray{typedVal}
		result = temp.Get(path)
	default:
		result.Error = fmt.Errorf("Value of type \"%T\" has no index!", typedVal)
	}
	return
}

func (tree *DTree) Set(path string, newValue interface{}) (result DTree){
	path = strings.Trim(path, " ")
	if path == "" {
		tree.Value = newValue
		result.Value = tree.Value
		return
	}
	result.RestPath = path
	var firstKey string
	firstKey, _, result.Error = ProcessPath(path)
	switch tree.Value.(type){
	case map[string]interface{}, []interface{} :
	default:
		if result.Error != nil {
			return
		}
		_, result.Error = strconv.Atoi(firstKey)
		if result.Error == nil || firstKey == "+" {
			tree.Value = []interface{}{}
		} else {
			tree.Value = map[string]interface{}{}
		}
	}
	switch typedVal := tree.Value.(type){
	case map[string]interface{}:
		temp := DMap{typedVal}
		result = temp.Set(path, newValue)
		tree.Value = temp.Value
	case []interface{}:
		temp := DArray{typedVal}
		result = temp.Set(path, newValue)
		tree.Value = temp.Value
	default:
		result.Error = fmt.Errorf("Value of type \"%T\" has no index!", typedVal)
	}
	return
}

func (tree *DTree) Update(path string, newValue interface{}) (result DTree){
	path = strings.Trim(path, " ")
	tmp := tree.Get(path)
	if tmp.Error == nil {
		result = tree.Set(path, newValue)
	} else {
		result.Error = fmt.Errorf("Element with key \"%v\" doesn't exist!", path)
	}
	return
}

func (tree *DTree) Add(path string, newValue interface{}) (result DTree){
	path = strings.Trim(path, " ")
	tmp := tree.Get(path)
	if tmp.Error != nil {
		result = tree.Set(path, newValue)
	} else {
		result.UsedPath = path
		result.Error = fmt.Errorf("Element with key \"%v\" already exists!", path)
	}
	return
}
