package datatree

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

const BREAKPOINT = ""

type DTreeHandler struct{
	DTree
	DirName string
	FileName string
	FileContent []byte
	Decode func() error
}

/* func (tree *DTree) Get(path string) (restPath string, value interface{}, err error){
	restPath, value, err = tree.GetValue(path)
	return
}

func (tree *DTree) Set(path string, newValue interface{}) (restPath string, value interface{}, err error){
	restPath, value, err = tree.SetValue(path, newValue)
	return
}

func (tree *DTree) Update(path string, newValue interface{}) (restPath string, value interface{}, err error){
	restPath, value, err = tree.UpdateValue(path, newValue)
	return
}

func (tree *DTree) Add(path string, newValue interface{}) (restPath string, value interface{}, err error){
	restPath, value, err = tree.AddValue(path, newValue)
	return
}
 */
 
func (tree *DTreeHandler) Get(path string) (restPath string, value interface{}, err error){
	restPath, value, err = tree.GetValue(path)
	return
}

func (tree *DTreeHandler) Set(path string, newValue interface{}) (restPath string, value interface{}, err error){
	restPath, value, err = tree.SetValue(path, newValue)
	return
}

func (tree *DTreeHandler) Update(path string, newValue interface{}) (restPath string, value interface{}, err error){
	restPath, value, err = tree.UpdateValue(path, newValue)
	return
}

func (tree *DTreeHandler) Add(path string, newValue interface{}) (restPath string, value interface{}, err error){
	restPath, value, err = tree.AddValue(path, newValue)
	return
}
 
func (handler *DTreeHandler) SetFileName(fullName string) (err error) {
	if fullName = strings.Trim(fullName, " "); fullName == "" {
		err = fmt.Errorf("DTreeHandler.SetFileName.fullName is empty!")
		return
	}
	dir, _ := filepath.Split(fullName)
	handler.DirName = filepath.ToSlash(strings.Trim(dir, " "))
	handler.FileName = filepath.ToSlash(strings.Trim(fullName, " "))
	return
}

func (handler *DTreeHandler) ReadBytes(bytes []byte) (err error) {
	if len(bytes) < 1 {
		err = fmt.Errorf("DTreeHandler.ReadBytes.bytes is empty!")
		return
	}
	handler.FileContent = bytes
	return
}

func (handler *DTreeHandler) ReadFile(fullName string) (err error) {
	err = handler.SetFileName(fullName)
	if err != nil || handler.FileName == "" {
		err = fmt.Errorf("DTreeHandler.FileName is empty!")
		return
	}
	handler.FileContent, err = ioutil.ReadFile(handler.FileName)
	return
}

func (handler *DTreeHandler) WriteFile(fullName string) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
