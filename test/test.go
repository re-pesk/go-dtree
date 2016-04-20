package main

import (
	. "github.com/re-pe/go/dtree"
	. "github.com/re-pe/go/output"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
)

const (
	//BREAKPOINT     = "breakpoint" // BREAKPOINT for modified godebug github.com/re-pe/godebug

	tmpDirName      = "_tmp/"
	logDirName      = tmpDirName + "log/"
	runFileName     = "test.go"
	logFileExt      = ".log"

	confName        = "start.conf"
	xmlName         = "start.xml"
	testLen         = -1
)

const (
	errMap    = "?:Map has no element with key \"%v\"!"
	errArr    = "?:Index \"%v\" is out of range of the array!"
	errVal    = "?:Value of type \"%s\" has no index!"
	errAlr    = "?:Element with key \"%s\" already exists!"
)

var (
	flags Flags
	logFileName string
)

type Value interface{}

type Expected struct{
	Value    Value
	UsedPath string
	RestPath string
	Error    error
}

type Test struct{
	Key      string
	NewValue Value
	Expected Expected
	Result   DTree
}

type TestList []Test

type Index []interface{} 
type Indices []Index


type FinalValue DTree

type FinalValueMap map[string]DTree

func JValue(jsonString string) (result Value){
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil && jsonString != "" {
		panic(err)
	}
	return
}

func Error(args ...interface{}) (err error){
	err = fmt.Errorf(Format(args...).Text)
	return
}

func OutputTest(testList *TestList, testLen int){
	Out(`FY.B?:
                 no. key = test.Key
				 
               new value = test.NewValue
			   
          expected value = test.Expected.Value
          returned Value = test.Result.Value
                                                            test summary 
      expected used path = test.Expected.UsedPath
      returned used path = test.Result.UsedPath
                                                            test summary 
      expected rest path = test.Expected.RestPath
      returned rest path = test.Result.RestPath
                                                            test summary 
          expected error = test.Expected.Error
          returned error = test.Result.Error
                                                            test summary 
-----------------------------------------------------------------------
`)
	for no, test := range (*testList)[:] {
		if testLen > -1 && no >= testLen {
			break
		}
		Out("FHY.B?:%24s = \"%v\"\n", Format("?:%3d. %s", no, "key").Text, test.Key)
		Out("\n")
		Out("?:%24s = %v\n"    , "new value"         , test.NewValue)
		Out("\n")
		Out("?:%24s = %v\n"    , "expected value"    , test.Expected.Value)
		Out("?:%24s = %v\n"    , "returned Value"    , test.Result.Value)
		if eq := reflect.DeepEqual(test.Expected.Value, test.Result.Value); eq {
			Out("FHC.B?: %80s \n", "Values are equal. Test is passed.")
		} else {
			Out("FHR.B?: %80s \n", "Values are not equal. Test is failed.")
		}
		Out("?:%24s = \"%v\"\n", "expected used path", test.Expected.UsedPath)
		Out("?:%24s = \"%v\"\n", "returned used path", test.Result.UsedPath)
		if eq := reflect.DeepEqual(test.Expected.UsedPath, test.Result.UsedPath); eq {
			Out("FHC.B?: %80s \n", "UsedPaths are equal. Test is passed.")
		} else {
			Out("FHR.B?: %80s \n", "UsedPaths are not equal. Test is failed.")
		}
		Out("?:%24s = \"%v\"\n", "expected rest path", test.Expected.RestPath)
		Out("?:%24s = \"%v\"\n", "returned rest path", test.Result.RestPath)
		if eq := reflect.DeepEqual(test.Expected.RestPath, test.Result.RestPath); eq {
			Out("FHC.B?: %80s \n", "RestPaths are equal. Test is passed.")
		} else {
			Out("FHR.B?: %80s \n", "RestPaths are not equal. Test is failed.")
		}
		Out("?:%24s = %v\n"    , "expected error"    , test.Expected.Error)
		Out("?:%24s = %v\n"    , "returned error"    , test.Result.Error)
		if eq := reflect.DeepEqual(test.Expected.Error, test.Result.Error); eq {
			Out("FHC.B?: %80s \n", "Errors are equal. Test is passed.")
		} else {
			Out("FHR.B?: %80s \n", "Errors are not equal. Test is failed.")
		}
		Out("\n")
	}
}

func TestGet(testList *TestList, tree I_DTreeHandler, testLen int) {
	Debug("FY.B?:Begin function", "TestGet()")
	for no, test := range (*testList)[:] {
		if testLen > -1 && no >= testLen {
			break
		}
_ = BREAKPOINT
		result := tree.Get(test.Key)
		(*testList)[no].Result = result
	}
	Out(`FY.B?:-----------------------------------------------------------------------
 Fields of testing results of DTree.Get():
`   )
	OutputTest(testList, testLen)
	Debug("End fnction", "TestGet()")
	return
}

