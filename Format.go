package slog

import (
	"time"
	"fmt"
)

type Format func() string


var FormatDefault = func() string{

	return `%s`
}

var FormatTimed Format = func() string{

	return fmt.Sprintf(`[%s] %%s`, time.Now().Format(time.RFC822))
}
