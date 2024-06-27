package logs

import "os"

const (
	TraceIdKey          = "X-RHO-TRACEID"
	ParentSpanIdKey     = "X-RHO-PARENTSPANID"
	GrpcTraceIdKey      = "rho-traceid"
	GrpcParentSpanIdKey = "rho-parentspanid"
)

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
