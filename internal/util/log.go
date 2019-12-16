package util

import "log"

// Log provides logging methods. Provide the given context (file or package name for example)
// to be displayed in the log messages.
type Log struct {
	Context string
}

// E logs given error.
func (l *Log) E(method string, err error) {
	log.Printf("[ERROR]: %s: %s: %s\n", l.Context, method, err.Error())
}

// W logs given message.
func (l *Log) W(method string, msg string) {
	log.Printf("[Warn]: %s: %s: %s\n", l.Context, method, msg)
}
