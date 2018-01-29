package slog

type Level struct{
	Name    string
	Value   int
}

var LvlNone     =  Level{"none",   -1}
var LvlAll      =  Level{"all",    0}
var LvlDebug    =  Level{"debug",  10}
var LvlInfo     =  Level{"info",   20}
var LvlNotice   =  Level{"notice", 30}
var LvlWarn     =  Level{"warn",   40}
var LvlError    =  Level{"err",    50}
var LvlCrit     =  Level{"crit",   60}