package datatree

import (
	"fmt"
	"strconv"
	"strings"
)

type DTree struct{
	Value interface{}
}

func (tree *DTree) GetValue(path string) (restPath string, value interface{}, err error){
	path = strings.Trim(path, " ") 
	if path == "" {
		value = tree.Value
		return
	}
	restPath = path
	switch typedVal := tree.Value.(type){
	case map[string]interface{}:
		temp := DMap{typedVal}
		restPath, value, err = temp.GetValue(path)
	case []interface{}:
		temp := DArray{typedVal}
		restPath, value, err = temp.GetValue(path)
	default:
		err = fmt.Errorf("Value of type \"%T\" has no index!", typedVal)
	}
	return
}

func (tree *DTree) SetValue(path string, newValue interface{}) (restPath string, value interface{}, err error){
	path = strings.Trim(path, " ")
	if path == "" {
		tree.Value = newValue
		value = tree.Value
		return
	}
	restPath = path
	var firstKey string
	firstKey, _, err = ProcessPath(path)
	switch tree.Value.(type){
	case map[string]interface{}:
	case []interface{}:
	default:
		if err != nil {
			return
		}
		_, err = strconv.Atoi(firstKey)
		if err == nil || firstKey == "+" {
			tree.Value = []interface{}{}
		} else {
			tree.Value = map[string]interface{}{}
		}
	}
	switch typedVal := tree.Value.(type){
	case map[string]interface{}:
		temp := DMap{typedVal}
		restPath, value, err = temp.SetValue(path, newValue)
		tree.Value = temp.Value
	case []interface{}:
		temp := DArray{typedVal}
		restPath, value, err = temp.SetValue(path, newValue)
		tree.Value = temp.Value
	default:
		err = fmt.Errorf("Value of type \"%T\" has no index!", typedVal)
	}
	return
}

func (tree *DTree) UpdateValue(path string, newValue interface{}) (restPath string, value interface{}, err error){
	if path = strings.Trim(path, " "); path == "" {
		tree.Value = newValue
		value = tree.Value
		return
	}
	restPath = path
	switch typedVal := tree.Value.(type){
	case map[string]interface{}:
		temp := DMap{typedVal}
		restPath, value, err = temp.UpdateValue(path, newValue)
	case []interface{}:
		temp := DArray{typedVal}
		restPath, value, err = temp.UpdateValue(path, newValue)
	default:
		err = fmt.Errorf("Value of type \"%T\" has no index!", typedVal)
	}
	return
}

func (tree *DTree) AddValue(path string, newValue interface{}) (restPath string, value interface{}, err error){
	if path = strings.Trim(path, " "); path == "" {
		tree.Value = newValue
		value = tree.Value
		return
	}
	restPath = path
	switch typedVal := tree.Value.(type){
	case map[string]interface{}:
		temp := DMap{typedVal}
		restPath, value, err = temp.AddValue(path, newValue)
	case []interface{}:
		temp := DArray{typedVal}
		restPath, value, err = temp.AddValue(path, newValue)
	default:
		err = fmt.Errorf("Value of type \"%T\" has no index!", typedVal)
	}
	return
}

