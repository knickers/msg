package msg

import (
	"fmt"
)

var Verbose *bool

func Log(msg string, args ...interface{}) {
	if *Verbose {
		fmt.Printf(msg, args...)
	}
}
