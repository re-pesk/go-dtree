package datatree

import (
	"fmt"
	"strings"
)

type I_Node interface{
	Get(string) (string, interface{}, error)
	Set(string, interface{}) (string, interface{}, error)
	Update(string, interface{}) (string, interface{}, error)
	Add(string, interface{}) (string, interface{}, error)
}

type I_DTreeHandler interface{
	I_Node
 	SetFileName(string) error
	ReadBytes([]byte) error
	ReadFile(string) error
	WriteFile(string) error
	Decode() error
}

func ProcessPath(path string) (firstKey string, restPath string, err error) {
	path = strings.Trim(path, " ")
	if path == "" {
		err = fmt.Errorf("Path cannot be empty!")
		return
	} 
	splitedPath := strings.SplitN(path, ".", 2)
	splitedPathLen := len(splitedPath)
	if splitedPathLen < 1 {
		err = fmt.Errorf("Path is wrong!")
		restPath = path
		return
	}
	firstKey = strings.Trim(splitedPath[0], " ")
	if splitedPathLen > 1 {
		restPath = strings.Trim(splitedPath[1], " ")
	}
	return
}