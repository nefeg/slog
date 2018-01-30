# slog
[![GoDoc](https://godoc.org/github.com/umbrella-evgeny-nefedkin/slog)](https://godoc.org/github.com/umbrella-evgeny-nefedkin/slog)

Simple logger. Just import and use.

Install
-------
    go get github.com/umbrella-evgeny-nefedkin/slog

Usage
-----
Import the library and use:

    import "github.com/umbrella-evgeny-nefedkin/slog"
    
    func main(){
        ...
        slog.SetLevel(slog.LvlAll) // default LvlNone
        slog.SetFormat(slog.FormatTimmed)
        slog.Info("any info")
        ...
    }

Custom setting
--------------

  __Custom log format__
  
        // import "github.com/umbrella-evgeny-nefedkin/slog"
        
        var TimeFormatted slog.Format = func() string{
        
            return fmt.Sprintf(`[%s] %%s`, time.Now().Format(time.RFC822))
        }
        
        slog.SetFormat(TimeFormatted)
        
        slog.Info("any info") // Output: [23 Jan 18 21:55 UTC] any info
        
        
  *build-in formats*: 
  
        slog.FormatTimed
        slog.FormatTimed_RFC822
        slog.FormatDefault



  __Custom log function__
  
  *Using wrapper:*

    // import . "github.com/umbrella-evgeny-nefedkin/slog"

	var myCutomFunction SLogger  = func(data ...interface{}) interface{} { fmt.Sprint(data...); return nil }
	var myCutomLevel             = Level{"myCutomLevel", 5}

	Wrap(&myCutomFunction, myCutomLevel, false)
	
	SetLevel(LvlAll)
	
	myCutomFunction("some data for print")
	
	
  *Using binding:*

    // import . "github.com/umbrella-evgeny-nefedkin/slog"

	var myStdBindFun SLogger  = func(data ...interface{}) interface{} { fmt.Println(data...); return nil }


	var myCustomDebugFunc SLogger
	var myCustomDebugLevel   = Level{"myCustomDebugLevel", 5}
	Bind(&myCustomDebugFunc, myStdBindFun, myCustomDebugLevel, false)

	var myCustomNoticeFunc SLogger
	var myCustomNoticeLevel  = Level{"myCustomNoticeLevel", 15}
	Bind(&myCustomNoticeFunc, myStdBindFun, myCustomNoticeLevel, false)

	myCustomDebugFunc("some debug data")
	myCustomDebugFunc("some info data")
	

 *Function overriding:*

Just use override-flag:

	Wrap(&myCutomFunction, LvlInfo, TRUE)
	
	
or


    Bind(&myCustomNoticeFunc, myStdBindFun, LvlInfo, TRUE)
