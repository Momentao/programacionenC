package main

// This file implements the client main function.
// Invoked whenever the user types executes "clip".
// The client merely forwards the CLI arguments
// to the clip daemon and returns the response to
// the user.

import (
	"fmt"
	"net/rpc"
	"os"
	"os/exec"
	"strings"
	