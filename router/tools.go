package router

import (
	"fmt"

	"github.com/256dpi/gomqtt/packet"
)

type logWriter struct {
	w ResponseWriter
}

func (w *logWriter) Publish(msg *packet.Message) {
	fmt.Printf("Publishing: %s\n", msg)

	w.w.Publish(msg)
}

// Logger is a middleware that prints requests and published messages.
func Logger(next Handler) Handler {
	return func(w ResponseWriter, r *Request) {
		fmt.Printf("New Request: %s\n", r.Message)

		next(&logWriter{w}, r)
	}
}
