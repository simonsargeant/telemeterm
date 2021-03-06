package display

import "time"

// SpanSnapshot allows unmarshalling the stdout exporter data. This is based on
// the output currently returned by the stdout exporter as the structs used by
// the exporter do not unmarshal back to the same struct. There are no guarantees
// that this structure won't change in the future.
//
// See https://github.com/open-telemetry/opentelemetry-go/issues/1819#issuecomment-825798804
type SpanSnapshot struct {
	SpanContext              SpanContext
	Parent                   SpanContext
	SpanKind                 int
	Name                     string
	StartTime                time.Time
	EndTime                  time.Time
	Attributes               []KeyValue
	MessageEvents            interface{} // ???
	Links                    interface{} // ???
	StatusCode               string
	StatusMessage            string
	DroppedAttributeCount    int
	DroppedMessageEventCount int
	DroppedLinkCount         int
	ChildSpanCount           int
	Resource                 []KeyValue
	InstrumentationLibrary   InstrumentationLibrary
}

type SpanContext struct {
	TraceID    string
	SpanID     string
	TraceFlags string
	TraceState interface{} // ???
	Remote     bool
}

type KeyValue struct {
	Key   string
	Value Value
}

type Value struct {
	Type  string
	Value interface{}
}

type InstrumentationLibrary struct {
	Name    string
	Version string
}
