package dtree

import (
	//"encoding/xml"
	"fmt"
	"github.com/clbanning/mxj"
)

type XMLHandler struct{
	DTreeHandler
}


func (tree *XMLHandler) Decode() (err error) {
	if len(tree.FileContent) == 0 {
		err = fmt.Errorf("XMLHandler.FileContent is empty!")
		return
	}
	//err = xml.Unmarshal(tree.FileContent, &tree.Value)
	tree.Value, err = mxj.NewMapXml(tree.FileContent, true)
	return
}
