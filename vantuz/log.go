package vantuz

import "fmt"

type Logger struct {
	enabled bool
}

func (l *Logger) log(format string, a ...any) {
	if !l.enabled {
		return
	}
	var msg = fmt.Sprintf("[vantuz] "+format, a...)
	fmt.Println(msg)
}
