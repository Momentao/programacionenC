package main

// This file implements the player state.
// The player communicates over RPC (Remote Procedure Call)
// with the client. The client forks a player daemon (clip -d)
// if none is yet running and sends RPC calls to it.

import (
//"sync"
)

type Player struct {
	Lib // the player's library
	//playlist ItemArray
	//	current  int // current track
	//	playing  bool
	//	backend  Backend
	port string // default RPC port
	//	sync.Mutex
}

// Constructor
func NewPlayer() *Player {
	p := new(Player)
	p.init()
	retur