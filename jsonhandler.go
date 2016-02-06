package datatree

import (
	"fmt"
	"encoding/json"
)

type JsonHandler struct{
	DTreeHandler
}

func (tree *JsonHandler) Decode() (err error) {
	if len(tree.FileContent) == 0 {
		err = fmt.Errorf("JsonHandler.FileContent is empty!")
		return
	}
	err = json.Unmarshal(tree.FileContent, &tree.Value)
	return
}

func (tree *JsonHandler) Encode() (err error) {
	tree.FileContent, err = json.MarshalIndent(tree.Value, "", "\t")
	return
}
