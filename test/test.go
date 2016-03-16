package main

import (
	"dtree"
	"github.com/re-pe/output"
	"log"
	"os"
	"strconv"
)

const (
	BREAKPOINT     = "breakpoint"

	tmpDirName      = "_tmp/"
	logDirName      = tmpDirName + "log/"
	logFileName     = logDirName + "_settings.log"

	keyConfFileName = "ConfFile"
	keyConfDirName  = "ConfDir"

	confName        = "start.conf"
	xmlName         = "start.xml"
	breakNo         = 100
)

var (
	flagDebug   bool
	flagVerbose bool
)

func Out(selector string, args ...interface{}){
	output.Out(selector, args...)
}

func OutDeb(selector string, args ...interface{}){
	selector = "R.20." + selector + output.Env.GetPreColor(selector)
	//output.Out("lPr."+selector, args...)
	output.Out("fPr." + selector, args...)
}

func OutAll(selector string, args ...interface{}){
	output.Out("lPr." + selector, args...)
	output.Out("fPr." + selector, args...)
}

func OutLog(selector string, args ...interface{}){
	output.Out("lPr."+selector, args...)
}

func OutScr(selector string, args ...interface{}){
	output.Out("fPr."+selector, args...)
}

type Index []interface{} 
type Indices []Index


type FinalValue dtree.DTree

type FinalValueMap map[string]dtree.DTree

//func CollectFinalValues(indices *Indices, tree *dtree.DTreeHandler /* *dtree.JsonHandler*/, breakNo int) (finalValues FinalValueMap) {
func CollectFinalValues(indices *Indices, tree dtree.I_DTreeHandler, breakNo int) (finalValues FinalValueMap) {
	OutDeb("BeginLoop", "CollectFinalValues()")
//	tree := *treePtr
	OutDeb("Out", "indices:"); OutScr("R.20.?F:%v", indices)
	OutScr("B.00.Out.FgYellow.Bold", "Indices are:\n        no.           index\n-----------------------------------------------------------------------"); 
	finalValues = FinalValueMap{}
	for no, index := range *indices {
		if breakNo > -1 && no > breakNo {
			break
		}
		strKey := index[0].(string)
		OutScr("S.00.?F:%10d. %15s\n", no, strKey)

//_ = BREAKPOINT
		var result dtree.DTree
		switch len(index){
		case 1 :
			result = tree.Get(strKey)
		case 2 : 
			result = tree.Update(strKey, index[1])
		}
		if result.Error == nil {
			OutScr("S.22.?F:tree.Get(\"%s\") ==\n\n", strKey); OutScr("R.22.?F:%v\n", result.Value)
		} else {
			OutScr("S.22.?F:Error!!! Result of tree.Get(\"%s\") is ", strKey)
			OutScr("S.22.?F:%v ", result.Value); OutScr("S.22.?F:%v", result.RestPath)
			OutScr("R.22.?F:.%v", result.Error.Error())
		}
		OutScr("R.22")
		finalValues[strKey] = result
	}
	OutDeb("EndLoop", "CollectFinalValues()")
	return
}

func PrintFinalValues(indices *Indices, finalValueMap *FinalValueMap, breakNo int){
	OutDeb("B.00.Out.FgYellow.Bold", "Values for indices are:\n        no.            path = result.Value\n        no.            path = --  Wrong path : rest path  Error : error\n-----------------------------------------------------------------------")
	OutDeb("BeginLoop", "valLoop")
	finalValues := *finalValueMap
	for no, index := range *indices {
		if breakNo > -1 && no > breakNo {
			break
		}
		strKey := index[0].(string)
		if intKey, cnvErr := strconv.Atoi(strKey); cnvErr == nil {
			OutScr("S.00.?F:%10d. %15d = ", no, intKey)
		} else {
			OutScr("S.00.?F:%10d. %15s = ", no, "\"" + strKey + "\"")
		}
		if finalValues[strKey].Error == nil {
			OutScr("S.00.?F:%v", finalValues[strKey].Value)
		} else {
			OutScr("S.00.?F:--  ")
		}
		if(len(finalValues[strKey].RestPath) > 0){
			OutScr("S.00.?F:Wrong path : %-10s", finalValues[strKey].RestPath); OutScr("R.00.?F: Error : %v", finalValues[strKey].Error)
		} else {
			OutScr("R.00")
		}
	}
	OutDeb("EndLoop", "valLoop")
	return
}

