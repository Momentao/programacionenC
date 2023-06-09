
package main

// This file implements the Remote Procedure Call between
// the clip daemon and client front-end

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"reflect"
	"unicode"
)

// Start serving RPC calls from client instances.
func (player *Player) serveRPC() {
	rpc.Register(player.RPC())
	rpc.HandleHTTP()
	conn, err := net.Listen("tcp", port)
	if err != nil {
		Err("listen error:", err)
	}
	Debug("Listening on port " + port)
	http.Serve(conn, nil)
	//TODO: log errors.
}

// Aliased type to define RPC methods on.
type RPC struct {
	player *Player
}

// RPC-exported function used for normal operation mode.
// The command-line arguments are passed (e.g. "play jazz")
// and a response to the user is returned in *resp.
// Here, run-time reflection is used to match the user command
// to a method on the API type.
func (rpc RPC) Call(args []string, resp *string) (err error) {
	Debug("PlayerRPC.Call", args)

	if len(args) == 0 {
		args = []string{"help"}
	}
	cmd := args[0]  // first arg is command (e.g.: "play")
	args = args[1:] // rest are arguments (e.g.: "jazz")

	*resp, err = rpc.call(cmd, args)

	return
}

func (rpc RPC) call(cmd string, args []string) (resp string, err error) {
	// convert first character to uppercase
	first := unicode.ToUpper(rune(cmd[0]))
	Cmd := string(first) + cmd[1:] // (e.g.: Play)

	// resolve the command using reflection
	player := rpc.player
	p := reflect.ValueOf(player.API())
	m := p.MethodByName(Cmd)
	Debug("MethodByName", Cmd, ":", m)
	if m.Kind() == reflect.Invalid {
		err = errors.New("clip: '" + cmd + "' is not a clip command. See clip help.")
		return
	}

	// set up method arguments
	ins := m.Type().NumIn()
	var callArgs []reflect.Value
	switch ins {
	default:
		err = errors.New(fmt.Sprint("Bug: wrong number of ins: ", ins))
		return
	case 0:
		if len(args) > 0 {
			err = errors.New(fmt.Sprint(cmd, ": does not take arugments"))
			return
		}
		callArgs = []reflect.Value{}
	case 1:
		if len(args) == 0 {
			err = errors.New(fmt.Sprint(cmd, ": need an argument"))
			return
		}
		callArgs = []reflect.Value{reflect.ValueOf(args)}
	}

	// call the method
	r := m.Call(callArgs)
	resp = r[0].Interface().(string)    // by convention, response is 1st return value
	errStr := r[1].Interface().(string) // by convention, error is 2nd return value
	if errStr != "" {
		err = errors.New(errStr)
	}
	return
}