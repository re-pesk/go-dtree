package datatree

import (
	"fmt"
	"strconv"
)


type DArray struct {
	Value []interface{}
}

func (arr *DArray) GetValue(path string) (restPath string, value interface{}, err error){
	var firstKey string
	firstKey, restPath, err = ProcessPath(path)
	if err != nil{
		restPath = path
		return
	}
	var key int
	key, err = strconv.Atoi(firstKey)
	if err != nil {
		restPath = path
		err = fmt.Errorf("Index of array must be integer!")
		return
	}
	if len(arr.Value) <= key {
		restPath = path
		err = fmt.Errorf("Index \"%d\" of array is out of range!", key)
		return
	}

	value = arr.Value[key]
	if restPath != "" {
		temp := DTree{value}
		restPath, value, err = temp.GetValue(restPath)
	}
	return
}

func (tree *DArray) SetValue(path string, newValue interface{}) (restPath string, value interface{}, err error){
	var firstKey string
	firstKey, restPath, err = ProcessPath(path)
	if err != nil {
		restPath = path
		return
	}
	var key int
	treeValLen := len(tree.Value)
	if firstKey == "+" {
		key, err = treeValLen, nil
	} else {
		key, err = strconv.Atoi(firstKey)
	}
	if err != nil {
		restPath = path
		err = fmt.Errorf("Index of array must be integer!")
		return
	}

	_, _, err = tree.GetValue(firstKey)
	if err != nil {
		for i:= treeValLen; i<= key; i++ {
			tree.Value = append(tree.Value, nil)
		}
		err = nil
	}
	if restPath == "" {
		tree.Value[key] = newValue
		value = tree.Value[key]
	} else {
		temp := DTree{tree.Value[key]}
		restPath, value, err = temp.SetValue(restPath, newValue)
		tree.Value[key] = temp.Value
	}
	return
}

func (tree *DArray) UpdateValue(path string, newValue interface{}) (restPath string, value interface{}, err error){
	if path == "" {
		err = fmt.Errorf("Index of array cannot be empty!")
		restPath = path
		return
	}
	var firstKey string
	firstKey, restPath, err = ProcessPath(path)
	key, err := strconv.Atoi(firstKey)
	if err != nil {
		restPath = path
		err = fmt.Errorf("Index of array must be integer!")
		return
	}
	_, _, err = tree.GetValue(firstKey)
	if err != nil {
		restPath = path
		return
	}
	if restPath == "" {
		tree.Value[key] = newValue
		value = tree.Value[key]
	} else {
		temp := DTree{tree.Value[key]}
		restPath, value, err = temp.UpdateValue(restPath, newValue)
	}
	return
}

func (tree *DArray) AddValue(path string, newValue interface{}) (restPath string, value interface{}, err error){
	if path == "" {
		err = fmt.Errorf("DArray.AddValue.path is empty!")
		return
	}
	var firstKey string
	firstKey, restPath, err = ProcessPath(path)
	key, err := strconv.Atoi(firstKey)
	if err != nil {
		restPath = path
		err = fmt.Errorf("Index of array must be integer!")
		return
	}
	_, _, err = tree.GetValue(firstKey)
	if restPath == "" {
		if err != nil {
			for i:= len(tree.Value); i<= key; i++ {
				tree.Value = append(tree.Value, nil)
			}
		}
		tree.Value[key] = newValue
		value = tree.Value[key]
	} else {
		temp := DTree{tree.Value[key]}
		restPath, value, err = temp.AddValue(restPath, newValue)
	}
	return
}
