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
	"time"
)

// RPC port
const port = ":2527"

// Main loop for "client" mode (the normal mode).
// Simply passes the arguments to the da