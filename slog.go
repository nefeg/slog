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

	Fatal(v ...interface{})
	FatalLn(v ...interface{})
	FatalF(format string, v ...interface{})

	Panic(v ...interface{})
	PanicLn(v ...interface{})
	PanicF(format string, v ...interface{})

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
type slog struct{
	CurrentLevel    Level
	CurrentFormat   Format

	loggers         map[Level][]*SLogger
}

var s *slog

var Debug,  DebugLn,    DebugF  SLogger
var Info,   InfoLn,     InfoF   SLogger
var Notice, NoticeLn,   NoticeF SLogger
var Warn,   WarnLn,     WarnF   SLogger
var Err,    ErrLn,      ErrF    SLogger
var Crit,   CritLn,     CritF   SLogger
var Fatal,  FatalLn,    FatalF  SLogger
var Panic,  PanicLn,    PanicF  SLogger

func init(){

	s = &slog{}
	s.loggers = map[Level][]*SLogger{}

	SetLevel( LvlNone )
	SetFormat( FormatDefault )

	Bind(&Debug,    logStd,     LvlDebug,   false )
	Bind(&DebugLn,  logStdLn,   LvlDebug,   false )
	Bind(&DebugF,   logStdF,    LvlDebug,   false )

	Bind(&Info,     logStd,     LvlInfo,    false )
	Bind(&InfoLn,   logStdLn,   LvlInfo,    false )
	Bind(&InfoF,    logStdF,    LvlInfo,    false )

	Bind(&Notice,   logStd,     LvlNotice,  false )
	Bind(&NoticeLn, logStdLn,   LvlNotice,  false )
	Bind(&NoticeF,  logStdF,    LvlNotice,  false )

	Bind(&Warn,     logStd,     LvlWarn,    false )
	Bind(&WarnLn,   logStdLn,   LvlWarn,    false )
	Bind(&WarnF,    logStdF,    LvlWarn,    false )

	Bind(&Err,      logStd,     LvlError,   false )
	Bind(&ErrLn,    logStdLn,   LvlError,   false )
	Bind(&ErrF,     logStdF,    LvlError,   false )

	Bind(&Crit,     logStd,     LvlCrit,    false )
	Bind(&CritLn,   logStdLn,   LvlCrit,    false )
	Bind(&CritF,    logStdF,    LvlCrit,    false )

	Bind(&Panic,    logPanic,     LvlNone,    false )
	Bind(&PanicLn,  logPanicLn,   LvlNone,    false )
	Bind(&PanicF,   logPanicF,    LvlNone,    false )

	Bind(&Fatal,    logFatal,     LvlNone,    false )
	Bind(&FatalLn,  logFatalLn,   LvlNone,    false )
	Bind(&FatalF,   logFatalF,    LvlNone,    false )

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
		if (s.CurrentLevel.Value >= 0 && s.CurrentLevel.Value <= level.Value) || level.Value <0{
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



