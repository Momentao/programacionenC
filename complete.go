
package main

// This file implements bash programmable completion.

import (
	"fmt"
	"strconv"
)

func init() {
	help["complete"] = `do bash autocompletion on arguments`
}