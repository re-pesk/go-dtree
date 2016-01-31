package datatree

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type DTreeHandler struct{
	DTree
	DirName string
	FileName string
	FileContent []byte
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

func (handler *DTreeHandler) ReadFile() (err error) {
	if handler.FileName == "" {
		err = fmt.Errorf("DTreeHandler.FileName is empty!")
		return
	}
	handler.FileContent, err = ioutil.ReadFile(handler.FileName)
	return
}

func (handler *DTreeHandler) WriteFile() (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
