package log

import "github.com/opentracing/opentracing-go/log"

var (
	// ErrorKind is the type or "kind" of an error (only for event="error" logs).
	// E.g., "Exception", "OSError"
	ErrorKind = stringLogName("error.kind")

	// ErrorObject for the actual error instance itself.
	ErrorObject = errorLogName("error.object")

	// Event is a stable identifier for some notable moment in the lifetime of a Span.
	// For instance, a mutex lock acquisition or release or the sorts of lifetime
	// events in a browser page load described in the Performance.timing specification.
	// E.g., from Zipkin, "cs", "sr", "ss", or "cr". Or, more generally, "initialized"
	// or "timed out". For errors, "error"
	Event = stringLogName("event")

	// Message a concise, human-readable, one-line message explaining the event.
	// E.g., "Could not connect to backend", "Cache invalidation succeeded"
	Message = stringLogName("message")

	// Stack a stack trace in platform-conventional format; may or may not pertain
	// to an error. E.g., "File \"example.py\", line 7, in \<module\>\ncaller()\nFile
	// \"example.py\", line 5, in caller\ncallee()\nFile \"example.py\", line 2, in
	// callee\nraise Exception(\"Yikes\")\n"
	Stack = stringLogName("stack")
)

func stringLogName(k string) func(string) log.Field {
	return func(v string) log.Field {
		return log.String(k, v)
	}
}

func errorLogName(k string) func(error) log.Field {
	return func(v error) log.Field {
		return log.String(k, v.Error())
	}
}
