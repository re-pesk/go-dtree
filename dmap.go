package datatree

import (
	"fmt"
)

type DMap struct {
	Value map[string]interface{}
}

func (tree *DMap) GetValue(path string) (restPath string, value interface{}, err error){
	var firstKey string
	firstKey, restPath, err = ProcessPath(path)
	if err != nil{
		restPath = path
		return
	}
	value, succ := tree.Value[firstKey]
	if succ {
		if restPath != "" {
			temp := DTree{value}
			restPath, value, err = temp.GetValue(restPath)
		}
	} else {
		restPath = path
		err = fmt.Errorf("Map has no element with key \"%s\"!", firstKey)
	}
	return
}

func (tree *DMap) CheckFirstKey(path string) (firstKey string, restPath string, err error){
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

func (tree *DMap) UpdateValue(path string, newValue interface{}) (restPath string, value interface{}, err error){
	var firstKey string
	firstKey, restPath, err = tree.CheckFirstKey(path)
	if err != nil {
		restPath = path
		return
	}
	if restPath == "" {
		tree.Value[firstKey] = newValue
		value = tree.Value[firstKey]
	} else {
		temp := DTree{tree.Value[firstKey]}
		restPath, value, err = temp.UpdateValue(restPath, newValue)
	}
	return
}
