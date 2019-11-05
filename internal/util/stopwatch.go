package util

import (
	"fmt"
	"time"
)

// TimeTrack prints the execution time of a specific function.
//
// Example:
// func myFunction() int {
//    defer timeTrack(time.Now(), "myFunction")
//    // ... do some things, maybe even return under some condition
//    return result
//}
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s", name, elapsed)
}
