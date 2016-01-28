package datatree

import (
	"fmt"
	"strings"
)

func ProcessPath(path string) (firstKey string, restPath string, err error) {
	path = strings.Trim(path, " ")
	splitedPath := strings.SplitN(path, ".", 2)
	splitedPathLen := len(splitedPath)
	if splitedPathLen < 1 {
		err = fmt.Errorf("Path is wrong!")
		restPath = path
	}
	firstKey = strings.Trim(splitedPath[0], " ")
	if splitedPathLen > 1 {
		restPath = strings.Trim(splitedPath[1], " ")
	}
	return
}

