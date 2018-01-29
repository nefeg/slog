package slog

import (
	"testing"
	"fmt"
)

var errFormat = "Fail: expected \"%v\", but get \"%v\""

func getTestLogger(level Level, override bool) *SLogger{

	var TestLogger SLogger  = func(data ...interface{}) interface{} { return fmt.Sprint(data...) }
	var TestLevel           = level

	Wrap(&TestLogger, TestLevel, override)

	return &TestLogger
}


func TestWrapNew(t *testing.T){

	var found               = false
	var TestLogger SLogger  = func(data ...interface{}) interface{} { return fmt.Sprint(data...) }
	var TestLevel           = Level{"TestWrapNew", 2}

	Wrap(&TestLogger, TestLevel, false)

	for _,ref := range s.loggers[TestLevel]{
		if ref == &TestLogger{
			found = true
		}
	}

	if !found{
		t.Error("Wrap fail: no handlers added level")
	}
}

func TestWrapOverride(t *testing.T){

	var found               = false
	var TestLogger SLogger  = func(data ...interface{}) interface{} { return fmt.Sprint(data...) }
	var TestLevel           = Level{"TestWrapOverride", 3}

	Wrap(&TestLogger, TestLevel, true) // add handler
	Wrap(&TestLogger, TestLevel, true) // override handler

	if len(s.loggers[TestLevel]) >1{
		t.Error("Wrap fail: handlers not overrided (more then one handler found)")
	}

	for _,ref := range s.loggers[TestLevel]{
		if ref == &TestLogger{
			found = true
		}
	}

	if !found{
		t.Error("Wrap fail: no handlers found")
	}
}

func TestWrapAppend(t *testing.T){

	var found               = false
	var TestLogger SLogger  = func(data ...interface{}) interface{} { return fmt.Sprint(data...) }
	var TestLevel           = Level{"TestWrapAppend", 4}

	Wrap(&TestLogger, TestLevel, false)

	currentHndCount := len(s.loggers[TestLevel])

	Wrap(&TestLogger, TestLevel, false)


	if expected := currentHndCount +1; len(s.loggers[TestLevel]) != expected{
		t.Errorf("Wrap fail: expected handlers count: %v, but get: %v", expected, len(s.loggers[TestLevel]))
	}

	for _,ref := range s.loggers[TestLevel]{
		if ref == &TestLogger{
			found = true
		}
	}

	if !found{
		t.Error("Wrap fail: no handlers found")
	}
}


func TestMustLog(t *testing.T){

	var TestLevel           = Level{"test", 5}
	var TestLogger          = *getTestLogger(TestLevel, false)
	var testString          = "test string"

	// must log
	SetLevel(LvlAll)
	if r := TestLogger(testString); r == nil || r.(string) != testString {
		t.Errorf(errFormat, testString, r)
	}
}


func TestMustNotLog(t *testing.T){

	var TestLevel           = Level{"test", 5}
	var TestLogger          = *getTestLogger(TestLevel, false)
	var testString          = "test string"

	SetLevel(LvlDebug)
	if r := TestLogger(testString); r != nil{
		t.Errorf(errFormat, "nil", r)
	}

}

func TestMustNeverLog(t *testing.T){

	var TestLevel           = LvlAll
	var TestLogger          = *getTestLogger(TestLevel, false)
	var testString          = "test string"

	SetLevel(LvlNone)
	if r := TestLogger(testString); r != nil{
		t.Errorf(errFormat, "nil", r)
	}
}


func TestBind(t *testing.T){

	var bindFunc SLogger  = func(data ...interface{}) interface{} { return fmt.Sprint(data...) }
	var TestLogger SLogger
	var TestLevel           = Level{"TestBind", 16}
	var found = false

	Bind(&TestLogger, bindFunc, TestLevel, false)

	for _,ref := range s.loggers[TestLevel]{
		if ref == &TestLogger{
			found = true
		}
	}

	if !found{
		t.Error("Wrap fail: no function added to s.logger")
		t.Fail()
	}
}


func TestLogPanic(t *testing.T) {

	var panicMsg = "Testing \"logPanic\"... OK \n"

	defer func(){
		if r := recover(); r == nil{
			t.Error("The code did not panic while calling slog.Panic()")

		}else if r != panicMsg{
			t.Error("Unexpected message while calling slog.Panic()")
		}
	}()

	SetLevel(LvlNone)
	Panic(panicMsg)
}