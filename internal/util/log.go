package util

import "log"

// Error logs given error.
func Error(context string, err error) {
	log.Printf("[ERROR]: %s: %s\n", context, err.Error())
}
