package msg

import (
	"fmt"
)

var Verbose *bool

// Uses fmt.Print
func Log(msgs ...interface{}) {
	if *Verbose {
		fmt.Print(msgs...)
	}
}

// Uses fmt.Printf
func Logf(msg string, args ...interface{}) {
	if *Verbose {
		fmt.Printf(msg, args...)
	}
}

// Uses fmt.Println
func Logln(msgs ...interface{}) {
	if *Verbose {
		fmt.Println(msgs...)
	}
}
