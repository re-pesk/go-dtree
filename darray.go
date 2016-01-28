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

func (tree *DArray) CheckFirstKey(path string) (firstKey string, restPath string, err error){
	firstKey, restPath, err = ProcessPath(path)
	if err != nil{
		return
	}
	_, _, err = tree.GetValue(firstKey)
	if err != nil {
		firstKey = ""
		restPath = path
	}
	return
}

func (tree *DArray) UpdateValue(path string, newValue interface{}) (restPath string, value interface{}, err error){
	var firstKey string
	firstKey, restPath, err = tree.CheckFirstKey(path)
	if err != nil {
		restPath = path
		return
	}
	key, _ := strconv.Atoi(firstKey)
	if restPath == "" {
		tree.Value[key] = newValue
		value = tree.Value[key]
	} else {
		temp := DTree{tree.Value[key]}
		restPath, value, err = temp.UpdateValue(restPath, newValue)
	}
	return
}