func main(){
	flagDebug = false
	flagVerbose = false
	args := os.Args
	for _, b := range args {
        if b == "--debug" {
            flagDebug = true
        }
		if b == "--verbose" {
			flagVerbose = true
		}
    }
//_ = BREAKPOINT
	output.Env.Init("_tmp/log/_settings.log", flagDebug, flagVerbose, output.States)
	defer output.Env.LogFile.Close()
	log.SetOutput(output.Env.LogFile)
	OutScr("R.00", "")
	OutDeb("BeginProg", "_Settings")
	//var result.RestPath string
	//var value interface{}
	OutAll("B.00.FgCyan.Bold.?F: JSON test\n-----------")
	var appConf dtree.JsonHandler
	
/*  ltrenatyva - funkcijos priskyrimas	
	appConf := dtree.DTreeHandler{}
	appConf.Decode = func() (err error) {
		tree := &appConf
		if len((*tree).FileContent) == 0 {
			err = fmt.Errorf("DTreeHandler.FileContent is empty!")
			return
		}
		err = json.Unmarshal((*tree).FileContent, &(*tree).Value)
		return
	}
*/
	
	var err error
	if err = appConf.ReadFile(confName); err != nil {
		panic(err)
	}
	OutAll("R.00.?F: File \"%v\" exists and was read.\n", appConf.FileName)
	if err = appConf.Decode(); err != nil {
		panic(err)
	}
	OutAll("R.00.?F: A structure with data tree of \"%v\" file was successfully created.\n", appConf.FileName)
	OutAll("R.20.Out", "Content of the structure is:\n\n                           key : value\n---------------------------------------------------------\n" )
	
//_ = BREAKPOINT
  	switch typedValue := appConf.Value.(type){
	case map[string]interface{}:
		for i, u := range typedValue {
			OutScr("S.20.?F: %30s : ", i); OutScr("R.20", u)
		}
	case []interface{}:
		for i, u := range typedValue {
			OutScr("S.20.?F: %30d : ", i); OutScr("R.20", u)
		}
	default:
		OutScr("S.20.?F:%30s : ", " "); OutScr("R.20", typedValue)
	}
	
	OutAll("R.20", "\nEnd of reading of", confName, "\n")

	OutScr("L.00.FgYellow.Bold.?F:appConf.DirName is "); OutScr("R.00", appConf.DirName)
	OutScr("L.00.FgYellow.Bold.?F:appConf.FileName is "); OutScr("R.00", appConf.FileName)
	OutScr("B.00.FgYellow.Bold.?F:appConf.Value is:\n"); OutScr("R.00", appConf.Value, "\n")

	indices := Indices{
/*        ".",
		"a.1.v.f",
		"AppName.a", 
		"AppName.b", 
		"Database.DirName", 
		"Database.Config", 
		"Database.Source", 
		"Desktop.DirName", 
		"Desktop.Config", 
		"Desktop.Source",
		"Apps.SourceDir",
		"Apps.ConfPattern",
		"Apps.ConfFile",
		"Database.DBName",
		"Database.ConfDir",
		"Database.ConfFile",
		"phpDesktop.DefaultsDir",
		"phpDesktop.ConfDir",
		"phpDesktop.ConfFile",
*/

/**/	{"root"},
/**/	{"root.Map"},
/**/	{"root.Map.a"},
/**/	{"root.Map.b"},
		{"root.Map.c"},
		{"root.Map.0"},
/**/	{"root.Map.a.1"},
/**/	{"root.Map.s"},
		{"root.Map.s.15"},

/**/	{"root.Arr"},
/**/	{"root.Arr.0"},
/**/	{"root.Arr.1"},
/**/	{"root.Arr.2"},
/**/	{"root.Arr.3"},
/**/	{"root.Arr.4"},
/**/ 	{"root.Arr.5"},
/**/	{"root.Arr.0.0"},
/**/	{"root.Arr.6"},
/**/	{"root.Arr.6.11"},
/**/
	}
	breakNo := -1
	finalValues := CollectFinalValues(&indices, &appConf, breakNo)
	PrintFinalValues(&indices, &finalValues, breakNo)
	OutScr("R.00")
	
	indices = Indices{
		//{"Map.a", "new map value"},
		//{"Map.3", "new map value"},
		//{"Arr.3", []interface{}{nil, nil, "new array value"}},
		//{"Arr.4", "new array value"},
		//{"Arr.4.2", "new array value"},
		//{"", "Labas"},
		//{"Arr.n.m", "dddd"},
		//{"Map.n.m", "dddd"},
	}
	
	finalValues = FinalValueMap{}
	
//_ = BREAKPOINT
	finalValues = CollectFinalValues(&indices, &appConf, breakNo)
	PrintFinalValues(&indices, &finalValues, breakNo)

_ = BREAKPOINT
	path := "root.Map.n.m"; newValue := "dddd"
	OutAll("B.00", "path : ",  path, "newValue : ", newValue)
	result := appConf.Set(path, newValue)
	OutAll("B.00", result.RestPath, result.Value, result.Error)
	
	path = "root.Arr.+"; newValue = "0000"
	OutAll("B.00", "path : ",  path, "newValue : ", newValue)
	result = appConf.Set(path, newValue)
	OutAll("B.00", result.RestPath, result.Value, result.Error)

	path = "root.Arr.+.+"; newValue = "7777"
	OutAll("B.00", "path : ",  path, "newValue : ", newValue)
	result = appConf.Set(path, newValue)
	OutAll("B.00", result.RestPath, result.Value, result.Error)

	path = "root.Map.+"; newValue = "0000"
	OutAll("B.00", "path : ",  path, "newValue : ", newValue)
	result = appConf.Set(path, newValue)
	OutAll("B.00", result.RestPath, result.Value, result.Error)

	OutScr("B.00.Out.FgYellow.Bold.?F:appConf.Value is: "); OutScr("B.00.?F: %v", appConf.Value)
 
	OutAll("B.00.Out.FgCyan.Bold.?F: XML test\n-----------")
 
	var xmlConf dtree.XMLHandler
	if result.Error = xmlConf.ReadFile(xmlName); result.Error != nil {
		panic(result.Error)
	}
	OutAll("R.00.?F: File \"%v\" exists and was read.\n", xmlConf.FileName)
	
	if result.Error = xmlConf.Decode(); result.Error != nil {
		panic(result.Error)
	}
	OutAll("R.00.?F: A structure with data tree of \"%v\" file was successfully created.\n", xmlConf.FileName)
	OutScr("R.00.Out.FgYellow.Bold.?F:xmlConf.Value is: "); OutScr("B.00.?F: %v", xmlConf.Value)

 
	OutDeb("EndProg", "_Settings")
}
 