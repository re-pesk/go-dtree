package datatree

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type DTHandler struct{
	DTree
	DirName string
	FileName string
	FileContent []byte
}

func (handler *DTHandler) SetFileName(fullName string) (err error) {
	if fullName = strings.Trim(fullName, " "); fullName == "" {
		err = fmt.Errorf("DTHandler.SetFileName.fullName is empty!")
		return
	}
	dir, _ := filepath.Split(fullName)
	handler.DirName = filepath.ToSlash(strings.Trim(dir, " "))
	handler.FileName = filepath.ToSlash(strings.Trim(fullName, " "))
	return
}

func (handler *DTHandler) ReadBytes(bytes []byte) (err error) {
	if len(bytes) < 1 {
		err = fmt.Errorf("DTHandler.ReadBytes.bytes is empty!")
		return
	}
	handler.FileContent = bytes
	return
}

func (handler *DTHandler) ReadFile() (err error) {
	if handler.FileName == "" {
		err = fmt.Errorf("DTHandler.FileName is empty!")
		return
	}
	handler.FileContent, err = ioutil.ReadFile(handler.FileName)
	return
}

func (handler *DTHandler) WriteFile() (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
