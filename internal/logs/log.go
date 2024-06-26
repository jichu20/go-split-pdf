package logs

import "os"

var Default = NewLog()

type Log struct {
	*Spitter
}

func NewLog() *Log {
	return &Log{
		Spitter: &Spitter{
			Output: os.Stdout,
		},
	}
}
