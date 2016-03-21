package dtree

import (
	"fmt"
	"encoding/json"
)

type JsonHandler struct{
	DTreeHandler
}

func (handler *JsonHandler) Decode() (err error) {
	if len(handler.FileContent) == 0 {
		err = fmt.Errorf("JsonHandler.FileContent is empty!")
		return
	}
	err = json.Unmarshal(handler.FileContent, &handler.Value)
	if err != nil {
		err = fmt.Errorf(
			"JsonHandler.Decode(): json.Unmarshal() error decoding file \"%s\":\n\n  %s", 
			handler.FileName, 
			err.Error(),
		)
	}
	return
}

func (handler *JsonHandler) Encode() (err error) {
	handler.FileContent, err = json.MarshalIndent(handler.Value, "", "\t")
	return
}

func (handler *JsonHandler) ToValue(jsonString string) (result interface{}){
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil && jsonString != "" {
		panic(err)
	}
	return
}