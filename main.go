package main

// This file implements the CLIP main function

import (
	"flag"
	"fmt"
	"runtime"
)

// Command-line flags for special modes
// not normally used by the user.
var (
	//flag_complete *bool = flag.Bool