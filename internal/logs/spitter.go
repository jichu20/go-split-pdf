package logs

import (
	"fmt"
	"io"
	"time"
)

type Spitter struct {
	CreationDate time.Time
	Message      string
	Output       io.Writer
	LogLevel     Level
}

// func (s *Spitter) spitLog(l Level, a ...interface{}) {
func (s *Spitter) spitLog(l Level, a string) {

	if l <= s.LogLevel {
		fmt.Printf("%s - %s - %s\n",
			time.Now().Format("02-01-2006 15:04:05.000000"), l.String(), a)
	}
	// fmt.Fprintf(s.Output, "\033[0;31m %s - %s - %s\n",
	// 	time.Now().Format("02-01-2006 15:04:05.000000"), l.String(), a)
}
