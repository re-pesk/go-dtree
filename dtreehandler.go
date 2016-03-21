package dtree

import (
	"fmt"
	"io/ioutil"
//	"os"
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

func (handler *DTreeHandler) ReadFile(fullName string) (err error) {
	err = handler.SetFileName(fullName)
	if err != nil || handler.FileName == "" {
		err = fmt.Errorf("DTreeHandler.FileName is empty!")
		return
	}
	if pathExists(handler.FileName, isFile).Index < 2 {
		err = fmt.Errorf("File %s does not exist!", handler.FileName)
		return
	}
	handler.FileContent, err = ioutil.ReadFile(handler.FileName)
	return
}

func (handler *DTreeHandler) NewFile(fullName string) (err error) {
	if len(handler.FileContent) < 1 {
		err = fmt.Errorf("DTreeHandler.FileContent is empty!")
		return
	}
	if fullName == "" {
		err = fmt.Errorf("DTreeHandler.NewFile().fullName is empty!")
		return
	}
	err = ioutil.WriteFile(fullName, handler.FileContent, 0777)
	return
}

func (handler *DTreeHandler) WriteFile() (err error) {
	if len(handler.FileContent) < 1 {
		err = fmt.Errorf("DTreeHandler.FileContent is empty!")
		return
	}
	if handler.FileName == "" {
		err = fmt.Errorf("DTreeHandler.FileName is empty!")
		return
	}
	err = ioutil.WriteFile(handler.FileName, handler.FileContent, 0777)
	return
}
