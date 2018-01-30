package slog

import (
	"time"
	"fmt"
)

type Format func() string


var FormatDefault Format = func() string{

	return `%s`
}

var FormatTimed Format = func() string{

	return fmt.Sprintf(`%s %%s`, time.Now().Format("2006/01/02 15:04:05"))
}

var FormatTimed_RFC822 Format = func() string{

	return fmt.Sprintf(`[%s] %%s`, time.Now().Format(time.RFC822))
}
