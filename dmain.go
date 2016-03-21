package dtree

import (
	"fmt"
	"os"
	"strings"
)

const (
	BREAKPOINT        = "breakpoint" //breakoint for modified godebug
	isDir             = true
	isFile            = false
)

type I_Node interface{
	Get(string) DTree
	Set(string, interface{}) DTree
	Update(string, interface{}) DTree
	Add(string, interface{}) DTree
}

type I_DTreeHandler interface{
	I_Node
 	SetFileName(string) error
	ReadBytes([]byte) error
	ReadFile(string) error
	WriteFile() error
	NewFile(string) error
	Decode() error
	Encode() error
}

type intResult struct {
	Index int
	Error error
}

func pathExists(path string, dir bool) (result intResult) {
	result.Index = -1
	path = strings.TrimSpace(path)
	if path == "" {
		result.Error = fmt.Errorf("PathExists.args.path is empty!")
		return
	}

	dirType := "Directory"
	if !dir {
		dirType = "File" 
	}

	finfo, err := os.Stat(path)
	if err != nil {
		// no such file or dir
		result.Index, result.Error = 0, fmt.Errorf("%s does not exist!", path)
		return
	}

	if dir == finfo.IsDir() { 
		result.Index = 2
	} else {
		result.Index = 1
	}

	if result.Index < 2 {
		result.Error = fmt.Errorf("%s is not %s!", path, strings.ToLower(dirType))
	}
	return
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

