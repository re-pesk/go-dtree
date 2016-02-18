package dtree

import (
	"fmt"
	"strconv"
)

type NArray []interface{}

type DArray struct {
	Value []interface{}
}

func (arr *DArray) Get(path string) (result DTree){
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
	if len(arr.Value) <= key {
		result.RestPath = path
		result.Error = fmt.Errorf("Index \"%d\" of array is out of range!", key)
		return
	}

	result.Value = arr.Value[key]
	if result.RestPath != "" {
		//var temp DTree
		//temp.Value = result.Value
		result = result.Get(result.RestPath)
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
		tree.Value[key] = result.Value
	} else {
		var temp DTree
		temp.Value = tree.Value[key]
		result = temp.Set(result.RestPath, newValue)
		tree.Value[key] = temp.Value
	}
	return
}

func (tree *DArray) Update(path string, newValue interface{}) (result DTree){
	if path == "" {
		result.Error = fmt.Errorf("Index of array cannot be empty!")
		result.RestPath = path
		return
	}
	var firstKey string
	firstKey, result.RestPath, result.Error = ProcessPath(path)
	var key int ; key, result.Error = strconv.Atoi(firstKey)
	if result.Error != nil {
		result.RestPath = path
		result.Error = fmt.Errorf("Index of array must be integer!")
		return
	}
	result.Error = tree.Get(firstKey).Error
	if result.Error != nil {
		result.RestPath = path
		return
	}
	if result.RestPath == "" {
		result.Value = newValue
		tree.Value[key] = result.Value
	} else {
		var temp DTree
		temp.Value = tree.Value[key]
		result = temp.Update(result.RestPath, newValue)
		tree.Value[key] = temp.Value
	}
	return
}

func (tree *DArray) Add(path string, newValue interface{}) (result DTree){
	if path == "" {
		result.Error = fmt.Errorf("DArray.AddValue.path is empty!")
		return
	}
	var firstKey string
	firstKey, result.RestPath, result.Error = ProcessPath(path)
	var key int ; key, result.Error = strconv.Atoi(firstKey)
	if result.Error != nil {
		result.RestPath = path
		result.Error = fmt.Errorf("Index of array must be integer!")
		return
	}
	result.Error = tree.Get(firstKey).Error
	if result.RestPath == "" {
		if result.Error != nil {
			for i:= len(tree.Value); i<= key; i++ {
				tree.Value = append(tree.Value, nil)
			}
		}
		result.Value = newValue
		tree.Value[key] = result.Value
	} else {
		var temp DTree
		temp.Value = tree.Value[key]
		result = temp.Add(result.RestPath, newValue)
		tree.Value[key] = temp.Value
	}
	return
}
