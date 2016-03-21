package dtree

import (
	"fmt"
)

type NMap map[string]interface{}

type DMap struct {
	Value map[string]interface{}
}

func (tree *DMap) Get(path string) (result DTree){
	var firstKey string
	firstKey, result.RestPath, result.Error = ProcessPath(path)
	if result.Error != nil{
		result.RestPath = path
		return
	}
	var succ bool ; result.Value, succ = tree.Value[firstKey]
	if succ {

		if result.RestPath != "" {
			result = result.Get(result.RestPath)
		}
		if result.UsedPath != "" {
			result.UsedPath = firstKey + "." + result.UsedPath
		} else {
			result.UsedPath = firstKey
		}
	} else {
		result.RestPath = path
		result.Error = fmt.Errorf("Map has no element with key \"%s\"!", firstKey)
	}
	return
}

func (tree *DMap) Set(path string, newValue interface{}) (result DTree){
	var firstKey string
	firstKey, result.RestPath, result.Error = ProcessPath(path)
	if result.Error != nil {
		result.RestPath = path
		return
	}
	if result.RestPath == "" {
		result.Value = newValue
		result.UsedPath = firstKey
		tree.Value[firstKey] = result.Value
	} else {
		result.Error = tree.Get(firstKey).Error
		if result.Error != nil {
			tree.Value[firstKey] = nil
		}
		var temp DTree
		temp.Value = tree.Value[firstKey]
		result = temp.Set(result.RestPath, newValue)
		if result.UsedPath != "" {
			result.UsedPath = firstKey + "." + result.UsedPath
		} else {
			result.UsedPath = firstKey
		}
		tree.Value[firstKey] = temp.Value

	}
	return
}