func TestSet(testList *TestList, tree I_DTreeHandler, testLen int) {
	Debug("FY.B?:Begin function", "TestSet()")
	for no, test := range (*testList)[:] {
		if testLen > -1 && no >= testLen {
			break
		}
_ = BREAKPOINT
		result := tree.Set(test.Key, test.NewValue)
		(*testList)[no].Result = result
	}

	Out(`FY.B?:-----------------------------------------------------------------------
 Fields of testing results of DTree.Set():
`   )
	OutputTest(testList, testLen)
	Debug("End function", "TestSet()")
	return
}

/* func TestAdd(testList *TestList, tree I_DTreeHandler, testLen int) {
	Debug("FY.B?:Begin function", "TestAdd()")
	for no, test := range (*testList)[:] {
		if testLen > -1 && no >= testLen {
			break
		}
_ = BREAKPOINT
		result := tree.Add(test.Key, test.NewValue)
		(*testList)[no].Result = result
	}

	Out(`FY.B?:-----------------------------------------------------------------------
 Fields of testing results of DTree.Add():
`   )
	OutputTest(testList, testLen)
	Debug("End function", "TestAdd()")
	return
}
*/
 
func main(){
	args := os.Args
	for _, arg := range args {
        switch arg {
		case "--debug" :
            flags.Debug = true
		case "--verbose" :
			flags.Verbose = true
        }
    }

	OuputInit(&flags)
	
// logfailo sukūrimas
	logFile, err := NewLogFile(logDirName, runFileName, logFileExt)
	if err != nil { return }
	logFileInfo, err := logFile.Stat()
	if err != nil { return }
	logFileName = logFileInfo.Name()
	defer logFile.Close()
	log.SetOutput(logFile)

	Out("\n File ", runFileName, " is running.\n\n")

	Out("FHM.B?: JSON test\n-----------\n\n")

	// json'o skaitymo įrankio sukūrimas
	var appConf JsonHandler
	
	if err = appConf.ReadFile(confName); err != nil {
		panic(err)
	}
	Out("?: File \"%v\" exists and was read.\n\n", appConf.FileName)
	if err = appConf.Decode(); err != nil {
		panic(err)
	}
	Out("?: A structure with data tree of \"%v\" file was successfully created.\n\n", appConf.FileName)

	Out("FY.B?: appConf.Value is:\n");  Out(appConf.Value, "\n\n")

	// Preparing list of DTree.Get() tests
	// []Test{ 
	//     key, new value(nil), 
	//     Expected { expected value, expected used path, expected rest path, expected error },
	//     result
	// }
	
	//_ = BREAKPOINT 
	
	testList := TestList {
 		Test{"root"          , nil, Expected{
        JValue(`{"Map" : {"a" : 0, "b" : 1, "c" : 2, "i0" : 3},"Arr" : ["a", "b", "c", 0, 1, 2]}`),
    		                                                 "root"      , ""     , nil                    }, DTree{}},
 		Test{"root.Map"      , nil, Expected{
		    JValue(`{"a" : 0, "b" : 1, "c" : 2, "i0" : 3}`), "root.Map"  , ""     , nil                    }, DTree{}},
		Test{"root.Map.a"    , nil, Expected{JValue(`0`)   , "root.Map.a", ""     , nil                    }, DTree{}},
 		Test{"root.Map.b"    , nil, Expected{JValue(`1`)   , "root.Map.b", ""     , nil                    }, DTree{}},
		Test{"root.Map.c"    , nil, Expected{JValue(`2`)   , "root.Map.c", ""     , nil                    }, DTree{}},
		Test{"root.Map.0"    , nil, Expected{JValue(``)    , "root.Map"  , "0"    , Error(errMap, 0)       }, DTree{}},
		Test{"root.Map.a.1"  , nil, Expected{JValue(``)    , "root.Map.a", "1"    , Error(errVal,"float64")}, DTree{}},
		Test{"root.Map.s"    , nil, Expected{JValue(``)    , "root.Map"  , "s"    , Error(errMap,"s")      }, DTree{}},
		Test{"root.Map.s.1"  , nil, Expected{JValue(``)    , "root.Map"  , "s.1"  , Error(errMap,"s")      }, DTree{}},
		Test{"root.Arr"      , nil, Expected{
        JValue(`["a", "b", "c", 0, 1, 2]`)             , "root.Arr"  , ""     , nil                    }, DTree{}},
		Test{"root.Arr.0"    , nil, Expected{JValue(`"a"`) , "root.Arr.0", ""     , nil                    }, DTree{}},
		Test{"root.Arr.1"    , nil, Expected{JValue(`"b"`) , "root.Arr.1", ""     , nil                    }, DTree{}},
		Test{"root.Arr.2"    , nil, Expected{JValue(`"c"`) , "root.Arr.2", ""     , nil                    }, DTree{}},
		Test{"root.Arr.3"    , nil, Expected{JValue(`0`)   , "root.Arr.3", ""     , nil                    }, DTree{}},
		Test{"root.Arr.4"    , nil, Expected{JValue(`1`)   , "root.Arr.4", ""     , nil                    }, DTree{}},
		Test{"root.Arr.5"    , nil, Expected{JValue(`2`)   , "root.Arr.5", ""     , nil                    }, DTree{}},
		Test{"root.Arr.0.0"  , nil, Expected{JValue(``)    , "root.Arr.0", "0"    , Error(errVal, "string")}, DTree{}},
		Test{"root.Arr.6"    , nil, Expected{JValue(``)    , "root.Arr"  , "6"    , Error(errArr, 6)       }, DTree{}},
		Test{"root.Arr.6.11" , nil, Expected{JValue(``)    , "root.Arr"  , "6.11" , Error(errArr, 6)       }, DTree{}},
	}
	
	TestGet(&testList, &appConf, testLen)

	Out("\n")
	
	// Preparing list of DTree.Set() tests
	// []Test{ 
	//     key, new value(nil), 
	//     Expected { expected value, expected used path, expected rest path, expected error },
	//     result
	// }
	
  	testList = TestList {
		Test{"root.Map.n.m"  , JValue(`"dddd"`), Expected{JValue(`"dddd"`), "root.Map.n.m"  ,"", nil}, DTree{}},
		Test{"root.Arr.+"    , JValue(`"0000"`), Expected{JValue(`"0000"`), "root.Arr.+"    ,"", nil}, DTree{}},
		Test{"root.Arr.+.+"  , JValue( `7777` ), Expected{JValue( `7777` ), "root.Arr.+.+"  ,"", nil}, DTree{}},
		Test{"root.Map.+"    , JValue( `15`   ), Expected{JValue( `15`   ), "root.Map.+"    ,"", nil}, DTree{}},
	}
	
	TestSet(&testList, &appConf, testLen)

	Out("\n")

	// Preparing list of DTree.Add() tests
	// []Test{ 
	//     key, new value(nil), 
	//     Expected { expected value, expected used path, expected rest path, expected error },
	//     result
	// }

/*   	testList = TestList {
		Test{"root.Map.n.m"  , JValue(`"cccc"`), Expected{JValue(`"cccc"`), "root.Map.n.m"  ,"", nil}, DTree{}},
		Test{"root.Map.+"    , JValue( `1555` ), Expected{JValue( `1555` ), "root.Map.+"    ,"", nil}, DTree{}},
		Test{"root.Map.d"    , JValue( `3`    ), Expected{JValue( `3`    ), "root.Map.d"    ,"", nil}, DTree{}},
		Test{"root.Arr.+"    , JValue(`"117"` ), Expected{JValue(`"117"` ), "root.Arr.+"    ,"", nil}, DTree{}},
		Test{"root.Arr.0"    , JValue(`"----"`), Expected{nil             , "root.Arr.0"    ,"", 
		                                                                  Error(errAlr,"root.Arr.0")}, DTree{}},
		Test{"root.Arr.0.0"  , JValue(`"----"`), Expected{nil             , "root.Arr.0"    ,"0", 
		                                                                  Error(errAlr,"root.Arr.0")}, DTree{}},
		Test{"root.Arr.6.11" , JValue(`"0611"`), Expected{JValue(`"0611"`), "root.Arr.6.11" ,"", nil}, DTree{}},
	}

_ = BREAKPOINT

	TestAdd(&testList, &appConf, testLen)
 */
	Out("FY.B?:appConf.Value is: "); Out("?: %v\n\n", appConf.Value)
 
	Out("FHM.B?: XML test\n-----------\n\n")
 
	var xmlConf XMLHandler
	if err = xmlConf.ReadFile(xmlName); err != nil {
		panic(err)
	}
	Out("?: File \"%v\" exists and was read.\n\n", xmlConf.FileName)
	
	if err = xmlConf.Decode(); err != nil {
		panic(err)
	} 
	Out("?: A structure with data tree of \"%v\" file was successfully created.\n\n", xmlConf.FileName)
	Out("FY.B.?:xmlConf.Value is: "); Out("?: %v\n", xmlConf.Value)
 
 
	Debug("EndProg", runFileName)
}
 