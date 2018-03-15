package slog

/**
Methods:
	Debug(v ...interface{})
	Debugln(v ...interface{})
	Debugf(format string, v ...interface{})

	Info(v ...interface{})
	Infoln(v ...interface{})
	Infof(format string, v ...interface{})

	Notice(v ...interface{})
	Noticeln(v ...interface{})
	Noticef(format string, v ...interface{})

	Warn(v ...interface{})
	Warnln(v ...interface{})
	Warnf(format string, v ...interface{})

	Err(v ...interface{})
	Errln(v ...interface{})
	Errf(format string, v ...interface{})

	Crit(v ...interface{})
	Critln(v ...interface{})
	Critf(format string, v ...interface{})

	Fatal(v ...interface{})
	Fatalln(v ...interface{})
	Fatalf(format string, v ...interface{})

	Panic(v ...interface{})
	Panicln(v ...interface{})
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

var Debug,  Debugln,    Debugf  SLogger
var Info,   Infoln,     Infof   SLogger
var Notice, Noticeln,   Noticef SLogger
var Warn,   Warnln,     Warnf   SLogger
var Err,    Errln,      Errf    SLogger
var Crit,   Critln,     Critf   SLogger
var Fatal,  Fatalln,    Fatalf  SLogger
var Panic,  Panicln,    Panicf  SLogger
var DebugPanic,  DebugPanicln,    DebugPanicf  SLogger


func init(){

	s = &slog{}
	s.loggers = map[Level][]*SLogger{}

	SetLevel( LvlNone )
	SetFormat( FormatDefault )

	Bind(&Debug,    logStd,     LvlDebug,   false )
	Bind(&Debugln,  logStdLn,   LvlDebug,   false )
	Bind(&Debugf,   logStdF,    LvlDebug,   false )

	Bind(&Info,     logStd,     LvlInfo,    false )
	Bind(&Infoln,   logStdLn,   LvlInfo,    false )
	Bind(&Infof,    logStdF,    LvlInfo,    false )

	Bind(&Notice,   logStd,     LvlNotice,  false )
	Bind(&Noticeln, logStdLn,   LvlNotice,  false )
	Bind(&Noticef,  logStdF,    LvlNotice,  false )

	Bind(&Warn,     logStd,     LvlWarn,    false )
	Bind(&Warnln,   logStdLn,   LvlWarn,    false )
	Bind(&Warnf,    logStdF,    LvlWarn,    false )

	Bind(&Err,      logStd,     LvlError,   false )
	Bind(&Errln,    logStdLn,   LvlError,   false )
	Bind(&Errf,     logStdF,    LvlError,   false )

	Bind(&Crit,     logStd,     LvlCrit,    false )
	Bind(&Critln,   logStdLn,   LvlCrit,    false )
	Bind(&Critf,    logStdF,    LvlCrit,    false )

	Bind(&Panic,    logPanic,     LvlNone,    false )
	Bind(&Panicln,  logPanicLn,   LvlNone,    false )
	Bind(&Panicf,   logPanicF,    LvlNone,    false )

	Bind(&Fatal,    logFatal,     LvlNone,    false )
	Bind(&Fatalln,  logFatalLn,   LvlNone,    false )
	Bind(&Fatalf,   logFatalF,    LvlNone,    false )

	Bind(&DebugPanic,    logPanic,     LvlDebug,   false )
	Bind(&DebugPanicln,  logPanicLn,   LvlDebug,   false )
	Bind(&DebugPanicf,   logPanicF,    LvlDebug,   false )

	Debugf("[slog] init with level: %s\n", GetLevel().Name)
}

// Bind custom logger to log-level via reference
//
// var MyPanicLogger slog.SLogger
// slog.Bind(&MyPanicLogger, 6, func(data ...interface{}) {log.Panic(data)} )
//
// !if level == LvlNone function will call always
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



