
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