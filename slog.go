package slog

/**
Methods:
	Debug(v ...interface{})
	DebugLn(v ...interface{})
	DebugF(format string, v ...interface{})

	Info(v ...interface{})
	InfoLn(v ...interface{})
	InfoF(format string, v ...interface{})

	Notice(v ...interface{})
	NoticeLn(v ...interface{})
	NoticeF(format string, v ...interface{})

	Warn(v ...interface{})
	WarnLn(v ...interface{})
	WarnF(format string, v ...interface{})

	Err(v ...interface{})
	ErrLn(v ...interface{})
	ErrF(format string, v ...interface{})

	Crit(v ...interface{})
	CritLn(v ...interface{})
	CritF(format string, v ...interface{})

Levels:
	LvlNone     -1
	LvlAll      0
	LvlDebug    10
	LvlInfo     20
	LvlNotice   30
	LvlWarn     40
	LvlError    50
	LvlCrit     60
 */

import (
	"fmt"
	"time"
)

var s *slog

var LvlNone     =  Level{"none",   -1}
var LvlAll      =  Level{"all",    0}
var LvlDebug    =  Level{"debug",  10}
var LvlInfo     =  Level{"info",   20}
var LvlNotice   =  Level{"notice", 30}
var LvlWarn     =  Level{"warn",   40}
var LvlError    =  Level{"err",    50}
var LvlCrit     =  Level{"crit",   60}

var Debug,  DebugLn,    DebugF  SLogger
var Info,   InfoLn,     InfoF   SLogger
var Notice, NoticeLn,   NoticeF SLogger
var Warn,   WarnLn,     WarnF   SLogger
var Err,    ErrLn,      ErrF    SLogger
var Crit,   CritLn,     CritF   SLogger

type slog struct{
	CurrentLevel    Level
	CurrentFormat   Format

	loggers         map[Level][]*SLogger
}

type Level struct{
	Name    string
	Value   int
}

type SLogger func(data ...interface{}) interface{}

type Format func() string


var FormatDefault = func() string{

	return `%s`
}

var FormatTimed Format = func() string{

	return fmt.Sprintf(`[%s] %%s`, time.Now().Format(time.RFC822))
}


func init(){

	s = &slog{}
	s.loggers = map[Level][]*SLogger{}

	SetLevel( LvlAll )
	SetFormat( FormatDefault )

	var stdLog, stdLogLn, stdLogF  SLogger
	stdLog      = func(data ...interface{}) interface{} { fmt.Printf(s.CurrentFormat(), fmt.Sprint(data...) ) ;return nil}
	stdLogLn    = func(data ...interface{}) interface{} { fmt.Printf(s.CurrentFormat() +"\n", fmt.Sprint(data...) ) ;return nil}
	stdLogF     = func(data ...interface{}) interface{} {
		fmt.Printf(
			fmt.Sprintf(s.CurrentFormat(), data[0].(string)),
			fmt.Sprint(data[1:]...),
		)
		return nil
	}


	Bind(&Debug,    stdLog,     LvlDebug,   false )
	Bind(&DebugLn,  stdLogLn,   LvlDebug,   false )
	Bind(&DebugF,   stdLogF,    LvlDebug,   false )

	Bind(&Info,     stdLog,     LvlInfo,    false )
	Bind(&InfoLn,   stdLogLn,   LvlInfo,    false )
	Bind(&InfoF,    stdLogF,    LvlInfo,    false )

	Bind(&Notice,   stdLog,     LvlNotice,  false )
	Bind(&NoticeLn, stdLogLn,   LvlNotice,  false )
	Bind(&NoticeF,  stdLogF,    LvlNotice,  false )

	Bind(&Warn,     stdLog,     LvlWarn,    false )
	Bind(&WarnLn,   stdLogLn,   LvlWarn,    false )
	Bind(&WarnF,    stdLogF,    LvlWarn,    false )

	Bind(&Err,      stdLog,     LvlError,   false )
	Bind(&ErrLn,    stdLogLn,   LvlError,   false )
	Bind(&ErrF,     stdLogF,    LvlError,   false )

	Bind(&Crit,     stdLog,     LvlCrit,    false )
	Bind(&CritLn,   stdLogLn,   LvlCrit,    false )
	Bind(&CritF,    stdLogF,    LvlCrit,    false )

	DebugF("[slog] init with level: %s\n", GetLevel().Name)
}

// Bind custom logger to log-level via reference
//
// var MyPanicLogger slog.SLogger
// slog.Bind(&MyPanicLogger, 6, func(data ...interface{}) {log.Panic(data)} )
func Bind(value *SLogger, fn SLogger, level Level, override bool ){

	*value = fn

	Wrap(value, level, override)
}



// Get custom logger to log-level
//
// var MyPanicLogger slog.SLogger = slog.GetSLogger(6, func(data ...interface{}) {log.Panic(data)} )
func Wrap(fn *SLogger, level Level, override bool  ) {

	if override && len(s.loggers[level]) >0{
		s.loggers[level] = []*SLogger{}
	}

	fnSource := *fn // for avoid circular reference errors
	*fn = func(data ...interface{}) interface{}{
		if s.CurrentLevel.Value >= 0 && s.CurrentLevel.Value <= level.Value{
			return fnSource(data...)
		}

		return nil
	}

	s.loggers[level] = append(s.loggers[level], fn)
}

func SetLevel(level Level){
	s.CurrentLevel = level
}

func GetLevel() Level{
	return s.CurrentLevel
}

func GetLevels() []Level{

	levels := make([]Level, 0, len(s.loggers))
	for l := range s.loggers {
		levels = append(levels, l)
	}

	return levels
}

func SetFormat(format Format){
	s.CurrentFormat = format
}

func GetFormat() Format{
	return s.CurrentFormat
}



