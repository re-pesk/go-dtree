package datatree

import (
	"fmt"
	"strings"
)

type Value interface{}
type DTree struct{
	Value
}

func (tree *DTree) Get(path string) (restPath string, value interface{}, err error){
	restPath, value, err = tree.GetValue(path)
	return
}

func (tree *DTree) GetValue(path string) (restPath string, value interface{}, err error){
	if path = strings.Trim(path, " "); path == "" {
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

func (tree *DTree) Update(path string, newValue interface{}) (restPath string, value interface{}, err error){
	restPath, value, err = tree.UpdateValue(path, newValue)
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
