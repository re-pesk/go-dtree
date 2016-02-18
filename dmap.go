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
		tree.Value[firstKey] = result.Value
	} else {
		result.Error = tree.Get(firstKey).Error
		if result.Error != nil {
			tree.Value[firstKey] = nil
		}
		var temp DTree
		temp.Value = tree.Value[firstKey]
		result = temp.Set(result.RestPath, newValue)
		tree.Value[firstKey] = temp.Value

	}
	return
}




func (tree *DMap) Update(path string, newValue interface{}) (result DTree){
	if path == "" {
		result.Error = fmt.Errorf("Map cannot have value with key \"\"!")
		result.RestPath = path
		return
	}
	var firstKey string
	firstKey, result.RestPath, result.Error = ProcessPath(path)
	result.Error = tree.Get(firstKey).Error
	if result.Error != nil {
		result.RestPath = path
		return
	}
	if result.RestPath == "" {
		result.Value = newValue
		tree.Value[firstKey] = result.Value
	} else {
		var temp DTree
		temp.Value = tree.Value[firstKey]
		result = temp.Update(result.RestPath, newValue)
		tree.Value[firstKey] = temp.Value
	}
	return
}



func (tree *DMap) Add(path string, newValue interface{}) (result DTree){
	if path == "" {
		result.Error = fmt.Errorf("Map cannot have value with key \"\"!")
		result.RestPath = path
		return
	}
	var firstKey string
	firstKey, result.RestPath, result.Error = ProcessPath(path)
	result.Error = tree.Get(firstKey).Error
	if result.RestPath == "" {
		if result.Error == nil {
			result.Error = fmt.Errorf("Map already has value with key %s!", firstKey)
			result.RestPath = path
			return
		}
		result.Value = newValue
		tree.Value[firstKey] = result.Value
	} else {
		if result.Error != nil {
			tree.Value[firstKey] = nil
		}
		var temp DTree
		temp.Value = tree.Value[firstKey]
		result = temp.Add(result.RestPath, newValue)
		tree.Value[firstKey] = temp.Value
	}
	return
}

