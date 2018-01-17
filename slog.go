package slog

import "fmt"

var s *slog

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

	loggers         map[Level][]*SLogger
}

type Level struct{
	Name    string
	Value   int
}

type SLogger func(data ...interface{})


func init(){

	s = &slog{}
	s.loggers = map[Level][]*SLogger{}

	SetLevel( LvlCrit )

	stdLog      := func(data ...interface{}){ fmt.Print(data...) }
	stdLogLn    := func(data ...interface{}){ fmt.Println(data...) }
	stdLogF     := func(data ...interface{}){ fmt.Printf(data[0].(string), data[1:]...) }

	Bind(&Debug,    LvlDebug,   stdLog,     false )
	Bind(&DebugLn,  LvlDebug,   stdLogLn,   false )
	Bind(&DebugF,   LvlDebug,   stdLogF,    false )

	Bind(&Info,     LvlInfo,    stdLog,     false )
	Bind(&InfoLn,   LvlInfo,    stdLogLn,   false )
	Bind(&InfoF,    LvlInfo,    stdLogF,    false )

	Bind(&Notice,   LvlNotice,  stdLog,     false )
	Bind(&NoticeLn, LvlNotice,  stdLogLn,   false )
	Bind(&NoticeF,  LvlNotice,  stdLogF,    false )

	Bind(&Warn,     LvlWarn,    stdLog,     false )
	Bind(&WarnLn,   LvlWarn,    stdLogLn,   false )
	Bind(&WarnF,    LvlWarn,    stdLogF,    false )

	Bind(&Err,      LvlError,   stdLog,     false )
	Bind(&ErrLn,    LvlError,   stdLogLn,   false )
	Bind(&ErrF,     LvlError,   stdLogF,    false )

	Bind(&Crit,     LvlCrit,    stdLog,     false )
	Bind(&CritLn,   LvlCrit,    stdLogLn,   false )
	Bind(&CritF,    LvlCrit,    stdLogF,    false )

	DebugF("[slog] init with level: %s\n", GetLevel().Name)
}

// Bind custom logger to log-level via reference
//
// var MyPanicLogger slog.SLogger
// slog.Bind(&MyPanicLogger, 6, func(data ...interface{}) {log.Panic(data)} )
func Bind(value *SLogger, level Level, fn func(data ...interface{}), override bool ){

	*value = GetSLogger(level, fn, override)
}

// Get custom logger to log-level
//
// var MyPanicLogger slog.SLogger = slog.GetSLogger(6, func(data ...interface{}) {log.Panic(data)} )
func GetSLogger( level Level, fn func(data ...interface{}), override bool  ) (slogger SLogger){

	if override && len(s.loggers[level]) >0{
		s.loggers[level] = []*SLogger{}
	}

	slogger = func(data ...interface{}) {
		if s.CurrentLevel.Value <= level.Value{
			fn(data...)
		}
	}

	s.loggers[level] = append(s.loggers[level], &slogger)

	return slogger
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





