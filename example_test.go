package slog

import (
	"fmt"
)

func ExampleWrap(){

	var myCutomFunction SLogger  = func(data ...interface{}) interface{} { fmt.Print(data...); return nil }
	var myCutomLevel             = Level{"myCutomLevel", 5}

	Wrap(&myCutomFunction, myCutomLevel, false)

	SetLevel(LvlAll)

	myCutomFunction("some data for print")
	// Output: some data for print

}

func ExampleBind(){

	var myStdBindFun SLogger  = func(data ...interface{}) interface{} { fmt.Println(data...); return nil }


	var myCustomDebugFunc SLogger
	var myCustomDebugLevel   = Level{"myCustomDebugLevel", 5}

	Bind(&myCustomDebugFunc, myStdBindFun, myCustomDebugLevel, false)


	var myCustomNoticeFunc SLogger
	var myCustomNoticeLevel  = Level{"myCustomNoticeLevel", 15}

	Bind(&myCustomNoticeFunc, myStdBindFun, myCustomNoticeLevel, false)

	SetLevel(LvlAll)

	myCustomDebugFunc("some debug data")
	myCustomDebugFunc("some info data")

	// Output:
	// some debug data
	// some info data
}

func ExampleFormattedLog(){

	SetLevel(LvlInfo)
	SetFormat(FormatDefault)

	Infof("One string: %s, two string: %s, three int: %d", "1str", "2str", 3)

	// Output: One string: 1str, two string: 2str, three int: 3
}