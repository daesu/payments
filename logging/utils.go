package logging

import (
	"fmt"
	"runtime"
	"strings"
)

// Trace returns the source code line and function name (of the calling function)
func Trace() (line string) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return fmt.Sprintf("%s,:%d %s\n", frame.File, frame.Line, frame.Function)
}

// StripSpecialChars strips newlines and tabs from a string
func StripSpecialChars(s string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case '\t', '\n':
			return ' '
		default:
			return r
		}
	}, s)
}
