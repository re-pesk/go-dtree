package datatree

import (
	"fmt"
)

type DMap struct {
	Value map[string]interface{}
}

func (tree *DMap) Get(path string) (restPath string, value interface{}, err error){
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
			restPath, value, err = temp.Get(restPath)
		}
	} else {
		restPath = path
		err = fmt.Errorf("Map has no element with key \"%s\"!", firstKey)
	}
	return
}

func (tree *DMap) Set(path string, newValue interface{}) (restPath string, value interface{}, err error){
	var firstKey string
	firstKey, restPath, err = ProcessPath(path)
/* 	if firstKey == "+" {
		err = fmt.Errorf("Map key \"+\" is wrong!")
	}
 */	if err != nil {
		restPath = path
		return
	}
	if restPath == "" {
		tree.Value[firstKey] = newValue
		value = tree.Value[firstKey]
	} else {
		_, _, err = tree.Get(firstKey)
		if err != nil {
			tree.Value[firstKey] = nil
		}
		temp := DTree{tree.Value[firstKey]}
		restPath, value, err = temp.Set(restPath, newValue)
		tree.Value[firstKey] = temp.Value
	}
	return
}




func (tree *DMap) Update(path string, newValue interface{}) (restPath string, value interface{}, err error){
	if path == "" {
		err = fmt.Errorf("Map cannot have value with key \"\"!")
		restPath = path
		return
	}
	var firstKey string
	firstKey, restPath, err = ProcessPath(path)
	_, _, err = tree.Get(firstKey)
	if err != nil {
		restPath = path
		return
	}
	if restPath == "" {
		tree.Value[firstKey] = newValue
		value = tree.Value[firstKey]
	} else {
		temp := DTree{tree.Value[firstKey]}
		restPath, value, err = temp.Update(restPath, newValue)
	}
	return
}



func (tree *DMap) Add(path string, newValue interface{}) (restPath string, value interface{}, err error){
	if path == "" {
		err = fmt.Errorf("Map cannot have value with key \"\"!")
		restPath = path
		return
	}
	var firstKey string
	firstKey, restPath, err = ProcessPath(path)
	_, _, err = tree.Get(firstKey)
	if restPath == "" {
		if err == nil {
			err = fmt.Errorf("Map already has value with key %s!", firstKey)
			restPath = path
			return
		}
		tree.Value[firstKey] = newValue
		value = tree.Value[firstKey]
	} else {
		if err != nil {
			tree.Value[firstKey] = nil
			value = tree.Value[firstKey]
		}
		temp := DTree{tree.Value[firstKey]}
		restPath, value, err = temp.Add(restPath, newValue)
	}
	return
}

