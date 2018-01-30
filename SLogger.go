package slog

import (
	"fmt"
	"os"
)

type SLogger func(data ...interface{}) interface{}


var logStd      SLogger = func(data ...interface{}) interface{} { fmt.Printf(s.CurrentFormat(), fmt.Sprint(data...) ) ;return nil}
var logStdLn    SLogger = func(data ...interface{}) interface{} { fmt.Printf(s.CurrentFormat() +"\n", fmt.Sprint(data...) ) ;return nil}
var logStdF     SLogger = func(data ...interface{}) interface{} {
	f := fmt.Sprintf(s.CurrentFormat(), data[0].(string))
	fmt.Printf(
		f,
		data[1:]...,
	)
	return nil
}


// log to STDERR and make panic
var logPanic    SLogger = func(data ...interface{}) interface{} {
	e := fmt.Sprintf(s.CurrentFormat(), fmt.Sprint(data...) )

	fmt.Fprint(os.Stderr, e)

	panic(e)
	return nil
}

var logPanicLn  SLogger = func(data ...interface{}) interface{} {
	e := fmt.Sprintf(s.CurrentFormat() +"\n", fmt.Sprint(data...) )

	fmt.Fprint(os.Stderr, e)

	panic(e)
	return nil

}

var logPanicF   SLogger = func(data ...interface{}) interface{} {
	e := fmt.Sprintf(
		fmt.Sprintf(s.CurrentFormat(), data[0].(string)),
		data[1:]...,
	)

	fmt.Fprint(os.Stderr, e)

	panic(e)
	return nil
}


// log to STDERR and exit(1)
var logFatal    SLogger = func(data ...interface{}) interface{} {
	fmt.Fprintf(os.Stderr,s.CurrentFormat(), fmt.Sprint(data...) )
	os.Exit(1)
	return nil
}


var logFatalLn  SLogger = func(data ...interface{}) interface{} {
	fmt.Fprintf(os.Stderr, s.CurrentFormat() +"\n", fmt.Sprint(data...) )
	os.Exit(1)
	return nil
}

var logFatalF   SLogger = func(data ...interface{}) interface{} {
	fmt.Fprintf(os.Stderr,
		fmt.Sprintf(s.CurrentFormat() + data[0].(string)),
		data[1:]...,
	)
	os.Exit(1)
	return nil
}



