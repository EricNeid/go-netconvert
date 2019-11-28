package util

import "log"

// Error logs given error.
func Error(context string, err error) {
	log.Printf("[ERROR]: %s: %s\n", context, err.Error())
}

// Warn logs given message.
func Warn(context string, msg string) {
	log.Printf("[Warn]: %s: %s\n", context, msg)
}
