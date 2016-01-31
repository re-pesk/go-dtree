package datatree

import (
	"fmt"
	"encoding/json"
)

type JsonHandler struct{
	DTHandler
}

func (tree *JsonHandler) Decode() (err error) {
	if len(tree.FileContent) == 0 {
		err = fmt.Errorf("JsonHandler.FileContent is empty!")
		return
	}
	err = json.Unmarshal(tree.FileContent, &tree.Value)
	return
}
