
package main

// This file implements log and debug functions

import (
	"fmt"
	"os"
)

// Print debug message to stderr.
func Debug(msg ...interface{}) {