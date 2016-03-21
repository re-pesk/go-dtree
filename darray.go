package dtree

import (
	"fmt"
	"strconv"
)

type NArray []interface{}

type DArray struct {
	Value []interface{}
}

func (tree *DArray) Get(path string) (result DTree){
	var firstKey string
	firstKey, result.RestPath, result.Error = ProcessPath(path)
	if result.Error != nil{
		result.RestPath = path
		return
	}
	var key int
	key, result.Error = strconv.Atoi(firstKey)
	if result.Error != nil {
		result.RestPath = path
		result.Error = fmt.Errorf("Index of array must be integer!")
		return
	}
	if len(tree.Value) <= key {
		result.RestPath = path
		result.Error = fmt.Errorf("Index \"%d\" is out of range of the array!", key)
		return
	}

	result.Value = tree.Value[key]
	if result.RestPath != "" {
		result = result.Get(result.RestPath)
	}
	if result.UsedPath != "" {
		result.UsedPath = firstKey + "." + result.UsedPath
	} else {
		result.UsedPath = firstKey
	}
	return
}

func (tree *DArray) Set(path string, newValue interface{}) (result DTree){
	var firstKey string
	firstKey, result.RestPath, result.Error = ProcessPath(path)
	if result.Error != nil {
		result.RestPath = path
		return
	}
	var key int
	treeValLen := len(tree.Value)
	if firstKey == "+" {
		key, result.Error = treeValLen, nil
	} else {
		key, result.Error = strconv.Atoi(firstKey)
	}
	if result.Error != nil {
		result.RestPath = path
		result.Error = fmt.Errorf("Index of array must be integer!")
		return
	}

	result.Error = tree.Get(firstKey).Error
	
	if result.Error != nil {
		for i:= treeValLen; i<= key; i++ {
			tree.Value = append(tree.Value, nil)
		}
		result.Error = nil
	}
	if result.RestPath == "" {
		result.Value = newValue
		result.UsedPath = firstKey
		tree.Value[key] = result.Value
	} else {
		var temp DTree
		temp.Value = tree.Value[key]
		result = temp.Set(result.RestPath, newValue)
		if result.UsedPath != "" {
			result.UsedPath = firstKey + "." + result.UsedPath
		} else {
			result.UsedPath = firstKey
		}
		tree.Value[key] = temp.Value
	}
	return
}

