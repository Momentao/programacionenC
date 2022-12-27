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
// Simply passes the arguments to the daemon and
// displays the result.
func MainClient(args []string) {
	client := dialDaemon()
	var resp string
	err := client.Call("RPC.Call", args, &resp)
	if err != nil {
		fmt.Fprint(os.Stderr, cleanup(err.Error()))
	}
	fmt.Print(cleanup(resp))
}

// cleanup newlines so string can be printed to stdout without redundant/missing newlines
func cleanup(str string) string {
	str = strings.Trim(str, "\n")
	if str != "" {
		return str + "\n"
	}
	return str
}

// Connect to the clip daemon for RPC communication.
// Starts the daemon if he's not yet running.
func dialDaemon() *rpc.Client {
	// try to call the daemon
	client, err 